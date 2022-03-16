# allure-testify

## Оглавление

- [Demo](#demo)
- [Getting started](#getting-started)
- [Examples](#examples)
- [Global environments keys](#global-environment-keys)
- [How to use runner](#how-to-use-runner)
- [How to use provider](#how-to-use-providert)
  - [Allure info](#allure-info)
      - [Test info](#test-info)
      - [Label](#label)
      - [Link](#link)
  - [Allure Actions](#allure-actions)
      - [Step](#step)
      - [Attachment](#attachment)
      - [Parameter](#parameter)
      - [Before/After (experimental)](#beforeafter)
  - [Test Behavior](#test-behavior)
      - [Example t.Run()](#example-trun)
- [How to use suite](#how-to-use-suite)
  - [Behavior](#behavior)
  - [Suite Before/Afters](#suite-beforeafters)
  - [Allure forward](#allure-forward)
- [How to use allure](#how-to-use-allure)
    - [Allure Objects](#allure-objects)
        - [Steps](#steps)
        - [Attachments](#attachments)

## Demo

Чтобы выкачать проект:

```bash
git clone https://github.com/ozontech/allure-go
```

Чтобы установить необходимые для примеров инструменты (mac)

```bash
make install
```

Чтобы сформировать аллюр отчет из примеров:

```bash
make demo
``` 

## Getting started

1. Установить пакет <br>

```bash
go get github.com/ozontech/allure-go
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
	"github.com/ozontech/allure-go/pkg/suite"
)
``` 

3. Заменить функции <br>

* `SetupSuite` -> `BeforeAll` <br>
* `SetupTest` -> `BeforeEach` <br>
* `TearDownTest` -> `AfterEach` <br>
* `TearDownSuite` -> `AfterAll` <br>

4. Запустить go test!

## Examples

### [Тест с вложенными шагами](examples_old/suite_demo/step_tree_test.go):

Код теста:

```go
package examples

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/suite"
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

![](_resources/example_step_tree.png)

### [Тест с Attachment](examples_old/suite_demo/attachments_test.go)

Код теста:

```go
package examples

import (
	"encoding/json"
	"testing"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/suite"
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

![](_resources/example_attachments.png)

### [Запуск нескольких suites](examples_old/suite_demo/running_test.go)

Код теста:

```go
package examples

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/suite"
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

![](_resources/example_multiple_suites_run.png)

## [Setup hooks](examples_old/suite_demo/befores_afters_test.go)

Код теста:

```go
package examples

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/suite"
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

![](_resources/example_befores_afters.png)

## Global Environment keys

| Key | Meaning | Default |
|---|---|---|
|`ALLURE_OUTPUT_PATH`|Указывает путь до папки для печати результатов.|`.` (Папка с тестами)|
|`ALLURE_OUTPUT_FOLDER`|Указывает имя папки для печати результатов.|`/allure-results`|
|`ALLURE_ISSUE_PATTERN`|Указывает URL паттерн для Issue. **Обязательно должен содержать один `%s`**| |
|`ALLURE_TESTCASE_PATTERN`|Указывает URL паттерн для TestCase. **Обязательно должен содержать один `%s`**.| |
|`ALLURE_LAUNCH_TAGS`|Указывает дефолтные тэги, которыми будут помечаться все тесты в прогоне. Тэги должны быть указаны через запятую.| |

## HOW TO USE runner

| Method Signature| Meaning |
|---|---|
|`RunSuite(*testing.T, framework.TestSuite)`| Запускает переданный Suite в контексте переданного `*testing.T` |
|`RunTest(*testing.T, string, func(t *provider.T), ...string) bool`| Запускает переданный тест в контексте переданного `*testing.T`. Тесты запущенные через `RunTest` могут быть параллельны относительно друг друга. См пример ниже.  |

## HOW TO USE provider.T

### Allure info

#### Test info

| Method Signature| Meaning |
|---|---|
|`Suite.Title(string)`|устанавливает имя теста|
|`Suite.Description(string)`|устанавливает описание теста|

#### Label

| Method Signature| Meaning |
|---|---|
|`T.Epic(string)`| устанавливает Epic для теста|
|`T.Feature(string)`|устанавливает Feature теста|
|`T.Story(string)`|устанавливает Story теста|
|`T.FrameWork(string)`|устанавливает FrameWork теста (по умолчанию проставляет allure-testify)|
|`T.Host(string)`|устанавливает Host теста|
|`T.Thread(string)`|устанавливает Thread для теста.(по умолчанию устанавливает имя сьюта + UUID.<a name="label_note_question1">[[1]](#label_note_answer1)</a> | 
|`T.ID(string)`|устанавливает ID теста|
|`T.AddSuiteLabel(string)`|добавляет Suite, к которому относится тест|
|`T.AddSubSuite(string)`|добавляет SubSuite, к которому относится тест|
|`T.AddParentSuite(string)`|добавляет ParentSuite, к которому относится тест|
|`T.Severity(string)`|устанавливает Severity теста (по умолчанию - `normal`)|
|`T.Tag(string)`|добавляет Tag теста|
|`T.Tags(...string)`|добавляет несколько Tag теста|
|`T.Package(string)`|устанавливает Package к которому относится тест (по умолчанию выбирается Caller функция)|
|`T.Owner(string)`|устанавливает Owner'а, которому принадлежит тест|
|`T.Label(string, string)`|добавляет произвольный Лейбл к тесту. Имя лейбла это первый аргумент, значение - второй.|

___________________________________  
***NOTES:***
<a name="label_note_answer1">[1]</a> - Это Knowing Issue - в golang пока не представляется целесообразным пытаться
достать имя текущей goroutine, как и нельзя задать ей имя
___________________________________

#### Link

| Method Signature| Meaning |
|---|---|
|`T.SetIssue(string)`|добавляет ссылку на Issue (чтобы задать паттерн, нужно установить глобальную переменную `ALLURE_ISSUE_PATTERN`)|
|`T.SetTestCase(string)`|добавляет ссылку на TestCase (чтобы задать паттерн, нужно установить глобальную переменную `ALLURE_TEST_CASE`)|
|`T.Link(allure.Link)`|добавляет произвольную ссылку к тесту|

### Allure Actions

#### Step

| Method Signature| Meaning |
|---|---|
|`T.Step(*allure.Step)`|добавляет в текущий тест новый шаг. Принимает в себя тип `allure.Step`|
|`T.NewStep(string)`|добавляет в текущий тест пустой шаг с именем, переданным строкой|
|`T.InnerStep(*allure.Step, *allure.Step)`|добавляет в текущий тест шаг, переданный вторым аргументом. Первым аргументом передается родительский шаг.|
|`T.WithStep(*allure.Step, func())`|объявляет, что все последующие шаги, объявленные в анонимной функции (второй аргумент) будут считаться вложенными в шаг, переданный первым аргументом.|
|`T.WithNewStep(string, func())`|объявляет, что все последующие шаги будут считаться вложенными в новый шаг с именем, переданным первым аргументом.|

#### Attachment

| Method Signature| Meaning |
|---|---|
|`T.Attachment(*allure.Attachment)`|добавляет к текущему тесту `allure.Attachment`. Если вызван в функциях-hook'ах, создаст новый шаг, к которому будет прикреплен аттачмент.|
|`T.AddNestedAttachment(*allure.Attachment)`|добавляет к текущему вложенному шагу `allure.Attachment`. Можно вызвать **ТОЛЬКО** внутри функции `WithStep`/`WithNewStep`.|

#### Parameter

| Method Signature| Meaning |
|---|---|
|`T.AddParameterToNested(*allure.Parameter)`|добавляет к текущему родительскому шагу `allure.Parameter`. Можно вызвать **ТОЛЬКО** внутри функции `WithStep`/`WithNewStep`.|
|`T.AddParametersToNested([]allure.Parameter)`|добавляет к текущему родительскому шагу массив `allure.Parameter`. Можно вызвать **ТОЛЬКО** внутри функции `WithStep`/`WithNewStep`.|
|`T.AddNewParameterToNested(string, string)`|добавляет к текущему родительскому шагу `allure.Parameter` с именем, переданным первым аргументом и значением, переданным вторым. Можно вызвать **ТОЛЬКО** внутри функции `WithStep`/`WithNewStep`.|
|`T.AddNewParametersToNested(...string)`|добавляет к текущему родительскому шагу массив `allure.Parameter` с именами и значениями через запятую (`k1, v1, k2, v2... kn, vn`)|

### Test behavior

| Method Signature| Meaning |
|---|---|
|`T.Skip(...interface{})` | Позволяет пропустить тест. Причина пропуска будет залогирована в статус аллюр отчета. |
|`T.Run(string, func(*T), ...string) bool` | Позволяет запускать тесты. Тесты запущенные через T.Run могут быть параллельны относительно друг друга. См пример ниже. |

#### Example `T.Run`

[Код теста:](examples_old/provider_demo/sample_test.go)

```go
package Test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/provider"
)

func TestOtherSampleDemo(realT *testing.T) {
	t := provider.NewT(realT, realT.Name(), "")
	t.Run("My test", func(t *provider.T) {
		t.WithBeforeTest(func() {
			t.NewStep("This is before test step")
		})
		defer t.WithAfterTest(func() {
			t.NewStep("This is after test step")
		})

		t.Epic("Only Provider Demo")
		t.Feature("T.Run()")

		t.Title("Some Other Sample test")
		t.Description("allure-testify allows you to use allure without suites")

		t.WithNewStep("Some nested step", func() {
			t.WithNewStep("Some inner step 1", func() {
				t.WithNewStep("Some inner step 1.1", func() {

				})
			})
			t.WithNewStep("Some inner step 2", func() {
				t.WithNewStep("Some inner step 2.1", func() {

				})
			})
		})
	}, "Sample", "Provider-only", "with provider initialization")
}
```

Вывод в Allure:

![](_resources/example_sample_provider.png)

### Before/After

**ВАЖНО**: данный раздел является экспериментальным. Крайне не рекомендуется к необдуманному использованию.

#### Test

| Method Signature| Meaning |
|---|---|
|`T.WithBeforeTest(func())`|Все действия, производимые в прокинутой функции, будут логироваться в контейнер теста в раздел `Set Up`|
|`T.WithAfterTest(func())`|Все действия, производимые в прокинутой функции, будут логироваться в контейнер теста в раздел `Tear Down`|

#### Suite

| Method Signature| Meaning |
|---|---|
|`T.WithBeforeSuite(func())`| Все действия, производимые в прокинутой функции, будут логироваться в контейнер сьюта в раздел `Set Up`|
|`T.WithAfterSuite(func())`| Все действия, производимые в прокинутой функции, будут логироваться в контейнер сьюта в раздел `Tear Down`|

## HOW TO USE suite

### Behavior

| Method Signature| Meaning |
|---|---|
|`Suite.T() *provider.T`|Возвращает указатель на контекст исполняемого теста|
|`Suite.RunSuite(*provider.T, AllureSuite)`| Позволяет запустить переданный сьют от лица родительского сьюта.|
|`Suite.RunTest(string, func(*provider.T), ...string) bool`| Позволяет запустить переданный тест в контексте `*provider.T`. **Эти тесты могут быть параллельными**. Использовать сьют, как провайдер allure для таких тестов **НЕЛЬЗЯ**.|
|`Suite.Run(string, func(), ...string) bool`| Позволяет запустить переданный тест в контексте текущего теста. **Эти тесты __НЕ(!)__ могут быть параллельными**. Можно использовать сьют, как провайдер allure. |
|`Suite.SkipOnPrint()`|Позволяет пропустить вывод allure.Result для текущего теста.|

### Suite Before/Afters

Следующие методы нужно переопределить в Вашем сьюте, чтобы повлиять на поведение исполнения тестов. [Пример](#setup-hooksexamplessuite_demobefores_afters_testgo)

| Method Signature| Meaning |
|---|---|
|`Suite.BeforeEach()`|Этот метод будет исполнятся **перед КАЖДЫМ** запуском теста.|
|`Suite.AfterEach()`|Этот метод будет исполнятся **после КАЖДОГО** окончания теста.|
|`Suite.BeforeSuite()`|Этот метод исполнится **ОДИН РАЗ** перед запуском сьюта.|
|`Suite.AfterSuite()`|Этот метод исполнится **ОДИН РАЗ** после окончания сьюта.|

### Allure forward

В структуру Suite прокинуты все методы `provider.T` для взаимодействия с allure-отчетом. Полный список прокинутых
методов:

| Method Signature|
|:---:|
|**Allure Info**|
|[`Suite.Title(string)`](t_title)|
|[`Suite.Description(string)`](t_description)|
|**Allure Label**|
|[`Suite.Epic(string)`](t_epic)|
|[`Suite.Feature(string)`](t_feature)|
|[`Suite.Story(string)`](t_story)|
|[`Suite.FrameWork(string)`](t_frameWork)|
|[`Suite.Host(string)`](t_host)|
|[`Suite.Thread(string)`](t_thread)| 
|[`Suite.ID(string)`](t_id)|
|[`Sutie.AddSuiteLabel(string)`](t_addsuitelabel)|
|[`Suite.AddSubSuite(string)`](t_addsubsuite)|
|[`Suite.AddParentSuite(string)`](t_addparentsuite)|
|[`Suite.Severity(string)`](t_severity)|
|[`Suite.Tag(string)`](t_tag)|
|[`Suite.Tags(...string)`](t_tags)|
|[`Suite.Package(string)`](t_package)|
|[`Suite.Owner(string)`](t_owner)|
|[`Suite.Label(string, string)`](t_label)|
|**Allure Links**|
|[`Suite.SetIssue(string)`](t_setissue)|
|[`Suite.SetTestCase(string)`](t_settestcase)|
|[`Suite.Link(allure.Link)`](t_link)|
|**Allure Actions**|
|[`Suite.Step(*allure.Step)`](t_step)|
|[`Suite.NewStep(string)`](t_newstep)|
|[`Suite.InnerStep(*allure.Step, *allure.Step)`](t_innerstep)|
|[`Suite.WithStep(*allure.Step, func())`](t_withstep)|
|[`Suite.WithNewStep(string, func())`](t_withnewstep)|
|[`Suite.Attachment(*allure.Attachment)`](t_attachment)|
|[`Suite.AddNestedAttachment(*allure.Attachment)`](t_addnestedattachment)|
|[`Suite.AddParameterToNested(*allure.Parameter)`](t_addparametertonested)|
|[`Suite.AddParametersToNested([]allure.Parameter)`](t_addparameterstonested)|
|[`Suite.AddNewParameterToNested(string, string)`](t_addnewparametertonested)|
|[`Suite.AddNewParametersToNested(...string)`](t_addnewparameterstonested)|

## HOW TO USE allure

### Allure Objects

#### Steps

##### Static functions

| Method Signature| Meaning |
|---|---|
|`NewStep(string, allure.Status, int64, int64, []allure.Parameter) *allure.Step`|возвращает указатель на `allure.Step`. Принимает имя, статус, время старта, время окончания и массив `allure.Parameter`|
|`NewSimpleStep(string) *allure.Step`|возвращает указатель на `allure.Step`. Принимает имя шага. Все остальные аттрибуты конструктора `allure.Step` останутся `0`, `""` и `nil` соответственно|
|`NewSimpleInnerStep(string, *allure.Step) *allure.Step`|возвращает указатель на `allure.Step`, у которого указан второй аргумент в качестве родителя|
|`NewStepWithStart(string) *allure.Step`|возвращает указатель на `allure.Step`, у которого проставлен `Start` как Now()|

##### Step functions

| Method Signature| Meaning |
|---|---|
|`Step.Attachment(*allure.Attachment)`|добавляет `allure.Attachment` к шагу.|
|`Step.AddParameter(allure.Parameter)`|добавляет `allure.Parameter` к шагу.|
|`Step.AddParameters(...allure.Parameter)`|добавляет массив `allure.Parameter` к шагу|
|`Step.AddNewParameter(string, string)`|добавляет `allure.Parameter` к шагу с именем первого аргумента и значением второго|
|`Step.AddNewParameters(...string)`|добавляет массив `allure.Parameter` к шагу. Каждая четная строка - имя нового аргумента, каждая нечетная - его значение. Если передано нечетное количество аргументов, последняя строка откинется.|
|`Step.WithAttachment(*allure.Attachment) *allure.Step`|добавляет `allure.Attachment` к шагу и возвращает указатель.|
|`Step.WithParemeter(allure.Paramter) *allure.Step`|добавляет `allure.Parameter` к шагу и возвращает указатель.|
|`Step.WithParameters(...allure.Parameter) *allure.Step`|добавляет массив `allure.Parameter` к шагу и возвращает указатель.|
|`Step.WithNewParameter(string, string) *allure.Step`|добавляет `allure.Parameter` к шагу с именем первого аргумента и значением второго. Возвращает указатель.|
|`Step.WithNewParameters(...string) *allure.Step`|добавляет массив `allure.Parameter` к шагу. Каждая четная строка - имя нового аргумента, каждая нечетная - его значение. Если передано нечетное количество аргументов, последняя строка откинется. Возвращает указатель.|
|`Step.WithStart() *allure.Step`|проставляет `allure.Step.Start` для текущего шага как `allure.GetNow()`. Возвращает указатель.|
|`Step.WithStop() *allure.Step`|проставляет `allure.Step.Stop` для текущего шага как `allure.GetNow()`. Возвращает указатель.|
|`Step.WithParent(*Step) *Step`|проставляет `allure.Step.Parent` для текущего шага как `Step.GetUUID()` от аргумента. Возвращает указатель.|
|`Step.Passed() *Step`|проставляет `allure.Step.Status` для текущего шага как `allure.Passed()`.|
|`Step.Failed() *Step`|проставляет `allure.Step.Status` для текущего шага как `allure.Failed()`.|
|`Step.Skipped() *Step`|проставляет `allure.Step.Status` для текущего шага как `allure.Skipped()`.|

#### Attachments

##### Static functions

| Method Signature| Meaning |
|---|---|
|`NewAttachment(string, MimeType, []byte) *Attachment`|принимает имя документа, его mimeType и контент в байтах. Возвращает указатель на `allure.Attachment`.|

##### Attachment functions

| Method Signature| Meaning |
|---|---|
|`Attachment.Print()`|распечатывает `allure.Attachment` в файл.|

TODO:

~~1. Ронять шаг если внутри была паника или ошибка~~ Done<br>
~~2. Релиз версии 0.0.1~~ Done<br>
~~3. Отрефачить~~ Done<br>
~~5. Научиться получать имя goroutine~~ Done<br>
~~6. Научиться получать имя пакета для запуска тестов~~ Done<br>
~~7. Readme.MD~~ Done<br>
~~8. Добить до состояния плагина (уже почти)~~ (не наш путь!) Cancelled <br>
~~9. Добиться частичной параллелизации~~ Done<br>
~~10. Релиз версии 0.1.0~~ Done <br>
~~11. Отделить зерна от плевел. Позволить получать аллюр без сьюта.~~ Done <br>

1. Дописать комменты к интерфейсам <br>
2. Довести до ума Before/After вне сьютов <br>
3. Больше параллельности! <br>
4. ???
5. ~~Profit!~~ релиз версии 1.0.0
