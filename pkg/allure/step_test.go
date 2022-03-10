package allure

import (
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
	assert.Nil(t, step.GetParent())
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
	assert.Nil(t, step.GetParent())
	assert.Nil(t, step.Parameters)
	assert.Nil(t, step.Attachments)
	assert.NotNil(t, step.uuid)
}
