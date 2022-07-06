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
	once     = sync.Once{}
	testPlan = initTestPlan()
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

func newTestPlan() (*TestPlan, error) {
	filePath := os.Getenv(testPlanPath)
	if filePath == "" {
		return nil, fmt.Errorf("{%s} environment variable not set", testPlanPath)
	}
	if !strings.HasSuffix(filePath, ".json") {
		return nil, fmt.Errorf("%s environment variable has a wrong format. Please, set path to .json file. Current path:%s", testPlanPath, filePath)
	}

	testPlanRaw, readFileErr := findTestPlan(filePath)
	if readFileErr != nil {
		return nil, readFileErr
	}

	plan := &TestPlan{}
	err := json.Unmarshal(testPlanRaw, plan)
	if err != nil {
		return nil, err
	}

	if plan != nil && len(plan.Tests) == 0 {
		return nil, fmt.Errorf("no any tests found in %s", filePath)
	}
	return plan, nil
}

// IsSelected returns true if selector matches with testplan selector
// TODO: ID parsing from TestOps
func (p *TestPlan) IsSelected(id, selector string) bool {
	for _, t := range p.Tests {
		if t.Selector == selector {
			return true
		}
	}
	return false
}

func initTestPlan() *TestPlan {
	var (
		err   error
		tPlan *TestPlan
	)

	testPlanOnce := func() {
		tPlan, err = newTestPlan()
		if err == nil {
			fmt.Printf("TestPlan found!")
		}
	}
	once.Do(testPlanOnce)
	return tPlan
}

// GetTestPlan ...
func GetTestPlan() *TestPlan {
	return testPlan
}

func findTestPlan(path string) (testPlanRaw []byte, readFileErr error) {
	testPlanRaw, readFileErr = ioutil.ReadFile(filepath.Clean(path))
	if readFileErr == nil && testPlanRaw != nil {
		return testPlanRaw, nil
	}
	dir, getWdErr := os.Getwd()
	if getWdErr != nil {
		return nil, getWdErr
	}

	pathParts := strings.Split(dir, string(os.PathSeparator))

	// windows absolute path workaround
	// issue describing: https://github.com/golang/go/issues/26953#issuecomment-412447719
	if strings.HasSuffix(pathParts[0], ":") {
		pathParts[0] = pathParts[0] + "/"
	}

	// os.Getwd() returns current test folder.
	// trying to walk up the absolute path to find testplan.json
	tmpPathParts := pathParts
	for range pathParts {
		basicPath := filepath.Join(tmpPathParts...)
		absolutePath := filepath.Join(basicPath, filepath.Clean(path))
		if pathParts[0] == "" && len(pathParts) > 1 {
			absolutePath = "/" + absolutePath
		}

		//nolint:gosec // already cleared
		testPlanRaw, readFileErr = ioutil.ReadFile(absolutePath)
		if readFileErr != nil {
			// stop looking if project root found
			//nolint:gosec // already cleared
			_, gErr := ioutil.ReadFile(filepath.Join(basicPath, "go.mod"))
			if gErr == nil {
				return
			}
			tmpPathParts = tmpPathParts[:len(tmpPathParts)-1]
			continue
		}
		break
	}
	return
}
