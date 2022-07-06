package runner

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/adapter"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/testplan"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// AllureBeforeTest has a BeforeEach method, which will run before each
// test in the suite.
type AllureBeforeTest interface {
	BeforeEach(t provider.T)
}

// AllureAfterTest has a AfterEach method, which will run after
// each test in the suite.
type AllureAfterTest interface {
	AfterEach(t provider.T)
}

// AllureBeforeSuite has a BeforeAll method, which will run before the
// tests in the suite are run.
type AllureBeforeSuite interface {
	BeforeAll(t provider.T)
}

// AllureAfterSuite has a AfterAll method, which will run after
// all the tests in the suite have been run.
type AllureAfterSuite interface {
	AfterAll(t provider.T)
}

type InternalSuite interface {
	GetRunner() TestRunner
	SetRunner(runner TestRunner)
}

type suiteRunner struct {
	*runner

	packageName string
	suiteName   string
	suite       InternalSuite
}

func NewSuiteRunnerWithParent(realT TestingT, packageName, suiteName, parentSuite string, suite InternalSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, parentSuite, suite)
}

func NewSuiteRunner(realT TestingT, packageName, suiteName string, suite InternalSuite) TestRunner {
	return newSuiteRunner(realT, packageName, suiteName, "", suite)
}

func newSuiteRunner(realT TestingT, packageName, suiteName, parentSuite string, suite InternalSuite) TestRunner {
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
		tests:     make(map[string]*test),
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

func collectTests(runner *suiteRunner, suite InternalSuite) *suiteRunner {
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
		runner.tests[method.Name] = &test{
			testMeta: testMeta,
			testBody: func(testT provider.T) {
				callArgs := []reflect.Value{
					reflect.ValueOf(suite),
					reflect.ValueOf(testT),
				}
				method.Func.Call(callArgs)
			},
		}
	}
	return runner
}

func collectHooks(runner *suiteRunner, suite InternalSuite) *suiteRunner {
	if beforeAll, ok := suite.(AllureBeforeSuite); ok {
		runner.BeforeAll(beforeAll.BeforeAll)
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

var matchMethod = flag.String("allure-go.m", "", "regular expression to select tests of the testify suite to run")

// Filtering method according to set regular expression
// specified command-line argument -m
func methodFilter(name string) (bool, error) {
	if ok, _ := regexp.MatchString("^Test", name); !ok {
		return false, nil
	}
	return regexp.MatchString(*matchMethod, name)
}
