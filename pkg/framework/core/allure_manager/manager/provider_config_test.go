package manager

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProviderConfigConstants(t *testing.T) {
	require.Equal(t, "suite", string(SuiteName))
	require.Equal(t, "package", string(PackageName))
	require.Equal(t, "fullName", string(FullName))
	require.Equal(t, "suitePath", string(SuitePath))
	require.Equal(t, "runner", string(Runner))
}

func TestNewProviderConfig(t *testing.T) {
	cfg := NewProviderConfig()
	require.NotNil(t, cfg)
}

func TestProviderConfig_values(t *testing.T) {
	cfg := providerConfig{map[ConfigKey]string{}}

	cfg.WithRunner("runner")
	require.NotEmpty(t, cfg.cfg[Runner])
	require.Equal(t, "runner", cfg.cfg[Runner])
	require.NotEmpty(t, cfg.Runner())
	require.Equal(t, "runner", cfg.Runner())

	cfg.WithSuiteName("suiteName")
	require.NotEmpty(t, cfg.cfg[SuiteName])
	require.Equal(t, "suiteName", cfg.cfg[SuiteName])
	require.NotEmpty(t, cfg.SuiteName())
	require.Equal(t, "suiteName", cfg.SuiteName())

	cfg.WithFullName("fullName")
	require.NotEmpty(t, cfg.cfg[FullName])
	require.Equal(t, "fullName", cfg.cfg[FullName])
	require.NotEmpty(t, cfg.FullName())
	require.Equal(t, "fullName", cfg.FullName())

	cfg.WithPackageName("packageName")
	require.NotEmpty(t, cfg.cfg[PackageName])
	require.Equal(t, "packageName", cfg.cfg[PackageName])
	require.NotEmpty(t, cfg.PackageName())
	require.Equal(t, "packageName", cfg.PackageName())

	cfg.WithSuitePath("suitePath")
	require.NotEmpty(t, cfg.cfg[SuitePath])
	require.Equal(t, "suitePath", cfg.cfg[SuitePath])
	require.NotEmpty(t, cfg.SuitePath())
	require.Equal(t, "suitePath", cfg.SuitePath())
}
