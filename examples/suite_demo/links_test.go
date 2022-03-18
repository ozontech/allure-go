//go:build examples_new
// +build examples_new

package suite_demo

import (
	"os"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type LinkDemoSuite struct {
	suite.Suite
}

func (s *LinkDemoSuite) BeforeAll(t provider.T) {
	_ = os.Setenv("ALLURE_TESTCASE_PATTERN", "https://tour.golang.org/welcome/%s")
	_ = os.Setenv("ALLURE_ISSUE_PATTERN", "https://pkg.go.dev/%s")
}

func (s *LinkDemoSuite) AfterAll(t provider.T) {
	_ = os.Setenv("ALLURE_TESTCASE_PATTERN", "")
	_ = os.Setenv("ALLURE_ISSUE_PATTERN", "")
}

func (s *LinkDemoSuite) TestLinks(t provider.T) {
	t.Epic("Demo")
	t.Feature("Links")
	t.Title("Test contains links")
	t.Description(`
		This test contains link with ISSUE, TEST CASE and LINK.
			Test case link: https://tour.golang.org/welcome/1
			Issue link:     https://pkg.go.dev/github.com/stretchr/testify
			Link link:      https://www.makeuseof.com/tag/8-purrfect-cat-websites/`)

	t.Tags("Links")

	t.SetTestCase("1")
	t.SetIssue("github.com/stretchr/testify")

	link := "https://www.makeuseof.com/tag/8-purrfect-cat-websites/"
	t.Link(allure.LinkLink("Demo Link", link))
}

func TestLinks(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(LinkDemoSuite))
}
