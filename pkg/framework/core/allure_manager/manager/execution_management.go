package manager

import "github.com/ozontech/allure-go/pkg/framework/core/allure_manager/ctx"

// TestContext initiate test context
func (a *allureManager) TestContext() {
	a.executionContext = ctx.NewTestCtx(a.testMeta.GetResult())
}

// BeforeEachContext initiate before each context
func (a *allureManager) BeforeEachContext() {
	a.executionContext = ctx.NewBeforeEachCtx(a.testMeta.GetContainer())
}

// AfterEachContext initiate after each context
func (a *allureManager) AfterEachContext() {
	a.executionContext = ctx.NewAfterEachCtx(a.testMeta.GetContainer())
}

// BeforeAllContext initiate before all context
func (a *allureManager) BeforeAllContext() {
	a.executionContext = ctx.NewBeforeAllCtx(a.suiteMeta.GetContainer())
}

// AfterAllContext initiate after all context
func (a *allureManager) AfterAllContext() {
	a.executionContext = ctx.NewAfterAllCtx(a.suiteMeta.GetContainer())
}
