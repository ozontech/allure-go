package common

import (
	"fmt"
	"runtime/debug"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/internal"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
	"github.com/ozontech/allure-go/pkg/provider/pkg/steps"
)

/*
Steps
*/

// Step adds step to test result
func (c *common) Step(step *allure.Step) {
	c.Provider().ExecutionContext().AddStep(step)
}

// NewStep creates new step and adds it to test result
func (c *common) NewStep(stepName string, params ...allure.Parameter) {
	c.Provider().ExecutionContext().AddStep(allure.NewSimpleStep(stepName, params...))
}

// WithNewStep opens nesting for struct.Step
// Any other struct.Step that will be added to struct.AllureResult object will be added as child step
func (c *common) WithNewStep(stepName string, step func(ctx provider.StepCtx), params ...allure.Parameter) {
	stCtx := steps.NewStepCtx(c, stepName, params...)
	defer c.Step(stCtx.CurrentStep())
	defer func() {
		r := recover()
		if r != nil {
			ctxName := c.Provider().ExecutionContext().GetName()
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, r, debug.Stack())
			stCtx.Broken()
			internal.TestError(c.Provider().ExecutionContext().GetName(), errMsg, c)
		}
	}()
	step(stCtx)
}

// WithNewAsyncStep opens nesting for struct.Step
// Any other struct.Step that will be added to struct.AllureResult object will be added as child step
func (c *common) WithNewAsyncStep(stepName string, step func(ctx provider.StepCtx), params ...allure.Parameter) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.WithNewStep(stepName, step, params...)
	}()
}

func (c *common) Attachment(attachment *allure.Attachment) {
	c.Provider().ExecutionContext().AddAttachment(attachment)
}

func (c *common) NewAttachment(name string, mimeType allure.MimeType, content []byte) {
	c.Provider().ExecutionContext().AddAttachment(allure.NewAttachment(name, mimeType, content))
}
