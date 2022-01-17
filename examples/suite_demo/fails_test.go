//go:build examples
// +build examples

package suite_demo

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type FailsDemoSuite struct {
	suite.Suite
}

func (s *FailsDemoSuite) BeforeEach() {
	s.Epic("Demo")
	s.Feature("Failures")
}

func (s *FailsDemoSuite) TestAssertionFail() {
	s.Title("This test failed by assert with message")
	s.Description(`
		This Test will be failed with assert Error.
		Error text: Assertion Failed`)
	s.Tags("fail", "assertions")

	t := s.T()
	require.Equal(t, 1, 2, "Assertion Failed")
}

func (s *FailsDemoSuite) TestXSkipFail() {
	s.Title("This test skipped by assert with message")
	s.Description(`
		This Test will be skipped with assert Error.
		Error text: Assertion Failed`)
	s.Tags("fail", "xskip", "assertions")

	t := s.T()
	t.XSkip()
	require.Equal(t, 1, 2, "Assertion Failed")
}

func (s *FailsDemoSuite) TestAssertionFailNoMessage() {
	s.Title("This test failed by assert without message")
	s.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2`)

	s.Tags("fail", "assertions")

	t := s.T()
	require.Equal(t, 1, 2)
}

func (s *FailsDemoSuite) TestAssertionFailInnerSteps() {
	s.Title("This test failed by assert with inner step")
	s.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2`)

	s.Tags("fail", "assertions", "nesting")

	t := s.T()

	s.WithNewStep("Failed parent step", func() {
		s.WithNewStep("Failed child step", func() {
			require.Equal(t, 1, 2, "Failed inside step")
		})
	})
}

func (s *FailsDemoSuite) TestPanic() {
	s.Title("This test panicked")
	s.Description(`
		This Test will Failed by panic.
		Error text:
	test panicked: runtime error: index out of range [0] with length 0 goroutine 8 [running]:...`)

	s.Tags("fail", "panic")

	t := s.T()

	var test []string
	test2 := test[0]
	require.Equal(t, test2, test2, "Never reach this")
}

func (s *FailsDemoSuite) TestPanicInnerSteps() {
	s.Title("This test panicked with inner steps")
	s.Description(`
		This Test will Failed by panic.
		All steps that includes error will be failed
		Error text:
	test panicked: runtime error: index out of range [0] with length 0 goroutine 8 [running]:...`)

	s.Tags("fail", "panic", "nesting")

	t := s.T()

	s.WithNewStep("Check 1", func() {
		s.WithNewStep("Check 1.1", func() {

		})
		s.WithNewStep("Check 1.2", func() {
			s.WithNewStep("Check 1.2.1", func() {
				var test []string
				test2 := test[0]
				require.Equal(t, test2, test2, "Never reach this")
			})
		})
	})
}

func TestFails(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(FailsDemoSuite))
}
