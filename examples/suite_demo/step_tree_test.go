//go:build examples_new
// +build examples_new

package suite_demo

import (
	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type StepTreeDemoSuite struct {
	suite.Suite
}

func (s *StepTreeDemoSuite) TestInnerSteps(t provider.T) {
	t.Epic("Demo")
	t.Feature("Inner Steps")
	t.Title("Simple Nesting")
	t.Description(`
		Step A is parent step for Step B and Step C
		Call order will be saved in allure report
		A -> (B, C)`)

	t.Tags("Steps", "Nesting")

	t.WithNewStep("Step A", func(ctx provider.StepCtx) {
		ctx.NewStep("Step B")
		ctx.NewStep("Step C")
	})
}

func (s *StepTreeDemoSuite) TestComplexStepTree(t provider.T) {
	t.Epic("Demo")
	t.Feature("Inner Steps")
	t.Title("Complex Nesting")
	t.Description(`
		Step A is parent for Step B, Step C and Step F
		Step C is parent for Step D and Step E
		Step F is parent for Step G and Step H
		Call order will be saved in allure report
		A -> (B, C -> (D, E), F -> (G, H), I)`)

	t.Tags("Steps", "Nesting")

	t.WithNewStep("Step A", func(ctx provider.StepCtx) {
		ctx.NewStep("Step B")
		ctx.WithNewStep("Step C", func(ctx provider.StepCtx) {
			ctx.NewStep("Step D")
			ctx.NewStep("Step E")
		})
		ctx.WithNewStep("Step F", func(ctx provider.StepCtx) {
			ctx.NewStep("Step G")
			ctx.NewStep("Step H")
		})
		ctx.NewStep("Step I")
	})
}

func TestStepTree(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(StepTreeDemoSuite))
}
