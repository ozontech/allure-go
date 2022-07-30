//go:build examples_new
// +build examples_new

package suite_demo

import (
	"fmt"
	"log"
	"os"
	"testing"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type StepTreeDemoSuite struct {
	suite.Suite
}

func (s *StepTreeDemoSuite) TestInnerSteps(t provider.T) {
	t.Epic("Demo")
	t.Feature("Inner Steps")
	t.Title("Simple Nesting")
	t.Description(`
		Step A is parent step for Step B and Step C
		Call order will be saved in allure report
		A -> (B, C)`)

	t.Tags("Steps", "Nesting")

	t.WithNewStep("Step A", func(ctx provider.StepCtx) {
		ctx.NewStep("Step B")
		ctx.NewStep("Step C")
	})
}

func (s *StepTreeDemoSuite) TestComplexStepTree(t provider.T) {
	t.Epic("Demo")
	t.Feature("Inner Steps")
	t.Title("Complex Nesting")
	t.Description(`
		Step A is parent for Step B, Step C and Step F
		Step C is parent for Step D and Step E
		Step F is parent for Step G and Step H
		Call order will be saved in allure report
		A -> (B, C -> (D, E), F -> (G, H), I)`)

	t.Tags("Steps", "Nesting")

	t.WithNewStep("Step A", func(ctx provider.StepCtx) {
		ctx.NewStep("Step B")
		ctx.WithNewStep("Step C", func(ctx provider.StepCtx) {
			ctx.NewStep("Step D")
			ctx.NewStep("Step E")
		})
		ctx.WithNewStep("Step F", func(ctx provider.StepCtx) {
			ctx.NewStep("Step G")
			ctx.NewStep("Step H")
		})
		ctx.NewStep("Step I")
	})
}

func TestStepTree(t *testing.T) {
	res := suite.RunSuite(t, new(StepTreeDemoSuite))
	processResult(res)
}

func processResult(results map[string]*allure.Result) {
	token := os.Getenv("TG_TOKEN")
	channel := os.Getenv("TG_CHANNEL")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for name, result := range results {
		text := fmt.Sprintf("*%v*\n\n*Status:* %v\n\n*Description:* %v", name, result.Status, result.Description)
		msg := tgbotapi.NewMessageToChannel(channel, text)
		msg.ParseMode = "Markdown"
		_, err = bot.Send(msg)
		if err != nil {
			log.Panic(err)
		}
	}
}
