package asserts

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
	wrapper.NewAsserts(t).Exactly(t, expected, actual, msgAndArgs...)
}

// Same ...
func Same(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Same(t, expected, actual, msgAndArgs...)
}

// NotSame ...
func NotSame(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotSame(t, expected, actual, msgAndArgs...)
}

// Equal ...
func Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Equal(t, expected, actual, msgAndArgs...)
}

// NotEqual ...
func NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotEqual(t, expected, actual, msgAndArgs...)
}

// EqualValues ...
func EqualValues(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).EqualValues(t, expected, actual, msgAndArgs...)
}

// NotEqualValues ...
func NotEqualValues(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotEqualValues(t, expected, actual, msgAndArgs...)
}

// Error ...
func Error(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Error(t, err, msgAndArgs...)
}

// NoError ...
func NoError(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NoError(t, err, msgAndArgs...)
}

// EqualError ...
func EqualError(t ProviderT, theError error, errString string, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).EqualError(t, theError, errString, msgAndArgs...)
}

// ErrorIs ...
func ErrorIs(t ProviderT, err error, target error, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).ErrorIs(t, err, target, msgAndArgs...)
}

// ErrorAs ...
func ErrorAs(t ProviderT, err error, target interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).ErrorAs(t, err, target, msgAndArgs...)
}

// NotNil ...
func NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotNil(t, object, msgAndArgs...)
}

// Nil ...
func Nil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Nil(t, object, msgAndArgs...)
}

// Len ...
func Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Len(t, object, length, msgAndArgs...)
}

// NotContains ...
func NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotContains(t, s, contains, msgAndArgs...)
}

// Contains ...
func Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Contains(t, s, contains, msgAndArgs...)
}

// Greater ...
func Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Greater(t, e1, e2, msgAndArgs...)
}

// GreaterOrEqual ...
func GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).GreaterOrEqual(t, e1, e2, msgAndArgs...)
}

// Less ...
func Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Less(t, e1, e2, msgAndArgs...)
}

// LessOrEqual ...
func LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).LessOrEqual(t, e1, e2, msgAndArgs...)
}

// Implements ...
func Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Implements(t, interfaceObject, object, msgAndArgs...)
}

// Empty ...
func Empty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Empty(t, object, msgAndArgs...)
}

// NotEmpty ...
func NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotEmpty(t, object, msgAndArgs...)
}

// WithinDuration ...
func WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

// JSONEq ...
func JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).JSONEq(t, expected, actual, msgAndArgs...)
}

// JSONContains ...
func JSONContains(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).JSONContains(t, expected, actual, msgAndArgs...)
}

// Subset ...
func Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Subset(t, list, subset, msgAndArgs...)
}

// NotSubset ...
func NotSubset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotSubset(t, list, subset, msgAndArgs...)
}

// IsType ...
func IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).IsType(t, expectedType, object, msgAndArgs...)
}

// True ...
func True(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).True(t, value, msgAndArgs...)
}

// False ...
func False(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).False(t, value, msgAndArgs...)
}

// Regexp ...
func Regexp(t ProviderT, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Regexp(t, rx, str, msgAndArgs...)
}

// ElementsMatch ...
func ElementsMatch(t ProviderT, listA interface{}, listB interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).ElementsMatch(t, listA, listB, msgAndArgs...)
}

// DirExists ...
func DirExists(t ProviderT, path string, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).DirExists(t, path, msgAndArgs...)
}

// Condition ...
func Condition(t ProviderT, condition assert.Comparison, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Condition(t, condition, msgAndArgs...)
}

// Zero ...
func Zero(t ProviderT, i interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Zero(t, i, msgAndArgs...)
}

// NotZero ...
func NotZero(t ProviderT, i interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotZero(t, i, msgAndArgs...)
}
