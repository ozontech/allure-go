package runner

import (
	"fmt"
	"runtime/debug"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/internal"
	"github.com/ozontech/allure-go/pkg/provider/pkg/common"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type TestRunner interface {
	NewTest(testName string, testBody func(provider.T), tags ...string)
	BeforeEach(hookBody func(provider.T))
	AfterEach(hookBody func(provider.T))
	BeforeAll(hookBody func(provider.T))
	AfterAll(hookBody func(provider.T))
	RunTests() map[string]bool
	FinishSuite()

	T() provider.InternalT
}

type runner struct {
	internalT provider.InternalT
	tests     map[string]*test
}

type test struct {
	testName string
	testBody func(provider.T)
	tags     []string
}

func newTest(testName string, testBody func(provider.T), tags ...string) *test {
	return &test{
		testName: testName,
		testBody: testBody,
		tags:     tags,
	}
}

func NewRunner(realT *testing.T, suiteName string) TestRunner {
	newT := common.NewT(realT, suiteName)
	if it, ok := newT.(provider.InternalT); ok {
		return &runner{internalT: it, tests: make(map[string]*test)}
	}
	panic("T does not implement provider.InternalT interface")
}

func (r *runner) T() provider.InternalT {
	return r.internalT
}

func (r *runner) NewTest(testName string, testBody func(provider.T), tags ...string) {
	fullName := fmt.Sprintf("%s/%s", r.T().Name(), testName)
	r.tests[fullName] = newTest(testName, testBody, tags...)
}

func (r *runner) BeforeEach(hookBody func(provider.T)) {
	r.T().Provider().GetTestMeta().SetBeforeEach(hookBody)
}

func (r *runner) AfterEach(hookBody func(provider.T)) {
	r.T().Provider().GetTestMeta().SetAfterEach(hookBody)
}

func (r *runner) BeforeAll(hookBody func(provider.T)) {
	r.T().Provider().GetSuiteMeta().SetBeforeAll(hookBody)
}

func (r *runner) AfterAll(hookBody func(provider.T)) {
	r.T().Provider().GetSuiteMeta().SetAfterAll(hookBody)
}

func (r *runner) RunTests() map[string]bool {
	var (
		started = false
		suiteWG = sync.WaitGroup{}
		result  = make(map[string]bool)
	)
	common.BeforeAllHook(r.T(), r.T().Provider())
	// wait for all BeforeAll's hooks over
	r.T().WG().Wait()
	for fullName, testData := range r.tests {
		suiteWG.Add(1)
		result[fullName] = r.T().RealT().Run(testData.testName, func(realT *testing.T) {
			var testT provider.InternalT
			newT := common.NewTestT(r.T(), realT, testData.testName, testData.tags...)
			if it, ok := newT.(provider.InternalT); ok {
				testT = it
			} else {
				panic("T does not implement provider.InternalT interface")
			}
			defer func() {
				suiteWG.Done()
				if !started {
					started = true
					suiteWG.Wait()
					common.AfterAllHook(r.T(), r.T().Provider())
					// wait for all AfterAll's hooks over
					r.T().WG().Wait()
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
					internal.TestError(testT.Provider().ExecutionContext().GetName(), errMsg, testT)
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
			testData.testBody(testT)

			// wait for all test's async steps over
			testT.WG().Wait()
		})
	}
	return result
}

func (r *runner) FinishSuite() {
	r.T().Provider().GetSuiteMeta().GetContainer().Finish()
	_ = r.T().Provider().GetSuiteMeta().GetContainer().Print()
}

func Run(t *testing.T, testName string, testBody func(provider.T), tags ...string) bool {
	newT := common.NewT(t, t.Name())
	return newT.Run(testName, testBody, tags...)
}
