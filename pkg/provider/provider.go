package provider

import (
	"fmt"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"testing"

	"github.com/koodeex/allure-testify/pkg/allure"
)

type IAllureProvider interface {
	AllureActions
	AllureInfo
	AllureLabels
	AllureLinks

	SetResult(result *allure.Result)
	GetResult() *allure.Result
	GetSuite() string
	GetPackage() string
	GetContainer() *allure.Container

	RealT() *testing.T
}

type T struct {
	runner      string
	suite       string
	packageName string

	result          *allure.Result
	parentContainer *allure.Container

	beforeEach func(*T)
	afterEach  func(*T)

	xskip bool

	*testing.T
	*contextController
}

func NewT(t *testing.T, runner, suite string) *T {
	newT := &T{T: t, runner: runner, suite: suite}
	newT.contextController = newStateMachine(newT)
	newT.parentContainer = allure.NewContainer()
	return newT
}

func NewTWithPackage(t *testing.T, runner, suite, packageName string) *T {
	newT := NewT(t, runner, suite)
	newT.packageName = packageName
	return newT
}

func NewTForTest(t *T, result *allure.Result, parentContainer *allure.Container) *T {
	newT := NewT(t.RealT(), t.runner, t.suite)
	newT.result = result
	newT.parentContainer = t.parentContainer
	if newT.parentContainer == nil {
		newT.parentContainer = parentContainer
	}
	return newT
}

func (t *T) Error(args ...interface{}) {
	fullMessage := fmt.Sprintf("%s", args...)
	t.registerError(fullMessage)
	t.RealT().Error(args...)
}

func (t *T) Errorf(format string, args ...interface{}) {
	fullMessage := fmt.Sprintf("%s", args...)
	t.registerError(fullMessage)
	t.RealT().Errorf(format, args...)
}

func (t *T) Skip(args ...interface{}) {
	t.safely(func(result *allure.Result) {
		skipMessage := fmt.Sprintln(args...)
		if len(skipMessage) > 100 {
			result.StatusDetails.Message = skipMessage[:100]
		} else {
			result.StatusDetails.Message = skipMessage
		}
		result.StatusDetails.Trace = skipMessage
		result.Status = allure.Skipped
	})
	t.RealT().Skip(args...)
}

func (t *T) Run(testName string, f func(*T), tags ...string) bool {
	var packageName string

	realT := t.RealT()

	if t.packageName == "" {
		packageName = getPackage(2)
	} else {
		packageName = t.packageName
	}

	result := allure.NewResultHelper().GetNewResult(t, testName, packageName, tags...)

	container := t.GetContainer()
	container.AddChild(result.UUID)
	newT := NewTForTest(t, result, container)
	return realT.Run(testName, func(testT *testing.T) {
		//dirty magic
		newT.T = testT
		defer func() {
			result.Done()
		}()

		defer func() {
			r := recover()
			if r != nil {
				errMsg := fmt.Sprintf("test panicked: %v\n%s", r, debug.Stack())
				newT.BreakResult(errMsg)
				newT.Errorf(errMsg)
				newT.FailNow()
			}
		}()
		defer func() {
			result.Finish()
			if t.afterEach != nil {
				t.afterEach(newT)
			}
			result.Container.Finish()
		}()

		result.Container.Begin()
		if t.beforeEach != nil {
			t.beforeEach(newT)
		}

		result.Begin()
		f(newT)
	})
}

func (t *T) RealT() *testing.T {
	return t.T
}

func (t *T) SetResult(result *allure.Result) {
	t.result = result
}

func (t *T) BreakResult(reason string) {
	if _result := t.GetResult(); _result != nil {
		_result.Status = allure.Broken
		_result.StatusDetails.Message = reason[:100]
		_result.StatusDetails.Trace = reason
		_result.Stop = allure.GetNow()
	}
}

func (t *T) GetSuite() string {
	return t.suite
}

func (t *T) GetPackage() string {
	return t.packageName
}

func (t *T) GetResult() *allure.Result {
	return t.result
}

func (t *T) GetContainer() *allure.Container {
	if t.parentContainer == nil {
		t.parentContainer = allure.NewContainer()
		return t.parentContainer
	}
	return t.parentContainer
}

func (t *T) safely(f func(result *allure.Result)) {
	if result := t.GetResult(); result != nil {
		f(result)
	}
}

func (t *T) registerError(fullMessage string) {
	xSkipPrefix := "[XSkip]"
	result := t.GetResult()
	if result != nil && result.Status != allure.Broken {
		if step := result.StepsQueue.Last(); step != nil {
			step.Status = allure.Failed
		}
		if t.xskip {
			t.result.Name = fmt.Sprintf("%s%s", xSkipPrefix, t.result.Name)
			t.Skip(fullMessage)
		}
		result.Status = allure.Failed
		result.StatusDetails.Message = extractErrorMessages(fullMessage)
		result.StatusDetails.Trace = fmt.Sprintf("%s\n%s", result.StatusDetails.Trace, fullMessage)
	}
}

func extractErrorMessages(output string) string {
	r := regexp.MustCompile(`Messages:(.*)`)
	result := strings.TrimPrefix(r.FindString(output), "Messages:   ")
	left := "\tError:"
	right := "\tTest:"
	if result == "" {
		r2 := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
		result = r2.FindString(output)
		result = strings.Trim(strings.TrimSuffix(result, "\tTest:"), " ")
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
