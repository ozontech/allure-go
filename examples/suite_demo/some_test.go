package suite_demo

import (
	"context"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type Suite struct {
	suite.Suite
}

func (s *Suite) TestTearDown(t provider.T) {
	ctx := context.Background()
	t.Parallel()
	t.WithTestSetup(func(t provider.T) {
		t.NewStep("Some before step", allure.NewParameters("p1", "v1")...)
		t.Logf("Complete before")
	})

	defer t.WithTestTeardown(func(t provider.T) {
		t.WithNewStep("Some after step", func(sCtx provider.StepCtx) {
			ctx.Done()
			sCtx.WithNewParameters("ctx", ctx)
		})
		t.Logf("Complete after")
	})
	t.Require().NotNil(nil)
}

func TestRunner2(t *testing.T) {
	suite.RunSuite(t, new(Suite))
}
