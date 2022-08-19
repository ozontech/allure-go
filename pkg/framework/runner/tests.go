package runner

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type suiteResult struct {
	container   *allure.Container
	testResults []TestResult
}

// NewSuiteResult Returns new SuiteResult
func NewSuiteResult(container *allure.Container) SuiteResult {
	return &suiteResult{container: container}
}

// NewResult appends test result to suite result
func (sr *suiteResult) NewResult(result TestResult) {
	sr.testResults = append(sr.testResults, result)
}

// GetContainer returns parent container
func (sr *suiteResult) GetContainer() *allure.Container {
	return sr.container
}

// GetAllTestResults returns all test results of suite
func (sr *suiteResult) GetAllTestResults() []TestResult {
	return sr.testResults
}

// GetResultByName searches result by name and returns it
func (sr *suiteResult) GetResultByName(name string) TestResult {
	for _, tr := range sr.testResults {
		if result := tr.GetResult(); result != nil {
			if result.Name == name {
				return tr
			}
		}
	}
	return nil
}

// GetResultByUUID searches result by UUID and returns it
func (sr *suiteResult) GetResultByUUID(uuid string) TestResult {
	for _, tr := range sr.testResults {
		if result := tr.GetResult(); result != nil {
			if result.UUID.String() == uuid {
				return tr
			}
		}
	}
	return nil
}

// ToJson marshall result to Json object
func (sr *suiteResult) ToJson() ([]byte, error) {
	return json.Marshal(sr)
}

type testResult struct {
	result    *allure.Result
	container *allure.Container
}

// NewTestResult returns new test result
func NewTestResult(result *allure.Result, container *allure.Container) TestResult {
	return &testResult{
		result:    result,
		container: container,
	}
}

// GetResult returns result
func (tr *testResult) GetResult() *allure.Result {
	return tr.result
}

// GetContainer returns container
func (tr *testResult) GetContainer() *allure.Container {
	return tr.container
}

// Print returns print
func (tr *testResult) Print() error {
	const errMessage = "failed to print Result. Reason: %s\nAlso failed to print Container. Reason: %s"
	var (
		result    *allure.Result
		container *allure.Container

		resultErr    error
		containerErr error
	)
	if result = tr.GetResult(); result != nil {
		resultErr = result.Done()
	}
	if result == nil {
		resultErr = fmt.Errorf("failed to print Result. Reason: *allure.Result is nil")
	}

	if container = tr.GetContainer(); container != nil {
		containerErr = container.Done()
	}
	if container == nil {
		containerErr = fmt.Errorf("failed to print Container. Reason: *allure.Container is nil")
	}
	if resultErr != nil && containerErr != nil {
		return fmt.Errorf(errMessage, resultErr.Error(), containerErr.Error())
	}
	if resultErr != nil {
		return resultErr
	}
	if containerErr != nil {
		return containerErr
	}
	return nil
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
