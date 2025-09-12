package runner

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/goccy/go-json"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type suiteResult struct {
	Container   *allure.Container `json:"container,omitempty"`
	TestResults []TestResult      `json:"test_results,omitempty"`

	mu sync.Mutex
}

// NewSuiteResult Returns new SuiteResult
func NewSuiteResult(container *allure.Container) SuiteResult {
	return &suiteResult{Container: container}
}

// NewResult appends test result to suite result
func (sr *suiteResult) NewResult(result TestResult) {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.TestResults = append(sr.TestResults, result)
}

// GetContainer returns parent Container
func (sr *suiteResult) GetContainer() *allure.Container {
	return sr.Container
}

// GetAllTestResults returns all test results of suite
func (sr *suiteResult) GetAllTestResults() []TestResult {
	return sr.TestResults
}

// GetResultByName searches result by name and returns it
func (sr *suiteResult) GetResultByName(name string) TestResult {
	for _, tr := range sr.TestResults {
		if r := tr.GetResult(); r != nil {
			if r.Name == name {
				return tr
			}
		}
	}

	return nil
}

// GetResultByUUID searches result by UUID and returns it
func (sr *suiteResult) GetResultByUUID(uuid string) TestResult {
	for _, tr := range sr.TestResults {
		if r := tr.GetResult(); r != nil {
			if r.UUID.String() == uuid {
				return tr
			}
		}
	}
	return nil
}

// ToJSON marshall result to Json object
//
// Deprecated: use [json.Marshal] instead.
func (sr *suiteResult) ToJSON() ([]byte, error) {
	return json.Marshal(sr)
}

type testResult struct {
	Result    *allure.Result    `json:"result,omitempty"`
	Container *allure.Container `json:"container,omitempty"`
}

// NewTestResult returns new test result
func NewTestResult(result *allure.Result, container *allure.Container) TestResult {
	return &testResult{
		Result:    result,
		Container: container,
	}
}

// GetResult returns result
func (tr *testResult) GetResult() *allure.Result {
	return tr.Result
}

// GetContainer returns Container
func (tr *testResult) GetContainer() *allure.Container {
	return tr.Container
}

// Print returns print
func (tr *testResult) Print() error {
	var (
		resultErr    error
		containerErr error
	)

	result := tr.GetResult()

	if result != nil {
		resultErr = result.Done()
	} else {
		resultErr = fmt.Errorf("failed to print Result. Reason: *allure.Result is nil")
	}

	container := tr.GetContainer()
	if container != nil {
		containerErr = container.Done()
	} else {
		containerErr = fmt.Errorf("failed to print Container. Reason: *allure.Container is nil")
	}

	if resultErr != nil && containerErr != nil {
		return fmt.Errorf("failed to print Result. Reason: %s\nAlso failed to print Container. Reason: %s", resultErr, containerErr)
	}

	if resultErr != nil {
		return resultErr
	}

	if containerErr != nil {
		return containerErr
	}

	return nil
}

// ToJSON marshall TestResult to the JSON
//
// Deprecated: use [json.Marshal] instead
func (tr *testResult) ToJSON() ([]byte, error) {
	return json.Marshal(tr)
}

type TestBody func(t provider.T)

type testMethod struct {
	testMeta provider.TestMeta
	testBody reflect.Method
	callArgs []reflect.Value
}

// GetArgs returns call args of the test
func (t *testMethod) GetArgs() []reflect.Value {
	return t.callArgs
}

// GetRawBody returns reflect.Method of the test
func (t *testMethod) GetRawBody() reflect.Method {
	return t.testBody
}

// GetBody returns wrapped function at the test
func (t *testMethod) GetBody() TestBody {
	return func(pT provider.T) {
		t.testBody.Func.Call(insert(t.callArgs, 1, reflect.ValueOf(pT)))
	}
}

// GetMeta returns provider.TestMeta of the test
func (t *testMethod) GetMeta() provider.TestMeta {
	return t.testMeta
}

type testFunc struct {
	testBody TestBody
	testMeta provider.TestMeta
}

// GetBody returns test function
func (t *testFunc) GetBody() TestBody {
	return t.testBody
}

// GetMeta returns provider.TestMeta of the test
func (t *testFunc) GetMeta() provider.TestMeta {
	return t.testMeta
}

func newTestFunc(body TestBody, testMeta provider.TestMeta) *testFunc {
	return &testFunc{
		testBody: body,
		testMeta: testMeta,
	}
}

func insert(a []reflect.Value, index int, value reflect.Value) []reflect.Value {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}

	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value

	return a
}
