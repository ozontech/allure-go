package provider

import "github.com/ozontech/allure-go/pkg/allure"

type Provider interface {
	NewTest(testName, packageName string, tags ...string)
	FinishTest()
	TestContext()
	BeforeEachContext()
	AfterEachContext()
	BeforeAllContext()
	AfterAllContext()

	GetTestMeta() TestContext
	GetSuiteMeta() SuiteContext
	ExecutionContext() ExecutionContext
}

type TestContext interface {
	GetResult() *allure.Result
	SetResult(result *allure.Result)

	GetContainer() *allure.Container

	SetBeforeEach(hook func(T))
	GetBeforeEach() func(T)
	SetAfterEach(hook func(T))
	GetAfterEach() func(T)
}

type SuiteContext interface {
	GetPackageName() string
	GetRunner() string
	GetSuiteName() string
	GetSuiteFullName() string
	GetContainer() *allure.Container

	SetBeforeAll(func(T))
	SetAfterAll(func(T))
	GetBeforeAll() func(T)
	GetAfterAll() func(T)
}

type ExecutionContext interface {
	AddStep(step *allure.Step)
	AddAttachment(attachment *allure.Attachment)
	GetName() string
}
