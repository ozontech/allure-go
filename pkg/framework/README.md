# pkg/framework

## Head of contents

+ [Interfaces](#Interfaces)
+ [Test as code](#Test as code)
+ [Asserts](#Asserts)
+ [Test Running](#Test Running)

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

| Method                          |           Description           |
|:--------------------------------|:-------------------------------:|
| Title(title string)             |    Sets `result.Name` field     |
| Description(description string) | Sets `result.Description` field |

##### Suite Methods (`SuiteLabels` interface)

| Method                       |           Description           |
|:-----------------------------|:-------------------------------:|
| AddSuiteLabel(value string)  |    Adds `suite` allure label    |
| AddSubSuite(value string)    |  Adds `subSuite` allure label   |
| AddParentSuite(value string) | Adds `parentSuite` allure label |

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

**NOTE**: Some labels (such as `languange`, `host`, `framework`, etc) have default values and cannot be set during test runtime any other way (`SystemLabels` interface) but ReplaceLabel method.

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

**Note**: Those methods **will create** file at your `allure-results` folder.

##### Steps methods (`AllureSteps` interface and some method in `T` interface)

| Method                                                                                   |                                                                    Description                                                                    |
|:-----------------------------------------------------------------------------------------|:-------------------------------------------------------------------------------------------------------------------------------------------------:|
| `Step(step *allure.Step)`                                                                |                                                       Adds `allure.Step` object to result.                                                        |
| `NewStep(stepName string, params ...allure.Parameter)`                                   |                                              Creates new `allure.Step` object and adds it to result.                                              |
| `WithNewStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)`      | Creates new `allure.Step` object and run anonymous function. With `StepCtx` interface you can work with step during anonymous function execution. |
| `WithNewAsyncStep(stepName string, step func(sCtx StepCtx), params ...allure.Parameter)` |                                          Same as `WithNewStep`, but it runs as async process with test.                                           |

TBD

### provider.StepCtx

TBD

## Test as code

## Asserts

## Test Running