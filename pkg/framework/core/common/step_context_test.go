package common

import (
	"sync"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"

	"github.com/stretchr/testify/require"
)

type providerTMockStep struct {
	steps   []*allure.Step
	error   bool
	errorF  bool
	log     bool
	logf    bool
	failNow bool
	failed  bool
	name    string

	testingT provider.TestingT
}

func (m *providerTMockStep) Break(args ...interface{}) {
	m.failed = true
}

func (m *providerTMockStep) Breakf(format string, args ...interface{}) {
	m.failed = true
}

func (m *providerTMockStep) Broken() {
	m.failed = true
}

func (m *providerTMockStep) BrokenNow() {
	m.failed = true
}

func newStepProviderMock() *providerTMockStep {
	return &providerTMockStep{steps: []*allure.Step{}}
}

func (m *providerTMockStep) Step(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *providerTMockStep) Errorf(format string, args ...interface{}) {
	m.errorF = true
	m.failed = true
}

func (m *providerTMockStep) FailNow() {
	m.failNow = true
	m.failed = true
}

func (m *providerTMockStep) Fail() {
	m.failed = true
}

func (m *providerTMockStep) Failed() bool {
	return m.failed
}

func (m *providerTMockStep) Error(args ...interface{}) {
	m.error = true
	m.failed = true
}

func (m *providerTMockStep) Log(args ...interface{}) {
	m.log = true
}

func (m *providerTMockStep) Logf(format string, args ...interface{}) {
	m.logf = true
}

func (m *providerTMockStep) Name() string {
	return m.name
}

func (m *providerTMockStep) GetRealT() provider.TestingT {
	return m.testingT
}

func (m *providerTMockStep) SetRealT(realT provider.TestingT) {
	m.testingT = realT
}

type providerMockStep struct {
	status allure.Status
	msg    string
	trace  string

	executionContext provider.ExecutionContext
}

func (m *providerMockStep) StopResult(status allure.Status) {
	m.status = status
}

func (m *providerMockStep) UpdateResultStatus(msg, trace string) {
	m.msg = msg
	m.trace = trace
}

func (m *providerMockStep) ExecutionContext() provider.ExecutionContext {
	return m.executionContext
}

func (m *providerMockStep) Helper() {
}

type executionCtxMock struct {
	name string

	steps       []*allure.Step
	attachments []*allure.Attachment
}

func newExecutionCtxMock(name string) *executionCtxMock {
	return &executionCtxMock{
		name:        name,
		steps:       []*allure.Step{},
		attachments: []*allure.Attachment{},
	}
}
func (m *executionCtxMock) AddStep(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *executionCtxMock) AddAttachments(attachments ...*allure.Attachment) {
	m.attachments = append(m.attachments, attachments...)
}

func (m *executionCtxMock) GetName() string {
	return m.name
}

func TestNewStepCtx(t *testing.T) {
	params := allure.NewParameters("p1", "v1", "p2", "v2")
	ctx := NewStepCtx(
		newStepProviderMock(),
		&providerMockStep{},
		"stepName",
		params...,
	)
	require.NotNil(t, ctx)
	require.NotNil(t, ctx.CurrentStep())
	require.Equal(t, "stepName", ctx.CurrentStep().Name)
	require.NotNil(t, ctx.CurrentStep().Parameters)
	require.Equal(t, params, ctx.CurrentStep().Parameters)
	require.NotNil(t, ctx.Assert())
	require.NotNil(t, ctx.Require())
}

func TestStepCtx_NewChildCtx(t *testing.T) {
	params := allure.NewParameters("p1", "v1", "p2", "v2")

	mockT := newStepProviderMock()
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, p: &providerMockStep{}, currentStep: step}

	childCtx := ctx.NewChildCtx("new step", params...)

	require.NotNil(t, childCtx)
	require.NotNil(t, childCtx.CurrentStep())
	require.Equal(t, "new step", childCtx.CurrentStep().Name)
	require.NotNil(t, childCtx.CurrentStep().Parameters)
	require.Equal(t, params, childCtx.CurrentStep().Parameters)

	require.NotNil(t, childCtx.Assert())
	require.NotNil(t, childCtx.Require())
}

func TestStepCtx_Broken_noParent(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.Broken()
	require.True(t, mockT.Failed())
	require.Equal(t, allure.Broken, step.Status)
}

func TestStepCtx_Broken_withParent(t *testing.T) {
	mockT := new(providerTMockStep)
	parentStep := allure.NewSimpleStep("parentStep")
	parentCtx := &stepCtx{t: mockT, currentStep: parentStep}
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step, parentStep: parentCtx}
	ctx.Broken()
	require.True(t, mockT.Failed())
	require.Equal(t, allure.Broken, step.Status)
	require.Equal(t, allure.Broken, parentStep.Status)
}

func TestStepCtx_FailNow(t *testing.T) {
	mockT := newStepProviderMock()
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.FailNow()
	require.True(t, mockT.failNow)
}

func TestStepCtx_Fail_noParent(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.Fail()
	require.False(t, mockT.Failed())
	require.Equal(t, allure.Failed, step.Status)
}

func TestStepCtx_Fail_withParent(t *testing.T) {
	mockT := new(providerTMockStep)
	parentStep := allure.NewSimpleStep("parentStep")
	parentCtx := &stepCtx{t: mockT, currentStep: parentStep}
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step, parentStep: parentCtx}
	ctx.Fail()
	require.False(t, mockT.Failed())
	require.Equal(t, allure.Failed, step.Status)
	require.Equal(t, allure.Failed, parentStep.Status)
}

func TestStepCtx_CurrentStep(t *testing.T) {
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{currentStep: step}
	require.Equal(t, step, ctx.CurrentStep())
}

func TestStepCtx_Step(t *testing.T) {
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{currentStep: step}
	newStep := allure.NewSimpleStep("childStep")
	ctx.Step(newStep)
	require.Len(t, step.Steps, 1)
	require.Equal(t, newStep, step.Steps[0])
}

func TestStepCtx_Errorf_withParent(t *testing.T) {
	mockT := new(providerTMockStep)
	mockT.SetRealT(t)
	parentStep := allure.NewSimpleStep("parentStep")
	parentCtx := &stepCtx{t: mockT, currentStep: parentStep}
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step, parentStep: parentCtx}
	ctx.Errorf("test")
	require.True(t, mockT.Failed())
	require.Equal(t, allure.Failed, step.Status)
	require.Equal(t, allure.Failed, parentStep.Status)
}

func TestStepCtx_Errorf_noParent(t *testing.T) {
	mockT := new(providerTMockStep)
	mockT.SetRealT(t)
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.Errorf("test")
	require.True(t, mockT.Failed())
	require.Equal(t, allure.Failed, step.Status)
}

func TestStepCtx_Error_withParent(t *testing.T) {
	mockT := new(providerTMockStep)
	mockT.SetRealT(t)

	parentStep := allure.NewSimpleStep("parentStep", allure.NewParameters("paramParent1", "v1", "paramParent2", "v2")...)
	parentCtx := &stepCtx{t: mockT, currentStep: parentStep}
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step, parentStep: parentCtx}
	ctx.Error("test")
	require.True(t, mockT.Failed())
	require.Equal(t, allure.Failed, step.Status)
	require.Equal(t, allure.Failed, parentStep.Status)
}

func TestStepCtx_Error_noParent(t *testing.T) {
	mockT := new(providerTMockStep)
	mockT.SetRealT(t)

	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.Error("test")
	require.True(t, mockT.Failed())
	require.Equal(t, allure.Failed, step.Status)
}

func TestStepCtx_Assert(t *testing.T) {
	providerMock := newStepProviderMock()
	test := helper.NewAssertsHelper(providerMock)
	ctx := stepCtx{asserts: test}
	require.Equal(t, test, ctx.Assert())
}

func TestStepCtx_Require(t *testing.T) {
	providerMock := newStepProviderMock()
	test := helper.NewRequireHelper(providerMock)
	ctx := stepCtx{require: test}
	require.Equal(t, test, ctx.Require())
}

func TestStepCtx_WG(t *testing.T) {
	test := sync.WaitGroup{}
	ctx := stepCtx{wg: test}
	require.Equal(t, &test, ctx.WG())
}

func TestStepCtx_WithParameters(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	params := allure.NewParameters("p1", "v1", "p2", "v2")

	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.WithParameters(params...)

	require.NotNil(t, step.Parameters)
	require.NotEmpty(t, step.Parameters)
	require.Equal(t, params, step.Parameters)
}

func TestStepCtx_WithNewParameters(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.WithNewParameters("p1", "v1", "p2", "v2")

	require.NotNil(t, step.Parameters)
	require.NotEmpty(t, step.Parameters)
	require.Len(t, step.Parameters, 2)
	require.Equal(t, "p1", step.Parameters[0].Name)
	require.Equal(t, "v1", step.Parameters[0].GetValue())
	require.Equal(t, "p2", step.Parameters[1].Name)
	require.Equal(t, "v2", step.Parameters[1].GetValue())
}

func TestStepCtx_WithAttachments(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, currentStep: step}
	attachments := []*allure.Attachment{
		allure.NewAttachment("attach1", allure.Text, []byte("attach text 1")),
		allure.NewAttachment("attach2", allure.Text, []byte("attach text 2")),
	}

	ctx.WithAttachments(attachments...)
	require.NotNil(t, step.Attachments)
	require.NotEmpty(t, step.Attachments)
	require.Len(t, step.Attachments, 2)
	require.Equal(t, attachments, step.Attachments)
}

func TestStepCtx_WithNewAttachment(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.WithNewAttachment("attach1", allure.Text, []byte("attach text 1"))
	require.NotNil(t, step.Attachments)
	require.NotEmpty(t, step.Attachments)
	require.Len(t, step.Attachments, 1)
	require.Equal(t, "attach1", step.Attachments[0].Name)
	require.Equal(t, allure.Text, step.Attachments[0].Type)
	require.Equal(t, []byte("attach text 1"), step.Attachments[0].GetContent())
}

func TestStepCtx_NewStep(t *testing.T) {
	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.NewStep("New Step", allure.NewParameter("p1", "v1"))
	require.NotNil(t, ctx.currentStep.Steps)
	require.NotEmpty(t, ctx.currentStep.Steps)
	require.Len(t, ctx.currentStep.Steps, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Name, "New Step")
	require.NotNil(t, ctx.currentStep.Steps[0].Parameters)
	require.NotEmpty(t, ctx.currentStep.Steps[0].Parameters)
	require.Len(t, ctx.currentStep.Steps[0].Parameters, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].Name, "p1")
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].GetValue(), "v1")
}

func TestStepCtx_WithNewStep(t *testing.T) {
	flag := false
	stepF := func(ctx provider.StepCtx) {
		flag = true
	}

	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.WithNewStep("new step", stepF, allure.NewParameter("p1", "v1"))
	require.True(t, flag)
	require.NotNil(t, ctx.currentStep.Steps)
	require.NotEmpty(t, ctx.currentStep.Steps)
	require.Len(t, ctx.currentStep.Steps, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Name, "new step")
	require.NotNil(t, ctx.currentStep.Steps[0].Parameters)
	require.NotEmpty(t, ctx.currentStep.Steps[0].Parameters)
	require.Len(t, ctx.currentStep.Steps[0].Parameters, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].Name, "p1")
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].GetValue(), "v1")
}

func TestStepCtx_WithNewAsyncStep(t *testing.T) {
	wg := sync.WaitGroup{}
	flag := false
	wg.Add(1)
	stepF := func(ctx provider.StepCtx) {
		flag = true
		defer wg.Done()
	}

	mockT := new(providerTMockStep)
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, currentStep: step}
	ctx.WithNewAsyncStep("new step", stepF, allure.NewParameter("p1", "v1"))
	wg.Wait()
	require.True(t, flag)
	require.NotNil(t, ctx.currentStep.Steps)
	require.NotEmpty(t, ctx.currentStep.Steps)
	require.Len(t, ctx.currentStep.Steps, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Name, "new step")
	require.NotNil(t, ctx.currentStep.Steps[0].Parameters)
	require.NotEmpty(t, ctx.currentStep.Steps[0].Parameters)
	require.Len(t, ctx.currentStep.Steps[0].Parameters, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].Name, "p1")
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].GetValue(), "v1")
}

func TestStepCtx_WithNewStep_panic(t *testing.T) {
	flag := false
	stepF := func(ctx provider.StepCtx) {
		flag = true
		panic("whoops")
	}

	mockT := newStepProviderMock()
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, p: &providerMockStep{executionContext: newExecutionCtxMock(constants.TestContextName)}, currentStep: step}

	ctx.WithNewStep("new step", stepF, allure.NewParameter("p1", "v1"))
	require.True(t, flag)
	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)

	require.NotNil(t, ctx.currentStep.Steps)
	require.NotEmpty(t, ctx.currentStep.Steps)
	require.Len(t, ctx.currentStep.Steps, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Name, "new step")
	require.NotNil(t, ctx.currentStep.Steps[0].Parameters)
	require.NotEmpty(t, ctx.currentStep.Steps[0].Parameters)
	require.Len(t, ctx.currentStep.Steps[0].Parameters, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].Name, "p1")
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].GetValue(), "v1")
}

func TestStepCtx_WithNewAsyncStep_panic(t *testing.T) {
	flag := false
	stepF := func(ctx provider.StepCtx) {
		flag = true
		panic("whoops")
	}

	mockT := newStepProviderMock()
	step := allure.NewSimpleStep("testStep")

	ctx := stepCtx{t: mockT, p: &providerMockStep{executionContext: newExecutionCtxMock(constants.TestContextName)}, currentStep: step}
	ctx.WithNewAsyncStep("new step", stepF, allure.NewParameter("p1", "v1"))

	// wg doesn't help cause panic
	time.Sleep(100 * time.Millisecond)

	require.True(t, flag)
	require.True(t, mockT.errorF)
	require.True(t, mockT.failNow)

	require.NotNil(t, ctx.currentStep.Steps)
	require.NotEmpty(t, ctx.currentStep.Steps)
	require.Len(t, ctx.currentStep.Steps, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Name, "new step")
	require.NotNil(t, ctx.currentStep.Steps[0].Parameters)
	require.NotEmpty(t, ctx.currentStep.Steps[0].Parameters)
	require.Len(t, ctx.currentStep.Steps[0].Parameters, 1)
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].Name, "p1")
	require.Equal(t, ctx.currentStep.Steps[0].Parameters[0].GetValue(), "v1")
}

func TestStepCtx_Name(t *testing.T) {
	mockT := newStepProviderMock()
	mockT.name = "test"
	var actualName string

	stepF := func(ctx provider.StepCtx) {
		t.Logf(ctx.Name())
		actualName = ctx.Name()
	}
	step := allure.NewSimpleStep("testStep")
	ctx := stepCtx{t: mockT, p: &providerMockStep{executionContext: newExecutionCtxMock(constants.TestContextName)}, currentStep: step}
	ctx.WithNewStep("new step", stepF)
	require.Equal(t, mockT.name, actualName)
}
