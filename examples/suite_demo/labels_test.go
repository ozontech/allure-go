//go:build examples
// +build examples

package suite_demo

import (
	"runtime"
	"testing"

	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type LabelsDemoSuite struct {
	suite.Suite
}

func (s *LabelsDemoSuite) BeforeEach() {
	s.Epic("Demo")
	s.Feature("Labels")
	s.Story("Story Label")

	s.Host("host.example.com")
	s.Thread("goroutine thread-example")
	s.Package("package#example:example")

	s.FrameWork("allure-testify")
	s.Language(runtime.Version())
	s.Owner("John Doe")
	s.Lead("John Doe's Boss")

	s.Tag("EachTestTag")
}

func (s *LabelsDemoSuite) TestLabelsExample1() {
	s.Title("Labels Demo Example 1")
	s.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example1"
			Severity = "blocker"
			Unique tag = "Example1"
		Also this test has additional "suite" label`)

	s.ID("Id example1")
	s.Severity(allure.BLOCKER)

	s.AddSuiteLabel("AnotherSuite")

	s.Tag("Example1")
}

func (s *LabelsDemoSuite) TestLabelsExample2() {
	s.Title("Labels Demo Example 2")
	s.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example2"
			Severity = "critical"
			Unique tag = "Example2"
		Also this test has additional "parentSuite" label`)

	s.ID("example2")
	s.Severity(allure.CRITICAL)

	s.AddParentSuite("AnotherParentSuite")

	s.Tag("Example2")
}

func (s *LabelsDemoSuite) TestLabelsExample3() {
	s.Title("Labels Demo Example 3")
	s.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example3"
			Severity = "normal"
			Unique tag = "Example3"
		Also this test has additional "subSuite" label`)

	s.ID("example3")
	s.Severity(allure.NORMAL)

	s.AddSubSuite("SomeSubSuite")

	s.Tag("Example3")
}

func (s *LabelsDemoSuite) TestLabelsExample4() {
	s.Title("Labels Demo Example 4")
	s.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example4"
			Severity = "minor"
			Unique tag = "Example4"`)

	s.ID("example4")
	s.Severity(allure.MINOR)

	s.Tag("Example4")
}

func (s *LabelsDemoSuite) TestLabelsExample5() {
	s.Title("Labels Demo Example 5")
	s.Description(`
		This Test will have all labels from SetupTest function
		Unique labels:
			ID = "example5"
			Severity = "trivial"
			Unique tag = "Example5"`)

	s.ID("example5")
	s.Severity(allure.TRIVIAL)

	s.Tag("Example5")
}

func TestLabels(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(LabelsDemoSuite))
}
