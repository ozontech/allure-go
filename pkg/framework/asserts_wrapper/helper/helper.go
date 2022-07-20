package helper

import (
	"time"

	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/wrapper"
)

type a struct {
	t       ProviderT
	asserts wrapper.AssertsWrapper
}

// Equal ...
func (a *a) Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a.asserts.Equal(a.t, expected, actual, msgAndArgs...)
}

// NotEqual ...
func (a *a) NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	a.asserts.NotEqual(a.t, expected, actual, msgAndArgs...)
}

// Error ...
func (a *a) Error(err error, msgAndArgs ...interface{}) {
	a.asserts.Error(a.t, err, msgAndArgs...)
}

// NoError ...
func (a *a) NoError(err error, msgAndArgs ...interface{}) {
	a.asserts.NoError(a.t, err, msgAndArgs...)
}

// Nil ...
func (a *a) Nil(object interface{}, msgAndArgs ...interface{}) {
	a.asserts.Nil(a.t, object, msgAndArgs...)
}

// NotNil ...
func (a *a) NotNil(object interface{}, msgAndArgs ...interface{}) {
	a.asserts.NotNil(a.t, object, msgAndArgs...)
}

// Len ...
func (a *a) Len(object interface{}, length int, msgAndArgs ...interface{}) {
	a.asserts.Len(a.t, object, length, msgAndArgs...)
}

// Contains ...
func (a *a) Contains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	a.asserts.Contains(a.t, s, contains, msgAndArgs...)
}

// NotContains ...
func (a *a) NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	a.asserts.NotContains(a.t, s, contains, msgAndArgs...)
}

// Greater ...
func (a *a) Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a.asserts.Greater(a.t, e1, e2, msgAndArgs...)
}

// GreaterOrEqual ...
func (a *a) GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a.asserts.GreaterOrEqual(a.t, e1, e2, msgAndArgs...)
}

// Less ...
func (a *a) Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a.asserts.Less(a.t, e1, e2, msgAndArgs...)
}

// LessOrEqual ...
func (a *a) LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	a.asserts.LessOrEqual(a.t, e1, e2, msgAndArgs...)
}

// Implements ...
func (a *a) Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	a.asserts.Implements(a.t, interfaceObject, object, msgAndArgs...)
}

// Empty ...
func (a *a) Empty(object interface{}, msgAndArgs ...interface{}) {
	a.asserts.Empty(a.t, object, msgAndArgs...)
}

// NotEmpty ...
func (a *a) NotEmpty(object interface{}, msgAndArgs ...interface{}) {
	a.asserts.NotEmpty(a.t, object, msgAndArgs...)
}

// WithinDuration ...
func (a *a) WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	a.asserts.WithinDuration(a.t, expected, actual, delta, msgAndArgs...)
}

// JSONEq ...
func (a *a) JSONEq(expected, actual string, msgAndArgs ...interface{}) {
	a.asserts.JSONEq(a.t, expected, actual, msgAndArgs...)
}

// JSONContains ...
func (a *a) JSONContains(expected, actual string, msgAndArgs ...interface{}) {
	a.asserts.JSONContains(a.t, expected, actual, msgAndArgs...)
}

// Subset ...
func (a *a) Subset(list, subset interface{}, msgAndArgs ...interface{}) {
	a.asserts.Subset(a.t, list, subset, msgAndArgs...)
}

// IsType ...
func (a *a) IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	a.asserts.IsType(a.t, expectedType, object, msgAndArgs...)
}

// True ...
func (a *a) True(value bool, msgAndArgs ...interface{}) {
	a.asserts.True(a.t, value, msgAndArgs...)
}

// False ...
func (a *a) False(value bool, msgAndArgs ...interface{}) {
	a.asserts.False(a.t, value, msgAndArgs...)
}

// Regexp ...
func (a *a) Regexp(rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	a.asserts.Regexp(a.t, rx, str, msgAndArgs...)
}
