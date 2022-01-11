package allure

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStepImplements(t *testing.T) {
	assert.Implements(t, (*WithAttachments)(nil), new(Step))
	assert.Implements(t, (*WithTimer)(nil), new(Step))
	assert.Implements(t, (*IStep)(nil), new(Step))
}

func TestNewStep(t *testing.T) {
	stepName := "Step A"
	stepStatus := Passed
	stepStart := time.Now().UnixNano() / int64(time.Millisecond)
	stepStop := time.Now().UnixNano()/int64(time.Millisecond) + 1
	parameters := []Parameter{
		{"Param1", "val1"},
		{"Param2", "val2"},
	}
	step := NewStep(stepName, stepStatus, stepStart, stepStop, parameters)
	assert.Equal(t, stepName, step.Name)
	assert.Equal(t, stepStatus, step.Status)
	assert.Equal(t, stepStart, step.Start)
	assert.Equal(t, stepStop, step.Stop)
	require.Equal(t, 2, len(step.Parameters))
	assert.Equal(t, parameters[0], step.Parameters[0])
	assert.Equal(t, parameters[1], step.Parameters[1])
	assert.Equal(t, "", step.Parent)
	assert.Nil(t, step.Attachments)
	assert.NotNil(t, step.uuid)
}

func TestNewSimpleStep(t *testing.T) {
	stepName := "Step A"
	step := NewSimpleStep(stepName)
	assert.Equal(t, stepName, step.Name)
	assert.Equal(t, Passed, step.Status)
	assert.NotEqual(t, 0, step.Start)
	assert.NotEqual(t, 0, step.Stop)
	assert.Equal(t, "", step.Parent)
	assert.Nil(t, step.Parameters)
	assert.Nil(t, step.Attachments)
	assert.NotNil(t, step.uuid)
}

func TestNewSimpleInnerStep(t *testing.T) {
	stepNameParent := "Step A"
	stepNameChild := "Step B"
	stepParent := NewSimpleStep(stepNameParent)
	stepChild := NewSimpleInnerStep(stepNameChild, stepParent)
	assert.Equal(t, stepNameChild, stepChild.Name)
	assert.Equal(t, Passed, stepChild.Status)
	assert.NotEqual(t, 0, stepChild.Start)
	assert.NotEqual(t, 0, stepChild.Stop)
	assert.Equal(t, stepParent.GetUUID(), stepChild.Parent)
	assert.Nil(t, stepChild.Parameters)
	assert.Nil(t, stepChild.Attachments)
	assert.NotNil(t, stepChild.uuid)
}

func TestNewStepWithStart(t *testing.T) {
	stepName := "Step A"
	step := NewStepWithStart(stepName)
	assert.Equal(t, stepName, step.Name)
	assert.Equal(t, Passed, step.Status)
	assert.NotEqual(t, 0, step.Start)
	assert.Equal(t, int64(0), step.Stop)
	assert.Equal(t, "", step.Parent)
	assert.Nil(t, step.Parameters)
	assert.Nil(t, step.Attachments)
	assert.NotNil(t, step.uuid)
}
