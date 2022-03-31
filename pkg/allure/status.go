package allure

// Status is Step's Status info
type Status string

// Status constants
const (
	Passed  Status = "passed"
	Failed  Status = "failed"
	Skipped Status = "skipped"
	Broken  Status = "broken"
	Unknown Status = "unknown"
)
