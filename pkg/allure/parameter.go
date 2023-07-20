package allure

import (
	//"encoding/json"
	"fmt"
	"strings"
)

// Parameter is an implementation of the Parameter entity,
// which Allure uses as additional information describing the test Step
// (for example - request host or server address)
type Parameter struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// NewParameter Constructor. Builds and returns a new `Parameter` object,
// using `name` as the parameter name and `value`, as the value.
func NewParameter(name string, value ...interface{}) *Parameter {
	val := trimBrackets(messageFromMsgAndArgs(value))
	return &Parameter{
		Name:  name,
		Value: val,
	}
}

// NewParameters Constructor. Accepts a list of strings, separated by commas.
// Each even string is considered a parameter name, and each  odd-value of the parameter.
// If an odd number of lines is passed, the last line is discarded.
// Returns the list of parameters received after processing the passed list.
func NewParameters(kv ...interface{}) []*Parameter {
	if len(kv)%2 != 0 {
		kv = kv[:len(kv)-1]
	}
	result := make([]*Parameter, len(kv)/2)
	for i := 0; i < len(kv); i += 2 {
		val := trimBrackets(messageFromMsgAndArgs(kv[i+1]))
		result[i/2] = NewParameter(messageFromMsgAndArgs(kv[i]), val)
	}
	return result
}

// GetValue returns param value as string
func (p *Parameter) GetValue() string {
	return strings.Trim(fmt.Sprintf("%s", p.Value), "\"")
}

func trimBrackets(val string) string {
	if strings.HasSuffix(val, "]") && strings.HasPrefix(val, "[") {
		return strings.TrimSuffix(strings.TrimPrefix(val, "["), "]")
	}
	return val
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}
	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		switch m := msg.(type) {
		case string:
			return m
		case int:
			return fmt.Sprintf("%d", m)
		default:
			return fmt.Sprintf("%+v", m)
		}
	}
	if len(msgAndArgs) > 1 {
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
	return ""
}
