package runner

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/adapter"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/testplan"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type runner struct {
	internalT internalT
	testPlan  *testplan.TestPlan
	tests     map[string]Test
}

func NewRunner(realT TestingT, suiteName string) TestRunner {
	newT := common.NewT(realT)

	callers := strings.Split(realT.Name(), "/")
	providerCfg := manager.NewProviderConfig().
		WithFullName(realT.Name()).
		WithPackageName(getPackage(defaultPackageDepth)).
		WithSuiteName(suiteName).
		WithRunner(callers[0])
	newT.SetProvider(manager.NewProvider(providerCfg))

	testPlan := testplan.GetTestPlan()
	return &runner{internalT: newT, tests: make(map[string]Test), testPlan: testPlan}
}

func (r *runner) t() internalT {
	return r.internalT
}

func (r *runner) realT() TestingT {
	return r.t().RealT()
}

func (r *runner) toRun(result *allure.Result) bool {
	if r.testPlan != nil {
		return r.testPlan.IsSelected(result.TestCaseID, result.FullName)
	}
	return true
}

func (r *runner) filterByTestPlan() map[string]Test {
	if plan := r.testPlan; plan != nil {
		tests := make(map[string]Test)
		for fullName, testData := range r.tests {
			if r.testPlan.IsSelected(testData.GetMeta().GetResult().TestCaseID, testData.GetMeta().GetResult().FullName) {
				tests[fullName] = testData
			}
		}
		return tests
	}
	return r.tests
}

func (r *runner) NewTest(testName string, testBody func(provider.T), tags ...string) {
	fullName := fmt.Sprintf("%s/%s", r.t().Name(), testName)

	testMeta := adapter.NewTestMeta(
		r.t().GetProvider().GetSuiteMeta().GetSuiteFullName(),
		r.t().GetProvider().GetSuiteMeta().GetSuiteName(),
		testName,
		getPackage(defaultPackageDepth),
		tags...,
	)

	if !r.toRun(testMeta.GetResult()) {
		return
	}

	r.tests[fullName] = newTestFunc(testBody, testMeta)
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

func (r *runner) RunTests() SuiteResult {
	var (
		wg = &sync.WaitGroup{}

		parentSuiteMeta = r.t().GetProvider().GetSuiteMeta()
		parentTestMeta  = r.t().GetProvider().GetTestMeta()

		result         = NewSuiteResult(parentSuiteMeta.GetContainer())
		beforeAllHook  = common.CarriedHook(common.BeforeAll, parentSuiteMeta.GetBeforeAll)
		afterAllHook   = common.CarriedHook(common.AfterAll, parentSuiteMeta.GetAfterAll)
		beforeEachHook = common.CarriedHook(common.BeforeEach, parentTestMeta.GetBeforeEach)
		afterEachHook  = common.CarriedHook(common.AfterEach, parentTestMeta.GetAfterEach)
	)

	r.realT().Run(parentSuiteMeta.GetSuiteName(), func(t *testing.T) {
		oldParentT := r.realT()
		r.t().SetRealT(t)
		defer r.t().SetRealT(oldParentT)

		r.tests = r.filterByTestPlan()

		if len(r.tests) == 0 {
			r.t().Skipf("No tests to run for suite %s", r.t().Name())
			return
		}

		defer func() {
			wg.Wait()
			finishSuite(r.internalT.GetProvider())
		}()

		// after all hook
		defer func() {
			wg.Wait()
			_, _ = runHook(r.t(), afterAllHook)
		}()

		for _, test := range r.tests {
			result.GetContainer().AddChild(test.GetMeta().GetResult().UUID)
		}

		// before all hook
		ok, err := runHook(r.t(), beforeAllHook)
		if err != nil {
			for _, test := range r.tests {
				result = setupErrorHandler(fmt.Sprintf("%v setup was failed", r.t().Name()), err, test.GetMeta(), result)
			}
			return
		}
		if !ok {
			for _, test := range r.tests {
				result = setupErrorHandler(fmt.Sprintf("%v setup was failed", r.t().Name()), fmt.Errorf("something goes wrong in beforeAll"), test.GetMeta(), result)
			}
			return
		}

		// THE MOST dirty hack in history
		// t.Parallel() waits for parent-test reach its defer function
		// Unfortunately it's impossible to reach this function if parent-test waits for other tests complete
		// So if we run child test from test-runner
		// tests from suite will wait defer func of test-runner child instead of test-runner itself
		r.realT().Run("Tests", func(t *testing.T) {
			oldTestT := r.internalT.RealT()
			r.t().SetRealT(t)
			defer r.t().SetRealT(oldTestT)

			for _, testData := range r.tests {
				test := testData
				wg.Add(1)
				r.realT().Run(test.GetMeta().GetResult().Begin().Name, func(t *testing.T) {
					defer wg.Done()
					defer func() {
						result.NewResult(finishTest(t, test.GetMeta()))
					}()
					testT := setupTest(t, r.t().GetProvider(), test.GetMeta())

					// after each hook
					defer func() {
						_, _ = runHook(testT, afterEachHook)
					}()

					// catch panic in test body context
					defer func() {
						rec := recover()
						if rec != nil {
							ctxName := testT.GetProvider().ExecutionContext().GetName()
							errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, rec, debug.Stack())
							common.TestError(testT, testT.GetProvider(), testT.GetProvider().ExecutionContext().GetName(), errMsg)
						}
					}()

					// before each hook
					ok, err = runHook(testT, beforeEachHook)
					if err != nil {
						result = setupErrorHandler("Test Setup failed", err, test.GetMeta(), result)
						return
					}
					if !ok {
						result = setupErrorHandler("Test Setup failed", fmt.Errorf("assertion error due test setup"), test.GetMeta(), result)
						return
					}

					testT.GetProvider().TestContext()
					defer testT.WG().Wait()
					test.GetBody()(testT)
				})
			}
		})
	})
	return result
}

func Run(t *testing.T, testName string, testBody func(provider.T), tags ...string) *allure.Result {
	var (
		newT        = common.NewT(t)
		callers     = strings.Split(t.Name(), "/")
		providerCfg = manager.NewProviderConfig().
				WithFullName(t.Name()).
				WithPackageName(getPackage(2)).
				WithSuiteName(t.Name()).
				WithRunner(callers[0])
		newProvider = manager.NewProvider(providerCfg)
	)
	newT.SetProvider(newProvider)
	newT.Provider.TestContext()

	return newT.Run(testName, testBody, tags...)
}

func setupTest(t TestingT, parentProvider provider.Provider, meta provider.TestMeta) *common.Common {
	var (
		testT = common.NewT(t)

		parentSuiteMeta = parentProvider.GetSuiteMeta()
		parentTestMeta  = parentProvider.GetTestMeta()

		packageName     = parentSuiteMeta.GetPackageName()
		suiteName       = parentSuiteMeta.GetSuiteName()
		parentSuiteName = parentSuiteMeta.GetParentSuite()

		callers = strings.Split(t.Name(), "/")
		cfg     = manager.NewProviderConfig().
			WithFullName(t.Name()).
			WithPackageName(packageName).
			WithSuiteName(suiteName).
			WithParentSuite(parentSuiteName).
			WithRunner(callers[0])
	)
	testT.SetProvider(manager.NewProvider(cfg))

	testT.Provider.TestContext()
	meta.SetBeforeEach(parentTestMeta.GetBeforeEach())
	meta.SetAfterEach(parentTestMeta.GetAfterEach())
	if parentSuite := testT.Provider.GetSuiteMeta().GetParentSuite(); parentSuite != "" {
		meta.GetResult().WithParentSuite(parentSuite)
	}
	meta.SetResult(copyLabels(parentProvider.GetResult(), meta.GetResult()))
	testT.Provider.SetTestMeta(meta)

	return testT
}

func finishTest(t TestingT, meta provider.TestMeta) TestResult {
	testRes := NewTestResult(meta.GetResult(), meta.GetContainer())
	defer func() {
		if err := testRes.Print(); err != nil {
			t.Error(err.Error())
		}
	}()
	return testRes
}

func finishSuite(p provider.Provider) {
	p.GetSuiteMeta().GetContainer().Finish()
	_ = p.GetSuiteMeta().GetContainer().Print()
}

func setupErrorHandler(msg string, err error, meta provider.TestMeta, result SuiteResult) SuiteResult {
	mtx := sync.Mutex{}
	mtx.Lock()
	defer mtx.Unlock()

	tRes := NewTestResult(meta.GetResult(), meta.GetContainer())
	tRes.GetResult().Status = allure.Failed
	tRes.GetResult().SetStatusMessage(msg)
	tRes.GetResult().SetStatusTrace(fmt.Sprintf("%s. Reason:\n%s", msg, err.Error()))
	_ = tRes.Print()
	result.NewResult(tRes)
	return result
}

func runHook(t internalT, hookFunc common.HookFunc) (res bool, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			ctxName := t.GetProvider().ExecutionContext().GetName()
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, rec, debug.Stack())
			err = fmt.Errorf(errMsg)
			common.TestError(t, t.GetProvider(), t.GetProvider().ExecutionContext().GetName(), errMsg)
		}
	}()
	return hookFunc(t, t.GetProvider())
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

func copyLabels(input, target *allure.Result) *allure.Result {
	if input == nil || target == nil {
		return target
	}

	if epic, ok := input.GetFirstLabel(allure.Epic); ok {
		target.AddLabel(epic)
	}

	if parentSuite, ok := input.GetFirstLabel(allure.ParentSuite); ok {
		target.AddLabel(parentSuite)
	}

	if lead, ok := input.GetFirstLabel(allure.Lead); ok {
		target.AddLabel(lead)
	}

	if owner, ok := input.GetFirstLabel(allure.Owner); ok {
		target.AddLabel(owner)
	}

	return target
}
