# allure-go

Allure-Go - the project that provides a complete allure provider in go, without overloading the interface of usage. <br>
The project started as a fork of testify, but over time it got its own runner and its own features. <br>

## Other Languages README.md

- [English Readme](README_en.md)

## Head of contents

- [Other Languages README.md](#other-languages-readmemd)
- [Head of contents](#head-of-contents)
- [Features](#features)
- [Getting started](#getting-started)
- [Demo](#demo)
    - [Demo installation](#demo-installation)
    - [Run examples](#run-examples)
- [How to use](#how-to-use)
    - [Installation](#installation)
    - [Configure Behavior](#configure-behavior)
    - [Configure Test](#configure-test)
    - [Configure Suite](#configure-suite)
    - [Configure Your Report](#configure-your-report)
        - [Test Info](#test-info)
        - [Labels](#labels)
        - [Links](#links)
        - [Allure Steps](#allure-steps)
        - [Nested Only Functions](#nested-only-functions)
        - [Allure Attachments](#allure-attachments)
        - [Test Behavior](#test-behaviour)
        - [Forwarded actions to the suite.Suite](#forward-to-suitesuite)
- [Documentation](#documentation)
- [Examples](#examples)
    - [Test with nested steps](#test-with-nested-steps)
    - [Test with attachment](#test-with-attachment)
    - [Run few parallel suites](#run-few-parallel-suites)
    - [Setup hooks](#setup-hooks)
    - [XSkip](#xskip)

## Features

### `pkg/allure`

The package containing the data model for Allure. <br>
Complete list of allure objects:

- `Attachment`
- `Container`
- `Label`
- `Link`
- `Parameter`
- `Result`
- `Step`

Providing a separate package allows you to customize your work with allure.<br>
For more details see [here](#how-to-use-allure). <br>

### `pkg/framework/provider/T`

Враппер контекста теста (`testing.T`). <br>
Основные преимущества и особенности:

- Has its own test runner (`T.Run(testName string, test func(t provider.T), tags ...string)`), which allows to use the
  advantages of the `testing` library.
- The functionality of analogues in other languages, without losing the convenience and ease of use.
- Completely integrated with `allure`. Your go-tests have never been so informative!

For more details see [here] (#how-to-use-providert). <br>

### `pkg/framework/runner`

The package provides functions for running test suites and individual tests. <br>
Tests run with these functions will generate an allure report at the end of the execution.<br>
For more details you can read [here] (#how-to-use-runner). <br>

### `pkg/framework/suite`

The package provides an Suite structure in which you can describe tests by grouping them into test suites. This can be
useful if you have a lot of different tests and it is difficult to navigate through them without having additional "
layers nesting levels" of test calls. <br>
For more details you can read [here] (#how-to-use-suite). <br>

## Getting Started

1. Install package <br>

```bash
go get github.com/ozontech/allure-go
```

2. If you already use testify, you need to replace the imports

```go
package tests

import (
	"github.com/stretchr/testify/suite"
)
``` 

на

```go
package tests

import (
	"github.com/ozontech/allure-go/pkg/framework/suite"
)
``` 

3. Replace functions <br>

* `SetupSuite` -> `BeforeAll` <br>
* `SetupTest` -> `BeforeEach` <br>
* `TearDownTest` -> `AfterEach` <br>
* `TearDownSuite` -> `AfterAll` <br>

4. С Since version 0.5.0 it's required to throw in each test and hook function of the provider.T interface

```go
package tests

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type SomeSuite struct {
	suite.Suite
}

func (s *SomeSuite) BeforeAll(t provider.T) {
	// ...
}

func (s *SomeSuite) BeforeEach(t provider.T) {
	// ...
}

func (s *SomeSuite) AfterEach(t provider.T) {
	// ...
}

func (s *SomeSuite) AfterAll(t provider.T) {
	// ...
}

func (s *SomeSuite) TestSome(t provider.T) {
	// ...
}
```

5. Run go test!

## Demo

### Demo Installation

```bash
 git clone https://github.com/ozontech/allure-go.git
```

### Run Examples

```bash
make demo
```

## How to use

### Installation

```bash
go get github.com/ozontech/allure-go
```

### Configure Behavior

The path to allure reports is gathered from the two global variables `$ALLURE_OUTPUT_FOLDER/$ALLURE_OUTPUT_PATH

- The `ALLURE_OUTPUT_FOLDER` is the name of the folder where the allure reports will be stored (by
  default, `allure-results`).
- The `ALLURE_OUTPUT_PATH` is the path where the `ALLURE_OUTPUT_FOLDER` will be created (by default this is the root
  folder root folder of the test launcher).

You can also specify several global configurations to integrate with your TMS or Task Tracker:

- `ALLURE_ISSUE_PATTERN` - Specifies the url pattern for your Issues. Has no default value. **Mandatory**. Must
  contain `%s`.

If `ALLURE_ISSUE_PATTERN` is not specified, the link will be read in its entirety.

Example:

```go
package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSampleDemo(t *testing.T) {
	runner.Run(t, "Just Link", func(t provider.T) {
		t.SetIssue("https://pkg.go.dev/github.com/stretchr/testify")
	})

	runner.Run(t, "With Pattern", func(t provider.T) {
		_ = os.Setenv("ALLURE_ISSUE_PATTERN", "https://pkg.go.dev/github.com/stretchr/%s")
		t.SetIssue("testify")
	})

}

```

- ``ALLURE_TESTCASE_PATTERN`` - Specifies the url pattern for your TestCases. Has no default value. **Mandatory**. Must
  contain `%s`.

If `ALLURE_TESTCASE_PATTERN` is not specified, the link will be read in its entirety.

Example:

```go
package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSampleDemo(t *testing.T) {
	runner.Run(t, "Just Link", func(t provider.T) {
		t.SetTestCase("https://pkg.go.dev/github.com/stretchr/testify")
	})

	runner.Run(t, "With Pattern", func(t provider.T) {
		_ = os.Setenv("ALLURE_TESTCASE_PATTERN", "https://pkg.go.dev/github.com/stretchr/%s")
		t.SetTestCase("testify")
	})

}

```

- ``ALLURE_LAUNCH_TAGS`` - Sheds a list of tags that will be applied to each test by default. It has no default value.

**Tip:** ``ALLURE_LAUNCH_TAGS`` - Very handy to use with CI/CD. For example, you can define test groups in it by your
ci-jobs, or you can roll the name of a branch.

### Configure Test

1) Using the runner package:

```go
package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSampleDemo(t *testing.T) {
	runner.Run(t, "My test", func(t provider.T) {
		// Test Body 
	})
}

```

2) Using the context declaration ``TestRunner'':

```go
package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewRunner(realT, realT.Name())
	r.NewTest("My test", func(t provider.T) {
		// Test Body 
	})
	r.RunTests()
}

```

The second option will allow the use of BeforeEach/AfterEach:

```go
package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewRunner(realT, "SuiteName")

	r.BeforeAll(func(t provider.T) {
		// BeforeAll body 
	})

	r.BeforeEach(func(t provider.T) {
		// Before Each body 
	})

	r.AfterEach(func(t provider.T) {
		// After Each body 
	})

	r.AfterAll(func(t provider.T) {
		// AfterAll body 
	})

	r.NewTest("My test", func(t *provider.T) {
		// Test Body 
	})
}

```

A feature of the `testing` library has also been preserved, allowing you to run tests from other tests:

```go
package provider_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewRunner(realT, "SuiteName")
	r.NewTest("My test", func(t provider.T) {
		r2 := runner.NewRunner(t, "SuiteName")
		// Test Body 
		r2.BeforeEach(func(t provider.T) {
			// inner Before Each body 
		})
		r2.AfterEach(func(t provider.T) {
			// inner After Each body 
		})
		r2.NewTest("My test", func(t provider.T) {
			// inner test body 
		})
	})
}

```

### Configure Suite

To group tests into test suites, you must:

1) declare a structure whose methods will be your tests

```go
package suite_demo

type DemoSuite struct {
}
```

2) Extend the declared structure with the `suite.Suite` structure

```go
package suite_demo

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type DemoSuite struct {
	suite.Suite
}
```

3) Describe tests

```go
package suite_demo

import (
	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type DemoSuite struct {
	suite.Suite
}

func (s *DemoSuite) TestSkip(t provider.T) {
	t.Epic("Demo")
	t.Feature("Suites")
	t.Title("My first test")
	t.Description(`
This test will be attached to the suite DemoSuite`)
}
```

4) Run the tests.

To do this, you need to describe a function that will run your test and call ``runner.RunSuite``:

```go
package suite_demo

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
	"github.com/ozontech/allure-go/pkg/framework/suite"
)

type DemoSuite struct {
	suite.Suite
}

func (s *DemoSuite) TestSkip(t provider.T) {
	t.Epic("Demo")
	t.Feature("Suites")
	t.Title("My first test")
	t.Description(`
		This test will be attached to the suite DemoSuite`)
}

func TestSkipDemo(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(SkipDemoSuite))
}
```

And run the tests with `go test`.

```bash
go test ${TEST_PATH}
```

Then the `allure-results` folder will be initialized in the root folder of the tests at the end of the run, containing
allure-reports.

### Configure Your Report

Allure-Go provides a wide range of options for interacting with allure.<br>
Most of the actions are done with the help of the `provider.T` structure, which is a wrapper over `testing.T`.<br>
Also, the `suite.Suite` structure allows you to use the Suite interface to interact with allure-report.

#### Test Info

Full list of supported methods to put the method information:

- `*T.Title`
- `*T.Description`

**Note:** By default, the test name is set to the name of the test function.

#### Labels

Full list of supported labels:

- `*T.Epic`
- `*T.Feature`
- `*T.Story`
- `*T.ID`
- `*T.Severity`
- `*T.ParentSuite`
- `*T.Suite`
- `*T.SubSuite`
- `*T.Package`
- `*T.Thread`
- `*T.Host`
- `*T.Tag`
- `*T.Framework`
- `*T.Language`
- `*T.Owner`
- `*T.Lead`
- `*T.AllureID`

You can read more about the methods [here](/pkg/allure/README.md#allurelabel)

##### Default label values

| Label | Default Value |
|---|---|
|`ParentSuite`| - For suite.Suite the name of the function in which suite was run is put.<br><br> - For independent tests, it is not set by default, but if the test was called within another test, this label will indicate the test that started the parent for the current test.|
|`Suite`|- For suite.Suite the name of the suite to which the current test belongs is set. <br><br>  For an independent test, the name of the parent test (from which the current test was run).
|`Package`|The package in which the tests were run|
|`Thread`| Sets `Result.FullName` [[1]](#label_note_answer1).|
|`Host`|`os.Host()`|
|`Framework`|`Allure-Go@v0.3.x`|
|`Language`|`runtime.Version()`|

___________________________________  
***NOTES:***
<a name="label_note_answer1">[1]</a> - This is a Knowing Issue - in golang, it does not yet seem reasonable (or possible
by adequate means) to try to get the name of the current goroutine, nor is it possible to give it a name.
___________________________________

#### Links

Full list of supported actions:

- `*T.SetIssue`
- `*T.SetTestCase`
- `*T.Link`

You can read more about methods [here](/pkg/allure/README.md#allurelink).

The variables you can interact with to simplify your work have been described [above](#configure-behavior).

#### Allure Steps

Full list of supported actions:

- `*T.Step` - adds the transmitted Step to the report.
- `*T.NewStep` - creates a new empty Step with the passed name and adds it to the report.
- `*T.WithStep` - wraps the function passed to f with the passed Step and adds Step to the report.
- `*T.WithNewStep` - creates a new Step, wraps the passed in f function with the created Step and adds it to the report.

**Note:** Functions with the suffix `ToNested` can be called **THERE** inside the `WithStep` / `WithNewStep` function. В
Otherwise nothing will happen.

#### Allure Attachments

Full list of supported actions:

- `*T.Attachment` - adds an Attachment to the current test

#### Test Behaviour

Full list of supported actions:

- `*T.Skip` - skips the current test. The report status will include the passed text.
- `*T.Errorf` - marks the selected test as Failed. The transmitted text will be attached to the report status,
- `*T.XSkip` - skips the selected test if `*T.Error` / `*T.Errorf` is called during its execution (e.g, assert fails)

#### Forward to suite.Suite

Full list of supported actions:

- [Test Info](#test-info)
    - `*Suite.Title`.
    - `*Suite.Description`.
- [Allure Labels](#labels)
    - `*Suite.Epic`
    - `*Suite.Feature`
    - `*Suite.Story`
    - `*Suite.ID`
    - `*Suite.Severity`
    - `*Suite.ParentSuite`
    - `*Suite.Suite`
    - `*Suite.SubSuite`
    - `*Suite.Package`
    - `*Suite.Thread`
    - `*Suite.Host`
    - `*Suite.Tag`
    - `*Suite.Framework`
    - `*Suite.Language`
    - `*Suite.Owner`
    - `*Suite.Lead`
- [Allure Links](#links)
    - `*Suite.SetIssue`
    - `*Suite.SetTestCase`
    - `*Suite.Link`
- [Allure Steps](#allure-steps)
    - `*Suite.Step`
    - `*Suite.NewStep`
    - `*Suite.WithStep`
    - `*Suite.WithNewStep`
- [Allure Attachments](#allure-attachments)
    - `*Suite.Attachment`

## Documentation

Detailed documentation for each public package can be found in that package's directory.

- [allure](/pkg/allure/README.md)
- [runner](/pkg/framework/runner/README.md)
- [suite](/pkg/framework/suite/README.md)

## Examples

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
