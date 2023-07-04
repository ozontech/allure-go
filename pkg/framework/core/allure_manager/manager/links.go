package manager

import (
	"github.com/ozontech/allure-go/pkg/allure"
)

// SetIssue adds issue link due environment variable ALLURE_ISSUE_PATTERN
func (a *allureManager) SetIssue(issue string) {
	a.Link(allure.IssueLink(issue))
}

// SetTestCase adds test case link due environment variable ALLURE_TEST_CASE_PATTERN
func (a *allureManager) SetTestCase(testCase string) {
	a.Link(allure.TestCaseLink(testCase))
}

// TmsLink adds allure external test case link due environment variable ALLURE_TMS_LINK_PATTERN
func (a *allureManager) TmsLink(testCase string) {
	a.Link(allure.TmsLink(testCase))
}

// TmsLinks adds multiple external test case links due environment variable ALLURE_TMS_LINK_PATTERN
func (a *allureManager) TmsLinks(testCase ...string) {
	a.Links(allure.TmsLinks(testCase...))
}

// Link adds Link to struct.AllureResult
func (a *allureManager) Link(link *allure.Link) {
	a.safely(func(result *allure.Result) {
		result.Links = append(result.Links, link)
	})
}

// Links adds multiple Link to struct.AllureResult
func (a *allureManager) Links(links []*allure.Link) {
	a.safely(func(result *allure.Result) {
		result.Links = append(result.Links, links...)
	})
}
