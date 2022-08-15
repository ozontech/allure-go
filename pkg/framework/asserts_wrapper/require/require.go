package require

import (
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/wrapper"
	"github.com/stretchr/testify/assert"
)

type ProviderT interface {
	Step(step *allure.Step)
	Errorf(format string, args ...interface{})
	FailNow()
}

// Exactly ...
func Exactly(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Exactly(t, expected, actual, msgAndArgs...)
}

// Same ...
func Same(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Same(t, expected, actual, msgAndArgs...)
}

// NotSame ...
func NotSame(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotSame(t, expected, actual, msgAndArgs...)
}

// Equal ...
func Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Equal(t, expected, actual, msgAndArgs...)
}

// NotEqual ...
func NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotEqual(t, expected, actual, msgAndArgs...)
}

// EqualValues ...
func EqualValues(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).EqualValues(t, expected, actual, msgAndArgs...)
}

// NotEqualValues ...
func NotEqualValues(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotEqualValues(t, expected, actual, msgAndArgs...)
}

// Error ...
func Error(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Error(t, err, msgAndArgs...)
}

// NoError ...
func NoError(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NoError(t, err, msgAndArgs...)
}

// EqualError ...
func EqualError(t ProviderT, theError error, errString string, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).EqualError(t, theError, errString, msgAndArgs...)
}

// ErrorIs ...
func ErrorIs(t ProviderT, err error, target error, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).ErrorIs(t, err, target, msgAndArgs...)
}

// ErrorAs ...
func ErrorAs(t ProviderT, err error, target interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).ErrorAs(t, err, target, msgAndArgs...)
}

// NotNil ...
func NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotNil(t, object, msgAndArgs...)
}

// Nil ...
func Nil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Nil(t, object, msgAndArgs...)
}

// Len ...
func Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Len(t, object, length, msgAndArgs...)
}

// NotContains ...
func NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotContains(t, s, contains, msgAndArgs...)
}

// Contains ...
func Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Contains(t, s, contains, msgAndArgs...)
}

// Greater ...
func Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Greater(t, e1, e2, msgAndArgs...)
}

// GreaterOrEqual ...
func GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).GreaterOrEqual(t, e1, e2, msgAndArgs...)
}

// Less ...
func Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Less(t, e1, e2, msgAndArgs...)
}

// LessOrEqual ...
func LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).LessOrEqual(t, e1, e2, msgAndArgs...)
}

// Implements ...
func Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Implements(t, interfaceObject, object, msgAndArgs...)
}

// Empty ...
func Empty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Empty(t, object, msgAndArgs...)
}

// NotEmpty ...
func NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotEmpty(t, object, msgAndArgs...)
}

// WithinDuration ...
func WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

// JSONEq ...
func JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).JSONEq(t, expected, actual, msgAndArgs...)
}

// JSONContains ...
func JSONContains(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).JSONContains(t, expected, actual, msgAndArgs...)
}

// Subset ...
func Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Subset(t, list, subset, msgAndArgs...)
}

// NotSubset ...
func NotSubset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotSubset(t, list, subset, msgAndArgs...)
}

// IsType ...
func IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).IsType(t, expectedType, object, msgAndArgs...)
}

// True ...
func True(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).True(t, value, msgAndArgs...)
}

// False ...
func False(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).False(t, value, msgAndArgs...)
}

// Regexp ...
func Regexp(t ProviderT, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Regexp(t, rx, str, msgAndArgs...)
}

// ElementsMatch ...
func ElementsMatch(t ProviderT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).ElementsMatch(t, listA, listB, msgAndArgs...)
}

// DirExists ...
func DirExists(t ProviderT, path string, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).DirExists(t, path, msgAndArgs...)
}

// Condition ...
func Condition(t ProviderT, condition assert.Comparison, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Condition(t, condition, msgAndArgs...)
}

// Zero ...
func Zero(t ProviderT, i interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Zero(t, i, msgAndArgs...)
}

// NotZero ...
func NotZero(t ProviderT, i interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotZero(t, i, msgAndArgs...)
}
