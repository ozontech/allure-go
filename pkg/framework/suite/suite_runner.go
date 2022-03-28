package suite

import (
	"fmt"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type suiteTest struct {
	testName string
	testBody reflect.Method
	tags     []string
}

func newSuiteTest(testName string, testBody reflect.Method, tags ...string) *suiteTest {
	return &suiteTest{
		testName: testName,
		testBody: testBody,
		tags:     tags,
	}
}

type suiteRunner struct {
	runner.TestRunner
	packageName string
	suiteName   string
	suite       InternalSuite
	tests       map[string]*suiteTest
}

func NewSuiteRunner(t TestingT, packageName, suiteName string, suite InternalSuite) runner.TestRunner {
	r := &suiteRunner{
		TestRunner:  runner.NewRunner(t, suiteName),
		packageName: packageName,
		suiteName:   suiteName,
		suite:       suite,
		tests:       make(map[string]*suiteTest),
	}
	r = collectTests(r, suite)
	r = collectHooks(r, suite)

	return r
}

func (r *suiteRunner) AddTest(testName string, testBody reflect.Method, tags ...string) {
	name := fmt.Sprintf("%s/%s", r.T().Name(), testName)
	r.tests[name] = newSuiteTest(testName, testBody, tags...)
}

func (r *suiteRunner) RunTests() map[string]bool {
	var (
		started = false
		suiteWG = sync.WaitGroup{}
		result  = make(map[string]bool)
	)
	parentT := r.T()
	common.BeforeAllHook(parentT, parentT.GetProvider())
	// wait for all BeforeAll's hooks over
	parentT.WG().Wait()
	for fullName, testData := range r.tests {
		suiteWG.Add(1)
		result[fullName] = parentT.RealT().Run(testData.testName, func(realT *testing.T) {
			testT := common.NewTestT(realT, parentT.GetProvider(), parentT, r.packageName, testData.testName, testData.tags...)
			defer func() {
				suiteWG.Done()
				if !started {
					started = true
					suiteWG.Wait()
					common.AfterAllHook(parentT, parentT.GetProvider())
					// wait for all AfterAll's hooks over
					parentT.WG().Wait()
					r.FinishSuite()
				}
			}()

			// print test result
			defer func() {
				testT.GetProvider().FinishTest()
				testT.GetProvider().GetTestMeta().GetContainer().Finish()
				_ = testT.GetProvider().GetTestMeta().GetContainer().Print()
			}()

			defer func() {
				rec := recover()
				if rec != nil {
					ctxName := testT.GetProvider().ExecutionContext().GetName()
					errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, rec, debug.Stack())
					common.TestError(testT, testT.Provider, testT.Provider.ExecutionContext().GetName(), errMsg)
				}
			}()

			defer func() {
				common.AfterEachHook(testT, testT.GetProvider())
				// wait for all AfterEachHook's async steps over
				testT.WG().Wait()
			}()

			common.BeforeEachHook(testT, testT.GetProvider())
			// wait for all BeforeEachHook's async steps over
			testT.WG().Wait()

			testT.GetProvider().TestContext()
			testData.testBody.Func.Call([]reflect.Value{reflect.ValueOf(r.suite), reflect.ValueOf(testT)})

			// wait for all test's async steps over
			testT.WG().Wait()
		})
	}
	return result
}

func getPackage(depth int) string {
	pc, _, _, _ := runtime.Caller(depth)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash
	return funcName[:lastDot]
}
