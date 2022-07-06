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

type StepAsyncDemo struct {
	suite.Suite
}

func (s *StepAsyncDemo) BeforeEach(t provider.T) {
	t.Epic("Async")
	t.Feature("Async Steps")
	t.Tags("async", "suite", "steps")
}

func (s *StepAsyncDemo) TestAsyncStepDemo1(t provider.T) {
	t.Title("Test with async steps 1")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.Logf("Step 2 Stopped At: %s", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo2(t provider.T) {
	t.Title("Test with async steps 2")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo3(t provider.T) {
	t.Title("Test with async steps 3")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		t.Title("Test with async steps 1")
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.Logf("Step 2 Stopped At: %s", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo4(t provider.T) {
	t.Title("Test with async steps 4")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo5(t provider.T) {
	t.Title("Test with async steps 5")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		panic("Whoops")
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo6(t provider.T) {
	t.Title("Test with async steps 6")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			panic("Whoops")
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo7(t provider.T) {
	t.Title("Test with async steps 7")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		ctx.Assert().False(true)
	})
}

func (s *StepAsyncDemo) TestAsyncStepDemo8(t provider.T) {
	t.Title("Test with async steps 8")

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			ctx.Assert().False(true)
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

type AsyncSuiteStepDemo struct {
	suite.Suite
}

func (s *AsyncSuiteStepDemo) BeforeEach(t provider.T) {
	t.Epic("Async")
	t.Feature("Async Steps")
	t.Tags("async", "suite", "steps")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo1(t provider.T) {
	t.Title("Async Test with async steps 1")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.Logf("Step 2 Stopped At: %s", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo2(t provider.T) {
	t.Title("Async Test with async steps 2")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo3(t provider.T) {
	t.Title("Async Test with async steps 3")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.Logf("Step 2 Stopped At: %s", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo4(t provider.T) {
	t.Title("Async Test with async steps 4")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo5(t provider.T) {
	t.Title("Async Test with async steps 5")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		panic("Whoops")
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo6(t provider.T) {
	t.Title("Async Test with async steps 6")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))

		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			panic("Whoops")
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo7(t provider.T) {
	t.Title("Async Test with async steps 7")

	t.Parallel()
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		time.Sleep(3 * time.Second)
		defer func() {
			ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
		}()
		ctx.Assert().False(true)
	})
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo8(t provider.T) {
	t.Title("Async Test with async steps 8")

	t.Parallel()

	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			ctx.WithNewParameters("Start", fmt.Sprintf("%s", time.Now()))
			time.Sleep(3 * time.Second)
			defer func() {
				ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
			}()
			ctx.Assert().False(true)
		})
		time.Sleep(3 * time.Second)
		ctx.WithNewParameters("Stop", fmt.Sprintf("%s", time.Now()))
	})
}

func TestStepAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(StepAsyncDemo))
}

func TestSuiteStepAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(AsyncSuiteStepDemo))
}
