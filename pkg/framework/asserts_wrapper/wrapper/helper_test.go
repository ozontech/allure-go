package wrapper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
)

type tMock struct {
	steps        []*allure.Step
	errorF       bool
	errorFString string
	failNow      bool
}

func newTMock() *tMock {
	return &tMock{steps: make([]*allure.Step, 0)}
}

func (p *tMock) Step(step *allure.Step) {
	p.steps = append(p.steps, step)
}

func (p *tMock) Errorf(format string, msgAndArgs ...interface{}) {
	p.errorFString = format
	p.errorF = true
}

func (p *tMock) FailNow() {
	p.failNow = true
}

func TestAssertHelper_getStepName(t *testing.T) {
	a := &assertHelper{}
	require.Equal(t, "ASSERT: Test", a.getStepName("Test"))

	b := &assertHelper{required: true}
	require.Equal(t, "REQUIRE: Test", b.getStepName("Test"))
}

func TestAssertHelper_withNewStep_requireFalse(t *testing.T) {
	a := &assertHelper{}
	mock := newTMock()
	param := allure.NewParameters("pName", "pValue")
	result := a.withNewStep(mock, mock, "Test", func(t TestingT) bool { return true }, param)
	require.True(t, result)
	require.NotEmpty(t, mock.steps)
	require.Len(t, mock.steps, 1)
	require.Equal(t, "ASSERT: Test", mock.steps[0].Name)
	require.Equal(t, allure.Passed, mock.steps[0].Status)
	require.NotEmpty(t, mock.steps[0].Parameters)
	require.Len(t, mock.steps[0].Parameters, 1)
	require.Equal(t, param[0].Name, mock.steps[0].Parameters[0].Name)
	require.Equal(t, param[0].Value, mock.steps[0].Parameters[0].Value)

	mock2 := newTMock()
	param2 := allure.NewParameter("pName", "pValue")
	result2 := a.withNewStep(mock2, mock2, "Test", func(t TestingT) bool { return false }, param)
	require.False(t, result2)
	require.NotEmpty(t, mock2.steps)
	require.Len(t, mock2.steps, 1)
	require.Equal(t, "ASSERT: Test", mock2.steps[0].Name)
	require.Equal(t, allure.Failed, mock2.steps[0].Status)
	require.NotEmpty(t, mock2.steps[0].Parameters)
	require.Len(t, mock2.steps[0].Parameters, 1)
	require.Equal(t, param2.Name, mock2.steps[0].Parameters[0].Name)
	require.Equal(t, param2.Value, mock2.steps[0].Parameters[0].Value)
}

func TestAssertHelper_withNewStep_requireTrue(t *testing.T) {
	a := &assertHelper{required: true}
	mock := newTMock()
	param := allure.NewParameters("pName", "pValue")
	result := a.withNewStep(mock, mock, "Test", func(t TestingT) bool { return true }, param)
	require.True(t, result)
	require.NotEmpty(t, mock.steps)
	require.Len(t, mock.steps, 1)
	require.Equal(t, "REQUIRE: Test", mock.steps[0].Name)
	require.Equal(t, allure.Passed, mock.steps[0].Status)
	require.NotEmpty(t, mock.steps[0].Parameters)
	require.Len(t, mock.steps[0].Parameters, 1)
	require.Equal(t, param[0].Name, mock.steps[0].Parameters[0].Name)
	require.Equal(t, param[0].Value, mock.steps[0].Parameters[0].Value)

	mock2 := newTMock()
	param2 := allure.NewParameter("pName", "pValue")
	result2 := a.withNewStep(mock2, mock2, "Test", func(t TestingT) bool { return false }, param)
	require.False(t, result2)
	require.NotEmpty(t, mock2.steps)
	require.Len(t, mock2.steps, 1)
	require.Equal(t, "REQUIRE: Test", mock2.steps[0].Name)
	require.Equal(t, allure.Failed, mock2.steps[0].Status)
	require.NotEmpty(t, mock2.steps[0].Parameters)
	require.Len(t, mock2.steps[0].Parameters, 1)
	require.Equal(t, param2.Name, mock2.steps[0].Parameters[0].Name)
	require.Equal(t, param2.Value, mock2.steps[0].Parameters[0].Value)
}
