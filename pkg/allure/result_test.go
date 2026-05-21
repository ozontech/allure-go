package allure

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
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
	require.Equal(t, runtime.Version(), result.Labels[0].GetValue())
	require.Equal(t, now, result.Start)
}

func TestResult_GetLabel(t *testing.T) {
	result := NewResult(testName, testFullName)
	require.NotNil(t, result)

	langLabel := result.GetLabels(Language)
	require.NotNil(t, langLabel)
	require.Len(t, langLabel, 1)
	require.Equal(t, Language.ToString(), langLabel[0].Name)
	require.Equal(t, runtime.Version(), langLabel[0].GetValue())
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
	require.Equal(t, labelValue, label[0].GetValue())
}

func TestResult_WithHost(t *testing.T) {
	labelValue := "TestHost"
	result := new(Result)
	result.WithHost(labelValue)

	label := result.GetLabels(Host)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Host.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].GetValue())
}

func TestResult_WithLanguage(t *testing.T) {
	labelValue := "TestLanguage"
	result := new(Result)
	result.WithLanguage(labelValue)

	label := result.GetLabels(Language)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Language.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].GetValue())
}

func TestResult_WithPackage(t *testing.T) {
	labelValue := "TestPackage"
	result := new(Result)
	result.WithPackage(labelValue)

	label := result.GetLabels(Package)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Package.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].GetValue())
}

func TestResult_WithParentSuite(t *testing.T) {
	labelValue := "TestParentSuite"
	result := new(Result)
	result.WithParentSuite(labelValue)

	label := result.GetLabels(ParentSuite)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, ParentSuite.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].GetValue())
}

func TestResult_WithSuite(t *testing.T) {
	labelValue := "TestSuite"
	result := new(Result)
	result.WithSuite(labelValue)

	label := result.GetLabels(Suite)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Suite.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].GetValue())
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
	require.Equal(t, labelValue1, label[0].GetValue())
	require.Equal(t, SubSuite.ToString(), label[1].Name)
	require.Equal(t, labelValue2, label[1].GetValue())
}

func TestResult_WithThread(t *testing.T) {
	labelValue := "TestThread"
	result := new(Result)
	result.WithThread(labelValue)

	label := result.GetLabels(Thread)
	require.NotNil(t, label)
	require.Len(t, label, 1)
	require.Equal(t, Thread.ToString(), label[0].Name)
	require.Equal(t, labelValue, label[0].GetValue())
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
	require.Equal(t, labelValue1, label[0].GetValue())
	require.Equal(t, Host.ToString(), label[1].Name)
	require.Equal(t, labelValue2, label[1].GetValue())
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
	require.Equal(t, labelValue1, labelFramework[0].GetValue())

	labelHost := result.GetLabels(Host)
	require.NotNil(t, labelFramework)
	require.Len(t, labelFramework, 1)
	require.Equal(t, Host.ToString(), labelHost[0].Name)
	require.Equal(t, labelValue2, labelHost[0].GetValue())
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
	require.Equal(t, "tag1", labels[0].GetValue())
	require.Equal(t, Tag.ToString(), labels[1].Name)
	require.Equal(t, "tag2", labels[1].GetValue())
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
	require.Equal(t, labelValue1, labels[0].GetValue())
	require.Equal(t, Language.ToString(), labels[1].Name)
	require.Equal(t, labelValue2, labels[1].GetValue())
}

func TestResult_PrintAttachments(t *testing.T) {
	attachmentText := `THIS IS A TEXT ATTACHMENT`
	result := new(Result)
	result.Attachments = append(result.Attachments, NewAttachment("Text Attachment if TestAttachment", Text, []byte(attachmentText)))
	result.PrintAttachments()
	defer os.RemoveAll(allureDir)

	files, _ := os.ReadDir(allureDir)
	require.Len(t, files, 1)
	var attachFile *os.File
	defer attachFile.Close()

	f := files[0]
	attachFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := io.ReadAll(attachFile)
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
	files, _ := os.ReadDir(allureDir)
	require.Len(t, files, 1)

	var resultFile *os.File
	defer resultFile.Close()

	f := files[0]
	emptyResult := &Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, f.Name()))
	bytes, readErr := io.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := sonic.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)

	require.Equal(t, result.Name, emptyResult.Name)
	require.Equal(t, result.FullName, emptyResult.FullName)
	require.Equal(t, result.UUID, emptyResult.UUID)
	require.Equal(t, result.TestCaseID, emptyResult.TestCaseID)
	require.Equal(t, result.HistoryID, emptyResult.HistoryID)

	require.Len(t, emptyResult.Labels, 1)
	require.Equal(t, result.Labels[0].Name, emptyResult.Labels[0].Name)
	require.Equal(t, result.Labels[0].GetValue(), emptyResult.Labels[0].GetValue())
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
	files, _ := os.ReadDir(allureDir)
	require.Len(t, files, 2)

	var resultFile *os.File
	defer resultFile.Close()

	var (
		fileByte   os.FileInfo
		attachByte os.FileInfo
	)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-result.json") {
			info, err := f.Info()
			require.NoError(t, err)

			fileByte = info
			continue
		}

		if strings.HasSuffix(f.Name(), "-attachment.txt") {
			info, err := f.Info()
			require.NoError(t, err)

			attachByte = info
			continue
		}
	}

	emptyResult := &Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, fileByte.Name()))
	bytes, readErr := io.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := sonic.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)

	require.Equal(t, result.Name, emptyResult.Name)
	require.Equal(t, result.FullName, emptyResult.FullName)
	require.Equal(t, result.UUID, emptyResult.UUID)
	require.Equal(t, result.TestCaseID, emptyResult.TestCaseID)
	require.Equal(t, result.HistoryID, emptyResult.HistoryID)

	require.Len(t, emptyResult.Labels, 1)
	require.Equal(t, result.Labels[0].Name, emptyResult.Labels[0].Name)
	require.Equal(t, result.Labels[0].GetValue(), emptyResult.Labels[0].GetValue())
	require.Equal(t, result.Start, emptyResult.Start)

	attachFile, _ := os.Open(fmt.Sprintf("%s/%s", allureDir, attachByte.Name()))
	bytes, readErr = io.ReadAll(attachFile)
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
	files, _ := os.ReadDir(allureDir)
	require.Len(t, files, 2)

	var resultFile *os.File
	defer resultFile.Close()

	var (
		fileByte   os.FileInfo
		attachByte os.FileInfo
	)

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-result.json") {
			info, err := f.Info()
			require.NoError(t, err)

			fileByte = info
			continue
		}

		if strings.HasSuffix(f.Name(), "-attachment.txt") {
			info, err := f.Info()
			require.NoError(t, err)

			attachByte = info
			continue
		}
	}

	emptyResult := &Result{}
	resultFile, _ = os.Open(fmt.Sprintf("%s/%s", allureDir, fileByte.Name()))
	bytes, readErr := io.ReadAll(resultFile)
	require.NoError(t, readErr)
	unMarshallErr := sonic.Unmarshal(bytes, emptyResult)
	require.NoError(t, unMarshallErr)

	require.Equal(t, result.Name, emptyResult.Name)
	require.Equal(t, result.FullName, emptyResult.FullName)
	require.Equal(t, result.UUID, emptyResult.UUID)
	require.Equal(t, result.TestCaseID, emptyResult.TestCaseID)
	require.Equal(t, result.HistoryID, emptyResult.HistoryID)

	require.Len(t, emptyResult.Labels, 1)
	require.Equal(t, result.Labels[0].Name, emptyResult.Labels[0].Name)
	require.Equal(t, result.Labels[0].GetValue(), emptyResult.Labels[0].GetValue())
	require.Equal(t, result.Start, emptyResult.Start)
	require.Equal(t, now, emptyResult.Stop)

	attachFile, _ := os.Open(fmt.Sprintf("%s/%s", allureDir, attachByte.Name()))
	bytes, readErr = io.ReadAll(attachFile)
	require.NoError(t, readErr)
	require.Equal(t, attachmentText, string(bytes))
}
