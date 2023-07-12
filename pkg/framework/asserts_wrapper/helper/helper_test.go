package helper

import (
	"fmt"
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

func TestAssertEqual_Success(t *testing.T) {
	mockT := newMock()
	NewAssertsHelper(mockT).Equal(1, 1)
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
	NewAssertsHelper(mockT).Equal(1, 2)
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
	NewAssertsHelper(mockT).NotEqual(1, 2)
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
	NewAssertsHelper(mockT).NotEqual(1, 1)
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

func TestAssertError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NewAssertsHelper(mockT).Error(err)
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
	NewAssertsHelper(mockT).Error(nil)
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
	NewAssertsHelper(mockT).NoError(nil)
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
	NewAssertsHelper(mockT).NoError(err)
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

func TestAssertNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NewAssertsHelper(mockT).NotNil(object)

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

	NewAssertsHelper(mockT).NotNil(nil)

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

	NewAssertsHelper(mockT).Nil(nil)

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

	NewAssertsHelper(mockT).Nil(object)

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
	NewAssertsHelper(mockT).Len(str, 4)

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

	NewAssertsHelper(mockT).Len(str, 4)

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
	NewAssertsHelper(mockT).NotContains(str, "4")

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

	NewAssertsHelper(mockT).NotContains(str, "est")

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
	NewAssertsHelper(mockT).Contains(str, "est")

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

	NewAssertsHelper(mockT).Contains(str, "4")

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

	NewAssertsHelper(mockT).Greater(test, 3)

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

	NewAssertsHelper(mockT).Greater(test, 5)

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

	NewAssertsHelper(mockT).GreaterOrEqual(test, 4)

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

	NewAssertsHelper(mockT).GreaterOrEqual(test, 5)

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

	NewAssertsHelper(mockT).Less(test, 4)

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

	NewAssertsHelper(mockT).Less(test, 5)

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

	NewAssertsHelper(mockT).LessOrEqual(test, 4)

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

	NewAssertsHelper(mockT).LessOrEqual(test, 5)

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

	NewAssertsHelper(mockT).Implements(ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*helper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*helper.testStructSuc(%#v)", ts), params[1].GetValue())

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

	NewAssertsHelper(mockT).Implements(ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*helper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*helper.testStructSuc(%#v)", ts), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	NewAssertsHelper(mockT).Empty(test)

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
	NewAssertsHelper(mockT).Empty(test)

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
	NewAssertsHelper(mockT).NotEmpty(test)

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
	NewAssertsHelper(mockT).NotEmpty(test)

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
	NewAssertsHelper(mockT).WithinDuration(test, test2, delta)

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
	NewAssertsHelper(mockT).WithinDuration(test, test2, delta)

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

	NewAssertsHelper(mockT).JSONEq(exp, exp)

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

	NewAssertsHelper(mockT).JSONEq(exp, actual)

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

	NewAssertsHelper(mockT).JSONContains(exp, actual)

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

	NewAssertsHelper(mockT).JSONContains(exp, actual)

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
	NewAssertsHelper(mockT).Subset(test, subset)

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
	NewAssertsHelper(mockT).Subset(test, subset)

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

	NewAssertsHelper(mockT).IsType(test, test)

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

	NewAssertsHelper(mockT).IsType(test, act)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*helper.testStruct(%#v)", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*helper.failTestStruct(%#v)", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertTrue_Success(t *testing.T) {
	mockT := newMock()

	NewAssertsHelper(mockT).True(true)

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

	NewAssertsHelper(mockT).True(false)

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

	NewAssertsHelper(mockT).False(false)

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

	NewAssertsHelper(mockT).False(true)

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
	NewAssertsHelper(mockT).Regexp(rx, str)

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
	NewAssertsHelper(mockT).Regexp(rx, str)

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

func TestRequireEqual_Success(t *testing.T) {
	mockT := newMock()
	NewRequireHelper(mockT).Equal(1, 1)
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
	NewRequireHelper(mockT).Equal(1, 2)
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
	NewRequireHelper(mockT).NotEqual(1, 2)
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
	NewRequireHelper(mockT).NotEqual(1, 1)
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

func TestRequireError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NewRequireHelper(mockT).Error(err)
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
	NewRequireHelper(mockT).Error(nil)
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
	NewRequireHelper(mockT).NoError(nil)
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
	NewRequireHelper(mockT).NoError(err)
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

func TestRequireNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NewRequireHelper(mockT).NotNil(object)

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

	NewRequireHelper(mockT).NotNil(nil)

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

	NewRequireHelper(mockT).Nil(nil)

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

	NewRequireHelper(mockT).Nil(object)

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
	NewRequireHelper(mockT).Len(str, 4)

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

	NewRequireHelper(mockT).Len(str, 4)

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
	NewRequireHelper(mockT).NotContains(str, "4")

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

	NewRequireHelper(mockT).NotContains(str, "est")

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
	NewRequireHelper(mockT).Contains(str, "est")

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

	NewRequireHelper(mockT).Contains(str, "4")

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

	NewRequireHelper(mockT).Greater(test, 3)

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

	NewRequireHelper(mockT).Greater(test, 5)

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

	NewRequireHelper(mockT).GreaterOrEqual(test, 4)

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

	NewRequireHelper(mockT).GreaterOrEqual(test, 5)

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

	NewRequireHelper(mockT).Less(test, 4)

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

	NewRequireHelper(mockT).Less(test, 5)

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

	NewRequireHelper(mockT).LessOrEqual(test, 4)

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

	NewRequireHelper(mockT).LessOrEqual(test, 5)

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

	NewRequireHelper(mockT).Implements(ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*helper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*helper.testStructSuc(%#v)", ts), params[1].GetValue())

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

	NewRequireHelper(mockT).Implements(ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*helper.testInterface(%#v)", ti), params[0].GetValue())
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*helper.testStructSuc(%#v)", ts), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	NewRequireHelper(mockT).Empty(test)

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
	NewRequireHelper(mockT).Empty(test)

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
	NewRequireHelper(mockT).NotEmpty(test)

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
	NewRequireHelper(mockT).NotEmpty(test)

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
	NewRequireHelper(mockT).WithinDuration(test, test2, delta)

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
	NewRequireHelper(mockT).WithinDuration(test, test2, delta)

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

	NewRequireHelper(mockT).JSONEq(exp, exp)

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

	NewRequireHelper(mockT).JSONEq(exp, actual)

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

	NewRequireHelper(mockT).JSONContains(exp, actual)

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

	NewRequireHelper(mockT).JSONContains(exp, actual)

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
	NewRequireHelper(mockT).Subset(test, subset)

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
	NewRequireHelper(mockT).Subset(test, subset)

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

	NewRequireHelper(mockT).IsType(test, test)

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

	NewRequireHelper(mockT).IsType(test, act)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*helper.testStruct(%#v)", test), params[0].GetValue())

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*helper.failTestStruct(%#v)", act), params[1].GetValue())

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireTrue_Success(t *testing.T) {
	mockT := newMock()

	NewRequireHelper(mockT).True(true)

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

	NewRequireHelper(mockT).True(false)

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

	NewRequireHelper(mockT).False(false)

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

	NewRequireHelper(mockT).False(true)

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
	NewRequireHelper(mockT).Regexp(rx, str)

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
	NewRequireHelper(mockT).Regexp(rx, str)

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
