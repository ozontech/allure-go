package common

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/internal"
	"github.com/ozontech/allure-go/pkg/provider/pkg/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type common struct {
	provider.TestingT

	assert  provider.Asserts
	require provider.Asserts

	provider Provider
	xSkip    bool

	wg sync.WaitGroup
}

func NewT(realT provider.TestingT, suiteName string) provider.T {
	packageName := getPackage(2)
	callers := strings.Split(realT.Name(), "/")
	newT := &common{TestingT: realT, provider: newProvider(packageName, callers[0], realT.Name(), suiteName)}
	newT.assert = helper.NewAssertsHelper(newT)
	newT.require = helper.NewRequireHelper(newT)
	return newT
}

func NewTestT(parentT provider.InternalT, realT *testing.T, testName string, tags ...string) provider.T {
	newT := NewT(realT, parentT.Provider().GetSuiteMeta().GetSuiteName()).(provider.InternalT)
	newT.Provider().NewTest(testName, parentT.GetPackage(), tags...)
	newT.Provider().TestContext()
	newT.Provider().GetTestMeta().SetBeforeEach(parentT.Provider().GetTestMeta().GetBeforeEach())
	newT.Provider().GetTestMeta().SetAfterEach(parentT.Provider().GetTestMeta().GetAfterEach())

	parentT.Provider().GetSuiteMeta().GetContainer().AddChild(newT.GetResult().UUID)
	newT.Provider().GetTestMeta().SetResult(copyLabels(parentT.GetResult(), newT.Provider().GetTestMeta().GetResult()))
	return newT.(provider.T)
}

func (c *common) RealT() provider.TestingT {
	return c.TestingT
}

func (c *common) WG() *sync.WaitGroup {
	return &c.wg
}

func (c *common) Provider() provider.Provider {
	return c.provider
}

func (c *common) SkipOnPrint() {
	c.GetResult().SkipOnPrint()
}

func (c *common) XSkip() {
	c.xSkip = true
}

func (c *common) Error(args ...interface{}) {
	fullMessage := fmt.Sprintf("%s", args...)
	c.registerError(fullMessage)
	c.RealT().Error(args...)
}

func (c *common) Errorf(format string, args ...interface{}) {
	fullMessage := fmt.Sprintf(format, args...)
	c.registerError(fullMessage)
	c.RealT().Errorf(format, args...)
}

func (c *common) Skip(args ...interface{}) {
	c.safely(func(result *allure.Result) {
		skipMessage := fmt.Sprintln(args...)
		if len(skipMessage) > 100 {
			result.StatusDetails.Message = skipMessage[:100]
		} else {
			result.StatusDetails.Message = skipMessage
		}
		result.StatusDetails.Trace = skipMessage
		result.Status = allure.Skipped
	})
	c.RealT().Skip(args...)
}

func (c *common) Assert() provider.Asserts {
	return c.assert
}

func (c *common) Require() provider.Asserts {
	return c.require
}

func (c *common) GetResult() *allure.Result {
	return c.provider.GetTestMeta().GetResult()
}

func (c *common) GetSuite() string {
	return c.provider.GetSuiteMeta().GetSuiteName()
}

func (c *common) GetPackage() string {
	return c.provider.GetSuiteMeta().GetPackageName()
}

func (c *common) BreakResult(reason string) {
	if _result := c.GetResult(); _result != nil {
		_result.Status = allure.Broken
		_result.StatusDetails.Message = reason[:100]
		_result.StatusDetails.Trace = reason
		_result.Stop = allure.GetNow()
	}
}

func (c *common) Run(testName string, testBody func(provider.T), tags ...string) bool {
	return c.RealT().Run(testName, func(realT *testing.T) {
		testT := NewTestT(c, realT, testName, tags...).(provider.InternalT)

		// print test result
		defer testT.Provider().FinishTest()

		defer func() {
			rec := recover()
			// wait for all test's async steps over
			testT.WG().Wait()
			if rec != nil {
				errMsg := fmt.Sprintf("Test panicked: %v\n%s", rec, debug.Stack())
				internal.TestError(c.Provider().ExecutionContext().GetName(), errMsg, testT)
			}
		}()

		testT.Provider().TestContext()
		testBody(testT.(provider.T))
	})
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

	return target
}

func (c *common) safely(f func(result *allure.Result)) {
	if result := c.GetResult(); result != nil {
		f(result)
	}
}

func (c *common) registerError(fullMessage string) {
	xSkipPrefix := "[XSkip]"
	result := c.GetResult()
	if result != nil && result.Status != allure.Broken {
		if c.xSkip {
			result.Name = fmt.Sprintf("%s%s", xSkipPrefix, result.Name)
			c.Skip(fullMessage)
		}
		result.Status = allure.Failed
		result.StatusDetails.Message = internal.ExtractErrorMessages(fullMessage)
		result.StatusDetails.Trace = fmt.Sprintf("%s\n%s", result.StatusDetails.Trace, fullMessage)
	}
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
