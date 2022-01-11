# allure

Пакет allure предлагает имплементацию всех сущностей, используемых Allure для работы с тест-репортами.<br>
Узнать больше про Allure Framework можно [здесь](https://docs.qameta.io/allure/).

## Head of contents

- [head of contents](#head-of-contents)
- [Global environments keys](#global-environment-keys)
- [interfaces](#allure-interfaces)
    - [Printable](#printable)
    - [WithAttachments](#withattachments)
    - [WithTimer](#withtimer)
    - [WithSteps](#withsteps)
- [attachments](#allureattachment)
    - [func NewAttachment(name string, mimeType MimeType, content []byte) *Attachment](#func-newattachmentname-string-mimetype-mimetype-content-byte-attachment)
    - [func (a *Attachment) Print() error](#func-a-attachment-print-error)
- [container](#allurecontainer)
    - [func NewContainer() *Container](#func-newcontainer-container)
    - [func (container *Container) AddChild(childUUID uuid.UUID)](#func-container-container-addchildchilduuid-uuiduuid)
    - [func (container *Container) IsEmpty() bool](#func-container-container-isempty-bool)
    - [func (container *Container) Print() error](#func-container-container-print-error)
    - [func (container *Container) PrintAttachments()](#func-container-container-printattachments)
    - [func (container *Container) MatchSteps()](#func-container-container-matchsteps)
    - [func (container *Container) Begin()](#func-container-container-begin)
    - [func (container *Container) Finish()](#func-container-container-finish)
- [label](#allurelabel)
    - [Supported label types](#supported-label-types)
    - [func NewLabel(labelType LabelType, value string) Label](#func-newlabellabeltype-labeltype-value-string-label)
    - [func LanguageLabel(language string) Label](#func-languagelabellanguage-string-label)
    - [func FrameWorkLabel(framework string) Label](#func-frameworklabelframework-string-label)
    - [func IDLabel(testID string) Label](#func-idlabeltestid-string-label)
    - [func TagLabel(tag string) Label](#func-taglabeltag-string-label)
    - [func TagLabels(tags ...string) []Label](#func-taglabelstags-string-label)
    - [func HostLabel(host string) Label](#func-hostlabelhost-string-label)
    - [func ThreadLabel(thread string) Label](#func-threadlabelthread-string-label)
    - [func SeverityLabel(severity SeverityType) Label](#func-severitylabelseverity-severitytype-label)
    - [func SubSuiteLabel(subSuite string) Label](#func-subsuitelabelsubsuite-string-label)
    - [func ParentSuiteLabel(parent string) Label](#func-parentsuitelabelparent-string-label)
    - [func SuiteLabel(suite string) Label](#func-suitelabelsuite-string-label)
    - [func PackageLabel(packageName string) Label](#func-packagelabelpackagename-string-label)
    - [func OwnerLabel(ownerName string) Label](#func-ownerlabelownername-string-label)
- [link](#allurelink)
    - [Supported link types](#supported-link-types)
    - [func NewLink(name string, _type LinkTypes, url string) Link](#func-newlinkname-string-_type-linktypes-url-string-link)
    - [func TestCaseLink(testCase string) Link](#func-testcaselinktestcase-string-link)
    - [func IssueLink(issue string) Link](#func-issuelinkissue-string-link)
    - [func LinkLink(linkname, link string) Link](#func-linklinklinkname-link-string-link)
- [parameter](#allureparameter)
    - [func NewParameter(name string, value string) Parameter](#func-newparametername-string-value-string-parameter)
    - [func NewParameters(kv ...string) []Parameter](#func-newparameterskv-string-parameter)
- [result](#allureresult)
    - [func NewResult(testName, fullName string) *Result](#func-newresulttestname-fullname-string-result)
    - [func (result *Result) GetLabel(labelType LabelType) []Label](#func-result-result-getlabellabeltype-labeltype-label)
    - [func (result *Result) SetLabel(label Label)](#func-result-result-setlabellabel-label)
    - [func (result *Result) SetLabels(labels ...Label)](#func-result-result-setlabelslabels-label)
    - [func (result *Result) WithParentSuite(parentName string) *Result](#func-result-result-withparentsuiteparentname-string-result)
    - [func (result *Result) WithSuite(suiteName string) *Result](#func-result-result-withsuitesuitename-string-result)
    - [func (result *Result) WithFrameWork(framework string) *Result](#func-result-result-withsubsuiteschildren-string-result)
    - [func (result *Result) WithFrameWork(framework string) *Result](#func-result-result-withframeworkframework-string-result)
    - [func (result *Result) WithThread(thread string) *Result](#func-result-result-withthreadthread-string-result)
    - [func (result *Result) WithPackage(pkg string) *Result](#func-result-result-withpackagepkg-string-result)
    - [func (result *Result) WithLaunchTags() *Result](#func-result-result-withlaunchtags-result)
    - [func (result *Result) Begin()](#func-result-result-begin)
    - [func (result *Result) Finish()](#func-result-result-finish)
    - [func (result *Result) Done()](#func-result-result-done)
    - [func (result *Result) SkipOnPrint()](#func-result-result-skiponprint)
    - [func (result *Result) Print() error](#func-result-result-print-error)
    - [func (result *Result) MatchSteps()](#func-result-result-matchsteps)
    - [func (result *Result) PrintAttachments()](#func-result-result-printattachments)
- [step](#allurestep)
    - [func NewStep(name string, status Status, start int64, stop int64, parameters []Parameter) *Step](#func-newstepname-string-status-status-start-int64-stop-int64-parameters-parameter-step)
    - [func NewSimpleStep(name string) *Step](#func-newsimplestepname-string-step)
    - [func NewSimpleInnerStep(name string, parent *Step) *Step](#func-newsimpleinnerstepname-string-parent-step-step)
    - [func NewStepWithStart(name string) *Step](#func-newstepwithstartname-string-step)
    - [func (st *Step) GetUUID() string](#func-st-step-getuuid-string)
    - [func (st *Step) Attachment(attachment *Attachment)](#func-st-step-attachmentattachment-attachment)
    - [func (st *Step) AddParameter(param Parameter)](#func-st-step-addparameterparam-parameter)
    - [func (st *Step) AddNewParameter(key, value string)](#func-st-step-addnewparameterkey-value-string)
    - [func (st *Step) AddNewParameters(kv ...string)](#func-st-step-addnewparameterskv-string)
    - [func (st *Step) WithParent(parent *Step) *Step](#func-st-step-withparentparent-step-step)
    - [func (st *Step) WithStart() *Step](#func-st-step-withstart-step)
    - [func (st *Step) WithStop() *Step](#func-st-step-withstop-step)
    - [func (st *Step) Passed() *Step](#func-st-step-passed-step)
    - [func (st *Step) Failed() *Step](#func-st-step-failed-step)
    - [func (st *Step) Skipped() *Step](#func-st-step-skipped-step)
    - [func (st *Step) Begin()](#func-st-step-begin)
    - [func (st *Step) Finish()](#func-st-step-finish)
    - [func (st *Step) PrintAttachments()](#func-st-step-printattachments)

## Global Environment Keys

| Key | Meaning | Default |
|---|---|---|
|`ALLURE_OUTPUT_PATH`|Указывает путь до папки для печати результатов.|`.` (Папка с тестами)|
|`ALLURE_OUTPUT_FOLDER`|Указывает имя папки для печати результатов.|`/allure-results`|
|`ALLURE_ISSUE_PATTERN`|Указывает URL паттерн для Issue. **Обязательно должен содержать один `%s`**| |
|`ALLURE_TESTCASE_PATTERN`|Указывает URL паттерн для TestCase. **Обязательно должен содержать один `%s`**.| |
|`ALLURE_LAUNCH_TAGS`|Указывает дефолтные тэги, которыми будут помечаться все тесты в прогоне. Тэги должны быть указаны через запятую.| |

## allure interfaces

### `Printable`

Интерфейс `Printable` обозначает, что объект имеет файловый эквивалент в allure.

Методы:

- `Print() error` - создает файл в файловой системе.

Интерфейс реализуют:

- `allure.Attachment`
- `allure.Result`
- `allure.Container`

### `WithAttachments`

Интерфейс обозначает, что объект может иметь `allure.Attachment`. По скольку, `allure.Attachment` реализуют `Printable`,
должна быть возможность удобно распечатать все `allure.Attachment` сущности разом.

Методы:

- `PrintAttachments()` - вызывает у каждого `Attachment` метод `Print()`

Реализуют:

- `allure.Result`
- `allure.Container`
- `allure.Step`

### `WithSteps`

Обозначает, что объект имеет `allure.Steps`. На данный момент реализация вложенности такова, что шаги, это список,
порядок которого гарантирован экосистемой go. Соответственно каждый ребенок лежит __сразу__ после родителя. Чтобы
превратить линейный список в дерево - реализуйте метод MatchSteps().

Методы:

- `MatchSteps()` - превращает список шагов в дерево.

Реализуют:

- `allure.Result`
- `allure.Container`
- `allure.Step`

### `WithTimer`

Обозначает, что объект имеет некое измерение времени.

Методы:

- `Begin()` - обозначает начало исполнения
- `Finish()` - обозначает конец исполнения

Реализуют:

- `allure.Result`
- `allure.Container`
- `allure.Step`

## [allure.Attachment](attachment.go)

Описание структуры:

```go
package allure

type IAttachment interface {
	Printable
}

// Attachment ...
type Attachment struct {
	Name    string   `json:"name"`   // Имя приложения
	Source  string   `json:"source"` // Путь до файла приложения (имя)
	Type    MimeType `json:"type"`   // Mime-type приложения
	uuid    string   // Уникальный идентификатор приложения
	content []byte   // Содержимое приложения в байтах 
}

// MimeType is Attachment's mime type
type MimeType string
```

`allure.Attachment` - является имплементацией приложений к отчету в allure. Чаще всего используется, чтобы содержать
скриншоты, ответы, файлы и другие данные, полученные во время выполнения теста.

Реализует интерфейсы:

- IAttachment
- [Printable](#printable)

### Supported types

| Key | Mime type | File type |
|:---:|:---:|---|
|Text|"text/plain"|`.txt`|
|Csv|"text/csv"|`.csv`|
|Tsv|"text/tab-separated-values"|`.tsv`|
|URIList|"text/uri-list"|`.uri`|
|HTML|"text/html"|`.html`|
|XML|"application/xml"|`.xml`|
|JSON|"application/json"|`.json`|
|Yaml|"application/yaml"|`.yaml`|
|Pcap|"application/vnd.tcpdump.pcap"|`.pcap`|
|Png|"image/png"|`.png`|
|Jpg|"image/jpg"|`.jpg`|
|Svg|"image/svg-xml"|`.svg`|
|Gif|"image/gif"|`.gif`|
|Bmp|"image/bmp"|`.bmp`|
|Tiff|"image/tiff"|`.tiff`|
|Mp4|"video/mp4"|`.mp4`|
|Ogg|"video/ogg"|`.ogg`|
|Webm|"video/webm"|`.webm`|
|Mpeg|"video/mpeg"|`.mpeg`|
|Pdf|"application/pdf"|`.pdf`|

### `func NewAttachment(name string, mimeType MimeType, content []byte) *Attachment`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя нового attachment'а|
|`mimeType MimeType`|mime type нового attachment'а|
|`content []byte`|содержимое нового attachment'а|

#### return value

| Return Value | Meaning |
|---|---|
|`*Attachment`|указатель на новый attachment|

Конструктор. Возвращает указатель на новый объект attachment.

### `func (a *Attachment) Print() error`

| Return Value | Meaning |
|---|---|
|`error`|если создать файл не удалось - возвращает ошибку|

Создает файл из `Attachment.content`. Тип файла определяется его `Attachment.mimeType`.

## [allure.Container](container.go)

Описание структуры:

```go
package allure

type IContainer interface {
	Printable
	WithAttachments
	WithSteps
	WithTimer
	AddChild(childUUID uuid.UUID)
	IsEmpty() bool
}

// Container ...
type Container struct {
	UUID         uuid.UUID    `json:"uuid"`     // Уникальный идентификатор контейнера
	Children     []uuid.UUID  `json:"children"` // Массив uuid в котором указаны все отчеты, ссылающиеся на контейнер
	Befores      []*Step      `json:"befores"`  // Массив шагов в Test Setup
	Afters       []*Step      `json:"afters"`   // Массив шагов в Test Teardown
	Start        int64        `json:"start"`    // Время начала работы контейнера
	Stop         int64        `json:"stop"`     // Время завершения работы контейнера
	BeforesQueue NestingQueue `json:"-"`        // Очередь вложенных шагов в Befores
	AftersQueue  NestingQueue `json:"-"`        // Очередь вложенных шагов в Afters
}
```

Является реализацией сущности `Container`, используемой Allure для работы с хуками TestSetup и TestTeardown. Список
зависимых от контейнера тестов содержится в массиве `Container.Children`.

Реализует интерфейсы:

- IContainer
- [Printable](#printable)
- [WithAttachments](#withattachments)
- [WithSteps](#withsteps)
- [WithTimer](#withtimer)

### `func NewContainer() *Container`

| Return Value | Meaning |
|---|---|
|`*Container`|указатель на новый объект `allure.Container`|

Конструктор. Собирает и возвращает новый объект `allure.Container`.

### `func (container *Container) AddChild(childUUID uuid.UUID)`

#### arguments

| Argument | Meaning |
|---|---|
|`childUUID uuid.UUID`|uuid объекта allure.Result, к которому относится Container|

Добавляет к массиву `container.Children` нового потомка.

### `func (container *Container) IsEmpty() bool`

#### return value

| Return Value | Meaning |
|---|---|
|`bool`|возвращает `true`, если контейнер считается пустым|

Возвращает `true`, если массивы `container.Befores` и `container.Afters` пусты.

### `func (container *Container) Print() error`

| Return Value | Meaning |
|---|---|
|`error`|если создать файл не удалось - возвращает ошибку|

Проверяет файл с помощью функции [`container.IsEmpty`](#func-container-container-isempty-bool):

1) если контейнер пустой, исполнение функции завершается без ошибки.
2) если контейнер содержит шаги
    1) Дергает [`container.PrintAttachments`](#func-container-container-printattachments)
    2) Дергает [`container.MatchSteps`](#func-container-container-matchsteps)
    3) Сериализует файл в `uuid4-container.json`.
    4) Создает файл в файловой системе в папке вывода (`$ALLURE_OUTPUT_PATH`/`$ALLURE_OUTPUT_FOLDER`). Если вовремя
       исполнения произошла ошибка - возвращает ее

### `func (container *Container) PrintAttachments()`

Проходится по всем `Container.Befores` и `Container.Afters` контейнера и у каждого `allure.Step` вызывает
метод `Step.PrintAttachments()`

### `func (container *Container) MatchSteps()`

Собирает шаги в дерево как для массива `Container.Befores`, так и для `Container.Afters`.

### `func (container *Container) Begin()`

Устанавливает `Container.Start` = `time.GetNow()`

### `func (container *Container) Finish()`

Устанавливает `Container.Stop` = `time.GetNow()`

## [allure.Label](label.go)

Описание структуры:

```go
package allure

type ILabel interface {
}

// Label ...
type Label struct {
	Name  string `json:"name"`  // Имя лейбла
	Value string `json:"value"` // Значение лейбла
}

type LabelType string

type SeverityType string
```

`allure.Label` - является имплементацией лейбла. Лейбл - сущность, используемая Allure для составления метрик и
группировки тестов.

Реализует интерфейс:

- ILabel

### Supported Label Types

```
Epic        
Feature     
Story       
ID          
Severity    
ParentSuite 
Suite       
SubSuite    
Package     
Thread      
Host        
Tag         
Framework   
Language    
Owner       
```

### `func NewLabel(labelType LabelType, value string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`labelType LabelType`|тип лейбла (поддерживаемые типы смотри [здесь](#supported-label-types))|
|`value string`|Значение лейбла|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label`|

Собирает и возвращает новый `allure.Label`. Ключ лейбла зависит от переданного `labelType`

### `func LanguageLabel(language string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`language string`|язык, на котором написаны тесты. Например - `go1.16`|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Language`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Language`,
а`value` = `language`

### `func FrameWorkLabel(framework string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`framework string`|framework, на котором запущены тесты. (например - `allure-testify@v0.x.x`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Framework`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Framework`
., `value` = `framework`

### `func IDLabel(testID string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`testID string`|ID теста в TMS.|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.ID`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.ID`
., `value` = `id`

### `func TagLabel(tag string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`tag string`|tag теста. (Например - `parametrized`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Tag`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Tag`,
а `value` = `tag`

### `func TagLabels(tags ...string) []Label`

#### arguments

| Argument | Meaning |
|---|---|
|`tags ...string`|tag'и теста. (Например - `parametrized, functional, e2e`)|

#### return value

| Return Value | Meaning |
|---|---|
|`[]Label`|slice новых `allure.Label` с типом `allure.Tag`|

Для каждого tag из слайса tags вызывает [`TagLabel`](#func-taglabeltag-string-label), прокидывая `value` = `tag` для
каждого элемента слайса.

### `func HostLabel(host string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`host string`|host, с которого запущены тесты. (Например - `mbp-admin`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Host`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Host`,
а `value` = `host`

### `func ThreadLabel(thread string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`thread string`|thread теста, в котором прошел тест. (Например - `MainThread`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Thread`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Thread`,
а `value` = `thread`

### `func SeverityLabel(severity SeverityType) Label`

#### Supported Severity Types

```
BLOCKER  
CRITICAL 
NORMAL   
MINOR    
TRIVIAL  
```

#### arguments

| Argument | Meaning |
|---|---|
|`severity SeverityType`|severity теста. (Например - `allure.Trivial`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Severity`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Severity`,
а `value` = `severity.String()`

### `func SubSuiteLabel(subSuite string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`subSuite string`|subSuite теста, в котором прошел тест. (Например - `Some SubSuite`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.SubSuite`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.SubSuite`,
а `value` = `subSuite`

### `func EpicLabel(epic string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`epic string`|epic теста, в котором прошел тест. (Например - `Some Epic`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Epic`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Epic`,
а `value` = `epic`

### `func StoryLabel(story string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`story string`|epic теста, в котором прошел тест. (Например - `Some Story`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Story`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Story`,
а `value` = `story`

### `func FeatureLabel(feature string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`feature string`|feature теста, в котором прошел тест. (Например - `Some Feature`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Feature`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Feature`,
а `value` = `feature`

### `func ParentSuiteLabel(parent string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`parent string`|parent suite теста, в котором прошел тест. (Например - `Some Parent Suite`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.ParentSuite`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType`
= `allure.ParentSuite`, а `value` = `parent`

### `func SuiteLabel(suite string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`suite string`|parent suite теста, в котором прошел тест. (Например - `Some Suite`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Suite`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Suite`
, а `value` = `suite`

### `func PackageLabel(packageName string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`packageName string`|parent suite теста, в котором прошел тест. (Например - `some/package/test`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Package`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Package`
, а `value` = `package`

### `func OwnerLabel(ownerName string) Label`

#### arguments

| Argument | Meaning |
|---|---|
|`ownerName string`|owner suite теста, в котором прошел тест. (Например - `jdoe@acme.com`)|

#### return value

| Return Value | Meaning |
|---|---|
|`Label`|новый `allure.Label` с типом `allure.Owner`|

Вызывает [`NewLabel`](#func-newlabellabeltype-labeltype-value-string-label), прокидывая `labelType` = `allure.Owner`
, а `value` = `ownerName`

## [allure.Link](link.go)

Описание структуры:

```go
package allure

type ILink interface {
}

// Link ...
type Link struct {
	Name string `json:"name"` // Имя ссылки
	Type string `json:"type"` // Тип ссылки
	URL  string `json:"url"`  // URL ссылки
}

// LinkTypes ...
type LinkTypes string
```

`allure.Link` - является реализацией сущности Link, используемой Allure для указания ссылок, необходимых для тестовых
отчетностей.

Таких, как:

- ссылка на задачу в Issue-трекере.
- ссылка на тест кейс в TMS
- Любая другая ссылка (например, ссылка на pod окружения)

Реализует интерфейсы:

- ILink

### Supported Link Types

```
LINK    
ISSUE   
TESTCASE
```

### `func NewLink(name string, _type LinkTypes, url string) Link`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя новой ссылки|
|`_type LinkTypes`|тип ссылки|
|`url string`|url, на который ссылается объект Link|

#### return value

| Return Value | Meaning |
|---|---|
|`Link`|Новый объект `Link`|

Конструктор. Собирает и возвращает новый объект `allure.Link`.

### `func TestCaseLink(testCase string) Link`

#### arguments

| Argument | Meaning |
|---|---|
|`testCase string`|урл/id тест-кейса в TMS|

#### return value

| Return Value | Meaning |
|---|---|
|`Link`|Новый объект `Link` с `LinkType`=`TESTCASE`|

Собирает объект ссылки, указывающий на описание тест-кейса в вашей TMS.<br>
Если `ALLURE_TESTCASE_PATTERN` не заполнен, берет переданный аргумент как полный url. Иначе - пытается подставить
переданный аргумент в строку `ALLURE_TESTCASE_PATTERN`.

### `func IssueLink(issue string) Link`

#### arguments

| Argument | Meaning |
|---|---|
|`issue string`|урл/id issue в системе контроля кейсов|

#### return value

| Return Value | Meaning |
|---|---|
|`Link`|Новый объект `Link` с `LinkType`=`ISSUE`|

Собирает объект ссылки, указывающий на описание тест-кейса в вашей TMS.<br>
Если `ALLURE_ISSUE_PATTERN` не заполнен, берет переданный аргумент как полный url. Иначе - пытается подставить
переданный аргумент в строку `ALLURE_ISSUE_PATTERN`.

### `func LinkLink(linkname, link string) Link`

#### arguments

| Argument | Meaning |
|---|---|
|`linkname string`|имя ссылки|
|`link string`|url ссылки|

#### return value

| Return Value | Meaning |
|---|---|
|`Link`|Новый объект `Link` с `LinkType`=`LINK`|

Собирает объект ссылки, используя `linkName` в качестве имени и `link` в качестве url.

## [allure.Parameter](parameter.go)

Описание структуры:

```go
package allure

type IParameter interface {
}

// Parameter ...
type Parameter struct {
	Name  string `json:"name"`  // Имя параметра
	Value string `json:"value"` // Значение параметра
}
```

`allure.Parameter` - является имплементацией сущности Parameter, которую Allure использует как дополнительную
информацию, описывающую шаг теста (например - хост запроса или адрес сервера)

Реализует интерфейс:

- IParameter

### `func NewParameter(name string, value string) Parameter`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя параметра|
|`value string`|значение параметра|

#### return value

| Return Value | Meaning |
|---|---|
|`Parameter`|Новый объект `Parameter`|

Конструктор. Собирает и возвращает новый объект `Parameter`, используя `name` в качестве имени параметра, а `value`, в
качестве значения.

### `func NewParameters(kv ...string) []Parameter`

#### arguments

| Argument | Meaning |
|---|---|
|`kv ...string`|пары ключ/значение|

#### return value

| Return Value | Meaning |
|---|---|
|`[]Parameter`|массив новых объектов `Parameter`, полученных после обработки массива входных строк|

Конструктор. Принимает в себя список строк, через запятую. Каждая четная строка считается именем параметра, а каждая
нечетная - значением параметра. Если передано нечетное количество строк, последняя строка отбрасывается.<br>
Возвращает список параметров, полученных после обработки переданного списка.

## [allure.Result](result.go)

Описание структуры:

```go
package allure

type IResult interface {
	Printable
	WithSteps
	WithTimer
	WithParentSuite(parentName string) *Result
	WithFrameWork(framework string) *Result
	WithThread(thread string) *Result
	WithLanguage(language string) *Result
	WithPackage(pkg string) *Result
	WithSuite(suiteName string) *Result
	WithSubSuites(children ...string) *Result
}

type Result struct {
	Name          string        `json:"name"`          // Имя теста
	FullName      string        `json:"fullName"`      // Полный путь до теста
	Status        Status        `json:"status"`        // Статус выполнения теста
	StatusDetails StatusDetail  `json:"statusDetails"` // Подробности о тесте (например, ошибки при выполнении будут записаны сюда)
	Start         int64         `json:"start"`         // Начало выполнения теста
	Stop          int64         `json:"stop"`          // Окончание выполнения теста
	UUID          uuid.UUID     `json:"uuid"`          // Уникальный идентификатор теста
	HistoryID     string        `json:"historyId"`     // ID в истории allure
	TestCaseID    string        `json:"testCaseId"`    // ID тест кейса (основан на хеше полного вызова)
	Description   string        `json:"description"`   // Описание теста
	Attachments   []*Attachment `json:"attachments"`   // Приложения к тесту
	Labels        []Label       `json:"labels"`        // Массив лейблов
	Links         []Link        `json:"links"`         // Массив ссылок
	Steps         []*Step       `json:"steps"`         // Массив шагов
	StepsQueue    NestingQueue  `json:"-"`             // Очередь вложенности
	NestedSteps   []string      `json:"-"`             // Массив содержащий все текущие uuid.UUID незавершенных вложенных шагов 
	Container     *Container    `json:"-"`             // Контейнер для Before/After Test хука
	toPrint       bool          // Если false - отчет не будет сохранен в файл
}

type Status string

type StatusDetail struct {
	Message string `json:"message"` // Укороченная версия сообщения
	Trace   string `json:"trace"`   // Полное сообщение
}
```

`allure.Result` - является имплементацией сущности Result, используемой Allure для хранения информации о тесте. Содержит
в себе информацию об имени теста, приложениях, описании, статусе, ссылках, лейблах, шагах, контейнерах и времени
исполнения теста.

Реализует интерфейсы:

- IResult
- [Printable](#printable)
- [WithAttachments](#withattachments)
- [WithSteps](#withsteps)
- [WithTimer](#withtimer)

### `func NewResult(testName, fullName string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`testName string`|имя теста|
|`fullName string`|путь до теста|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на новый объект `allure.Result`|

Конструктор Собирает новый `allure.Result`. Проставляет дефолтные значения для структуры.

|Field Value|Default|
|---|---|
|UUID|random `uuid4` value|
|Name|testName from args|
|FullName|fullName from args|
|TestCaseID|md5 hash of `Result.FullName`|
|HistoryID|md5 hash from `result.TestCaseID`|
|Container|new empty `allure.Container`|
|StepsQueue|new `StepQueue` object|
|Labels|add new `allure.Language` label|
|Start|`time.Now()`|
|toPrint|`true`|

Проставляет child для объекта container.

### `func (result *Result) GetLabel(labelType LabelType) []Label`

#### arguments

| Argument | Meaning |
|---|---|
|`labelType LabelType`|тип лейбла|

#### return value

| Return Value | Meaning |
|---|---|
|`[]Label`|список лейблов с указанным `LabelType`|

Возвращает все `allure.Label`, у которых `LabelType` совпадает с указанным в аргументе.

### `func (result *Result) SetLabel(label Label)`

#### arguments

| Argument | Meaning |
|---|---|
|`label Label`|лейбл, который нужно прикрепить к отчету|

Добавляет переданный в аргументах `allure.Label` к отчету

### `func (result *Result) SetLabels(labels ...Label)`

#### arguments

| Argument | Meaning |
|---|---|
|`labels ...Label`|лейблы, который нужно прикрепить к отчету|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет все переданные в аргументах `allure.Label` к отчету. Возвращает указатель на текущий `allure.Result` (для
Fluent Interface).

### `func (result *Result) WithParentSuite(parentName string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`parentName string`|имя сьюта-родителя|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `ParentSuite` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent
Interface).

### `func (result *Result) WithSuite(suiteName string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`suiteName string`|имя сьюта|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `Suite` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent
Interface).

### `func (result *Result) WithHost(hostName string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`hostName string`|имя хоста|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `Host` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent Interface)
.

### `func (result *Result) WithSubSuites(children ...string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`children ...string`|имена дочерних сьютов|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет все `allure.Label` (для каждого из списка `children`) с типом `SubSuite` к отчету. Возвращает указатель на
текущий `allure.Result` (для Fluent Interface).

### `func (result *Result) WithFrameWork(framework string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`framework string`|имя фреймворка|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `FrameWork` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent
Interface).

### `func (result *Result) WithLanguage(language string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`language string`|название языка|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `Language` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent
Interface).

### `func (result *Result) WithThread(thread string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`thread string`|имя текущего потока|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `Thread` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent
Interface).

### `func (result *Result) WithPackage(pkg string) *Result`

#### arguments

| Argument | Meaning |
|---|---|
|`pkg string`|имя пакета|

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет `allure.Label` с типом `Package` к отчету. Возвращает указатель на текущий `allure.Result` (для Fluent
Interface).

### `func (result *Result) WithLaunchTags() *Result`

#### return value

| Return Value | Meaning |
|---|---|
|`*Result`|указатель на текущий `allure.Result`|

Добавляет все Launch Tags из глобальной переменной `ALLURE_LAUNCH_TAGS` как лейблы с типом `Tag` к отчету. Возвращает
указатель на текущий `allure.Result` (для Fluent Interface).

### `func (result *Result) Begin()`

Устанавливает `result.Start` как текущее время

### `func (result *Result) Finish()`

Устанавливает `result.Stop` как текущее время

### `func (result *Result) Done()`

Проверяет статус отчета. Если `result.Status` не заполнено, считаем, что тест успешно завершен (нет ошибок).

### `func (result *Result) SkipOnPrint()`

Ставит переменную `result.toPrint` как true.

### `func (result *Result) Print() error`

#### return value

| Return Value | Meaning |
|---|---|
|`error`|если создать файл не удалось - возвращает ошибку|

Если `result.toPrint` = `true` - метод завершается без создания каких либо файлов. Иначе:

- Вызывает `result.PrintAttachments()`
- Вызывает `result.MatchSteps()`
- Сохраняет файл `uuid4-result.json`
- Вызывает `result.Container.Print()`
- Возвращает ошибку (если есть)

### `func (result *Result) MatchSteps()`

Собирает все шаги массива `result.Steps` в дерево (для сохранения вложенности шагов).

### `func (result *Result) PrintAttachments()`

Проходится по всем `result.Steps` отчета и у каждого шага вызывает метод `PrintAttachments()`.<br>
После чего вызывает `Print()` у всех `allure.Attachments` списка `result.Attachments`

## [allure.Step](step.go)

Описание структуры:

```go
package allure

type IStep interface {
	WithAttachments
	WithTimer
	Attachment(*Attachment)
	AddParameter(Parameter)
	AddParameters(...Parameter)
	AddNewParameter(string, string)
	AddNewParameters(...string)

	WithAttachment(*Attachment) *Step
	WithParameter(Parameter) *Step
	WithNewParameter(string, string) *Step
	WithParameters(...Parameter) *Step
	WithNewParameters(...string) *Step
	WithParent(*Step) *Step
	WithStart() *Step
	WithStop() *Step

	Passed() *Step
	Failed() *Step
	Skipped() *Step
}

// Step ...
type Step struct {
	Name        string        `json:"name"`        // Имя шага
	Status      Status        `json:"status"`      // Статус шага 
	Attachments []*Attachment `json:"attachments"` // Приложения шага
	Start       int64         `json:"start"`       // Время начала шага
	Stop        int64         `json:"stop"`        // Время завершения шага
	Steps       []*Step       `json:"steps"`       // Массив вложенных шагов
	Parameters  []Parameter   `json:"parameters"`  // Массив параметров шага
	Parent      string        `json:"-"`           // Родительский uuid
	uuid        string        // Уникальный идентификатор шага
}
```

`allure.Step` - является имплементацией сущности Step, используемой Allure для определения и описания шагов теста. Шаги
могут быть вложены друг в друга, имеют статус (успешен, провален, пропущен, сломан), могут содержать приложения и параметры, а так
же имеют время исполнения.<br> 
Allure-testify предлагает широкие возможности по созданию и изменению шагов, позволяя
собирать красивые и понятные отчеты об исполнении тестов.<br>
Крайне рекомендуется использовать шаги при описании вашего тест-сценария. Это позволяет Вашим тестам быть понятнее и
информативнее не только в отчетах, но и в коде.<br>
Однако, не рекомендуется чрезмерное использование, так как злоупотребление может превратить код Ваших тестов в полотно,
состоящее только из шагов. 


Реализует интерфейсы:

- IStep
- [WithAttachments](#withattachments)
- [WithTimer](#withtimer)

### `func NewStep(name string, status Status, start int64, stop int64, parameters []Parameter) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя шага|
|`status Status`|статус шага|
|`start int64`|время начала шага|
|`stop int64`|время завершения шага|
|`parameters []Parameter`|массив `allure.Parameter` шага |

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на новый объект `allure.Step`|

Конструктор. Создает новый объект `allure.Step` с переданными в аргументах значениями полей и возвращает указатель на
него.

### `func NewSimpleStep(name string) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя шага|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на новый объект `allure.Step`|

Конструктор. Создает объект `Step`, путем вызова `allure.NewStep` с определенными стандартными значениями (за
исключением имени шага)

|Field Value|Default|
|---|---|
|status|`passed`|
|start|`allure.GetNow()`|
|stop|`allure.GetNow()`|
|parameters|nil|

### `func NewSimpleInnerStep(name string, parent *Step) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя шага|
|`parent *Step`|указатель на шаг-родитель|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на новый объект `allure.Step`|

Вызывает `allure.NewSimpleStep`, после чего устанавливает ему `parentUUID` как UUID переданного в аргументе `parent`
шага

### `func NewStepWithStart(name string) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`name string`|имя шага|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на новый объект `allure.Step`|

Конструктор. Создает объект `Step`, путем вызова `allure.NewStep` с определенными стандартными значениями (за
исключением имени шага).

В отличие от `allure.NewSimpleStep`, не заполняет поле Stop.

|Field Value|Default|
|---|---|
|status|`passed`|
|start|`allure.GetNow()`|
|parameters|nil|

### `func (st *Step) GetUUID() string`

#### return value

| Return Value | Meaning |
|---|---|
|`string`|UUID шага|

Возвращает UUID шага

### `func (st *Step) Attachment(attachment *Attachment)`

#### arguments

| Argument | Meaning |
|---|---|
|`attachment *Attachment`|`allure.Attachment`, который будет добавлен к шагу|

Добавляет к массиву `Step.Attachments` переданный в аргументе `allure.Attachment`.

### `func (st *Step) AddParameter(param Parameter)`

#### arguments

| Argument | Meaning |
|---|---|
|`param Parameter`|`allure.Parameter`, который будет добавлен к шагу|

Добавляет к массиву `Step.Parameters` переданный в аргументе `param`.

### `func (st *Step) AddNewParameter(key, value string)`

#### arguments

| Argument | Meaning |
|---|---|
|`key string`|имя параметра, который будет добавлен к шагу|
|`value string`|значение параметра, который будет добавлен к шагу|

Создает новый `allure.Parameters` из переданных аргументов `key` и `value` и добавляет их к массиву `Step.Parameters`.

### `func (st *Step) AddParameters(params ...Parameter)`

#### arguments

| Argument | Meaning |
|---|---|
|`params ...Parameter`| массив `allure.Parameter`, который будет добавлен к шагу|

Добавляет к массиву `Step.Parameters` все `allure.Parameter`, переданные в аргументе `params`.

### `func (st *Step) AddNewParameters(kv ...string)`

#### arguments

| Argument | Meaning |
|---|---|
|`kv ...string`|пары ключ-значение, которые будут преобразованы в параметры|

Принимает в себя список строк, через запятую. Каждая четная строка считается именем параметра, а каждая нечетная -
значением параметра. Если передано нечетное количество строк, последняя строка отбрасывается.<br>
Добавляет к массиву `Step.Parameters` все `allure.Parameter`, полученные после преобразования `kv`.

### `func (st *Step) WithAttachment(attachment *Attachment) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`attachment *Attachment`|`allure.Attachment`, который будет добавлен к шагу|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Добавляет к массиву `Step.Attachments` переданный в аргументе `allure.Attachment`. Возвращает указатель на текущий шаг (
для Fluent Interface).

### `func (st *Step) WithParameter(param Parameter) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`param Parameter`|`allure.Parameter`, который будет добавлен к шагу|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Добавляет к массиву `Step.Parameters` переданный в аргументе `param`. Возвращает указатель на текущий шаг (для Fluent
Interface).

### `func (st *Step) WithParameters(params ...Parameter) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`params ...Parameter`| массив `allure.Parameter`, который будет добавлен к шагу|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Добавляет к массиву `Step.Parameters` все `allure.Parameter`, переданные в аргументе `params`. Возвращает указатель на
текущий шаг (для Fluent Interface).

### `func (st *Step) WithNewParameter(key, value string) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`key string`|имя параметра, который будет добавлен к шагу|
|`value string`|значение параметра, который будет добавлен к шагу|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Создает новый `allure.Parameters` из переданных аргументов `key` и `value` и добавляет их к массиву `Step.Parameters`.
Возвращает указатель на текущий шаг (для Fluent Interface).

### `func (st *Step) WithNewParameters(kv ...string) *Step`

| Argument | Meaning |
|---|---|
|`kv ...string`|пары ключ-значение, которые будут преобразованы в параметры|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Принимает в себя список строк, через запятую. Каждая четная строка считается именем параметра, а каждая нечетная -
значением параметра. Если передано нечетное количество строк, последняя строка отбрасывается.<br>
Добавляет к массиву `Step.Parameters` все `allure.Parameter`, полученные после преобразования `kv`. Возвращает указатель
на текущий шаг (для Fluent Interface).

### `func (st *Step) WithParent(parent *Step) *Step`

#### arguments

| Argument | Meaning |
|---|---|
|`parent *Step`|указатель на шаг-родитель|

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Устанавливает шагу `parentUUID`, как UUID переданного в аргументе `parent` шага. Возвращает указатель на текущий шаг (
для Fluent Interface).

### `func (st *Step) WithStart() *Step`

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Проставляет `Step.Start` = `GetNow()`
Возвращает указатель на текущий шаг (для Fluent Interface).

### `func (st *Step) WithStop() *Step`

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Проставляет `Step.Stop` = `GetNow()`
Возвращает указатель на текущий шаг (для Fluent Interface).

### `func (st *Step) Passed() *Step`

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Проставляет `Step.Status` = `passed`
Возвращает указатель на текущий шаг (для Fluent Interface).

### `func (st *Step) Failed() *Step`

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Проставляет `Step.Status` = `failed`
Возвращает указатель на текущий шаг (для Fluent Interface).

### `func (st *Step) Skipped() *Step`

#### return value

| Return Value | Meaning |
|---|---|
|`*Step`|указатель на текущий объект `allure.Step`|

Проставляет `Step.Status` = `skipped`
Возвращает указатель на текущий шаг (для Fluent Interface).

### `func (st *Step) Begin()`

Проставляет `Step.Start` = `GetNow()`

### `func (st *Step) Finish()`

Проставляет `Step.Stop` = `GetNow()`

### `func (st *Step) PrintAttachments()`

Проходится по всем `allure.Attachments` массива `Step.Attachments` и вызывает у `allure.Attachment` метод `Print()`.
