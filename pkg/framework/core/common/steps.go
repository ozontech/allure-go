package common

import (
	"fmt"
	"runtime/debug"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// WithNewStep opens nesting for struct.Step
// Any other struct.Step that will be added to struct.AllureResult object will be added as child step
func (c *Common) WithNewStep(stepName string, step func(ctx provider.StepCtx), params ...*allure.Parameter) {
	stCtx := NewStepCtx(c, c.Provider, stepName, params...)
	defer c.Step(stCtx.CurrentStep())
	defer func() {
		r := recover()
		stCtx.WG().Wait()
		stCtx.CurrentStep().Finish()
		if r != nil {
			ctxName := c.ExecutionContext().GetName()
			errMsg := fmt.Sprintf("%s panicked: %v\n%s", ctxName, r, debug.Stack())
			stCtx.Broken()
			TestError(c.TestingT, c.Provider, c.Provider.ExecutionContext().GetName(), errMsg)
		}
	}()
	step(stCtx)
}

// WithNewAsyncStep opens nesting for struct.Step
// Any other struct.Step that will be added to struct.AllureResult object will be added as child step
func (c *Common) WithNewAsyncStep(stepName string, step func(ctx provider.StepCtx), params ...*allure.Parameter) {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		c.WithNewStep(stepName, step, params...)
	}()
}
