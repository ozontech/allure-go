package suite

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/internal/framework"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

// RunSuite forward of runner.RunSuite function
func RunSuite(t *testing.T, suite framework.TestSuite) {
	runner.RunSuite(t, suite)
}
