//go:build provider_new
// +build provider_new

package provider_demo

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSampleDemo(t *testing.T) {
	runner.Run(t, "My test", func(t provider.T) {
		t.Epic("Only Provider Demo")
		t.Feature("runner.RunTest")

		t.Title("Some Sample test")
		t.Description("allure-go allows you to use allure without suites")
		t.WithParameters(allure.NewParameter("host", "localhost"))

		t.WithNewStep("Some nested step", func(ctx provider.StepCtx) {
			ctx.WithNewStep("Some inner step 1", func(ctx provider.StepCtx) {
				ctx.WithNewStep("Some inner step 1.1", func(ctx provider.StepCtx) {

				})
			})
			ctx.WithNewStep("Some inner step 2", func(ctx provider.StepCtx) {
				ctx.WithNewStep("Some inner step 2.1", func(ctx provider.StepCtx) {

				})
			})
		})
	}, "Sample", "Provider-only", "No provider initialization")
}

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewRunner(realT, realT.Name())

	r.BeforeEach(func(t provider.T) {
		t.NewStep(fmt.Sprintf("This is before test step for %s", t.Name()))
	})
	r.BeforeAll(func(t provider.T) {
		t.NewStep(fmt.Sprintf("This is BeforeAll test step for %s", t.Name()))
	})
	r.AfterEach(func(t provider.T) {
		t.NewStep(fmt.Sprintf("This is AfterEach test step for %s", t.Name()))
	})
	r.AfterAll(func(t provider.T) {
		t.NewStep(fmt.Sprintf("This is AfterAll test step for %s", t.Name()))
	})

	r.NewTest("My test 1", func(t provider.T) {
		t.Epic("Only Provider Demo")
		t.Feature("T.Run()")

		t.Title("Some Other Sample test")
		t.Description("allure-testify allows you to use allure without suites")

		t.WithNewStep("Some nested step", func(ctx provider.StepCtx) {
			ctx.WithNewStep("Some inner step 1", func(ctx provider.StepCtx) {
				ctx.WithNewStep("Some inner step 1.1", func(ctx provider.StepCtx) {

				})
			})
			ctx.WithNewStep("Some inner step 2", func(ctx provider.StepCtx) {
				ctx.WithNewStep("Some inner step 2.1", func(ctx provider.StepCtx) {

				})
			})
		})
	}, "Sample", "Provider-only", "with provider initialization")

	r.RunTests()
}
