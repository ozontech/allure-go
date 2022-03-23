package runner

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type InternalT interface {
	provider.T

	GetProvider() provider.Provider
	WG() *sync.WaitGroup
	GetResult() *allure.Result
}

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type TestRunner interface {
	NewTest(testName string, testBody func(provider.T), tags ...string)
	BeforeEach(hookBody func(provider.T))
	AfterEach(hookBody func(provider.T))
	BeforeAll(hookBody func(provider.T))
	AfterAll(hookBody func(provider.T))
	RunTests() map[string]bool
	FinishSuite()

	T() InternalT
}

type runner struct {
	internalT InternalT
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

func NewRunner(realT TestingT, suiteName string) TestRunner {
	newT := common.NewT(realT, getPackage(2), suiteName)
	return &runner{internalT: newT, tests: make(map[string]*test)}
}

func (r *runner) T() InternalT {
	return r.internalT
}

func (r *runner) NewTest(testName string, testBody func(provider.T), tags ...string) {
	fullName := fmt.Sprintf("%s/%s", r.T().Name(), testName)
	r.tests[fullName] = newTest(testName, testBody, tags...)
}

func (r *runner) BeforeEach(hookBody func(provider.T)) {
	r.internalT.GetProvider().GetTestMeta().SetBeforeEach(hookBody)
}

func (r *runner) AfterEach(hookBody func(provider.T)) {
	r.internalT.GetProvider().GetTestMeta().SetAfterEach(hookBody)
}

func (r *runner) BeforeAll(hookBody func(provider.T)) {
	r.internalT.GetProvider().GetSuiteMeta().SetBeforeAll(hookBody)
}

func (r *runner) AfterAll(hookBody func(provider.T)) {
	r.internalT.GetProvider().GetSuiteMeta().SetAfterAll(hookBody)
}

func (r *runner) RunTests() map[string]bool {
	var (
		started = false
		suiteWG = sync.WaitGroup{}
		result  = make(map[string]bool)
	)
	common.BeforeAllHook(r.T(), r.internalT.GetProvider())
	// wait for all BeforeAll's hooks over
	r.internalT.WG().Wait()
	for fullName, testData := range r.tests {
		suiteWG.Add(1)
		result[fullName] = r.T().RealT().Run(testData.testName, func(realT *testing.T) {
			testT := common.NewTestT(realT, r.internalT.GetProvider(), r.internalT, r.internalT.GetProvider().GetSuiteMeta().GetPackageName(), testData.testName, testData.tags...)
			defer func() {
				suiteWG.Done()
				if !started {
					started = true
					suiteWG.Wait()
					common.AfterAllHook(r.T(), r.internalT.GetProvider())
					// wait for all AfterAll's hooks over
					r.internalT.WG().Wait()
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
					common.TestError(testT, testT.Provider, errMsg)
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
			testData.testBody(testT)

			// wait for all test's async steps over
			testT.WG().Wait()
		})
	}
	return result
}

func (r *runner) FinishSuite() {
	r.internalT.GetProvider().GetSuiteMeta().GetContainer().Finish()
	_ = r.internalT.GetProvider().GetSuiteMeta().GetContainer().Print()
}

func Run(t *testing.T, testName string, testBody func(provider.T), tags ...string) bool {
	newT := common.NewT(t, getPackage(2), t.Name())
	return newT.Run(testName, testBody, tags...)
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
