//go:build examples_new
// +build examples_new

package suite_demo

import (
	"context"
	"fmt"
	"testing"

	"github.com/jackc/fake"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type Example struct {
	country string
	number  int
}

func (e *Example) String() string {
	return fmt.Sprintf("%s#%d", e.country, e.number)
}

type SetupSuite struct {
	suite.Suite
	ParamMyTest []*Example
}

func (s *SetupSuite) BeforeAll(t provider.T) {
	var params []*allure.Parameter
	for i := 0; i < 10; i++ {
		param := &Example{
			country: fake.Country(),
			number:  fake.Year(1900, 2000),
		}
		params = append(params, allure.NewParameter(fmt.Sprintf("Ex %d", i), param))
		s.ParamMyTest = append(s.ParamMyTest, param)
	}
	t.NewStep("BeforeAllStep", params...)
}

func (s *SetupSuite) BeforeEach(t provider.T) {
	t.Epic("Demo")
	t.Feature("BeforeAfter")
	t.NewStep("This Step will be before Each")
}

func (s *SetupSuite) AfterEach(t provider.T) {
	t.NewStep("AfterEach Step")
}

func (s *SetupSuite) AfterAll(t provider.T) {
	t.NewStep("AfterAll Step")
}

func (s *SetupSuite) TableTestMyTest(t provider.T, example *Example) {
	t.Titlef("TableTest With Setup - %s", example)
	t.Descriptionf(`
		Test will unpack all data from passed parameter to the variables in WithTestSetup func.
		After test finish, it will do ctx.Done() in TestTearDown.
		All Setup and TearDown tests will be add as Befores and Afters to test's container.
		Used Data: %s`, example)
	t.Tags("Parametrized", "Parallel", "Setup", "BeforeAfter")

	t.Parallel()
	var (
		country string
		year    int
		ctx     context.Context
	)

	defer t.WithTestTeardown(func(t provider.T) {
		t.WithNewStep("Close ctx", func(sCtx provider.StepCtx) {
			ctx.Done()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})

	t.WithTestSetup(func(t provider.T) {
		t.WithNewStep("init country", func(sCtx provider.StepCtx) {
			country = example.country
			sCtx.WithNewParameters("country", country)
		})
		t.WithNewStep("init year", func(sCtx provider.StepCtx) {
			year = example.number
			sCtx.WithNewParameters("year", year)
		})
		t.WithNewStep("init ctx", func(sCtx provider.StepCtx) {
			ctx = context.Background()
			sCtx.WithNewParameters("ctx", ctx)
		})
	})

	t.Require().NotEqual("PonyCountry", country, "No magic countries in the list")
	t.Require().NotEqual(2007, year, "No one returned to 2007")
	t.Require().NotNil(ctx, "Not empty context")
}

func (s *SetupSuite) TestMyOtherTest(t provider.T) {
	t.Title("Just Test WithSetup")
	t.Description(`
		Test will prepare some data at TestSetup.`)
	t.Tags("Parallel", "Setup", "BeforeAfter")

	t.Parallel()
	var (
		name string
		age  int
	)

	t.WithTestSetup(func(t provider.T) {
		t.WithNewStep("init args", func(sCtx provider.StepCtx) {
			name = fake.FullName()
			age = fake.Day()
			sCtx.WithNewParameters("name", name, "age", age)
		})
	})
}

func TestRunner(t *testing.T) {
	suite.RunSuite(t, new(SetupSuite))
}
