package common

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func TestError(t testing.TB, provider provider.Provider, errMsg string) {
	short := errMsg
	if len(errMsg) > 100 {
		short = errMsg[:100]
	}
	switch provider.ExecutionContext().GetName() {
	case constants.TestContextName, constants.BeforeEachContextName:
		provider.StopResult(allure.Broken)
		provider.UpdateResultStatus(short, errMsg)
		t.Errorf(errMsg)
		t.FailNow()
	case constants.AfterEachContextName, constants.AfterAllContextName:
		t.Logf(errMsg)
		provider.UpdateResultStatus(short, errMsg)
	case constants.BeforeAllContextName:
		t.Logf(errMsg)
		provider.UpdateResultStatus(short, errMsg)
		t.FailNow()
	}
}
