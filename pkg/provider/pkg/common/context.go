package common

import (
	"fmt"
	"runtime/debug"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/internal"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type testCtx struct {
	name   string
	result *allure.Result
}

func newTestCtx(result *allure.Result) provider.ExecutionContext {
	return &testCtx{result: result, name: internal.TestContextName}
}

func (ctx *testCtx) AddStep(newStep *allure.Step) {
	ctx.result.Steps = append(ctx.result.Steps, newStep)
}

func (ctx *testCtx) GetName() string {
	return ctx.name
}

func (ctx *testCtx) AddAttachment(attachment *allure.Attachment) {
	ctx.result.Attachments = append(ctx.result.Attachments, attachment)
}

type beforeEachCtx struct {
	name      string
	container *allure.Container
}

func newBeforeEachCtx(result *allure.Container) provider.ExecutionContext {
	return &beforeEachCtx{container: result, name: internal.BeforeEachContextName}
}

func (ctx *beforeEachCtx) AddStep(newStep *allure.Step) {
	ctx.container.Befores = append(ctx.container.Befores, newStep)
}

func (ctx *beforeEachCtx) GetName() string {
	return ctx.name
}

func (ctx *beforeEachCtx) AddAttachment(attachment *allure.Attachment) {
	newStep := allure.NewSimpleStep(
		fmt.Sprintf("Attachment %s", attachment.Name),
		allure.NewParameter("name", attachment.Name),
		allure.NewParameter("type", string(attachment.Type)),
		allure.NewParameter("source", attachment.Source))
	newStep.WithAttachments(attachment)
	ctx.AddStep(newStep)
}

type afterEachCtx struct {
	name      string
	container *allure.Container
}

func newAfterEachCtx(container *allure.Container) provider.ExecutionContext {
	return &afterEachCtx{container: container, name: internal.AfterEachContextName}
}

func (ctx *afterEachCtx) AddStep(newStep *allure.Step) {
	ctx.container.Afters = append(ctx.container.Afters, newStep)
}

func (ctx *afterEachCtx) GetName() string {
	return ctx.name
}

func (ctx *afterEachCtx) AddAttachment(attachment *allure.Attachment) {
	newStep := allure.NewSimpleStep(
		fmt.Sprintf("Attachment %s", attachment.Name),
		allure.NewParameter("name", attachment.Name),
		allure.NewParameter("type", string(attachment.Type)),
		allure.NewParameter("source", attachment.Source))
	newStep.WithAttachments(attachment)
	ctx.AddStep(newStep)
}

type beforeAllCtx struct {
	name      string
	container *allure.Container
}

func newBeforeAllCtx(container *allure.Container) provider.ExecutionContext {
	return &beforeAllCtx{container: container, name: internal.BeforeAllContextName}
}

func (ctx *beforeAllCtx) AddStep(newStep *allure.Step) {
	ctx.container.Befores = append(ctx.container.Befores, newStep)
}

func (ctx *beforeAllCtx) GetName() string {
	return ctx.name
}

func (ctx *beforeAllCtx) AddAttachment(attachment *allure.Attachment) {
	newStep := allure.NewSimpleStep(
		fmt.Sprintf("Attachment %s", attachment.Name),
		allure.NewParameter("name", attachment.Name),
		allure.NewParameter("type", string(attachment.Type)),
		allure.NewParameter("source", attachment.Source))
	newStep.WithAttachments(attachment)
	ctx.AddStep(newStep)
}

type afterAllCtx struct {
	name      string
	container *allure.Container
}

func newAfterAllCtx(container *allure.Container) provider.ExecutionContext {
	return &afterAllCtx{container: container, name: internal.AfterAllContextName}
}

func (ctx *afterAllCtx) AddStep(newStep *allure.Step) {
	ctx.container.Afters = append(ctx.container.Afters, newStep)
}

func (ctx *afterAllCtx) GetName() string {
	return ctx.name
}

func (ctx *afterAllCtx) AddAttachment(attachment *allure.Attachment) {
	newStep := allure.NewSimpleStep(
		fmt.Sprintf("Attachment %s", attachment.Name),
		allure.NewParameter("name", attachment.Name),
		allure.NewParameter("type", string(attachment.Type)),
		allure.NewParameter("source", attachment.Source))
	newStep.WithAttachments(attachment)
	ctx.AddStep(newStep)
}

func BeforeAllHook(t provider.InternalT, provider Provider) {
	t.WG().Add(1)
	defer t.WG().Done()
	if provider.GetSuiteMeta().GetBeforeAll() != nil {
		provider.BeforeAllContext()
		defer func() {
			r := recover()
			if r != nil {
				t.Errorf("BeforeAll hook panicked:%v\n%s", r, debug.Stack())
				t.FailNow()
			}
		}()
		provider.GetSuiteMeta().GetBeforeAll()(t)
	}
}

func AfterAllHook(t provider.InternalT, provider Provider) {
	t.WG().Add(1)
	defer t.WG().Done()
	if provider.GetSuiteMeta().GetAfterAll() != nil {
		provider.AfterAllContext()
		defer func() {
			r := recover()
			if r != nil {
				t.Errorf("AfterAll hook panicked:%v\n%s", r, debug.Stack())
				t.FailNow()
			}
		}()
		provider.GetSuiteMeta().GetAfterAll()(t)
	}
}

func BeforeEachHook(t provider.InternalT, provider Provider) {
	if provider.GetTestMeta().GetBeforeEach() != nil {
		t.Provider().BeforeEachContext()
		provider.GetTestMeta().GetBeforeEach()(t)
	}
}

func AfterEachHook(t provider.InternalT, provider Provider) {
	if provider.GetTestMeta().GetAfterEach() != nil {
		t.Provider().AfterEachContext()
		provider.GetTestMeta().GetAfterEach()(t)
	}
}
