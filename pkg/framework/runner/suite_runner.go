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
	return newSuiteRunner(realT, packageName, suiteName, parentSuite, suite, false)
}

func NewInitParamsSuiteRunnerWithParent(realT TestingT, packageName, suiteName, parentSuite string, suite TestSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, parentSuite, suite, true)
}

func NewSuiteRunner(realT TestingT, packageName, suiteName string, suite TestSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, "", suite, false)
}

func NewInitParamsSuiteRunner(realT TestingT, packageName, suiteName string, suite TestSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, "", suite, true)
}

func newSuiteRunner(realT TestingT, packageName, suiteName, parentSuite string, suite TestSuite, initParams bool) TestRunner {
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
	r = collectTests(r, suite)
	if initParams {
		r = collectParametrizedTestsBeforeHooks(r, suite)
	}
	r = collectHooks(r, suite, initParams)

	return r
}

// collectParametrizedTestsBeforeHooks executes InitTestParams function, finds test methods with tableTestPrefix,
// gets map with parameters, gets map with parameterized tests,
// replaces tests in runner with parameterized tests with results
func collectParametrizedTestsBeforeHooks(runner *suiteRunner, suite TestSuite) *suiteRunner {
	if initTestParamsSuit, ok := suite.(WithTestPramsSuite); ok {
		initTestParamsSuit.InitTestParams()
	}
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
				id, ok := suite.FindAllureID(tName)
				if ok {
					tResult.AddLabel(allure.IDAllureLabel(id))
				}
				newTests[tName] = body
				runner.internalT.GetProvider().GetSuiteMeta().GetContainer().AddChild(tResult.UUID)
			}
		}
	}
	runner.tests = newTests
	return runner
}

// collectTests filters suite methods according to set regular expression and
// adds filtered methods to tests of runner
func collectTests(runner *suiteRunner, suite TestSuite) *suiteRunner {
	var (
		methodFinder  = reflect.TypeOf(suite)
		packageName   = runner.packageName
		suiteName     = runner.internalT.GetProvider().GetSuiteMeta().GetSuiteName()
		suiteFullName = runner.internalT.GetProvider().GetSuiteMeta().GetSuiteFullName()
	)

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
		id, ok := suite.FindAllureID(method.Name)
		if ok {
			testMeta.GetResult().AddLabel(allure.IDAllureLabel(id))
		}
		runner.tests[method.Name] = &testMethod{
			testMeta: testMeta,
			testBody: method,
			callArgs: []reflect.Value{
				reflect.ValueOf(suite),
			},
		}
	}
	return runner
}

type parametrizedTest interface {
	GetRawBody() reflect.Method
	GetArgs() []reflect.Value
	GetMeta() provider.TestMeta
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

// parametrizedWrap executes beforeAll function, finds test methods with tableTestPrefix,
// gets map with parameters, gets map with parameterized tests,
// replaces tests in runner with parameterized tests with results
func parametrizedWrap(runner *suiteRunner, beforeAll func(provider.T)) func(t provider.T) {
	return func(t provider.T) {
		beforeAll(t)
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
					newTests[tName] = body
					runner.internalT.GetProvider().GetSuiteMeta().GetContainer().AddChild(body.GetMeta().GetResult().UUID)
				}
			}
		}
		runner.tests = newTests
	}
}

func collectHooks(runner *suiteRunner, suite TestSuite, initParams bool) *suiteRunner {
	if beforeAll, ok := suite.(AllureBeforeSuite); ok {
		if initParams {
			runner.BeforeAll(beforeAll.BeforeAll)
		} else {
			runner.BeforeAll(parametrizedWrap(runner, beforeAll.BeforeAll))
		}
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

	return runner
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
