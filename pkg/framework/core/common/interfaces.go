package common

import (
	"sync"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type ParentT interface {
	GetProvider() provider.Provider
	GetResult() *allure.Result
}

type HookProvider interface {
	BeforeEachContext()
	AfterEachContext()
	BeforeAllContext()
	AfterAllContext()

	GetSuiteMeta() provider.SuiteMeta
	GetTestMeta() provider.TestMeta
}

type InternalT interface {
	provider.T
	SetRealT(realT provider.TestingT)
	WG() *sync.WaitGroup
}
