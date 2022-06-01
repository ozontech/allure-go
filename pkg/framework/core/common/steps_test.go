package common

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type executionContextstepsCommMock struct {
	name        string
	steps       []*allure.Step
	attachments []*allure.Attachment
}

func newExecContextstepsCommMock(name string) *executionContextstepsCommMock {
	return &executionContextstepsCommMock{
		name:        name,
		steps:       []*allure.Step{},
		attachments: []*allure.Attachment{},
	}
}

func (m *executionContextstepsCommMock) AddStep(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *executionContextstepsCommMock) AddAttachments(attachments ...*allure.Attachment) {
	m.attachments = append(m.attachments, attachments...)
}

func (m *executionContextstepsCommMock) GetName() string {
	return m.name
}

type providerMockstepsCommon struct {
	provider.AllureForwardFull
	steps         []*allure.Step
	testMetaMock  provider.TestMeta
	suiteMetaMock *suiteMetaMockstepsCommon
	executionMock *executionContextstepsCommMock
}

func (m *providerMockstepsCommon) SetTestMeta(meta provider.TestMeta) {
	m.testMetaMock = meta
}

func (m *providerMockstepsCommon) GetResult() *allure.Result {
	return m.testMetaMock.GetResult()
}

func (m *providerMockstepsCommon) UpdateResultStatus(msg string, trace string) {}

func (m *providerMockstepsCommon) StopResult(status allure.Status) {}

func (m *providerMockstepsCommon) GetTestMeta() provider.TestMeta {
	return m.testMetaMock
}

func (m *providerMockstepsCommon) GetSuiteMeta() provider.SuiteMeta {
	return m.suiteMetaMock
}

func (m *providerMockstepsCommon) ExecutionContext() provider.ExecutionContext {
	return m.executionMock
}

func (m *providerMockstepsCommon) Step(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *providerMockstepsCommon) TestContext()                                         {}
func (m *providerMockstepsCommon) BeforeEachContext()                                   {}
func (m *providerMockstepsCommon) AfterEachContext()                                    {}
func (m *providerMockstepsCommon) BeforeAllContext()                                    {}
func (m *providerMockstepsCommon) AfterAllContext()                                     {}
func (m *providerMockstepsCommon) NewTest(testName, packageName string, tags ...string) {}
func (m *providerMockstepsCommon) FinishTest()                                          {}

type suiteMetaMockstepsCommon struct {
	namePrefix string
	name       string
	container  *allure.Container
	hook       func(t provider.T)
}

func (m *suiteMetaMockstepsCommon) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockstepsCommon) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockstepsCommon) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockstepsCommon) GetSuiteFullName() string {
	return fmt.Sprintf("%s/%s", m.namePrefix, m.name)
}

func (m *suiteMetaMockstepsCommon) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockstepsCommon) SetBeforeAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockstepsCommon) SetAfterAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockstepsCommon) GetBeforeAll() func(provider.T) {
	return m.hook
}

func (m *suiteMetaMockstepsCommon) GetAfterAll() func(provider.T) {
	return m.hook
}

type testMetaMockstepsCommon struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockstepsCommon) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockstepsCommon) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockstepsCommon) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockstepsCommon) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockstepsCommon) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockstepsCommon) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockstepsCommon) GetAfterEach() func(t provider.T) {
	return m.ae
}

type stepsStepsCommTMock struct {
	testing.TB
	t          *testing.T
	steps      []*allure.Step
	errorf     string
	errorfFlag bool
	failNow    bool
	parallel   bool
	run        bool
	skipped    bool
}

func newStepsCommonTMock() *stepsStepsCommTMock {
	return &stepsStepsCommTMock{steps: []*allure.Step{}}
}

func (m *stepsStepsCommTMock) Skip(args ...interface{}) {
	m.skipped = true
}

func (m *stepsStepsCommTMock) Step(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *stepsStepsCommTMock) Errorf(format string, args ...interface{}) {
	m.errorfFlag = true
}

func (m *stepsStepsCommTMock) Error(args ...interface{}) {
	m.errorfFlag = true
}

func (m *stepsStepsCommTMock) FailNow() {
	m.failNow = true
}

func (m *stepsStepsCommTMock) Parallel() {
	m.parallel = true
}

func (m *stepsStepsCommTMock) Run(testName string, testBody func(t *testing.T)) bool {
	m.run = true
	testBody(m.t)
	return m.run
}

func TestCommon_WithNewStep(t *testing.T) {
	mockT := newStepsCommonTMock()
	mockT.t = new(testing.T)
	p := &providerMockstepsCommon{
		testMetaMock:  &testMetaMockstepsCommon{},
		suiteMetaMock: &suiteMetaMockstepsCommon{},
		executionMock: newExecContextstepsCommMock(constants.TestContextName),
	}
	comm := Common{TestingT: mockT, Provider: p}
	params := allure.NewParameters("p1", "v1", "p2", "v2")
	comm.WithNewStep("step", func(ctx provider.StepCtx) {}, params...)
	require.NotEmpty(t, p.steps)
	require.Len(t, p.steps, 1)
	require.Equal(t, "step", p.steps[0].Name)
	require.Equal(t, params, p.steps[0].Parameters)
}

func TestCommon_WithNewStep_panic(t *testing.T) {
	mockT := newStepsCommonTMock()
	mockT.t = new(testing.T)
	p := &providerMockstepsCommon{
		testMetaMock:  &testMetaMockstepsCommon{result: &allure.Result{}},
		suiteMetaMock: &suiteMetaMockstepsCommon{},
		executionMock: newExecContextstepsCommMock(constants.TestContextName),
	}
	comm := Common{TestingT: mockT, Provider: p}
	params := allure.NewParameters("p1", "v1", "p2", "v2")
	comm.WithNewStep("step", func(ctx provider.StepCtx) { panic("whoops") }, params...)
	require.NotEmpty(t, p.steps)
	require.Len(t, p.steps, 1)
	require.Equal(t, "step", p.steps[0].Name)
	require.Equal(t, params, p.steps[0].Parameters)
}

func TestCommon_WithNewAsyncStep(t *testing.T) {
	mockT := newStepsCommonTMock()
	mockT.t = new(testing.T)
	p := &providerMockstepsCommon{
		testMetaMock:  &testMetaMockstepsCommon{},
		suiteMetaMock: &suiteMetaMockstepsCommon{},
		executionMock: newExecContextstepsCommMock(constants.TestContextName),
	}
	comm := Common{TestingT: mockT, Provider: p}
	params := allure.NewParameters("p1", "v1", "p2", "v2")
	comm.WithNewAsyncStep("step", func(ctx provider.StepCtx) {}, params...)
	time.Sleep(100 * time.Millisecond)
	require.NotEmpty(t, p.steps)
	require.Len(t, p.steps, 1)
	require.Equal(t, "step", p.steps[0].Name)
	require.Equal(t, params, p.steps[0].Parameters)
}

func TestCommon_WithNewAsyncStep_panic(t *testing.T) {
	mockT := newStepsCommonTMock()
	mockT.t = new(testing.T)
	p := &providerMockstepsCommon{
		testMetaMock:  &testMetaMockstepsCommon{result: &allure.Result{}},
		suiteMetaMock: &suiteMetaMockstepsCommon{},
		executionMock: newExecContextstepsCommMock(constants.TestContextName),
	}
	comm := Common{TestingT: mockT, Provider: p}
	params := allure.NewParameters("p1", "v1", "p2", "v2")
	comm.WithNewAsyncStep("step", func(ctx provider.StepCtx) { panic("whoops") }, params...)
	time.Sleep(100 * time.Millisecond)
	require.NotEmpty(t, p.steps)
	require.Len(t, p.steps, 1)
	require.Equal(t, "step", p.steps[0].Name)
	require.Equal(t, params, p.steps[0].Parameters)
}
