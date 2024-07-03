package runner

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// AllureBeforeTest has a BeforeEach method, which will run before each
// test in the suite.
type AllureBeforeTest interface {
	BeforeEach(t provider.T)
}

// AllureAfterTest has a AfterEach method, which will run after
// each test in the suite.
type AllureAfterTest interface {
	AfterEach(t provider.T)
}

// AllureBeforeSuite has a BeforeAll method, which will run before the
// tests in the suite are run.
type AllureBeforeSuite interface {
	BeforeAll(t provider.T)
}

// AllureAfterSuite has a AfterAll method, which will run after
// all the tests in the suite have been run.
type AllureAfterSuite interface {
	AfterAll(t provider.T)
}

// AllureIDSuite has a GetAllureID method,
// which will produce allureIDs for the test by its name
type AllureIDSuite interface {
	GetAllureID(testName string) string
}

// ParametrizedSuite suit can initialize parameters for
// parametrized test before running hooks
type ParametrizedSuite interface {
	InitializeTestsParams()
}

// ParametrizedTestParam parameter for parametrized test
// with custom AllureId and Title
type ParametrizedTestParam interface {
	GetAllureID() string
	GetAllureTitle() string
}

type TestSuite interface {
	GetRunner() TestRunner
	SetRunner(runner TestRunner)
}

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type TestRunner interface {
	NewTest(testName string, testBody func(provider.T), tags ...string)
	BeforeEach(hookBody func(provider.T))
	AfterEach(hookBody func(provider.T))
	BeforeAll(hookBody func(provider.T))
	AfterAll(hookBody func(provider.T))
	RunTests() SuiteResult
}

type Test interface {
	GetBody() TestBody
	GetMeta() provider.TestMeta
}

type SuiteResult interface {
	NewResult(result TestResult)
	GetContainer() *allure.Container
	GetAllTestResults() []TestResult
	GetResultByName(name string) TestResult
	GetResultByUUID(uuid string) TestResult
	ToJSON() ([]byte, error)
}

type TestResult interface {
	GetResult() *allure.Result
	GetContainer() *allure.Container
	Print() error
	ToJSON() ([]byte, error)
}

type internalT interface {
	provider.T

	SetRealT(t provider.TestingT)
	GetProvider() provider.Provider
	WG() *sync.WaitGroup
	GetResult() *allure.Result
}
