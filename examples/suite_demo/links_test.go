//go:build examples
// +build examples

package suite_demo

import (
	"os"
	"testing"

	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type LinkDemoSuite struct {
	suite.Suite
}

func (s *LinkDemoSuite) BeforeAll() {
	_ = os.Setenv("ALLURE_TESTCASE_PATTERN", "https://tour.golang.org/welcome/%s")
	_ = os.Setenv("ALLURE_ISSUE_PATTERN", "https://pkg.go.dev/%s")
}

func (s *LinkDemoSuite) AfterAll() {
	_ = os.Setenv("ALLURE_TESTCASE_PATTERN", "")
	_ = os.Setenv("ALLURE_ISSUE_PATTERN", "")
}

func (s *LinkDemoSuite) TestLinks() {
	s.Epic("Demo")
	s.Feature("Links")
	s.Title("Test contains links")
	s.Description(`
		This test contains link with ISSUE, TEST CASE and LINK.
			Test case link: https://tour.golang.org/welcome/1
			Issue link:     https://pkg.go.dev/github.com/stretchr/testify
			Link link:      https://www.makeuseof.com/tag/8-purrfect-cat-websites/`)

	s.Tags("Links")

	s.SetTestCase("1")
	s.SetIssue("github.com/stretchr/testify")

	link := "https://www.makeuseof.com/tag/8-purrfect-cat-websites/"
	s.Link(allure.LinkLink("Demo Link", link))
}

func TestLinks(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(LinkDemoSuite))
}
