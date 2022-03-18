package allure

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFileManager(t *testing.T) {
	fm := NewFileManager()
	require.NotNil(t, fm)
}

func TestCreateOutputDir(t *testing.T) {
	fm := &fileManager{allureDir}
	fm.createOutputDir()
	defer os.RemoveAll(allureDir)
	require.DirExists(t, allureDir)
}

func TestGetOutputFolderName_noEnv(t *testing.T) {
	require.Equal(t, "allure-results", getOutputFolderName())
}

func TestGetOutputFolderName_Env(t *testing.T) {
	os.Setenv(outputFolderEnvKey, "not_allure_results")
	defer os.Setenv(outputFolderEnvKey, "")
	require.Equal(t, "not_allure_results", getOutputFolderName())
}

func TestGetResultPath_noEnv(t *testing.T) {
	require.Equal(t, allureDir, getResultPath())
}

func TestGetResultPath_Env(t *testing.T) {
	os.Setenv(resultsPathEnvKey, "not_allure_results")
	defer os.Setenv(resultsPathEnvKey, "")
	require.Equal(t, "not_allure_results/allure-results", getResultPath())
}

func TestFileManager_CreateFile(t *testing.T) {
	fileContent := `SOME TEXT`
	fm := NewFileManager()
	fm.CreateFile("test.txt", []byte(fileContent))
	require.DirExists(t, allureDir)
	defer os.RemoveAll(allureDir)

	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 1)
	var file *os.File
	defer file.Close()

	f := files[0]
	file, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := ioutil.ReadAll(file)
	require.NoError(t, readErr)
	require.Equal(t, fileContent, string(bytes))
}
