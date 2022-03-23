package adapter

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// SuiteContext describes behavior of the suite
// such as before/after all functions, package name, runner name, suite path and suite name
type SuiteContext struct {
	packageName   string
	runner        string
	fullSuiteName string
	suiteName     string

	beforeAll func(provider.T)
	afterAll  func(provider.T)

	container *allure.Container
}

// NewSuiteMeta returns SuiteContext pointer
func NewSuiteMeta(packageName, runner, fullSuiteName, suiteName string) *SuiteContext {
	return &SuiteContext{
		packageName:   packageName,
		runner:        runner,
		fullSuiteName: fullSuiteName,
		suiteName:     suiteName,
		container:     allure.NewContainer(),
	}
}

// GetPackageName returns package name
func (ctx *SuiteContext) GetPackageName() string {
	return ctx.packageName
}

// GetRunner returns runner name
func (ctx *SuiteContext) GetRunner() string {
	return ctx.runner
}

// GetSuiteName returns suite name
func (ctx *SuiteContext) GetSuiteName() string {
	return ctx.suiteName
}

// GetSuiteFullName returns full name
func (ctx *SuiteContext) GetSuiteFullName() string {
	return ctx.fullSuiteName
}

// GetContainer returns container
func (ctx *SuiteContext) GetContainer() *allure.Container {
	return ctx.container
}

// SetBeforeAll sets before all func
func (ctx *SuiteContext) SetBeforeAll(hook func(provider.T)) {
	ctx.beforeAll = hook
}

// SetAfterAll sets after all func
func (ctx *SuiteContext) SetAfterAll(hook func(provider.T)) {
	ctx.afterAll = hook
}

// GetBeforeAll returns before all func
func (ctx *SuiteContext) GetBeforeAll() func(provider.T) {
	return ctx.beforeAll
}

// GetAfterAll returns after all func
func (ctx *SuiteContext) GetAfterAll() func(provider.T) {
	return ctx.afterAll
}
