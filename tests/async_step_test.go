package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestAsyncStep_race(t *testing.T) {
	t.Run("steps race via async steps", func(t *testing.T) {
		r := runner.NewRunner(t, "suite: async step race")
		r.NewTest("async step race", func(t provider.T) {
			for i := 0; i < 64; i++ {
				paramVal := i
				t.WithNewAsyncStep(fmt.Sprintf("step_%d", paramVal), func(sCtx provider.StepCtx) {
					sCtx.WithNewParameters("param", paramVal)
				})
			}

			// Wait for all async steps to be flushed by internal wait groups.
			time.Sleep(300 * time.Millisecond)
		})
		r.RunTests()
	})

	t.Run("nested steps race via async steps", func(t *testing.T) {
		r := runner.NewRunner(t, "suite: async step nested stepctx race")
		r.NewTest("async step nested stepctx race", func(t provider.T) {
			for i := 0; i < 64; i++ {
				idx := i
				t.WithNewAsyncStep(fmt.Sprintf("step_add_step_%d", idx), func(ctx provider.StepCtx) {
					ctx.WithNewStep(fmt.Sprintf("inner_%d", idx), func(inner provider.StepCtx) {
						inner.WithNewParameters("idx", idx)
					})
				})
			}

			time.Sleep(300 * time.Millisecond)
		})
		r.RunTests()
	})

	t.Run("attachments race via async steps", func(t *testing.T) {
		r := runner.NewRunner(t, "suite: async step attachment race")
		r.NewTest("async step attachment race", func(t provider.T) {
			for i := 0; i < 64; i++ {
				idx := i
				t.WithNewAsyncStep(fmt.Sprintf("step_attach_%d", idx), func(ctx provider.StepCtx) {
					t.WithNewAttachment(
						fmt.Sprintf("attach_%d", idx),
						"text/plain",
						[]byte(fmt.Sprintf("payload-%d", idx)),
					)
					ctx.WithNewAttachment(
						fmt.Sprintf("inner_attach_%d", idx),
						"text/plain",
						[]byte(fmt.Sprintf("inner-payload-%d", idx)),
					)
				})
			}

			time.Sleep(300 * time.Millisecond)
		})
		r.RunTests()
	})

	t.Run("parameters race via async steps", func(t *testing.T) {
		r := runner.NewRunner(t, "suite: async step parameter race")
		r.NewTest("async step parameter race", func(t provider.T) {
			for i := 0; i < 64; i++ {
				idx := i
				t.WithNewAsyncStep(fmt.Sprintf("step_param_%d", idx), func(ctx provider.StepCtx) {
					t.WithNewParameters("idx", idx)
					ctx.WithNewParameters("inner_idx", idx)
				})
			}

			time.Sleep(300 * time.Millisecond)
		})
		r.RunTests()
	})

	t.Run("labels race via async steps", func(t *testing.T) {
		r := runner.NewRunner(t, "suite: async step label race")
		r.NewTest("async step label race", func(t provider.T) {
			for i := 0; i < 64; i++ {
				idx := i
				t.WithNewAsyncStep(fmt.Sprintf("step_label_%d", idx), func(ctx provider.StepCtx) {
					t.Tag(fmt.Sprintf("tag_%d", idx))
					t.ReplaceLabel(allure.NewLabel(allure.Feature, fmt.Sprintf("feature_%d", idx)))
					ctx.WithNewParameters("label_idx", idx)
				})
			}

			time.Sleep(300 * time.Millisecond)
		})
		r.RunTests()
	})

	t.Run("links race via async steps", func(t *testing.T) {
		r := runner.NewRunner(t, "suite: async step link race")
		r.NewTest("async step link race", func(t provider.T) {
			for i := 0; i < 64; i++ {
				idx := i
				t.WithNewAsyncStep(fmt.Sprintf("step_link_%d", idx), func(ctx provider.StepCtx) {
					t.Link(allure.NewLink(
						fmt.Sprintf("link_%d", idx),
						allure.LINK,
						fmt.Sprintf("https://localhost/%d", idx),
					))
					ctx.WithNewParameters("link_idx", idx)
				})
			}

			time.Sleep(300 * time.Millisecond)
		})
		r.RunTests()
	})

}
