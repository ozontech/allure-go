package provider

import (
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
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
	SkipOnPrint()
	Assert() Asserts
	Require() Asserts
	Run(testName string, testBody func(T), tags ...string) bool

	WithNewStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)
	WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)
}

type StepCtx interface {
	Step(step *allure.Step)
	NewStep(stepName string, parameters ...allure.Parameter)
	WithNewStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)
	WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)

	WithParameters(parameters ...allure.Parameter)
	WithNewParameters(kv ...interface{})

	WithAttachments(attachment ...*allure.Attachment)
	WithNewAttachment(name string, mimeType allure.MimeType, content []byte)

	Assert() Asserts
	Require() Asserts

	CurrentStep() *allure.Step

	Broken()

	Fail()
	FailNow()
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Name() string
}

// Asserts ...
type Asserts interface {
	Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{})
	Error(err error, msgAndArgs ...interface{})
	NoError(err error, msgAndArgs ...interface{})
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
	IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{})
	True(value bool, msgAndArgs ...interface{})
	False(value bool, msgAndArgs ...interface{})
	Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{})
}
