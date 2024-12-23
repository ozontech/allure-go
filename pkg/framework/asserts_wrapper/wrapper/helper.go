package wrapper

import (
	"fmt"

	"github.com/ozontech/allure-go/pkg/allure"
)

type Provider interface {
	Step(step *allure.Step)
}

type assertHelper struct {
	prefix string
}

func (h *assertHelper) getStepName(assertName string, msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 {
		return fmt.Sprintf("%s: %s", h.prefix, assertName)
	}
	return fmt.Sprintf("%s: %s", h.prefix, messageFromMsgAndArgs(msgAndArgs...))
}

func (h *assertHelper) WithNewStep(t TestingT, provider Provider, assertName string, assert func(t TestingT) bool, params []*allure.Parameter, msgAndArgs ...interface{}) bool {
	var (
		step   = allure.NewSimpleStep(h.getStepName(assertName, msgAndArgs...), params...)
		result = assert(t)
	)

	provider.Step(step)
	if !result {
		step.Failed()
	}

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
