# Allure-Go

![](.resources/allure_go_icon.svg)

Allure-Go - the project that provides a complete allure provider in go, without overloading the interface of usage. <br>
The project started as a fork of testify, but over time it got its own runner and its own features. <br>

## Head of contents

## Features

Providing a separate package allows you to customize your work with allure.<br>

### pkg/allure

The package containing the data model for Allure. <br>
Complete list of allure objects:

+ `Attachment`
+ `Container`
+ `Label`
+ `Link`
+ `Parameter`
+ `Result`
+ `Step`

### pkg/framework

The package provides a fully integrated with Allure JUNIT-like framework for working with tests.<br>
Main features:

+ **Allure support**
    + Test plan support (Allure TestOps feature)
    + Tests as code
    + Extensive configuration options for test steps
    + Testify's asserts already wrapped with `allure.Step`!
    + xSkip support (you can mark test as `t.XSkip()` and it will be skipped on fail)
+ **Suite support**
    + Before/After feature
    + Suite as go-struct
    + Suite as sub-test
+ **Parallel running**
    + Parallel tests in suite structs
    + Parallel steps in test functions

## Getting Started with framework!

**Step 0.** Install package

```bash
go get github.com/ozontech/allure-go/pkg/framework
```

### No Suite tests

**NOTE:** No suite tests doesn't support before after hooks

**Step 1.** Describe tests

```go
package test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestRunner(t *testing.T) {
	runner.Run(t, "My first test", func(t provider.T) {
		t.NewStep("My First Step!")
	})
	runner.Run(t, "My second test", func(t provider.T) {
		t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
			sCtx.NewStep("My First SubStep!")
		})
	})
}
```

**Step 2.** Run it!

```bash
go test ./test/... 
```

### Suite

**Step 1.** Make your first test suite

```go
package tests

import (
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MyFirstSuite struct {
	suite.Suite
}
```

**Step 2.** Extend it with tests

```go
package tests

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MyFirstSuite struct {
	suite.Suite
}

func (s *MyFirstSuite) TestMyFirstTest(t provider.T) {
	t.NewStep("My First Step!")
}

func (s *MyFirstSuite) TestMySecondTest(t provider.T) {
	t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
		sCtx.NewStep("My First SubStep!")
	})
}
```

**Step 3.** Describe suite runner function

```go
package test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type MyFirstSuite struct {
	suite.Suite
}

func (s *MyFirstSuite) TestMyFirstTest(t provider.T) {
	t.NewStep("My First Step!")
}

func (s *MyFirstSuite) TestMySecondTest(t provider.T) {
	t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
		sCtx.NewStep("My First SubStep!")
	})
}

func TestSuiteRunner(t *testing.T) {
	suite.RunSuite(t, new(MyFirstSuite))
}
```

**Step 4.** Run it!

```bash
go test ./test/... 
```

### Runner

**Step 1.** Init runner object

```go
package test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestRunner(t *testing.T) {
	r := runner.NewRunner(t, "My First Suite!")
}
```

**Step 2.** Extend it with tests

```go
package test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestRunner(t *testing.T) {
	r := runner.NewRunner(t, "My First Suite!")
	r.NewTest("My first test", func(t provider.T) {
		t.NewStep("My First Step!")
	})

	r.NewTest("My second test", func(t provider.T) {
		t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
			sCtx.NewStep("My First SubStep!")
		})
	})
}
```

**Step 3.** Call RunTests function from runner

```go
package test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestRunner(t *testing.T) {
	r := runner.NewRunner(t, "My First Suite!")
	r.NewTest("My first test", func(t provider.T) {
		t.NewStep("My First Step!")
	})

	r.NewTest("My second test", func(t provider.T) {
		t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
			sCtx.NewStep("My First SubStep!")
		})
	})
	r.RunTests()
}
```

**Step 4.** Run it!

```bash
go test ./test/... 
```

## Going Deeper...

### pkg/allure

TBD

### pkg/framework

TBD

## Few more examples

### [Test with nested steps](examples/suite_demo/step_tree_test.go):

Test code:

```go
package examples

import (
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
```

Output to Allure:

![](.resources/example_step_tree.png)

### [Test with Attachment](examples/suite_demo/attachments_test.go)

Test code:

```go
package examples

import (
	"encoding/json"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type JSONStruct struct {
	Message string `json:"message"`
}

type AttachmentTestDemoSuite struct {
	suite.Suite
}

func (s *AttachmentTestDemoSuite) TestAttachment(t provider.T) {
	t.Epic("Demo")
	t.Feature("Attachments")
	t.Title("Test Attachments")
	t.Description(`
		Test's test body and all steps inside can contain attachments`)

	t.Tags("Attachments", "BeforeAfter", "Steps")

	attachmentText := `THIS IS A TEXT ATTACHMENT`
	t.Attachment(allure.NewAttachment("Text Attachment if TestAttachment", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Step A")
	var ExampleJson = JSONStruct{"this is JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for Step A", allure.JSON, attachmentJSON))
	t.Step(step)
}
```

Output to Allure:

![](.resources/example_attachments.png)

### [Run few parallel suites](examples/suite_demo/running_test.go)

Test code:

```go
package examples

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type TestRunningDemoSuite struct {
	suite.Suite
}

func (s *TestRunningDemoSuite) TestBeforesAfters(t provider.T) {
	t.Parallel()
	// use RunInner to run suite of tests
	s.RunSuite(t, new(BeforeAfterDemoSuite))
}

func (s *TestRunningDemoSuite) TestFails(t provider.T) {
	t.Parallel()
	s.RunSuite(t, new(FailsDemoSuite))
}

func (s *TestRunningDemoSuite) TestLabels(t provider.T) {
	t.Parallel()
	s.RunSuite(t, new(LabelsDemoSuite))
}

func TestRunDemo(t *testing.T) {
	// use RunSuites to run suite of suites
	suite.RunSuite(t, new(TestRunningDemoSuite))
}
```

Output to Allure:

![](.resources/example_multiple_suites_run.png)

### [Setup hooks](examples/suite_demo/befores_afters_test.go)

Test code:

```go
package examples

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type BeforeAfterDemoSuite struct {
	suite.Suite
}

func (s *BeforeAfterDemoSuite) BeforeEach(t provider.T) {
	t.NewStep("Before Test Step")
}

func (s *BeforeAfterDemoSuite) AfterEach(t provider.T) {
	t.NewStep("After Test Step")
}

func (s *BeforeAfterDemoSuite) BeforeAll(t provider.T) {
	t.NewStep("Before suite Step")
}

func (s *BeforeAfterDemoSuite) AfterAll(t provider.T) {
	t.NewStep("After suite Step")
}

func (s *BeforeAfterDemoSuite) TestBeforeAfterTest(t provider.T) {
	t.Epic("Demo")
	t.Feature("BeforeAfter")
	t.Title("Test wrapped with SetUp & TearDown")
	t.Description(`
		This test wrapped with SetUp and TearDown containert.`)

	t.Tags("BeforeAfter")
}

func TestBeforesAfters(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(BeforeAfterDemoSuite))
}
```

Output to Allure:

![](.resources/example_befores_afters.png)

### [XSkip](examples/suite_demo/fails_test.go)

Test code:

```go
package examples

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type DemoSuite struct {
	suite.Suite
}

func (s *DemoSuite) TestXSkipFail(t provider.T) {
	t.Title("This test skipped by assert with message")
	t.Description(`
		This Test will be skipped with assert Error.
		Error text: Assertion Failed`)
	t.Tags("fail", "xskip", "assertions")

	t.XSkip()
	t.Require().Equal(1, 2, "Assertion Failed")
}

func TestDemoSuite(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(DemoSuite))
}
```

Output to Allure:

![](.resources/example_xskip.png)
