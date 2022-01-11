//go:build examples
// +build examples

package suite_demo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
	"github.com/koodeex/allure-testify/pkg/provider"
)

type ParametrizedTestDemo struct {
	suite.Suite
}

func (s *ParametrizedTestDemo) BeforeEach() {
	s.Epic("Demo")
	s.Feature("Paramterized")
}

func (s *ParametrizedTestDemo) TestParameterized() {
	s.Title("Parent Test")
	s.Description(`
		This test is parent for all Parametrized`)

	for i := 0; i < 10; i++ {
		newI := i
		s.RunTest(fmt.Sprintf("Parametrized #%d", newI), func(t *provider.T) {
			t.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			t.Tag("Parameterized")
			t.Parallel()
			t.WithNewStep(fmt.Sprintf("Step %d", i), func() {
				require.Equal(t, 1, newI)
			})
		})
	}

	for i := 0; i < 10; i++ {
		newI := i
		s.RunTest(fmt.Sprintf("Parametrized 2#%d", newI), func(t *provider.T) {
			t.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			t.Tag("Parameterized")
			t.Parallel()
			t.WithNewStep(fmt.Sprintf("Step %d", newI), func() {
				require.Less(t, 1, newI)
			})
		})
	}
}

func TestParametrizedDemo(t *testing.T) {
	runner.RunSuite(t, new(ParametrizedTestDemo))
}

type ParametrizedTestDemo2 struct {
	suite.Suite
}

func (s *ParametrizedTestDemo2) BeforeEach() {
	s.Epic("Demo")
	s.Feature("Paramterized")
}

func (s *ParametrizedTestDemo2) TestParameterized2() {
	s.Title("Parent Test")
	s.Description(`
		This test is parent for all Parametrized`)

	for i := 0; i < 10; i++ {
		newI := i
		s.Run(fmt.Sprintf("Parametrized #%d", newI), func() {
			s.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			s.Tag("Parameterized")
			s.WithNewStep(fmt.Sprintf("Step %d", newI), func() {
				require.Equal(s.T(), 1, newI)
			})
		})
	}

	for i := 0; i < 10; i++ {
		newI := i
		s.Run(fmt.Sprintf("Parametrized 2#%d", newI), func() {
			s.Description(fmt.Sprintf("This test checks that 1 Equal %d", newI))
			s.Tag("Parameterized")
			s.WithNewStep(fmt.Sprintf("Step %d", newI), func() {
				if newI == 4 {
					panic("WHOOPS")
				}
				require.Less(s.T(), 1, newI)
			})
		})
	}
}

func TestParametrizedDemo2(t *testing.T) {
	runner.RunSuite(t, new(ParametrizedTestDemo2))
}
