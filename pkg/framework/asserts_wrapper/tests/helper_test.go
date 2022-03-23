package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type MockHelperT struct {
	*testing.T
	failed bool
}

func (t *MockHelperT) Failed() bool {
	return t.failed
}

func (t *MockHelperT) FailNow() {
	t.failed = true
}

func (t *MockHelperT) Errorf(format string, args ...interface{}) {
	_, _ = format, args
}

func getHelperT(testName string) provider.T {
	mockRealT := new(MockHelperT)
	mockRealT.T = new(testing.T)
	mockT := common.NewT(mockRealT, "package", "TestSuite")
	mockT.Provider.NewTest("FakeTest", testName)
	mockT.Provider.TestContext()
	return mockT
}

func TestHelperEqual_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).Equal(1, 1)
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

func TestHelperEqual_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).Equal(1, 2)
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

func TestHelperNotEqual_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).NotEqual(1, 2)
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

func TestHelperNotEqual_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).NotEqual(1, 1)
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

func TestHelperError_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	err := errors.New("Some Error")

	helper.NewRequireHelper(mockT).Error(err)
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

func TestHelperError_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).Error(nil)
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

func TestHelperNoError_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).NoError(nil)
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

func TestHelperNoError_Fail(t *testing.T) {
	err := errors.New("Some Error")

	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).NoError(err)
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

func TestHelperNotNil_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	object := struct{}{}

	helper.NewRequireHelper(mockT).NotNil(object)
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

func TestHelperNotNil_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).NotNil(nil)
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

func TestHelperNil_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).Nil(nil)
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

func TestHelperNil_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())
	object := struct{}{}

	helper.NewRequireHelper(mockT).Nil(object)
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

func TestHelperLen_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	str := "test"
	helper.NewRequireHelper(mockT).Len(str, 4)
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

func TestHelperLen_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())
	str := "test1"

	helper.NewRequireHelper(mockT).Len(str, 4)
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

func TestHelperNotContains_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	str := "test"
	helper.NewRequireHelper(mockT).NotContains(str, "4")
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

func TestHelperNotContains_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())
	str := "test"

	helper.NewRequireHelper(mockT).NotContains(str, "est")
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

func TestHelperContains_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	str := "test"
	helper.NewRequireHelper(mockT).Contains(str, "est")

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

func TestHelperContains_Failed(t *testing.T) {
	mockT := getHelperT(t.Name())
	str := "test"

	helper.NewRequireHelper(mockT).Contains(str, "4")
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

func TestHelperGreater_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 4

	helper.NewRequireHelper(mockT).Greater(test, 3)
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

func TestHelperGreater_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 4

	helper.NewRequireHelper(mockT).Greater(test, 5)
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

func TestHelperGreaterOrEqual_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 4

	helper.NewRequireHelper(mockT).GreaterOrEqual(test, 4)
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

func TestHelperGreaterOrEqual_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 4

	helper.NewRequireHelper(mockT).GreaterOrEqual(test, 5)
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

func TestHelperLess_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 3

	helper.NewRequireHelper(mockT).Less(test, 4)
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

func TestHelperLess_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 5

	helper.NewRequireHelper(mockT).Less(test, 5)
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

func TestHelperLesOrEqual_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 4

	helper.NewRequireHelper(mockT).LessOrEqual(test, 4)
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

func TestHelperLessOrEqual_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())
	test := 6

	helper.NewRequireHelper(mockT).LessOrEqual(test, 5)
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

type testHelperStructSuc struct {
}

func (t *testHelperStructSuc) test() {
}

func TestHelperImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := getHelperT(t.Name())
	ti := new(testInterface)
	ts := &testHelperStructSuc{}

	helper.NewRequireHelper(mockT).Implements(ti, ts)
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
	require.Equal(t, fmt.Sprintf("*tests.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*tests.testHelperStructSuc(%#v)", ts), params[1].Value)
}

func TestHelperImplements_Failed(t *testing.T) {
	type testInterface interface {
		test2()
	}

	mockT := getHelperT(t.Name())
	ti := new(testInterface)
	ts := &testHelperStructSuc{}

	helper.NewRequireHelper(mockT).Implements(ti, ts)
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
	require.Equal(t, fmt.Sprintf("*tests.testInterface(%#v)", ti), params[0].Value)
	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*tests.testHelperStructSuc(%#v)", ts), params[1].Value)
}

func TestHelperEmpty_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := ""
	helper.NewRequireHelper(mockT).Empty(test)
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

func TestHelperEmpty_False(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := "123"
	helper.NewRequireHelper(mockT).Empty(test)
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

func TestHelperNotEmpty_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := "123"
	helper.NewRequireHelper(mockT).NotEmpty(test)
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

func TestHelperNotEmpty_False(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := ""
	helper.NewRequireHelper(mockT).NotEmpty(test)
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

func TestHelperWithDuration_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	helper.NewRequireHelper(mockT).WithinDuration(test, test2, delta)
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

func TestHelperWithDuration_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	helper.NewRequireHelper(mockT).WithinDuration(test, test2, delta)
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

func TestHelperJSONEq_Success(t *testing.T) {
	mockT := getHelperT(t.Name())
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	helper.NewRequireHelper(mockT).JSONEq(exp, exp)
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

func TestHelperJSONEq_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	helper.NewRequireHelper(mockT).JSONEq(exp, actual)
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

func TestHelperSubset_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	helper.NewRequireHelper(mockT).Subset(test, subset)
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

func TestHelperSubset_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	helper.NewRequireHelper(mockT).Subset(test, subset)
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

func TestHelperIsType_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	type testStruct struct {
	}
	test := new(testStruct)

	helper.NewRequireHelper(mockT).IsType(test, test)
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

func TestHelperIsType_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())

	type testStruct struct {
	}
	type failTestStruct struct {
	}
	test := new(testStruct)
	act := new(failTestStruct)

	helper.NewRequireHelper(mockT).IsType(test, act)
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
	require.Equal(t, fmt.Sprintf("*tests.testStruct(%#v)", test), params[0].Value)

	require.Equal(t, "Object", params[1].Name)
	require.Equal(t, fmt.Sprintf("*tests.failTestStruct(%#v)", act), params[1].Value)
}

func TestHelperTrue_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).True(true)
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

func TestHelperTrue_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).True(false)
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

func TestHelperFalse_Success(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).False(false)
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

func TestHelperFalse_Fail(t *testing.T) {
	mockT := getHelperT(t.Name())

	helper.NewRequireHelper(mockT).False(true)
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
