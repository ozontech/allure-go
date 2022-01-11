package allure

import "os"

// createOutputFolder ...
func createOutputFolder(folder string) {
	isExists, err := exists(folder)
	if err != nil {
		panic(err)
	}
	if !isExists {
		_ = os.MkdirAll(folder, os.ModePerm)
	}
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
