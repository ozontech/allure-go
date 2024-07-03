package runner

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"unsafe"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/adapter"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/testplan"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type suiteRunner struct {
	*runner

	packageName string
	suiteName   string
	suite       TestSuite
}

func NewSuiteRunnerWithParent(realT TestingT, packageName, suiteName, parentSuite string, suite TestSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, parentSuite, suite)
}

func NewSuiteRunner(realT TestingT, packageName, suiteName string, suite TestSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, "", suite)
}

func newSuiteRunner(realT TestingT, packageName, suiteName, parentSuite string, suite TestSuite) TestRunner {
	newT := common.NewT(realT)

	callers := strings.Split(realT.Name(), "/")
	fullName := fmt.Sprintf("%s/%s", realT.Name(), suiteName)
	providerCfg := manager.NewProviderConfig().
		WithFullName(fullName).
		WithPackageName(packageName).
		WithSuiteName(suiteName).
		WithParentSuite(parentSuite).
		WithRunner(callers[0])
	newT.SetProvider(manager.NewProvider(providerCfg))
	testPlan := testplan.GetTestPlan()
	if testPlan != nil {
		fmt.Printf("Test plan found. It will be used for test filters\n")
	}

	testRunner := &runner{
		internalT: newT,
		testPlan:  testPlan,
		tests:     make(map[string]Test),
	}
	r := &suiteRunner{
		runner:      testRunner,
		packageName: packageName,
		suiteName:   suiteName,
		suite:       suite,
	}
	collectTests(r, suite)
	collectParametrizedTests(r, suite)
	collectHooks(r, suite)

	return r
}

// collectTests filters suite methods according to set regular expression and
// adds filtered methods to tests of runner
func collectTests(runner *suiteRunner, tSuite TestSuite) {
	var (
		methodFinder  = reflect.TypeOf(tSuite)
		packageName   = runner.packageName
		suiteName     = runner.internalT.GetProvider().GetSuiteMeta().GetSuiteName()
		suiteFullName = runner.internalT.GetProvider().GetSuiteMeta().GetSuiteFullName()
		getAllureID   func(string) string
	)

	if ais, ok := tSuite.(AllureIDSuite); ok {
		getAllureID = ais.GetAllureID
	}

	for i := 0; i < methodFinder.NumMethod(); i++ {
		method := methodFinder.Method(i)
		ok, err := methodFilter(method.Name)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "allure-go: invalid regexp for -m: %s\n", err)
			os.Exit(1)
		}

		if !ok {
			continue
		}

		testMeta := adapter.NewTestMeta(suiteFullName, suiteName, method.Name, packageName)

		if getAllureID != nil {
			id := getAllureID(method.Name)
			if len(id) > 0 {
				testMeta.GetResult().AddLabel(allure.IDAllureLabel(id))
			}
		}

		runner.tests[method.Name] = &testMethod{
			testMeta: testMeta,
			testBody: method,
			callArgs: []reflect.Value{
				reflect.ValueOf(tSuite),
			},
		}
	}
}

type parametrizedTest interface {
	GetRawBody() reflect.Method
	GetArgs() []reflect.Value
	GetMeta() provider.TestMeta
}

// parametrizedWrap executes beforeAll function, finds parametrized tests in the suite
func parametrizedWrap(runner *suiteRunner, beforeAll func(provider.T)) func(t provider.T) {
	return func(t provider.T) {
		beforeAll(t)
		initializeParametrizedTests(runner)
	}
}

// collectParametrizedTests finds test methods with tableTestPrefix,
// gets map with parameters, gets map with parameterized tests,
// replaces tests in runner with parameterized tests with results
func collectParametrizedTests(runner *suiteRunner, suite TestSuite) {
	if ps, ok := suite.(ParametrizedSuite); ok {
		ps.InitializeTestsParams()
	} else {
		return
	}
	initializeParametrizedTests(runner)
}

func initializeParametrizedTests(runner *suiteRunner) {
	newTests := make(map[string]Test)
	for k, v := range runner.tests {
		newTests[k] = v
	}
	for name, test := range runner.tests {
		if strings.HasPrefix(name, tableTestPrefix) {
			params, err := getParams(runner.suite, name)
			if err != nil {
				panic(err)
			}
			temp := getParamTests(test, params)
			delete(newTests, name)
			for tName, body := range temp {
				tResult := body.GetMeta().GetResult()
				newTests[tName] = body
				runner.internalT.GetProvider().GetSuiteMeta().GetContainer().AddChild(tResult.UUID)
			}
		}
	}
	runner.tests = newTests
}

// getParamTests create instance of TestAdapter for every param from params
// and returns map whose elements are a pair (<param name>, <pointer to instance of testMethod>)
func getParamTests(parentTest Test, params map[string]interface{}) map[string]Test {
	if paramTest, ok := parentTest.(parametrizedTest); ok {
		var (
			suiteName   string
			packageName string
			tags        []string

			res           = make(map[string]Test)
			parentMeta    = paramTest.GetMeta()
			result        = parentMeta.GetResult()
			suiteFullName = result.FullName
		)
		if suite, ok := result.GetFirstLabel(allure.Suite); ok {
			suiteName = suite.GetValue()
		}

		if _package, ok := result.GetFirstLabel(allure.Package); ok {
			packageName = _package.GetValue()
		}
		for _, tag := range result.GetLabels(allure.Tag) {
			tags = append(tags, tag.GetValue())
		}

		for pName, param := range params {
			meta := adapter.NewTestMeta(suiteFullName, suiteName, pName, packageName, tags...)
			if parentSuite, ok := result.GetFirstLabel(allure.ParentSuite); ok {
				meta.GetResult().ReplaceLabel(parentSuite)
			}

			if ptp, ok := param.(ParametrizedTestParam); ok {
				meta.GetResult().Name = ptp.GetAllureTitle()
				meta.GetResult().AddLabel(allure.IDAllureLabel(ptp.GetAllureID()))
			}

			res[pName] = &testMethod{
				testMeta: meta,
				testBody: paramTest.GetRawBody(),
				callArgs: append(paramTest.GetArgs(), reflect.ValueOf(param)),
			}
		}
		return res
	}
	panic(fmt.Sprintf("missing interface implementaion (parametrizedTest) for test: %s", parentTest.GetMeta().GetResult().Name))
}

// getParams checks that the parameter extending the suite is of the slice type
// and returns a map whose elements are a pair
// (<method name without tableTestPrefix>_<value of slice element>, <value of slice element>)
func getParams(suite TestSuite, methodName string) (res map[string]interface{}, err error) {
	var (
		structSuite = reflect.ValueOf(suite).Elem()
		paramName   = strings.TrimPrefix(methodName, tableTestPrefix)
	)
	res = make(map[string]interface{})
	params := structSuite.FieldByName(tableParamPrefix + paramName)
	if params.Kind() != reflect.Slice {
		err = fmt.Errorf("cannot find appropriate params for %s", methodName)
		return
	}
	for i := 0; i < params.Len(); i++ {
		paramV := params.Index(i)
		param := reflect.NewAt(paramV.Type(), unsafe.Pointer(paramV.UnsafeAddr())).Elem().Interface()
		pName := fmt.Sprintf("%s_%+v", paramName, param)
		if len(pName) > 150 { // workaround for t.TempDir()
			pName = pName[:150]
		}
		res[pName] = param
	}
	return
}

func collectHooks(runner *suiteRunner, suite TestSuite) {
	if beforeAll, ok := suite.(AllureBeforeSuite); ok {
		runner.BeforeAll(parametrizedWrap(runner, beforeAll.BeforeAll))
	}

	if beforeEach, ok := suite.(AllureBeforeTest); ok {
		runner.BeforeEach(beforeEach.BeforeEach)
	}

	if afterAll, ok := suite.(AllureAfterSuite); ok {
		runner.AfterAll(afterAll.AfterAll)
	}

	if afterEach, ok := suite.(AllureAfterTest); ok {
		runner.AfterEach(afterEach.AfterEach)
	}

}

var matchMethod = flag.String("allure-go.m", "", "regular expression to select tests of the allure-go suite to run")

// Filtering method according to set regular expression
// specified command-line argument -m
func methodFilter(name string) (bool, error) {
	var (
		validPrefixes = strings.Join([]string{testPrefix, tableTestPrefix}, "|")
		regFilter     = fmt.Sprintf("^(%s)", validPrefixes)
	)

	if ok, _ := regexp.MatchString(regFilter, name); !ok {
		return false, nil
	}
	return regexp.MatchString(*matchMethod, name)
}
