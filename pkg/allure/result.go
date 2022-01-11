package allure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const (
	defaultTagsEnvKey = "ALLURE_LAUNCH_TAGS" // Indicates the default tags that will mark all tests in the run. The tags must be specified separated by commas.
)

type IResult interface {
	Printable
	WithSteps
	WithTimer
	WithParentSuite(parentName string) *Result
	WithFrameWork(framework string) *Result
	WithThread(thread string) *Result
	WithLanguage(language string) *Result
	WithPackage(pkg string) *Result
	WithSuite(suiteName string) *Result
	WithSubSuites(children ...string) *Result
}

// Result is an implementation of the Result entity used by Allure to store information about the test. It contains
// information about the test name, applications, description, status, references, labels,
// steps, containers, and time of the test execution.
type Result struct {
	Name          string        `json:"name,omitempty"`          // Test name
	FullName      string        `json:"fullName,omitempty"`      // Full path to the test
	Status        Status        `json:"status,omitempty"`        // Status of the test execution
	StatusDetails StatusDetail  `json:"statusDetails,omitempty"` // Details about the test (for example, errors during test execution will be recorded here)
	Start         int64         `json:"start,omitempty"`         // Start of test execution
	Stop          int64         `json:"stop,omitempty"`          // End of test execution
	UUID          uuid.UUID     `json:"uuid,omitempty"`          // Unique test ID
	HistoryID     string        `json:"historyId,omitempty"`     // ID in the allure history
	TestCaseID    string        `json:"testCaseId,omitempty"`    // ID of the test case (based on the hash of the full call)
	Description   string        `json:"description,omitempty"`   // Test description
	Attachments   []*Attachment `json:"attachments,omitempty"`   // Test case attachments
	Labels        []Label       `json:"labels,omitempty"`        // Array of labels
	Links         []Link        `json:"links,omitempty"`         // Array of references
	Steps         []*Step       `json:"steps,omitempty"`         // Array of steps
	StepsQueue    NestingQueue  `json:"-"`                       // Array of nesting
	NestedSteps   []string      `json:"-"`                       // Array containing all the current uuid.UUID of the uncompleted nested steps
	Container     *Container    `json:"-"`                       // Container for Before/After Test hook
	toPrint       bool          // If false - the report will not be saved to a file
}

// StatusDetail ...
type StatusDetail struct {
	Message string `json:"message"` // Abridged version of the message
	Trace   string `json:"trace"`   // Full message
}

// NewResult Constructor Builds a new `allure.Result`. Sets the default values for the structure.
// ================================================
// |Field Value| Default                          |
// ================================================
// |UUID       | random `uuid4` value             |
// |Name       | testName from args               |
// |FullName   | fullName from args               |
// |TestCaseID | md5 hash of `Result.FullName`    |
// |HistoryID  | md5 hash from `result.TestCaseID`|
// |Container  | new empty `allure.Container`     |
// |StepsQueue | new `StepQueueue` object         |
// |Labels     | add new `allure.Language` label  |
// |Start      | allure.GetNow()                  |
// |toPrint    | `true`                           |
// ================================================
// Sets the child for the container object.
func NewResult(testName, fullName string) *Result {
	result := Result{
		UUID:       GetUUID(),
		Name:       testName,
		FullName:   fullName,
		TestCaseID: GetMD5Hash(fullName),
		Container:  NewContainer(),
		StepsQueue: NewNestingQueue(),
		toPrint:    true,
	}
	result.HistoryID = GetMD5Hash(result.TestCaseID)
	result.Labels = append(result.Labels, LanguageLabel(runtime.Version()))
	result.Container.AddChild(result.UUID)
	result.Begin()
	return &result
}

// GetLabel Returns all `allure.Label` whose `LabelType` matches the one specified in the argument.
func (result *Result) GetLabel(labelType LabelType) []Label {
	var labels []Label
	for _, label := range result.Labels {
		if label.Name == labelType.ToString() {
			labels = append(labels, label)
		}
	}
	return labels
}

// SetLabel Adds the `allure.Label` passed in arguments to the report
func (result *Result) SetLabel(label Label) {
	result.addLabel(LabelType(label.Name), label.Value)
}

// SetLabels Adds all passed in arguments `allure.Label` to the report
func (result *Result) SetLabels(labels ...Label) {
	for _, label := range labels {
		result.SetLabel(label)
	}
}

func (result *Result) addLabel(labelType LabelType, labelValue string) {
	label := NewLabel(labelType, labelValue)
	result.Labels = append(result.Labels, label)
}

// WithParentSuite Adds `allure.Label` with type `Parent` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithParentSuite(parentName string) *Result {
	if parentName == "" {
		return result
	}
	result.addLabel(ParentSuite, parentName)
	return result
}

// WithSuite Adds `allure.Label` with type `Suite` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithSuite(suiteName string) *Result {
	result.addLabel(Suite, suiteName)
	return result
}

// WithHost Adds `allure.Label` with type `Host` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithHost(hostName string) *Result {
	result.addLabel(Host, hostName)
	return result
}

// WithSubSuites Adds `allure.Label` with type `SubSuite` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithSubSuites(children ...string) *Result {
	for idx := range children {
		result.addLabel(SubSuite, children[idx])
	}
	return result
}

// WithFrameWork Adds `allure.Label` with type `Framework` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithFrameWork(framework string) *Result {
	result.addLabel(Framework, framework)
	return result
}

// WithLanguage Adds `allure.Label` with type `Language` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithLanguage(language string) *Result {
	result.addLabel(Language, language)
	return result
}

// WithThread Adds `allure.Label` with type `Thread` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithThread(thread string) *Result {
	result.addLabel(Thread, thread)
	return result
}

// WithPackage Adds `allure.Label` with type `Package` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithPackage(pkg string) *Result {
	result.addLabel(Package, pkg)
	return result
}

// WithLaunchTags Adds all Launch Tags from the global variable `ALLURE_LAUNCH_TAGS` as labels with type `Tag` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithLaunchTags() *Result {
	if tags := os.Getenv(defaultTagsEnvKey); tags != "" {
		for _, tag := range strings.Split(tags, ",") {
			result.Labels = append(result.Labels, TagLabel(strings.Trim(tag, " ")))
		}
	}
	return result
}

// Begin Sets `result.Start` as the current time
func (result *Result) Begin() {
	result.Start = GetNow()
}

// Finish Sets `result.Stop` as the current time
func (result *Result) Finish() {
	result.Stop = GetNow()
}

// Done Checks the status of the report.
// If `result.Status` is not filled in, consider the test successfully completed (no errors).
func (result *Result) Done() {
	if result.Status == "" {
		result.Status = Passed
	}
	_ = result.Print()
}

// SkipOnPrint Sets the `result.toPrint` variable to true.
func (result *Result) SkipOnPrint() {
	result.toPrint = false
}

// Print If `result.toPrint` = `true` - the method terminates without creating any files. Otherwise:
//	- Calls `result.PrintAttachments()`.
//	- Calls `result.MatchSteps()`.
//	- Saves the file `uuid4-result.json`.
//	- Calls `result.Container.Print()`
//	- Returns error (if any)
func (result *Result) Print() error {
	if !result.toPrint {
		return nil
	}
	var err error
	result.PrintAttachments()
	result.MatchSteps()
	err = result.printResult()
	if err != nil {
		return err
	}
	err = result.Container.Print()
	return err
}

// printResult marshals AllureResult to json and ioutil.WriteFile
func (result *Result) printResult() error {

	createOutputFolder(resultsPath)
	bResult, err := json.Marshal(result)
	if err != nil {
		return errors.Wrap(err, "Failed marshal result")
	}

	file := path.Join(resultsPath, fmt.Sprintf("%s-result.json", result.UUID))
	err = ioutil.WriteFile(file, bResult, fileSystemPermissionCode)

	if err != nil {
		return errors.Wrap(err, "Error write result")
	}
	return nil
}

// MatchSteps Assembles all steps of the `result.Steps` array into a tree (to keep the steps nested).
func (result *Result) MatchSteps() {
	result.Steps = matchInnerSteps(result.Steps)
}

// PrintAttachments Goes through all `result.Steps` of the report and
// for each allure.Step calls the `Step.PrintAttachments()` method.
// Then calls `Attachment.Print()` on all `allure.Attachment` of the `Result.Attachments` list.
func (result *Result) PrintAttachments() {
	for idx := range result.Steps {
		result.Steps[idx].PrintAttachments()
	}

	for idx := range result.Attachments {
		_ = result.Attachments[idx].Print()
	}
}
