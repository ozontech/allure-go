package runner

import (
	"fmt"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"reflect"
)

type suiteResult struct {
	Container   *allure.Container
	TestResults []TestResult
}

func NewSuiteResult(container *allure.Container) SuiteResult {
	return &suiteResult{Container: container}
}

func (sr *suiteResult) NewResult(result TestResult) {
	sr.TestResults = append(sr.TestResults, result)
}

func (sr *suiteResult) GetContainer() *allure.Container {
	return sr.Container
}

func (sr *suiteResult) GetAllTestResults() []TestResult {
	return sr.TestResults
}

func (sr *suiteResult) GetResultByName(name string) TestResult {
	for _, tr := range sr.TestResults {
		if result := tr.GetResult(); result != nil {
			if result.Name == name {
				return tr
			}
		}
	}
	return nil
}

func (sr *suiteResult) GetResultByUUID(uuid string) TestResult {
	for _, tr := range sr.TestResults {
		if result := tr.GetResult(); result != nil {
			if result.UUID.String() == uuid {
				return tr
			}
		}
	}
	return nil
}

type testResult struct {
	result    *allure.Result
	container *allure.Container
}

func NewTestResult(result *allure.Result, container *allure.Container) TestResult {
	return &testResult{
		result:    result,
		container: container,
	}
}

func (tr *testResult) GetResult() *allure.Result {
	return tr.result
}

func (tr *testResult) GetContainer() *allure.Container {
	return tr.container
}

func (tr *testResult) Print() error {
	const errMessage = "failed to print Result. Reason: %s\nAlso failed to print Container. Reason: %s"
	var (
		result    *allure.Result
		container *allure.Container

		errR error
		errC error
	)
	if result = tr.GetResult(); result != nil {
		errR = result.Done()
	}
	if result == nil {
		errR = fmt.Errorf("failed to print Result. Reason: *allure.Result is nil")
	}

	if container = tr.GetContainer(); container != nil {
		errC = container.Done()
	}
	if container == nil {
		errC = fmt.Errorf("failed to print Container. Reason: *allure.Container is nil")
	}
	if errR != nil && errC != nil {
		return fmt.Errorf(errMessage, errR.Error(), errC.Error())
	}
	if errR != nil {
		return errR
	}
	if errC != nil {
		return errC
	}
	return nil
}

type TestBody func(t provider.T)

type testMethod struct {
	testMeta          provider.TestMeta
	testBody          reflect.Method
	callArgs          []reflect.Value
	getAdditionalArgs func() []interface{}
}

func (t *testMethod) GetArgs() []reflect.Value {
	return t.callArgs
}

func (t *testMethod) GetRawBody() reflect.Method {
	return t.testBody
}

func (t *testMethod) GetBody() TestBody {
	return func(pT provider.T) {
		t.testBody.Func.Call(insert(t.callArgs, 1, reflect.ValueOf(pT)))
	}
}

func (t *testMethod) GetMeta() provider.TestMeta {
	return t.testMeta
}

type testFunc struct {
	testBody TestBody
	testMeta provider.TestMeta
	callArgs []reflect.Value
}

func (t *testFunc) GetBody() TestBody {
	return t.testBody
}

func (t *testFunc) GetMeta() provider.TestMeta {
	return t.testMeta
}

func newTestFunc(body TestBody, testMeta provider.TestMeta) *testFunc {
	return &testFunc{
		testBody: body,
		testMeta: testMeta,
	}
}

func newTestMethod(method reflect.Method, testMeta provider.TestMeta, args []reflect.Value) Test {
	return &testMethod{
		testMeta: testMeta,
		testBody: method,
		callArgs: args,
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
