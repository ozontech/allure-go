package adapter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

func TestNewTestMeta(t *testing.T) {
	host, _ := os.Hostname()

	adapter := NewTestMeta("fullName", "suiteName", "testName", "packageName", "tag1", "tag2")
	require.Equal(t, "testName", adapter.GetResult().Name)
	require.Equal(t, "fullName/testName", adapter.GetResult().FullName)

	require.Len(t, adapter.GetResult().GetLabels(allure.Host), 1)
	require.Equal(t, adapter.GetResult().GetLabels(allure.Host)[0].Value, host)

	require.Len(t, adapter.GetResult().GetLabels(allure.Framework), 1)
	require.Equal(t, adapter.GetResult().GetLabels(allure.Framework)[0].Value, allure.DefaultVersion)

	require.Len(t, adapter.GetResult().GetLabels(allure.Thread), 1)
	require.Equal(t, "fullName/testName", adapter.GetResult().GetLabels(allure.Thread)[0].Value)

	require.Len(t, adapter.GetResult().GetLabels(allure.Suite), 1)
	require.Equal(t, "suiteName", adapter.GetResult().GetLabels(allure.Suite)[0].Value)

	require.Len(t, adapter.GetResult().GetLabels(allure.Package), 1)
	require.Equal(t, "packageName", adapter.GetResult().GetLabels(allure.Package)[0].Value)

	require.Len(t, adapter.GetResult().GetLabels(allure.Tag), 2)
	require.Equal(t, "tag1", adapter.GetResult().GetLabels(allure.Tag)[0].Value)
	require.Equal(t, "tag2", adapter.GetResult().GetLabels(allure.Tag)[1].Value)

}

func TestTestAdapter_GetResult(t *testing.T) {
	test := &allure.Result{}
	adapter := TestAdapter{result: test}
	require.Equal(t, test, adapter.GetResult())
}

func TestTestAdapter_SetResult(t *testing.T) {
	test := &allure.Result{}
	adapter := TestAdapter{}
	adapter.SetResult(test)
	require.Equal(t, test, adapter.GetResult())
}

func TestTestAdapter_SetBeforeEach(t *testing.T) {
	adapter := TestAdapter{}
	adapter.SetBeforeEach(func(t provider.T) {})
	require.NotNil(t, adapter.GetBeforeEach())
}

func TestTestAdapter_SetAfterEach(t *testing.T) {
	adapter := TestAdapter{}
	adapter.SetAfterEach(func(t provider.T) {})
	require.NotNil(t, adapter.GetAfterEach())
}

func TestTestAdapter_GetBeforeEach(t *testing.T) {
	adapter := TestAdapter{beforeEach: func(t provider.T) {}}
	require.NotNil(t, adapter.GetBeforeEach())
}

func TestTestAdapter_GetAfterEach(t *testing.T) {
	adapter := TestAdapter{afterEach: func(t provider.T) {}}
	require.NotNil(t, adapter.GetAfterEach())
}

func TestTestAdapter_GetContainer(t *testing.T) {
	container := allure.NewContainer()
	adapter := TestAdapter{container: container}
	require.Equal(t, container, adapter.container)
}
