//go:build provider_new
// +build provider_new

package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestPanicProviderDemo(realT *testing.T) {
	r := runner.NewRunner(realT, realT.Name())
	r.BeforeEach(func(t provider.T) {
		t.Epic("Only Provider Demo")
		t.Feature("runner.RunTest")
	})
	r.NewTest("Test1", func(t provider.T) {
		t.Title("Some Panic test 1")
		t.Description("allure-go allows you to use allure without suites. Even if it panics or failed")
		panic("whoops")
	})

	r.NewTest("Test3", func(t provider.T) {
		t.Title("Some Failed test 2")
		t.Description("allure-go allows you to use allure without suites. Even if it panics or failed")
		t.Require().NotNil(nil)
	})

	r.NewTest("Test2", func(t provider.T) {
		t.Title("Some Normal test 2")
		t.Description("allure-go allows you to use allure without suites. Even if it panics or failed")
	})

	r.RunTests()
}
