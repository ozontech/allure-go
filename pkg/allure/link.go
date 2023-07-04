package allure

import (
	"fmt"
	"os"
	"strings"
)

// Link is an implementation of the Link entity used by Allure to specify the links needed for test reports.
// Such as:
// - A link to a task in Issue tracker.
// - A link to a test case in the TMS
// - Any other link (e.g. a link to an environment pod)
type Link struct {
	Name string `json:"name"` // Link name
	Type string `json:"type"` // Link's Type (issue, test case or any other)
	URL  string `json:"url"`  // Link URL
}

// LinkTypes ...
type LinkTypes string

// LinkTypes constants
const (
	LINK     LinkTypes = "link"
	ISSUE    LinkTypes = "issue"
	TESTCASE LinkTypes = "test_case"
	TMS      LinkTypes = "tms"
)

// NewLink Constructor. Builds and returns a new `allure.Link` object.
func NewLink(name string, linkType LinkTypes, url string) *Link {
	return &Link{name, string(linkType), url}
}

// TestCaseLink returns TESTCASE type link
func TestCaseLink(testCase string) *Link {
	linkName := fmt.Sprintf("TestCase[%s]", testCase)
	return NewLink(linkName, TESTCASE, fmt.Sprintf(getTestCasePattern(), testCase))
}

// IssueLink returns ISSUE type link
func IssueLink(issue string) *Link {
	linkName := fmt.Sprintf("Issue[%s]", issue)
	return NewLink(linkName, ISSUE, fmt.Sprintf(getIssuePattern(), issue))
}

// LinkLink returns LINK type link
func LinkLink(linkname, link string) *Link {
	return NewLink(linkname, LINK, link)
}

// TmsLink returns TMS type link
func TmsLink(testCase string) *Link {
	return NewLink(testCase, TMS, fmt.Sprintf(getTmsPattern(), testCase))
}

// TmsLinks returns multiple TmsLink type link
func TmsLinks(testCases ...string) []*Link {
	var result []*Link
	for _, testCase := range testCases {
		result = append(result, NewLink(testCase, TMS, fmt.Sprintf(getTmsPattern(), testCase)))
	}
	return result
}

func getIssuePattern() string {
	return getPattern(issuePatternEnvKey, "%s")
}

func getTestCasePattern() string {
	return getPattern(testCasePatternEnvKey, "%s")
}

func getTmsPattern() string {
	return getPattern(tmsLinkPatternEnvKey, "%s")
}

func getPattern(envKey string, defaultPattern string) string {
	pattern := os.Getenv(envKey)
	if pattern == "" && !strings.Contains(pattern, "%s") {
		fmt.Printf("Provided pattern (%s) not contains '%%s' or empty.\n", pattern)
		fmt.Printf("Please provide correct one. Use %s environment variable.\n", envKey)
		fmt.Printf("Until this default pattern will be used (%s).\n", defaultPattern)
		return defaultPattern
	}
	return pattern
}
