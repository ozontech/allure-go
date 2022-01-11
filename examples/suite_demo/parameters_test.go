//go:build examples
// +build examples

package suite_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type ParametersDemoSuite struct {
	suite.Suite
}

func (s *ParametersDemoSuite) TestAddParameterToStep() {
	s.Epic("Demo")
	s.Feature("Parameters")
	s.Title("Add Parameters to step")
	s.Description(`
		Step A will contain following parameters:
			Param1 = Val1
			Param2 = Val2
			Param3 = Val3
			Param4 = Val4
			Param5 = Val5
			Param6 = Val6
			Param7 = Val7`)

	s.Tags("Steps", "Nesting", "Parameters")

	step := allure.NewSimpleStep("Step A")
	// with step.AddParameter(s) function
	step.AddParameter(allure.NewParameter("Param1", "Val1"))
	step.AddParameters(allure.NewParameters("Param2", "Val2", "Param3", "Val3", "Param4", "Val4")...)

	// with step.AddNewParameter(s) function
	step.AddNewParameter("Param5", "Val5")
	step.AddNewParameters("Param6", "Val6", "Param7", "Val7")

	// don't forget register your step :)
	s.Step(step)
}

func (s *ParametersDemoSuite) TestAddParameterToNestedStep() {
	s.Epic("Demo")
	s.Feature("Parameters")
	s.Title("Add parameters to Nested Steps")
	s.Description(`
		Step A is parent step for Step B
		Step A contains Param 1 and Param 4
		Step B contains Param 2 and Param 3`)

	s.Tags("Steps", "Nesting", "Parameters")

	s.WithNewStep("Step A", func() {
		s.AddNewParameterToNested("Param 1", "Value 1")
		s.WithNewStep("Step B", func() {
			s.AddNewParametersToNested("Param 2", "Value 2", "Param 3", "Value 3")
		})
		s.AddNewParameterToNested("Param 4", "Value 3")
	})
}

func TestParameters(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(ParametersDemoSuite))
}
