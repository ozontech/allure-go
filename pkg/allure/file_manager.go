package allure

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileManager interface {
	CreateFile(name string, content []byte) error
}

type fileManager struct {
	resultsPath string
}

func NewFileManager() FileManager {
	resultsPath := getResultPath()
	fm := &fileManager{resultsPath: resultsPath}
	fm.createOutputDir()
	return fm
}

func (m *fileManager) CreateFile(name string, content []byte) error {
	file := fmt.Sprintf("%s/%s", m.resultsPath, name)
	return ioutil.WriteFile(file, content, fileSystemPermissionCode)
}

func (m *fileManager) createOutputDir() {
	isExists, err := exists(m.resultsPath)
	if err != nil {
		panic(err)
	}
	if !isExists {
		_ = os.MkdirAll(m.resultsPath, os.ModePerm)
	}
}

func getOutputFolderName() string {
	outputFolderName := os.Getenv(outputFolderEnvKey)
	if outputFolderName != "" {
		return outputFolderName
	}
	return "allure-results"
}

func getResultPath() string {
	resultsPathToOutput := os.Getenv(resultsPathEnvKey)
	outputFolderName := getOutputFolderName()
	if resultsPathToOutput != "" {
		return fmt.Sprintf("%s/%s", resultsPathToOutput, outputFolderName)
	}
	return fmt.Sprintf("./%s", outputFolderName)
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
