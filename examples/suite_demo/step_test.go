//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type StepDemoSuite struct {
	suite.Suite
}

func (s *StepDemoSuite) TestAddSteps(t provider.T) {
	t.Epic("Demo")
	t.Feature("Steps")
	t.Title("Base. Add steps to allure result")
	t.Description(`
		Step A, Step B and Step C will be add to Allure Result`)

	t.Tags("Steps")

	stepA := allure.NewSimpleStep("Step A")
	t.Step(stepA)

	stepB := allure.NewStep("Step B", // Step's Name
		allure.Passed,                           // Step Status
		allure.GetNow(),                         // Step Start
		allure.GetNow(),                         // Step Finish
		allure.NewParameters("paramB", "value")) // Step Parameters

	t.Step(stepB)

	stepC := allure.NewSimpleStep("Step C")
	stepC.Start = allure.GetNow()
	stepC.WithNewParameters("paramC", "value")
	stepC.Stop = allure.GetNow()
	t.Step(stepC)
}

func (s *StepDemoSuite) TestQuickWorkWithSteps(t provider.T) {
	t.Epic("Demo")
	t.Feature("Steps")
	t.Title("Base. Add steps to allure result")
	t.Description(`
		Step A, Step B, Step C and Step D will be add to Allure Result`)

	t.Tags("Steps")

	t.Step(allure.NewSimpleStep("Step A").Passed())  // This step will be passed
	t.Step(allure.NewSimpleStep("Step B").Failed())  // This step will be failed
	t.Step(allure.NewSimpleStep("Step C").Skipped()) // This step will be skipped

	stepD := allure.NewSimpleStep("Step D").Begin()
	time.Sleep(1 * time.Second) // Do some
	stepD = stepD.Finish().Passed()
	t.Step(stepD)
}

func (s *StepDemoSuite) TestInnerStep(t provider.T) {
	t.Epic("Demo")
	t.Feature("Steps")
	t.Title("Add child steps to existed step.")
	t.Description(`
		Step A is parent step for Step B and Step C
		Step D is parent step for Step E and Step F
		Call order will be saved in allure report
		A -> (B, C), D -> (E, F)`)

	t.Tags("Steps", "Nesting")

	// use allure.NewSimpleStep constructor
	stepA := allure.NewSimpleStep("Step A")
	stepB := allure.NewSimpleStep("Step B")
	stepC := allure.NewSimpleStep("Step C")
	stepA.WithChild(stepB)
	stepA.WithChild(stepC)
	t.Step(stepA)

	// use InnerStep function
	stepD := allure.NewSimpleStep("Step D")
	t.Step(stepD)
	stepD.WithChild(allure.NewSimpleStep("Step E"))
	stepF := allure.NewStep("Step F", // Step's Name
		allure.Passed,                           // Step Status
		allure.GetNow(),                         // Step Start
		allure.GetNow(),                         // Step Finish
		allure.NewParameters("paramF", "value")) // Step Parameters
	stepF.WithParent(stepD)

	// forward way
	stepG := allure.NewSimpleStep("Step G")
	stepH := allure.NewSimpleStep("Step H")
	stepI := allure.NewSimpleStep("Step I")
	stepG.WithChild(stepH)
	stepG.WithChild(stepI)
}

func TestStepDemo(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(StepDemoSuite))
}
