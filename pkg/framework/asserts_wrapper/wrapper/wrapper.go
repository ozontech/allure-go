package wrapper

import (
	"bufio"
	"fmt"
	"reflect"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	coreAssert "github.com/ozontech/allure-go/pkg/framework/core/assert"
	"github.com/stretchr/testify/assert"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
	FailNow()
}

type asserts struct {
	t TestingT

	resultHelper *assertHelper
}

// NewAsserts inits new Assert interface
func NewAsserts(t TestingT) AssertsWrapper {
	return &asserts{
		t:            t,
		resultHelper: &assertHelper{},
	}
}

// NewRequire inits new Require interface
func NewRequire(t TestingT) AssertsWrapper {
	return &asserts{
		t:            t,
		resultHelper: &assertHelper{required: true},
	}
}

// Equal ...
func (a *asserts) Equal(provider Provider, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assertName := "Equal"
	expString, actString := formatUnequalValues(expected, actual)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Equal(a.t, expected, actual, msgAndArgs...) },
		allure.NewParameters("Expected", expString, "Actual", actString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// NotEqual ...
func (a *asserts) NotEqual(provider Provider, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Equal"
	expString, actString := formatUnequalValues(expected, actual)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.NotEqual(t, expected, actual, msgAndArgs...) },
		allure.NewParameters("Expected", expString, "Actual", actString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Error ...
func (a *asserts) Error(provider Provider, err error, msgAndArgs ...interface{}) {
	assertName := "Error"
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Error(t, err, msgAndArgs...) },
		allure.NewParameters("Actual", fmt.Sprintf("%+v", err)),
		msgAndArgs...,
	)

	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// NoError ...
func (a *asserts) NoError(provider Provider, err error, msgAndArgs ...interface{}) {
	assertName := "No Error"
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.NoError(t, err, msgAndArgs...) },
		allure.NewParameters("Actual", fmt.Sprintf("%+v", err)),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Nil ...
func (a *asserts) Nil(provider Provider, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Nil"
	_, objString := formatUnequalValues(nil, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Nil(t, object, msgAndArgs...) },
		allure.NewParameters("Actual", objString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// NotNil ...
func (a *asserts) NotNil(provider Provider, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Nil"
	_, objString := formatUnequalValues(nil, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.NotNil(t, object, msgAndArgs...) },
		allure.NewParameters("Actual", objString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Len ...
func (a *asserts) Len(provider Provider, object interface{}, length int, msgAndArgs ...interface{}) {
	assertName := "Length"
	lenString, objString := formatUnequalValues(length, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Len(t, object, length, msgAndArgs...) },
		allure.NewParameters("Actual", objString, "Expected Len", lenString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Contains ...
func (a *asserts) Contains(provider Provider, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	assertName := "Contains"
	sString, containsString := formatUnequalValues(s, contains)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Contains(t, s, contains, msgAndArgs...) },
		allure.NewParameters("Target Struct", sString, "Should Contains", containsString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// NotContains ...
func (a *asserts) NotContains(provider Provider, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Contains"
	sString, containsString := formatUnequalValues(s, contains)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.NotContains(t, s, contains, msgAndArgs...) },
		allure.NewParameters("Target Struct", sString, "Should Not Contains", containsString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Greater ...
func (a *asserts) Greater(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Greater"
	e1String, e2String := formatUnequalValues(e1, e2)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Greater(t, e1, e2, msgAndArgs...) },
		allure.NewParameters("First Element", e1String, "Second Element", e2String),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// GreaterOrEqual ...
func (a *asserts) GreaterOrEqual(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Greater Or Equal"
	e1String, e2String := formatUnequalValues(e1, e2)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.GreaterOrEqual(t, e1, e2, msgAndArgs...) },
		allure.NewParameters("First Element", e1String, "Second Element", e2String),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Less ...
func (a *asserts) Less(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Less"
	e1String, e2String := formatUnequalValues(e1, e2)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Less(t, e1, e2, msgAndArgs...) },
		allure.NewParameters("First Element", e1String, "Second Element", e2String),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// LessOrEqual ...
func (a *asserts) LessOrEqual(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Less Or Equal"
	e1String, e2String := formatUnequalValues(e1, e2)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.LessOrEqual(t, e1, e2, msgAndArgs...) },
		allure.NewParameters("First Element", e1String, "Second Element", e2String),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Implements ...
func (a *asserts) Implements(provider Provider, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Implements"
	interfaceObjectString, objectString := formatUnequalValues(interfaceObject, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Implements(t, interfaceObject, object, msgAndArgs...) },
		allure.NewParameters("Interface Object", interfaceObjectString, "Object", objectString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Empty ...
func (a *asserts) Empty(provider Provider, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Empty"
	_, objectString := formatUnequalValues(nil, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Empty(t, object, msgAndArgs...) },
		allure.NewParameters("Object", objectString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// NotEmpty ...
func (a *asserts) NotEmpty(provider Provider, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Empty"
	_, objectString := formatUnequalValues(nil, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.NotEmpty(t, object, msgAndArgs...) },
		allure.NewParameters("Object", objectString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// WithinDuration ...
func (a *asserts) WithinDuration(provider Provider, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	assertName := "Within Duration"

	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.WithinDuration(t, expected, actual, delta, msgAndArgs...) },
		allure.NewParameters("Expected", expected.String(), "Actual", actual.String(), "Delta", delta.String()),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// JSONEq ...
func (a *asserts) JSONEq(provider Provider, expected, actual string, msgAndArgs ...interface{}) {
	assertName := "JSON Equal"
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.JSONEq(t, expected, actual, msgAndArgs...) },
		allure.NewParameters("Expected", expected, "Actual", actual),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// JSONContains ...
func (a *asserts) JSONContains(provider Provider, expected, actual string, msgAndArgs ...interface{}) {
	assertName := "JSON Contains"
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return coreAssert.JSONContains(t, expected, actual, msgAndArgs...) },
		allure.NewParameters("Expected", expected, "Actual", actual),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Subset ...
func (a *asserts) Subset(provider Provider, list, subset interface{}, msgAndArgs ...interface{}) {
	assertName := "Subset"
	listString, subsetString := formatUnequalValues(list, subset)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Subset(t, list, subset, msgAndArgs...) },
		allure.NewParameters("List", listString, "Subset", subsetString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// IsType ...
func (a *asserts) IsType(provider Provider, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Is Type"
	expectedTypeString, objectString := formatUnequalValues(expectedType, object)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.IsType(t, expectedType, object, msgAndArgs...) },
		allure.NewParameters("Expected Type", expectedTypeString, "Object", objectString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// True ...
func (a *asserts) True(provider Provider, value bool, msgAndArgs ...interface{}) {
	assertName := "True"
	_, valueString := formatUnequalValues(nil, value)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.True(t, value, msgAndArgs...) },
		allure.NewParameters("Actual Value", valueString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// False ...
func (a *asserts) False(provider Provider, value bool, msgAndArgs ...interface{}) {
	assertName := "False"
	_, valueString := formatUnequalValues(nil, value)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.False(t, value, msgAndArgs...) },
		allure.NewParameters("Actual Value", valueString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// Regexp ...
func (a *asserts) Regexp(provider Provider, rx interface{}, str interface{}, msgAndArgs ...interface{}) {
	assertName := "Regexp"
	expString, actString := formatUnequalValues(rx, str)
	success := a.resultHelper.withNewStep(
		a.t,
		provider,
		assertName,
		func(t TestingT) bool { return assert.Regexp(a.t, rx, str, msgAndArgs...) },
		allure.NewParameters("Expected", expString, "Actual", actString),
		msgAndArgs...,
	)
	if !success && a.resultHelper.required {
		a.t.FailNow()
	}
}

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parenthesis similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual interface{}) (e string, a string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return fmt.Sprintf("%T(%s)", expected, truncatingFormat(expected)),
			fmt.Sprintf("%T(%s)", actual, truncatingFormat(actual))
	}
	switch expected.(type) {
	case time.Duration:
		return fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual)
	}
	return truncatingFormat(expected), truncatingFormat(actual)
}

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(data interface{}) string {
	value := fmt.Sprintf("%#v", data)
	max := bufio.MaxScanTokenSize - 100 // Give us some space the type info too if needed.
	if len(value) > max {
		value = value[0:max] + "<... truncated>"
	}
	return value
}
