# pkg/framework

## Head of contents

+ [Head of contents](#head-of-contents)
+ [Interfaces](#interfaces)
  + [provider.T](#providert)
    + [Extended methods](#extended-methods)
    + [Specific methods](#specific-methods)
      + [Description methods](#description-methods-descriptionfields-interface)
      + [Suite methods](#suite-methods-suitelabels-interface)
      + [Description label methods](#description-label-methods-descriptionlabels-interface)
      + [Link methods](#link-methods-links-interface)
      + [Attachment methods](#attachment-methods-attachments-interface)
      + [Assertion methods](#assertion-methods-t-interface)
      + [Steps methods](#steps-methods-alluresteps-interface-and-some-method-in-t-interface)
      + [Test Run methods](#test-run-function-t-interface)
      + [Behaviour manipulation methods](#behaviour-manipulation-methods-t-interface)
  + [provider.StepCtx](#providerstepctx)
    + [Steps methods](#steps-methods)
    + [Parameters methods](#parameters-methods)
    + [Attachments methods](#attachments-methods)
    + [Assertion methods](#assertion-methods)
    + [Step condition and log methods](#step-condition-and-log-methods)
  + [provider.Asserts](#providerasserts)
+ [Test Running](#test-running)
  + [No suite running](#no-suite-running)
  + [Suite with runner object](#suite-with-runner-object)
  + [Suite with struct](#suite-with-struct)

## Interfaces

Main interfaces for test working are `provider.T` and `provider.StepCtx`.

### provider.T

#### Extended methods

`provider.T` extends [`testing.TB`](https://pkg.go.dev/testing#TB) interface and supports all it methods.

|      Method      |                                                                   Description                                                                    |
|:----------------:|:------------------------------------------------------------------------------------------------------------------------------------------------:|
|      `Name`      |                                                              Returns `result.Name`                                                               |
|      `Fail`      |                                   Fails test. Marks test as `Failed`. `Fail` **DOESN'T STOPS** test execution.                                   |
|    `FailNow`     |                                       Fails test. Marks test as `Failed`. `Fail` **STOPS** test execution.                                       |                                                                                             |
| `Error`/`Errorf` | Fails test, marks result.Status as Failed and add error message to `result.StatusDetails`. `Error` and `Errorf` **DOESN'T STOP** test execution. |
| `Fatal`/`Fatalf` |    Fails test, marks result.Status as Failed and add error message to `result.StatusDetails`. `Fatal` and `Fatalf` **STOPS** test execution.     |
|  `Skip`/`Skipf`  |                         Skips test and add skip message to `result.Status`. `Skip` and `Skipf`**STOPS** test execution.                          |

#### Specific methods

`provider.T` suggests a lot of methods for describing your allure report, run tests and control your tests.

##### Description Methods (`DescriptionFields` interface)

| Method                            |           Description           |
|:----------------------------------|:-------------------------------:|
| `Title(title string)`             |    Sets `result.Name` field     |
| `Description(description string)` | Sets `result.Description` field |

##### Suite Methods (`SuiteLabels` interface)

| Method                         |           Description           |
|:-------------------------------|:-------------------------------:|
| `AddSuiteLabel(value string)`  |    Adds `suite` allure label    |
| `AddSubSuite(value string)`    |  Adds `subSuite` allure label   |
| `AddParentSuite(value string)` | Adds `parentSuite` allure label |

##### Description Label Methods (`DescriptionLabels` interface)

| Method                                       |                      Description                       |
|:---------------------------------------------|:------------------------------------------------------:|
| `ID(value string)`                           |                 Adds `id` allure label                 |
| `AllureID(value string)`                     |             Adds `ALLURE_ID` allure label              |
| `Epic(value string)`                         |                Adds `epic` allure label                |
| `Feature(value string)`                      |              Adds `feature` allure label               |
| `Story(value string)`                        |               Adds `story` allure label                |
| `Severity(severityType allure.SeverityType)` |              Adds `severity` allure label              |
| `Tag(value string)`                          |                Adds `tag`  allure label                |
| `Tags(values ...string)`                     |           Adds multiple `tag` allure labels            |
| `Owner(value string)`                        |                Adds `owner`allure label                |
| `Lead(value string)`                         |                Adds `lead` allure label                |
| `Label(label allure.Label)`                  |                Adds custom allure label                |
| `Labels(labels ...allure.Label)`             |              Adds multiple allure labels               |
| `ReplaceLabel(label allure.Label)`           | Replace any label with same name as passed to argument |

:warning: **NOTE**: Some labels (such as `languange`, `host`, `framework`, etc) have default values and cannot be set during test runtime any other way (`SystemLabels` interface) but ReplaceLabel method.

##### Link Methods (`Links` interface)

| Method                         |                                              Description                                              |
|:-------------------------------|:-----------------------------------------------------------------------------------------------------:|
| `SetIssue(issue string)`       |    Sets `issue` link. You can use `ALLURE_ISSUE_PATTERN` environment variable to set link pattern.    |
| `SetTestCase(testCase string)` | Sets `testCase` link. You can use `ALLURE_TESTCASE_PATTERN` environment variable to set link pattern. |
| `Link(link allure.Link)`       |                                           Sets custom link.                                           |

##### Attachment methods (`Attachments` interface)

| Method                                                                     |                      Description                       |
|:---------------------------------------------------------------------------|:------------------------------------------------------:|
| `WithNewAttachment(name string, mimeType allure.MimeType, content []byte)` | Creates new `allure.Attachment` and adds it to result. |
| `WithAttachments(attachment ...*allure.Attachment)`                        |   Adds multiple `allure.Attachment`s to the `result`   |

:warning: **Note**: Those methods **will create** file at your `allure-results` folder.

##### Steps methods (`AllureSteps` interface and some method in `T` interface)

| Method                                                                                   |                                                                    Description                                                                    |
|:-----------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------:|
| `Step(step *allure.Step)`                                                                |                                                       Adds `allure.Step` object to result.                                                        |
| `NewStep(stepName string, params ...allure.Parameter)`                                   |                                              Creates new `allure.Step` object and adds it to result.                                              |
| `WithNewStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)`      | Creates new `allure.Step` object and run anonymous function. With `StepCtx` interface you can work with step during anonymous function execution. |
| `WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)` |                                          Same as `WithNewStep`, but it runs as async process with test.                                           |

##### Assertion methods (`T` interface)

| Method               |                                                                                 Description                                                                                 |
|:---------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| 	`Assert() Asserts`  |                   Returns struct, that contains a lot of asserts that fails test, but **NOT STOPS** its execution. Creates step with assert description.                    |
| 	`Require() Asserts` |                      Returns struct, that contains a lot of asserts that fails test and **STOPS** its execution. Creates step with assert description.                      |

##### Test run function (`T` interface)

| Method                                                        |                                                       Description                                                        |
|:--------------------------------------------------------------|:------------------------------------------------------------------------------------------------------------------------:|
| `Run(testName string, testBody func(T), tags ...string) bool` | Runs passed anonymous function as test. Returns true if test succeed, false if not. Also it adds passed tags for report. |

##### Behaviour manipulation methods (`T` interface)

| Method          |                                                                Description                                                                 |
|:----------------|:------------------------------------------------------------------------------------------------------------------------------------------:|
| `XSkip()`       |                     Marks test as expected to fail. If test going to fail with assert it will be marked skip instead.                      |
| `SkipOnPrint()` | Marks report as skip on print. That means that report won't be created for current test. Use it for clean reports from parent of subtests. |

### provider.StepCtx

`StepCtx` interface provides wide list of ways to work with step and test.

#### Steps methods

| Method                                                                                   |                                                                              Description                                                                              |
|:-----------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| `Step(step *allure.Step)`                                                                |                                                               Adds created `allure.Step` as a substep.                                                                |
| `NewStep(stepName string, parameters ...allure.Parameter)`                               |                                                       Creates new allure.Step object and adds it as a substep.                                                        |
| `WithNewStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)`      | Creates new `allure.Step` object and run anonymous function. With `StepCtx` interface you can work with step during anonymous function execution. Adds it as substep. |
| `WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)` |                                                 Same as `WithNewStep`, but runs anonymous function as async process.                                                  |
| `CurrentStep() *allure.Step`                                                             |                                                         Returns pointer to the current `allure.Step` object.                                                          |

#### Parameters methods

| Method                                           |                                                Description                                                |
|:-------------------------------------------------|:---------------------------------------------------------------------------------------------------------:|
| `WithParameters(parameters ...allure.Parameter)` |                          Add passed list of `allure.Parameter` to current step.                           |
| `WithNewParameters(kv ...interface{})`           | Create new parameters from passed strings. All odd arguments are keys, and all even arguments are values. |

#### Attachments methods

| Method                                                                     |                             Description                              |
|:---------------------------------------------------------------------------|:--------------------------------------------------------------------:|
| `WithAttachments(attachment ...*allure.Attachment)`                        |             Add `allure.Attachment` to the current step.             |
| `WithNewAttachment(name string, mimeType allure.MimeType, content []byte)` | Create new `allure.Attachment` file and adds it to the current step. |

#### Assertion methods

| Method              |                                                                Description                                                                |
|:--------------------|:-----------------------------------------------------------------------------------------------------------------------------------------:|
| `Assert() Asserts`  | Returns struct, that contains a lot of asserts that fails test, but **NOT STOPS** its execution. Creates substep with assert description. |
| `Require() Asserts` |   Returns struct, that contains a lot of asserts that fails test and **STOPS** its execution. Creates substep with assert description.    |

#### Step condition and log methods

| Method                                                                    |                                       Description                                        |
|:--------------------------------------------------------------------------|:----------------------------------------------------------------------------------------:|
| `Broken()`                                                                |                       Marks step and all parent steps as `broken`.                       |
| `Fail()`                                                                  |                       Marks step and all parent steps as `failed`.                       |
| `FailNow()`                                                               | **STOPS** test execution immediately. Marks step, all parent steps and test as `failed`. |
| `Error(args ...interface{})`/`Errorf(format string, args ...interface{})` |   **DOESN'T STOP test execution.** Marks step, all parent steps and test as `failed`.    |
| `Log(args ...interface{})`/`Logf(format string, args ...interface{})`     |                               Same as `testing.TB` analog.                               |
| `Name() string      `                                                     |                                    Returns test name.                                    |

### provider.Asserts

allure-go provides implementation of most usable [testify](https://github.com/stretchr/testify) asserts. There are full list of supported asserts:

| Method                                                                                       |
|:---------------------------------------------------------------------------------------------|
| `Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{})`                 | 
| `NotEqual(expected interface{}, actual interface{}, msgAndArgs ...interface{})`              | 
| `Error(err error, msgAndArgs ...interface{})`                                                | 
| `NoError(err error, msgAndArgs ...interface{})`                                              | 
| `NotNil(object interface{}, msgAndArgs ...interface{})`                                      | 
| `Nil(object interface{}, msgAndArgs ...interface{})`                                         | 
| `Len(object interface{}, length int, msgAndArgs ...interface{})`                             | 
| `NotContains(s interface{}, contains interface{}, msgAndArgs ...interface{})`                | 
| `Contains(s interface{}, contains interface{}, msgAndArgs ...interface{})`                   | 
| `Greater(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                         | 
| `GreaterOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                  | 
| `Less(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                            | 
| `LessOrEqual(e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                     | 
| `Implements(interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})`     | 
| `Empty(object interface{}, msgAndArgs ...interface{})`                                       | 
| `NotEmpty(object interface{}, msgAndArgs ...interface{})`                                    | 
| `WithinDuration(expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{})` | 
| `JSONEq(expected, actual string, msgAndArgs ...interface{})`                                 | 
| `Subset(list, subset interface{}, msgAndArgs ...interface{})`                                | 
| `IsType(expectedType interface{}, object interface{}, msgAndArgs ...interface{})`            | 
| `True(value bool, msgAndArgs ...interface{})`                                                | 
| `False(value bool, msgAndArgs ...interface{})`                                               |

:information_desk_person: **NOTE:** allure-go supports assert/require separation. User `T.Assert()`/`T.Require()` to get asserts you need.

:information_desk_person: **NOTE:** If you need assert that does not supported in allure-go, but it supported in testify (or something special and unique), please,
create an [issue](https://github.com/ozontech/allure-go/issues/new?assignees=&labels=&template=feature_request.md&title=), and we will add this assert as soon as possible.

Also, allure-go supports assert/require functionality that not attached to the `T` or `StepCtx` interfaces from `asserts`/`require` packages:

| Method                                                                                                    |
|:----------------------------------------------------------------------------------------------------------|
| `Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{})`                 | 
| `NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{})`              | 
| `Error(t ProviderT, err error, msgAndArgs ...interface{})`                                                | 
| `NoError(t ProviderT, err error, msgAndArgs ...interface{})`                                              | 
| `NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{})`                                      | 
| `Nil(t ProviderT, object interface{}, msgAndArgs ...interface{})`                                         | 
| `Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{})`                             | 
| `NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{})`                | 
| `Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{})`                   | 
| `Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                         | 
| `GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                  | 
| `Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                            | 
| `LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})`                     | 
| `Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})`     | 
| `Empty(t ProviderT, object interface{}, msgAndArgs ...interface{})`                                       | 
| `NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{})`                                    | 
| `WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{})` | 
| `JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{})`                                 | 
| `Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{})`                                | 
| `IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{})`            | 
| `True(t ProviderT, value bool, msgAndArgs ...interface{})`                                                | 
| `False(t ProviderT, value bool, msgAndArgs ...interface{})`                                               |

:information_desk_person: **NOTE:** `ProviderT` interface:

```go
package asserts

type ProviderT interface {
	Step(step *allure.Step)
	Errorf(format string, args ...interface{})
	FailNow()
}
```

:warning: **NOTE:** USING REQUIRE ASSERTS WITH ASYNC STEPS ARE NOT RECOMMENDED. Reason: `testing.T.FailNow()` makes `go.Exit()` and It's impossible to handle this situation, so you can lose your step or test data. 

## Test Running

allure-go provides wide list of ways to run your tests. There are few simple examples:

:information_desk_person: **NOTE:** For more examples [click here](../../examples).

### No suite running

```go
package test

import (
  "testing"

  "github.com/ozontech/allure-go/pkg/framework/provider"
  "github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestMyTest(t *testing.T) {
  runner.Run(t, "My first test", func(t provider.T) {
    // test body...
  }, "sampleTag1", "sampleTag2")
}
```

### Suite with runner object

:information_desk_person: **FYI** runner supports before/after each/all functions

```go
package test

import (
  "testing"

  "github.com/ozontech/allure-go/pkg/framework/provider"
  "github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestMyTest(t *testing.T) {
  r := runner.NewRunner(t, t.Name())

  r.BeforeAll(func(t provider.T) {
     // This will be executed before all tests start ...
  })
  r.BeforeEach(func(t provider.T) {
    // This will be executed before each test start ...
  })
  r.AfterEach(func(t provider.T) {
     // This will be executed after each test ...
  })
  r.AfterAll(func(t provider.T) {
     // This will be executed when all tests over ...
  })

  r.NewTest("My test 1", func(t provider.T) {
     // Test Body...
  }, "sampleTag1", "sampleTag2")

  r.RunTests()
}
```

### Suite with struct

```go
package suite_demo

import (
  "testing"

  "github.com/ozontech/allure-go/pkg/framework/provider"
  "github.com/ozontech/allure-go/pkg/framework/suite"
)

type SampleSuite struct {
  suite.Suite
}

func (s *SampleSuite) BeforeAll(t provider.T) {
  // This will be executed before all tests start ...
}

func (s *SampleSuite) AfterAll(t provider.T) {
  // This will be executed when all tests over ...
}

func (s *SampleSuite) BeforeEach(t provider.T) {
  // This will be executed before each test start ...
}

func (s *SampleSuite) AfterEach(t provider.T) {
  // This will be executed after each test ...
}

func (s *SampleSuite) TestBeforeAfterTest(t provider.T) {
  // Test Body ...
}

func TestRunner(t *testing.T) {
  suite.RunSuite(t, new(SampleSuite))
}
```
