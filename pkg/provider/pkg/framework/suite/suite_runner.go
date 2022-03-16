package suite

import (
	"fmt"
	"reflect"
	"runtime/debug"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/internal"
	"github.com/ozontech/allure-go/pkg/provider/pkg/common"
	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

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
	suiteName string
	suite     InternalSuite
	tests     map[string]*suiteTest
}

func NewSuiteRunner(t *testing.T, suiteName string, suite InternalSuite) runner.TestRunner {
	r := &suiteRunner{
		TestRunner: runner.NewRunner(t, suiteName),
		suiteName:  suiteName,
		suite:      suite,
		tests:      make(map[string]*suiteTest),
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
	common.BeforeAllHook(parentT, parentT.Provider())
	// wait for all BeforeAll's hooks over
	parentT.WG().Wait()
	for fullName, testData := range r.tests {
		suiteWG.Add(1)
		result[fullName] = parentT.RealT().Run(testData.testName, func(realT *testing.T) {
			var testT provider.InternalT
			newT := common.NewTestT(parentT, realT, testData.testName, testData.tags...)
			if it, ok := newT.(provider.InternalT); ok {
				testT = it
			} else {
				panic("T does not implements InternalT")
			}
			defer func() {
				suiteWG.Done()
				if !started {
					started = true
					suiteWG.Wait()
					common.AfterAllHook(parentT, parentT.Provider())
					// wait for all AfterAll's hooks over
					parentT.WG().Wait()
					r.FinishSuite()
				}
			}()

			// print test result
			defer func() {
				testT.Provider().FinishTest()
				testT.Provider().GetTestMeta().GetContainer().Finish()
				_ = testT.Provider().GetTestMeta().GetContainer().Print()
			}()

			defer func() {
				rec := recover()
				if rec != nil {
					ctxName := testT.Provider().ExecutionContext().GetName()
					errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, rec, debug.Stack())
					internal.TestError(errMsg, testT)
				}
			}()

			defer func() {
				common.AfterEachHook(testT, testT.Provider())
				// wait for all AfterEachHook's async steps over
				testT.WG().Wait()
			}()

			common.BeforeEachHook(testT, testT.Provider())
			// wait for all BeforeEachHook's async steps over
			testT.WG().Wait()

			testT.Provider().TestContext()
			testData.testBody.Func.Call([]reflect.Value{reflect.ValueOf(r.suite), reflect.ValueOf(testT)})

			// wait for all test's async steps over
			testT.WG().Wait()
		})
	}
	return result
}
