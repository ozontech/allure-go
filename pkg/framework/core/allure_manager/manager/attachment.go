package manager

import "github.com/ozontech/allure-go/pkg/allure"

// Attachment adds attachment to report in case of current execution context
func (a *allureManager) Attachment(attachment *allure.Attachment) {
	a.ExecutionContext().AddAttachment(attachment)
}

// NewAttachment creates and adds attachment to report in case of current execution context
func (a *allureManager) NewAttachment(name string, mimeType allure.MimeType, content []byte) {
	a.ExecutionContext().AddAttachment(allure.NewAttachment(name, mimeType, content))
}
