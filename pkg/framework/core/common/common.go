package common

import (
	"fmt"
	"regexp"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/testplan"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

const logMsgSetupTearDown = "%s will be skipped. Reason: wrong context. Expected: %s; Actual: %s"

type Common struct {
	provider.TestingT
	provider.Provider

	assert  provider.Asserts
	require provider.Asserts

	xSkip bool

	wg sync.WaitGroup

	tempDirMu  sync.Mutex
	tempDir    string
	tempDirErr error
	tempDirSeq int32
}

// NewT returns Common instance that implementing provider.T interface
func NewT(realT provider.TestingT) *Common {
	newT := &Common{TestingT: realT}
	newT.assert = helper.NewAssertsHelper(newT)
	newT.require = helper.NewRequireHelper(newT)
	return newT
}

func (c *Common) registerError(fullMessage string) {
	xSkipPrefix := "[XSkip]"
	res := c.GetResult()

	if res != nil && res.Status != allure.Broken {
		if c.xSkip {
			res.Name = fmt.Sprintf("%s%s", xSkipPrefix, res.Name)
			c.Skip(fullMessage)
		}
		res.Status = allure.Failed
	}

	if res != nil {
		res.StatusDetails.Message = extractErrorMessages(fullMessage)
		res.StatusDetails.Trace = fmt.Sprintf("%s\n%s", res.StatusDetails.Trace, fullMessage)
	}
}

func (c *Common) withResult(f func(result *allure.Result)) {
	if r := c.GetResult(); r != nil {
		f(r)
	}
}

func (c *Common) SetProvider(provider provider.Provider) {
	c.Provider = provider
}

// WG ...
func (c *Common) WG() *sync.WaitGroup {
	return &c.wg
}

// RealT returns instance of testing.T
func (c *Common) RealT() provider.TestingT {
	return c.TestingT
}

// Assert ...
func (c *Common) Assert() provider.Asserts {
	return c.assert
}

// Require ...
func (c *Common) Require() provider.Asserts {
	return c.require
}

// XSkip marks current test as XSkip that means that in case of assert fail this test will be marked skipped
func (c *Common) XSkip() {
	c.xSkip = true
}

// GetProvider ...
func (c *Common) GetProvider() provider.Provider {
	return c.Provider
}

// SkipOnPrint skips creating of report for current test
func (c *Common) SkipOnPrint() {
	c.GetResult().SkipOnPrint()
}

// LogStep ...
func (c *Common) LogStep(args ...interface{}) {
	c.Step(allure.NewSimpleStep(fmt.Sprintln(args...)))
	c.Log(args...)
}

// LogfStep ...
func (c *Common) LogfStep(format string, args ...interface{}) {
	c.Step(allure.NewSimpleStep(fmt.Sprintf(format, args...)))
	c.Logf(format, args...)
}

// Error ...
func (c *Common) Error(args ...interface{}) {
	c.Helper()

	fullMessage := fmt.Sprint(args...)
	c.registerError(fullMessage)
	c.TestingT.Error(args...)
}

// Errorf ...
func (c *Common) Errorf(format string, args ...interface{}) {
	c.Helper()

	fullMessage := fmt.Sprintf(format, args...)
	c.registerError(fullMessage)
	c.TestingT.Errorf(format, args...)
}

// Fatal ...
func (c *Common) Fatal(args ...interface{}) {
	c.Helper()

	fullMessage := fmt.Sprintf("%s", args...)
	c.registerError(fullMessage)
	c.TestingT.Fatal(args...)
}

// Fatalf ...
func (c *Common) Fatalf(format string, args ...interface{}) {
	c.Helper()

	fullMessage := fmt.Sprintf(format, args...)
	c.registerError(fullMessage)
	c.TestingT.Fatalf(format, args...)
}

// Break ...
func (c *Common) Break(args ...interface{}) {
	c.withResult(func(r *allure.Result) {
		r.Status = allure.Broken
	})

	fullMessage := fmt.Sprint(args...)
	c.registerError(fullMessage)
	c.FailNow()
}

// Breakf ...
func (c *Common) Breakf(format string, args ...interface{}) {
	c.withResult(func(r *allure.Result) {
		r.Status = allure.Broken
	})

	fullMessage := fmt.Sprintf(format, args...)
	c.registerError(fullMessage)
	c.FailNow()
}

// Name ...
func (c *Common) Name() string {
	if c.GetProvider() != nil && c.GetProvider().GetResult() != nil {
		return c.GetProvider().GetResult().Name
	}

	return c.TestingT.Name()
}

// Fail ...
func (c *Common) Fail() {
	c.GetProvider().GetResult().Status = allure.Failed
	c.TestingT.Fail()
}

// FailNow ...
func (c *Common) FailNow() {
	c.withResult(func(r *allure.Result) {
		if r.Status != allure.Broken {
			r.Status = allure.Failed
		}
	})
	c.TestingT.FailNow()
}

// Broken ...
func (c *Common) Broken() {
	c.withResult(func(r *allure.Result) {
		r.Status = allure.Broken
	})
	c.TestingT.Fail()
}

// BrokenNow ...
func (c *Common) BrokenNow() {
	c.withResult(func(r *allure.Result) {
		r.Status = allure.Broken
	})

	c.TestingT.FailNow()
}

// Skip ...
func (c *Common) Skip(args ...interface{}) {
	c.withResult(func(r *allure.Result) {
		skipMessage := fmt.Sprintln(args...)
		if len(skipMessage) > 100 {
			// TODO: trim properly to avoid invalid UTF-8
			r.StatusDetails.Message = skipMessage[:100]
		} else {
			r.StatusDetails.Message = skipMessage
		}

		r.StatusDetails.Trace = skipMessage
		r.Status = allure.Skipped
	})

	c.TestingT.Skip(args...)
}

// Skipf ...
func (c *Common) Skipf(format string, args ...interface{}) {
	c.withResult(func(r *allure.Result) {
		skipMessage := fmt.Sprintf(format, args...)
		if len(skipMessage) > 100 {
			// TODO: trim properly to avoid invalid UTF-8
			r.StatusDetails.Message = skipMessage[:100]
		} else {
			r.StatusDetails.Message = skipMessage
		}

		r.StatusDetails.Trace = skipMessage
		r.Status = allure.Skipped
	})

	c.TestingT.Skipf(format, args...)
}

// WithTestSetup ...
func (c *Common) WithTestSetup(setup func(provider.T)) {
	currentContext := c.GetProvider().ExecutionContext().GetName()
	if currentContext != constants.TestContextName {
		c.Logf(logMsgSetupTearDown, "WithTestSetup", constants.TestContextName, currentContext)
		return
	}

	c.subContextWrapper(currentContext, BeforeEach, setup)
}

// WithTestTeardown ...
func (c *Common) WithTestTeardown(teardown func(provider.T)) {
	currentContext := c.GetProvider().ExecutionContext().GetName()
	if currentContext != constants.TestContextName {
		c.Logf(logMsgSetupTearDown, "WithTestTeardown", constants.TestContextName, currentContext)
		return
	}

	c.subContextWrapper(currentContext, AfterEach, teardown)
}

func (c *Common) subContextWrapper(currentContext string, newContext HookType, hook func(t provider.T)) {
	defer func() {
		rec := recover()
		c.GetProvider().TestContext()
		if rec != nil {
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", currentContext, rec, debug.Stack())
			TestError(c, c.GetProvider(), currentContext, errMsg)
		}
	}()

	switch newContext {
	case BeforeEach:
		c.GetProvider().BeforeEachContext()

	case AfterEach:
		c.GetProvider().AfterEachContext()

	default:
		panic(fmt.Sprintf("Wrong execution context used to run func. Expected: %s OR %s; Actual: %s", BeforeEach, AfterEach, newContext))
	}

	hook(c)
}

// Run runs test body as test with passed tags
func (c *Common) Run(testName string, testBody func(provider.T), tags ...string) (res *allure.Result) {
	parentCallers := strings.Split(c.RealT().Name(), "/")
	suiteName := parentCallers[len(parentCallers)-1]

	c.TestingT.Run(testName, func(realT *testing.T) {
		var (
			testT = NewT(realT)

			packageName = c.Provider.GetSuiteMeta().GetPackageName()
			parentSuite = c.Provider.GetSuiteMeta().GetSuiteName()

			callers = strings.Split(realT.Name(), "/")
		)

		if result := c.Provider.GetTestMeta().GetResult(); result != nil {
			suiteName = result.Name
		}

		providerCfg := manager.NewProviderConfig().
			WithFullName(realT.Name()).
			WithPackageName(packageName).
			WithSuiteName(suiteName).
			WithRunner(callers[0])

		if parentSuite != "" && parentSuite != suiteName && parentSuite != callers[len(callers)-1] {
			providerCfg = providerCfg.WithParentSuite(parentSuite)
		}
		newProvider := manager.NewProvider(providerCfg)

		newProvider.NewTest(testName, packageName, tags...)
		if testPlan := testplan.GetTestPlan(); testPlan != nil {
			if !testPlan.IsSelected(newProvider.GetTestMeta().GetResult().TestCaseID, newProvider.GetResult().FullName) {
				realT.Skip("Test is not Selected in Test Plan")
			}
		}
		newProvider.TestContext()

		testT.SetProvider(newProvider)

		defer func() {
			res = testT.GetResult()
		}()

		// print test result
		defer func() {
			err := testT.FinishTest()
			if err != nil {
				testT.Error(err.Error())
			}
		}()

		defer func() {
			rec := recover()
			// wait for all tests async steps over
			testT.wg.Wait()
			if rec != nil {
				errMsg := fmt.Sprintf("Test panicked: %v\n%s", rec, debug.Stack())
				TestError(testT, testT.Provider, testT.Provider.ExecutionContext().GetName(), errMsg)
			}
		}()

		testT.TestContext()
		testBody(testT)
	})
	return
}

func (c *Common) SetRealT(realT provider.TestingT) {
	c.TestingT = realT
}

func (c *Common) GetRealT() provider.TestingT {
	return c.TestingT
}

func copyLabels(input, target *allure.Result) *allure.Result {
	if input == nil || target == nil {
		return target
	}

	if epic, ok := input.GetFirstLabel(allure.Epic); ok {
		target.ReplaceLabel(epic)
	}

	if parentSuite, ok := input.GetFirstLabel(allure.ParentSuite); ok {
		target.ReplaceLabel(parentSuite)
	}

	if lead, ok := input.GetFirstLabel(allure.Lead); ok {
		target.ReplaceLabel(lead)
	}

	if owner, ok := input.GetFirstLabel(allure.Owner); ok {
		target.ReplaceLabel(owner)
	}

	return target
}

var messagesRegexp = regexp.MustCompile(`Messages:(.*)`)

func extractErrorMessages(output string) string {
	result := strings.TrimSpace(strings.TrimPrefix(messagesRegexp.FindString(output), "Messages:   "))

	if result == "" {
		left := "\tError:"
		right := "\tTest:"

		r2 := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))

		result = r2.FindString(output)
		result = strings.TrimSpace(strings.TrimSuffix(result, "\tTest:"))

		if result != "" {
			return result
		}
	}

	return output
}
