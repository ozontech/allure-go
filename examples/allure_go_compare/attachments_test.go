//go:build allure_go
// +build allure_go

package allure_go_compare

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type AllureGoAttachments struct {
	suite.Suite
}

/* Allure-Go style:
func TestTextAttachment(t *testing.T) {
	allure.Test(t, allure.Description("Testing a text attachment"), allure .Action(func() {
		_ = allure.AddAttachment("text!", allure.TextPlain, []byte("Some text!"))
	}))
}
*/

func (s *AllureGoAttachments) TestTextAttachment() {
	s.Epic("Compare with allure-go")
	s.Description("Testing a text attachment")
	s.Attachment(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
}

/* TestTextAttachmentToStep Allure-Go style:
func TestTextAttachmentToStep(t *testing.T) {
	allure.Test(t, allure.Description("Testing a text attachment"), allure.Action(func() {
		allure.Step(allure.Description("adding a text attachment"), allure.Action(func() {
			_ = allure.AddAttachment("text!", allure.TextPlain, []byte("Some text!"))
		}))
	}))
}
*/

func (s *AllureGoAttachments) TestTextAttachmentToStepV1() {
	s.Epic("Compare with allure-go")
	s.Description("Testing a text attachment")
	s.WithNewStep("adding a text attachment", func() {
		s.AddAttachmentToNested(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
	})
}

func (s *AllureGoAttachments) TestTextAttachmentToStepV2() {
	s.Epic("Compare with allure-go")
	s.Description("Testing a text attachment")
	s.Step(allure.NewSimpleStep("adding a text attachment").
		WithAttachment(allure.NewAttachment("text!", allure.Text, []byte("Some text!"))))
}

func (s *AllureGoAttachments) TestTextAttachmentToStepV3() {
	s.Epic("Compare with allure-go")
	s.Description("Testing a text attachment")
	step := allure.NewSimpleStep("adding a text attachment")
	step.Attachment(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
	s.Step(step)
}

// Also we provide some complex actions with steps
// In this example we initialize step, declare it like nested and add new attachment to it.

func (s *AllureGoAttachments) TestTextAttachmentToStepV4() {
	s.Epic("Compare with allure-go")
	s.Description("Testing a text attachment")
	step := allure.NewSimpleStep("adding a text attachment")
	step.Attachment(allure.NewAttachment("text!", allure.Text, []byte("Some text!")))
	s.WithStep(step, func() {
		s.AddAttachmentToNested(allure.NewAttachment("text inside step!", allure.Text, []byte("Another text!")))
	})
}

func TestAllureGoAttachments(t *testing.T) {
	runner.RunSuite(t, new(AllureGoAttachments))
}
