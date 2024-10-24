package examples

import (
	"context"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
	"testing"
)

type SetupSuite struct {
	suite.Suite
}

func (s *SetupSuite) Test_Auth_RegisterEmployee(t provider.T) {
	t.Title("[Auth] RegisterEmployee")
	t.Descriptionf("Register employee tests")
	t.Tags("auth")
	t.Parallel()

	t.WithNewAsyncStep("Incorrect company ID", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		sCtx.WithNewParameters("ctx", ctx)
	})
	t.WithNewAsyncStep("Correct request", func(sCtx provider.StepCtx) {
		ctx := context.TODO()
		sCtx.WithNewParameters("ctx", ctx)
	})

}

func TestRunner(t *testing.T) {
	suite.RunSuite(t, new(SetupSuite))
}
