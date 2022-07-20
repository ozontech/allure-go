package assert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSONContains_EqualSONString(t *testing.T) {
	mockT := new(testing.T)
	assert.True(t, JSONContains(mockT, `{"hello": "world", "foo": "bar"}`, `{"hello": "world", "foo": "bar"}`))
}

func TestJSONContains_EquivalentButNotEqual(t *testing.T) {
	mockT := new(testing.T)
	assert.True(t, JSONContains(mockT, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`))
}

func TestJSONContains_HashOfArraysAndHashes(t *testing.T) {
	mockT := new(testing.T)
	assert.True(t, JSONContains(mockT, "{\r\n\t\"numeric\": 1.5,\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]],\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\"\r\n}",
		"{\r\n\t\"numeric\": 1.5,\r\n\t\"hash\": {\"nested\": \"hash\", \"nested_slice\": [\"this\", \"is\", \"nested\"]},\r\n\t\"string\": \"foo\",\r\n\t\"array\": [{\"foo\": \"bar\"}, 1, \"string\", [\"nested\", \"array\", 5.5]]\r\n}"))
}

func TestJSONContains_Array(t *testing.T) {
	mockT := new(testing.T)
	assert.True(t, JSONContains(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`))
}

func TestJSONContains_ExpectedArrayLessThanActual(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `[{"foobar": 1}, "foo", {"nested": "hash", "hello": "world"}]`))
}

func TestJSONContains_ExpectedArrayBiggerThanActual(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, `[{"foobar": 1}, "foo", {"hello": "world", "nested": "hash"}]`, `["foo", {"nested": "hash", "hello": "world"}]`))
}

func TestJSONContains_ExpectedLessThanActual(t *testing.T) {
	mockT := new(testing.T)
	assert.True(t, JSONContains(mockT, `{"hello": {"world": 1}, "foo": [{"bar": "baz"}]}`,
		`{"hello": {"world": 1, "missing": "key"}, "foo": [{"bar": "baz", "waldo": "fred"}], "foobar": 2}`))
}

func TestJSONContains_ExpectedBiggerThanActual(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, `{"hello": {"world": 1}, "foo": [{"bar": "baz", "waldo": "fred"}], "foobar": 2}`,
		`{"hello": {"world": 1, "missing": "key"}, "foo": [{"bar": "baz"}]}`))
}

func TestJSONContains_HashAndArrayNotEquivalent(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `{"foo": "bar", {"nested": "hash", "hello": "world"}}`))
}

func TestJSONContains_ActualIsNotJSON(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, `{"foo": "bar"}`, "Not JSON"))
}

func TestJSONContains_ExpectedIsNotJSON(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, "Not JSON", `{"foo": "bar", "hello": "world"}`))
}

func TestJSONContains_ExpectedAndActualNotJSON(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, "Not JSON", "Not JSON"))
}

func TestJSONContains_ArraysOfDifferentOrder(t *testing.T) {
	mockT := new(testing.T)
	assert.False(t, JSONContains(mockT, `["foo", {"hello": "world", "nested": "hash"}]`, `[{ "hello": "world", "nested": "hash"}, "foo"]`))
}
