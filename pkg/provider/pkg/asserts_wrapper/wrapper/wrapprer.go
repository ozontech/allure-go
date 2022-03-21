package wrapper

import (
	"bufio"
	"fmt"
	"reflect"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type asserts struct {
	resultHelper *assertHelper
}

// NewAsserts inits new Assert interface
func NewAsserts() AssertsWrapper {
	return &asserts{
		resultHelper: &assertHelper{},
	}
}

// NewRequire inits new Require interface
func NewRequire() AssertsWrapper {
	return &asserts{
		resultHelper: &assertHelper{required: true},
	}
}

// NewAssertsSubStep inits new Require interface for sub step
func NewAssertsSubStep(ctx provider.StepCtx) AssertsWrapper {
	return &asserts{
		resultHelper: &assertHelper{parentCtx: ctx},
	}
}

// NewRequireSubStep inits new Require interface for sub step
func NewRequireSubStep(ctx provider.StepCtx) AssertsWrapper {
	return &asserts{
		resultHelper: &assertHelper{parentCtx: ctx, required: true},
	}
}

// Equal ...
func (a *asserts) Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assertName := "Equal"
	expString, actString := formatUnequalValues(expected, actual)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Equal(ctx.T(), expected, actual, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Expected", expString),
		allure.NewParameter("Actual", actString),
	)
}

// NotEqual ...
func (a *asserts) NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Equal"
	expString, actString := formatUnequalValues(expected, actual)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.NotEqual(ctx.T(), expected, actual, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Expected", expString),
		allure.NewParameter("Actual", actString),
	)
}

// Error ...
func (a *asserts) Error(t ProviderT, err error, msgAndArgs ...interface{}) {
	assertName := "Error"
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Error(ctx.T(), err, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual", fmt.Sprintf("%+v", err)),
	)
}

// NoError ...
func (a *asserts) NoError(t ProviderT, err error, msgAndArgs ...interface{}) {
	assertName := "No Error"
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.NoError(ctx.T(), err, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual", fmt.Sprintf("%+v", err)),
	)
}

// Nil ...
func (a *asserts) Nil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Nil"
	_, objString := formatUnequalValues(nil, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Nil(ctx.T(), object, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual", objString),
	)
}

// NotNil ...
func (a *asserts) NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Nil"
	_, objString := formatUnequalValues(nil, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.NotNil(ctx.T(), object, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual", objString),
	)
}

// Len ...
func (a *asserts) Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{}) {
	assertName := "Length"
	lenString, objString := formatUnequalValues(length, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Len(ctx.T(), object, length, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual", objString),
		allure.NewParameter("Expected Len", lenString),
	)
}

// Contains ...
func (a *asserts) Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	assertName := "Contains"
	sString, containsString := formatUnequalValues(s, contains)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Contains(ctx.T(), s, contains, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Target Struct", sString),
		allure.NewParameter("Should Contains", containsString),
	)
}

// NotContains ...
func (a *asserts) NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Contains"
	sString, containsString := formatUnequalValues(s, contains)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.NotContains(ctx.T(), s, contains, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Target Struct", sString),
		allure.NewParameter("Should Not Contains", containsString),
	)
}

// Greater ...
func (a *asserts) Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Greater"
	e1String, e2String := formatUnequalValues(e1, e2)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Greater(ctx.T(), e1, e2, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("First Element", e1String),
		allure.NewParameter("Second Element", e2String),
	)
}

// GreaterOrEqual ...
func (a *asserts) GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Greater Or Equal"
	e1String, e2String := formatUnequalValues(e1, e2)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.GreaterOrEqual(ctx.T(), e1, e2, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("First Element", e1String),
		allure.NewParameter("Second Element", e2String),
	)
}

// Less ...
func (a *asserts) Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Less"
	e1String, e2String := formatUnequalValues(e1, e2)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Less(ctx.T(), e1, e2, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("First Element", e1String),
		allure.NewParameter("Second Element", e2String),
	)
}

// LessOrEqual ...
func (a *asserts) LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	assertName := "Less Or Equal"
	e1String, e2String := formatUnequalValues(e1, e2)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.LessOrEqual(ctx.T(), e1, e2, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("First Element", e1String),
		allure.NewParameter("Second Element", e2String),
	)
}

// Implements ...
func (a *asserts) Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Implements"
	interfaceObjectString, objectString := formatUnequalValues(interfaceObject, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Implements(ctx.T(), interfaceObject, object, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Interface Object", interfaceObjectString),
		allure.NewParameter("Object", objectString),
	)
}

// Empty ...
func (a *asserts) Empty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Empty"
	_, objectString := formatUnequalValues(nil, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Empty(ctx.T(), object, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Object", objectString),
	)
}

// NotEmpty ...
func (a *asserts) NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Not Empty"
	_, objectString := formatUnequalValues(nil, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.NotEmpty(ctx.T(), object, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Object", objectString),
	)
}

// WithinDuration ...
func (a *asserts) WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	assertName := "Within Duration"

	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.WithinDuration(ctx.T(), expected, actual, delta, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Expected", expected.String()),
		allure.NewParameter("Actual", actual.String()),
		allure.NewParameter("Delta", delta.String()),
	)
}

// JSONEq ...
func (a *asserts) JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	assertName := "JSON Equal"
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.JSONEq(ctx.T(), expected, actual, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Expected", expected),
		allure.NewParameter("Actual", actual),
	)
}

// Subset ...
func (a *asserts) Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	assertName := "Subset"
	listString, subsetString := formatUnequalValues(list, subset)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.Subset(ctx.T(), list, subset, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("List", listString),
		allure.NewParameter("Subset", subsetString),
	)
}

// IsType ...
func (a *asserts) IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	assertName := "Is Type"
	expectedTypeString, objectString := formatUnequalValues(expectedType, object)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.IsType(ctx.T(), expectedType, object, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Expected Type", expectedTypeString),
		allure.NewParameter("Object", objectString),
	)
}

// True ...
func (a *asserts) True(t ProviderT, value bool, msgAndArgs ...interface{}) {
	assertName := "True"
	_, valueString := formatUnequalValues(nil, value)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.True(ctx.T(), value, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual Value", valueString),
	)
}

// False ...
func (a *asserts) False(t ProviderT, value bool, msgAndArgs ...interface{}) {
	assertName := "False"
	_, valueString := formatUnequalValues(nil, value)
	a.resultHelper.withNewStep(
		t,
		a.resultHelper.getStepName(assertName),
		func(ctx provider.StepCtx) {
			result := assert.False(ctx.T(), value, msgAndArgs...)
			a.resultHelper.handleResult(ctx, result)
		},
		allure.NewParameter("Actual Value", valueString),
	)
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
