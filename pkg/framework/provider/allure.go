package provider

import (
	"github.com/ozontech/allure-go/pkg/allure"
)

type SystemLabels interface {
	Package(value string)
	FrameWork(value string)
	Host(value string)
	Thread(value string)
	Language(value string)
}

type SuiteLabels interface {
	AddSuiteLabel(value string)
	AddSubSuite(value string)
	AddParentSuite(value string)
}

type DescriptionLabels interface {
	ID(value string)
	AllureID(value string)
	Epic(value string)
	Feature(value string)
	Story(value string)
	Severity(severityType allure.SeverityType)
	Tag(value string)
	Tags(values ...string)
	Owner(value string)
	Lead(value string)
	Label(label allure.Label)
	Labels(labels ...allure.Label)
	ReplaceLabel(label allure.Label)
}

type Links interface {
	SetIssue(issue string)
	SetTestCase(testCase string)
	Link(link allure.Link)
}

type DescriptionFields interface {
	Title(title string)
	Description(description string)
}

type AllureSteps interface {
	Step(step *allure.Step)
	NewStep(stepName string, params ...allure.Parameter)
}

type Attachments interface {
	WithAttachments(attachment ...*allure.Attachment)
	WithNewAttachment(name string, mimeType allure.MimeType, content []byte)
}

type AllureForward interface {
	DescriptionLabels
	SuiteLabels
	Links
	DescriptionFields
	AllureSteps
	Attachments
}

type AllureForwardFull interface {
	AllureForward
	SystemLabels
}
