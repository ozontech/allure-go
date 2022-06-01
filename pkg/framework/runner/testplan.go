package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Path to testplan.json
const testPlanPath = "ALLURE_TESTPLAN_PATH"

type TestCase struct {
	ID       string `json:"id"`
	Selector string `json:"selector"`
}

type TestPlan struct {
	Version string      `json:"version"`
	Tests   []*TestCase `json:"tests"`
}

func NewTestPlan() (*TestPlan, error) {
	filePath := os.Getenv(testPlanPath)
	if filePath == "" {
		return nil, fmt.Errorf("{%s} environment variable not set", testPlanPath)
	}
	if !strings.HasSuffix(filePath, ".json") {
		return nil, fmt.Errorf("%s environment variable has a wrong format. Please, set path to .json file. Current path:%s", testPlanPath, filePath)
	}

	testPlanRaw, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return nil, err
	}

	testPlan := &TestPlan{}
	err = json.Unmarshal(testPlanRaw, testPlan)
	if err != nil {
		return nil, err
	}
	return testPlan, nil
}

func (p *TestPlan) IsSelected(id, selector string) bool {
	for _, t := range p.Tests {
		if t.ID == id || t.Selector == selector {
			return true
		}
	}
	return false
}
