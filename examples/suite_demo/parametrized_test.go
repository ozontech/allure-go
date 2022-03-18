//go:build examples_new
// +build examples_new

package suite_demo

import (
	"fmt"
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type ParametrizedTestDemo struct {
	suite.Suite
}

func (s *ParametrizedTestDemo) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Parametrized")
}

func (s *ParametrizedTestDemo) TestParameterized(t provider.T) {
	t.Title("Parent Test")
	t.Description(`
		This test is parent for all Parametrized`)

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized #%d", newI), func(t provider.T) {
			t.Feature("Parametrized")
			t.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			t.Tag("Parametrized")
			t.Parallel()
			t.WithNewStep(fmt.Sprintf("Step %d", i), func(ctx provider.StepCtx) {
				ctx.Require().Equal(1, newI)
			})
		})
	}

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized 2#%d", newI), func(t provider.T) {
			t.Feature("Parametrized")
			t.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			t.Tag("Parametrized")
			t.Parallel()
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func(ctx provider.StepCtx) {
				ctx.Require().Less(1, newI)
			})
		})
	}
}

func TestParametrizedDemo(t *testing.T) {
	suite.RunSuite(t, new(ParametrizedTestDemo))
}

type ParametrizedTestDemo2 struct {
	suite.Suite
}

func (s *ParametrizedTestDemo2) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Parametrized")
}

func (s *ParametrizedTestDemo2) TestParameterized2(t provider.T) {
	t.Title("Parent Test")
	t.Description(`
		This test is parent for all Parametrized`)

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized #%d", newI), func(t provider.T) {
			t.Feature("Parametrized")
			t.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			t.Tag("Parametrized")
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func(ctx provider.StepCtx) {
				ctx.Require().Equal(1, newI)
			})
		})
	}

	for i := 0; i < 10; i++ {
		newI := i
		t.Run(fmt.Sprintf("Parametrized 2#%d", newI), func(t provider.T) {
			t.Feature("Parametrized")
			t.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			t.Tag("Parametrized")
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func(ctx provider.StepCtx) {
				if newI == 4 {
					panic("WHOOPS")
				}
				ctx.Require().Less(1, newI)
			})
		})
	}
}

func TestParametrizedDemo2(t *testing.T) {
	suite.RunSuite(t, new(ParametrizedTestDemo2))
}
