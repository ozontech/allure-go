package testplan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	once = sync.Once{}

	testPlan *TestPlan
)

// Path to testplan.json
const testPlanPath = "ALLURE_TESTPLAN_PATH"

type TestCase struct {
	ID       int    `json:"id"`
	Selector string `json:"selector"`
}

type TestPlan struct {
	Version string      `json:"version"`
	Tests   []*TestCase `json:"tests"`
}

func GetTestPlan() *TestPlan {
	var (
		err error
	)

	testPlanOnce := func() {
		testPlan, err = newTestPlan()
		if err != nil {
			fmt.Printf("%s\n", err.Error())
		}
	}

	once.Do(testPlanOnce)
	return testPlan
}

func newTestPlan() (*TestPlan, error) {
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

	plan := &TestPlan{}
	err = json.Unmarshal(testPlanRaw, plan)
	if err != nil {
		return nil, err
	}
	return testPlan, nil
}

func (p *TestPlan) IsSelected(id, selector string) bool {
	for _, t := range p.Tests {
		if t.Selector == selector {
			return true
		}
	}
	return false
}
