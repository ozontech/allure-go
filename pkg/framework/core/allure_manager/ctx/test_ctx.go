package ctx

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type testCtx struct {
	name   string
	result *allure.Result
}

func NewTestCtx(result *allure.Result) provider.ExecutionContext {
	return &testCtx{result: result, name: constants.TestContextName}
}

func (ctx *testCtx) AddStep(newStep *allure.Step) {
	ctx.result.Steps = append(ctx.result.Steps, newStep)
}

func (ctx *testCtx) GetName() string {
	return ctx.name
}

func (ctx *testCtx) AddAttachments(attachments ...*allure.Attachment) {
	ctx.result.Attachments = append(ctx.result.Attachments, attachments...)
}
