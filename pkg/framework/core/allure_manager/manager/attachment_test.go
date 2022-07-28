package manager

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/stretchr/testify/require"
)

type execMockAttach struct {
	name   string
	steps  []*allure.Step
	attach []*allure.Attachment
}

func newExecMockAttach(name string) *execMockAttach {
	return &execMockAttach{
		name:   name,
		steps:  make([]*allure.Step, 0),
		attach: make([]*allure.Attachment, 0),
	}
}

func (m *execMockAttach) AddStep(step *allure.Step) {
	m.steps = append(m.steps, step)
}

func (m *execMockAttach) AddAttachments(attachments ...*allure.Attachment) {
	m.attach = append(m.attach, attachments...)
}

func (m *execMockAttach) GetName() string {
	return m.name
}

func TestAllureManager_Attachment(t *testing.T) {
	mock := newExecMockAttach(constants.TestContextName)
	attach := allure.NewAttachment("testAttach", allure.Text, []byte("test"))
	manager := allureManager{executionContext: mock}

	manager.WithAttachments(attach)
	require.NotEmpty(t, mock.attach)
	require.Len(t, mock.attach, 1)
	require.Equal(t, mock.attach[0], attach)
}

func TestAllureManager_NewAttachment(t *testing.T) {
	mock := newExecMockAttach(constants.TestContextName)
	manager := allureManager{executionContext: mock}
	manager.WithNewAttachment("testAttach", allure.Text, []byte("test"))
	require.NotEmpty(t, mock.attach)
	require.Len(t, mock.attach, 1)
	require.Equal(t, "testAttach", mock.attach[0].Name)
	require.Equal(t, allure.Text, mock.attach[0].Type)
	require.Equal(t, []byte("test"), mock.attach[0].GetContent())
}
