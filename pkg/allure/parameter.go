package allure

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"strconv"
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
	s := fmt.Sprint(p.Value)

	unquoted, err := strconv.Unquote(s)
	if err != nil {
		return s
	}

	return unquoted
}

func (p *Parameter) UnmarshalJSON(data []byte) error {
	// Since [Parameter.Value] is interface{} json will unmarshal any number as float64.
	// This might lead to unexpected behaviour, such as 83294782375982 unmarshalled as 8.3294782375982e+13
	// While these values are logically the same, when converted to string the later will result 8.3294782375982e+13 (with exponent)
	//
	// We could've checked if this float is convertable to int in [Parameter.GetValue], unless we can't - int(float32(99999999)) == 100000000
	// See: https://stackoverflow.com/questions/65417925/golang-weird-behavior-when-converting-float-to-int
	//
	// So using custom json unmarshalling seems like the only choice if we want to preserve backwards compatibility.
	//
	// TODO: refactor this in v2

	var aux struct {
		Name  string          `json:"name"`
		Value json.RawMessage `json:"value"`
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var (
		valueStr string
		valueNum json.Number
	)

	errStr := json.Unmarshal(aux.Value, &valueStr)
	if errStr == nil {
		*p = Parameter{
			Name:  aux.Name,
			Value: valueStr,
		}

		return nil
	}

	errNum := json.Unmarshal(aux.Value, &valueNum)
	if errNum == nil {
		if n, err := valueNum.Int64(); err == nil {
			*p = Parameter{
				Name:  aux.Name,
				Value: n,
			}

			return nil
		}

		if n, err := valueNum.Float64(); err == nil {
			*p = Parameter{
				Name:  aux.Name,
				Value: n,
			}

			return nil
		}

		// possibly unreachable
		*p = Parameter{
			Name:  aux.Name,
			Value: valueNum.String(),
		}

		return nil
	}

	// possibly unreachable
	return fmt.Errorf("unmarshal value: %w, %w", errStr, errNum)
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
