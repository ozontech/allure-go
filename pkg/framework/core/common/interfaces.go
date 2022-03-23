package common

import (
	"sync"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type InternalT interface {
	provider.T

	GetProvider() provider.Provider
	WG() *sync.WaitGroup
	GetResult() *allure.Result
}
