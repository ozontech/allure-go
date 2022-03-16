package common

import (
	"fmt"
	"os"
	"runtime"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type testContext struct {
	beforeEach func(provider.T)
	afterEach  func(provider.T)

	result    *allure.Result
	container *allure.Container
}

type TestContext interface {
	GetResult() *allure.Result
	SetResult(result *allure.Result)

	GetContainer() *allure.Container

	SetBeforeEach(hook func(provider.T))
	GetBeforeEach() func(provider.T)
	SetAfterEach(hook func(provider.T))
	GetAfterEach() func(provider.T)
}

func newTestContext(suiteFullName, suiteName, testName, packageName string, tags ...string) *testContext {
	host, _ := os.Hostname()
	fullName := fmt.Sprintf("%s/%s", suiteFullName, testName)

	var newTags []allure.Label
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

	return &testContext{result: result, container: container}
}

func (ctx *testContext) GetResult() *allure.Result {
	return ctx.result
}

func (ctx *testContext) SetResult(result *allure.Result) {
	ctx.result = result
}

func (ctx *testContext) GetContainer() *allure.Container {
	return ctx.container
}

func (ctx *testContext) SetBeforeEach(hook func(provider.T)) {
	ctx.beforeEach = hook
}
func (ctx *testContext) GetBeforeEach() func(provider.T) {
	return ctx.beforeEach
}
func (ctx *testContext) SetAfterEach(hook func(provider.T)) {
	ctx.afterEach = hook
}

func (ctx *testContext) GetAfterEach() func(provider.T) {
	return ctx.afterEach

}

type SuiteContext interface {
	GetPackageName() string
	GetRunner() string
	GetSuiteName() string
	GetSuiteFullName() string
	GetContainer() *allure.Container

	SetBeforeAll(func(provider.T))
	SetAfterAll(func(provider.T))
	GetBeforeAll() func(provider.T)
	GetAfterAll() func(provider.T)
}

type suiteContext struct {
	packageName   string
	runner        string
	fullSuiteName string
	suiteName     string

	beforeAll func(provider.T)
	afterAll  func(provider.T)

	container *allure.Container
}

func newSuiteContext(packageName, runner, fullSuiteName, suiteName string) *suiteContext {
	return &suiteContext{
		packageName:   packageName,
		runner:        runner,
		fullSuiteName: fullSuiteName,
		suiteName:     suiteName,
		container:     allure.NewContainer(),
	}
}

func (ctx *suiteContext) GetPackageName() string {
	return ctx.packageName
}

func (ctx *suiteContext) GetRunner() string {
	return ctx.runner
}

func (ctx *suiteContext) GetSuiteName() string {
	return ctx.suiteName
}

func (ctx *suiteContext) GetSuiteFullName() string {
	return ctx.fullSuiteName
}

func (ctx *suiteContext) GetContainer() *allure.Container {
	return ctx.container
}

func (ctx *suiteContext) SetBeforeAll(hook func(provider.T)) {
	ctx.beforeAll = hook
}

func (ctx *suiteContext) SetAfterAll(hook func(provider.T)) {
	ctx.afterAll = hook
}

func (ctx *suiteContext) GetBeforeAll() func(provider.T) {
	return ctx.beforeAll
}

func (ctx *suiteContext) GetAfterAll() func(provider.T) {
	return ctx.afterAll
}

type Provider interface {
	NewTest(testName, packageName string, tags ...string)
	FinishTest()
	TestContext()
	BeforeEachContext()
	AfterEachContext()
	BeforeAllContext()
	AfterAllContext()

	GetTestMeta() provider.TestContext
	GetSuiteMeta() provider.SuiteContext
	ExecutionContext() provider.ExecutionContext
}

type p struct {
	tCtx provider.TestContext
	sCtx provider.SuiteContext

	eCtx provider.ExecutionContext
}

func newProvider(packageName, runner, fullSuiteName, suiteName string) Provider {
	sCtx := newSuiteContext(packageName, runner, fullSuiteName, suiteName)
	return &p{sCtx: sCtx, tCtx: &testContext{}}
}

func (p *p) TestContext() {
	p.eCtx = newTestCtx(p.tCtx.GetResult())
}

func (p *p) BeforeEachContext() {
	p.eCtx = newBeforeEachCtx(p.tCtx.GetContainer())
}

func (p *p) AfterEachContext() {
	p.eCtx = newAfterEachCtx(p.tCtx.GetContainer())
}

func (p *p) BeforeAllContext() {
	p.eCtx = newBeforeAllCtx(p.sCtx.GetContainer())
}

func (p *p) AfterAllContext() {
	p.eCtx = newAfterAllCtx(p.sCtx.GetContainer())
}

func (p *p) NewTest(testName, packageName string, tags ...string) {
	p.tCtx = newTestContext(p.sCtx.GetSuiteFullName(), p.sCtx.GetSuiteName(), testName, packageName, tags...)
	p.sCtx.GetContainer().AddChild(p.tCtx.GetResult().UUID)
}

func (p *p) FinishTest() {
	p.tCtx.GetResult().Done()
}

func (p *p) GetTestMeta() provider.TestContext {
	return p.tCtx
}

func (p *p) GetSuiteMeta() provider.SuiteContext {
	return p.sCtx
}

func (p *p) ExecutionContext() provider.ExecutionContext {
	return p.eCtx
}
