package manager

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/core/allure_manager/adapter"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type allureManager struct {
	testMeta  provider.TestMeta
	suiteMeta provider.SuiteMeta

	executionContext provider.ExecutionContext
}

func NewProvider(cfg ProviderConfig) provider.Provider {
	suiteMeta := adapter.NewSuiteMetaWithParent(
		cfg.PackageName(),
		cfg.Runner(),
		cfg.FullName(),
		cfg.SuiteName(),
		cfg.ParentSuite(),
	)

	return &allureManager{
		suiteMeta: suiteMeta,
		testMeta:  adapter.NewTestMeta(cfg.FullName(), cfg.SuiteName(), "", cfg.PackageName()),
	}
}

func (a *allureManager) withResult(f func(r *allure.Result)) {
	if r := a.GetResult(); r != nil {
		f(r)
	}
}

func (a *allureManager) UpdateResultStatus(msg string, trace string) {
	a.GetResult().SetStatusMessage(msg)
	a.GetResult().SetStatusTrace(trace)
}

func (a *allureManager) StopResult(status allure.Status) {
	a.withResult(func(r *allure.Result) {
		r.Status = status
		r.Stop = allure.GetNow()
	})
}

func (a *allureManager) GetResult() *allure.Result {
	return a.testMeta.GetResult()
}

func (a *allureManager) SetTestMeta(meta provider.TestMeta) {
	a.testMeta = meta
}

func (a *allureManager) GetTestMeta() provider.TestMeta {
	return a.testMeta
}

func (a *allureManager) GetSuiteMeta() provider.SuiteMeta {
	return a.suiteMeta
}

func (a *allureManager) ExecutionContext() provider.ExecutionContext {
	return a.executionContext
}

func (a *allureManager) NewTest(testName, packageName string, tags ...string) {
	a.testMeta = adapter.NewTestMeta(
		a.suiteMeta.GetSuiteFullName(),
		a.suiteMeta.GetSuiteName(),
		testName,
		packageName,
		tags...,
	)

	if parentSuite := a.suiteMeta.GetParentSuite(); parentSuite != "" {
		a.testMeta.GetResult().WithParentSuite(parentSuite)
	}

	a.suiteMeta.GetContainer().AddChild(a.testMeta.GetResult().UUID)
}

func (a *allureManager) FinishTest() error {
	container := a.testMeta.GetContainer()
	if container != nil {
		if err := container.Done(); err != nil {
			return err
		}
	}

	return a.testMeta.GetResult().Done()
}
