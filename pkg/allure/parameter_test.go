package allure

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParameter(t *testing.T) {
	const paramName = "paramName"
	paramValue := "paramValue"
	param := NewParameter(paramName, paramValue)
	require.NotNil(t, param)
	require.Equal(t, paramName, param.Name)
	require.Equal(t, string(paramValue), param.GetValue())
}

func TestNewParameters_even(t *testing.T) {
	params := NewParameters("p1", "v1", "p2", "v2", "p3", "v3")
	require.NotNil(t, params)
	require.Len(t, params, 3)

	require.Equal(t, "p1", params[0].Name)
	require.Equal(t, "v1", params[0].GetValue())
	require.Equal(t, "p2", params[1].Name)
	require.Equal(t, "v2", params[1].GetValue())
	require.Equal(t, "p3", params[2].Name)
	require.Equal(t, "v3", params[2].GetValue())
}

func TestNewParameters_odd(t *testing.T) {
	params := NewParameters("p1", "v1", "p2", "v2", "p3", "v3", 24)
	require.NotNil(t, params)
	require.Len(t, params, 2)

	require.Equal(t, "p1", params[0].Name)
	require.Equal(t, "v1", params[0].GetValue())

	require.Equal(t, "p2", params[1].Name)
	require.Equal(t, "v2", params[1].GetValue())

	require.Equal(t, "p3", params[2].Name)
	require.Equal(t, "24", params[2].GetValue())
}

func TestParameterUnmarshal(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		const data = `{"name": "epic", "value": "\"very epic indeed\""}`

		var param Parameter

		require.NoError(t, json.Unmarshal([]byte(data), &param))

		require.Equal(t, Parameter{
			Name:  "epic",
			Value: "very epic indeed",
		}, param)

		require.Equal(t, "very epic indeed", param.GetValue())
	})

	t.Run("int", func(t *testing.T) {
		const data = `{"name": "epic", "value": 83294782375982}`

		var param Parameter

		require.NoError(t, json.Unmarshal([]byte(data), &param))

		require.Equal(t, Parameter{
			Name:  "epic",
			Value: 83294782375982,
		}, param)

		require.Equal(t, "83294782375982", param.GetValue())
	})
}
