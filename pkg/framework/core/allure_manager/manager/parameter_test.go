package manager

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type testMetaMockParameter struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockParameter) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockParameter) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockParameter) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockParameter) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockParameter) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockParameter) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockParameter) GetAfterEach() func(t provider.T) {
	return m.ae
}

func TestAllureManager_Parameter(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockParameter{result: &allure.Result{}}}
	manager.WithParameters(allure.NewParameter("host", "localhost"))
	require.Len(t, manager.GetResult().Parameters, 1)
	require.Equal(t, "host", manager.GetResult().Parameters[0].Name)
	require.Equal(t, "localhost", manager.GetResult().Parameters[0].Value)
}

func TestAllureManager_NewParameter(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockParameter{result: &allure.Result{}}}
	manager.WithNewParameters("host", "localhost", "os", "linux")
	require.Len(t, manager.GetResult().Parameters, 2)
	require.Equal(t, "host", manager.GetResult().Parameters[0].Name)
	require.Equal(t, "localhost", manager.GetResult().Parameters[0].Value)
	require.Equal(t, "os", manager.GetResult().Parameters[1].Name)
	require.Equal(t, "linux", manager.GetResult().Parameters[1].Value)
}
