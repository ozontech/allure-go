package manager

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
)

type testMetaMockDescription struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockDescription) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockDescription) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockDescription) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockDescription) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockDescription) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockDescription) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockDescription) GetAfterEach() func(t provider.T) {
	return m.ae
}

func TestAllureManager_Title(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockDescription{result: &allure.Result{}}}
	manager.Title("Test")
	require.Equal(t, "Test", manager.testMeta.GetResult().Name)
}

func TestAllureManager_Description(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockDescription{result: &allure.Result{}}}
	manager.Description("Test")
	require.Equal(t, "Test", manager.testMeta.GetResult().Description)
}
