package manager

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type testMetaMockLabels struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockLabels) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockLabels) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockLabels) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockLabels) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockLabels) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockLabels) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockLabels) GetAfterEach() func(t provider.T) {
	return m.ae
}

func TestAllureManager_Labels(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockLabels{result: &allure.Result{}}}

	manager.ID("id")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.ID))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.ID), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.ID)[0].Value, "id")

	manager.AllureID("allureID")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.AllureID))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.AllureID), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.AllureID)[0].Value, "allureID")

	manager.Epic("epic")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Epic))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Epic), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Epic)[0].Value, "epic")

	manager.Feature("feature")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Feature))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Feature), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Feature)[0].Value, "feature")

	manager.Story("story")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Story))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Story), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Story)[0].Value, "story")

	manager.Severity(allure.TRIVIAL)
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Severity))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Severity), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Severity)[0].Value, allure.TRIVIAL.ToString())

	manager.Tag("tag1")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Tag))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Tag), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[0].Value, "tag1")

	manager.Tags("tag2", "tag3")
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Tag), 3)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[1].Value, "tag2")
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[2].Value, "tag3")

	manager.Owner("owner")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Owner))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Owner), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Owner)[0].Value, "owner")

	manager.Lead("lead")
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Lead))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Lead), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Lead)[0].Value, "lead")

	manager.Label(allure.NewLabel(allure.Framework, "Framework"))
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Framework))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Framework), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Framework)[0].Value, "Framework")

	manager.Labels(allure.NewLabel(allure.Tag, "tag4"), allure.NewLabel(allure.Tag, "tag5"))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Tag), 5)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[3].Value, "tag4")
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Tag)[4].Value, "tag5")

	manager.ReplaceLabel(allure.NewLabel(allure.Framework, "NewFramework"))
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabel(allure.Framework))
	require.Len(t, manager.testMeta.GetResult().GetLabel(allure.Framework), 1)
	require.Equal(t, manager.testMeta.GetResult().GetLabel(allure.Framework)[0].Value, "NewFramework")
}
