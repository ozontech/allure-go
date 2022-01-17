//go:build examples
// +build examples

package suite_demo

import (
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type JSONStruct struct {
	Message string `json:"message"`
}

type AttachmentTestDemoSuite struct {
	suite.Suite
}

func (s *AttachmentTestDemoSuite) TestAttachment() {
	s.Epic("Demo")
	s.Feature("Attachments")
	s.Title("Test Attachments")
	s.Description(`
		Test's test body and all steps inside can contain attachments`)

	s.Tags("Attachments", "BeforeAfter", "Steps")

	attachmentText := `THIS IS A TEXT ATTACHMENT`
	s.Attachment(allure.NewAttachment("Text Attachment if TestAttachment", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Step A")
	var ExampleJson = JSONStruct{"this is JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for Step A", allure.JSON, attachmentJSON))
	s.Step(step)
}

type AttachmentDemoSuite struct {
	suite.Suite
}

func (s *AttachmentDemoSuite) BeforeAll() {
	// this action will create a step in Set up
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	s.Attachment(allure.NewAttachment("Text Attachment for Before suite", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Before suite Step")
	var ExampleJson = JSONStruct{"This is BeforeAll JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for Before suite Step", allure.JSON, attachmentJSON))
	s.Step(step)
}

func (s *AttachmentDemoSuite) BeforeEach() {
	// this action will create a step in Set up
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	s.Attachment(allure.NewAttachment("Text Attachment for Before Test", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Before Test Step")
	var ExampleJson = JSONStruct{"This is BeforeEach JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for Before Test Step", allure.JSON, attachmentJSON))
	s.Step(step)
}

func (s *AttachmentDemoSuite) AfterAll() {
	// this action will create a step in Tear down
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	s.Attachment(allure.NewAttachment("Text Attachment for After suite", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("After suite Step")
	var ExampleJson = JSONStruct{"This is AfterAll JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for After suite Step", allure.JSON, attachmentJSON))
	s.Step(step)
}

func (s *AttachmentDemoSuite) AfterEach() {
	// this action will create a step in Tear down
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	s.Attachment(allure.NewAttachment("Text Attachment for After Test", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("After Test Step")
	var ExampleJson = JSONStruct{"This is AfterEach JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for After Test Step", allure.JSON, attachmentJSON))
	s.Step(step)
}

func (s *AttachmentDemoSuite) TestAttachment_1() {
	s.Epic("Demo")
	s.Feature("Attachments")
	s.Title("Test with Before/After Attachments")
	s.Description(`
		Test "Set up", Test "Tear down", suite "Set up", suite "Tear down" 
		and all steps inside can contain attachments`)

	s.Tags("Attachments", "BeforeAfter", "Steps")
}

type NestedAttachmentDemoSuite struct {
	suite.Suite
}

func (s *NestedAttachmentDemoSuite) BeforeAll() {
	s.WithNewStep("SetupSuite step", func() {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		s.AddAttachmentToNested(allure.NewAttachment("Text Attachment for Before all", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) AfterAll() {
	s.WithNewStep("TearDownSuite step", func() {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		s.AddAttachmentToNested(allure.NewAttachment("Text Attachment for After all", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) BeforeEach() {
	s.WithNewStep("SetupTest step", func() {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		s.AddAttachmentToNested(allure.NewAttachment("Text Attachment for Before Test", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) AfterEach() {
	s.WithNewStep("TearDownTest step", func() {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		s.AddAttachmentToNested(allure.NewAttachment("Text Attachment for After Test", allure.Text, []byte(attachmentText)))
	})
}

func (s *NestedAttachmentDemoSuite) TestNestedAttachment() {
	s.Epic("Demo")
	s.Feature("Attachments")
	s.Title("Test NestedAttachments")
	s.Description(`
		Test "Set up", Test "Tear down", suite "Set up", suite "Tear down" test body has step with attachments.`)

	s.Tags("Attachments", "Nesting", "Steps", "BeforeAfter")

	s.WithNewStep("TestNestedAttachment step", func() {
		attachmentText := `THIS IS A TEXT ATTACHMENT`
		s.AddAttachmentToNested(allure.NewAttachment("Text Attachment for After Test", allure.Text, []byte(attachmentText)))
	})
}

func TestAttachments(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(AttachmentTestDemoSuite))
	runner.RunSuite(t, new(AttachmentDemoSuite))
	runner.RunSuite(t, new(NestedAttachmentDemoSuite))
}
