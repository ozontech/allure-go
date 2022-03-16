package wrapper

import (
	"fmt"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type assertHelper struct {
	parentCtx provider.StepCtx
	required  bool
}

func (h *assertHelper) getStepName(assertName string, msgAndArgs ...string) string {
	prefix := "ASSERT"
	if h.required {
		prefix = "REQUIRE"
	}
	return fmt.Sprintf("%s: %s", prefix, assertName)
}

func (h *assertHelper) withNewStep(t ProviderT, stepName string, f func(ctx provider.StepCtx), params ...allure.Parameter) {
	if h.parentCtx != nil {
		h.parentCtx.WithNewStep(stepName, f, params...)
		return
	}
	t.WithNewStep(stepName, f, params...)
}

func (h *assertHelper) handleResult(ctx provider.StepCtx, result bool) {
	if !result {
		ctx.Fail()
		if h.required {
			ctx.T().FailNow()
		}
	}
}
