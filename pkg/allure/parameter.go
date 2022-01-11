package allure

type IParameter interface {
}

// Parameter is an implementation of the Parameter entity,
// which Allure uses as additional information describing the test step
// (for example - request host or server address)
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// NewParameter Constructor. Builds and returns a new `Parameter` object,
// using `name` as the parameter name and `value`, as the value.
func NewParameter(name string, value string) Parameter {
	return Parameter{
		Name:  name,
		Value: value,
	}
}

// NewParameters Constructor. Accepts a list of strings, separated by commas.
// Each even string is considered a parameter name, and each  odd-value of the parameter.
// If an odd number of lines is passed, the last line is discarded.
// Returns the list of parameters received after processing the passed list.
func NewParameters(kv ...string) []Parameter {
	if len(kv)%2 != 0 {
		kv = kv[:len(kv)-1]
	}
	result := make([]Parameter, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		result[i/2] = NewParameter(kv[i], kv[i+1])
	}
	return result
}
