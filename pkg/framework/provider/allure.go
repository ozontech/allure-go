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
	Layer(value string)
	Feature(value string)
	Story(value string)
	Severity(severityType allure.SeverityType)
	Tag(value string)
	Tags(values ...string)
	Owner(value string)
	Lead(value string)
	Label(label *allure.Label)
	Labels(labels ...*allure.Label)
	ReplaceLabel(label *allure.Label)
}

type Links interface {
	SetIssue(issue string)
	SetTestCase(testCase string)
	Link(link *allure.Link)
	TmsLink(tmsCase string)
	TmsLinks(tmsCases ...string)
}

type DescriptionFields interface {
	Title(args ...interface{})
	Titlef(format string, args ...interface{})
	Description(args ...interface{})
	Descriptionf(format string, args ...interface{})
	Stage(args ...interface{})
	Stagef(format string, args ...interface{})
}

type AllureSteps interface {
	Step(step *allure.Step)
	NewStep(stepName string, params ...*allure.Parameter)
}

type Attachments interface {
	WithAttachments(attachment ...*allure.Attachment)
	WithNewAttachment(name string, mimeType allure.MimeType, content []byte)
}

type Parameters interface {
	WithParameters(params ...*allure.Parameter)
	WithNewParameters(kv ...interface{})
}

type AllureForward interface {
	DescriptionLabels
	SuiteLabels
	Links
	DescriptionFields
	AllureSteps
	Attachments
	Parameters
}

type AllureForwardFull interface {
	AllureForward
	SystemLabels
}
