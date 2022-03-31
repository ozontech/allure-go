//go:build allure_go_new
// +build allure_go_new

package allure_go_compare

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type AllureGoBeforesAfters struct {
	suite.Suite
}

/* Allure-Go style:
func TestAllureSetupTeardown(t *testing.T) {
	allure.BeforeTest(t,
		allure.Description("setup"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Setup step 1"),
				allure.Action(func() {}))
		}))

	allure.Test(t,
		allure.Description("actual test"),
		allure.Action(func() {
			allure.Step(
				allure.Description("Test step 1"),
				allure.Action(func() {}))
		}))

	allure.AfterTest(t,
		allure.Description("teardown"),
		allure.Action(func() {
			allure.Step(
				allure.Description("teardown step 1"),
				allure.Action(func() {}))
		}))
}
*/

// We provide even more! but without infinite breakers ({(((({(({()}))}))))})
//func (s AllureGoBeforesAfters) BeforeAll() {
//	s.NewStep("Setup suite step 1")
//}
//
//func (s AllureGoBeforesAfters) AfterAll() {
//	s.NewStep("Teardown suite step 1")
//}

func (s AllureGoBeforesAfters) BeforeEach(t provider.T) {
	t.NewStep("Setup test step 1")
}

func (s AllureGoBeforesAfters) AfterEach(t provider.T) {
	t.NewStep("Teardown test step 1")
}

func (s AllureGoBeforesAfters) TestBeforesAfters(t provider.T) {
	t.Epic("Compare with allure-go")
	t.NewStep("Test step 1")
}

func TestAllureGoBeforesAfters(t *testing.T) {
	suite.RunSuite(t, new(AllureGoBeforesAfters))
}
