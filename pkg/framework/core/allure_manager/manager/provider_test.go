package manager

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/bytedance/sonic"
	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type testMetaMockProvider struct {
	result    *allure.Result
	container *allure.Container
	be        func(t provider.T)
	ae        func(t provider.T)
}

func (m *testMetaMockProvider) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockProvider) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockProvider) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockProvider) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockProvider) GetBeforeEach() func(t provider.T) {
	return m.be
}

func (m *testMetaMockProvider) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockProvider) GetAfterEach() func(t provider.T) {
	return m.ae
}

type suiteMetaMockProvider struct {
	namePrefix string
	name       string
	container  *allure.Container
	hook       func(t provider.T)
}

func (m *suiteMetaMockProvider) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockProvider) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockProvider) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockProvider) GetSuiteFullName() string {
	return fmt.Sprintf("%s/%s", m.namePrefix, m.name)
}

func (m *suiteMetaMockProvider) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockProvider) SetBeforeAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockProvider) SetAfterAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockProvider) GetBeforeAll() func(provider.T) {
	return m.hook
}

func (m *suiteMetaMockProvider) GetAfterAll() func(provider.T) {
	return m.hook
}

func (m *suiteMetaMockProvider) GetParentSuite() string {
	return ""
}

func TestNewProvider(t *testing.T) {
	cfg := NewProviderConfig().
		WithRunner("WithRunner").
		WithPackageName("WithPackageName").
		WithSuiteName("WithSuiteName").
		WithFullName("WithFullName")

	manager := NewProvider(cfg)
	require.NotNil(t, manager.GetTestMeta())
	require.NotNil(t, manager.GetSuiteMeta())
	require.Equal(t, "WithSuiteName", manager.GetSuiteMeta().GetSuiteName())
	require.Equal(t, "WithPackageName", manager.GetSuiteMeta().GetPackageName())
	require.Equal(t, "WithFullName", manager.GetSuiteMeta().GetSuiteFullName())
	require.Equal(t, "WithRunner", manager.GetSuiteMeta().GetRunner())
}

func TestAllureManager_safely(t *testing.T) {
	result := &allure.Result{}
	manager := allureManager{testMeta: &testMetaMockProvider{result: result}}
	manager.withResult(func(result2 *allure.Result) {
		require.NotNil(t, result2)
		require.Equal(t, result, result2)
	})
}

func TestAllureManager_UpdateResultStatus(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockProvider{result: &allure.Result{}}}
	manager.UpdateResultStatus("TestMsg", "TestTrace")
	require.Equal(t, "TestMsg", manager.GetResult().StatusDetails.Message)
	require.Equal(t, "TestTrace", manager.GetResult().StatusDetails.Trace)
}

func TestAllureManager_StopResult(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockProvider{result: &allure.Result{}}}
	manager.StopResult(allure.Unknown)
	now := allure.GetNow()
	require.Equal(t, allure.Unknown, manager.GetResult().Status)
	require.Equal(t, now, manager.GetResult().Stop)
}

func TestAllureManager_GetResult(t *testing.T) {
	result := &allure.Result{}
	manager := allureManager{testMeta: &testMetaMockProvider{result: result}}
	require.Equal(t, result, manager.GetResult())
}

func TestAllureManager_ExecutionContext(t *testing.T) {
	manager := allureManager{testMeta: &testMetaMockProvider{result: &allure.Result{}}}
	manager.TestContext()
	require.NotNil(t, manager.ExecutionContext())
}

func TestAllureManager_GetTestMeta(t *testing.T) {
	testMeta := &testMetaMockProvider{}
	manager := allureManager{testMeta: testMeta}
	require.Equal(t, testMeta, manager.GetTestMeta())
}

func TestAllureManager_GetSuiteMeta(t *testing.T) {
	suiteMeta := &suiteMetaMockProvider{}
	manager := allureManager{suiteMeta: suiteMeta}
	require.Equal(t, suiteMeta, manager.GetSuiteMeta())
}

func TestAllureManager_FinishTest(t *testing.T) {
	const (
		attachmentText = `THIS IS A TEXT ATTACHMENT`
		testName       = "Test"
		testFullName   = "FullNameTest"
		allureDir      = "./allure-results"
	)

	result := allure.NewResult(testName, testFullName)
	result.Attachments = append(result.Attachments, allure.NewAttachment("Text Attachment if TestAttachment", allure.Text, []byte(attachmentText)))
	result.PrintAttachments()
	now := allure.GetNow()

	manager := allureManager{testMeta: &testMetaMockProvider{result: result}}

	manager.FinishTest()

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

	emptyResult := &allure.Result{}
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

func TestAllureManager_NewTest(t *testing.T) {
	container := allure.NewContainer()
	manager := allureManager{suiteMeta: &suiteMetaMockProvider{namePrefix: "prefix", name: "name", container: container}}

	manager.NewTest("TestName", "PackageName", "tag1", "tag2")
	require.Equal(t, "TestName", manager.testMeta.GetResult().Name)
	require.Equal(t, "prefix/name/TestName", manager.testMeta.GetResult().FullName)

	require.NotNil(t, manager.testMeta.GetResult().GetLabels(allure.Suite))
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Suite))
	require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Suite), 1)

	require.Equal(t, "name", manager.testMeta.GetResult().GetLabels(allure.Suite)[0].GetValue())
	require.NotEmpty(t, manager.testMeta.GetResult().GetLabels(allure.Package))
	require.Len(t, manager.testMeta.GetResult().GetLabels(allure.Package), 1)

	require.Equal(t, "PackageName", manager.testMeta.GetResult().GetLabels(allure.Package)[0].GetValue())

	require.NotNil(t, manager.testMeta.GetContainer())
	require.NotEmpty(t, manager.testMeta.GetContainer().Children)
	require.Len(t, manager.testMeta.GetContainer().Children, 1)
	require.Equal(t, manager.GetResult().UUID.String(), manager.testMeta.GetContainer().Children[0].String())
}
