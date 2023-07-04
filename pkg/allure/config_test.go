package allure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	require.Equal(t, "ALLURE_OUTPUT_PATH", resultsPathEnvKey)
	require.Equal(t, "ALLURE_OUTPUT_FOLDER", outputFolderEnvKey)
	require.Equal(t, "ALLURE_ISSUE_PATTERN", issuePatternEnvKey)
	require.Equal(t, "ALLURE_TESTCASE_PATTERN", testCasePatternEnvKey)
	require.Equal(t, "ALLURE_LAUNCH_TAGS", defaultTagsEnvKey)
	require.Equal(t, "ALLURE_LINK_TMS_PATTERN", tmsLinkPatternEnvKey)
	require.Equal(t, 0644, fileSystemPermissionCode)
}
