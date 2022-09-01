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

func (h *assertHelper) getStepName(assertName string, msgAndArgs ...interface{}) string {
	prefix := "ASSERT"
	if h.required {
		prefix = "REQUIRE"
	}
	if len(msgAndArgs) == 0 {
		return fmt.Sprintf("%s: %s", prefix, assertName)
	}
	return fmt.Sprintf("%s: %s", prefix, messageFromMsgAndArgs(msgAndArgs...))
}

func (h *assertHelper) withNewStep(t TestingT, provider Provider, assertName string, assert func(t TestingT) bool, params []*allure.Parameter, msgAndArgs ...interface{}) bool {
	var result bool
	step := allure.NewSimpleStep(h.getStepName(assertName, msgAndArgs...), params...)
	defer func() {
		if !result {
			step.Failed()
		}
		provider.Step(step)
	}()
	result = assert(t)

	return result
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
