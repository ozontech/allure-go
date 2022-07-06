package asserts

import (
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
)

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
	Equal(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertEqual_Fail(t *testing.T) {
	mockT := newMock()
	Equal(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotEqual_Success(t *testing.T) {
	mockT := newMock()
	NotEqual(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotEqual_Fail(t *testing.T) {
	mockT := newMock()
	NotEqual(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	Error(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertError_Fail(t *testing.T) {
	mockT := newMock()
	Error(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNoError_Success(t *testing.T) {
	mockT := newMock()
	NoError(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Passed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNoError_Fail(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NoError(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "ASSERT: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NotNil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotNil_Failed(t *testing.T) {
	mockT := newMock()

	NotNil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNil_Success(t *testing.T) {
	mockT := newMock()

	Nil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNil_Failed(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	Nil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertLen_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Length", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test\")", params[0].Value)
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertLen_Failed(t *testing.T) {
	mockT := newMock()
	str := "test1"

	Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Length", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test1\")", params[0].Value)
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NotContains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Not Contains", params[1].Name)
	require.Equal(t, "\"4\"", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NotContains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Not Contains", params[1].Name)
	require.Equal(t, "\"est\"", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	Contains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Contains", params[1].Name)
	require.Equal(t, "\"est\"", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	Contains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Contains", params[1].Name)
	require.Equal(t, "\"4\"", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertGreater_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	Greater(mockT, test, 3)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "3", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertGreater_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	Greater(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertGreaterOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	GreaterOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertGreaterOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	GreaterOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertLess_Success(t *testing.T) {
	mockT := newMock()
	test := 3

	Less(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "3", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertLess_Fail(t *testing.T) {
	mockT := newMock()
	test := 5

	Less(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "5", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertLesOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	LessOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertLessOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 6

	LessOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "6", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

type testStructSuc struct {
}

func (t *testStructSuc) test() {
}

func TestAssertImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*asserts.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*asserts.testStructSuc(%#v)", ts), params[1].Value)

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

	Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*asserts.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*asserts.testStructSuc(%#v)", ts), params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertEmpty_False(t *testing.T) {
	mockT := newMock()

	test := "123"
	Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertNotEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := "123"
	NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertNotEmpty_False(t *testing.T) {
	mockT := newMock()

	test := ""
	NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Not Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertWithDuration_Success(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Within Duration", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].Value)

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].Value)

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
	WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Within Duration", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 3)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, test.String(), params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, test2.String(), params[1].Value)

	require.Equal(t, "Delta", params[2].Name)
	require.Equal(t, delta.String(), params[2].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertJSONEq_Success(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	JSONEq(mockT, exp, exp)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, exp, params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertJSONEq_Fail(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	JSONEq(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertJSONContains_Success(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 123, "key3": ["foo", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertJSONContains_Fail(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 321, "key3": ["foobar", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertSubset_Success(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Subset", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertSubset_Fail(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Subset", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertIsType_Success(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	test := new(testStruct)

	IsType(mockT, test, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Is Type", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[1].Value)

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

	IsType(mockT, test, act)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*asserts.testStruct(%#v)", test), params[0].Value)

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*asserts.failTestStruct(%#v)", act), params[1].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertTrue_Success(t *testing.T) {
	mockT := newMock()

	True(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: True", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertTrue_Fail(t *testing.T) {
	mockT := newMock()

	True(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: True", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestAssertFalse_Success(t *testing.T) {
	mockT := newMock()

	False(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: False", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestAssertFalse_Fail(t *testing.T) {
	mockT := newMock()

	False(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "ASSERT: False", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].Value)

	require.True(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}
