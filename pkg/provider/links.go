package provider

import (
	"github.com/koodeex/allure-testify/pkg/allure"
)

/*
Links
*/

type AllureLinks interface {
	SetIssue(issue string)
	SetTestCase(testCase string)
	Link(link allure.Link)
}

// SetIssue adds issue link due environment variable ALLURE_ISSUE_PATTERN
func (t *T) SetIssue(issue string) {
	t.Link(allure.IssueLink(issue))
}

// SetTestCase adds test case link due environment variable ALLURE_TEST_CASE_PATTERN
func (t *T) SetTestCase(testCase string) {
	t.Link(allure.TestCaseLink(testCase))
}

// Link adds Link to struct.AllureResult
func (t *T) Link(link allure.Link) {
	t.safely(func(result *allure.Result) {
		result.Links = append(result.Links, link)
	})
}
