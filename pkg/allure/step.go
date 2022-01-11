package allure

type IStep interface {
	WithAttachments
	WithTimer
	Attachment(*Attachment)
	AddParameter(Parameter)
	AddParameters(...Parameter)
	AddNewParameter(string, string)
	AddNewParameters(...string)

	WithAttachment(*Attachment) *Step
	WithParameter(Parameter) *Step
	WithNewParameter(string, string) *Step
	WithParameters(...Parameter) *Step
	WithNewParameters(...string) *Step
	WithParent(*Step) *Step
	WithStart() *Step
	WithStop() *Step

	Passed() *Step
	Failed() *Step
	Skipped() *Step
}

// Step is an implementation of the Step entity used by Allure to define and describe test steps. Steps
// can be nested, have a status (successful, failed, skipped, broken),
// can contain attachments and parameters, and have an execution time.
// Allure-testify offers great possibilities for creating and modifying the steps allowing
// to collect beautiful and clear reports on test execution.
// It is highly recommended that you use steps when describing your test script.
// This allows your tests to be clearer and more informative not only in reports, but also in code.
// However, overuse is not recommended, as misuse can turn your tests code into a canvas,
// consisting only of steps.
type Step struct {
	Name        string        `json:"name,omitempty"`
	Status      Status        `json:"status,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
	Start       int64         `json:"start,omitempty"`
	Stop        int64         `json:"stop,omitempty"`
	Steps       []*Step       `json:"steps,omitempty"`
	Parameters  []Parameter   `json:"parameters,omitempty"`
	Parent      string        `json:"-"`
	uuid        string
}

// Status is step's Status info
type Status string

// Status constants
const (
	Passed  Status = "passed"
	Failed  Status = "failed"
	Skipped Status = "skipped"
	Broken  Status = "broken"
	Unknown Status = "unknown"
)

// NewStep Constructor. Creates a new `allure.Step` object with field values passed in arguments
// and returns a pointer to it.
func NewStep(name string, status Status, start int64, stop int64, parameters []Parameter) *Step {
	return &Step{
		Name:       name,
		Status:     status,
		Start:      start,
		Stop:       stop,
		Parameters: parameters,
		uuid:       GetUUID().String(),
	}
}

// NewSimpleStep Constructor. Creates a `Step` object, by calling `allure.NewStep` with certain standard values
//(except for the step name)
// =================================
// | Field Value| Default          |
// =================================
// | status     | `passed`         |
// | start      | `allure.GetNow()`|
// | stop       | `allure.GetNow()`|
// | parameters | nil              |
// =================================
func NewSimpleStep(name string) *Step {
	return NewStep(name, Passed, GetNow(), GetNow(), nil)
}

// NewSimpleInnerStep Calls `allure.NewSimpleStep`,
// then sets `parentUUID` as the UUID of the step that have been passed in the `parent` argument
func NewSimpleInnerStep(name string, parent *Step) *Step {
	step := NewSimpleStep(name)
	step.Parent = parent.GetUUID()
	return step
}

// NewStepWithStart Constructor. Creates a `Step` object, by calling `allure.NewStep` with defined standard values
// (except for the step name).
//
// Unlike `allure.NewSimpleStep`, does not fill Stop field.
// =================================
// | Field Value| Default          |
// =================================
// | status     | `passed`         |
// | start      | `allure.GetNow()`|
// | parameters | nil              |
// =================================
func NewStepWithStart(name string) *Step {
	return &Step{Name: name, Start: GetNow(), uuid: GetUUID().String(), Status: Passed}
}

// GetUUID returns step's UUID
func (st *Step) GetUUID() string {
	return st.uuid
}

// Attachment Adds to the array `Step.Attachments` passed in the argument `allure.Attachment`.
func (st *Step) Attachment(attachment *Attachment) {
	st.Attachments = append(st.Attachments, attachment)
}

// AddParameter Adds to the array `Step.Parameters` passed in the argument `param`.
func (st *Step) AddParameter(param Parameter) {
	st.Parameters = append(st.Parameters, param)
}

// AddNewParameter Creates a new `allure.Parameters` from the passed `key` and `value` arguments
// and adds them to the `Step.Parameters` array.
func (st *Step) AddNewParameter(key, value string) {
	st.AddParameter(NewParameter(key, value))
}

// AddParameters Adds to the `Step.Parameters` array all `allure.Parameter` passed in the `params` argument.
func (st *Step) AddParameters(params ...Parameter) {
	st.Parameters = append(st.Parameters, params...)
}

// AddNewParameters Accepts a list of strings, separated by commas.
// Each even-numbered string is considered a parameter name, and each odd-numbered string is
// parameter value. If odd number of lines is passed, the last line is discarded.
// Adds to the array `Step.Parameters` all `allure.Parameter` received after conversion `kv`.
func (st *Step) AddNewParameters(kv ...string) {
	st.AddParameters(NewParameters(kv...)...)
}

// WithAttachment Adds to the array `Step.Attachments` passed in the argument `allure.Attachment`.
// Returns a pointer to the current step (For Fluent Interface).
func (st *Step) WithAttachment(attachment *Attachment) *Step {
	st.Attachment(attachment)
	return st
}

// WithParameter Adds to the array `Step.Parameters` passed in the argument `param`.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) WithParameter(param Parameter) *Step {
	st.AddParameter(param)
	return st
}

// WithParameters Adds to the `Step.Parameters` array all `allure.Parameter` passed in the `params` argument.
// Returns a pointer to current step (for Fluent Interface).
func (st *Step) WithParameters(params ...Parameter) *Step {
	st.AddParameters(params...)
	return st
}

// WithNewParameter Creates a new `allure.Parameters` from the passed `key` and `value` arguments and adds them to the `Step.Parameters` array.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) WithNewParameter(key, value string) *Step {
	st.AddNewParameter(key, value)
	return st
}

// WithNewParameters Accepts a list of strings, separated by commas.
// Each even-numbered string is considered a parameter name, and each odd-numbered string is
// parameter value. If an odd number of lines is passed, the last line is discarded.
// Adds to the array `Step.Parameters` all `allure.Parameter` received after conversion `kv`.
// Returns pointer to the current step (for Fluent Interface).
func (st *Step) WithNewParameters(kv ...string) *Step {
	st.AddNewParameters(kv...)
	return st
}

// WithParent Sets the step `parentUUID` as the UUID of the step passed in the argument `parent`.
// Returns a pointer to the current step (For Fluent Interface).
func (st *Step) WithParent(parent *Step) *Step {
	st.Parent = parent.GetUUID()
	return st
}

// WithStart Puts `Step.Start` = `GetNow()`.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) WithStart() *Step {
	st.Begin()
	return st
}

// WithStop Puts `Step.Stop` = `GetNow()`.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) WithStop() *Step {
	st.Finish()
	return st
}

// Passed Puts `Step.Status` = `passed`.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) Passed() *Step {
	st.Status = Passed
	return st
}

// Failed Puts `Step.Status` = `failed`.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) Failed() *Step {
	st.Status = Failed
	return st
}

// Skipped Puts `Step.Status` = `skipped`.
// Returns a pointer to the current step (for Fluent Interface).
func (st *Step) Skipped() *Step {
	st.Status = Skipped
	return st
}

// Begin Puts `Step.Start` = `GetNow()`.
func (st *Step) Begin() {
	st.Start = GetNow()
}

// Finish Puts `Step.Stop` = `GetNow()`.
func (st *Step) Finish() {
	st.Stop = GetNow()
}

// PrintAttachments Goes through all `allure.Attachments` of the `Step.Attachments`
// array and calls `Print()` method on `allure.Attachment`.
func (st *Step) PrintAttachments() {
	for idx := range st.Attachments {
		_ = st.Attachments[idx].Print()
	}
}
