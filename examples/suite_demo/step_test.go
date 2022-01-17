//go:build examples
// +build examples

package suite_demo

import (
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type StepDemoSuite struct {
	suite.Suite
}

func (s *StepDemoSuite) TestAddSteps() {
	s.Epic("Demo")
	s.Feature("Steps")
	s.Title("Base. Add steps to allure result")
	s.Description(`
		Step A, Step B and Step C will be add to Allure Result`)

	s.Tags("Steps")

	stepA := allure.NewSimpleStep("Step A")
	s.Step(stepA)

	stepB := allure.NewStep("Step B", // Step's Name
		allure.Passed,                           // Step Status
		allure.GetNow(),                         // Step Start
		allure.GetNow(),                         // Step Finish
		allure.NewParameters("paramB", "value")) // Step Parameters

	s.Step(stepB)

	stepC := allure.NewSimpleStep("Step C")
	stepC.Start = allure.GetNow()
	stepC.AddNewParameter("paramC", "value")
	stepC.Stop = allure.GetNow()
	s.Step(stepC)
}

func (s *StepDemoSuite) TestQuickWorkWithSteps() {
	s.Epic("Demo")
	s.Feature("Steps")
	s.Title("Base. Add steps to allure result")
	s.Description(`
		Step A, Step B, Step C and Step D will be add to Allure Result`)

	s.Tags("Steps")

	s.Step(allure.NewSimpleStep("Step A").Passed())  // This step will be passed
	s.Step(allure.NewSimpleStep("Step B").Failed())  // This step will be failed
	s.Step(allure.NewSimpleStep("Step C").Skipped()) // This step will be skipped

	stepD := allure.NewSimpleStep("Step D").WithStart()
	time.Sleep(1 * time.Second) // Do some
	stepD = stepD.WithStop().Passed()
	s.Step(stepD)
}

func (s *StepDemoSuite) TestInnerStep() {
	s.Epic("Demo")
	s.Feature("Steps")
	s.Title("Add child steps to existed step.")
	s.Description(`
		Step A is parent step for Step B and Step C
		Step D is parent step for Step E and Step F
		Call order will be saved in allure report
		A -> (B, C), D -> (E, F)`)

	s.Tags("Steps", "Nesting")

	// use allure.NewSimpleStep constructor
	stepA := allure.NewSimpleStep("Step A")
	s.Step(stepA)
	stepB := allure.NewSimpleInnerStep("Step B", stepA)
	s.Step(stepB)
	stepC := allure.NewSimpleInnerStep("Step C", stepA)
	s.Step(stepC)

	// use InnerStep function
	stepD := allure.NewSimpleStep("Step D")
	s.Step(stepD)
	s.InnerStep(stepD, allure.NewSimpleStep("Step E"))
	stepF := allure.NewStep("Step F", // Step's Name
		allure.Passed,                           // Step Status
		allure.GetNow(),                         // Step Start
		allure.GetNow(),                         // Step Finish
		allure.NewParameters("paramF", "value")) // Step Parameters
	s.InnerStep(stepD, stepF)

	// forward way
	stepG := allure.NewSimpleStep("Step G")
	stepH := allure.NewSimpleStep("Step H")
	stepI := allure.NewSimpleStep("Step I")
	stepH.Parent = stepG.GetUUID()
	stepI.Parent = stepG.GetUUID()
}

func TestStepDemo(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(StepDemoSuite))
}
