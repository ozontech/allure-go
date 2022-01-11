# allure-testify

Allure-Testify - проект, предоставляющий полноценный провайдер allure в go, без перегрузки интерфейса
использования. <br>
Проект начинался как форк от testify, но со временем обзавелся своим раннером и своими особенностями. <br>

## Other Languages README.md

- [English Readme](README_en.md)

## Head of contents

- [Other Languages README.md](#other-languages-readmemd)
- [Head of contents](#head-of-contents)
- [Features](#features)
- [Getting started](#getting-started)
- [Demo](#demo)
    - [Demo installation](#demo-installation)
    - [Run examples](#run-examples)
- [How to use](#how-to-use)
    - [Installation](#installation)
    - [Configure Behavior](#configure-behavior)
    - [Configure Test](#configure-test)
    - [Configure Suite](#configure-suite)
    - [Configure Your Report](#configure-your-report)
        - [Test Info](#test-info)
        - [Labels](#labels)
        - [Links](#links)
        - [Allure Steps](#allure-steps)
        - [Nested Only Functions](#nested-only-functions)
        - [Allure Attachments](#allure-attachments)
        - [Test Behavior](#test-behaviour)
        - [Forwarded actions to the suite.Suite](#forward-to-suitesuite)
- [Documentation](#documentation)
- [Examples](#examples)
    - [Test with nested steps](#test-with-nested-steps)
    - [Test with attachment](#test-with-attachment)
    - [Run few parallel suites](#run-few-parallel-suites)
    - [Setup hooks](#setup-hooks)
    - [XSkip](#xskip)

## Features

### `pkg/allure`

Пакет, содержащий модель данных для Allure. <br>
Полный список allure-объектов:

- `Attachment`
- `Container`
- `Label`
- `Link`
- `Parameter`
- `Result`
- `Step`

Предоставление отдельного пакета позволяет кастомизировать работу с allure.<br>
Подробно можно почитать [тут](#how-to-use-allure). <br>

### `pkg/provider.T`

Враппер контекста теста (`testing.T`). <br>
Основные преимущества и особенности:

- Имеет свой раннер тестов (`T.Run(testName string, test func(t *provider.T), tags ...string)`), что позволяет
  использовать преимущества библиотеки `testing`.
- Функциональность аналогов на других языках, без потери удобства и простоты использования.
- Полностью интегрирован с `allure`. Ваши go-тесты еще никогда не были такими информативными!

Подробно можно почитать [тут](#how-to-use-providert). <br>

### `pkg/framework/runner`

Пакет предоставляет функции для запуска тестовых структур (Suite) и отдельных тестов. <br>
Тесты, запущенные с помощью этих функций по окончанию исполнения будут создавать allure отчет.<br>
Подробно можно почитать [тут](#how-to-use-runner). <br>

### `pkg/framework/suite`

Пакет предоставляет структуру Suite, в которой можно описывать тесты, группируя их в тест-комплекты.<br>
Это может быть удобным, если у вас много разных тестов и вам сложно в них ориентироваться, без дополнительных "уровней
вложения" вызовов тестов. <br>
Подробно можно почитать [тут](#how-to-use-suite). <br>

## Getting Started

1. Установить пакет <br>

```bash
go get github.com/koodeex/allure-testify
```

2. Если Вы уже используете testify, то нужно заменить импорты

```go
package tests

import (
	"github.com/stretchr/testify/suite"
)
``` 

на

```go
package tests

import (
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)
``` 

3. Заменить функции <br>

* `SetupSuite` -> `BeforeAll` <br>
* `SetupTest` -> `BeforeEach` <br>
* `TearDownTest` -> `AfterEach` <br>
* `TearDownSuite` -> `AfterAll` <br>

4. Запустить go test!

## Demo

### Demo Installation

```bash
  git clone https://github.com/koodeex/allure-testify.git
```

### Run Examples

```bash
make demo
```

## How to use

### Installation

```bash
go get github.com/koodeex/allure-testify
```

### Configure Behavior

Путь до allure отчетов собирается из двух глобальных переменных `$ALLURE_OUTPUT_FOLDER/$ALLURE_OUTPUT_PATH`

- `ALLURE_OUTPUT_FOLDER` - это имя папки, в которую будут складываться allure-отчеты (по умолчанию - `allure-results`).
- `ALLURE_OUTPUT_PATH` - это путь, в котором будет создана `ALLURE_OUTPUT_FOLDER` (по умолчанию это корневая папка
  запуска тестов).

Так же, можно указать несколько глобальных конфигураций, для интеграции с вашей TMS или Task Tracker:

- `ALLURE_ISSUE_PATTERN` - Указывает урл-паттерн для ваших Issues. Не имеет значения по умолчанию. **Обязательно**
  должен содержать `%s`.

Если `ALLURE_ISSUE_PATTERN` не задан, ссылка будет читаться целиком.

Пример:

```go
package provider_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/provider"
)

func TestSampleDemo(t *testing.T) {
	runner.RunTest(t, "Just Link", func(t *provider.T) {
		t.SetIssue("https://pkg.go.dev/github.com/stretchr/testify")
	})

	runner.RunTest(t, "With Pattern", func(t *provider.T) {
		_ = os.Setenv("ALLURE_ISSUE_PATTERN", "https://pkg.go.dev/github.com/stretchr/%s")
		t.SetIssue("testify")
	})
}
```

- `ALLURE_TESTCASE_PATTERN` - Указывает урл-паттерн для ваших TestCases. Не имеет значения по умолчанию. **Обязательно**
  должен содержать `%s`.

Если `ALLURE_TESTCASE_PATTERN` не задан, ссылка будет читаться целиком.

Пример:

```go
package provider_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/provider"
)

func TestSampleDemo(t *testing.T) {
	runner.RunTest(t, "Just Link", func(t *provider.T) {
		t.SetTestCase("https://pkg.go.dev/github.com/stretchr/testify")
	})

	runner.RunTest(t, "With Pattern", func(t *provider.T) {
		_ = os.Setenv("ALLURE_TESTCASE_PATTERN", "https://pkg.go.dev/github.com/stretchr/%s")
		t.SetTestCase("testify")
	})
}
```

- `ALLURE_LAUNCH_TAGS` - Прокидывает список тэгов, которые будут применены к каждому тесту по умолчанию. Не имеет
  значения по умолчанию.

**Совет:** `ALLURE_LAUNCH_TAGS` - очень удобен в использовании с CI/CD. Например, в нем можно определять группы тестов
по вашим ci-jobs или же прокидывать имя ветки.

### Configure Test

1) Используя пакет runner:

```go
package provider_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/provider"
)

func TestSampleDemo(t *testing.T) {
	runner.RunTest(t, "My test", func(t *provider.T) {
		// Test Body
	})
}
```

2) Используя декларирование контекста `TestRunner`:

```go
package provider_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
)

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewTestRunner(realT)
	r.Run("My test", func(t *provider.T) {
		// Test Body
	})
}
```

Второй вариант позволит использовать BeforeEach/AfterEach:

```go
package provider_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
)

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewTestRunner(realT)
	r.WithBeforeEach(func(t *provider.T) {
		// Before Each body 
	})
	r.WithAfterEach(func(t *provider.T) {
		// After Each body
	})
	r.Run("My test", func(t *provider.T) {
		// Test Body
	})
}
```

Так же сохранена особенность библиотеки `testing`, позволяющая запускать тесты из других тестов:

```go
package provider_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/provider"
)

func TestOtherSampleDemo(realT *testing.T) {
	r := runner.NewTestRunner(realT)
	r.Run("My test", func(t *provider.T) {
		// Test Body
		t.WithBeforeTest(func(t *provider.T) {
			// inner Before Each body
		})
		t.WithAfterTest(func(t *provider.T) {
			// inner After Each body
		})
		t.Run("My test", func(t *provider.T) {
			// inner test body
		})
	})
}
```

### Configure Suite

Чтобы группировать тесты в тест-комплекты, необходимо:

1) объявить структуру, методами которой будут ваши тесты

```go
package suite_demo

type DemoSuite struct {
}
```

2) расширить объявленную структуру структурой `suite.Suite`

```go
package suite_demo

import "github.com/koodeex/allure-testify/pkg/framework/suite"

type DemoSuite struct {
	suite.Suite
}
```

3) описать тесты

```go
package suite_demo

import "github.com/koodeex/allure-testify/pkg/framework/suite"

type DemoSuite struct {
	suite.Suite
}

func (s *DemoSuite) TestSkip() {
	s.Epic("Demo")
	s.Feature("Suites")
	s.Title("My first test")
	s.Description(`
		This test will be attached to the suite DemoSuite`)
}
```

4) Запустить тесты.

Для этого нужно описать функцию, которая запустить Ваш тест и вызвать `runner.RunSuite`:

```go
package suite_demo

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/runner"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type DemoSuite struct {
	suite.Suite
}

func (s *DemoSuite) TestSkip() {
	s.Epic("Demo")
	s.Feature("Suites")
	s.Title("My first test")
	s.Description(`
		This test will be attached to the suite DemoSuite`)
}

func TestSkipDemo(t *testing.T) {
	t.Parallel()
	runner.RunSuite(t, new(SkipDemoSuite))
}
```

И запустить тесты с помощью `go test`

```bash
go test ${TEST_PATH}
```

Тогда в корневой папке тестов по окончанию прогона будет проинициализирована папка `allure-results`, содержащая в себе
allure-отчеты.

### Configure Your Report

Allure-Testify предоставляет широкие возможности взаимодействия с allure.<br>
Большинство действий осуществляется с помощью структуры `provider.T`, являющейся оберткой над `testing.T`.<br>
Так же структура `suite.Suite` позволяет использовать интерфейс Suite для взаимодействия с allure-report.

#### Test Info

Полный список поддерживаемых методов для проставления информации о методе:

- `*T.Title`
- `*T.Description`

**Note:** По умолчанию имя теста ставится в соответствии с именем функции теста.

#### Labels

Полный список поддерживаемых лейблов:

- `*T.Epic`
- `*T.Feature`
- `*T.Story`
- `*T.ID`
- `*T.Severity`
- `*T.ParentSuite`
- `*T.Suite`
- `*T.SubSuite`
- `*T.Package`
- `*T.Thread`
- `*T.Host`
- `*T.Tag`
- `*T.Framework`
- `*T.Language`
- `*T.Owner`
- `*T.Lead`

Более подробно про методы можно почитать [здесь](/pkg/allure/README.md#allurelabel)

##### Default label values

| Label | Default Value |
|---|---|
|`ParentSuite`| - Для suite.Suite ставится имя функции, в которой suite был запущен.<br><br> - Для независимых тестов - по умолчанию не ставится, однако если тест был вызван внутри другого теста, в этом лейбле будет указан тест, запустивший родительский для текущего тест.|
|`Suite`|- Для suite.Suite ставится имя suite, которому текущий тест принадлежит. <br><br> - Для независимого теста - имя теста - родителя (из которого был запущен текущий тест).|
|`Package`|Пакет, в котором были запущены тесты|
|`Thread`|Ставится `Result.FullName` [[1]](#label_note_answer1).|
|`Host`|`os.Host()`|
|`Framework`|`Allure-Testify@v0.3.x`|
|`Language`|`runtime.Version()`|

___________________________________  
***NOTES:***
<a name="label_note_answer1">[1]</a> - Это Knowing Issue - в golang пока не представляется целесообразным (или возможным
адекватными способами) пытаться достать имя текущей goroutine, как и невозможно задать ей имя.
___________________________________

#### Links

Полный список поддерживаемых действий:

- `*T.SetIssue`
- `*T.SetTestCase`
- `*T.Link`

Более подробно про методы можно почитать [здесь](/pkg/allure/README.md#allurelink).

Про переменные, с которыми можно взаимодействовать для упрощения работы было указано [выше](#configure-behavior).

#### Allure Steps

Полный список поддерживаемых действий:

- `*T.Step` - добавляет к отчету переданный Step.
- `*T.NewStep` - создает новый пустой Step с переданным именем и добавляет его к отчету.
- `*T.InnerStep` - добавляет к отчету переданный Step, проставив переданный ParentStep как родительский.
- `*T.NewInnerStep` - создает новый пустой Step с переданным именем и добавляет его к отчету, проставляет переданный
  ParentStep как родительский и добавляет его к отчету.
- `*T.WithStep` - оборачивает переданную в f функцию переданным Step и добавляет Step к отчету.
- `*T.WithNewStep` - создает новый Step, оборачивает переданную в f функцию созданным Step и добавляет его к отчету.

#### Nested-Only Functions

- `*T.AddAttachmentToNested` - добавляет к вложенному шагу Attachment
- `*T.AddParameterToNested` - добавляет к вложенному шагу Parameter
- `*T.AddParametersToNested` - добавляет к вложенному шагу все элементы массива Parameter
- `*T.AddNewParameterToNested` - инициализирует новый Parameter и добавляет его к вложенному шагу
- `*T.AddNewParametersToNested` - инициализирует новый массив Parameter и добавляет все его элементы к вложенному шагу

**Note:** Функции с суффиксом `ToNested` могут быть вызваны **ТОЛЬКО** внутри функции `WithStep`/`WithNewStep`. В
противном случае ничего не произойдет.

#### Allure Attachments

Полный список поддерживаемых действий:

- `*T.Attachment` - добавляет к текущему тесту Attachment

#### Test Behaviour

Полный список поддерживаемых действий:

- `*T.Skip` - пропускает текущий тест. В статус отчета будет указан переданный текст.
- `*T.Errorf` - помечает выбранный тест, как Failed. В статус отчета будет прикреплен переданный текст,
- `*T.XSkip` - пропускает выбранный тест, если в процессе его исполнения вызывается `*T.Error`/`*T.Errorf` (например,
  падает assert)

#### Forward to suite.Suite

Полный список поддерживаемых действий:

- [Test Info](#test-info)
    - `*Suite.Title`
    - `*Suite.Description`
- [Allure Labels](#labels)
    - `*Suite.Epic`
    - `*Suite.Feature`
    - `*Suite.Story`
    - `*Suite.ID`
    - `*Suite.Severity`
    - `*Suite.ParentSuite`
    - `*Suite.Suite`
    - `*Suite.SubSuite`
    - `*Suite.Package`
    - `*Suite.Thread`
    - `*Suite.Host`
    - `*Suite.Tag`
    - `*Suite.Framework`
    - `*Suite.Language`
    - `*Suite.Owner`
    - `*Suite.Lead`
- [Allure Links](#links)
    - `*Suite.SetIssue`
    - `*Suite.SetTestCase`
    - `*Suite.Link`
- [Allure Steps](#allure-steps)
    - `*Suite.Step`
    - `*Suite.NewStep`
    - `*Suite.InnerStep`
    - `*Suite.InnerNewStep`
    - `*Suite.WithStep`
    - `*Suite.WithNewStep`
- [Nested Only Functions](#nested-only-functions)
    - `*Suite.AddNestedAttachment`
    - `*Suite.AddParameterToNested`
    - `*Suite.AddParametersToNested`
    - `*Suite.AddNewParameterToNested`
    - `*Suite.AddNewParametersToNested`
- [Allure Attachments](#allure-attachments)
    - `*Suite.Attachment`

## Documentation

Подробная документация по каждому публичному пакету может быть найдена в каталоге этого пакета.

- [allure](/pkg/allure/README.md)
- [provider](/pkg/provider/README.md)
- [runner](/pkg/framework/runner/README.md)
- [suite](/pkg/framework/suite/README.md)

## Examples

### [Test with nested steps](examples/suite_demo/step_tree_test.go):

Код теста:

```go
package examples

import (
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type StepTreeDemoSuite struct {
	suite.Suite
}

func (s *StepTreeDemoSuite) TestInnerSteps() {
	s.Epic("Demo")
	s.Feature("Inner Steps")
	s.Title("Simple Nesting")
	s.Description(`
		Step A is parent step for Step B and Step C
		Call order will be saved in allure report
		A -> (B, C)`)

	s.Tags("Steps", "Nesting")

	s.WithNewStep("Step A", func() {
		s.NewStep("Step B")
		s.NewStep("Step C")
	})
}
```

Вывод в Allure:

![](.resources/example_step_tree.png)

### [Test with Attachment](examples/suite_demo/attachments_test.go)

Код теста:

```go
package examples

import (
	"encoding/json"

	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type JSONStruct struct {
	Message string `json:"message"`
}

type AttachmentTestDemoSuite struct {
	suite.Suite
}

func (s *AttachmentTestDemoSuite) TestAttachment() {
	s.Epic("Demo")
	s.Feature("Attachments")
	s.Title("Test Attachments")
	s.Description(`
		Test's test body and all steps inside can contain attachments`)

	s.Tags("Attachments", "BeforeAfter", "Steps")

	attachmentText := `THIS IS A TEXT ATTACHMENT`
	s.Attachment(allure.NewAttachment("Text Attachment if TestAttachment", allure.Text, []byte(attachmentText)))

	step := allure.NewSimpleStep("Step A")
	var ExampleJson = JSONStruct{"this is JSON message"}
	attachmentJSON, _ := json.Marshal(ExampleJson)
	step.Attachment(allure.NewAttachment("Json Attachment for Step A", allure.JSON, attachmentJSON))
	s.Step(step)
}
```

Вывод в Allure:

![](.resources/example_attachments.png)

### [Run few parallel suites](examples/suite_demo/running_test.go)

Код теста:

```go
package examples

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type TestRunningDemoSuite struct {
	suite.Suite
}

func (s *TestRunningDemoSuite) TestBeforesAfters() {
	t := s.T()
	t.Parallel()
	// use RunInner to run suite of tests
	s.RunSuite(t, new(BeforeAfterDemoSuite))
}

func (s *TestRunningDemoSuite) TestFails() {
	t := s.T()
	t.Parallel()
	s.RunSuite(t, new(FailsDemoSuite))
}

func (s *TestRunningDemoSuite) TestLabels() {
	t := s.T()
	t.Parallel()
	s.RunSuite(t, new(LabelsDemoSuite))
}

func TestRunDemo(t *testing.T) {
	// use RunSuites to run suite of suites
	suite.RunSuite(t, new(TestRunningDemoSuite))
}
```

Вывод в Allure:

![](.resources/example_multiple_suites_run.png)

### [Setup hooks](examples/suite_demo/befores_afters_test.go)

Код теста:

```go
package examples

import (
	"testing"

	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type BeforeAfterDemoSuite struct {
	suite.Suite
}

func (s *BeforeAfterDemoSuite) BeforeEach() {
	s.NewStep("Before Test Step")
}

func (s *BeforeAfterDemoSuite) AfterEach() {
	s.NewStep("After Test Step")
}

func (s *BeforeAfterDemoSuite) BeforeAll() {
	s.NewStep("Before suite Step")
}

func (s *BeforeAfterDemoSuite) AfterAll() {
	s.NewStep("After suite Step")
}

func (s *BeforeAfterDemoSuite) TestBeforeAfterTest() {
	s.Epic("Demo")
	s.Feature("BeforeAfter")
	s.Title("Test wrapped with SetUp & TearDown")
	s.Description(`
		This test wrapped with SetUp and TearDown containers.`)

	s.Tags("BeforeAfter")
}

func TestBeforesAfters(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(BeforeAfterDemoSuite))
}
```

Вывод в Allure:

![](.resources/example_befores_afters.png)

### [XSkip](examples/suite_demo/fails_test.go)

Код теста:

```go
package examples

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/koodeex/allure-testify/pkg/framework/suite"
)

type DemoSuite struct {
	suite.Suite
}

func (s *DemoSuite) TestXSkipFail() {
	s.Title("This test skipped by assert with message")
	s.Description(`
		This Test will be skipped with assert Error.
		Error text: Assertion Failed`)
	s.Tags("fail", "xskip", "assertions")

	t := s.T()
	t.XSkip()
	require.Equal(t, 1, 2, "Assertion Failed")
}

func TestDemoSuite(t *testing.T) {
	t.Parallel()
	suite.RunSuite(t, new(DemoSuite))
}
```

Вывод в Allure:

![](.resources/example_xskip.png)
