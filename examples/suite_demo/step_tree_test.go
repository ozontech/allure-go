//go:build examples
// +build examples

package suite_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type StepTreeDemoSuite struct {
	suite.Suite
}

func (s *StepTreeDemoSuite) TestInnerSteps() {
	s.Epic("Demo")
	s.Feature("Inner Steps")
	s.Title("Simple Nesting")
	s.Description(`
		Step A is parent step for Step B and Step C
		Call order will be saved in allure report
		A -> (B, C)`)

	s.Tags("Steps", "Nesting")

	s.WithNewStep("Step A", func() {
		s.NewStep("Step B")
		s.NewStep("Step C")
	})
}

func (s *StepTreeDemoSuite) TestComplexStepTree() {
	s.Epic("Demo")
	s.Feature("Inner Steps")
	s.Title("Complex Nesting")
	s.Description(`
		Step A is parent for Step B, Step C and Step F
		Step C is parent for Step D and Step E
		Step F is parent for Step G and Step H
		Call order will be saved in allure report
		A -> (B, C -> (D, E), F -> (G, H), I)`)

	s.Tags("Steps", "Nesting")

	s.WithNewStep("Step A", func() {
		s.NewStep("Step B")
		s.WithNewStep("Step C", func() {
			s.NewStep("Step D")
			s.NewStep("Step E")
		})
		s.WithNewStep("Step F", func() {
			s.NewStep("Step G")
			s.NewStep("Step H")
		})
		s.NewStep("Step I")
	})
}

func TestStepTree(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(StepTreeDemoSuite))
}
