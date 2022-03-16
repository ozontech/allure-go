package provider

import "github.com/ozontech/allure-go/pkg/allure"

type StepCtx interface {
	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	WithNewStep(stepName string, step func(ctx StepCtx), params ...allure.Parameter)
	WithNewAsyncStep(stepName string, step func(ctx StepCtx), params ...allure.Parameter)
	WithParameters(parameters ...allure.Parameter)
	WithNewParameters(kv ...string)
	WithAttachments(attachments ...*allure.Attachment)
	WithNewAttachment(name string, mimeType allure.MimeType, content []byte)
	NewStep(stepName string, parameters ...allure.Parameter)
	Fail()
	Broken()

	Assert() Asserts
	Require() Asserts

	T() T
	CurrentStep() *allure.Step
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}
