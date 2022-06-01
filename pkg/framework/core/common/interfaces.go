package common

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"sync"
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
