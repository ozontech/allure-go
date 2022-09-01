package allure

type Step struct {
	Name        string        `json:"name,omitempty"`
	Status      Status        `json:"status,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
	Start       int64         `json:"start,omitempty"`
	Stop        int64         `json:"stop,omitempty"`
	Steps       []*Step       `json:"steps,omitempty"`
	Parameters  []*Parameter  `json:"parameters,omitempty"`
	parent      *Step
}

// NewStep Constructor. Creates a new `allure.Step` object with field values passed in arguments
// and returns a pointer to it.
func NewStep(name string, status Status, start int64, stop int64, parameters []*Parameter) *Step {
	return &Step{
		Name:       name,
		Status:     status,
		Start:      start,
		Stop:       stop,
		Parameters: parameters,
	}
}

// NewSimpleStep Constructor. Creates a `Step` object, by calling `allure.NewStep` with certain standard values
//(except for the Step name and possible parameters)
// =================================
// | Field Value| Default          |
// =================================
// | status     | `passed`         |
// | start      | `allure.GetNow()`|
// | stop       | `allure.GetNow()`|
// | parameters | ...*Parameter     |
// =================================
func NewSimpleStep(name string, parameters ...*Parameter) *Step {
	return NewStep(name, Passed, GetNow(), GetNow(), parameters)
}

// GetParent returns step's parent
func (s *Step) GetParent() *Step {
	return s.parent
}

// WithAttachments Adds to the array `Step.Attachments` passed in the argument `allure.Attachment`.
// Returns a pointer to the current Step (For Fluent Interface).
func (s *Step) WithAttachments(attachments ...*Attachment) *Step {
	s.Attachments = append(s.Attachments, attachments...)
	return s
}

// WithParameters Adds to the `Step.Parameters` array all `allure.Parameter` passed in the `params` argument.
// Returns a pointer to current Step (for Fluent Interface).
func (s *Step) WithParameters(params ...*Parameter) *Step {
	s.Parameters = append(s.Parameters, params...)
	return s
}

// WithNewParameters Accepts a list of strings, separated by commas.
// Each even-numbered string is considered a parameter name, and each odd-numbered string is
// parameter value. If an odd number of lines is passed, the last line is discarded.
// Adds to the array `Step.Parameters` all `allure.Parameter` received after conversion `kv`.
// Returns pointer to the current Step (for Fluent Interface).
func (s *Step) WithNewParameters(kv ...interface{}) *Step {
	s.Parameters = append(s.Parameters, NewParameters(kv...)...)
	return s
}

// Passed Puts `Step.Status` = `passed`.
// Returns a pointer to the current Step (for Fluent Interface).
func (s *Step) Passed() *Step {
	s.Status = Passed
	return s
}

// Failed Puts `Step.Status` = `failed`.
// Returns a pointer to the current Step (for Fluent Interface).
func (s *Step) Failed() *Step {
	s.Status = Failed
	return s
}

// Skipped Puts `Step.Status` = `skipped`.
// Returns a pointer to the current Step (for Fluent Interface).
func (s *Step) Skipped() *Step {
	s.Status = Skipped
	return s
}

// Broken Puts `Step.Status` = `broken`.
// Returns a pointer to the current Step (for Fluent Interface).
func (s *Step) Broken() *Step {
	s.Status = Broken
	return s
}

// Begin Puts `Step.Start` = `GetNow()`.
// Returns a pointer to the current Step (for Fluent Interface).
func (s *Step) Begin() *Step {
	s.Start = GetNow()
	return s
}

// Finish Puts `Step.Start` = `GetNow()`.
// Returns a pointer to the current Step (for Fluent Interface).
func (s *Step) Finish() *Step {
	s.Stop = GetNow()
	return s
}

// WithParent Sets the step `parentUUID` as the UUID of the step passed in the argument `parent`.
// Returns a pointer to the current step (For Fluent Interface).
func (s *Step) WithParent(parent *Step) *Step {
	parent.Steps = append(parent.Steps, s)
	s.parent = parent
	return s
}

// WithChild Sets the step `parentUUID` as the UUID for the step passed in the argument `child`.
// Returns a pointer to the current step (For Fluent Interface).
func (s *Step) WithChild(child *Step) *Step {
	child.WithParent(s)
	return s
}

// PrintAttachments Goes through all `allure.Attachments` of the `Step.Attachments`
// array and calls `Print()` method on `allure.Attachment`.
func (s *Step) PrintAttachments() {
	for _, a := range s.Attachments {
		_ = a.Print()
	}
	if s.Steps == nil {
		return
	}
	for _, step := range s.Steps {
		step.PrintAttachments()
	}
}
