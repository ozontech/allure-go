package wrapper

import (
	"fmt"

	"github.com/ozontech/allure-go/pkg/allure"
)

type Provider interface {
	Step(step *allure.Step)
}

type assertHelper struct {
	required bool
}

func (h *assertHelper) getStepName(assertName string) string {
	prefix := "ASSERT"
	if h.required {
		prefix = "REQUIRE"
	}
	return fmt.Sprintf("%s: %s", prefix, assertName)
}

func (h *assertHelper) withNewStep(t TestingT, provider Provider, assertName string, assert func(t TestingT) bool, params ...allure.Parameter) bool {
	step := allure.NewSimpleStep(h.getStepName(assertName), params...)
	result := assert(t)
	if !result {
		step.Failed()
	}
	provider.Step(step)
	return result
}
