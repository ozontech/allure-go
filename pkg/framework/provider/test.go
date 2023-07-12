package provider

import (
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/stretchr/testify/assert"
)

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type T interface {
	testing.TB
	AllureForward

	Parallel()

	RealT() TestingT
	XSkip()
	Break(args ...interface{})
	Breakf(format string, args ...interface{})
	Broken()
	BrokenNow()
	SkipOnPrint()
	Assert() Asserts
	Require() Asserts
	Run(testName string, testBody func(T), tags ...string) *allure.Result

	LogStep(args ...interface{})
	LogfStep(format string, args ...interface{})
	WithNewStep(stepName string, step func(sCtx StepCtx), params ...*allure.Parameter)
	WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...*allure.Parameter)
	WithTestSetup(setup func(T))
	WithTestTeardown(teardown func(T))
}

type StepCtx interface {
	Step(step *allure.Step)
	NewStep(stepName string, parameters ...*allure.Parameter)
	WithNewStep(stepName string, step func(sCtx StepCtx), params ...*allure.Parameter)
	WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...*allure.Parameter)

	WithParameters(parameters ...*allure.Parameter)
	WithNewParameters(kv ...interface{})

	WithAttachments(attachment ...*allure.Attachment)
	WithNewAttachment(name string, mimeType allure.MimeType, content []byte)

	Assert() Asserts
	Require() Asserts

	LogStep(args ...interface{})
	LogfStep(format string, args ...interface{})
	CurrentStep() *allure.Step

	Broken()
	BrokenNow()

	Fail()
	FailNow()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Break(args ...interface{})
	Breakf(format string, args ...interface{})
	Name() string
}

// Asserts ...
type Asserts interface {
	Exactly(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	Same(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	NotSame(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	EqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	NotEqualValues(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	Error(err error, msgAndArgs ...interface{})
	NoError(err error, msgAndArgs ...interface{})
	EqualError(theError error, errString string, msgAndArgs ...interface{})
	ErrorIs(err error, target error, msgAndArgs ...interface{})
	ErrorAs(err error, target interface{}, msgAndArgs ...interface{})
	NotNil(object interface{}, msgAndArgs ...interface{})
	Nil(object interface{}, msgAndArgs ...interface{})
	Len(object interface{}, length int, msgAndArgs ...interface{})
	NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{})
	Contains(s interface{}, contains interface{}, msgAndArgs ...interface{})
	Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
	Empty(object interface{}, msgAndArgs ...interface{})
	NotEmpty(object interface{}, msgAndArgs ...interface{})
	WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{})
	JSONEq(expected, actual string, msgAndArgs ...interface{})
	JSONContains(expected, actual string, msgAndArgs ...interface{})
	Subset(list, subset interface{}, msgAndArgs ...interface{})
	NotSubset(list, subset interface{}, msgAndArgs ...interface{})
	IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{})
	True(value bool, msgAndArgs ...interface{})
	False(value bool, msgAndArgs ...interface{})
	Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{})
	ElementsMatch(listA interface{}, listB interface{}, msgAndArgs ...interface{})
	DirExists(path string, msgAndArgs ...interface{})
	Condition(condition assert.Comparison, msgAndArgs ...interface{})
	Zero(i interface{}, msgAndArgs ...interface{})
	NotZero(i interface{}, msgAndArgs ...interface{})
	InDelta(expected, actual interface{}, delta float64, msgAndArgs ...interface{})
}
