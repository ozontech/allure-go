package allure

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testName     = "Test"
	testFullName = "FullNameTest"
)

func TestNewResult(t *testing.T) {

	result := NewResult(testName, testFullName)
	now := GetNow()

	require.NotNil(t, result)
	require.Equal(t, testName, result.Name)
	require.Equal(t, testFullName, result.FullName)
	require.NotEmpty(t, result.UUID)
	require.Equal(t, getMD5Hash(testFullName), result.TestCaseID)
	require.True(t, result.ToPrint)
	require.Equal(t, getMD5Hash(getMD5Hash(testFullName)), result.HistoryID)
	require.Len(t, result.Labels, 1)
	require.Equal(t, Language.ToString(), result.Labels[0].Name)
	require.Equal(t, runtime.Version(), result.Labels[0].Value)
	require.Equal(t, now, result.Start)
}

func TestResult_GetLabel(t *testing.T) {
	result := NewResult(testName, testFullName)
	require.NotNil(t, result)

	langLabel := result.GetLabels(Language)
	require.NotNil(t, langLabel)
	require.Len(t, langLabel, 1)
	require.Equal(t, Language.ToString(), langLabel[0].Name)
	require.Equal(t, runtime.Version(), langLabel[0].Value)
}

func TestResult_Begin(t *testing.T) {
	result := new(Result)
	result.Begin()
	now := GetNow()
	require.Equal(t, now, result.Start)
}

func TestResult_Finish(t *testing.T) {
	result := new(Result)
	result.Finish()
	now := GetNow()
	require.Equal(t, now, result.Stop)
}

func TestResult_SkipOnPrint(t *testing.T) {
	result := new(Result)
	result.ToPrint = true

	result.SkipOnPrint()
	require.False(t, result.ToPrint)
}

func TestResult_WithFrameWork(t *testing.T) {
	labelValue := "TestFrameWork"
	result := new(Result)
	result.WithFrameWork(labelValue)

	label := result.GetLabels(Framework)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Framework.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_WithHost(t *testing.T) {
	labelValue := "TestHost"
	result := new(Result)
	result.WithHost(labelValue)

	label := result.GetLabels(Host)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Host.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_WithLanguage(t *testing.T) {
	labelValue := "TestLanguage"
	result := new(Result)
	result.WithLanguage(labelValue)

	label := result.GetLabels(Language)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Language.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_WithPackage(t *testing.T) {
	labelValue := "TestPackage"
	result := new(Result)
	result.WithPackage(labelValue)

	label := result.GetLabels(Package)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Package.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_WithParentSuite(t *testing.T) {
	labelValue := "TestParentSuite"
	result := new(Result)
	result.WithParentSuite(labelValue)

	label := result.GetLabels(ParentSuite)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, ParentSuite.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_WithSuite(t *testing.T) {
	labelValue := "TestSuite"
	result := new(Result)
	result.WithSuite(labelValue)

	label := result.GetLabels(Suite)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Suite.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_WithSubSuites(t *testing.T) {
	labelValue1 := "TestSubSuite1"
	labelValue2 := "TestSubsuite2"
	result := new(Result)
	result.WithSubSuites(labelValue1, labelValue2)

	label := result.GetLabels(SubSuite)
	require.NotNil(t, label)
	require.Len(t, label, 2)
	require.Equal(t, SubSuite.ToString(), label[0].Name)
	require.Equal(t, labelValue1, label[0].Value)
	require.Equal(t, SubSuite.ToString(), label[1].Name)
	require.Equal(t, labelValue2, label[1].Value)
}

func TestResult_WithThread(t *testing.T) {
	labelValue := "TestThread"
	result := new(Result)
	result.WithThread(labelValue)

	label := result.GetLabels(Thread)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Thread.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].Value)
}

func TestResult_SetStatusMessage(t *testing.T) {
	statusMessage := "statusMessageTest"
	result := new(Result)
	statusDetails := StatusDetail{}
	result.StatusDetails = statusDetails
	result.SetStatusMessage(statusMessage)
	require.Equal(t, statusMessage, result.StatusDetails.Message)
}

func TestResult_SetStatusTraceTest(t *testing.T) {
	statusTrace := "statusTraceTest"
	result := new(Result)
	statusDetails := StatusDetail{}
	result.StatusDetails = statusDetails
	result.SetStatusTrace(statusTrace)
	require.Equal(t, statusTrace, result.StatusDetails.Trace)
}

func TestResult_GetStatusMessage(t *testing.T) {
	statusMessage := "statusMessageTest"
	result := new(Result)
	statusDetails := StatusDetail{Message: statusMessage}
	result.StatusDetails = statusDetails
	require.Equal(t, statusMessage, result.GetStatusMessage())
}

func TestResult_GetStatusTrace(t *testing.T) {
	statusTrace := "statusTraceTest"
	result := new(Result)
	statusDetails := StatusDetail{Trace: statusTrace}
	result.StatusDetails = statusDetails
	require.Equal(t, statusTrace, result.GetStatusTrace())
}

func TestResult_SetLabel(t *testing.T) {
	labelValue1 := "TestValue1"
	labelValue2 := "TestValue2"
	result := new(Result)
	result.AddLabel(FrameWorkLabel(labelValue1), HostLabel(labelValue2))

	label := result.Labels
	require.NotNil(t, label)
	require.Len(t, label, 2)
	require.Equal(t, Framework.ToString(), label[0].Name)
	require.Equal(t, labelValue1, label[0].Value)
	require.Equal(t, Host.ToString(), label[1].Name)
	require.Equal(t, labelValue2, label[1].Value)
}

func TestResult_SetNewLabelMap(t *testing.T) {
	labelValue1 := "TestValue1"
	labelValue2 := "TestValue2"

	labelMap := map[LabelType]string{
		Framework: labelValue1,
		Host:      labelValue2,
	}
	result := new(Result)
	result.SetNewLabelMap(labelMap)
	labelFramework := result.GetLabels(Framework)
	require.NotNil(t, labelFramework)
	require.Len(t, labelFramework, 1)
	require.Equal(t, Framework.ToString(), labelFramework[0].Name)
	require.Equal(t, labelValue1, labelFramework[0].Value)

	labelHost := result.GetLabels(Host)
	require.NotNil(t, labelFramework)
	require.Len(t, labelFramework, 1)
	require.Equal(t, Host.ToString(), labelHost[0].Name)
	require.Equal(t, labelValue2, labelHost[0].Value)
}

func TestResult_WithLaunchTags_noTags(t *testing.T) {
	result := new(Result)
	result.WithLaunchTags()
	require.Nil(t, result.Labels)
}

func TestResult_WithLaunchTags_withTags(t *testing.T) {
	result := new(Result)
	os.Setenv(defaultTagsEnvKey, "tag1,tag2")
	defer os.Setenv(defaultTagsEnvKey, "")
	result.WithLaunchTags()
	require.NotNil(t, result.Labels)

	labels := result.GetLabels(Tag)
	require.Len(t, labels, 2)
	require.Equal(t, Tag.ToString(), labels[0].Name)
	require.Equal(t, "tag1", labels[0].Value)
	require.Equal(t, Tag.ToString(), labels[1].Name)
	require.Equal(t, "tag2", labels[1].Value)
}

func TestResult_WithLabels(t *testing.T) {
	labelValue1 := "TestValue1"
	labelValue2 := "TestValue2"

	result := new(Result)
	result.WithLabels(FrameWorkLabel(labelValue1), LanguageLabel(labelValue2))

	require.NotNil(t, result.Labels)

	labels := result.Labels
	require.Len(t, labels, 2)
	require.Equal(t, Framework.ToString(), labels[0].Name)
	require.Equal(t, labelValue1, labels[0].Value)
	require.Equal(t, Language.ToString(), labels[1].Name)
	require.Equal(t, labelValue2, labels[1].Value)
}

func TestResult_PrintAttachments(t *testing.T) {
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	result := new(Result)
	result.Attachments = append(result.Attachments, NewAttachment("Text Attachment if TestAttachment", Text, []byte(attachmentText)))
	result.PrintAttachments()
	defer os.RemoveAll(allureDir)

	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 1)
	var attachFile *os.File
	defer attachFile.Close()

	f := files[0]
	attachFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := ioutil.ReadAll(attachFile)
	require.NoError(t, readErr)
	require.Equal(t, attachmentText, string(bytes))
}

func TestResult_Print_toPrintFalse(t *testing.T) {
	result := new(Result)
	err := result.Print()

	require.NoError(t, err)
	require.NoDirExists(t, allureDir)
}

func TestResult_Print(t *testing.T) {
	result := NewResult(testName, testFullName)
	err := result.Print()

	require.NoError(t, err)

	defer os.RemoveAll(allureDir)
	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 1)

	var resultFile *os.File
	defer resultFile.Close()

	f := files[0]
	emptyResult := &Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := ioutil.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := json.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)

	require.Equal(t, result.Name, emptyResult.Name)
	require.Equal(t, result.FullName, emptyResult.FullName)
	require.Equal(t, result.UUID, emptyResult.UUID)
	require.Equal(t, result.TestCaseID, emptyResult.TestCaseID)
	require.Equal(t, result.HistoryID, emptyResult.HistoryID)

	require.Len(t, emptyResult.Labels, 1)
	require.Equal(t, result.Labels[0].Name, emptyResult.Labels[0].Name)
	require.Equal(t, result.Labels[0].Value, emptyResult.Labels[0].Value)
	require.Equal(t, result.Start, emptyResult.Start)
}

func TestResult_Print_withAttachment(t *testing.T) {
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	result := NewResult(testName, testFullName)
	result.Attachments = append(result.Attachments, NewAttachment("Text Attachment if TestAttachment", Text, []byte(attachmentText)))
	result.PrintAttachments()
	err := result.Print()

	require.NoError(t, err)

	defer os.RemoveAll(allureDir)
	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 2)

	var resultFile *os.File
	defer resultFile.Close()

	var (
		fileByte   os.FileInfo
		attachByte os.FileInfo
	)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-result.json") {
			fileByte = f
			continue
		}

		if strings.HasSuffix(f.Name(), "-attachment.txt") {
			attachByte = f
			continue
		}
	}

	emptyResult := &Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, fileByte.Name()))
	bytes, readErr := ioutil.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := json.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)

	require.Equal(t, result.Name, emptyResult.Name)
	require.Equal(t, result.FullName, emptyResult.FullName)
	require.Equal(t, result.UUID, emptyResult.UUID)
	require.Equal(t, result.TestCaseID, emptyResult.TestCaseID)
	require.Equal(t, result.HistoryID, emptyResult.HistoryID)

	require.Len(t, emptyResult.Labels, 1)
	require.Equal(t, result.Labels[0].Name, emptyResult.Labels[0].Name)
	require.Equal(t, result.Labels[0].Value, emptyResult.Labels[0].Value)
	require.Equal(t, result.Start, emptyResult.Start)

	attachFile, _ := os.Open(fmt.Sprintf("%s/%s", allureDir, attachByte.Name()))
	bytes, readErr = ioutil.ReadAll(attachFile)
	require.NoError(t, readErr)
	require.Equal(t, attachmentText, string(bytes))
}

func TestResult_Done(t *testing.T) {
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	result := NewResult(testName, testFullName)
	result.Attachments = append(result.Attachments, NewAttachment("Text Attachment if TestAttachment", Text, []byte(attachmentText)))
	result.PrintAttachments()
	now := GetNow()
	result.Done()

	defer os.RemoveAll(allureDir)
	files, _ := ioutil.ReadDir(allureDir)
	require.Len(t, files, 2)

	var resultFile *os.File
	defer resultFile.Close()

	var (
		fileByte   os.FileInfo
		attachByte os.FileInfo
	)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-result.json") {
			fileByte = f
			continue
		}

		if strings.HasSuffix(f.Name(), "-attachment.txt") {
			attachByte = f
			continue
		}
	}

	emptyResult := &Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, fileByte.Name()))
	bytes, readErr := ioutil.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := json.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)

	require.Equal(t, result.Name, emptyResult.Name)
	require.Equal(t, result.FullName, emptyResult.FullName)
	require.Equal(t, result.UUID, emptyResult.UUID)
	require.Equal(t, result.TestCaseID, emptyResult.TestCaseID)
	require.Equal(t, result.HistoryID, emptyResult.HistoryID)

	require.Len(t, emptyResult.Labels, 1)
	require.Equal(t, result.Labels[0].Name, emptyResult.Labels[0].Name)
	require.Equal(t, result.Labels[0].Value, emptyResult.Labels[0].Value)
	require.Equal(t, result.Start, emptyResult.Start)
	require.Equal(t, now, emptyResult.Stop)

	attachFile, _ := os.Open(fmt.Sprintf("%s/%s", allureDir, attachByte.Name()))
	bytes, readErr = ioutil.ReadAll(attachFile)
	require.NoError(t, readErr)
	require.Equal(t, attachmentText, string(bytes))
}

func TestUnmarshall(t *testing.T) {
	input :=
		`{
	"uuid": "ac2bdad0-9f7b-4b38-a64b-e7210b74c912",
	"historyId": "Общие настройки перед выполением тестовых сценариев.(stg_zup_ir_daily) _02 Настройка обновления для новостей",
	"name": "(stg_zup_ir_daily) _02 Настройка обновления для новостей",
	"fullName": null,
	"start": 1659265291384,
	"stop": 1659265305014,
	"statusDetails": {
		"known": false,
		"muted": false,
		"flaky": false
	},
	"status": "passed",
	"stage": "finished",
	"steps": [
		{
			"name": "И Настройка событий и новостей для пользователей",
			"start": 1659265292607,
			"stop": 1659265305014,
			"status": "passed",
			"stage": "finished",
			"statusDetails": {
				"known": false,
				"muted": false,
				"flaky": false
			},
			"parameters": [],
			"attachments": [],
			"steps": [
				{
					"name": "И я открываю окно функции для технического специалиста (расширение)",
					"start": 1659265292607,
					"stop": 1659265292607,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [],
					"attachments": [],
					"steps": []
				},
				{
					"name": "Когда открылось окно 'Функции для технического специалиста'",
					"start": 1659265292607,
					"stop": 1659265298859,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Функции для технического специалиста"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"Table\" я перехожу к строке:",
					"start": 1659265298859,
					"stop": 1659265298872,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Table"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "813798d0-5df4-4b2b-af41-ac5fb5e2cb31-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"Table\" я разворачиваю строку:",
					"start": 1659265298872,
					"stop": 1659265298887,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Table"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "70a5546f-684c-4e92-8fd3-3d3f0e3f46b6-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"Table\" я перехожу к строке:",
					"start": 1659265298887,
					"stop": 1659265299028,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Table"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "34d7f7ef-4f00-473b-8f2a-9088535c04f3-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"Table\" я нажимаю на кнопку с именем 'CommandOpen'",
					"start": 1659265299028,
					"stop": 1659265299424,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Table"
						},
						{
							"name": "Парам2",
							"value": "CommandOpen"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "Тогда открылось окно 'Управление настройками и обновлением для новостей'",
					"start": 1659265299424,
					"stop": 1659265301245,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Управление настройками и обновлением для новостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я перехожу к закладке с именем \"СтраницаНастройкаПользователей\"",
					"start": 1659265301261,
					"stop": 1659265301261,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "СтраницаНастройкаПользователей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я перехожу к строке:",
					"start": 1659265301261,
					"stop": 1659265301520,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "17e78d03-44d1-43ca-9679-8a14a02b4a9b-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я нажимаю на кнопку с именем 'НастройкиОтборовПользователямиКомандаНастройкиПользователя'",
					"start": 1659265301520,
					"stop": 1659265301567,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						},
						{
							"name": "Парам2",
							"value": "НастройкиОтборовПользователямиКомандаНастройкиПользователя"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я перехожу к закладке с именем \"СтраницаЛентыНовостей\"",
					"start": 1659265301568,
					"stop": 1659265301838,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "СтраницаЛентыНовостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"ЛентыНовостей\" я нажимаю на кнопку с именем 'ЛентыНовостейКомандаВыключитьВсе'",
					"start": 1659265301838,
					"stop": 1659265301890,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "ЛентыНовостей"
						},
						{
							"name": "Парам2",
							"value": "ЛентыНовостейКомандаВыключитьВсе"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я нажимаю на кнопку с именем 'КнопкаОК'",
					"start": 1659265301890,
					"stop": 1659265301908,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "КнопкаОК"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я перехожу к строке:",
					"start": 1659265301909,
					"stop": 1659265302655,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "86b3da76-2483-4fdc-9c4f-93a3f2776ddf-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я нажимаю на кнопку с именем 'НастройкиОтборовПользователямиКомандаНастройкиПользователя'",
					"start": 1659265302655,
					"stop": 1659265302688,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						},
						{
							"name": "Парам2",
							"value": "НастройкиОтборовПользователямиКомандаНастройкиПользователя"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я перехожу к закладке с именем \"СтраницаЛентыНовостей\"",
					"start": 1659265302688,
					"stop": 1659265302841,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "СтраницаЛентыНовостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"ЛентыНовостей\" я нажимаю на кнопку с именем 'ЛентыНовостейКомандаВыключитьВсе'",
					"start": 1659265302841,
					"stop": 1659265302875,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "ЛентыНовостей"
						},
						{
							"name": "Парам2",
							"value": "ЛентыНовостейКомандаВыключитьВсе"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я нажимаю на кнопку с именем 'КнопкаОК'",
					"start": 1659265302876,
					"stop": 1659265302893,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "КнопкаОК"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я перехожу к строке:",
					"start": 1659265302894,
					"stop": 1659265303070,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "54d3126c-57d8-41a1-b424-5a53805bf8b5-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я нажимаю на кнопку с именем 'НастройкиОтборовПользователямиКомандаНастройкиПользователя'",
					"start": 1659265303071,
					"stop": 1659265303105,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						},
						{
							"name": "Парам2",
							"value": "НастройкиОтборовПользователямиКомандаНастройкиПользователя"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я перехожу к закладке с именем \"СтраницаЛентыНовостей\"",
					"start": 1659265303106,
					"stop": 1659265303249,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "СтраницаЛентыНовостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"ЛентыНовостей\" я нажимаю на кнопку с именем 'ЛентыНовостейКомандаВыключитьВсе'",
					"start": 1659265303249,
					"stop": 1659265303281,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "ЛентыНовостей"
						},
						{
							"name": "Парам2",
							"value": "ЛентыНовостейКомандаВыключитьВсе"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я нажимаю на кнопку с именем 'КнопкаОК'",
					"start": 1659265303281,
					"stop": 1659265303296,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "КнопкаОК"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я перехожу к строке:",
					"start": 1659265303296,
					"stop": 1659265303631,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "e99c6bd7-ee8c-4512-89ae-6ad6cfd8bbaf-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я нажимаю на кнопку с именем 'НастройкиОтборовПользователямиКомандаНастройкиПользователя'",
					"start": 1659265303631,
					"stop": 1659265303662,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						},
						{
							"name": "Парам2",
							"value": "НастройкиОтборовПользователямиКомандаНастройкиПользователя"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я перехожу к закладке с именем \"СтраницаЛентыНовостей\"",
					"start": 1659265303678,
					"stop": 1659265303807,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "СтраницаЛентыНовостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"ЛентыНовостей\" я нажимаю на кнопку с именем 'ЛентыНовостейКомандаВыключитьВсе'",
					"start": 1659265303807,
					"stop": 1659265303838,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "ЛентыНовостей"
						},
						{
							"name": "Парам2",
							"value": "ЛентыНовостейКомандаВыключитьВсе"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я нажимаю на кнопку с именем 'КнопкаОК'",
					"start": 1659265303838,
					"stop": 1659265303854,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "КнопкаОК"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я перехожу к строке:",
					"start": 1659265303854,
					"stop": 1659265304438,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						}
					],
					"attachments": [
						{
							"name": "table",
							"source": "d3225e2a-e901-4525-bc8b-421a48fb4fff-attachment.csv",
							"type": "text/csv"
						}
					],
					"steps": []
				},
				{
					"name": "И в таблице \"НастройкиОтборовПользователями\" я нажимаю на кнопку с именем 'НастройкиОтборовПользователямиКомандаНастройкиПользователя'",
					"start": 1659265304438,
					"stop": 1659265304470,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "НастройкиОтборовПользователями"
						},
						{
							"name": "Парам2",
							"value": "НастройкиОтборовПользователямиКомандаНастройкиПользователя"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я перехожу к закладке с именем \"СтраницаЛентыНовостей\"",
					"start": 1659265304470,
					"stop": 1659265304620,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "СтраницаЛентыНовостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И в таблице \"ЛентыНовостей\" я нажимаю на кнопку с именем 'ЛентыНовостейКомандаВыключитьВсе'",
					"start": 1659265304620,
					"stop": 1659265304651,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "ЛентыНовостей"
						},
						{
							"name": "Парам2",
							"value": "ЛентыНовостейКомандаВыключитьВсе"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я нажимаю на кнопку с именем 'КнопкаОК'",
					"start": 1659265304651,
					"stop": 1659265304667,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "КнопкаОК"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "Тогда открылось окно 'Управление настройками и обновлением для новостей'",
					"start": 1659265304667,
					"stop": 1659265304746,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Управление настройками и обновлением для новостей"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я нажимаю на кнопку с именем 'ФормаЗакрыть'",
					"start": 1659265304746,
					"stop": 1659265304762,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "ФормаЗакрыть"
						}
					],
					"attachments": [],
					"steps": []
				},
				{
					"name": "И я жду закрытия окна 'Управление настройками и обновлением для новостей' в течение 20 секунд",
					"start": 1659265304762,
					"stop": 1659265305014,
					"status": "passed",
					"stage": "finished",
					"statusDetails": {
						"known": false,
						"muted": false,
						"flaky": false
					},
					"parameters": [
						{
							"name": "Парам1",
							"value": "Управление настройками и обновлением для новостей"
						},
						{
							"name": "Парам2",
							"value": 20
						}
					],
					"attachments": [],
					"steps": []
				}
			]
		}
	],
	"parameters": [],
	"labels": [
		{
			"name": "story",
			"value": "_02 Настройка обновления для новостей"
		},
		{
			"name": "feature",
			"value": "Общие настройки перед выполением тестовых сценариев"
		},
		{
			"name": "epic",
			"value": "Internal"
		},
		{
			"name": "suite",
			"value": "Общие настройки перед выполением тестовых сценариев"
		},
		{
			"name": "tag",
			"value": "settings"
		},
		{
			"name": "tag",
			"value": "tree"
		},
		{
			"name": "package",
			"value": "0_Settings"
		},
		{
			"name": "host",
			"value": "accapp10z501"
		}
	],
	"links": [],
	"attachments": [],
	"description": null
}`
	var result Result
	err := json.Unmarshal([]byte(input), &result)
	require.NoError(t, err)
}
