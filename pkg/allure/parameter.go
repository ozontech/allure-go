package allure

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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

// TODO: remove this in v2
//
// See explanation in [Parameter.UnmarshalJSON]
type parameterValue struct {
	inner interface{}
}

func (p parameterValue) Inner() interface{} {
	switch i := p.inner.(type) {
	case string, float64, int64, bool:
		return p.inner
	case []parameterValue:
		values := make([]interface{}, 0, len(i))

		for _, v := range i {
			values = append(values, v.Inner())
		}

		return values
	case map[string]parameterValue:
		values := make(map[string]interface{}, len(i))

		for k, v := range i {
			values[k] = v.Inner()
		}

		return values
	default:
		// unreachable
		return nil
	}
}

func (p *parameterValue) UnmarshalJSON(data []byte) error {
	var valueStr string

	errStr := sonic.Unmarshal(data, &valueStr)
	if errStr == nil {
		p.inner = valueStr

		return nil
	}

	var valueBool bool

	errBool := sonic.Unmarshal(data, &valueBool)
	if errBool == nil {
		p.inner = valueBool

		return nil
	}

	var valueNum json.Number

	errNum := sonic.Unmarshal(data, &valueNum)
	if errNum == nil {
		if n, err := valueNum.Int64(); err == nil {
			p.inner = n

			return nil
		}

		if n, err := valueNum.Float64(); err == nil {
			p.inner = n

			return nil
		}

		// possibly unreachable
		p.inner = valueNum.String()

		return nil
	}

	var valueMap map[string]parameterValue

	errMap := sonic.Unmarshal(data, &valueMap)
	if errMap == nil {
		p.inner = valueMap

		return nil
	}

	var valueSlice []parameterValue

	errSlice := sonic.Unmarshal(data, &valueSlice)
	if errSlice == nil {
		p.inner = valueSlice

		return nil
	}

	return &joinError{
		errs: []error{errStr, errBool, errNum, errMap, errSlice},
	}
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
		Name  string         `json:"name"`
		Value parameterValue `json:"value"`
	}

	if err := sonic.Unmarshal(data, &aux); err != nil {
		return err
	}

	*p = Parameter{
		Name:  aux.Name,
		Value: aux.Value.Inner(),
	}

	return nil
}

func (p *Parameter) MarshalJSON() ([]byte, error) {
	var raw json.RawMessage

	switch v := p.Value.(type) {
	case proto.Message:
		res, err := protojson.MarshalOptions{
			AllowPartial:      true,
			EmitDefaultValues: true,
			EmitUnpopulated:   true,
		}.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("protojson marshal: %w", err)
		}
		raw = res

	default:
		res, err := sonic.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("json marshal: %w", err)
		}
		raw = res
	}

	aux := struct {
		Name  string          `json:"name"`
		Value json.RawMessage `json:"value"`
	}{
		Name:  p.Name,
		Value: raw,
	}

	return sonic.Marshal(aux)
}

func trimBrackets(val string) string {
	if strings.HasSuffix(val, "]") && strings.HasPrefix(val, "[") {
		return strings.TrimSuffix(strings.TrimPrefix(val, "["), "]")
	}

	return val
}

func messageFromMsgAndArgs(msgAndArgs ...interface{}) string {
	switch len(msgAndArgs) {
	case 0:
		return ""
	case 1:
		msg := msgAndArgs[0]

		switch m := msg.(type) {
		case string:
			return m
		case int:
			return fmt.Sprintf("%d", m)
		default:
			return fmt.Sprintf("%+v", m)
		}
	default:
		return fmt.Sprintf(msgAndArgs[0].(string), msgAndArgs[1:]...)
	}
}
