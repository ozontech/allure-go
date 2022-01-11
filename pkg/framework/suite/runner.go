package suite

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/internal/framework"
	"github.com/koodeex/allure-testify/pkg/framework/runner"
)

// RunSuite forward of runner.RunSuite function
func RunSuite(t *testing.T, suite framework.TestSuite) {
	runner.RunSuite(t, suite)
}
