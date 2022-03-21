package provider

import (
	"sync"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
)

type TestingT interface {
	testing.TB
	Parallel()
	Run(testName string, testBody func(t *testing.T)) bool
}

type T interface {
	testing.TB
	AllureForward

	Parallel()

	RealT() TestingT
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
