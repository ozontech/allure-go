//go:build examples_new
// +build examples_new

package suite_demo

import (
	"context"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type SetupSuite struct {
	suite.Suite
}

func (s *SetupSuite) TestMyTest(t provider.T) {
	var (
		v1  string
		v2  int
		ctx context.Context
	)
	t.WithTestSetup(func(t provider.T) {
		t.WithNewStep("init v1", func(sCtx provider.StepCtx) {
			v1 = "string"
			sCtx.WithNewParameters("v1", v1)
		})
		t.WithNewStep("init v2", func(sCtx provider.StepCtx) {
			v2 = 123
			sCtx.WithNewParameters("v2", v2)
		})
		t.WithNewStep("init ctx", func(sCtx provider.StepCtx) {
			ctx = context.Background()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})
	defer t.WithTestTeardown(func(t provider.T) {
		t.WithNewStep("Close ctx", func(sCtx provider.StepCtx) {
			ctx.Done()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})
}

func TestRunner(t *testing.T) {
	suite.RunSuite(t, new(SetupSuite))
}
