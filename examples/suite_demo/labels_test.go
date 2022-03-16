//go:build examples_new
// +build examples_new

package suite_demo

import (
	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type LabelsDemoSuite struct {
	suite.Suite
}

func (s *LabelsDemoSuite) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Labels")
	t.Story("Story Label")

	//t.Host("host.example.com")
	//t.Thread("goroutine thread-example")
	//t.Package("package#example:example")
	//
	//t.FrameWork("allure-testify")
	//t.Language(runtime.Version())
	t.Owner("John Doe")
	t.Lead("John Doe's Boss")

	t.Tag("EachTestTag")
}

func (s *LabelsDemoSuite) TestLabelsExample1(t provider.T) {
	t.Title("Labels Demo Example 1")
	t.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example1"
			Severity = "blocker"
			Unique tag = "Example1"
		Also this test has additional "suite" label`)

	t.ID("Id example1")
	t.Severity(allure.BLOCKER)

	t.AddSuiteLabel("AnotherSuite")

	t.Tag("Example1")
}

func (s *LabelsDemoSuite) TestLabelsExample2(t provider.T) {
	t.Title("Labels Demo Example 2")
	t.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example2"
			Severity = "critical"
			Unique tag = "Example2"
		Also this test has additional "parentSuite" label`)

	t.ID("example2")
	t.Severity(allure.CRITICAL)

	t.AddParentSuite("AnotherParentSuite")

	t.Tag("Example2")
}

func (s *LabelsDemoSuite) TestLabelsExample3(t provider.T) {
	t.Title("Labels Demo Example 3")
	t.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example3"
			Severity = "normal"
			Unique tag = "Example3"
		Also this test has additional "subSuite" label`)

	t.ID("example3")
	t.Severity(allure.NORMAL)

	t.AddSubSuite("SomeSubSuite")

	t.Tag("Example3")
}

func (s *LabelsDemoSuite) TestLabelsExample4(t provider.T) {
	t.Title("Labels Demo Example 4")
	t.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example4"
			Severity = "minor"
			Unique tag = "Example4"`)

	t.ID("example4")
	t.Severity(allure.MINOR)

	t.Tag("Example4")
}

func (s *LabelsDemoSuite) TestLabelsExample5(t provider.T) {
	t.Title("Labels Demo Example 5")
	t.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example5"
			Severity = "trivial"
			Unique tag = "Example5"`)

	t.ID("example5")
	t.Severity(allure.TRIVIAL)

	t.Tag("Example5")
}

func TestLabels(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(LabelsDemoSuite))
}
