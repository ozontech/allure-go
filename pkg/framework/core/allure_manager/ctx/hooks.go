package ctx

import (
	"fmt"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type hooksCtx struct {
	name      string
	container *allure.Container
}

// NewAfterAllCtx returns after all context
func NewAfterAllCtx(container *allure.Container) provider.ExecutionContext {
	return &hooksCtx{container: container, name: constants.AfterAllContextName}
}

// NewAfterEachCtx returns after each context
func NewAfterEachCtx(container *allure.Container) provider.ExecutionContext {
	return &hooksCtx{container: container, name: constants.AfterEachContextName}
}

// NewBeforeAllCtx returns before all context
func NewBeforeAllCtx(container *allure.Container) provider.ExecutionContext {
	return &hooksCtx{container: container, name: constants.BeforeAllContextName}
}

// NewBeforeEachCtx returns before each context
func NewBeforeEachCtx(result *allure.Container) provider.ExecutionContext {
	return &hooksCtx{container: result, name: constants.BeforeEachContextName}
}

// AddStep adds step to current execution container
func (ctx *hooksCtx) AddStep(newStep *allure.Step) {
	switch ctx.name {
	case constants.BeforeAllContextName, constants.BeforeEachContextName:
		ctx.container.Befores = append(ctx.container.Befores, newStep)

	case constants.AfterAllContextName, constants.AfterEachContextName:
		ctx.container.Afters = append(ctx.container.Afters, newStep)
	}
}

// GetName returns context name
func (ctx *hooksCtx) GetName() string {
	return ctx.name
}

// AddAttachments adds attachment to the execution context
func (ctx *hooksCtx) AddAttachments(attachments ...*allure.Attachment) {
	if len(attachments) == 0 {
		return
	}

	first := attachments[0]

	newStep := allure.NewSimpleStep(
		fmt.Sprintf("Attachment %s", first.Name),
		allure.NewParameter("name", first.Name),
		allure.NewParameter("type", string(first.Type)),
		allure.NewParameter("source", first.Source),
	).WithAttachments(attachments...)

	ctx.AddStep(newStep)
}
