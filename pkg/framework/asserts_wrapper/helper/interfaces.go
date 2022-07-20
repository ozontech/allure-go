package helper

import (
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
)

type ProviderT interface {
	Step(step *allure.Step)
	Errorf(format string, args ...interface{})
	FailNow()
}

// AssertsHelper ...
type AssertsHelper interface {
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
