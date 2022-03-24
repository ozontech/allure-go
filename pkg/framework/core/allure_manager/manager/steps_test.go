package manager

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
	"testing"
)

type testMetaMockSteps struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockSteps) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockSteps) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockSteps) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockSteps) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockSteps) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockSteps) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockSteps) GetAfterEach() func(t provider.T) {
	return m.ae
}

type suiteMetaMockSteps struct {
	name      string
	container *allure.Container
	hook      func(t provider.T)
}

func (m *suiteMetaMockSteps) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockSteps) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockSteps) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockSteps) GetSuiteFullName() string {
	return m.name
}

func (m *suiteMetaMockSteps) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockSteps) SetBeforeAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockSteps) SetAfterAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockSteps) GetBeforeAll() func(provider.T) {
	return m.hook
}

func (m *suiteMetaMockSteps) GetAfterAll() func(provider.T) {
	return m.hook
}

func TestAllureManager_Step(t *testing.T) {
	manager := allureManager{
		testMeta:  &testMetaMockLabels{result: &allure.Result{}, container: allure.NewContainer()},
		suiteMeta: &suiteMetaMockExecM{container: allure.NewContainer()},
	}
	testStep := allure.NewSimpleStep("Step")
	manager.TestContext()
	manager.Step(testStep)
	require.NotEmpty(t, manager.testMeta.GetResult().Steps)
	require.Len(t, manager.testMeta.GetResult().Steps, 1)
	require.Equal(t, manager.testMeta.GetResult().Steps[0], testStep)

	manager.BeforeEachContext()
	manager.Step(testStep)
	require.NotEmpty(t, manager.testMeta.GetContainer().Befores)
	require.Len(t, manager.testMeta.GetContainer().Befores, 1)
	require.Equal(t, manager.testMeta.GetContainer().Befores[0], testStep)

	manager.BeforeAllContext()
	manager.Step(testStep)
	require.NotEmpty(t, manager.suiteMeta.GetContainer().Befores)
	require.Len(t, manager.suiteMeta.GetContainer().Befores, 1)
	require.Equal(t, manager.suiteMeta.GetContainer().Befores[0], testStep)

	manager.AfterEachContext()
	manager.Step(testStep)
	require.NotEmpty(t, manager.testMeta.GetContainer().Afters)
	require.Len(t, manager.testMeta.GetContainer().Afters, 1)
	require.Equal(t, manager.testMeta.GetContainer().Afters[0], testStep)

	manager.AfterAllContext()
	manager.Step(testStep)
	require.NotEmpty(t, manager.suiteMeta.GetContainer().Afters)
	require.Len(t, manager.suiteMeta.GetContainer().Afters, 1)
	require.Equal(t, manager.suiteMeta.GetContainer().Afters[0], testStep)
}

func TestAllureManager_NewStep(t *testing.T) {
	manager := allureManager{
		testMeta:  &testMetaMockLabels{result: &allure.Result{}, container: allure.NewContainer()},
		suiteMeta: &suiteMetaMockExecM{container: allure.NewContainer()},
	}
	manager.TestContext()
	manager.NewStep("testStep")
	require.NotEmpty(t, manager.testMeta.GetResult().Steps)
	require.Len(t, manager.testMeta.GetResult().Steps, 1)
	require.Equal(t, manager.testMeta.GetResult().Steps[0].Name, "testStep")

	manager.BeforeEachContext()
	manager.NewStep("testStep")
	require.NotEmpty(t, manager.testMeta.GetContainer().Befores)
	require.Len(t, manager.testMeta.GetContainer().Befores, 1)
	require.Equal(t, manager.testMeta.GetContainer().Befores[0].Name, "testStep")

	manager.BeforeAllContext()
	manager.NewStep("testStep")
	require.NotEmpty(t, manager.suiteMeta.GetContainer().Befores)
	require.Len(t, manager.suiteMeta.GetContainer().Befores, 1)
	require.Equal(t, manager.suiteMeta.GetContainer().Befores[0].Name, "testStep")

	manager.AfterEachContext()
	manager.NewStep("testStep")
	require.NotEmpty(t, manager.testMeta.GetContainer().Afters)
	require.Len(t, manager.testMeta.GetContainer().Afters, 1)
	require.Equal(t, manager.testMeta.GetContainer().Afters[0].Name, "testStep")

	manager.AfterAllContext()
	manager.NewStep("testStep")
	require.NotEmpty(t, manager.suiteMeta.GetContainer().Afters)
	require.Len(t, manager.suiteMeta.GetContainer().Afters, 1)
	require.Equal(t, manager.suiteMeta.GetContainer().Afters[0].Name, "testStep")
}
