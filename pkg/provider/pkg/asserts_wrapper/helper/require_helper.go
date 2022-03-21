package helper

import (
	"github.com/ozontech/allure-go/pkg/provider/pkg/asserts_wrapper/wrapper"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

// NewRequireHelper inits new Require interface
func NewRequireHelper(t ProviderT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewRequire(),
	}
}

// NewRequireSubStepHelper inits new Require interface for sub step
func NewRequireSubStepHelper(t ProviderT, ctx provider.StepCtx) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewRequireSubStep(ctx),
	}
}
