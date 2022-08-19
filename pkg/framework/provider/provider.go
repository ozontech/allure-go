package provider

import (
	"github.com/ozontech/allure-go/pkg/allure"
)

type Provider interface {
	AllureForwardFull

	GetResult() *allure.Result
	UpdateResultStatus(msg string, trace string)
	StopResult(status allure.Status)

	SetTestMeta(meta TestMeta)
	GetTestMeta() TestMeta
	GetSuiteMeta() SuiteMeta
	ExecutionContext() ExecutionContext

	TestContext()
	BeforeEachContext()
	AfterEachContext()
	BeforeAllContext()
	AfterAllContext()

	NewTest(testName, packageName string, tags ...string)
	FinishTest() error
}

type TestMeta interface {
	GetResult() *allure.Result
	SetResult(result *allure.Result)

	GetContainer() *allure.Container

	SetBeforeEach(hook func(T))
	GetBeforeEach() func(T)
	SetAfterEach(hook func(T))
	GetAfterEach() func(T)
}

type SuiteMeta interface {
	GetPackageName() string
	GetRunner() string
	GetSuiteName() string
	GetParentSuite() string
	GetSuiteFullName() string
	GetContainer() *allure.Container

	SetBeforeAll(func(T))
	SetAfterAll(func(T))
	GetBeforeAll() func(T)
	GetAfterAll() func(T)
}

type ExecutionContext interface {
	AddStep(step *allure.Step)
	AddAttachments(attachment ...*allure.Attachment)
	GetName() string
}
