//go:build examples
// +build examples

package suite_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type SkipDemoSuite struct {
	suite.Suite
}

func (s *SkipDemoSuite) TestSkip() {
	s.Epic("Demo")
	s.Feature("Skip Test")
	s.Title("Skip test")
	s.Description(`
		This test will be skipped`)

	s.Tags("Test", "Skip")
	s.T().Skip("Skip Reason")
}

func TestSkipDemo(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(SkipDemoSuite))
}
