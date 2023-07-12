package allure

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStep(t *testing.T) {
	stepName := "Step A"
	stepStatus := Passed
	stepStart := time.Now().UnixNano() / int64(time.Millisecond)
	stepStop := time.Now().UnixNano()/int64(time.Millisecond) + 1
	parameters := []*Parameter{
		{"Param1", []byte("val1")},
		{"Param2", []byte("val2")},
	}
	step := NewStep(stepName, stepStatus, stepStart, stepStop, parameters)
	assert.Equal(t, stepName, step.Name)
	assert.Equal(t, stepStatus, step.Status)
	assert.Equal(t, stepStart, step.Start)
	assert.Equal(t, stepStop, step.Stop)
	require.Equal(t, 2, len(step.Parameters))
	assert.Equal(t, parameters[0].Name, step.Parameters[0].Name)
	assert.Equal(t, parameters[0].GetValue(), step.Parameters[0].GetValue())
	assert.Equal(t, parameters[1].Name, step.Parameters[1].Name)
	assert.Equal(t, parameters[1].GetValue(), step.Parameters[1].GetValue())
	assert.Nil(t, step.GetParent())
	assert.Nil(t, step.Attachments)
}

func TestNewSimpleStep(t *testing.T) {
	stepName := "Step A"
	step := NewSimpleStep(stepName)
	assert.Equal(t, stepName, step.Name)
	assert.Equal(t, Passed, step.Status)
	assert.NotEqual(t, 0, step.Start)
	assert.NotEqual(t, 0, step.Stop)
	assert.Nil(t, step.GetParent())
	assert.Nil(t, step.Parameters)
	assert.Nil(t, step.Attachments)
}

func TestStep_Begin(t *testing.T) {
	step := new(Step)
	now := GetNow()
	step.Begin()
	require.Equal(t, now, step.Start)
}

func TestStep_Finish(t *testing.T) {
	step := new(Step)
	now := GetNow()
	step.Finish()
	require.Equal(t, now, step.Stop)
}

func TestStep_Passed(t *testing.T) {
	step := new(Step)
	step.Passed()
	require.Equal(t, Passed, step.Status)
}

func TestStep_Failed(t *testing.T) {
	step := new(Step)
	step.Failed()
	require.Equal(t, Failed, step.Status)
}

func TestStep_Skipped(t *testing.T) {
	step := new(Step)
	step.Skipped()
	require.Equal(t, Skipped, step.Status)
}

func TestStep_Broken(t *testing.T) {
	step := new(Step)
	step.Broken()
	require.Equal(t, Broken, step.Status)
}

func TestStep_PrintAttachments(t *testing.T) {
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	step := new(Step)
	step.Attachments = append(step.Attachments, NewAttachment("Text Attachment if TestAttachment", Text, []byte(attachmentText)))

	step.PrintAttachments()

	defer os.RemoveAll(allureDir)

	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 1)
	var attachFile *os.File
	defer attachFile.Close()

	f := files[0]
	attachFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := ioutil.ReadAll(attachFile)
	require.NoError(t, readErr)
	require.Equal(t, attachmentText, string(bytes))
}

func TestStep_WithAttachments(t *testing.T) {
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	attachment := NewAttachment("Text Attachment if TestAttachment", Text, []byte(attachmentText))
	step := new(Step)
	step.WithAttachments(attachment)
	require.NotNil(t, step.Attachments)
	require.Len(t, step.Attachments, 1)

	att1 := step.Attachments[0]
	require.Equal(t, attachment.Name, att1.Name)
	require.Equal(t, attachment.Type, att1.Type)
	require.Equal(t, attachment.Source, att1.Source)
	require.Equal(t, attachment.content, att1.content)
}

func TestStep_WithChild(t *testing.T) {
	childStep := NewSimpleStep("Child Step")
	step := new(Step)
	step.WithChild(childStep)

	require.Len(t, step.Steps, 1)
	st := step.Steps[0]
	require.Equal(t, childStep, st)
}

func TestStep_WithNewParameters_even(t *testing.T) {
	step := new(Step)
	step.WithNewParameters("param1", "val1", "param2", "val2")
	require.NotNil(t, step.Parameters)
	require.Len(t, step.Parameters, 2)
	require.Equal(t, "param1", step.Parameters[0].Name)
	require.Equal(t, "val1", step.Parameters[0].GetValue())
	require.Equal(t, "param2", step.Parameters[1].Name)
	require.Equal(t, "val2", step.Parameters[1].GetValue())
}

func TestStep_WithNewParameters_odd(t *testing.T) {
	step := new(Step)
	step.WithNewParameters("param1", "val1", "param2")
	require.NotNil(t, step.Parameters)
	require.Len(t, step.Parameters, 1)
	require.Equal(t, "param1", step.Parameters[0].Name)
	require.Equal(t, "val1", step.Parameters[0].GetValue())
}

func TestStep_WithParameters(t *testing.T) {
	step := new(Step)
	step.WithParameters(NewParameter("param1", "val1"), NewParameter("param2", "val2"))
	require.NotNil(t, step.Parameters)
	require.Len(t, step.Parameters, 2)
	require.Equal(t, "param1", step.Parameters[0].Name)
	require.Equal(t, "val1", step.Parameters[0].GetValue())
	require.Equal(t, "param2", step.Parameters[1].Name)
	require.Equal(t, "val2", step.Parameters[1].GetValue())
}

func TestStep_WithParent(t *testing.T) {
	parentStep := NewSimpleStep("Parent Step")

	step := new(Step)
	parentStep.WithParent(step)
	require.Len(t, step.Steps, 1)
	st := step.Steps[0]
	require.Equal(t, parentStep, st)
}

func TestStep_GetParent(t *testing.T) {
	parentStep := NewSimpleStep("Parent Step")

	step := new(Step)
	parentStep.WithParent(step)
	require.Len(t, step.Steps, 1)
	st := parentStep.GetParent()
	require.Equal(t, step, st)
}
