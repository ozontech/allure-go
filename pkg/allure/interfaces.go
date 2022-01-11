package allure

// Printable interface indicates that the object has a file equivalent in allure.
// The interface is implemented by:
//	allure.Attachment;
//	allure.Result;
//	allure.Container.
type Printable interface {
	// Print creates a file in the file system.
	Print() error
}

// WithAttachments interface indicates that the object can have allure.Attachment.
// Since allure.Attachment implements Printable,
// it should be possible to conveniently print all allure.Attachment entities at once.
// Implemented by:
//	allure.Result;
//	allure.Container;
//	allure.Step.
type WithAttachments interface {
	// PrintAttachments calls Attachment.Print() method for each attachment at "Attachments" array
	PrintAttachments()
}

// WithSteps Denotes that the object has allure.Steps. The current implementation of nesting is that steps, is a list,
// whose order is guaranteed by the ecosystem go. Accordingly, each child lies right after the parent. To
// turn a linear list into a tree - implement method MatchSteps().
// Implement:
//	allure.Result;
//	allure.Container;
//	allure.Step.
type WithSteps interface {
	// MatchSteps turns the list of steps into a tree.
	MatchSteps()
}

// WithTimer Denotes that the object has some dimension of time.
type WithTimer interface {
	// Begin denotes the start of execution
	Begin()
	// Finish denotes the end of execution
	Finish()
}
