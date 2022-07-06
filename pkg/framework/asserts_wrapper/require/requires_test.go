package require

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

func TestRequireEqual_Success(t *testing.T) {
	mockT := newMock()
	Equal(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal", mockT.steps[0].Name)
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

func TestRequireEqual_Fail(t *testing.T) {
	mockT := newMock()
	Equal(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEqual_Success(t *testing.T) {
	mockT := newMock()
	NotEqual(mockT, 1, 2)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", mockT.steps[0].Name)
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

func TestRequireNotEqual_Fail(t *testing.T) {
	mockT := newMock()
	NotEqual(mockT, 1, 1)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireError_Success(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	Error(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error", mockT.steps[0].Name)
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

func TestRequireError_Fail(t *testing.T) {
	mockT := newMock()
	Error(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNoError_Success(t *testing.T) {
	mockT := newMock()
	NoError(mockT, nil)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: No Error", mockT.steps[0].Name)
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

func TestRequireNoError_Fail(t *testing.T) {
	mockT := newMock()
	err := errors.New("kek")
	NoError(mockT, err)
	require.Len(t, mockT.steps, 1)
	require.Equal(t, "REQUIRE: No Error", mockT.steps[0].Name)
	require.Equal(t, allure.Failed, mockT.steps[0].Status)

	params := mockT.steps[0].Parameters
	require.NotEmpty(t, params)
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, fmt.Sprintf("%+v", err), params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotNil_Success(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	NotNil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotNil_Failed(t *testing.T) {
	mockT := newMock()

	NotNil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNil_Success(t *testing.T) {
	mockT := newMock()

	Nil(mockT, nil)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNil_Failed(t *testing.T) {
	mockT := newMock()
	object := struct{}{}

	Nil(mockT, object)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLen_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
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

func TestRequireLen_Failed(t *testing.T) {
	mockT := newMock()
	str := "test1"

	Len(mockT, str, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test1\")", params[0].Value)
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	NotContains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
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

func TestRequireNotContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	NotContains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Not Contains", params[1].Name)
	require.Equal(t, "\"est\"", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireContains_Success(t *testing.T) {
	mockT := newMock()
	str := "test"
	Contains(mockT, str, "est")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
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

func TestRequireContains_Failed(t *testing.T) {
	mockT := newMock()
	str := "test"

	Contains(mockT, str, "4")

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Contains", params[1].Name)
	require.Equal(t, "\"4\"", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireGreater_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	Greater(mockT, test, 3)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
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

func TestRequireGreater_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	Greater(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireGreaterOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	GreaterOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
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

func TestRequireGreaterOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 4

	GreaterOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLess_Success(t *testing.T) {
	mockT := newMock()
	test := 3

	Less(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
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

func TestRequireLess_Fail(t *testing.T) {
	mockT := newMock()
	test := 5

	Less(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "5", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireLesOrEqual_Success(t *testing.T) {
	mockT := newMock()
	test := 4

	LessOrEqual(mockT, test, 4)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
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

func TestRequireLessOrEqual_Fail(t *testing.T) {
	mockT := newMock()
	test := 6

	LessOrEqual(mockT, test, 5)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "6", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

type testStructSuc struct {
}

func (t *testStructSuc) test() {
}

func TestRequireImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := newMock()
	ti := new(testInterface)
	ts := &testStructSuc{}

	Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.testStructSuc(%#v)", ts), params[1].Value)

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

	Implements(mockT, ti, ts)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.testStructSuc(%#v)", ts), params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := ""
	Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireEmpty_False(t *testing.T) {
	mockT := newMock()

	test := "123"
	Empty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireNotEmpty_Success(t *testing.T) {
	mockT := newMock()

	test := "123"
	NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireNotEmpty_False(t *testing.T) {
	mockT := newMock()

	test := ""
	NotEmpty(mockT, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireWithDuration_Success(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Within Duration", steps[0].Name)
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

func TestRequireWithDuration_Fail(t *testing.T) {
	mockT := newMock()

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	WithinDuration(mockT, test, test2, delta)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Within Duration", steps[0].Name)
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
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireJSONEq_Success(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	JSONEq(mockT, exp, exp)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
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

func TestRequireJSONEq_Fail(t *testing.T) {
	mockT := newMock()
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	JSONEq(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireJSONContains_Success(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 123, "key3": ["foo", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`
	JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Contains", steps[0].Name)
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

func TestRequireJSONContains_Fail(t *testing.T) {
	mockT := newMock()
	exp := `{"key1": 321, "key3": ["foobar", "bar"]}`
	actual := `{"key1": 123, "key2": "foobar", "key3": ["foo", "bar"]}`

	JSONContains(mockT, exp, actual)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireSubset_Success(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
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

func TestRequireSubset_Fail(t *testing.T) {
	mockT := newMock()

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	Subset(mockT, test, subset)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireIsType_Success(t *testing.T) {
	mockT := newMock()

	type testStruct struct {
	}
	test := new(testStruct)

	IsType(mockT, test, test)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
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

func TestRequireIsType_Fail(t *testing.T) {
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
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testStruct(%#v)", test), params[0].Value)

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.failTestStruct(%#v)", act), params[1].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireTrue_Success(t *testing.T) {
	mockT := newMock()

	True(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireTrue_Fail(t *testing.T) {
	mockT := newMock()

	True(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}

func TestRequireFalse_Success(t *testing.T) {
	mockT := newMock()

	False(mockT, false)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].Value)

	require.False(t, mockT.errorF)
	require.False(t, mockT.failNow)
	require.Empty(t, mockT.errorFString)
}

func TestRequireFalse_Fail(t *testing.T) {
	mockT := newMock()

	False(mockT, true)

	steps := mockT.steps
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].Value)

	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)
	require.Equal(t, "\n%s", mockT.errorFString)
}
