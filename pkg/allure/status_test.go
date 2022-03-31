package allure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	require.Equal(t, "passed", string(Passed))
	require.Equal(t, "failed", string(Failed))
	require.Equal(t, "skipped", string(Skipped))
	require.Equal(t, "broken", string(Broken))
	require.Equal(t, "unknown", string(Unknown))
}
