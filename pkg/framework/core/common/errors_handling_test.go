package common

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/stretchr/testify/require"
	"testing"
)

type errorTMock struct {
	logF    bool
	errorF  bool
	failNow bool
}

func (e *errorTMock) Errorf(format string, args ...interface{}) {
	e.errorF = true
}

func (e *errorTMock) Logf(format string, args ...interface{}) {
	e.logF = true
}

func (e *errorTMock) FailNow() {
	e.failNow = true
}

type errorProviderMock struct {
	status allure.Status
	msg    string
	trace  string
}

func (m *errorProviderMock) StopResult(status allure.Status) {
	m.status = status
}

func (m *errorProviderMock) UpdateResultStatus(msg string, trace string) {
	m.msg = msg
	m.trace = trace
}

func TestTestError_less100(t *testing.T) {
	errTMock := &errorTMock{}
	errProviderMock := &errorProviderMock{}

	TestError(errTMock, errProviderMock, constants.TestContextName, "errMsg")
	require.Equal(t, "errMsg", errProviderMock.msg)
	require.Equal(t, "errMsg", errProviderMock.trace)
	require.Equal(t, allure.Broken, errProviderMock.status)
	require.True(t, errTMock.errorF)
	require.True(t, errTMock.failNow)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}

	TestError(errTMock, errProviderMock, constants.BeforeEachContextName, "errMsg")
	require.Equal(t, "errMsg", errProviderMock.msg)
	require.Equal(t, "errMsg", errProviderMock.trace)
	require.Equal(t, allure.Broken, errProviderMock.status)
	require.True(t, errTMock.errorF)
	require.True(t, errTMock.failNow)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}

	TestError(errTMock, errProviderMock, constants.BeforeAllContextName, "errMsg")
	require.Equal(t, "errMsg", errProviderMock.msg)
	require.Equal(t, "errMsg", errProviderMock.trace)
	require.Empty(t, errProviderMock.status)
	require.True(t, errTMock.logF)
	require.True(t, errTMock.failNow)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}

	TestError(errTMock, errProviderMock, constants.AfterEachContextName, "errMsg")
	require.Equal(t, "errMsg", errProviderMock.msg)
	require.Equal(t, "errMsg", errProviderMock.trace)
	require.Empty(t, errProviderMock.status)
	require.True(t, errTMock.logF)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}

	TestError(errTMock, errProviderMock, constants.AfterAllContextName, "errMsg")
	require.Equal(t, "errMsg", errProviderMock.msg)
	require.Equal(t, "errMsg", errProviderMock.trace)
	require.Empty(t, errProviderMock.status)
	require.True(t, errTMock.logF)
}

func TestTestError_more100(t *testing.T) {
	errTMock := &errorTMock{}
	errProviderMock := &errorProviderMock{}
	errMsg := `errMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMserrMs`
	TestError(errTMock, errProviderMock, constants.TestContextName, errMsg)
	require.Equal(t, errMsg[:100], errProviderMock.msg)
	require.Equal(t, errMsg, errProviderMock.trace)
	require.Equal(t, allure.Broken, errProviderMock.status)
	require.True(t, errTMock.errorF)
	require.True(t, errTMock.failNow)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}
	TestError(errTMock, errProviderMock, constants.BeforeAllContextName, errMsg)
	require.Equal(t, errMsg[:100], errProviderMock.msg)
	require.Equal(t, errMsg, errProviderMock.trace)
	require.Empty(t, errProviderMock.status)
	require.True(t, errTMock.logF)
	require.True(t, errTMock.failNow)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}
	TestError(errTMock, errProviderMock, constants.BeforeEachContextName, errMsg)
	require.Equal(t, errMsg[:100], errProviderMock.msg)
	require.Equal(t, errMsg, errProviderMock.trace)
	require.Equal(t, allure.Broken, errProviderMock.status)
	require.True(t, errTMock.errorF)
	require.True(t, errTMock.failNow)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}
	TestError(errTMock, errProviderMock, constants.AfterEachContextName, errMsg)
	require.Equal(t, errMsg[:100], errProviderMock.msg)
	require.Equal(t, errMsg, errProviderMock.trace)
	require.Empty(t, errProviderMock.status)
	require.True(t, errTMock.logF)

	errTMock = &errorTMock{}
	errProviderMock = &errorProviderMock{}
	TestError(errTMock, errProviderMock, constants.AfterAllContextName, errMsg)
	require.Equal(t, errMsg[:100], errProviderMock.msg)
	require.Equal(t, errMsg, errProviderMock.trace)
	require.Empty(t, errProviderMock.status)
	require.True(t, errTMock.logF)
}
