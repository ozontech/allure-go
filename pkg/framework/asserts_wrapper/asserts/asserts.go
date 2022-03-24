package asserts

import (
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/wrapper"
)

type ProviderT interface {
	Step(step *allure.Step)
	Errorf(format string, args ...interface{})
	FailNow()
}

// Equal ...
func Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Equal(t, expected, actual, msgAndArgs...)
}

// NotEqual ...
func NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NotEqual(t, expected, actual, msgAndArgs...)
}

// Error ...
func Error(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Error(t, err, msgAndArgs...)
}

// NoError ...
func NoError(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).NoError(t, err, msgAndArgs...)
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

// Subset ...
func Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewAsserts(t).Subset(t, list, subset, msgAndArgs...)
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
