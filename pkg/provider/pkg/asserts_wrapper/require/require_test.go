package require

import (
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/common"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type resultInterface interface {
	GetResult() *allure.Result
	Provider() provider.Provider
}

type MockT struct {
	*testing.T
	failed bool
}

func (t *MockT) Failed() bool {
	return t.failed
}

func (t *MockT) FailNow() {
	t.failed = true
}

func (t *MockT) Errorf(format string, args ...interface{}) {
	_, _ = format, args
}

func getT(testName string) provider.T {
	mockRealT := new(MockT)
	mockRealT.T = new(testing.T)
	mockT := common.NewT(mockRealT, "TestSuite")
	mockT.(resultInterface).Provider().NewTest("FakeTest", testName)
	mockT.(resultInterface).Provider().TestContext()
	return mockT
}

func TestEqual_Success(t *testing.T) {
	mockT := getT(t.Name())

	Equal(mockT, 1, 1)
	require.False(t, mockT.Failed())
	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].Value)
}

func TestEqual_Failed(t *testing.T) {
	mockT := getT(t.Name())

	Equal(mockT, 1, 2)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)
	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].Value)
}

func TestNotEqual_Success(t *testing.T) {
	mockT := getT(t.Name())

	NotEqual(mockT, 1, 2)
	require.False(t, mockT.Failed())
	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "2", params[1].Value)
}

func TestNotEqual_Failed(t *testing.T) {
	mockT := getT(t.Name())

	NotEqual(mockT, 1, 1)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)
	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, "1", params[0].Value)
	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, "1", params[1].Value)
}

func TestError_Success(t *testing.T) {
	mockT := getT(t.Name())

	err := errors.New("Some Error")

	Error(mockT, err)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Error", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Contains(t, params[0].Value, err.Error())
}

func TestError_Fail(t *testing.T) {
	mockT := getT(t.Name())

	Error(mockT, nil)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, result.Status, allure.Failed)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Error", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)
}

func TestNoError_Success(t *testing.T) {
	mockT := getT(t.Name())

	NoError(mockT, nil)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: No Error", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)
}

func TestNoError_Fail(t *testing.T) {
	err := errors.New("Some Error")

	mockT := getT(t.Name())

	NoError(mockT, err)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, result.Status, allure.Failed)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: No Error", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Contains(t, params[0].Value, err.Error())
}

func TestNotNil_Success(t *testing.T) {
	mockT := getT(t.Name())
	object := struct{}{}

	NotNil(mockT, object)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].Value)
}

func TestNotNil_Failed(t *testing.T) {
	mockT := getT(t.Name())

	NotNil(mockT, nil)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)
}

func TestNil_Success(t *testing.T) {
	mockT := getT(t.Name())

	Nil(mockT, nil)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "<nil>", params[0].Value)
}

func TestNil_Failed(t *testing.T) {
	mockT := getT(t.Name())
	object := struct{}{}

	Nil(mockT, object)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Nil", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "struct {}(struct {}{})", params[0].Value)
}

func TestLen_Success(t *testing.T) {
	mockT := getT(t.Name())
	str := "test"
	Len(mockT, str, 4)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test\")", params[0].Value)
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].Value)
}

func TestLen_Failed(t *testing.T) {
	mockT := getT(t.Name())
	str := "test1"

	Len(mockT, str, 4)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Length", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Actual", params[0].Name)
	require.Equal(t, "string(\"test1\")", params[0].Value)
	require.Equal(t, "Expected Len", params[1].Name)
	require.Equal(t, "int(4)", params[1].Value)
}

func TestNotContains_Success(t *testing.T) {
	mockT := getT(t.Name())
	str := "test"
	NotContains(mockT, str, "4")
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Not Contains", params[1].Name)
	require.Equal(t, "\"4\"", params[1].Value)
}

func TestNotContains_Failed(t *testing.T) {
	mockT := getT(t.Name())
	str := "test"

	NotContains(mockT, str, "est")
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Not Contains", params[1].Name)
	require.Equal(t, "\"est\"", params[1].Value)
}

func TestContains_Success(t *testing.T) {
	mockT := getT(t.Name())
	str := "test"
	Contains(mockT, str, "est")

	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Contains", params[1].Name)
	require.Equal(t, "\"est\"", params[1].Value)
}

func TestContains_Failed(t *testing.T) {
	mockT := getT(t.Name())
	str := "test"

	Contains(mockT, str, "4")
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Contains", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Target Struct", params[0].Name)
	require.Equal(t, "\"test\"", params[0].Value)
	require.Equal(t, "Should Contains", params[1].Name)
	require.Equal(t, "\"4\"", params[1].Value)
}

func TestGreater_Success(t *testing.T) {
	mockT := getT(t.Name())
	test := 4

	Greater(mockT, test, 3)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "3", params[1].Value)
}

func TestGreater_Fail(t *testing.T) {
	mockT := getT(t.Name())
	test := 4

	Greater(mockT, test, 5)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)
}

func TestGreaterOrEqual_Success(t *testing.T) {
	mockT := getT(t.Name())
	test := 4

	GreaterOrEqual(mockT, test, 4)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].Value)
}

func TestGreaterOrEqual_Fail(t *testing.T) {
	mockT := getT(t.Name())
	test := 4

	GreaterOrEqual(mockT, test, 5)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Greater Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)
}

func TestLess_Success(t *testing.T) {
	mockT := getT(t.Name())
	test := 3

	Less(mockT, test, 4)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "3", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].Value)
}

func TestLess_Fail(t *testing.T) {
	mockT := getT(t.Name())
	test := 5

	Less(mockT, test, 5)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "5", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)
}

func TestLesOrEqual_Success(t *testing.T) {
	mockT := getT(t.Name())
	test := 4

	LessOrEqual(mockT, test, 4)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "4", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "4", params[1].Value)
}

func TestLessOrEqual_Fail(t *testing.T) {
	mockT := getT(t.Name())
	test := 6

	LessOrEqual(mockT, test, 5)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Less Or Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "First Element", params[0].Name)
	require.Equal(t, "6", params[0].Value)
	require.Equal(t, "Second Element", params[1].Name)
	require.Equal(t, "5", params[1].Value)
}

type testStructSuc struct {
}

func (t *testStructSuc) test() {
}

func TestImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := getT(t.Name())
	ti := new(testInterface)
	ts := &testStructSuc{}

	Implements(mockT, ti, ts)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.testStructSuc(%#v)", ts), params[1].Value)
}

func TestImplements_Failed(t *testing.T) {
	type testInterface interface {
		test2()
	}

	mockT := getT(t.Name())
	ti := new(testInterface)
	ts := &testStructSuc{}

	Implements(mockT, ti, ts)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Implements", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)
	require.Equal(t, "Interface Object", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.testStructSuc(%#v)", ts), params[1].Value)
}

func TestEmpty_Success(t *testing.T) {
	mockT := getT(t.Name())

	test := ""
	Empty(mockT, test)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].Value)
}

func TestEmpty_False(t *testing.T) {
	mockT := getT(t.Name())

	test := "123"
	Empty(mockT, test)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].Value)
}

func TestNotEmpty_Success(t *testing.T) {
	mockT := getT(t.Name())

	test := "123"
	NotEmpty(mockT, test)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"123\")", params[0].Value)
}

func TestNotEmpty_False(t *testing.T) {
	mockT := getT(t.Name())

	test := ""
	NotEmpty(mockT, test)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Not Empty", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)
	require.Equal(t, "Object", params[0].Name)
	require.Equal(t, "string(\"\")", params[0].Value)
}

func TestWithDuration_Success(t *testing.T) {
	mockT := getT(t.Name())

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	WithinDuration(mockT, test, test2, delta)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
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
}

func TestWithDuration_Fail(t *testing.T) {
	mockT := getT(t.Name())

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	WithinDuration(mockT, test, test2, delta)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
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
}

func TestJSONEq_Success(t *testing.T) {
	mockT := getT(t.Name())
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	JSONEq(mockT, exp, exp)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, exp, params[1].Value)
}

func TestJSONEq_Fail(t *testing.T) {
	mockT := getT(t.Name())
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	JSONEq(mockT, exp, actual)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: JSON Equal", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected", params[0].Name)
	require.Equal(t, exp, params[0].Value)

	require.Equal(t, "Actual", params[1].Name)
	require.Equal(t, actual, params[1].Value)
}

func TestSubset_Success(t *testing.T) {
	mockT := getT(t.Name())

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	Subset(mockT, test, subset)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].Value)
}

func TestSubset_Fail(t *testing.T) {
	mockT := getT(t.Name())

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	Subset(mockT, test, subset)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Subset", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "List", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Subset", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", subset), params[1].Value)
}

func TestIsType_Success(t *testing.T) {
	mockT := getT(t.Name())

	type testStruct struct {
	}
	test := new(testStruct)

	IsType(mockT, test, test)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[0].Value)

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("%#v", test), params[1].Value)
}

func TestIsType_Fail(t *testing.T) {
	mockT := getT(t.Name())

	type testStruct struct {
	}
	type failTestStruct struct {
	}
	test := new(testStruct)
	act := new(failTestStruct)

	IsType(mockT, test, act)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: Is Type", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 2)

	require.Equal(t, "Expected Type", params[0].Name)
	require.Equal(t, fmt.Sprintf("*require.testStruct(%#v)", test), params[0].Value)

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*require.failTestStruct(%#v)", act), params[1].Value)
}

func TestTrue_Success(t *testing.T) {
	mockT := getT(t.Name())

	True(mockT, true)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].Value)
}

func TestTrue_Fail(t *testing.T) {
	mockT := getT(t.Name())

	True(mockT, false)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: True", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].Value)
}

func TestFalse_Success(t *testing.T) {
	mockT := getT(t.Name())

	False(mockT, false)
	require.False(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Empty(t, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Passed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(false)", params[0].Value)
}

func TestFalse_Fail(t *testing.T) {
	mockT := getT(t.Name())

	False(mockT, true)
	require.True(t, mockT.Failed())

	result := mockT.(resultInterface).GetResult()
	steps := result.Steps
	require.Equal(t, allure.Failed, result.Status)
	require.Len(t, steps, 1)
	require.Equal(t, "REQUIRE: False", steps[0].Name)
	require.Equal(t, allure.Failed, steps[0].Status)

	params := steps[0].Parameters
	require.Len(t, params, 1)

	require.Equal(t, "Actual Value", params[0].Name)
	require.Equal(t, "bool(true)", params[0].Value)
}
