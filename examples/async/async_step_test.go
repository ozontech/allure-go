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

func (s *StepAsyncDemo) TestAsyncStepDemo1(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 2 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo2(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo5(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 2 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo6(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo3(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		defer func() {
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		}()
		panic("Whoops")
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo4(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			defer func() {
				stopSign := fmt.Sprintf("%s", time.Now())
				ctx.Logf("Step 2 Stopped At: %s", stopSign)
				ctx.WithNewParameters("Stop", stopSign)
			}()
			panic("Whoops")
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo7(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		defer func() {
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		}()
		ctx.Assert().False(true)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *StepAsyncDemo) TestAsyncStepDemo8(t provider.T) {
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			defer func() {
				stopSign := fmt.Sprintf("%s", time.Now())
				ctx.Logf("Step 2 Stopped At: %s", stopSign)
				ctx.WithNewParameters("Stop", stopSign)
			}()
			ctx.Assert().False(true)
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

type AsyncSuiteStepDemo struct {
	suite.Suite
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo1(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 2 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo2(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo5(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 2 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo6(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo3(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		defer func() {
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		}()
		panic("Whoops")
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo4(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			defer func() {
				stopSign := fmt.Sprintf("%s", time.Now())
				ctx.Logf("Step 2 Stopped At: %s", stopSign)
				ctx.WithNewParameters("Stop", stopSign)
			}()
			panic("Whoops")
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo7(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 1 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})

	t.WithNewAsyncStep("Async Step 2", func(ctx provider.StepCtx) {
		startSign := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", startSign)
		ctx.Logf("Step 2 Started At: %s", startSign)
		time.Sleep(3 * time.Second)
		defer func() {
			stopSign := fmt.Sprintf("%s", time.Now())
			ctx.Logf("Step 2 Stopped At: %s", stopSign)
			ctx.WithNewParameters("Stop", stopSign)
		}()
		ctx.Assert().False(true)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func (s *AsyncSuiteStepDemo) TestAsyncStepDemo8(t provider.T) {
	t.Parallel()
	testStart := fmt.Sprintf("%s", time.Now())
	t.Logf("Test Started at %s", testStart)
	t.WithNewAsyncStep("Async Step 1", func(ctx provider.StepCtx) {
		outerStepStart := fmt.Sprintf("%s", time.Now())
		ctx.WithNewParameters("Start", outerStepStart)
		ctx.Logf("Step 1 Started At: %s", outerStepStart)
		ctx.WithNewAsyncStep("Async Step 1.1", func(ctx provider.StepCtx) {
			innerStepStart := fmt.Sprintf("%s", time.Now())
			ctx.WithNewParameters("Start", innerStepStart)
			ctx.Logf("Step 2 Started At: %s", innerStepStart)
			time.Sleep(3 * time.Second)
			defer func() {
				stopSign := fmt.Sprintf("%s", time.Now())
				ctx.Logf("Step 2 Stopped At: %s", stopSign)
				ctx.WithNewParameters("Stop", stopSign)
			}()
			ctx.Assert().False(true)
		})
		time.Sleep(3 * time.Second)
		stopSign := fmt.Sprintf("%s", time.Now())
		ctx.Logf("Step 1 Stopped At: %s", stopSign)
		ctx.WithNewParameters("Stop", stopSign)
	})
	t.Logf("Test already here at %s.", time.Now())
	t.Logf("But it still running cause of async steps!")
}

func TestStepAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(StepAsyncDemo))
}

func TestSuiteStepAsyncRunner(t *testing.T) {
	suite.RunSuite(t, new(AsyncSuiteStepDemo))
}
