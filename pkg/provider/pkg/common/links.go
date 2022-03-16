package common

import "github.com/ozontech/allure-go/pkg/allure"

// SetIssue adds issue link due environment variable ALLURE_ISSUE_PATTERN
func (c *common) SetIssue(issue string) {
	c.Link(allure.IssueLink(issue))
}

// SetTestCase adds test case link due environment variable ALLURE_TEST_CASE_PATTERN
func (c *common) SetTestCase(testCase string) {
	c.Link(allure.TestCaseLink(testCase))
}

// Link adds Link to struct.AllureResult
func (c *common) Link(link allure.Link) {
	c.safely(func(result *allure.Result) {
		result.Links = append(result.Links, link)
	})
}
