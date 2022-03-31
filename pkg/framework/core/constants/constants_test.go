package constants

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstants(t *testing.T) {
	require.Equal(t, "test", TestContextName)
	require.Equal(t, "beforeEach", BeforeEachContextName)
	require.Equal(t, "afterEach", AfterEachContextName)
	require.Equal(t, "beforeAll", BeforeAllContextName)
	require.Equal(t, "afterAll", AfterAllContextName)
}
