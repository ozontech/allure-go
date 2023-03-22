package allure

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Result is an implementation of the Result entity used by Allure to store information about the test. It contains
// information about the test name, applications, description, status, references, labels,
// steps, containers, and time of the test execution.
type Result struct {
	Name          string        `json:"name,omitempty"`          // Test name
	FullName      string        `json:"fullName,omitempty"`      // Full path to the test
	Stage         string        `json:"stage,omitempty"`         // Stage of test execution
	Status        Status        `json:"status,omitempty"`        // Status of the test execution
	StatusDetails StatusDetail  `json:"statusDetails,omitempty"` // Details about the test (for example, errors during test execution will be recorded here)
	Start         int64         `json:"start,omitempty"`         // Start of test execution
	Stop          int64         `json:"stop,omitempty"`          // End of test execution
	UUID          uuid.UUID     `json:"uuid,omitempty"`          // Unique test ID
	HistoryID     string        `json:"historyId,omitempty"`     // ID in the allure history
	TestCaseID    string        `json:"testCaseId,omitempty"`    // ID of the test case (based on the hash of the full call)
	Description   string        `json:"description,omitempty"`   // Test description
	Attachments   []*Attachment `json:"attachments,omitempty"`   // Test case attachments
	Parameters    []*Parameter  `json:"parameters,omitempty"`    // Test case parameters
	Labels        []*Label      `json:"labels,omitempty"`        // Array of labels
	Links         []*Link       `json:"links,omitempty"`         // Array of references
	Steps         []*Step       `json:"steps,omitempty"`         // Array of steps
	ToPrint       bool          `json:"-"`                       // If false - the report will not be saved to a file
}

// NewResult Constructor Builds a new `allure.Result`. Sets the default values for the structure.
// ================================================
// |Field Value| Default                          |
// ================================================
// |UUID       | random `uuid4` value             |
// |Name       | testName from args               |
// |FullName   | fullName from args               |
// |TestCaseID | md5 hash of `Result.FullName`    |
// |HistoryID  | md5 hash from `Result.TestCaseID`|
// |Container  | new empty `allure.Container`     |
// |Labels     | add new `allure.Language` label  |
// |Start      | allure.GetNow()                  |
// |ToPrint    | `true`                           |
// ================================================
// Sets the child for the container object.
func NewResult(testName, fullName string) *Result {
	result := Result{
		UUID:       getUUID(),
		Name:       testName,
		FullName:   fullName,
		TestCaseID: getMD5Hash(fullName),
		ToPrint:    true,
	}
	result.HistoryID = getMD5Hash(result.TestCaseID)
	result.Labels = append(result.Labels, LanguageLabel(runtime.Version()))
	result.Begin()
	return &result
}

func (result *Result) SetStatusMessage(msg string) {
	result.StatusDetails.Message = msg
}

func (result *Result) GetStatusMessage() string {
	return result.StatusDetails.Message
}

func (result *Result) SetStatusTrace(trace string) {
	result.StatusDetails.Trace = trace
}

func (result *Result) GetStatusTrace() string {
	return result.StatusDetails.Trace
}

func (result *Result) addLabel(labelType LabelType, labelValue string) {
	label := NewLabel(labelType, labelValue)
	result.Labels = append(result.Labels, label)
}

// AddLabel Adds all passed in arguments `allure.Label` to the report
func (result *Result) AddLabel(labels ...*Label) {
	result.Labels = append(result.Labels, labels...)
}

// GetFirstLabel returns first label in labels list and true if something have been found, return false if nothing was found
func (result *Result) GetFirstLabel(labelType LabelType) (label *Label, ok bool) {
	if labels := result.GetLabels(labelType); len(labels) > 0 {
		return labels[0], true
	}
	return
}

// GetLabels Returns all `allure.Label` whose `LabelType` matches the one specified in the argument.
func (result *Result) GetLabels(labelType LabelType) []*Label {
	labels := make([]*Label, 0)
	for _, label := range result.Labels {
		if label.Name == labelType.ToString() {
			labels = append(labels, label)
		}
	}
	return labels
}

// SetNewLabelMap Adds all passed in arguments `allure.Label` to the report
func (result *Result) SetNewLabelMap(kv map[LabelType]string) {
	var labels []*Label
	for k, v := range kv {
		labels = append(labels, NewLabel(k, v))
	}
	result.AddLabel(labels...)
}

// WithStage sets Stage field to result
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithStage(stage string) *Result {
	result.Stage = stage
	return result
}

// WithParentSuite Adds `allure.Label` with type `Parent` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithParentSuite(parentName string) *Result {
	if parentName == "" {
		return result
	}
	result.ReplaceNewLabel(ParentSuite, parentName)
	return result
}

// WithSuite Adds `allure.Label` with type `Suite` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithSuite(suiteName string) *Result {
	result.ReplaceNewLabel(Suite, suiteName)
	return result
}

// WithHost Adds `allure.Label` with type `Host` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithHost(hostName string) *Result {
	result.ReplaceNewLabel(Host, hostName)
	return result
}

// WithSubSuites Adds `allure.Label` with type `SubSuite` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithSubSuites(children ...string) *Result {
	for _, child := range children {
		result.addLabel(SubSuite, child)
	}
	return result
}

// WithFrameWork Adds `allure.Label` with type `Framework` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithFrameWork(framework string) *Result {
	result.ReplaceNewLabel(Framework, framework)
	return result
}

// WithLanguage Adds `allure.Label` with type `Language` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithLanguage(language string) *Result {
	result.ReplaceNewLabel(Language, language)
	return result
}

// WithThread Adds `allure.Label` with type `Thread` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithThread(thread string) *Result {
	result.ReplaceNewLabel(Thread, thread)
	return result
}

// WithPackage Adds `allure.Label` with type `Package` to the report.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithPackage(pkg string) *Result {
	result.ReplaceNewLabel(Package, pkg)
	return result
}

// WithLabels Adds an array of `allure.Label`.
// Returns a pointer to the current `allure.Result` (for Fluent Interface).
func (result *Result) WithLabels(label ...*Label) *Result {
	result.AddLabel(label...)
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

// Begin Sets `Result.Start` as the current time
func (result *Result) Begin() *Result {
	result.Start = GetNow()
	return result
}

// Finish Sets `Result.Stop` as the current time
func (result *Result) Finish() *Result {
	result.Stop = GetNow()
	return result
}

// SkipOnPrint Sets the `Result.ToPrint` variable to false.
func (result *Result) SkipOnPrint() {
	result.ToPrint = false
}

// Print If `Result.ToPrint` = `true` - the method terminates without creating any files. Otherwise:
//	- Calls `Result.PrintAttachments()`.
//	- Saves the file `uuid4-Result.json`.
//	- Calls `Result.Container.Print()`
//	- Returns error (if any)
func (result *Result) Print() error {
	if !result.ToPrint {
		return nil
	}
	result.PrintAttachments()
	return result.printResult()
}

// printResult marshals allure.Result to json and do ioutil.WriteFile
func (result *Result) printResult() error {
	bResult, err := json.Marshal(result)
	if err != nil {
		return errors.Wrap(err, "Failed marshal Result")
	}

	err = NewFileManager().CreateFile(fmt.Sprintf("%s-result.json", result.UUID), bResult)
	if err != nil {
		return errors.Wrap(err, "Cannot save Result")
	}

	return nil
}

// PrintAttachments Goes through all `Result.Steps` of the report and
// for each allure.Step calls the `Step.PrintAttachments()` method.
// Then calls `Attachment.Print()` on all `allure.Attachment` of the `Result.Attachments` list.
func (result *Result) PrintAttachments() {
	for _, step := range result.Steps {
		step.PrintAttachments()
	}

	for _, attachment := range result.Attachments {
		_ = attachment.Print()
	}
}

// Done Checks the status of the report.
// If `Result.Status` is not filled in, consider the test successfully completed (no errors).
// After that - it calls Finish() and Print() methods.
func (result *Result) Done() error {
	if result.Status == "" {
		result.Status = Passed
	}
	result.Finish()
	return result.Print()
}

// ReplaceNewLabel creates new label and replaces it in the allure.Result object by label's name
func (result *Result) ReplaceNewLabel(name LabelType, value string) {
	result.ReplaceLabel(NewLabel(name, value))
}

// ReplaceLabel replaces label in the allure.Result object by label's name
func (result *Result) ReplaceLabel(label *Label) {
	for _, l := range result.Labels {
		if label.Name == l.Name {
			l.Value = label.Value
			return
		}
	}
	result.Labels = append(result.Labels, label)
}

// ToJSON marshall allure.Result to json file
func (result *Result) ToJSON() ([]byte, error) {
	return json.Marshal(result)
}

// getMD5Hash ...
func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
