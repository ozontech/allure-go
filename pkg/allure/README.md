# pkg/allure

![](../../.resources/allure_icon.svg)

The allure package offers an implementation of all the entities that Allure uses to handle test reports.<br>

:information_desk_person: Learn more about the [**Allure Framework**](https://docs.qameta.io/allure/).


## Head of contents

+ [:mortar_board: Head of contents](#head-of-contents)
+ [:earth_americas: Global Environment Keys](#global-environment-keys)
+ [:briefcase: Status](#status)
+ [:page_facing_up: Attachment](#attachment)
  + [Attachment's Supported Types](#attachments-supported-types)
  + [Attachment's Constructors](#attachments-constructors)
  + [Attachment's Methods](#attachments-methods)
+ [:mailbox_with_no_mail: Container](#container)
  + [Container's Constructors](#containers-constructors)
  + [Container's Methods](#containers-methods)
+ [:speech_balloon: Label](#label)
  + [Supported Label Types And Severity Levels](#supported-label-types-and-severity-levels)
  + [Label's Constructors](#labels-constructors)
+ [:email: Link](#link)
  + [Link Types](#link-types)
  + [Link's Constructors](#links-constructors)
+ [:nut_and_bolt: Parameter](#parameter)
  + [Parameter's Constructors](#parameters-constructors)
+ [:rocket: Result](#result)
  + [Result's Constructors](#results-methods)
  + [Result's Methods](#results-constructors)
+ [:walking: Step](#step)
  + [Step's Constructors](#steps-constructors)
  + [Step's Methods](#steps-methods)

## Global Environment Keys

| Key                       | Meaning                                                                                                                    | Default           |
|---------------------------|----------------------------------------------------------------------------------------------------------------------------|-------------------|
| `ALLURE_OUTPUT_PATH`      | Specifies the path to the folder to print the results.                                                                     | `.` (test folder) |
| `ALLURE_OUTPUT_FOLDER`    | Specifies the name of a folder for printing results.                                                                       | `/allure-results` |
| `ALLURE_ISSUE_PATTERN`    | Specifies the URL pattern for Issue. **Must contain exactly one `%s`**.                                                    |                   |
| `ALLURE_TESTCASE_PATTERN` | Specifies the URL pattern for TestCase. **Must contain exactly one `%s`**.                                                 |                   |
| `ALLURE_LAUNCH_TAGS`      | Specifies the default tags that will be used to mark all tests in the run. The tags must be specified separated by commas. |                   |

## Status

Supported test statuses:

|   Name    |    Key    |
|:---------:|:---------:|
| `Passed`  | `passed`  |
| `Failed`  | `failed`  |
| `Skipped` | `skipped` |
| `Broken`  | `broken`  |
| `Unknown` | `unknown` |

NOTE: Tests failed in the BeforeAll/BeforeEach functions have the status Unknown

## Attachment

[`allure.Attachment`](attachment.go) - is the implementation of the appendices to the report in allure. It is most often used to contain
screenshots, api-answers, files and other data obtained during the test.

### Attachment's Supported Types

|    Key    |           Mime type            |  File type  |
|:---------:|:------------------------------:|:-----------:|
|  `Text`   |          "text/plain"          |   `.txt`    |
|   `Csv`   |           "text/csv"           |   `.csv`    |
|   `Tsv`   |  "text/tab-separated-values"   |   `.tsv`    |
| `URIList` |        "text/uri-list"         |   `.uri`    |
|  `HTML`   |          "text/html"           |   `.html`   |
|   `XML`   |       "application/xml"        |   `.xml`    |
|  `JSON`   |       "application/json"       |   `.json`   |
|  `Yaml`   |       "application/yaml"       |   `.yaml`   |
|  `Pcap`   | "application/vnd.tcpdump.pcap" |   `.pcap`   |
|   `Png`   |          "image/png"           |   `.png`    |
|   `Jpg`   |          "image/jpg"           |   `.jpg`    |
|   `Svg`   |        "image/svg-xml"         |   `.svg`    |
|   `Gif`   |          "image/gif"           |   `.gif`    |
|   `Bmp`   |          "image/bmp"           |   `.bmp`    |
|  `Tiff`   |          "image/tiff"          |   `.tiff`   |
|   `Mp4`   |          "video/mp4"           |   `.mp4`    |
|   `Ogg`   |          "video/ogg"           |   `.ogg`    |
|  `Webm`   |          "video/webm"          |   `.webm`   |
|  `Mpeg`   |          "video/mpeg"          |   `.mpeg`   |
|   `Pdf`   |       "application/pdf"        |   `.pdf`    |

### Attachment's Constructors

| Function                                                                    |                      Description                       |
|:----------------------------------------------------------------------------|:------------------------------------------------------:|
| `NewAttachment(name string, mimeType MimeType, content []byte) *Attachment` | Returns pointer to the new `allure.Attachment` object. |

### Attachment's Methods

| Method                |                                             Description                                             |
|:----------------------|:---------------------------------------------------------------------------------------------------:|
| `GetUUID() string`    |                                      Returns attachment's UUID                                      |
| `GetContent() []byte` |                                    Returns attachment's Content                                     |
| `Print() error`       | Creates a file from `Attachment.content`. The file type is determined by its `Attachment.mimeType`. |

## Container

[`allure.Container`](container.go) - This is an implementation of the `Container` entity used by Allure to handle TestSetup and TestTeardown hooks.
The list of The list of container-dependent tests is contained in the array `Container.Children`.

### Container's Constructors


| Function                    |                      Description                      |
|:----------------------------|:-----------------------------------------------------:|
| `NewContainer() *Container` | Returns pointer to the new `allure.Container` object. |

### Container's Methods

| Method                          |                                                                                 Description                                                                                 |
|:--------------------------------|:---------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| `GetUUID() string`              |                                                                          Returns container's UUID.                                                                          |
| `AddChild(childUUID uuid.UUID)` |                                                                         Adds passed UUID as child.                                                                          |
| `IsEmpty() bool`                |                                               Returns `true` if arrays `Container.Befores` and `Container.Afters` are empty.                                                |
| `Print() error`                 |                     Creates `xxxxxx-container.json` file and call `PrintAttachments` if any step exists in `Container.Befores` and `Container.Afters`.                      |
| `ToJSON() ([]byte, error)`      |                                                     Marshall `allure.Container` to the JSON. Returns error if has any.                                                      |
| `PrintAttachments()`            | PrintAttachments It goes through all `Container.Befores` and `Container.Afters` of the Container and calls the `Container.PrintAttachments()` method at each `allure.Step`. |
| `Begin()`                       |                                                                 Sets `Container.Start` = `allure.GetNow()`.                                                                 |
| `Finish()`                      |                                                                 Sets `Container.Stop` = `allure.GetNow()`.                                                                  |
| `Done() error`                  |                                                     Calls `Finish()` and `Print()` methods.  Returns error if has any.                                                      |


## Label

[`Label`](label.go) - implementation of the `Label` entity used by Allure for metrics and test grouping.

### Supported Label Types and Severity levels

Label Types:

| Name          |      Key      |
|:--------------|:-------------:|
| `Epic`        |    `epic`     |       
| `Feature`     |   `feature`   |    
| `Story`       |    `story`    |      
| `ID`          |    `as_id`    |         
| `Severity`    |  `severity`   |   
| `ParentSuite` | `parentSuite` |
| `Suite`       |    `suite`    |      
| `SubSuite`    |  `subSuite`   |   
| `Package`     |   `package`   |    
| `Thread`      |   `thread`    |     
| `Host`        |    `host`     |       
| `Tag`         |     `tag`     |        
| `Framework`   |  `framework`  |  
| `Language`    |  `language`   |   
| `Owner`       |    `owner`    |      
| `Lead`        |    `lead`     |       
| `AllureID`    |  `ALLURE_ID`  |

Severity levels:

| Name       |    Key     |
|:-----------|:----------:|
| `BLOCKER`  | `blocker`  |
| `CRITICAL` | `critical` |
| `NORMAL`   |  `normal`  |
| `MINOR`    |  `minor`   |
| `TRIVIAL`  | `trivial`  |

### Label's Constructors

| Function                                            |                                        Description                                        |
|:----------------------------------------------------|:-----------------------------------------------------------------------------------------:|
| `NewLabel(labelType LabelType, value string) Label` | Builds and returns a new `allure.Label`. The label key depends on the passed `LabelType`. |
| `EpicLabel (epic string) Label`                     |                                 Returns new `Epic` label                                  |     
| `FeatureLabel (feature string) Label`               |                                Returns new `Feature` label                                |  
| `StoryLabel (story string) Label`                   |                                 Returns new `Story` label                                 |    
| `IDLabel (testID string) Label`                     |                                  Returns new `ID` label                                   |       
| `SeverityLabel (severity SeverityType) Label`       |                               Returns new `Severity` label                                |
| `ParentSuitLabel (parent string) Label`             |                              Returns new `ParentSuite` label                              |
| `SuiteLabel (suite string) Label`                   |                                 Returns new `Suite` label                                 |    
| `SubSuiteLabel (subSuite string) Label`             |                               Returns new `SubSuite` label                                |
| `PackageLabel (package string) Label`               |                                Returns new `Package` label                                |  
| `ThreadLabel (thread string) Label`                 |                                Returns new `Thread` label                                 |   
| `HostLabel (host string) Label`                     |                                 Returns new `Host` label                                  |     
| `TagLabel (tag string) Label`                       |                                  Returns new `Tag` label                                  |    
| `TagLabels(tags ...string) []Label`                 |                Returns as many new `Tag` labels as passed strings in args                 |
| `FrameworkLabel (framework string) Label`           |                               Returns new `Framework` label                               |
| `LanguageLabel (language string) Label`             |                               Returns new `Language` label                                |
| `OwnerLabel (owner string) Label`                   |                                 Returns new `Owner` label                                 |
| `LeadLabel (lead string) Label`                     |                                 Returns new `Lead` label                                  |    
| `IDAllureLabel (allureID string) Label`             |                               Returns new `AllureID` label                                |     

## Link

[`Link`](link.go) - is an implementation of the Link entity used by Allure to specify the links needed for test reporting.

### Link Types

| Name       |     Key     |            Description             |
|:-----------|:-----------:|:----------------------------------:|
| `LINK`     |   `link`    |            Custom link             |
| `ISSUE`    |   `issue`   |   Link to the Issue in your CMS    |
| `TESTCASE` | `test_case` |  Link to the TestCase in your TMS  |

### Link's Constructors

| Function                                                 |                                           Description                                            |
|:---------------------------------------------------------|:------------------------------------------------------------------------------------------------:|
| `NewLink(name string, _type LinkTypes, url string) Link` |                           Base constructor. Returns new `allure.Link`.                           |
| `TestCaseLink(testCase string) Link`                     | Returns `TESTCASE` type link. It uses environment variable `ALLURE_TESTCASE_PATTERN` as pattern. |
| `IssueLink(issue string) Link`                           |    Returns `ISSUE` type link. It uses environment variable `ALLURE_ISSUE_PATTERN` as pattern.    |
| `LinkLink(linkname, link string) Link`                   |                                    Returns `LINK` type link.                                     |

**NOTE:** Check more about patterns [here](#global-environment-keys)

## Parameter

[`Parameter`](parameter.go) - is an implementation of the `Parameter` entity, which Allure uses as additional information describing the test step (e.g. request host or server address).

### Parameter's Constructors

| Function                                                    |                                                            Description                                                             |
|:------------------------------------------------------------|:----------------------------------------------------------------------------------------------------------------------------------:|
| `NewParameter(name string, value ...interface{}) Parameter` |                              Builds new `Parameter` object. Value **must** be able to cast to string.                              |
| `NewParameters(kv ...interface{}) []Parameter`              | Returns list of `allure.Parameter` objects. Each even string is considered a parameter name, and each  odd-value of the parameter. |

## Result

[`Result`](result.go) - is an implementation of the Result entity used by Allure to store information about the test. It contains information about the test name, applications, description, status, references, labels, steps, containers, and time test execution time.

### Result's Constructors

| Function                                       |                               Description                                |
|:-----------------------------------------------|:------------------------------------------------------------------------:|
| `NewResult(testName, fullName string) *Result` | Builds a new `allure.Result`. Sets the default values for the structure. |

### Result's methods

| Method                                       |                                                                                                  Description                                                                                                   |
|:---------------------------------------------|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------:|
| `SetStatusMessage(msg string)`               |                                                                                      Sets `Result.StatusDetails.Message`.                                                                                      |
| `ToJSON() ([]byte, error)`                   |                                                                           Marshall `allure.Result` to JSON, returns error if has any                                                                           |
 | `GetStatusMessage() string`                  |                                                                                     Returns `Result.StatusDetails.Trace`.                                                                                      |
 | `SetStatusTrace(trace string)`               |                                                                                       Sets `Result.StatusDetails.Trace`.                                                                                       |
 | `GetStatusTrace() string`                    |                                                                                     Returns `Result.StatusDetails.Trace`.                                                                                      |
 | `SetLabel(labels ...Label)`                  |                                                                            Sets all labels passed as arguments to `allure.Result`.                                                                             |
 | `GetLabel(labelType LabelType) []Label`      |                                                                                Returns all labels with keys to `allure.Result`.                                                                                |
 | `SetNewLabelMap(kv map[LabelType]string)`    |                                                                                  Sets new labels with map to `allure.Result`.                                                                                  |
 | `WithParentSuite(parentName string) *Result` |                                                                                   Sets `ParentSuite` label `allure.Result`.                                                                                    |
 | `WithSuite(suiteName string) *Result`        |                                                                                      Sets `Suite` label `allure.Result`.                                                                                       |
 | `WithHost(hostName string) *Result`          |                                                                                       Sets `Host` label `allure.Result`.                                                                                       |
 | `WithSubSuites(children ...string) *Result`  |                                                                                 Sets all `SubSuite` labels to `allure.Result`.                                                                                 |
 | `WithFrameWork(framework string) *Result`    |                                                                                   Sets `Framework` label to `allure.Result`.                                                                                   |
 | `WithLanguage(language string) *Result`      |                                                                                   Sets `Language` label to `allure.Result`.                                                                                    |
 | `WithThread(thread string) *Result`          |                                                                                    Sets `Thread` label to `allure.Result`.                                                                                     |
 | `WithPackage(pkg string) *Result`            |                                                                                    Sets `Package` label to `allure.Result`.                                                                                    |
 | `WithLabels(label ...Label) *Result`         |                                                                             Sets all labels passed as arguments to `allure.Result`                                                                             |
 | `WithLaunchTags() *Result`                   |                                                  Adds all Launch Tags from the global variable `ALLURE_LAUNCH_TAGS` as labels with type `Tag` to the report.                                                   |
 | `Begin() *Result`                            |                                                                                   Sets `Result.Start` == `allure.GetNow()`.                                                                                    |
 | `Finish() *Result`                           |                                                                                    Sets `Result.Stop` == `allure.GetNow()`.                                                                                    |
 | `SkipOnPrint()`                              |                                                                            Skips result from printing when `Result.Print()` called.                                                                            |
 | `Print() error`                              |                                        If `Result.ToPrint` == `false` creates `uuid4-result.json` and call `Print()` method for all attachments and step's attachments.                                        |
 | `PrintAttachments()`                         | Goes through all `Result.Steps` of the report and for each allure.Step calls the `Step.PrintAttachments()` method.Then calls `Attachment.Print()` on all `allure.Attachment` of the `Result.Attachments` list. |
 | `Done() error`                               |                  If `Result.Status` is not filled in, consider the test successfully completed (no errors). After that - it calls `Finish()` and `Print()` methods. Returns error if has any.                  |

## Step

[`Step`](step.go) - is an implementation of the `Step` entity used by Allure to define and describe test steps. 

Steps can be nested, have a status (successful, failed, skipped, broken), can contain `Attachment`s and `Parameter`s and have an execution time

### Step's Constructors

| Function                                                                                     |                         Description                         |
|:---------------------------------------------------------------------------------------------|:-----------------------------------------------------------:|
| `NewStep(name string, status Status, start int64, stop int64, parameters []Parameter) *Step` |      Returns pointer to the new `allure.Step` object.       |
| `NewSimpleStep(name string, parameters ...Parameter) *Step`                                  | Same as `NewStep` but the most of fields are pre populated. |


### Step's methods

| Method                                              |                                             Description                                             |
|:----------------------------------------------------|:---------------------------------------------------------------------------------------------------:|
| `GetParent() *Step`                                 |                              Returns pointer to parent step (if any).                               |
| `WithAttachments(attachments ...*Attachment) *Step` |                                      Adds attachments to step.                                      |
| `WithParameters(params ...Parameter) *Step`         |                                Adds `Allure.Parameter`s to the step.                                |
| `WithNewParameters(kv ...interface{}) *Step`        |                    Creates new `Allure.Parameters` and attach them to the step.                     |
| `Passed() *Step`                                    |                                       Marks step as `Passed`.                                       |
| `Failed() *Step`                                    |                                       Marks step as `Failed`.                                       |
| `Skipped() *Step`                                   |                                      Marks step as `Skipped`.                                       |
| `Broken() *Step`                                    |                                       Marks test as `Broken`.                                       |
| `Begin() *Step`                                     |                               Sets `Step.Start` == `allure.GetNow()`.                               |
| `Finish() *Step`                                    |                               Sets `Step.Start` == `allure.GetNow()`.                               |
| `WithParent(parent *Step) *Step`                    |                                     Sets passed step as parent.                                     |
| `WithChild(child *Step) *Step`                      |                                     Sets passed step as child.                                      |
| `PrintAttachments()`                                | Iterate throw list of attachments, attached to `allure.Step` and call `Print()` at each attachment. |

**NOTE:** The most step methods can be called in chain of calls. Example:

```go
package test

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

// ...

func (s *SomeSuite) SomeTest(t provider.T) {
	step := allure.NewSimpleStep("My First Step").WithNewParameters("k1", "v1").Passed().Finish()
	t.Step(step)
}
```
