//go:build provider
// +build provider

package provider_demo

import (
	"fmt"
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/provider"
)

func TestSampleDemo(t *testing.T) {
	runner.RunTest(t, "My test", func(t *provider.T) {
		t.Epic("Only Provider Demo")
		t.Feature("runner.RunTest")

		t.Title("Some Sample test")
		t.Description("allure-testify allows you to use allure without suites")

		t.WithNewStep("Some nested step", func() {
			t.WithNewStep("Some inner step 1", func() {
				t.WithNewStep("Some inner step 1.1", func() {

				})
			})
			t.WithNewStep("Some inner step 2", func() {
				t.WithNewStep("Some inner step 2.1", func() {

				})
			})
		})
	}, "Sample", "Provider-only", "No provider initialization")
}

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewTestRunner(realT)

	r.WithBeforeEach(func(t *provider.T) {
		t.NewStep(fmt.Sprintf("This is before test step for %s", t.Name()))
	})
	r.WithAfterEach(func(t *provider.T) {
		t.NewStep(fmt.Sprintf("This is after test step for %s", t.Name()))
	})

	r.Run("My test 1", func(t *provider.T) {
		t.Epic("Only Provider Demo")
		t.Feature("T.Run()")

		t.Title("Some Other Sample test")
		t.Description("allure-testify allows you to use allure without suites")

		t.WithNewStep("Some nested step", func() {
			t.WithNewStep("Some inner step 1", func() {
				t.WithNewStep("Some inner step 1.1", func() {

				})
			})
			t.WithNewStep("Some inner step 2", func() {
				t.WithNewStep("Some inner step 2.1", func() {

				})
			})
		})
	}, "Sample", "Provider-only", "with provider initialization")
}
