package allure

import (
	"fmt"
	"os"
	"strings"
)

const (
	// DefaultVersion - allure-testify current Version
	DefaultVersion = "Allure-Testify@v0.3.1"

	resultsPathEnvKey     = "ALLURE_OUTPUT_PATH"      // Indicates the path to the results print folder
	outputFolderEnvKey    = "ALLURE_OUTPUT_FOLDER"    // Indicates the name of the folder to print the results.
	issuePatternEnvKey    = "ALLURE_ISSUE_PATTERN"    // Indicates the URL pattern for Issue. It must contain exactly one `%s`
	testCasePatternEnvKey = "ALLURE_TESTCASE_PATTERN" // Indicates the URL pattern for TestCase. It must contain exactly one `%s`

	fileSystemPermissionCode = 0644 // Attachment permission
)

var resultsPath = getResultPath()
var outputFolder = getOutputFolderName()

func getOutputFolderName() string {
	outputFolderName := os.Getenv(outputFolderEnvKey)
	if outputFolderName != "" {
		return outputFolderName
	}
	return "allure-results"
}

func getResultPath() string {
	resultsPathToOutput := os.Getenv(resultsPathEnvKey)
	if resultsPathToOutput != "" {
		return fmt.Sprintf("%s/%s", resultsPathToOutput, outputFolder)
	}
	return fmt.Sprintf("./%s", outputFolder)
}

func getIssuePattern() string {
	return getPattern(issuePatternEnvKey, "%s")
}

func getTestCasePattern() string {
	return getPattern(testCasePatternEnvKey, "%s")
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
