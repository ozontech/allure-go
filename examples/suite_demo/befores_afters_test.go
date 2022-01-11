//go:build examples
// +build examples

package suite_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type BeforeAfterDemoSuite struct {
	suite.Suite
}

func (s *BeforeAfterDemoSuite) BeforeEach() {
	s.NewStep("Before Test Step")
}

func (s *BeforeAfterDemoSuite) AfterEach() {
	s.NewStep("After Test Step")
}

func (s *BeforeAfterDemoSuite) BeforeAll() {
	s.NewStep("Before suite Step")
}

func (s *BeforeAfterDemoSuite) AfterAll() {
	s.NewStep("After suite Step")
}

func (s *BeforeAfterDemoSuite) TestBeforeAfterTest() {
	s.Epic("Demo")
	s.Feature("BeforeAfter")
	s.Title("Test wrapped with SetUp & TearDown")
	s.Description(`
		This test wrapped with SetUp and TearDown containers.`)

	s.Tags("BeforeAfter")
}

func TestBeforesAfters(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(BeforeAfterDemoSuite))
}
