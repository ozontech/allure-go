package internal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConstants(t *testing.T) {
	require.Equal(t, "test", TestContextName)
	require.Equal(t, "beforeEach", BeforeEachContextName)
	require.Equal(t, "afterEach", AfterEachContextName)
	require.Equal(t, "beforeAll", BeforeAllContextName)
	require.Equal(t, "afterAll", AfterAllContextName)
}
