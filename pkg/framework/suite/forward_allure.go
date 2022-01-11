package suite

import "github.com/koodeex/allure-testify/pkg/allure"

/*
	Forward Actions
*/

// Attachment ...
func (suite *Suite) Attachment(attachment *allure.Attachment) {
	suite.T().Attachment(attachment)
}

// AddAttachmentToNested ...
func (suite *Suite) AddAttachmentToNested(attachment *allure.Attachment) {
	suite.T().AddAttachmentToNested(attachment)
}

// Step ...
func (suite *Suite) Step(step *allure.Step) {
	suite.T().Step(step)
}

// NewStep ...
func (suite *Suite) NewStep(stepName string) {
	suite.T().NewStep(stepName)
}

// InnerStep ...
func (suite *Suite) InnerStep(stepParent *allure.Step, step *allure.Step) {
	suite.T().InnerStep(stepParent, step)
}

// InnerNewStep ...
func (suite *Suite) InnerNewStep(stepParent *allure.Step, stepName string) {
	suite.T().InnerNewStep(stepParent, stepName)
}

// WithStep ...
func (suite *Suite) WithStep(step *allure.Step, f func()) {
	suite.T().WithStep(step, f)
}

// WithNewStep ...
func (suite *Suite) WithNewStep(stepName string, f func()) {
	suite.T().WithNewStep(stepName, f)
}

// AddParameterToNested ...
func (suite *Suite) AddParameterToNested(param allure.Parameter) {
	suite.T().AddParameterToNested(param)
}

// AddParametersToNested ...
func (suite *Suite) AddParametersToNested(params []allure.Parameter) {
	suite.T().AddParametersToNested(params)
}

// AddNewParameterToNested ...
func (suite *Suite) AddNewParameterToNested(key, value string) {
	suite.T().AddNewParameterToNested(key, value)
}

// AddNewParametersToNested ...
func (suite *Suite) AddNewParametersToNested(kv ...string) {
	suite.T().AddNewParametersToNested(kv...)
}

/*
	Forward Info
*/

// Title ...
func (suite *Suite) Title(title string) {
	suite.T().Title(title)
}

// Description ...
func (suite *Suite) Description(description string) {
	suite.T().Description(description)
}

/*
	Forward Labels
*/

// ID ...
func (suite *Suite) ID(value string) {
	suite.T().ID(value)
}

// Epic ...
func (suite *Suite) Epic(value string) {
	suite.T().Epic(value)
}

// AddSuiteLabel ...
func (suite *Suite) AddSuiteLabel(value string) {
	suite.T().AddSuiteLabel(value)
}

// AddSubSuite ...
func (suite *Suite) AddSubSuite(value string) {
	suite.T().AddSubSuite(value)
}

// AddParentSuite ...
func (suite *Suite) AddParentSuite(value string) {
	suite.T().AddParentSuite(value)
}

// Feature ...
func (suite *Suite) Feature(value string) {
	suite.T().Feature(value)
}

// Story ...
func (suite *Suite) Story(value string) {
	suite.T().Story(value)
}

// Tag ...
func (suite *Suite) Tag(value string) {
	suite.T().Tag(value)
}

// Tags ...
func (suite *Suite) Tags(values ...string) {
	suite.T().Tags(values...)
}

// Package ...
func (suite *Suite) Package(value string) {
	suite.T().Package(value)
}

// Severity ...
func (suite *Suite) Severity(value allure.SeverityType) {
	suite.T().Severity(value)
}

// FrameWork ...
func (suite *Suite) FrameWork(value string) {
	suite.T().FrameWork(value)
}

// Host ...
func (suite *Suite) Host(value string) {
	suite.T().Host(value)
}

// Thread ...
func (suite *Suite) Thread(value string) {
	suite.T().Thread(value)
}

// Language ...
func (suite *Suite) Language(value string) {
	suite.T().Language(value)
}

// Owner ...
func (suite *Suite) Owner(value string) {
	suite.T().Owner(value)
}

// Lead ...
func (suite *Suite) Lead(value string) {
	suite.T().Lead(value)
}

// Label ...
func (suite *Suite) Label(label allure.Label) {
	suite.T().Label(label)
}

// Labels ...
func (suite *Suite) Labels(labels ...allure.Label) {
	suite.T().Labels(labels...)
}

/*
	Forward Links
*/

// SetIssue ...
func (suite *Suite) SetIssue(issue string) {
	suite.T().SetIssue(issue)
}

// SetTestCase ...
func (suite *Suite) SetTestCase(testCase string) {
	suite.T().SetTestCase(testCase)
}

// Link ...
func (suite *Suite) Link(link allure.Link) {
	suite.T().Link(link)
}
