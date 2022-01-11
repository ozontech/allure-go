//go:build allure_go
// +build allure_go

package allure_go_compare

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
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

func (s *SuiteStruct) TestNewTest() {
	s.Epic("Compare with allure-go")
	s.Description("New Test Description")
	s.WithNewStep("Step description", func() {

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

func doSomething(s *SuiteStruct) {
	s.WithNewStep("Something", func() {
		s.NewStep("Because We Can!")
	})
}

// Works even better! each step will have his real status, but parent will be failed anyway

func (s *SuiteStruct) TestWithIntricateSubsteps() {
	s.Epic("Compare with allure-go")
	s.Description("Test")

	t := s.T()

	s.WithNewStep("Step 1", func() {
		doSomething(s)
		s.WithNewStep("Sub-step 1.1", func() {
			t.Errorf("Failure")
		})
		s.WithNewStep("Sub-step 1.2", func() {

		})
		s.WithStep(allure.NewSimpleStep("Sub-step 1.3").Skipped(), func() {

		})
	})
	s.WithNewStep("Step 2", func() {
		s.WithNewStep("Sub-Step 2.1", func() {
			s.WithNewStep("Sub-step 2.1.1", func() {
				s.WithNewStep("Sub-step 2.1.1.1", func() {
					t.Errorf("Failure")
				})
				s.WithNewStep("Sub-step 2.1.1.2", func() {
					t.Errorf("Failed like this")
				})
			})
		})
		s.WithNewStep("Sub-Step 2.2", func() {

		})
	})
}

func TestRun(t *testing.T) {
	runner.RunSuite(t, new(SuiteStruct))
}
