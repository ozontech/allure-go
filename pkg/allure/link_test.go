package allure

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLinkTypes(t *testing.T) {
	link := "link"
	issue := "issue"
	testCase := "test_case"
	tms := "tms"

	require.Equal(t, link, string(LINK))
	require.Equal(t, issue, string(ISSUE))
	require.Equal(t, testCase, string(TESTCASE))
	require.Equal(t, tms, string(TMS))
}

func TestNewLink(t *testing.T) {
	link := NewLink("testLink", LINK, "https://www.testLink.com")
	issue := NewLink("issueLink", ISSUE, "https://www.testIssue.com")
	testCase := NewLink("testCaseLink", TESTCASE, "https://www.testCase.com")
	tms := NewLink("tmsLink", TMS, "https://www.tmslink.com")

	require.NotNil(t, link)
	require.Equal(t, "testLink", link.Name)
	require.Equal(t, string(LINK), link.Type)
	require.Equal(t, "https://www.testLink.com", link.URL)

	require.NotNil(t, issue)
	require.Equal(t, "issueLink", issue.Name)
	require.Equal(t, string(ISSUE), issue.Type)
	require.Equal(t, "https://www.testIssue.com", issue.URL)

	require.NotNil(t, testCase)
	require.Equal(t, "testCaseLink", testCase.Name)
	require.Equal(t, string(TESTCASE), testCase.Type)
	require.Equal(t, "https://www.testCase.com", testCase.URL)

	require.NotNil(t, tms)
	require.Equal(t, "tmsLink", tms.Name)
	require.Equal(t, string(TMS), tms.Type)
	require.Equal(t, "https://www.tmslink.com", tms.URL)
}

func TestTestCaseLink_noEnv(t *testing.T) {
	testCase := TestCaseLink("TEST-112")
	require.NotNil(t, testCase)
	require.Equal(t, "TestCase[TEST-112]", testCase.Name)
	require.Equal(t, string(TESTCASE), testCase.Type)
	require.Equal(t, "TEST-112", testCase.URL)
}

func TestTestCaseLink_Env(t *testing.T) {
	os.Setenv(testCasePatternEnvKey, "https://jira-mock.com/%s")
	defer os.Setenv(testCasePatternEnvKey, "")
	testCase := TestCaseLink("TEST-112")
	require.NotNil(t, testCase)
	require.Equal(t, "TestCase[TEST-112]", testCase.Name)
	require.Equal(t, string(TESTCASE), testCase.Type)
	require.Equal(t, "https://jira-mock.com/TEST-112", testCase.URL)
}

func TestIssueLink_noEnv(t *testing.T) {
	testCase := IssueLink("TEST-112")
	require.NotNil(t, testCase)
	require.Equal(t, "Issue[TEST-112]", testCase.Name)
	require.Equal(t, string(ISSUE), testCase.Type)
	require.Equal(t, "TEST-112", testCase.URL)
}

func TestIssueLink_Env(t *testing.T) {
	os.Setenv(issuePatternEnvKey, "https://jira-mock.com/%s")
	defer os.Setenv(issuePatternEnvKey, "")
	testCase := IssueLink("TEST-112")
	require.NotNil(t, testCase)
	require.Equal(t, "Issue[TEST-112]", testCase.Name)
	require.Equal(t, string(ISSUE), testCase.Type)
	require.Equal(t, "https://jira-mock.com/TEST-112", testCase.URL)
}

func TestTmsLink_noEnv(t *testing.T) {
	testCase := TmsLink("TMS-112")
	require.NotNil(t, testCase)
	require.Equal(t, "TMS-112", testCase.Name)
	require.Equal(t, string(TMS), testCase.Type)
	require.Equal(t, "TMS-112", testCase.URL)
}

func TestTmsLink_Env(t *testing.T) {
	os.Setenv(tmsLinkPatternEnvKey, "https://tms-mock.com/%s")
	defer os.Setenv(tmsLinkPatternEnvKey, "")
	testCase := TmsLink("TMS-112")
	require.NotNil(t, testCase)
	require.Equal(t, "TMS-112", testCase.Name)
	require.Equal(t, string(TMS), testCase.Type)
	require.Equal(t, "https://tms-mock.com/TMS-112", testCase.URL)
}

func TestTmsLinks_noEnv(t *testing.T) {
	testCase := TmsLinks("TMS-110", "TMS-112")
	require.NotNil(t, testCase)
	require.Len(t, testCase, 2)
	require.Equal(t, "TMS-110", testCase[0].Name)
	require.Equal(t, "TMS-112", testCase[1].Name)
	require.Equal(t, string(TMS), testCase[0].Type)
	require.Equal(t, string(TMS), testCase[1].Type)
	require.Equal(t, "TMS-110", testCase[0].URL)
	require.Equal(t, "TMS-112", testCase[1].URL)
}

func TestTmsLinks_Env(t *testing.T) {
	os.Setenv(tmsLinkPatternEnvKey, "https://tms-mock.com/%s")
	defer os.Setenv(tmsLinkPatternEnvKey, "")
	testCase := TmsLinks("TEST-110", "TEST-112")
	require.NotNil(t, testCase)
	require.Len(t, testCase, 2)
	require.Equal(t, "TEST-110", testCase[0].Name)
	require.Equal(t, "TEST-112", testCase[1].Name)
	require.Equal(t, string(TMS), testCase[0].Type)
	require.Equal(t, string(TMS), testCase[1].Type)
	require.Equal(t, "https://tms-mock.com/TEST-110", testCase[0].URL)
	require.Equal(t, "https://tms-mock.com/TEST-112", testCase[1].URL)
}
