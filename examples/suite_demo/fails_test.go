//go:build examples_new
// +build examples_new

package suite_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/provider/pkg/framework/suite"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type FailsDemoSuite struct {
	suite.Suite
}

func (s *FailsDemoSuite) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("Failures")
}

func (s *FailsDemoSuite) TestAssertionFail(t provider.T) {
	t.Title("This test failed by assert with message")
	t.Description(`
		This Test will be failed with assert Error.
		Error text: Assertion Failed`)
	t.Tags("fail", "assertions")

	t.Require().Equal(1, 2, "Assertion Failed")
}

func (s *FailsDemoSuite) TestXSkipFail(t provider.T) {
	t.Title("This test skipped by assert with message")
	t.Description(`
		This Test will be skipped with assert Error.
		Error text: Assertion Failed`)
	t.Tags("fail", "xskip", "assertions")

	t.XSkip()
	t.Require().Equal(1, 2, "Assertion Failed")
}

func (s *FailsDemoSuite) TestAssertionFailNoMessage(t provider.T) {
	t.Title("This test failed by assert without message")
	t.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2`)

	t.Tags("fail", "assertions")

	t.Require().Equal(1, 2)
}

func (s *FailsDemoSuite) TestAssertionFailInnerSteps(t provider.T) {
	t.Title("This test failed by assert with inner step")
	t.Description(`
		This Test will be failed with assert Error.
		Error text:
					Not equal:
					expected: 1
					actual  : 2`)

	t.Tags("fail", "assertions", "nesting")

	t.WithNewStep("Failed parent step", func(ctx provider.StepCtx) {
		ctx.WithNewStep("Failed child step", func(ctx provider.StepCtx) {
			ctx.Require().Equal(1, 2, "Failed inside step")
		})
	})
}

func (s *FailsDemoSuite) TestPanic(t provider.T) {
	t.Title("This test panicked")
	t.Description(`
		This Test will Failed by panic.
		Error text:
	test panicked: runtime error: index out of range [0] with length 0 goroutine 8 [running]:...`)

	t.Tags("fail", "panic")

	var test []string
	test2 := test[0]
	t.Require().Equal(test2, test2, "Never reach this")
}

func (s *FailsDemoSuite) TestPanicInnerSteps(t provider.T) {
	t.Title("This test panicked with inner steps")
	t.Description(`
		This Test will Failed by panic.
		All steps that includes error will be failed
		Error text:
	test panicked: runtime error: index out of range [0] with length 0 goroutine 8 [running]:...`)

	t.Tags("fail", "panic", "nesting")

	t.WithNewStep("Check 1", func(ctx provider.StepCtx) {
		ctx.WithNewStep("Check 1.1", func(ctx provider.StepCtx) {

		})
		ctx.WithNewStep("Check 1.2", func(ctx provider.StepCtx) {
			ctx.WithNewStep("Check 1.2.1", func(ctx provider.StepCtx) {
				var test []string
				test2 := test[0]
				ctx.Require().Equal(test2, test2, "Never reach this")
			})
		})
	})
}

func TestFails(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(FailsDemoSuite))
}
