//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

//TestRunningDemoSuite demonstrate parallel running
type TestRunningDemoSuite struct {
	suite.Suite
}

func (s *TestRunningDemoSuite) TestBeforesAfters(t provider.T) {
	t.Parallel()
	s.RunSuite(t, new(BeforeAfterDemoSuite))
}

func (s *TestRunningDemoSuite) TestFails(t provider.T) {
	t.Parallel()
	s.RunSuite(t, new(FailsDemoSuite))
}

func (s *TestRunningDemoSuite) TestLabels(t provider.T) {
	t.Parallel()
	s.RunSuite(t, new(LabelsDemoSuite))
}

func TestRunDemo(t *testing.T) {
	// use RunSuites to run suite of suites
	t.Parallel()
	suite.RunSuite(t, new(TestRunningDemoSuite))
}
