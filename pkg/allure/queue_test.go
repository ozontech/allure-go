package allure

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TODO: make it thread-safe

func TestQueueImplements(t *testing.T) {
	assert.Implements(t, (*FiloQueue)(nil), new(NestingQueue))
}

func TestNewNestingQueue(t *testing.T) {
	queue := NewNestingQueue()
	require.NotNil(t, queue.queue)
	require.Equal(t, queue.count, 0)
}

func TestNestingQueue_Push(t *testing.T) {
	queue := NewNestingQueue()
	queue.Push(NewSimpleStep("Step"))
	require.Len(t, queue.queue, 1)
	require.Equal(t, queue.count, 1)
}

func TestNestingQueue_Last(t *testing.T) {
	queue := NewNestingQueue()
	stepToPush := NewSimpleStep("Step")
	queue.Push(stepToPush)
	step := queue.Last()
	require.Len(t, queue.queue, 1)
	require.Equal(t, queue.count, 1)

	require.Equal(t, stepToPush.Name, step.Name)
	require.Equal(t, stepToPush.uuid, step.uuid)
	require.Equal(t, stepToPush.Start, step.Start)
	require.Equal(t, stepToPush.Stop, step.Stop)
	require.Equal(t, stepToPush.Steps, step.Steps)
	require.Equal(t, stepToPush, step)
}

func TestNestingQueue_Pop(t *testing.T) {
	queue := NewNestingQueue()
	stepToPush := NewSimpleStep("Step")
	queue.Push(stepToPush)
	step := queue.Pop()
	require.Len(t, queue.queue, 0)
	require.Equal(t, queue.count, 0)

	require.Equal(t, stepToPush.Name, step.Name)
	require.Equal(t, stepToPush.uuid, step.uuid)
	require.Equal(t, stepToPush.Start, step.Start)
	require.Equal(t, stepToPush.Stop, step.Stop)
	require.Equal(t, stepToPush.Steps, step.Steps)
	require.Equal(t, stepToPush, step)
}
