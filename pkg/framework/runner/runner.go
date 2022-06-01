package runner

import (
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/adapter"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"
)

type InternalT interface {
	provider.T

	SetRealT(t provider.TestingT)
	GetProvider() provider.Provider
	WG() *sync.WaitGroup
	GetResult() *allure.Result
}

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type testFunc func(t provider.T)

type test struct {
	testBody testFunc
	testMeta provider.TestMeta
}

func newTest(body testFunc, testMeta provider.TestMeta) *test {
	return &test{
		testBody: body,
		testMeta: testMeta,
	}
}

type TestRunner interface {
	NewTest(testName string, testBody func(provider.T), tags ...string)
	BeforeEach(hookBody func(provider.T))
	AfterEach(hookBody func(provider.T))
	BeforeAll(hookBody func(provider.T))
	AfterAll(hookBody func(provider.T))
	RunTests() map[string]bool
	T() InternalT
}

type runner struct {
	internalT InternalT
	testPlan  *TestPlan
	tests     map[string]*test
}

func NewRunner(realT TestingT, suiteName string) TestRunner {
	newT := common.NewT(realT)

	callers := strings.Split(realT.Name(), "/")
	providerCfg := manager.NewProviderConfig().
		WithFullName(realT.Name()).
		WithPackageName(getPackage(2)).
		WithSuiteName(suiteName).
		WithRunner(callers[0])
	newT.SetProvider(manager.NewProvider(providerCfg))

	testPlan, err := NewTestPlan()
	if err != nil {
		fmt.Printf("Cannot find test plan. Reason: %s\n", err.Error())
	}
	return &runner{internalT: newT, tests: make(map[string]*test), testPlan: testPlan}
}

func (r *runner) IsRun(result *allure.Result) bool {
	if r.testPlan != nil {
		return r.testPlan.IsSelected(result.TestCaseID, result.FullName)
	}
	return true
}

func (r *runner) T() InternalT {
	return r.internalT
}

func (r *runner) NewTest(testName string, testBody func(provider.T), tags ...string) {
	fullName := fmt.Sprintf("%s/%s", r.T().Name(), testName)

	testMeta := adapter.NewTestMeta(
		r.T().GetProvider().GetSuiteMeta().GetSuiteFullName(),
		r.T().GetProvider().GetSuiteMeta().GetSuiteName(),
		testName,
		getPackage(2),
		tags...,
	)
	if !r.IsRun(testMeta.GetResult()) {
		return
	}

	r.tests[fullName] = newTest(testBody, testMeta)
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

func (r *runner) realT() TestingT {
	return r.internalT.RealT()
}

func (r *runner) RunTests() map[string]bool {
	var (
		wg     = &sync.WaitGroup{}
		result = make(map[string]bool)

		beforeAllHook  = common.CarriedHook(common.BeforeAll, r.internalT.GetProvider().GetSuiteMeta().GetBeforeAll)
		afterAllHook   = common.CarriedHook(common.AfterAll, r.internalT.GetProvider().GetSuiteMeta().GetAfterAll)
		beforeEachHook = common.CarriedHook(common.BeforeEach, r.internalT.GetProvider().GetTestMeta().GetBeforeEach)
		afterEachHook  = common.CarriedHook(common.AfterEach, r.internalT.GetProvider().GetTestMeta().GetAfterEach)
	)

	finishSuite := func(p provider.Provider) {
		p.GetSuiteMeta().GetContainer().Finish()
		_ = p.GetSuiteMeta().GetContainer().Print()
	}

	runHook := func(t InternalT, wg *sync.WaitGroup, hookFunc common.HookFunc) (bool, error) {
		defer t.WG().Wait()
		return hookFunc(t, t.GetProvider(), wg)
	}

	setupTest := func(t *testing.T, meta provider.TestMeta) *common.Common {
		testT := common.NewT(t)
		packageName := r.internalT.GetProvider().GetSuiteMeta().GetPackageName()
		suiteName := r.internalT.GetProvider().GetSuiteMeta().GetSuiteName()
		callers := strings.Split(t.Name(), "/")
		cfg := manager.NewProviderConfig().
			WithFullName(t.Name()).
			WithPackageName(packageName).
			WithSuiteName(suiteName).
			WithRunner(callers[0])
		testT.SetProvider(manager.NewProvider(cfg))

		testT.Provider.TestContext()
		meta.SetBeforeEach(r.internalT.GetProvider().GetTestMeta().GetBeforeEach())
		meta.SetAfterEach(r.internalT.GetProvider().GetTestMeta().GetAfterEach())
		meta.SetResult(copyLabels(r.internalT.GetResult(), meta.GetResult()))
		testT.Provider.SetTestMeta(meta)

		return testT
	}

	handleError := func(msg string, err error, allureResult *allure.Result) {
		allureResult.Status = allure.Unknown
		allureResult.SetStatusMessage(msg)
		allureResult.SetStatusTrace(fmt.Sprintf("%s. Reason:\n%s", msg, err.Error()))
	}

	defer finishSuite(r.internalT.GetProvider())
	defer func() {
		for _, testMeta := range r.tests {
			testMeta.testMeta.GetResult().Done()
			testMeta.testMeta.GetContainer().Finish()
			_ = testMeta.testMeta.GetContainer().Print()
		}
	}()
	defer func() {
		rec := recover()
		if rec != nil {
			ctxName := r.internalT.GetProvider().ExecutionContext().GetName()
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, rec, debug.Stack())
			common.TestError(r.internalT, r.internalT.GetProvider(), r.internalT.GetProvider().ExecutionContext().GetName(), errMsg)
		}
	}()

	// after all hook
	defer func() {
		runHook(r.internalT, wg, afterAllHook)
	}()
	for _, testMeta := range r.tests {
		r.internalT.GetProvider().GetSuiteMeta().GetContainer().AddChild(testMeta.testMeta.GetResult().UUID)
	}

	// before all hook
	ok, err := runHook(r.internalT, wg, beforeAllHook)
	if err != nil {
		for _, testMeta := range r.tests {
			handleError("Suite Setup failed", err, testMeta.testMeta.GetResult())
		}
		return result
	}
	if !ok {
		for _, testMeta := range r.tests {
			handleError("Suite Setup failed", fmt.Errorf("some assertion error during Suite Setup"), testMeta.testMeta.GetResult())
		}
		return result
	}

	for fullName, testData := range r.tests {
		wg.Add(1)
		result[fullName] = r.realT().Run(testData.testMeta.GetResult().Name, func(t *testing.T) {
			defer wg.Done()
			testT := setupTest(t, testData.testMeta)

			// after each hook
			defer runHook(testT, testT.WG(), afterEachHook)
			defer func() {
				rec := recover()
				if rec != nil {
					ctxName := testT.GetProvider().ExecutionContext().GetName()
					errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, rec, debug.Stack())
					common.TestError(testT, testT.Provider, testT.Provider.ExecutionContext().GetName(), errMsg)
				}
			}()

			// before each hook
			ok, err = runHook(testT, r.internalT.WG(), beforeEachHook)
			if err != nil {
				handleError("Test Setup failed", err, testData.testMeta.GetResult())
				return
			}
			if !ok {
				handleError("Test Setup failed", fmt.Errorf("assertion error due test setup"), testData.testMeta.GetResult())
				return
			}

			testT.GetProvider().TestContext()
			testData.testBody(testT)
			testT.WG().Wait()
		})
	}
	return result
}

func Run(t *testing.T, testName string, testBody func(provider.T), tags ...string) bool {
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

	if epics := input.GetLabel(allure.Epic); len(epics) > 0 {
		target.SetLabel(epics[0])
	}

	if parentSuites := input.GetLabel(allure.ParentSuite); len(parentSuites) > 0 {
		target.SetLabel(parentSuites[0])
	}

	if leads := input.GetLabel(allure.Lead); len(leads) > 0 {
		target.SetLabel(leads[0])
	}

	if owners := input.GetLabel(allure.Owner); len(owners) > 0 {
		target.SetLabel(owners[0])
	}

	return target
}
