package adapter

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// TestAdapter describes behavior of the test
// such as before/after each function, container, test result object, and container object
type TestAdapter struct {
	beforeEach func(provider.T)
	afterEach  func(provider.T)

	result    *allure.Result
	container *allure.Container
}

// NewTestMeta returns pointer to instance of TestAdapter
func NewTestMeta(suiteFullName, suiteName, testName, packageName string, tags ...string) *TestAdapter {
	host, _ := os.Hostname()
	fullName := suiteFullName
	// ex: suiteFullName=TestRunner/My_Test, testName=My Test => after split and replace: My_test == My_test
	// why? to avoid TestRunner/My_Test/My Test
	if callers := strings.Split(suiteFullName, "/"); callers[len(callers)-1] != strings.ReplaceAll(testName, " ", "_") {
		fullName = fmt.Sprintf("%s/%s", fullName, testName)
	}

	var newTags []*allure.Label
	for _, tag := range tags {
		newTags = append(newTags, allure.NewLabel(allure.Tag, tag))
	}

	result := allure.NewResult(testName, fullName).
		WithFrameWork(allure.DefaultVersion).
		WithHost(host).
		WithThread(fullName).
		WithLanguage(runtime.Version()).
		WithLaunchTags().
		WithSuite(suiteName).
		WithPackage(packageName).
		WithLabels(newTags...)

	container := allure.NewContainer()
	container.AddChild(result.UUID)

	return &TestAdapter{result: result, container: container}
}

// GetResult returns allure.Result pointer
func (ctx *TestAdapter) GetResult() *allure.Result {
	return ctx.result
}

// SetResult sets allure.Result
func (ctx *TestAdapter) SetResult(result *allure.Result) {
	ctx.result = result
}

// GetContainer returns allure.Container pointer
func (ctx *TestAdapter) GetContainer() *allure.Container {
	return ctx.container
}

// SetBeforeEach sets before each function
func (ctx *TestAdapter) SetBeforeEach(hook func(provider.T)) {
	ctx.beforeEach = hook
}

// GetBeforeEach returns before each function
func (ctx *TestAdapter) GetBeforeEach() func(provider.T) {
	return ctx.beforeEach
}

// SetAfterEach sets after each function
func (ctx *TestAdapter) SetAfterEach(hook func(provider.T)) {
	ctx.afterEach = hook
}

// GetAfterEach returns after each function
func (ctx *TestAdapter) GetAfterEach() func(provider.T) {
	return ctx.afterEach
}
