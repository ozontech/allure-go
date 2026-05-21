package allure

import (
	"fmt"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	allureDir = "allure-results"
)

func TestNewContainer(t *testing.T) {
	container := NewContainer()
	require.NotNil(t, container)
	require.Zero(t, container.Start)
	require.Zero(t, container.Stop)
	require.NotNil(t, container.UUID)
}

func TestContainer_AddChild(t *testing.T) {
	container := NewContainer()
	result := NewResult("Mock Name", "Full Mock Name")
	container.AddChild(result.UUID)
	require.NotEmpty(t, container.Children)
	require.Len(t, container.Children, 1)
	require.Equal(t, result.UUID, container.Children[0])
}

func TestContainer_Begin(t *testing.T) {
	container := NewContainer()
	begin := GetNow()
	container.Begin()
	assert.Equal(t, container.Start, begin)
}

func TestContainer_Finish(t *testing.T) {
	container := NewContainer()
	finish := GetNow()
	container.Finish()
	assert.Equal(t, container.Stop, finish)
}

func TestContainer_Print(t *testing.T) {
	container := NewContainer()
	containerBefore := NewSimpleStep("Before")
	containerAfter := NewSimpleStep("After")
	container.Befores = append(container.Befores, containerBefore)
	container.Afters = append(container.Afters, containerAfter)

	_ = container.Print()
	defer os.RemoveAll(allureDir)
	require.DirExists(t, allureDir)
	files, _ := os.ReadDir(allureDir)
	require.Len(t, files, 1)
	var jsonFile *os.File
	defer jsonFile.Close()

	f := files[0]
	emptyContainer := &Container{}
	jsonFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := io.ReadAll(jsonFile)
	require.NoError(t, readErr)
	unMarshallErr := sonic.Unmarshal(bytes, emptyContainer)
	require.NoError(t, unMarshallErr)
	require.Equal(t, container.UUID, emptyContainer.UUID)

	require.NotEmpty(t, emptyContainer.Befores)
	require.NotEmpty(t, emptyContainer.Afters)
	require.Len(t, emptyContainer.Befores, 1)
	require.Len(t, emptyContainer.Afters, 1)

	before := emptyContainer.Befores[0]
	after := emptyContainer.Afters[0]

	require.Equal(t, containerBefore, before)
	require.Equal(t, containerAfter, after)
}

func TestMarshallingContainer(t *testing.T) {
	container := NewContainer()
	b := `grpc_cli call --json_input --json_output myendpoint.com:82 bla.bla/GetBla '{"bla":"blabla"}' -metadata 'x-bla::x-app-bla:bla-qa/bla-tests:x-bla-version::x-nocache:true:x-pf-nocache:true:x-s2s:*****'`
	container.Befores = append(container.Befores, &Step{
		Name:        "name",
		Status:      Passed,
		Attachments: nil,
		Start:       time.Now().UnixNano(),
		Stop:        time.Now().UnixNano(),
		Steps:       nil,
		Parameters:  []*Parameter{NewParameter("p", b)},
	})

	_, err := container.ToJSON()
	require.NoError(t, err)
}
