package common

import (
	"fmt"
	"runtime/debug"
	"sync"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/helper"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type StepProvider interface {
	StopResult(status allure.Status)
	UpdateResultStatus(msg string, trace string)
	ExecutionContext() provider.ExecutionContext
}

type StepT interface {
	FailNow()
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Log(args ...interface{})
	Logf(format string, args ...interface{})
	Break(args ...interface{})
	Breakf(format string, args ...interface{})
	Broken()
	BrokenNow()
	Name() string
	GetRealT() provider.TestingT
}

type InternalStepCtx interface {
	provider.StepCtx

	ExecutionContextName() string
	WG() *sync.WaitGroup
}

type stepCtx struct {
	t StepT
	p StepProvider

	currentStep *allure.Step
	parentStep  InternalStepCtx

	asserts provider.Asserts
	require provider.Asserts

	wg sync.WaitGroup
}

func NewStepCtx(t StepT, p StepProvider, stepName string, params ...*allure.Parameter) InternalStepCtx {
	currentStep := allure.NewSimpleStep(stepName, params...)
	newCtx := &stepCtx{t: t, p: p, currentStep: currentStep, wg: sync.WaitGroup{}}
	newCtx.asserts = helper.NewAssertsHelper(newCtx)
	newCtx.require = helper.NewRequireHelper(newCtx)
	return newCtx
}

func (ctx *stepCtx) NewChildCtx(stepName string, params ...*allure.Parameter) InternalStepCtx {
	currentStep := allure.NewSimpleStep(stepName, params...)
	newCtx := &stepCtx{t: ctx.t, p: ctx.p, currentStep: currentStep, parentStep: ctx, wg: sync.WaitGroup{}}
	newCtx.asserts = helper.NewAssertsHelper(newCtx)
	newCtx.require = helper.NewRequireHelper(newCtx)
	return newCtx
}

func (ctx *stepCtx) Name() string {
	return ctx.t.Name()
}

func (ctx *stepCtx) Assert() provider.Asserts {
	return ctx.asserts
}

func (ctx *stepCtx) Require() provider.Asserts {
	return ctx.require
}

func (ctx *stepCtx) WG() *sync.WaitGroup {
	return &ctx.wg
}

func (ctx *stepCtx) ExecutionContextName() string {
	return ctx.p.ExecutionContext().GetName()
}

func (ctx *stepCtx) FailNow() {
	ctx.Fail()
	ctx.t.FailNow()
}

func (ctx *stepCtx) Error(args ...interface{}) {
	ctx.t.GetRealT().Helper()

	ctx.Fail()
	ctx.t.Error(args...)
}

func (ctx *stepCtx) Errorf(format string, args ...interface{}) {
	ctx.t.GetRealT().Helper()

	ctx.Fail()
	ctx.t.Errorf(format, args...)
}

func (ctx *stepCtx) Log(args ...interface{}) {
	ctx.t.GetRealT().Helper()

	ctx.t.Log(args...)
}

func (ctx *stepCtx) Logf(format string, args ...interface{}) {
	ctx.t.GetRealT().Helper()

	ctx.t.Logf(format, args...)
}

func (ctx *stepCtx) CurrentStep() *allure.Step {
	return ctx.currentStep
}

func (ctx *stepCtx) WithParameters(parameters ...*allure.Parameter) {
	ctx.currentStep.WithParameters(parameters...)
}

func (ctx *stepCtx) WithNewParameters(kv ...interface{}) {
	ctx.currentStep.WithNewParameters(kv...)
}

func (ctx *stepCtx) WithAttachments(attachments ...*allure.Attachment) {
	ctx.currentStep.WithAttachments(attachments...)
}

func (ctx *stepCtx) WithNewAttachment(name string, mimeType allure.MimeType, content []byte) {
	ctx.currentStep.WithAttachments(allure.NewAttachment(name, mimeType, content))
}

func (ctx *stepCtx) LogStep(args ...interface{}) {
	newStep := allure.NewSimpleStep(fmt.Sprintln(args...))
	ctx.currentStep.WithChild(newStep)
	ctx.Log(args...)
}

func (ctx *stepCtx) LogfStep(format string, args ...interface{}) {
	newStep := allure.NewSimpleStep(fmt.Sprintf(format, args...))
	ctx.currentStep.WithChild(newStep)
	ctx.Logf(format, args...)
}

func (ctx *stepCtx) Step(step *allure.Step) {
	ctx.currentStep.WithChild(step)
}

func (ctx *stepCtx) NewStep(stepName string, parameters ...*allure.Parameter) {
	newStep := allure.NewSimpleStep(stepName, parameters...)
	ctx.currentStep.WithChild(newStep)
}

func (ctx *stepCtx) WithNewStep(stepName string, step func(ctx provider.StepCtx), params ...*allure.Parameter) {
	newCtx := ctx.NewChildCtx(stepName, params...)
	defer ctx.currentStep.WithChild(newCtx.CurrentStep())
	defer func() {
		r := recover()
		newCtx.WG().Wait()
		newCtx.CurrentStep().Finish()
		if r != nil {
			ctxName := newCtx.ExecutionContextName()
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, r, debug.Stack())
			newCtx.Broken()
			TestError(ctx.t, ctx.p, ctxName, errMsg)
		}
	}()
	step(newCtx)
}

func (ctx *stepCtx) WithNewAsyncStep(stepName string, step func(ctx provider.StepCtx), params ...*allure.Parameter) {
	var wg *sync.WaitGroup
	wg = &ctx.wg
	if ctx.parentStep != nil {
		wg = ctx.parentStep.WG()
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		ctx.WithNewStep(stepName, step, params...)
	}()
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
	ctx.t.Broken()
}

func (ctx *stepCtx) BrokenNow() {
	ctx.currentStep.Broken()
	if ctx.parentStep != nil {
		ctx.parentStep.Broken()
	}
	ctx.t.BrokenNow()
}

func (ctx *stepCtx) Break(args ...interface{}) {
	ctx.Broken()
	ctx.t.Break(args...)
}

func (ctx *stepCtx) Breakf(format string, args ...interface{}) {
	ctx.Broken()
	ctx.t.Breakf(format, args...)
}

func (ctx *stepCtx) GetRealT() provider.TestingT {
	return ctx.t.GetRealT()
}
