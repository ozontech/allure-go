package provider

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
)

type T interface {
	testing.TB
	AllureForward

	Parallel()

	RealT() *testing.T
	XSkip()
	SkipOnPrint()
	Assert() Asserts
	Require() Asserts
	Run(testName string, testBody func(T), tags ...string) bool
}

type InternalT interface {
	T

	Provider() Provider
	WG() *sync.WaitGroup
	GetResult() *allure.Result
	GetPackage() string
	BreakResult(string)
}
