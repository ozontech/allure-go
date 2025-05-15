package allure

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// Container This is an implementation of the `Container` entity used by Allure to handle TestSetup and TestTeardown hooks.
// The list of container-dependent tests is contained in the `Container.Children` array.
// Note: For Before/After Test hooks, the Container.Children array will contain one element (one container per test).
// For Before/After Suite hooks the Container.Children array will contain UUIDs of all tests for which the hook was executed.
type Container struct {
	UUID     uuid.UUID   `json:"uuid,omitempty"`     // Unique identifier of the container
	Children []uuid.UUID `json:"children,omitempty"` // UUID array containing all reports referring to the container
	Befores  []*Step     `json:"befores,omitempty"`  // Array of pointers to allure.Step in Test Setup
	Afters   []*Step     `json:"afters,omitempty"`   // Array of pointers to allure.Step in Test TearDown
	Start    int64       `json:"start,omitempty"`    // Start time of the container
	Stop     int64       `json:"stop,omitempty"`     // Stop time of the container
}

// NewContainer - Constructor. Builds and returns a new `allure.Container` object.
func NewContainer() *Container {
	return &Container{
		UUID: uuid.New(),
	}
}

// AddChild Adds a new child to the Container.Children array.
func (container *Container) AddChild(child uuid.UUID) {
	container.Children = append(container.Children, child)
}

// IsEmpty Returns `true` if arrays Container.Befores and Container.Afters are empty.
func (container *Container) IsEmpty() bool {
	return len(container.Befores) == 0 && len(container.Afters) == 0
}

// Print Checks the file with the function Container.IsEmpty:
//
//  1. If the container is empty, execution of the function completes without error.
//
//  2. If the container contains steps
//     2.1. Call Container.PrintAttachments()
//     2.2. Serializes the file into `uuid4-container.json`.
//     2.3. Creates a file in the file system in the output folder (`$ALLURE_OUTPUT_PATH`/`$ALLURE_OUTPUT_FOLDER`). If there is an error during
//
// If error occurs during execution - returns it
func (container *Container) Print() error {
	if !container.IsEmpty() {
		container.PrintAttachments()

		return container.printContainer()
	}

	return nil
}

// PrintAttachments It goes through all Container.Befores and Container.Afters
// of the Container and calls the Container.PrintAttachments() method at each allure.Step.
func (container *Container) PrintAttachments() {
	for _, step := range container.Befores {
		step.PrintAttachments()
	}

	for _, step := range container.Afters {
		step.PrintAttachments()
	}
}

// Begin Sets `Container.Start` = allure.GetNow()
func (container *Container) Begin() {
	container.Start = GetNow()
}

// Finish Sets Container.Stop = allure.GetNow()
func (container *Container) Finish() {
	container.Stop = GetNow()
}

// Done calls Finish and Print
func (container *Container) Done() error {
	container.Finish()

	return container.Print()
}

// ToJSON marshal allure.Result to json file
//
// Deprecated: use [json.Marshal] instead
func (container *Container) ToJSON() ([]byte, error) {
	return json.Marshal(container)
}

// Print prints all attachments of [Container.Befores] and [Container.Afters]
// after that marshals [Container] and [os.WriteFile]
func (container *Container) printContainer() error {
	bResult, err := json.Marshal(container)
	if err != nil {
		return errors.Wrap(err, "Failed marshal Result")
	}

	err = NewFileManager().CreateFile(container.UUID.String()+"-container.json", bResult)
	if err != nil {
		return errors.Wrap(err, "Error write Result")
	}

	return nil
}
