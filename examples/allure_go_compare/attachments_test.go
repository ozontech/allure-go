//go:build allure_go_new
// +build allure_go_new

package allure_go_compare

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type AllureGoAttachments struct {
	suite.Suite
}

/* Allure-Go style:
func TestTextAttachment(t *testing.T) {
	allure.Test(t, allure.Description("Testing a text attachment"), allure .Action(func() {
		_ = allure.AddAttachments("text!", allure.TextPlain, []byte("Some text!"))
	}))
}
*/

func (s *AllureGoAttachments) TestTextAttachment(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("Testing a text attachment")
	t.WithAttachments(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
}

/* TestTextAttachmentToStep Allure-Go style:
func TestTextAttachmentToStep(t *testing.T) {
	allure.Test(t, allure.Description("Testing a text attachment"), allure.Action(func() {
		allure.Step(allure.Description("adding a text attachment"), allure.Action(func() {
			_ = allure.AddAttachments("text!", allure.TextPlain, []byte("Some text!"))
		}))
	}))
}
*/

func (s *AllureGoAttachments) TestTextAttachmentToStepV1(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("Testing a text attachment")
	t.WithNewStep("adding a text attachment", func(ctx provider.StepCtx) {
		ctx.WithAttachments(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
	})
}

func (s *AllureGoAttachments) TestTextAttachmentToStepV2(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("Testing a text attachment")
	t.Step(allure.NewSimpleStep("adding a text attachment").
		WithAttachments(allure.NewAttachment("text!", allure.Text, []byte("Some text!"))))
}

func (s *AllureGoAttachments) TestTextAttachmentToStepV3(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("Testing a text attachment")
	step := allure.NewSimpleStep("adding a text attachment")
	step.WithAttachments(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
	t.Step(step)
}

// Also we provide some complex actions with steps
// In this example we initialize step, declare it like nested and add new attachment to it.

func (s *AllureGoAttachments) TestTextAttachmentToStepV4(t provider.T) {
	t.Epic("Compare with allure-go")
	t.Description("Testing a text attachment")
	step := allure.NewSimpleStep("adding a text attachment")
	step.WithAttachments(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
	t.WithNewStep("step", func(ctx provider.StepCtx) {
		ctx.WithAttachments(allure.NewAttachment("text inside step!", allure.Text, []byte("Another text!")))
	})
}

func TestAllureGoAttachments(t *testing.T) {
	suite.RunSuite(t, new(AllureGoAttachments))
}
