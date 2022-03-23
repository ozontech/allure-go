# suite

Пакет `suite` представляет собой имплементацию аналогов тест классов из JUnit, TestNG, pytest или их аналогов.

Вдохновлено аналогичным пакетом из библиотеки `testify`.

Позволяет группировать тесты в единый тест комплект, расширяя свою структуру через `suite.Suite`. Удобно для группировки
и запуска тестов по тегам, постфиксам и тд.

## Head of contents

- static
    - [func RunSuite(t *testing.T, suite framework.TestSuite)](#func-runsuitet-testingt-suite-frameworktestsuite)
- struct
    - [func (suite *Suite) SetT(t *provider.T)](#func-suite-suite-settt-providert)
    - [func (suite *Suite) GetName() string](#func-suite-suite-getname-string)
    - [func (suite *Suite) GetPackage() string](#func-suite-suite-getpackage-string)
    - [func (suite *Suite) SkipOnPrint()](#func-suite-suite-skiponprint)
    - [func (suite *Suite) T() *provider.T](#func-suite-suite-t-providert)
    - [func (suite *Suite) RunTest(testName string, test func(t *provider.T), tags ...string) bool](#func-suite-suite-runtesttestname-string-test-funct-providert-tags-string-bool)
    - [func (suite *Suite) Run(testName string, test func(), tags ...string) bool](#func-suite-suite-runtestname-string-test-func-tags-string-bool)
    - [func (suite *Suite) RunSuite(t *provider.T, newSuite AllureSuite)](#func-suite-suite-runsuitet-providert-newsuite-alluresuite)

## [runner](runner.go)

### func RunSuite(t *testing.T, suite framework.TestSuite)

## [suite](suite.go)

Описание структуры.

```go
package suite

import (
	"github.com/koodeex/allure-testify/pkg/provider"
)

// AllureSuite is an interface that describes Suite behaviour
type AllureSuite interface {
	SetT(t *provider.T)
	T() *provider.T
	GetName() string
	setName(string)
	GetPackage() string
	setPackage(string)
	GetParent() string
	setParent(string)
}

// Suite is test-class like object, that allows group tests in test suites.
type Suite struct {
	name        string
	parent      string
	packageName string

	t *provider.T
}
```

### func (suite *Suite) SetT(t *provider.T)

### func (suite *Suite) GetName() string

### func (suite *Suite) GetPackage() string

### func (suite *Suite) GetParent() string

### func (suite *Suite) SkipOnPrint()

### func (suite *Suite) T() *provider.T

### func (suite *Suite) RunTest(testName string, test func(t *provider.T), tags ...string) bool

### func (suite *Suite) Run(testName string, test func(), tags ...string) bool

### func (suite *Suite) RunSuite(t *provider.T, newSuite AllureSuite)

## [forward_allure](forward_allure.go)

`suite.Suite` поддерживает все allure-методы провайдера для упрощения взаимодействия с интерфейсом.

Подробное описание методов можно найти [здесь](../../provider/README.md)