package allure

// StatusDetail ...
type StatusDetail struct {
	Message string `json:"message"` // Abridged version of the message
	Trace   string `json:"trace"`   // Full message
}
