//go:build examples_new
// +build examples_new

package suite_demo

import (
	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type ParametersDemoSuite struct {
	suite.Suite
}

func (s *ParametersDemoSuite) TestAddParameterToStep(t provider.T) {
	t.Epic("Demo")
	t.Feature("Parameters")
	t.Title("Add Parameters to step")
	t.Description(`
		Step A will contain following parameters:
			Param1 = Val1
			Param2 = Val2
			Param3 = Val3
			Param4 = Val4
			Param5 = Val5
			Param6 = Val6
			Param7 = Val7`)

	t.Tags("Steps", "Nesting", "Parameters")

	step := allure.NewSimpleStep("Step A")

	// with step.WithParameters(s) function
	step.WithParameters(allure.NewParameter("Param1", "Val1"))
	step.WithParameters(allure.NewParameters("Param2", "Val2", "Param3", "Val3", "Param4", "Val4")...)

	// don't forget register your step :)
	t.Step(step)
}

func (s *ParametersDemoSuite) TestAddParameterToNestedStep(t provider.T) {
	t.Epic("Demo")
	t.Feature("Parameters")
	t.Title("Add parameters to Nested Steps")
	t.Description(`
		Step A is parent step for Step B
		Step A contains Param 1 and Param 4
		Step B contains Param 2 and Param 3`)

	t.Tags("Steps", "Nesting", "Parameters")

	t.WithNewStep("Step A", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Param 1", "Value 1")
		ctx.WithNewStep("Step B", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Param 2", "Value 2", "Param 3", "Value 3")
		})
		ctx.WithNewParameters("Param 4", "Value 3")
	})
}

func TestParameters(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(ParametersDemoSuite))
}
