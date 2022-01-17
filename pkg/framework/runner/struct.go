package runner

import (
	"strings"
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/internal/file_manager"
	"github.com/ozontech/allure-go/pkg/provider"
)

// ITestRunner ...
type ITestRunner interface {
	WithBeforeEach(beforeEach func(t *provider.T))
	WithAfterEach(afterEach func(t *provider.T))
	Run(testName string, test func(t *provider.T), tags ...string) bool
}

// TestRunner - a structure that allows to run a number of tests using Before/After Test hooks.
// It contains pointer to test context `provider.T`. Supports parallel launching of tests.
type TestRunner struct {
	t *provider.T
}

// NewTestRunner Constructor. Creates a new instance of `TestRunner` by initializing `provider.T` based on passed realT
// Returns pointer to the created instance of `TestRunner`.
func NewTestRunner(realT *testing.T) *TestRunner {
	callers := strings.Split(realT.Name(), "/")
	return &TestRunner{provider.NewTWithPackage(realT, realT.Name(), callers[len(callers)-1], file_manager.GetPackage(2))}
}

// WithBeforeEach - Sets the function to be executed before each test for a particular instance of `TestRunner`.
// If function is not specified, beforeEach will be skipped during test execution.
func (runner *TestRunner) WithBeforeEach(beforeEach func(t *provider.T)) {
	runner.t.WithBeforeTest(beforeEach)
}

// WithAfterEach - Sets the function to be executed after each test for a particular instance of `TestRunner`.
// If function is not specified, afterEach will be skipped during test execution.
func (runner *TestRunner) WithAfterEach(afterEach func(t *provider.T)) {
	runner.t.WithAfterTest(afterEach)
}

// Run Works similarly to static function runner.RunTest
// except that the context `T` will belong to an instance of the `TestRunner` structure.
// If `BeforeEach` or `AfterEach` are specified for TestRunner they will be executed before/after the test respectively.
// Returns true if the test was executed successfully. False - in case of error or panic.
func (runner *TestRunner) Run(testName string, test func(t *provider.T), tags ...string) bool {
	return runner.t.Run(testName, test, tags...)
}
