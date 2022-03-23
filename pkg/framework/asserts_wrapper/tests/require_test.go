package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	r "github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/require"
	"github.com/ozontech/allure-go/pkg/framework/core/common"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

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

func getRequireT(testName string) provider.T {
	mockRealT := new(MockT)
	mockRealT.T = new(testing.T)
	mockT := common.NewT(mockRealT, "package", "TestSuite")
	mockT.Provider.NewTest("FakeTest", testName)
	mockT.Provider.TestContext()
	return mockT
}

func TestRequireEqual_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.Equal(mockT, 1, 1)
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

func TestRequireEqual_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.Equal(mockT, 1, 2)
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

func TestRequireNotEqual_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.NotEqual(mockT, 1, 2)
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

func TestRequireNotEqual_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.NotEqual(mockT, 1, 1)
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

func TestRequireError_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	err := errors.New("Some Error")

	r.Error(mockT, err)
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

func TestRequireError_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.Error(mockT, nil)
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

func TestRequireNoError_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.NoError(mockT, nil)
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

func TestRequireNoError_Fail(t *testing.T) {
	err := errors.New("Some Error")

	mockT := getRequireT(t.Name())

	r.NoError(mockT, err)
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

func TestRequireNotNil_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	object := struct{}{}

	r.NotNil(mockT, object)
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

func TestRequireNotNil_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.NotNil(mockT, nil)
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

func TestRequireNil_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.Nil(mockT, nil)
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

func TestRequireNil_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())
	object := struct{}{}

	r.Nil(mockT, object)
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

func TestRequireLen_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	str := "test"
	r.Len(mockT, str, 4)
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

func TestRequireLen_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())
	str := "test1"

	r.Len(mockT, str, 4)
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

func TestRequireNotContains_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	str := "test"
	r.NotContains(mockT, str, "4")
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

func TestRequireNotContains_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())
	str := "test"

	r.NotContains(mockT, str, "est")
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

func TestRequireContains_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	str := "test"
	r.Contains(mockT, str, "est")

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

func TestRequireContains_Failed(t *testing.T) {
	mockT := getRequireT(t.Name())
	str := "test"

	r.Contains(mockT, str, "4")
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

func TestRequireGreater_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 4

	r.Greater(mockT, test, 3)
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

func TestRequireGreater_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 4

	r.Greater(mockT, test, 5)
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

func TestRequireGreaterOrEqual_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 4

	r.GreaterOrEqual(mockT, test, 4)
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

func TestRequireGreaterOrEqual_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 4

	r.GreaterOrEqual(mockT, test, 5)
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

func TestRequireLess_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 3

	r.Less(mockT, test, 4)
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

func TestRequireLess_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 5

	r.Less(mockT, test, 5)
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

func TestRequireLesOrEqual_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 4

	r.LessOrEqual(mockT, test, 4)
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

func TestRequireLessOrEqual_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())
	test := 6

	r.LessOrEqual(mockT, test, 5)
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

type testRequireStructSuc struct {
}

func (t *testRequireStructSuc) test() {
}

func TestRequireImplements_Success(t *testing.T) {
	type testInterface interface {
		test()
	}

	mockT := getRequireT(t.Name())
	ti := new(testInterface)
	ts := &testRequireStructSuc{}

	r.Implements(mockT, ti, ts)
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
	require.Equal(t, fmt.Sprintf("*tests.testRequireStructSuc(%#v)", ts), params[1].Value)
}

func TestRequireImplements_Failed(t *testing.T) {
	type testInterface interface {
		test2()
	}

	mockT := getRequireT(t.Name())
	ti := new(testInterface)
	ts := &testRequireStructSuc{}

	r.Implements(mockT, ti, ts)
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
	require.Equal(t, fmt.Sprintf("*tests.testRequireStructSuc(%#v)", ts), params[1].Value)
}

func TestRequireEmpty_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := ""
	r.Empty(mockT, test)
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

func TestRequireEmpty_False(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := "123"
	r.Empty(mockT, test)
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

func TestRequireNotEmpty_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := "123"
	r.NotEmpty(mockT, test)
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

func TestRequireNotEmpty_False(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := ""
	r.NotEmpty(mockT, test)
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

func TestRequireWithDuration_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	r.WithinDuration(mockT, test, test2, delta)
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

func TestRequireWithDuration_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := time.Now()
	test2 := test.Add(100)
	delta := test2.Sub(test)
	test = test.Add(1000000)
	r.WithinDuration(mockT, test, test2, delta)
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

func TestRequireJSONEq_Success(t *testing.T) {
	mockT := getRequireT(t.Name())
	exp := "{\"key1\": 123, \"key2\": \"test\"}"

	r.JSONEq(mockT, exp, exp)
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

func TestRequireJSONEq_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())
	exp := "{\"key1\": 123, \"key2\": \"test\"}"
	actual := "{\"key1\": 1232, \"key2\": \"test2\"}"

	r.JSONEq(mockT, exp, actual)
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

func TestRequireSubset_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := []int{1, 2, 3}
	subset := []int{2, 3}
	r.Subset(mockT, test, subset)
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

func TestRequireSubset_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())

	test := []int{1, 2, 3}
	subset := []int{4, 3}
	r.Subset(mockT, test, subset)
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

func TestRequireIsType_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	type testStruct struct {
	}
	test := new(testStruct)

	r.IsType(mockT, test, test)
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

func TestRequireIsType_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())

	type testStruct struct {
	}
	type failTestStruct struct {
	}
	test := new(testStruct)
	act := new(failTestStruct)

	r.IsType(mockT, test, act)
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

func TestRequireTrue_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.True(mockT, true)
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

func TestRequireTrue_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.True(mockT, false)
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

func TestRequireFalse_Success(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.False(mockT, false)
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

func TestRequireFalse_Fail(t *testing.T) {
	mockT := getRequireT(t.Name())

	r.False(mockT, true)
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
