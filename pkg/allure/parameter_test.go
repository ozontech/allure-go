package allure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParameter(t *testing.T) {
	const paramName = "paramName"
	var paramValue = "paramValue"
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
	params := NewParameters("p1", "v1", "p2", "v2", "p3")
	require.NotNil(t, params)
	require.Len(t, params, 2)

	require.Equal(t, "p1", params[0].Name)
	require.Equal(t, "v1", params[0].GetValue())
	require.Equal(t, "p2", params[1].Name)
	require.Equal(t, "v2", params[1].GetValue())
}
