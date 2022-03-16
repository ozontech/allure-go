//go:build allure_go_new
// +build allure_go_new

package allure_go_compare

import (
	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type SuiteStruct struct {
	suite.Suite
}

/* Allure-Go style:
func TestNewTest(t *testing.T) {
	allure.Test(
		t,
		allure.Description("New Test Description"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Step description"),
				allure.Action(func() {

				}))
		}))
}
*/

func (s *SuiteStruct) TestNewTest(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("New Test Description")
	t.WithNewStep("Step description", func(ctx provider.StepCtx) {

	})
}

/* Allure-Go style:
func TestWithIntricateSubsteps(t *testing.T) {
	allure.Test(t, allure.Description("Test"),
		allure.Action(func() {
			allure.Step(allure.Description("Step 1"), allure.Action(func() {
				doSomething()
				allure.Step(allure.Description("Sub-step 1.1"), allure.Action(func() {
					t.Errorf("Failure")
				}))
				allure.Step(allure.Description("Sub-step 1.2"), allure.Action(func() {}))
				allure.SkipStep(allure.Description("Sub-step 1.3"), allure.Reason("Skip this step because of defect to be fixed"), allure.Action(func() {}))
			}))
			allure.Step(allure.Description("Step 2"), allure.Action(func() {
				allure.Step(allure.Description("Sub-step 2.1"), allure.Action(func() {
					allure.Step(allure.Description("Step 2.1.1"), allure.Action(func() {
						allure.Step(allure.Description("Sub-step 2.1.1.1"), allure.Action(func() {
							t.Errorf("Failure")
						}))
						allure.Step(allure.Description("Sub-step 2.1.1.2"), allure.Action(func() {
							t.Error("Failed like this")
						}))
					}))
				}))
				allure.Step(allure.Description("Sub-step 2.2"), allure.Action(func() {}))
			}))
		}))
}
*/

func doSomething(ctx provider.StepCtx) {
	ctx.WithNewStep("Something", func(ctx provider.StepCtx) {
		ctx.WithNewStep("Because We Can!", func(ctx provider.StepCtx) {

		})
	})
}

// Works even better! each step will have his real status, but parent will be failed anyway

func (s *SuiteStruct) TestWithIntricateSubsteps(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("Test")

	t.WithNewStep("Step 1", func(ctx provider.StepCtx) {
		doSomething(ctx)
		ctx.WithNewStep("Sub-step 1.1", func(ctx provider.StepCtx) {
			ctx.Error("Failure Sub-step 1.1")
		})
		ctx.WithNewStep("Sub-step 1.2", func(ctx provider.StepCtx) {

		})
		ctx.WithNewStep("Sub-step 1.3", func(ctx provider.StepCtx) {

		})
	})
	t.WithNewStep("Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewStep("Sub-Step 2.1", func(ctx provider.StepCtx) {
			ctx.WithNewStep("Sub-step 2.1.1", func(ctx provider.StepCtx) {
				ctx.WithNewStep("Sub-step 2.1.1.1", func(ctx provider.StepCtx) {
					ctx.Error("Failure Sub-step 2.1.1.1")
				})
				ctx.WithNewStep("Sub-step 2.1.1.2", func(ctx provider.StepCtx) {
					ctx.Error("This will be in report-status. Sub-step 2.1.1.2")
				})
			})
		})
		ctx.WithNewStep("Sub-Step 2.2", func(ctx provider.StepCtx) {

		})
	})
}

func TestRun(t *testing.T) {
	suite.RunSuite(t, new(SuiteStruct))
}
