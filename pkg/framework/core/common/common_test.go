package common

import (
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/manager"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

type providerMockCommon struct {
	provider.AllureForwardFull

	testMetaMock  *testMetaMockCommon
	suiteMetaMock *suiteMetaMockCommon
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
	return nil
}
func (m *providerMockCommon) TestContext()                                         {}
func (m *providerMockCommon) BeforeEachContext()                                   {}
func (m *providerMockCommon) AfterEachContext()                                    {}
func (m *providerMockCommon) BeforeAllContext()                                    {}
func (m *providerMockCommon) AfterAllContext()                                     {}
func (m *providerMockCommon) NewTest(testName, packageName string, tags ...string) {}
func (m *providerMockCommon) FinishTest()                                          {}

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
	steps      []*allure.Step
	errorf     string
	errorfFlag bool
	failNow    bool
	parallel   bool
	run        bool
	skipped    bool
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
	return m.run
}

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
	cfg := manager.NewProviderConfig().
		WithRunner("WithRunner").
		WithPackageName("WithPackageName").
		WithSuiteName("WithSuiteName").
		WithFullName("WithFullName")

	comm := Common{TestingT: mock, Provider: manager.NewProvider(cfg)}
	comm.Error("test")
	require.True(t, mock.errorfFlag)
}

func TestCommon_Errorf(t *testing.T) {
	mock := newCommonTMock()
	cfg := manager.NewProviderConfig().
		WithRunner("WithRunner").
		WithPackageName("WithPackageName").
		WithSuiteName("WithSuiteName").
		WithFullName("WithFullName")

	comm := Common{TestingT: mock, Provider: manager.NewProvider(cfg)}
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
	mockT := newCommonTMock()
	comm := Common{TestingT: mockT}
	result := comm.Run("myTest", func(t provider.T) {}, "tag1", "tag2")

	require.True(t, result)
	require.True(t, mockT.run)
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
	c := NewT(mockT, "packageName", "suiteName")
	require.NotNil(t, c)
	require.NotNil(t, c.wg)

	require.NotNil(t, c.require)
	require.NotNil(t, c.assert)
	require.False(t, c.xSkip)

	require.NotNil(t, c.Provider)
	require.NotNil(t, c.Provider.GetTestMeta())
	require.NotNil(t, c.Provider.GetSuiteMeta())
	require.NotNil(t, c.Provider.GetSuiteMeta().GetContainer())
	require.Equal(t, "suiteName", c.Provider.GetSuiteMeta().GetSuiteName())
	require.Equal(t, "packageName", c.Provider.GetSuiteMeta().GetPackageName())
}

func TestNewTestT(t *testing.T) {
	mockT := new(testing.T)
	c := NewT(mockT, "packageName", "suiteName")
	newMockT := new(testing.T)
	cNew := NewTestT(newMockT, c.Provider, c, "packageName2", "testName")
	require.NotNil(t, cNew)
	require.NotNil(t, cNew.wg)

	require.NotNil(t, cNew.require)
	require.NotNil(t, cNew.assert)
	require.False(t, cNew.xSkip)

	require.NotNil(t, cNew.Provider)
	require.NotNil(t, cNew.Provider.GetTestMeta())
	require.NotNil(t, cNew.Provider.GetTestMeta().GetResult())
	require.NotNil(t, cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Suite))
	require.NotEmpty(t, cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Suite))
	require.Len(t, cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Suite), 1)
	require.Equal(t, "suiteName", cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Suite)[0].Value)

	require.Equal(t, "testName", cNew.Provider.GetTestMeta().GetResult().Name)
	require.NotNil(t, cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Package))
	require.NotEmpty(t, cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Package))
	require.Len(t, cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Package), 1)
	require.Equal(t, "packageName2", cNew.Provider.GetTestMeta().GetResult().GetLabel(allure.Package)[0].Value)

	require.NotNil(t, cNew.Provider.GetSuiteMeta())
	require.NotNil(t, cNew.Provider.GetSuiteMeta().GetContainer())
	require.Equal(t, "suiteName", cNew.Provider.GetSuiteMeta().GetSuiteName())
	require.Equal(t, "packageName2", cNew.Provider.GetSuiteMeta().GetPackageName())
}
