package provider

import (
	"fmt"
	"runtime/debug"

	"github.com/ozontech/allure-go/pkg/allure"
)

type AllureActions interface {
	Attachment(attachment *allure.Attachment)
	Step(step *allure.Step)
	NewStep(stepName string)
	InnerStep(stepParent *allure.Step, step *allure.Step)
	WithStep(step *allure.Step, f func())
	WithNewStep(stepName string, f func())

	AddAttachmentToNested(attachment *allure.Attachment)
	AddParameterToNested(param allure.Parameter)
	AddParametersToNested(params []allure.Parameter)
	AddNewParameterToNested(key, value string)
	AddNewParametersToNested(kv ...string)
}

/*
Attachments
*/

// Attachment adds attachment to test result
func (t *T) Attachment(attachment *allure.Attachment) {
	t.getState().addAttachment(attachment)
}

// AddAttachmentToNested adds attachment to last opened nested step
func (t *T) AddAttachmentToNested(attachment *allure.Attachment) {
	t.getState().addNestedAttachment(attachment)
}

/*
Steps
*/

// Step adds step to test result TODO: SuiteSetup handling
func (t *T) Step(step *allure.Step) {
	t.getState().addStep(step)
}

func (t *T) NewStep(stepName string) {
	t.getState().addStep(allure.NewSimpleStep(stepName))
}

// InnerNewStep inits new step and attaches step as child to the stepParent
func (t *T) InnerNewStep(stepParent *allure.Step, stepName string) {
	step := allure.NewSimpleStep(stepName)
	step.Parent = stepParent.GetUUID()
	t.Step(step)
}

// InnerStep attaches step as child to the stepParent
func (t *T) InnerStep(stepParent *allure.Step, step *allure.Step) {
	step.Parent = stepParent.GetUUID()
	t.Step(step)
}

// WithStep opens nesting for struct.Step
// Any other struct.Step that will be added to struct.AllureResult object will be added as child step
func (t *T) WithStep(step *allure.Step, f func()) {
	t.nestedStep(step)
	defer func() {
		r := recover()
		if r != nil {
			errMsg := fmt.Sprintf("test panicked: %v\n%s", r, debug.Stack())
			if result := t.GetResult(); result != nil {
				step.Status = allure.Broken
				result.Status = allure.Broken
				result.StatusDetails.Message = errMsg[:100]
				result.StatusDetails.Trace = errMsg
			}
			t.finishNesting()
			t.Errorf(errMsg)
			t.FailNow()
		}
		t.finishNesting()
	}()
	f()
}

// WithNewStep same as WithStep but consumes stepName to initiate struct.Step with struct.NewSimpleStep
func (t *T) WithNewStep(stepName string, f func()) {
	t.WithStep(allure.NewSimpleStep(stepName), f)
}

// NestedStep announces that all steps that will be before FinishNesting(t *testing.T)
// will be children of nested step
func (t *T) nestedStep(step *allure.Step) {
	t.getState().addNested(step)
}

// finishNesting close last nested step
func (t *T) finishNesting() {
	t.getState().finishNesting()
}

/*
Parameters
*/

// AddParameterToNested adds struct.Parameter to opened nested struct.Step
func (t *T) AddParameterToNested(param allure.Parameter) {
	t.AddParametersToNested([]allure.Parameter{param})
}

// AddParametersToNested adds few struct.Parameter to opened nested struct.Step
func (t *T) AddParametersToNested(params []allure.Parameter) {
	t.getState().addNestedParam(params...)
}

// AddNewParameterToNested adds struct.Parameter to opened nested struct.Step
func (t *T) AddNewParameterToNested(key, value string) {
	t.AddParametersToNested(allure.NewParameters(key, value))
}

// AddNewParametersToNested adds struct.Parameter to opened nested struct.Step
func (t *T) AddNewParametersToNested(kv ...string) {
	t.AddParametersToNested(allure.NewParameters(kv...))
}
