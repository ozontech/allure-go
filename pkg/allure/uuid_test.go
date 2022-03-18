package allure

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetUUID(t *testing.T) {
	require.NotZero(t, getUUID())
}
