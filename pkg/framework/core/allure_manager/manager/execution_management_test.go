package manager

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
)

type testMetaMockExecM struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockExecM) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockExecM) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockExecM) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockExecM) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockExecM) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockExecM) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockExecM) GetAfterEach() func(t provider.T) {
	return m.ae
}

type suiteMetaMockExecM struct {
	name      string
	container *allure.Container
	hook      func(t provider.T)
}

func (m *suiteMetaMockExecM) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockExecM) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockExecM) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockExecM) GetParentSuite() string {
	return ""
}

func (m *suiteMetaMockExecM) GetSuiteFullName() string {
	return m.name
}

func (m *suiteMetaMockExecM) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockExecM) SetBeforeAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockExecM) SetAfterAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockExecM) GetBeforeAll() func(provider.T) {
	return m.hook
}

func (m *suiteMetaMockExecM) GetAfterAll() func(provider.T) {
	return m.hook
}

func TestAllureManager_AfterAllContext(t *testing.T) {
	manager := allureManager{suiteMeta: &suiteMetaMockExecM{container: allure.NewContainer()}}
	manager.AfterAllContext()
	require.NotNil(t, manager.executionContext)
	require.Equal(t, constants.AfterAllContextName, manager.executionContext.GetName())
}

func TestAllureManager_BeforeAllContext(t *testing.T) {
	manager := allureManager{suiteMeta: &suiteMetaMockExecM{container: allure.NewContainer()}}
	manager.BeforeAllContext()
	require.NotNil(t, manager.executionContext)
	require.Equal(t, constants.BeforeAllContextName, manager.executionContext.GetName())
}

func TestAllureManager_BeforeEachContext(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockExecM{container: allure.NewContainer()}}
	manager.BeforeEachContext()
	require.NotNil(t, manager.executionContext)
	require.Equal(t, constants.BeforeEachContextName, manager.executionContext.GetName())
}

func TestAllureManager_AfterEachContext(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockExecM{container: allure.NewContainer()}}
	manager.AfterEachContext()
	require.NotNil(t, manager.executionContext)
	require.Equal(t, constants.AfterEachContextName, manager.executionContext.GetName())
}

func TestAllureManager_TestContext(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockExecM{result: &allure.Result{}}}
	manager.TestContext()
	require.NotNil(t, manager.executionContext)
	require.Equal(t, constants.TestContextName, manager.executionContext.GetName())
}
