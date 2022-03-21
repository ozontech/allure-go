package steps

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/internal"
	"github.com/ozontech/allure-go/pkg/provider/pkg/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type stepCtx struct {
	t           ProviderT
	currentStep *allure.Step
	parentStep  *stepCtx

	asserts provider.Asserts
	require provider.Asserts

	wg sync.WaitGroup
}

func NewStepCtx(t ProviderT, stepName string, params ...allure.Parameter) provider.StepCtx {
	currentStep := allure.NewSimpleStep(stepName, params...)
	newCtx := &stepCtx{t: t, currentStep: currentStep, wg: sync.WaitGroup{}}
	newCtx.asserts = helper.NewAssertsSubStepHelper(t, newCtx)
	newCtx.require = helper.NewRequireSubStepHelper(t, newCtx)
	return newCtx
}

func (ctx *stepCtx) newChildCtx(stepName string, params ...allure.Parameter) *stepCtx {
	currentStep := allure.NewSimpleStep(stepName, params...)
	newCtx := &stepCtx{t: ctx.t, currentStep: currentStep, parentStep: ctx, wg: sync.WaitGroup{}}
	newCtx.asserts = helper.NewAssertsSubStepHelper(ctx.t, newCtx)
	newCtx.require = helper.NewRequireSubStepHelper(ctx.t, newCtx)
	return newCtx
}

func (ctx *stepCtx) T() provider.T {
	return ctx.t.(provider.T)
}

func (ctx *stepCtx) Error(args ...interface{}) {
	ctx.Fail()
	ctx.T().Error(args...)
}

func (ctx *stepCtx) Errorf(format string, args ...interface{}) {
	ctx.Fail()
	ctx.T().Errorf(format, args...)
}
func (ctx *stepCtx) CurrentStep() *allure.Step {
	return ctx.currentStep
}

func (ctx *stepCtx) Log(args ...interface{}) {
	ctx.t.Log(args...)
}

func (ctx *stepCtx) Logf(format string, args ...interface{}) {
	ctx.t.Logf(format, args...)
}

func (ctx *stepCtx) Assert() provider.Asserts {
	return ctx.asserts
}

func (ctx *stepCtx) Require() provider.Asserts {
	return ctx.require
}

func (ctx *stepCtx) NewStep(stepName string, parameters ...allure.Parameter) {
	newStep := allure.NewSimpleStep(stepName, parameters...)
	ctx.currentStep.WithChild(newStep)
}

func (ctx *stepCtx) WithNewStep(stepName string, step func(ctx provider.StepCtx), params ...allure.Parameter) {
	newCtx := ctx.newChildCtx(stepName, params...)
	defer ctx.currentStep.WithChild(newCtx.currentStep)
	defer func() {
		r := recover()
		if r != nil {
			ctxName := newCtx.t.Provider().ExecutionContext().GetName()
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, r, debug.Stack())
			newCtx.Broken()
			internal.TestError(newCtx.t.Provider().ExecutionContext().GetName(), errMsg, newCtx.t)
		}
	}()
	step(newCtx)
}

func (ctx *stepCtx) WithNewAsyncStep(stepName string, step func(ctx provider.StepCtx), params ...allure.Parameter) {
	var wg *sync.WaitGroup
	wg = &ctx.wg
	if ctx.parentStep != nil {
		wg = &ctx.parentStep.wg
		defer wg.Wait()
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		ctx.WithNewStep(stepName, step, params...)
	}()
}

func (ctx *stepCtx) WithParameters(parameters ...allure.Parameter) {
	ctx.currentStep.WithParameters(parameters...)
}

func (ctx *stepCtx) WithNewParameters(kv ...string) {
	ctx.currentStep.WithNewParameters(kv...)
}

func (ctx *stepCtx) WithAttachments(attachments ...*allure.Attachment) {
	ctx.currentStep.WithAttachments(attachments...)
}

func (ctx *stepCtx) WithNewAttachment(name string, mimeType allure.MimeType, content []byte) {
	ctx.currentStep.WithAttachments(allure.NewAttachment(name, mimeType, content))
}

func (ctx *stepCtx) Fail() {
	ctx.currentStep.Failed()
	if ctx.parentStep != nil {
		ctx.parentStep.Fail()
	}
}

func (ctx *stepCtx) Broken() {
	ctx.currentStep.Broken()
	if ctx.parentStep != nil {
		ctx.parentStep.Broken()
	}
}
