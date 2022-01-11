package allure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type IContainer interface {
	Printable
	WithAttachments
	WithSteps
	WithTimer
	AddChild(childUUID uuid.UUID)
	IsEmpty() bool
}

// Container This is an implementation of the `Container` entity used by Allure to handle TestSetup and TestTeardown hooks.
// The list of container-dependent tests is contained in the `Container.Children` array.
// Note: For Before/After Test hooks, the Container.Children array will contain one element (one container per test).
// For Before/After Suite hooks the Container.Children array will contain UUIDs of all tests for which the hook was executed.
type Container struct {
	UUID         uuid.UUID    `json:"uuid,omitempty"`     // Unique identifier of the container
	Children     []uuid.UUID  `json:"children,omitempty"` // UUID array containing all reports referring to the container
	Befores      []*Step      `json:"befores,omitempty"`  // Array of pointers to allure.Step in Test Setup
	Afters       []*Step      `json:"afters,omitempty"`   // Array of pointers to allure.Step in Test TearDown
	Start        int64        `json:"start,omitempty"`    // Start time of the container
	Stop         int64        `json:"stop,omitempty"`     // Stop time of the container
	BeforesQueue NestingQueue `json:"-"`                  // Queue of nested allure.Step in Befores
	AftersQueue  NestingQueue `json:"-"`                  // Queue of nested allure.Step in Afters
}

// NewContainer - Constructor. Builds and returns a new `allure.Container` object.
func NewContainer() *Container {
	return &Container{
		UUID:         GetUUID(),
		BeforesQueue: NewNestingQueue(),
		AftersQueue:  NewNestingQueue(),
	}
}

// AddChild Adds a new child to the Container.Children array.
func (container *Container) AddChild(childUUID uuid.UUID) {
	container.Children = append(container.Children, childUUID)
}

// IsEmpty Returns `true` if arrays Container.Befores and Container.Afters are empty.
func (container *Container) IsEmpty() bool {
	return (container.Befores == nil || len(container.Befores) == 0) && (container.Afters == nil || len(container.Afters) == 0)
}

// Print Checks the file with the function Container.IsEmpty:
// 1) if the container is empty, execution of the function completes without error.
// 2) If the container contains steps
//    1) Call Container.PrintAttachments()
//    2) Call Container.MatchSteps()
//    3) Serializes the file into `uuid4-container.json`.
//    4) Creates a file in the file system in the output folder (`$ALLURE_OUTPUT_PATH`/`$ALLURE_OUTPUT_FOLDER`). If there is an error during
//       error occurs during execution - returns it
func (container *Container) Print() error {
	var err error
	if !container.IsEmpty() {
		container.PrintAttachments()
		container.MatchSteps()
		err = container.printContainer()
	}
	return err
}

// PrintAttachments It goes through all Container.Befores and Container.Afters
// of the Container and calls the Container.PrintAttachments() method at each allure.Step.
func (container *Container) PrintAttachments() {
	for idx := range container.Befores {
		container.Befores[idx].PrintAttachments()
	}

	for idx := range container.Afters {
		container.Afters[idx].PrintAttachments()
	}
}

// MatchSteps Assembles steps into a tree for both Container.Befores and Container.Afters arrays.
func (container *Container) MatchSteps() {
	container.Befores = matchInnerSteps(container.Befores)
	container.Afters = matchInnerSteps(container.Afters)
}

// Begin Sets `Container.Start` = allure.GetNow()
func (container *Container) Begin() {
	container.Start = GetNow()
}

// Finish Sets Container.Stop = allure.GetNow()
func (container *Container) Finish() {
	container.Stop = GetNow()
}

// Print prints all attachments of Container.Befores and Container.Afters
// after that marshals Container and ioutil.WriteFile
func (container *Container) printContainer() error {
	createOutputFolder(resultsPath)
	bResult, err := json.Marshal(container)
	if err != nil {
		return errors.Wrap(err, "Failed marshal result")
	}

	for _, step := range container.Befores {
		for _, attachment := range step.Attachments {
			_ = attachment.Print()
		}
	}

	for _, step := range container.Afters {
		for _, attachment := range step.Attachments {
			_ = attachment.Print()
		}
	}

	file := path.Join(resultsPath, fmt.Sprintf("%s-container.json", container.UUID))
	err = ioutil.WriteFile(file, bResult, fileSystemPermissionCode)

	if err != nil {
		return errors.Wrap(err, "Error write result")
	}
	return nil
}
