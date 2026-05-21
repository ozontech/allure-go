package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type executionContextCommMock struct {
	name        string
	steps       []*allure.Step
	attachments []*allure.Attachment
}

func newExecContextCommMock(name string) *executionContextCommMock {
	return &executionContextCommMock{
		name:        name,
		steps:       []*allure.Step{},
		attachments: []*allure.Attachment{},
	}
}

func (m *executionContextCommMock) AddStep(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *executionContextCommMock) AddAttachments(attachments ...*allure.Attachment) {
	m.attachments = append(m.attachments, attachments...)
}

func (m *executionContextCommMock) GetName() string {
	return m.name
}

func (m *executionContextCommMock) GetTestResult() *allure.Result {
	return nil
}

type providerMockCommon struct {
	provider.AllureForwardFull

	testMetaMock  provider.TestMeta
	suiteMetaMock *suiteMetaMockCommon
	executionMock provider.ExecutionContext // Changed to interface
}

func newProviderMockCommon(name, fullName string) *providerMockCommon {
	return &providerMockCommon{
		testMetaMock:  &testMetaMockCommon{result: allure.NewResult(name, fullName)},
		suiteMetaMock: nil,
		executionMock: nil,
	}
}

func (m *providerMockCommon) GetResult() *allure.Result {
	return m.testMetaMock.GetResult()
}

func (m *providerMockCommon) UpdateResultStatus(msg string, trace string) {}

func (m *providerMockCommon) StopResult(status allure.Status) {}

func (m *providerMockCommon) GetTestMeta() provider.TestMeta {
	return m.testMetaMock
}

func (m *providerMockCommon) GetSuiteMeta() provider.SuiteMeta {
	return m.suiteMetaMock
}

func (m *providerMockCommon) ExecutionContext() provider.ExecutionContext {
	return m.executionMock
}

func (m *providerMockCommon) SetTestMeta(meta provider.TestMeta) {
	m.testMetaMock = meta
}

func (m *providerMockCommon) TestContext()                                         {}
func (m *providerMockCommon) BeforeEachContext()                                   {}
func (m *providerMockCommon) AfterEachContext()                                    {}
func (m *providerMockCommon) BeforeAllContext()                                    {}
func (m *providerMockCommon) AfterAllContext()                                     {}
func (m *providerMockCommon) NewTest(testName, packageName string, tags ...string) {}
func (m *providerMockCommon) FinishTest() error {
	return nil
}

type suiteMetaMockCommon struct {
	namePrefix string
	name       string
	container  *allure.Container
	hook       func(t provider.T)
}

func (m *suiteMetaMockCommon) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockCommon) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockCommon) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockCommon) GetParentSuite() string {
	return ""
}

func (m *suiteMetaMockCommon) GetSuiteFullName() string {
	return fmt.Sprintf("%s/%s", m.namePrefix, m.name)
}

func (m *suiteMetaMockCommon) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockCommon) SetBeforeAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockCommon) SetAfterAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockCommon) GetBeforeAll() func(provider.T) {
	return m.hook
}

func (m *suiteMetaMockCommon) GetAfterAll() func(provider.T) {
	return m.hook
}

type testMetaMockCommon struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockCommon) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockCommon) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockCommon) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockCommon) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockCommon) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockCommon) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockCommon) GetAfterEach() func(t provider.T) {
	return m.ae
}

type commonTMock struct {
	testing.TB
	t          *testing.T
	steps      []*allure.Step
	errorf     string
	errorfFlag bool
	failNow    bool
	parallel   bool
	run        bool
	skipped    bool
	TestingT   *testing.T
}

func newCommonTMock() *commonTMock {
	return &commonTMock{steps: []*allure.Step{}}
}

func (m *commonTMock) Skip(args ...interface{}) {
	m.skipped = true
}

func (m *commonTMock) Step(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *commonTMock) Errorf(format string, args ...interface{}) {
	m.errorfFlag = true
}

func (m *commonTMock) Error(args ...interface{}) {
	m.errorfFlag = true
}

func (m *commonTMock) FailNow() {
	m.failNow = true
}

func (m *commonTMock) Parallel() {
	m.parallel = true
}

func (m *commonTMock) Run(testName string, testBody func(t *testing.T)) bool {
	m.run = true
	testBody(m.t)
	return m.run
}

func (m *commonTMock) Helper() {}

func TestCommon_Assert(t *testing.T) {
	asserts := helper.NewAssertsHelper(newCommonTMock())
	comm := Common{assert: asserts}
	require.NotNil(t, comm.Assert())
	require.Equal(t, asserts, comm.assert)
}

func TestCommon_Require(t *testing.T) {
	asserts := helper.NewRequireHelper(newCommonTMock())
	comm := Common{assert: asserts}
	require.NotNil(t, comm.Assert())
	require.Equal(t, asserts, comm.assert)
}

func TestCommon_Error(t *testing.T) {
	mock := newCommonTMock()
	comm := Common{TestingT: mock, Provider: newProviderMockCommon("name", "fullName")}
	comm.Error("test")
	require.True(t, mock.errorfFlag)
}

func TestCommon_Errorf(t *testing.T) {
	mock := newCommonTMock()
	comm := Common{TestingT: mock, Provider: newProviderMockCommon("name", "fullName")}
	comm.Errorf("test")
	require.True(t, mock.errorfFlag)
}

func TestCommon_WG(t *testing.T) {
	comm := Common{wg: sync.WaitGroup{}}
	require.NotNil(t, comm.WG())
}

func TestCommon_GetProvider(t *testing.T) {
	cfg := manager.NewProviderConfig().
		WithRunner("WithRunner").
		WithPackageName("WithPackageName").
		WithSuiteName("WithSuiteName").
		WithFullName("WithFullName")
	comm := Common{Provider: manager.NewProvider(cfg)}
	require.NotNil(t, comm.GetProvider())
}

func TestCommon_RealT(t *testing.T) {
	mockT := newCommonTMock()
	comm := Common{TestingT: mockT}
	require.NotNil(t, comm.RealT())
	require.Equal(t, mockT, comm.RealT())
}

func TestCommon_Parallel(t *testing.T) {
	mockT := newCommonTMock()
	comm := Common{TestingT: mockT}
	comm.Parallel()

	require.True(t, mockT.parallel)
}

func TestCommon_Run(t *testing.T) {
	t.Skip("This test must be reworked")
	mockT := newCommonTMock()
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	mockT.t = new(testing.T)

	comm := Common{TestingT: mockT, Provider: &providerMockCommon{
		testMetaMock:  &testMetaMockCommon{result: &allure.Result{}, container: allure.NewContainer()},
		suiteMetaMock: &suiteMetaMockCommon{container: allure.NewContainer()},
		executionMock: newExecContextCommMock(constants.TestContextName),
	}}
	result := comm.Run("myTest", func(t provider.T) {}, "tag1", "tag2")

	require.NotNil(t, result)
	require.True(t, mockT.run)

	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 1)

	var resultFile *os.File
	defer resultFile.Close()

	f := files[0]
	emptyResult := &allure.Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := ioutil.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := sonic.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)
}

func TestCommon_Run_panicHandle(t *testing.T) {
	t.Skip("This test must be fixed")
	mockT := newCommonTMock()
	allureDir := "./allure-results"

	defer os.RemoveAll(allureDir)

	mockT.t = new(testing.T)
	comm := Common{TestingT: mockT, Provider: &providerMockCommon{
		testMetaMock:  &testMetaMockCommon{result: &allure.Result{}, container: allure.NewContainer()},
		suiteMetaMock: &suiteMetaMockCommon{container: allure.NewContainer()},
		executionMock: newExecContextCommMock(constants.TestContextName),
	}}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// routine for runtime.Goexit() simulation
	go require.NotPanics(t, func() {
		defer wg.Done()
		comm.Run("myTest", func(t provider.T) { panic("whoops") }, "tag1", "tag2")
	})
	wg.Wait()

	require.True(t, mockT.t.Failed())
	require.True(t, mockT.run)

	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 1)

	var resultFile *os.File
	defer resultFile.Close()

	f := files[0]
	emptyResult := &allure.Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := ioutil.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := sonic.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)
}

func TestCommon_XSkip(t *testing.T) {
	comm := Common{}
	comm.XSkip()
	require.True(t, comm.xSkip)
}

func TestCommon_Skip(t *testing.T) {
	mockT := newCommonTMock()
	comm := Common{TestingT: mockT, Provider: &providerMockCommon{testMetaMock: &testMetaMockCommon{result: &allure.Result{}}}}

	comm.Skip("msg")
	require.Equal(t, "msg\n", comm.Provider.GetResult().StatusDetails.Message)
	require.Equal(t, "msg\n", comm.Provider.GetResult().StatusDetails.Trace)

	require.True(t, mockT.skipped)
}

func TestCommon_SkipOnPrint(t *testing.T) {
	mockT := newCommonTMock()
	comm := Common{TestingT: mockT, Provider: &providerMockCommon{testMetaMock: &testMetaMockCommon{result: &allure.Result{ToPrint: true}}}}

	comm.SkipOnPrint()
	require.False(t, comm.GetResult().ToPrint)
}

func TestCommon_GetResult(t *testing.T) {
	res := &allure.Result{ToPrint: true}
	mockT := newCommonTMock()
	comm := Common{TestingT: mockT, Provider: &providerMockCommon{testMetaMock: &testMetaMockCommon{result: res}}}
	require.Equal(t, res, comm.GetResult())
}

func TestNewT(t *testing.T) {
	mockT := new(testing.T)
	c := NewT(mockT)

	require.NotNil(t, c)
	require.NotNil(t, c.wg)

	require.NotNil(t, c.require)
	require.NotNil(t, c.assert)
	require.False(t, c.xSkip)

	require.Nil(t, c.Provider)
}

func TestCopyLabels(t *testing.T) {
	input := &allure.Result{}
	epic := allure.EpicLabel("EpicTest")
	parentSuite := allure.ParentSuiteLabel("ParentSuiteTest")
	lead := allure.LeadLabel("LeadTest")
	owner := allure.OwnerLabel("OwnerTest")
	input.WithLabels(epic, parentSuite, lead, owner)
	target := &allure.Result{}
	target = copyLabels(input, target)
	require.NotNil(t, target.Labels)
	require.Len(t, target.Labels, 4)

	require.NotEmpty(t, target.GetLabels(allure.Epic))
	require.Len(t, target.GetLabels(allure.Epic), 1)
	require.Equal(t, epic, target.GetLabels(allure.Epic)[0])

	require.NotEmpty(t, target.GetLabels(allure.ParentSuite))
	require.Len(t, target.GetLabels(allure.ParentSuite), 1)
	require.Equal(t, parentSuite, target.GetLabels(allure.ParentSuite)[0])

	require.NotEmpty(t, target.GetLabels(allure.Lead))
	require.Len(t, target.GetLabels(allure.Lead), 1)
	require.Equal(t, lead, target.GetLabels(allure.Lead)[0])

	require.NotEmpty(t, target.GetLabels(allure.Owner))
	require.Len(t, target.GetLabels(allure.Owner), 1)
	require.Equal(t, owner, target.GetLabels(allure.Owner)[0])
}

// Tests for GetCurrentTestResult functionality
func TestCommon_GetCurrentTestResult_AfterEachContext(t *testing.T) {
	testResult := allure.NewResult("TestName", "TestFullName")
	testResult.Status = allure.Failed
	testResult.SetStatusMessage("Test failed")

	// Mock execution context that returns test result (AfterEach context)
	execContext := &executionContextWithResult{
		name:       constants.AfterEachContextName,
		testResult: testResult,
	}

	provider := &providerMockCommon{
		testMetaMock:  &testMetaMockCommon{result: testResult},
		executionMock: execContext,
	}

	comm := Common{Provider: provider}

	result, _ := comm.GetCurrentTestResult()
	require.NotNil(t, result)
	require.Equal(t, allure.Failed, result.Status)
	require.Equal(t, "Test failed", result.GetStatusMessage())
}

func TestCommon_GetCurrentTestResult_NilInTestContext(t *testing.T) {
	testResult := allure.NewResult("TestName", "TestFullName")

	// Mock test context - should return nil
	execContext := &executionContextWithResult{
		name:       constants.TestContextName,
		testResult: testResult,
	}

	provider := &providerMockCommon{
		testMetaMock:  &testMetaMockCommon{result: testResult},
		executionMock: execContext,
	}

	comm := Common{Provider: provider}

	result, _ := comm.GetCurrentTestResult()
	require.Nil(t, result, "GetCurrentTestResult should return nil in test context")
}

func TestCommon_GetCurrentTestResult_NilInBeforeEachContext(t *testing.T) {
	testResult := allure.NewResult("TestName", "TestFullName")

	// Mock BeforeEach context - should return nil
	execContext := &executionContextWithResult{
		name:       constants.BeforeEachContextName,
		testResult: testResult,
	}

	provider := &providerMockCommon{
		testMetaMock:  &testMetaMockCommon{result: testResult},
		executionMock: execContext,
	}

	comm := Common{Provider: provider}

	result, _ := comm.GetCurrentTestResult()
	require.Nil(t, result, "GetCurrentTestResult should return nil in BeforeEach context")
}

func TestCommon_GetCurrentTestResult_NilProviderOrContext(t *testing.T) {
	// Test with nil Provider
	comm := Common{Provider: nil}
	res, _ := comm.GetCurrentTestResult()
	require.Nil(t, res)

	// Test with nil ExecutionContext
	provider := &providerMockCommon{
		executionMock: nil,
	}
	comm = Common{Provider: provider}
	res, _ = comm.GetCurrentTestResult()
	require.Nil(t, res)
}

// Mock execution context for testing GetCurrentTestResult
type executionContextWithResult struct {
	name        string
	testResult  *allure.Result
	steps       []*allure.Step
	attachments []*allure.Attachment
}

func (m *executionContextWithResult) AddStep(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *executionContextWithResult) AddAttachments(attachments ...*allure.Attachment) {
	m.attachments = append(m.attachments, attachments...)
}

func (m *executionContextWithResult) GetName() string {
	return m.name
}

func (m *executionContextWithResult) GetTestResult() *allure.Result {
	return m.testResult
}
