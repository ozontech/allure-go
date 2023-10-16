package manager

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/stretchr/testify/require"
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

	t.Run("ID", func(t *testing.T) {
		manager.ID("id")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.ID))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.ID), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.ID)[0].GetValue(), "id")
	})

	t.Run("AllureID", func(t *testing.T) {
		manager.AllureID("allureID")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.AllureID))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.AllureID), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.AllureID)[0].GetValue(), "allureID")
	})

	t.Run("Epic", func(t *testing.T) {
		manager.Epic("epic")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Epic))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Epic), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Epic)[0].GetValue(), "epic")
	})

	t.Run("Feature", func(t *testing.T) {
		manager.Feature("feature")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Feature))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Feature), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Feature)[0].GetValue(), "feature")
	})

	t.Run("Story", func(t *testing.T) {
		manager.Story("story")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Story))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Story), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Story)[0].GetValue(), "story")
	})

	t.Run("Severity", func(t *testing.T) {
		manager.Severity(allure.TRIVIAL)
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Severity))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Severity), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Severity)[0].GetValue(), allure.TRIVIAL.ToString())
	})

	t.Run("Tag", func(t *testing.T) {
		manager.Tag("tag1")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Tag))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Tag), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Tag)[0].GetValue(), "tag1")
	})

	t.Run("Tags", func(t *testing.T) {
		manager.Tags("tag2", "tag3")
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Tag), 3)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Tag)[1].GetValue(), "tag2")
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Tag)[2].GetValue(), "tag3")
	})

	t.Run("Owner", func(t *testing.T) {
		manager.Owner("owner")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Owner))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Owner), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Owner)[0].GetValue(), "owner")
	})

	t.Run("Lead", func(t *testing.T) {
		manager.Lead("lead")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Lead))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Lead), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Lead)[0].GetValue(), "lead")
	})

	t.Run("Label", func(t *testing.T) {
		manager.Label(allure.NewLabel(allure.Framework, "Framework"))
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Framework))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Framework), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Framework)[0].GetValue(), "Framework")
	})

	t.Run("Labels", func(t *testing.T) {
		manager.Labels(allure.NewLabel(allure.Tag, "tag4"), allure.NewLabel(allure.Tag, "tag5"))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Tag), 5)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Tag)[3].GetValue(), "tag4")
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Tag)[4].GetValue(), "tag5")
	})

	t.Run("AddSuiteLabel", func(t *testing.T) {
		manager.AddSuiteLabel("Suite")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Suite))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Suite), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Suite)[0].GetValue(), "Suite")
	})

	t.Run("AddSubSuite", func(t *testing.T) {
		manager.AddSubSuite("SubSuite")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.SubSuite))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.SubSuite), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.SubSuite)[0].GetValue(), "SubSuite")
	})

	t.Run("AddParentSuite", func(t *testing.T) {
		manager.AddParentSuite("ParentSuite")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.ParentSuite))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.ParentSuite), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.ParentSuite)[0].GetValue(), "ParentSuite")
	})

	t.Run("ID", func(t *testing.T) {
		manager.Host("Host")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Host))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Host), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Host)[0].GetValue(), "Host")
	})

	t.Run("Thread", func(t *testing.T) {
		manager.Thread("Thread")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Thread))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Thread), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Thread)[0].GetValue(), "Thread")
	})

	t.Run("Language", func(t *testing.T) {
		manager.Language("Language")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Language))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Language), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Language)[0].GetValue(), "Language")
	})

	t.Run("Package", func(t *testing.T) {
		manager.Package("Package")
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Package))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Package), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Package)[0].GetValue(), "Package")
	})

	t.Run("ReplaceLabel", func(t *testing.T) {
		manager.ReplaceLabel(allure.NewLabel(allure.Framework, "NewFramework"))
		require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Framework))
		require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Framework), 1)
		require.Equal(t, manager.testMeta.GetResult().GetLabels(allure.Framework)[0].GetValue(), "NewFramework")
	})
}
