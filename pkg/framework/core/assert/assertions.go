package assert

import (
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
)

// TestingT is an interface wrapper around *testing.T
type TestingT interface {
	Errorf(format string, args ...interface{})
}

type tHelper interface {
	Helper()
}

// JSONContains asserts that expected JSON contains fields and values of actual JSON which can be bigger.
//
//  assert.JSONContains(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world", "foobar": 1}`)
func JSONContains(t TestingT, expected string, actual string, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	var expectedJSONAsInterface, actualJSONAsInterface interface{}

	if err := json.Unmarshal([]byte(expected), &expectedJSONAsInterface); err != nil {
		return assert.Fail(t, fmt.Sprintf("Expected value ('%s') is not valid json.\nJSON parsing error: '%s'", expected, err.Error()), msgAndArgs...)
	}

	if err := json.Unmarshal([]byte(actual), &actualJSONAsInterface); err != nil {
		return assert.Fail(t, fmt.Sprintf("Input ('%s') needs to be valid json.\nJSON parsing error: '%s'", actual, err.Error()), msgAndArgs...)
	}

	return assert.Equal(t, expectedJSONAsInterface, extractActualJSON(t, expectedJSONAsInterface, actualJSONAsInterface), msgAndArgs...)
}

func extractActualJSON(t TestingT, expected interface{}, actual interface{}) interface{} {
	switch expected.(type) {
	case []interface{}:
		exp := expected.([]interface{})
		act, ok := actual.([]interface{})
		if !ok {
			assert.Fail(t, "Unexpected type of actual JSON element")
			return nil
		}
		if len(exp) > len(act) {
			assert.Fail(t, "Expected slice bigger than actual JSON element")
			return nil
		}
		result := make([]interface{}, 0, len(exp))

		for i, ev := range exp {
			result = append(result, extractActualJSON(t, ev, act[i]))
		}
		return result

	case map[string]interface{}:
		result := make(map[string]interface{}, len(expected.(map[string]interface{})))
		act, ok := actual.(map[string]interface{})
		if !ok {
			assert.Fail(t, "Unexpected type of actual JSON element")
			return nil
		}

		for k, ev := range expected.(map[string]interface{}) { // use type assertion to loop over map[string]interface{}
			result[k] = extractActualJSON(t, ev, act[k])
		}
		return result
	}
	return actual
}
