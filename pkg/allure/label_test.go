package allure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLabelType_ToString(t *testing.T) {
	epic := "epic"
	layer := "layer"
	feature := "feature"
	story := "story"
	as_id := "as_id"
	severity := "severity"
	parentSuite := "parentSuite"
	suite := "suite"
	subSuite := "subSuite"
	_package := "package"
	thread := "thread"
	host := "host"
	tag := "tag"
	framework := "framework"
	language := "language"
	owner := "owner"
	lead := "lead"
	allure_id := "ALLURE_ID"

	require.Equal(t, epic, Epic.ToString())
	require.Equal(t, layer, Layer.ToString())
	require.Equal(t, feature, Feature.ToString())
	require.Equal(t, story, Story.ToString())
	require.Equal(t, as_id, ID.ToString())
	require.Equal(t, severity, Severity.ToString())
	require.Equal(t, parentSuite, ParentSuite.ToString())
	require.Equal(t, suite, Suite.ToString())
	require.Equal(t, subSuite, SubSuite.ToString())
	require.Equal(t, _package, Package.ToString())
	require.Equal(t, thread, Thread.ToString())
	require.Equal(t, host, Host.ToString())
	require.Equal(t, tag, Tag.ToString())
	require.Equal(t, framework, Framework.ToString())
	require.Equal(t, language, Language.ToString())
	require.Equal(t, owner, Owner.ToString())
	require.Equal(t, lead, Lead.ToString())
	require.Equal(t, allure_id, AllureID.ToString())
}

func TestSeverityType_ToString(t *testing.T) {
	blocker := "blocker"
	critical := "critical"
	normal := "normal"
	minor := "minor"
	trivial := "trivial"

	require.Equal(t, blocker, BLOCKER.ToString())
	require.Equal(t, critical, CRITICAL.ToString())
	require.Equal(t, normal, NORMAL.ToString())
	require.Equal(t, minor, MINOR.ToString())
	require.Equal(t, trivial, TRIVIAL.ToString())
}

func TestLabelCreation(t *testing.T) {
	epic := EpicLabel("epicTest")
	layer := LayerLabel("layerTest")
	feature := FeatureLabel("featureTest")
	story := StoryLabel("storyTest")
	as_id := IDLabel("idTest")
	severity := SeverityLabel(BLOCKER)
	parentSuite := ParentSuiteLabel("parentSuiteTest")
	suite := SuiteLabel("suiteTest")
	subsuite := SubSuiteLabel("subSuiteTest")
	_package := PackageLabel("packageTest")
	thread := ThreadLabel("threadTest")
	host := HostLabel("hostTest")
	tag := TagLabel("tagTest")
	framework := FrameWorkLabel("frameWorkTest")
	language := LanguageLabel("languageTest")
	owner := OwnerLabel("ownerTest")
	lead := LeadLabel("leadTest")
	idAllure := IDAllureLabel("idAllureTest")

	require.Equal(t, epic.Name, Epic.ToString())
	require.Equal(t, layer.Name, Layer.ToString())
	require.Equal(t, feature.Name, Feature.ToString())
	require.Equal(t, story.Name, Story.ToString())
	require.Equal(t, as_id.Name, ID.ToString())
	require.Equal(t, severity.Name, Severity.ToString())
	require.Equal(t, parentSuite.Name, ParentSuite.ToString())
	require.Equal(t, suite.Name, Suite.ToString())
	require.Equal(t, subsuite.Name, SubSuite.ToString())
	require.Equal(t, _package.Name, Package.ToString())
	require.Equal(t, thread.Name, Thread.ToString())
	require.Equal(t, host.Name, Host.ToString())
	require.Equal(t, tag.Name, Tag.ToString())
	require.Equal(t, framework.Name, Framework.ToString())
	require.Equal(t, language.Name, Language.ToString())
	require.Equal(t, owner.Name, Owner.ToString())
	require.Equal(t, lead.Name, Lead.ToString())
	require.Equal(t, idAllure.Name, AllureID.ToString())

	require.Equal(t, "epicTest", epic.GetValue())
	require.Equal(t, "featureTest", feature.GetValue())
	require.Equal(t, "storyTest", story.GetValue())
	require.Equal(t, "idTest", as_id.GetValue())
	require.Equal(t, BLOCKER.ToString(), severity.GetValue())
	require.Equal(t, "parentSuiteTest", parentSuite.GetValue())
	require.Equal(t, "suiteTest", suite.GetValue())
	require.Equal(t, "subSuiteTest", subsuite.GetValue())
	require.Equal(t, "packageTest", _package.GetValue())
	require.Equal(t, "threadTest", thread.GetValue())
	require.Equal(t, "hostTest", host.GetValue())
	require.Equal(t, "tagTest", tag.GetValue())
	require.Equal(t, "frameWorkTest", framework.GetValue())
	require.Equal(t, "languageTest", language.GetValue())
	require.Equal(t, "ownerTest", owner.GetValue())
	require.Equal(t, "leadTest", lead.GetValue())
	require.Equal(t, "idAllureTest", idAllure.GetValue())
}
