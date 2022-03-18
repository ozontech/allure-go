package internal

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/stretchr/testify/require"
)

const (
	testErrMsg = `common.go:79: 
                        Error Trace:    wrapprer.go:55
                                                                steps.go:92
                                                                helper.go:25
                                                                wrapprer.go:51
                                                                helper.go:49
                                                                fails_test.go:70
                                                                steps.go:92
                                                                fails_test.go:69
                                                                actions.go:41
                                                                fails_test.go:68
                        Error:          Not equal: 
                                        expected: 1
                                        actual  : 2
                        Test:           TestRunDemo/TestFails/TestAssertionFailInnerSteps
                        Messages:       Failed inside step`
)

func TestExtractErrorMessages_default(t *testing.T) {
	require.Equal(t, "TestMessage", ExtractErrorMessages("TestMessage"))
}

func TestExtractErrorMessages_Messages(t *testing.T) {
	testString := `Messages:   TestMessage`
	require.Equal(t, "TestMessage", ExtractErrorMessages(testString))
}

func TestExtractErrorMessages(t *testing.T) {
	output := ExtractErrorMessages(testErrMsg)
	require.NotEmpty(t, output)
	require.Equal(t, "Failed inside step", output)
}

type mockT struct {
	contextName  string
	broken       bool
	errf         bool
	logf         bool
	failed       bool
	mockedResult *allure.Result
}

func (t2 *mockT) BreakResult(errMsg string) {
	t2.broken = true
	t2.mockedResult.Status = allure.Broken
	t2.mockedResult.SetStatusMessage(errMsg)
	t2.mockedResult.SetStatusTrace(errMsg)
}

func (t2 *mockT) Errorf(format string, msgAndArgs ...interface{}) {
	t2.errf = true
}

func (t2 *mockT) FailNow() {
	t2.failed = true
}

func (t2 *mockT) Logf(format string, msgAndArgs ...interface{}) {
	t2.logf = true
}

func (t2 *mockT) GetResult() *allure.Result {
	return t2.mockedResult
}

func TestTestError_TestBeforeEach(t *testing.T) {
	errMsg := "testErrMsg"
	mockedT := &mockT{
		contextName:  TestContextName,
		mockedResult: new(allure.Result),
	}

	TestError(mockedT.contextName, errMsg, mockedT)
	require.True(t, mockedT.broken)
	require.True(t, mockedT.failed)
	require.True(t, mockedT.errf)
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusMessage())
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusTrace())

	mockedT = &mockT{
		contextName:  BeforeEachContextName,
		mockedResult: new(allure.Result),
	}

	TestError(mockedT.contextName, errMsg, mockedT)
	require.True(t, mockedT.broken)
	require.True(t, mockedT.failed)
	require.True(t, mockedT.errf)
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusMessage())
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusTrace())
}

func TestTestError_AfterEachAll(t *testing.T) {
	errMsg := "testErrMsg"
	mockedT := &mockT{
		contextName:  AfterEachContextName,
		mockedResult: &allure.Result{StatusDetails: allure.StatusDetail{}},
	}

	TestError(mockedT.contextName, errMsg, mockedT)
	require.True(t, mockedT.logf)
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusMessage())
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusTrace())

	mockedT = &mockT{
		contextName:  AfterAllContextName,
		mockedResult: &allure.Result{StatusDetails: allure.StatusDetail{}},
	}

	TestError(mockedT.contextName, errMsg, mockedT)
	require.True(t, mockedT.logf)
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusMessage())
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusTrace())
}

func TestTestError_BeforeAll(t *testing.T) {
	errMsg := "testErrMsg"
	mockedT := &mockT{
		contextName:  BeforeAllContextName,
		mockedResult: &allure.Result{StatusDetails: allure.StatusDetail{}},
	}

	TestError(mockedT.contextName, errMsg, mockedT)
	require.True(t, mockedT.failed)
	require.True(t, mockedT.logf)
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusMessage())
	require.Equal(t, errMsg, mockedT.mockedResult.GetStatusTrace())
}
