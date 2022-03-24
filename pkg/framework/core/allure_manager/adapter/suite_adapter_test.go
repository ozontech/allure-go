package adapter

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func TestNewSuiteMeta(t *testing.T) {
	adapter := NewSuiteMeta("packageName", "runner", "fullName", "suiteName")
	require.NotNil(t, adapter)
	require.Equal(t, "packageName", adapter.GetPackageName())
	require.Equal(t, "suiteName", adapter.GetSuiteName())
	require.Equal(t, "runner", adapter.GetRunner())
	require.Equal(t, "fullName", adapter.GetSuiteFullName())
	require.NotNil(t, adapter.GetContainer())
}

func TestSuiteAdapter_GetPackageName(t *testing.T) {
	adapter := &SuiteAdapter{packageName: "packageName"}
	require.Equal(t, "packageName", adapter.GetPackageName())
}

func TestSuiteAdapter_GetSuiteName(t *testing.T) {
	adapter := &SuiteAdapter{suiteName: "suiteName"}
	require.Equal(t, "suiteName", adapter.GetSuiteName())
}

func TestSuiteAdapter_GetRunner(t *testing.T) {
	adapter := &SuiteAdapter{runner: "runner"}
	require.Equal(t, "runner", adapter.GetRunner())
}

func TestSuiteAdapter_GetSuiteFullName(t *testing.T) {
	adapter := &SuiteAdapter{fullSuiteName: "fullName"}
	require.Equal(t, "fullName", adapter.GetSuiteFullName())
}

func TestSuiteAdapter_GetBeforeAll(t *testing.T) {
	adapter := &SuiteAdapter{beforeAll: func(t provider.T) {}}
	require.NotNil(t, adapter.GetBeforeAll())
}

func TestSuiteAdapter_GetAfterAll(t *testing.T) {
	adapter := &SuiteAdapter{afterAll: func(t provider.T) {}}
	require.NotNil(t, adapter.GetAfterAll())
}

func TestSuiteAdapter_SetBeforeAll(t *testing.T) {
	adapter := &SuiteAdapter{}
	adapter.SetBeforeAll(func(t provider.T) {})
	require.NotNil(t, adapter.GetBeforeAll())
}

func TestSuiteAdapter_SetAfterAll(t *testing.T) {
	adapter := &SuiteAdapter{}
	adapter.SetAfterAll(func(t provider.T) {})
	require.NotNil(t, adapter.GetAfterAll())
}

func TestSuiteAdapter_GetContainer(t *testing.T) {
	c := allure.NewContainer()
	adapter := &SuiteAdapter{container: c}
	require.NotNil(t, adapter.GetContainer())
	require.Equal(t, c, adapter.GetContainer())
}
