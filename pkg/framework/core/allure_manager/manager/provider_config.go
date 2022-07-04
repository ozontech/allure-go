package manager

type ConfigKey string

// describes base provider configs
const (
	SuiteName   ConfigKey = "suite"
	PackageName ConfigKey = "package"
	FullName    ConfigKey = "fullName"
	SuitePath   ConfigKey = "suitePath"
	Runner      ConfigKey = "runner"
	ParentSuite ConfigKey = "parentSuite"
)

// ProviderConfig describes configuration interface
type ProviderConfig interface {
	SuitePath() string
	SuiteName() string
	FullName() string
	PackageName() string
	ParentSuite() string
	Runner() string

	WithSuitePath(suitePath string) ProviderConfig
	WithSuiteName(suiteName string) ProviderConfig
	WithFullName(fullName string) ProviderConfig
	WithPackageName(packageName string) ProviderConfig
	WithParentSuite(parentSuite string) ProviderConfig
	WithRunner(runner string) ProviderConfig
}

type providerConfig struct {
	cfg map[ConfigKey]string
}

// NewProviderConfig ...
func NewProviderConfig() ProviderConfig {
	return &providerConfig{make(map[ConfigKey]string)}
}

// SuitePath ...
func (cfg *providerConfig) SuitePath() string {
	return cfg.cfg[SuitePath]
}

// SuiteName ...
func (cfg *providerConfig) SuiteName() string {
	return cfg.cfg[SuiteName]
}

// PackageName ...
func (cfg *providerConfig) PackageName() string {
	return cfg.cfg[PackageName]
}

// FullName ...
func (cfg *providerConfig) FullName() string {
	return cfg.cfg[FullName]
}

// Runner ...
func (cfg *providerConfig) Runner() string {
	return cfg.cfg[Runner]
}

// ParentSuite ...
func (cfg *providerConfig) ParentSuite() string {
	return cfg.cfg[ParentSuite]
}

// WithSuitePath ...
func (cfg *providerConfig) WithSuitePath(suitePath string) ProviderConfig {
	cfg.cfg[SuitePath] = suitePath
	return cfg
}

// WithSuiteName ...
func (cfg *providerConfig) WithSuiteName(suiteName string) ProviderConfig {
	cfg.cfg[SuiteName] = suiteName
	return cfg
}

// WithPackageName ...
func (cfg *providerConfig) WithPackageName(packageName string) ProviderConfig {
	cfg.cfg[PackageName] = packageName
	return cfg
}

// WithFullName ...
func (cfg *providerConfig) WithFullName(fullName string) ProviderConfig {
	cfg.cfg[FullName] = fullName
	return cfg
}

// WithRunner ...
func (cfg *providerConfig) WithRunner(runner string) ProviderConfig {
	cfg.cfg[Runner] = runner
	return cfg
}

// WithParentSuite ...
func (cfg *providerConfig) WithParentSuite(parentSuite string) ProviderConfig {
	cfg.cfg[ParentSuite] = parentSuite
	return cfg
}
