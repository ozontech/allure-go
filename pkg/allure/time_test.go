package allure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNow(t *testing.T) {
	require.NotZero(t, GetNow())
}
