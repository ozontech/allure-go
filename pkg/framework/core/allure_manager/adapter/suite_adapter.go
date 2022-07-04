package adapter

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// SuiteAdapter describes behavior of the suite
// such as before/after all functions, package name, runner name, suite path and suite name
type SuiteAdapter struct {
	packageName   string
	runner        string
	fullSuiteName string
	suiteName     string
	parentSuite   string

	beforeAll func(provider.T)
	afterAll  func(provider.T)

	container *allure.Container
}

// NewSuiteMeta returns SuiteAdapter pointer
func NewSuiteMeta(packageName, runner, fullSuiteName, suiteName string) *SuiteAdapter {
	return &SuiteAdapter{
		packageName:   packageName,
		runner:        runner,
		fullSuiteName: fullSuiteName,
		suiteName:     suiteName,
		container:     allure.NewContainer(),
	}
}

func NewSuiteMetaWithParent(packageName, runner, fullSuiteName, suiteName, parentSuite string) *SuiteAdapter {
	return &SuiteAdapter{
		packageName:   packageName,
		runner:        runner,
		fullSuiteName: fullSuiteName,
		suiteName:     suiteName,
		parentSuite:   parentSuite,
		container:     allure.NewContainer(),
	}
}

// GetPackageName returns package name
func (ctx *SuiteAdapter) GetPackageName() string {
	return ctx.packageName
}

// GetRunner returns runner name
func (ctx *SuiteAdapter) GetRunner() string {
	return ctx.runner
}

// GetSuiteName returns suite name
func (ctx *SuiteAdapter) GetSuiteName() string {
	return ctx.suiteName
}

// GetParentSuite returns suite name
func (ctx *SuiteAdapter) GetParentSuite() string {
	return ctx.parentSuite
}

// GetSuiteFullName returns full name
func (ctx *SuiteAdapter) GetSuiteFullName() string {
	return ctx.fullSuiteName
}

// GetContainer returns container
func (ctx *SuiteAdapter) GetContainer() *allure.Container {
	return ctx.container
}

// SetBeforeAll sets before all func
func (ctx *SuiteAdapter) SetBeforeAll(hook func(provider.T)) {
	ctx.beforeAll = hook
}

// SetAfterAll sets after all func
func (ctx *SuiteAdapter) SetAfterAll(hook func(provider.T)) {
	ctx.afterAll = hook
}

// GetBeforeAll returns before all func
func (ctx *SuiteAdapter) GetBeforeAll() func(provider.T) {
	return ctx.beforeAll
}

// GetAfterAll returns after all func
func (ctx *SuiteAdapter) GetAfterAll() func(provider.T) {
	return ctx.afterAll
}
