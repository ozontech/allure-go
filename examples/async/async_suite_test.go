//go:build async
// +build async

package async

import (
	"fmt"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type SuiteAsyncDemo struct {
	suite.Suite
}

func (s *SuiteAsyncDemo) BeforeEach(t provider.T) {
	t.Epic("Async")
	t.Feature("Async Suite")
	t.Tags("async", "suite", "steps")
}

func (s *SuiteAsyncDemo) TestAsyncSuiteDemo1(t provider.T) {
	t.Title("Async Test 1")

	startSign := fmt.Sprintf("%s", time.Now())
	t.Parallel()
	t.WithNewStep("Sync Step Demo", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Test 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Test 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
}

func (s *SuiteAsyncDemo) TestAsyncSuiteDemo2(t provider.T) {
	t.Title("Async Test 2")

	startSign := fmt.Sprintf("%s", time.Now())
	t.Parallel()
	t.WithNewStep("Sync Step Demo", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Test 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Test 2 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
}

func (s *SuiteAsyncDemo) TestAsyncSuiteDemo3(t provider.T) {
	t.Title("Async Test 3")
	t.Description(`
		This test should be failed. But all logs a correct.`)

	startSign := fmt.Sprintf("%s", time.Now())
	t.Parallel()
	t.WithNewStep("Sync Step Demo", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Test 2 Started At: %s", startSign)
		time.Sleep(1 * time.Second)
		defer func() {
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Test 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		}()
		ctx.Require().False(true)
	})
}

func (s *SuiteAsyncDemo) TestAsyncSuiteDemo4(t provider.T) {
	t.Title("Async Test 4")
	t.Description(`
		This test should be panic. But all logs a correct.`)

	startSign := fmt.Sprintf("%s", time.Now())
	t.Parallel()
	t.WithNewStep("Sync Step Demo", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Test 2 Started At: %s", startSign)
		time.Sleep(1 * time.Second)
		defer func() {
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Test 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		}()
		panic("Whoops")
	})
}

func TestSuiteAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(SuiteAsyncDemo))
}
