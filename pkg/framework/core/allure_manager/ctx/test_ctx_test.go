package ctx

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/constants"
	"github.com/stretchr/testify/require"
)

func TestNewTestCtx(t *testing.T) {
	ctx := NewTestCtx(&allure.Result{})
	require.NotNil(t, ctx)
}

func TestTestCtx_GetName(t *testing.T) {
	th := testCtx{name: "test"}
	require.Equal(t, "test", th.GetName())
}

func TestTestCtx_AddStep(t *testing.T) {
	testStep := allure.NewSimpleStep("test")
	test := testCtx{name: constants.TestContextName, result: &allure.Result{}}
	test.AddStep(testStep)
	require.NotEmpty(t, test.result.Steps)
	require.Len(t, test.result.Steps, 1)
	require.Equal(t, testStep, test.result.Steps[0])
}

func TestTestCtx_AddAttachment(t *testing.T) {
	attach := allure.NewAttachment("testAttach", allure.Text, []byte("test"))
	test := testCtx{name: constants.TestContextName, result: &allure.Result{}}
	test.AddAttachments(attach)
	require.NotEmpty(t, test.result.Attachments)
	require.Len(t, test.result.Attachments, 1)
	require.Equal(t, attach, test.result.Attachments[0])
}
