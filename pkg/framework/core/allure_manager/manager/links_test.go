package manager

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
)

type testMetaMockLinks struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockLinks) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockLinks) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockLinks) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockLinks) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockLinks) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockLinks) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockLinks) GetAfterEach() func(t provider.T) {
	return m.ae
}

func TestAllureManager_Link(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockLinks{result: &allure.Result{}}}
	manager.Link(allure.NewLink("Name", allure.LINK, "http://test.com"))
	require.Len(t, manager.GetResult().Links, 1)
	require.Equal(t, "Name", manager.GetResult().Links[0].Name)
	require.Equal(t, string(allure.LINK), manager.GetResult().Links[0].Type)
	require.Equal(t, "http://test.com", manager.GetResult().Links[0].URL)

}

func TestAllureManager_SetTestCase(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockLinks{result: &allure.Result{}}}
	manager.SetTestCase("TestCase")
	require.Len(t, manager.GetResult().Links, 1)
	require.Equal(t, "TestCase[TestCase]", manager.GetResult().Links[0].Name)
	require.Equal(t, string(allure.TESTCASE), manager.GetResult().Links[0].Type)
}

func TestAllureManager_SetIssue(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockLinks{result: &allure.Result{}}}
	manager.SetIssue("Issue")
	require.NotEmpty(t, manager.GetResult().Links)
	require.Len(t, manager.GetResult().Links, 1)
	require.Equal(t, "Issue[Issue]", manager.GetResult().Links[0].Name)
	require.Equal(t, string(allure.ISSUE), manager.GetResult().Links[0].Type)
}
