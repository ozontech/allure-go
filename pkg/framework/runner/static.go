package runner

import (
	"strings"
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/internal/file_manager"
	"github.com/koodeex/allure-testify/pkg/framework/internal/framework"
	"github.com/koodeex/allure-testify/pkg/provider"
)

// RunSuite Finds all methods of the passed structure that begin with the prefix `Test` and executes them.
// If the structure implements methods `BeforeEach`, `BeforeAll`, `AfterEach` and/or `AfterAll`,
// they will also be executed.
func RunSuite(realT *testing.T, suite framework.TestSuite) {
	var packageName string
	if suite.GetPackage() == "" {
		packageName = file_manager.GetPackage(2)
	} else {
		packageName = suite.GetPackage()
	}
	suiteToRun := framework.NewInternalSuite(realT, packageName, suite)
	suiteToRun.Run(realT)
}

// RunTest Runs the test passed in the `f` argument.
// At the end of execution `result.json` and `container.json` will be created (if not empty).
func RunTest(realT *testing.T, testName string, f func(t *provider.T), tags ...string) bool {
	callers := strings.Split(realT.Name(), "/")
	t := provider.NewTWithPackage(realT, callers[0], callers[len(callers)-1], file_manager.GetPackage(2))
	return t.Run(testName, f, tags...)
}
