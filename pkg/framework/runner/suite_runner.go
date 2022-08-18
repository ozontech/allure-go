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
	r = collectTests(r, suite)
	r = collectHooks(r, suite)

	return r
}

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
			_, _ = fmt.Fprintf(os.Stderr, "allire-go: invalid regexp for -m: %s\n", err)
			os.Exit(1)
		}

		if !ok {
			continue
		}

		testMeta := adapter.NewTestMeta(suiteFullName, suiteName, method.Name, packageName)
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

type ParametrizedTest interface {
	GetRawBody() reflect.Method
	GetArgs() []reflect.Value
	GetMeta() provider.TestMeta
}

func parametrizedWrap(runner *suiteRunner, foo func(provider.T)) func(t provider.T) {
	return func(t provider.T) {
		foo(t)
		newTests := runner.tests
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
				}
			}
		}
		runner.tests = newTests
	}
}

func getParamTests(parentTest Test, params map[string]interface{}) map[string]Test {
	if paramTest, ok := parentTest.(ParametrizedTest); ok {
		var (
			suiteName   string
			packageName string
			tags        []string

			res           = make(map[string]Test)
			parentMeta    = paramTest.GetMeta()
			result        = parentMeta.GetResult()
			suiteFullName = result.FullName
		)
		if suites := result.GetLabel(allure.Suite); len(suites) > 0 {
			suiteName = suites[0].Value
		}
		if packages := result.GetLabel(allure.Package); len(packages) > 0 {
			packageName = packages[0].Value
		}
		for _, tag := range result.GetLabel(allure.Tag) {
			tags = append(tags, tag.Value)
		}

		for pName, param := range params {
			meta := adapter.NewTestMeta(fmt.Sprintf("%s/%s", suiteFullName, suiteName), parentMeta.GetResult().Name, pName, packageName, tags...)
			meta.GetResult().SetLabel(allure.ParentSuiteLabel(suiteName))
			res[pName] = &testMethod{
				testMeta: meta,
				testBody: paramTest.GetRawBody(),
				callArgs: append(paramTest.GetArgs(), reflect.ValueOf(param)),
			}
		}
		return res
	}
	panic("missing interface implementaion for passed test: ParametrizedTest")
}

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
		pName := fmt.Sprintf("%+v", param)
		res[pName] = param
	}
	return
}

func collectHooks(runner *suiteRunner, suite TestSuite) *suiteRunner {
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
