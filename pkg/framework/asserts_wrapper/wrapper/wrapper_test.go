package wrapper

import (
	"fmt"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
)

type testStructSuc struct {
}

func (t *testStructSuc) test() {
}

type providerTMock struct {
	steps        []*allure.Step
	errorF       bool
	errorFString string
	failNow      bool
}

func newMock() *providerTMock {
	return &providerTMock{steps: make([]*allure.Step, 0)}
}

func (p *providerTMock) Step(step *allure.Step) {
	p.steps = append(p.steps, step)
}

func (p *providerTMock) Errorf(format string, msgAndArgs ...interface{}) {
	p.errorFString = format
	p.errorF = true
}

func (p *providerTMock) FailNow() {
	p.failNow = true
}

func TestAssert_Decorate_Success(t *testing.T) {
	a := NewAsserts(t)
	dec, ok := a.(interface {
		Decorate(provider Provider, name string, assertFunc func(TestingT) bool, params []*allure.Parameter, msgAndArgs ...interface{})
	})
	require.True(t, ok)

	mockT := newMock()
	dec.Decorate(
		mockT,
		"TestDecorate",
		func(TestingT) bool { return true },
		allure.NewParameters("TestParam", "param_val"),
	)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: TestDecorate", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "TestParam", params[0].Name)
	require.Equal(t, "param_val", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssert_Decorate_Fail(t *testing.T) {
	mockT := newMock()
	a := NewAsserts(mockT)
	dec, ok := a.(interface {
		Decorate(provider Provider, name string, assertFunc func(TestingT) bool, params []*allure.Parameter, msgAndArgs ...interface{})
	})
	require.True(t, ok)

	assertFunc := func(t TestingT) bool {
		t.Errorf("\n%s", "err")
		return false
	}

	dec.Decorate(
		mockT,
		"TestDecorate",
		assertFunc,
		allure.NewParameters("TestParam", "param_val"),
	)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: TestDecorate", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "TestParam", params[0].Name)
	require.Equal(t, "param_val", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequire_Decorator_Success(t *testing.T) {
	mockT := newMock()
	a := NewRequire(mockT)
	dec, ok := a.(interface {
		Decorate(provider Provider, name string, assertFunc func(TestingT) bool, params []*allure.Parameter, msgAndArgs ...interface{})
	})
	require.True(t, ok)

	dec.Decorate(
		mockT,
		"TestDecorate",
		func(TestingT) bool { return true },
		allure.NewParameters("TestParam", "param_val"),
	)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: TestDecorate", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "TestParam", params[0].Name)
	require.Equal(t, "param_val", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssert_Decorate_Require(t *testing.T) {
	mockT := newMock()
	a := NewRequire(mockT)
	dec, ok := a.(interface {
		Decorate(provider Provider, name string, assertFunc func(TestingT) bool, params []*allure.Parameter, msgAndArgs ...interface{})
	})
	require.True(t, ok)

	assertFunc := func(t TestingT) bool {
		t.Errorf("\n%s", "err")
		return false
	}

	dec.Decorate(
		mockT,
		"TestDecorate",
		assertFunc,
		allure.NewParameters("TestParam", "param_val"),
	)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: TestDecorate", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "TestParam", params[0].Name)
	require.Equal(t, "param_val", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertExactly_Success(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).Exactly(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Exactly", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertExactly_Fail(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).Exactly(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Exactly", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertSame_Success(t *testing.T) {
	mockT := newMock()
	type someStr struct {
	}
	exp := &someStr{}
	act := exp
	NewAsserts(mockT).Same(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Same", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertSame_Fail(t *testing.T) {
	mockT := newMock()
	type someStr struct {
		someField string
	}
	exp := &someStr{}
	act := &someStr{}

	NewAsserts(mockT).Same(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Same", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotSame_Success(t *testing.T) {
	mockT := newMock()
	type someStr struct {
		someField string
	}
	exp := &someStr{}
	act := &someStr{}

	NewAsserts(mockT).NotSame(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Same", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotSame_Fail(t *testing.T) {
	mockT := newMock()
	type someStr struct {
	}
	exp := &someStr{}
	act := exp
	NewAsserts(mockT).NotSame(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Same", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertEqual_Success(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).Equal(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertEqual_Fail(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).Equal(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotEqual_Success(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).NotEqual(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotEqual_Fail(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).NotEqual(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertEqualValues_Success(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).EqualValues(mockT, uint32(123), int32(123))
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "uint32(0x7b)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "int32(123)", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertEqualValues_Fail(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).EqualValues(mockT, 1, "test")
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "int(1)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "string(\"test\")", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotEqualValues_Success(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).NotEqualValues(mockT, 1, "test")
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "int(1)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "string(\"test\")", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotEqualValues_Fail(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).NotEqualValues(mockT, uint32(123), int32(123))
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "uint32(0x7b)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "int32(123)", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NewAsserts(mockT).Error(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertError_Fail(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).Error(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNoError_Success(t *testing.T) {
	mockT := newMock()
	NewAsserts(mockT).NoError(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNoError_Fail(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NewAsserts(mockT).NoError(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertEqualError_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := errors.New(exp)
	NewAsserts(mockT).EqualError(mockT, err, exp)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, err.Error(), params[0].GetValue())
	require.Equal(t, "Expected", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertEqualError_Fail(t *testing.T) {
	mockT := newMock()
	exp := "testErr2"
	actual := "testErr"
	err := errors.New(actual)
	NewAsserts(mockT).EqualError(mockT, err, exp)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, err.Error(), params[0].GetValue())
	require.Equal(t, "Expected", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertErrorIs_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := fmt.Errorf(exp)
	errNew := errors.Wrap(err, "NewMessage")
	NewAsserts(mockT).ErrorIs(mockT, errNew, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error Is", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, errNew.Error(), params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

type fakeError struct {
	input string
}

func (f *fakeError) Error() string {
	return fmt.Sprintf("fake error: %s", f.input)
}

func TestAssertErrorIs_Fail(t *testing.T) {
	mockT := newMock()

	var err = fakeError{"some"}
	errNew := errors.Wrap(fmt.Errorf("other"), "NewMessage")
	NewAsserts(mockT).ErrorIs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error Is", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, "NewMessage: other", params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, "fake error: some", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertErrorAs_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := fmt.Errorf(exp)
	errNew := errors.Wrap(err, "NewMessage")
	NewAsserts(mockT).ErrorAs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error As", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, errNew.Error(), params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertErrorAs_Fail(t *testing.T) {
	mockT := newMock()

	var err *fakeError
	errNew := errors.Wrap(fmt.Errorf("other"), "NewMessage")
	NewAsserts(mockT).ErrorAs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error As", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, "NewMessage: other", params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, fmt.Sprintf("**wrapper.fakeError((**wrapper.fakeError)(%+v))", &err), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NewAsserts(mockT).NotNil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotNil_Failed(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).NotNil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNil_Success(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).Nil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNil_Failed(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NewAsserts(mockT).Nil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertLen_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NewAsserts(mockT).Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Length", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test\")", params[0].GetValue())
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertLen_Failed(t *testing.T) {
	mockT := newMock()
	str := "test1"

	NewAsserts(mockT).Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Length", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test1\")", params[0].GetValue())
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NewAsserts(mockT).NotContains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Not Contain", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NewAsserts(mockT).NotContains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Not Contain", params[1].Name)
	require.Equal(t, "est", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NewAsserts(mockT).Contains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Contain", params[1].Name)
	require.Equal(t, "est", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NewAsserts(mockT).Contains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Contain", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertGreater_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	NewAsserts(mockT).Greater(mockT, test, 3)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "3", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertGreater_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	NewAsserts(mockT).Greater(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertGreaterOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	NewAsserts(mockT).GreaterOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertGreaterOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	NewAsserts(mockT).GreaterOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertLess_Success(t *testing.T) {
	mockT := newMock()
	test := 3

	NewAsserts(mockT).Less(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "3", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertLess_Fail(t *testing.T) {
	mockT := newMock()
	test := 5

	NewAsserts(mockT).Less(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "5", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertLesOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	NewAsserts(mockT).LessOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertLessOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 6

	NewAsserts(mockT).LessOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "6", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	NewAsserts(mockT).Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testStructSuc(%#v)", ts), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertImplements_Failed(t *testing.T) {
	type testInterface interface {
		test2()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	NewAsserts(mockT).Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testStructSuc(%#v)", ts), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	NewAsserts(mockT).Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertEmpty_False(t *testing.T) {
	mockT := newMock()

	test := "123"
	NewAsserts(mockT).Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := "123"
	NewAsserts(mockT).NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotEmpty_False(t *testing.T) {
	mockT := newMock()

	test := ""
	NewAsserts(mockT).NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertWithDuration_Success(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	NewAsserts(mockT).WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Within Duration", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertWithDuration_Fail(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	NewAsserts(mockT).WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Within Duration", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertJSONEq_Success(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	NewAsserts(mockT).JSONEq(mockT, exp, exp)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertJSONEq_Fail(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	NewAsserts(mockT).JSONEq(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertJSONContains_Success(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 123, "key3": ["foo", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	NewAsserts(mockT).JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertJSONContains_Fail(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 321, "key3": ["foobar", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	NewAsserts(mockT).JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertSubset_Success(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	NewAsserts(mockT).Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Subset", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertSubset_Fail(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	NewAsserts(mockT).Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Subset", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertIsType_Success(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	test := new(testStruct)

	NewAsserts(mockT).IsType(mockT, test, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Is Type", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertIsType_Fail(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	type failTestStruct struct {
	}
	test := new(testStruct)
	act := new(failTestStruct)

	NewAsserts(mockT).IsType(mockT, test, act)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testStruct(%#v)", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.failTestStruct(%#v)", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertTrue_Success(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).True(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: True", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertTrue_Fail(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).True(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: True", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertFalse_Success(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).False(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: False", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertFalse_Fail(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).False(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: False", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertRegexp_Success(t *testing.T) {
	mockT := newMock()

	rx := `^start`
	str := "start of the line"
	NewAsserts(mockT).Regexp(mockT, rx, str)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Regexp", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, rx, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, str, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertRegexp_Failed(t *testing.T) {
	mockT := newMock()

	rx := `^end`
	str := "start of the line"
	NewAsserts(mockT).Regexp(mockT, rx, str)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Regexp", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, rx, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, str, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertElementsMatch_Success(t *testing.T) {
	mockT := newMock()

	listA := []int{1, 2, 3}
	listB := []int{1, 2, 3}
	NewAsserts(mockT).ElementsMatch(mockT, listA, listB)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Elements Match", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "ListA", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", listA), params[0].GetValue())

	require.Equal(t, "ListB", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", listB), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertElementsMatch_Fail(t *testing.T) {
	mockT := newMock()

	listA := []int{1, 2, 3}
	listB := []int{4, 3}
	NewAsserts(mockT).ElementsMatch(mockT, listA, listB)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Elements Match", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "ListA", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", listA), params[0].GetValue())

	require.Equal(t, "ListB", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", listB), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertDirExists_Success(t *testing.T) {
	dirName := "test"
	err := os.Mkdir(dirName, 0644)
	require.NoError(t, err, "Can't create folder to begin test")
	defer os.RemoveAll(dirName)

	mockT := newMock()
	NewAsserts(mockT).DirExists(mockT, dirName)
	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Dir Exists", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Path", params[0].Name)
	require.Equal(t, dirName, params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertDirExists_Fail(t *testing.T) {
	dirName := "test"

	mockT := newMock()
	NewAsserts(mockT).DirExists(mockT, dirName)
	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Dir Exists", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Path", params[0].Name)
	require.Equal(t, dirName, params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertCondition_Success(t *testing.T) {
	test := false
	conditionFunc := func() bool {
		test = true
		return test
	}
	mockT := newMock()
	NewAsserts(mockT).Condition(mockT, conditionFunc)
	steps := mockT.steps
	require.True(t, test)
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Condition", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Signature", params[0].Name)
	require.Equal(t, "assert.Comparison", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertCondition_Fail(t *testing.T) {
	test := false
	conditionFunc := func() bool {
		test = true
		return !test
	}
	mockT := newMock()
	NewAsserts(mockT).Condition(mockT, conditionFunc)
	steps := mockT.steps
	require.True(t, test)
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Condition", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Signature", params[0].Name)
	require.Equal(t, "assert.Comparison", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertZero_Success(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).Zero(mockT, 0)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Zero", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "0", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertZero_Fail(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).Zero(mockT, 1)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Zero", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotZero_Success(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).NotZero(mockT, 1)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Zero", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotZero_Fail(t *testing.T) {
	mockT := newMock()

	NewAsserts(mockT).NotZero(mockT, 0)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Zero", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "0", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertInDelta_Success(t *testing.T) {
	mockT := newMock()

	expected := 10.1
	actual := 9.9
	delta := 0.2
	NewAsserts(mockT).InDelta(mockT, expected, actual, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: In Delta", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", expected), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", actual), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, fmt.Sprintf("%v", delta), params[2].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertInDelta_Fail(t *testing.T) {
	mockT := newMock()

	expected := 10
	actual := 11.1
	delta := float64(1)
	NewAsserts(mockT).InDelta(mockT, expected, actual, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: In Delta", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", expected), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", actual), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, fmt.Sprintf("%v", delta), params[2].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireExactly_Success(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).Exactly(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Exactly", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireExactly_Fail(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).Exactly(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Exactly", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireSame_Success(t *testing.T) {
	mockT := newMock()
	type someStr struct {
	}
	exp := &someStr{}
	act := exp
	NewRequire(mockT).Same(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Same", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireSame_Fail(t *testing.T) {
	mockT := newMock()
	type someStr struct {
		someField string
	}
	exp := &someStr{}
	act := &someStr{}

	NewRequire(mockT).Same(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Same", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotSame_Success(t *testing.T) {
	mockT := newMock()
	type someStr struct {
		someField string
	}
	exp := &someStr{}
	act := &someStr{}

	NewRequire(mockT).NotSame(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Same", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotSame_Fail(t *testing.T) {
	mockT := newMock()
	type someStr struct {
	}
	exp := &someStr{}
	act := exp
	NewRequire(mockT).NotSame(mockT, exp, act)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Same", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%p", exp), params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%p", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEqual_Success(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).Equal(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqual_Fail(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).Equal(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEqual_Success(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).NotEqual(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEqual_Fail(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).NotEqual(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEqualValues_Success(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).EqualValues(mockT, uint32(123), int32(123))
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "uint32(0x7b)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "int32(123)", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqualValues_Fail(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).EqualValues(mockT, 1, "test")
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "int(1)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "string(\"test\")", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEqualValues_Success(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).NotEqualValues(mockT, 1, "test")
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "int(1)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "string(\"test\")", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEqualValues_Fail(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).NotEqualValues(mockT, uint32(123), int32(123))
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal Values", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "uint32(0x7b)", params[0].GetValue())
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "int32(123)", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NewRequire(mockT).Error(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireError_Fail(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).Error(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNoError_Success(t *testing.T) {
	mockT := newMock()
	NewRequire(mockT).NoError(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNoError_Fail(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NewRequire(mockT).NoError(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEqualError_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := errors.New(exp)
	NewRequire(mockT).EqualError(mockT, err, exp)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, err.Error(), params[0].GetValue())
	require.Equal(t, "Expected", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEqualError_Fail(t *testing.T) {
	mockT := newMock()
	exp := "testErr2"
	actual := "testErr"
	err := errors.New(actual)
	NewRequire(mockT).EqualError(mockT, err, exp)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, err.Error(), params[0].GetValue())
	require.Equal(t, "Expected", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireErrorIs_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := fmt.Errorf(exp)
	errNew := errors.Wrap(err, "NewMessage")
	NewRequire(mockT).ErrorIs(mockT, errNew, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error Is", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, errNew.Error(), params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireErrorIs_Fail(t *testing.T) {
	mockT := newMock()

	var err = fakeError{"some"}
	errNew := errors.Wrap(fmt.Errorf("other"), "NewMessage")
	NewRequire(mockT).ErrorIs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error Is", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, "NewMessage: other", params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, "fake error: some", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireErrorAs_Success(t *testing.T) {
	mockT := newMock()
	exp := "testErr"
	err := fmt.Errorf(exp)
	errNew := errors.Wrap(err, "NewMessage")
	NewRequire(mockT).ErrorAs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error As", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, errNew.Error(), params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireErrorAs_Fail(t *testing.T) {
	mockT := newMock()

	var err *fakeError
	errNew := errors.Wrap(fmt.Errorf("other"), "NewMessage")
	NewRequire(mockT).ErrorAs(mockT, errNew, &err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error As", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Error", params[0].Name)
	require.Equal(t, "NewMessage: other", params[0].GetValue())
	require.Equal(t, "Target", params[1].Name)
	require.Equal(t, fmt.Sprintf("**wrapper.fakeError((**wrapper.fakeError)(%+v))", &err), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NewRequire(mockT).NotNil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotNil_Failed(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).NotNil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNil_Success(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).Nil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNil_Failed(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NewRequire(mockT).Nil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLen_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NewRequire(mockT).Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test\")", params[0].GetValue())
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireLen_Failed(t *testing.T) {
	mockT := newMock()
	str := "test1"

	NewRequire(mockT).Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test1\")", params[0].GetValue())
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NewRequire(mockT).NotContains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Not Contain", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NewRequire(mockT).NotContains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Not Contain", params[1].Name)
	require.Equal(t, "est", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NewRequire(mockT).Contains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Contain", params[1].Name)
	require.Equal(t, "est", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NewRequire(mockT).Contains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "test", params[0].GetValue())
	require.Equal(t, "Should Contain", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireGreater_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	NewRequire(mockT).Greater(mockT, test, 3)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "3", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireGreater_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	NewRequire(mockT).Greater(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireGreaterOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	NewRequire(mockT).GreaterOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireGreaterOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	NewRequire(mockT).GreaterOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLess_Success(t *testing.T) {
	mockT := newMock()
	test := 3

	NewRequire(mockT).Less(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "3", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireLess_Fail(t *testing.T) {
	mockT := newMock()
	test := 5

	NewRequire(mockT).Less(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "5", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLesOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	NewRequire(mockT).LessOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireLessOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 6

	NewRequire(mockT).LessOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "6", params[0].GetValue())
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	NewRequire(mockT).Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testStructSuc(%#v)", ts), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireImplements_Failed(t *testing.T) {
	type testInterface interface {
		test2()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	NewRequire(mockT).Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testStructSuc(%#v)", ts), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	NewRequire(mockT).Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEmpty_False(t *testing.T) {
	mockT := newMock()

	test := "123"
	NewRequire(mockT).Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := "123"
	NewRequire(mockT).NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEmpty_False(t *testing.T) {
	mockT := newMock()

	test := ""
	NewRequire(mockT).NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireWithDuration_Success(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	NewRequire(mockT).WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Within Duration", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireWithDuration_Fail(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	NewRequire(mockT).WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Within Duration", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireJSONEq_Success(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	NewRequire(mockT).JSONEq(mockT, exp, exp)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, exp, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireJSONEq_Fail(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	NewRequire(mockT).JSONEq(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireJSONContains_Success(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 123, "key3": ["foo", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	NewRequire(mockT).JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireJSONContains_Fail(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 321, "key3": ["foobar", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	NewRequire(mockT).JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireSubset_Success(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	NewRequire(mockT).Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireSubset_Fail(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	NewRequire(mockT).Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireIsType_Success(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	test := new(testStruct)

	NewRequire(mockT).IsType(mockT, test, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireIsType_Fail(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	type failTestStruct struct {
	}
	test := new(testStruct)
	act := new(failTestStruct)

	NewRequire(mockT).IsType(mockT, test, act)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.testStruct(%#v)", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*wrapper.failTestStruct(%#v)", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireTrue_Success(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).True(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireTrue_Fail(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).True(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireFalse_Success(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).False(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireFalse_Fail(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).False(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireRegexp_Success(t *testing.T) {
	mockT := newMock()

	rx := `^start`
	str := "start of the line"
	NewRequire(mockT).Regexp(mockT, rx, str)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Regexp", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, rx, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, str, params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireRegexp_Failed(t *testing.T) {
	mockT := newMock()

	rx := `^end`
	str := "start of the line"
	NewRequire(mockT).Regexp(mockT, rx, str)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Regexp", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, rx, params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, str, params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireElementsMatch_Success(t *testing.T) {
	mockT := newMock()

	listA := []int{1, 2, 3}
	listB := []int{1, 2, 3}
	NewRequire(mockT).ElementsMatch(mockT, listA, listB)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Elements Match", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "ListA", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", listA), params[0].GetValue())

	require.Equal(t, "ListB", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", listB), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireElementsMatch_Fail(t *testing.T) {
	mockT := newMock()

	listA := []int{1, 2, 3}
	listB := []int{4, 3}
	NewRequire(mockT).ElementsMatch(mockT, listA, listB)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Elements Match", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "ListA", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", listA), params[0].GetValue())

	require.Equal(t, "ListB", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", listB), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireDirExists_Success(t *testing.T) {
	dirName := "test"
	err := os.Mkdir(dirName, 0644)
	require.NoError(t, err, "Can't create folder to begin test")
	defer os.RemoveAll(dirName)

	mockT := newMock()
	NewRequire(mockT).DirExists(mockT, dirName)
	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Dir Exists", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Path", params[0].Name)
	require.Equal(t, dirName, params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireDirExists_Fail(t *testing.T) {
	dirName := "test"

	mockT := newMock()
	NewRequire(mockT).DirExists(mockT, dirName)
	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Dir Exists", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Path", params[0].Name)
	require.Equal(t, dirName, params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireCondition_Success(t *testing.T) {
	test := false
	conditionFunc := func() bool {
		test = true
		return test
	}
	mockT := newMock()
	NewRequire(mockT).Condition(mockT, conditionFunc)
	steps := mockT.steps
	require.True(t, test)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Condition", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Signature", params[0].Name)
	require.Equal(t, "assert.Comparison", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireCondition_Fail(t *testing.T) {
	test := false
	conditionFunc := func() bool {
		test = true
		return !test
	}
	mockT := newMock()
	NewRequire(mockT).Condition(mockT, conditionFunc)
	steps := mockT.steps
	require.True(t, test)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Condition", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Signature", params[0].Name)
	require.Equal(t, "assert.Comparison", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireZero_Success(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).Zero(mockT, 0)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Zero", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "0", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireZero_Fail(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).Zero(mockT, 1)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Zero", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotZero_Success(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).NotZero(mockT, 1)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Zero", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "1", params[0].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotZero_Fail(t *testing.T) {
	mockT := newMock()

	NewRequire(mockT).NotZero(mockT, 0)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Zero", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Target", params[0].Name)
	require.Equal(t, "0", params[0].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireInDelta_Success(t *testing.T) {
	mockT := newMock()

	expected := 10.1
	actual := 9.9
	delta := 0.2
	NewRequire(mockT).InDelta(mockT, expected, actual, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: In Delta", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", expected), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", actual), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, fmt.Sprintf("%v", delta), params[2].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireInDelta_Fail(t *testing.T) {
	mockT := newMock()

	expected := 10
	actual := 11.1
	delta := float64(1)
	NewRequire(mockT).InDelta(mockT, expected, actual, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: In Delta", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", expected), params[0].GetValue())

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", actual), params[1].GetValue())

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, fmt.Sprintf("%v", delta), params[2].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertsEventually_Success(t *testing.T) {
	mockT := newMock()

	var (
		counter atomic.Int32
		waitFor = time.Second
		tick    = 10 * time.Millisecond
	)
	NewAsserts(mockT).Eventually(mockT, func() bool {
		if counter.Add(1) < 3 {
			time.Sleep(20 * time.Millisecond)
			return false
		}
		return true
	}, waitFor, tick)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Eventually", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "WaitFor", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", waitFor), params[0].GetValue())

	require.Equal(t, "Tick", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", tick), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertsEventually_Fail(t *testing.T) {
	mockT := newMock()

	var (
		counter atomic.Int32
		waitFor = 10 * time.Millisecond
		tick    = 10 * time.Millisecond
	)
	NewAsserts(mockT).Eventually(mockT, func() bool {
		if counter.Add(1) < 3 {
			time.Sleep(20 * time.Millisecond)
			return false
		}
		return true
	}, waitFor, tick)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Eventually", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "WaitFor", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", waitFor), params[0].GetValue())

	require.Equal(t, "Tick", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", tick), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEventually_Success(t *testing.T) {
	mockT := newMock()

	var (
		counter atomic.Int32
		waitFor = time.Second
		tick    = 10 * time.Millisecond
	)
	NewRequire(mockT).Eventually(mockT, func() bool {
		if counter.Add(1) < 3 {
			time.Sleep(20 * time.Millisecond)
			return false
		}
		return true
	}, waitFor, tick)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Eventually", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "WaitFor", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", waitFor), params[0].GetValue())

	require.Equal(t, "Tick", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", tick), params[1].GetValue())

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEventually_Fail(t *testing.T) {
	mockT := newMock()

	var (
		counter atomic.Int32
		waitFor = 10 * time.Millisecond
		tick    = 10 * time.Millisecond
	)
	NewRequire(mockT).Eventually(mockT, func() bool {
		if counter.Add(1) < 3 {
			time.Sleep(20 * time.Millisecond)
			return false
		}
		return true
	}, waitFor, tick)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Eventually", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "WaitFor", params[0].Name)
	require.Equal(t, fmt.Sprintf("%v", waitFor), params[0].GetValue())

	require.Equal(t, "Tick", params[1].Name)
	require.Equal(t, fmt.Sprintf("%v", tick), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}
