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
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type Common struct {
	provider.TestingT
	provider.Provider

	assert  provider.Asserts
	require provider.Asserts

	xSkip bool

	wg sync.WaitGroup
}

// NewT returns Common instance that implementing provider.T interface
func NewT(realT provider.TestingT, packageName, suiteName string) *Common {
	callers := strings.Split(realT.Name(), "/")
	cfg := manager.NewProviderConfig().
		WithFullName(realT.Name()).
		WithPackageName(packageName).
		WithSuiteName(suiteName).
		WithRunner(callers[0])
	newT := &Common{TestingT: realT, Provider: manager.NewProvider(cfg)}
	newT.assert = helper.NewAssertsHelper(newT)
	newT.require = helper.NewRequireHelper(newT)
	return newT
}

// NewTestT returns Common instance that implementing provider.T interface, that copy labels of parent instance
func NewTestT(realT provider.TestingT, provider provider.Provider, parentT ParentT, packageName string, testName string, tags ...string) *Common {
	newT := NewT(realT, packageName, provider.GetSuiteMeta().GetSuiteName())
	newT.Provider.NewTest(testName, packageName, tags...)
	newT.Provider.TestContext()
	newT.Provider.GetTestMeta().SetBeforeEach(provider.GetTestMeta().GetBeforeEach())
	newT.Provider.GetTestMeta().SetAfterEach(provider.GetTestMeta().GetAfterEach())

	parentT.GetProvider().GetSuiteMeta().GetContainer().AddChild(newT.GetResult().UUID)
	newT.Provider.GetTestMeta().SetResult(copyLabels(parentT.GetResult(), newT.Provider.GetTestMeta().GetResult()))
	return newT
}

func (c *Common) registerError(fullMessage string) {
	xSkipPrefix := "[XSkip]"
	result := c.GetResult()
	if result != nil && result.Status != allure.Broken {
		if c.xSkip {
			result.Name = fmt.Sprintf("%s%s", xSkipPrefix, result.Name)
			c.Skip(fullMessage)
		}
		result.Status = allure.Failed
		result.StatusDetails.Message = extractErrorMessages(fullMessage)
		result.StatusDetails.Trace = fmt.Sprintf("%s\n%s", result.StatusDetails.Trace, fullMessage)
	}
}

func (c *Common) safely(f func(result *allure.Result)) {
	if result := c.GetResult(); result != nil {
		f(result)
	}
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

// Error ...
func (c *Common) Error(args ...interface{}) {
	fullMessage := fmt.Sprintf("%s", args...)
	c.registerError(fullMessage)
	c.TestingT.Error(args...)
}

// Errorf ...
func (c *Common) Errorf(format string, args ...interface{}) {
	fullMessage := fmt.Sprintf(format, args...)
	c.registerError(fullMessage)
	c.TestingT.Errorf(format, args...)
}

// Run runs test body as test with passed tags
func (c *Common) Run(testName string, testBody func(provider.T), tags ...string) bool {
	return c.TestingT.Run(testName, func(realT *testing.T) {
		testT := NewTestT(realT, c.Provider, c, c.Provider.GetSuiteMeta().GetPackageName(), testName, tags...)

		// print test result
		defer testT.Provider.FinishTest()

		defer func() {
			rec := recover()
			// wait for all test's async steps over
			testT.wg.Wait()
			if rec != nil {
				errMsg := fmt.Sprintf("Test panicked: %v\n%s", rec, debug.Stack())
				TestError(testT, c.Provider, c.Provider.ExecutionContext().GetName(), errMsg)
			}
		}()

		testT.Provider.TestContext()
		testBody(testT)
	})
}

// Skip ...
func (c *Common) Skip(args ...interface{}) {
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
	c.TestingT.Skip(args...)
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

func extractErrorMessages(output string) string {
	r := regexp.MustCompile(`Messages:(.*)`)
	result := strings.Trim(strings.TrimPrefix(r.FindString(output), "Messages:   "), " ")
	if result == "" {
		left := "\tError:"
		right := "\tTest:"
		r2 := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
		result = r2.FindString(output)
		result = strings.Trim(strings.TrimSuffix(result, "\tTest:"), " ")
	}
	if result == "" {
		return output
	}
	return result
}
