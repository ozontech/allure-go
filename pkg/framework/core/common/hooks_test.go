package common

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/provider"
)

type suiteMetaMockHooks struct {
	namePrefix string
	name       string
	container  *allure.Container

	baFlag bool
	aaFlag bool
	hook   func(t provider.T)
}

func (m *suiteMetaMockHooks) GetPackageName() string {
	return m.name
}

func (m *suiteMetaMockHooks) GetRunner() string {
	return m.name
}

func (m *suiteMetaMockHooks) GetSuiteName() string {
	return m.name
}

func (m *suiteMetaMockHooks) GetSuiteFullName() string {
	return fmt.Sprintf("%s/%s", m.namePrefix, m.name)
}

func (m *suiteMetaMockHooks) GetContainer() *allure.Container {
	return m.container
}

func (m *suiteMetaMockHooks) SetBeforeAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockHooks) SetAfterAll(hook func(provider.T)) {
	m.hook = hook
}

func (m *suiteMetaMockHooks) GetBeforeAll() func(provider.T) {
	m.baFlag = true
	return m.hook
}

func (m *suiteMetaMockHooks) GetAfterAll() func(provider.T) {
	m.aaFlag = true
	return m.hook
}

type testMetaMockHooks struct {
	result    *allure.Result
	container *allure.Container

	beFlag bool
	be     func(t provider.T)
	aeFlag bool
	ae     func(t provider.T)
}

func (m *testMetaMockHooks) GetResult() *allure.Result {
	return m.result
}

func (m *testMetaMockHooks) SetResult(result *allure.Result) {
	m.result = result
}

func (m *testMetaMockHooks) GetContainer() *allure.Container {
	return m.container
}

func (m *testMetaMockHooks) SetBeforeEach(hook func(t provider.T)) {
	m.be = hook
}

func (m *testMetaMockHooks) GetBeforeEach() func(t provider.T) {
	m.beFlag = true
	return m.be
}

func (m *testMetaMockHooks) SetAfterEach(hook func(t provider.T)) {
	m.ae = hook
}

func (m *testMetaMockHooks) GetAfterEach() func(t provider.T) {
	m.aeFlag = true
	return m.ae
}

type realTMock struct {
	testing.TB
	parallelFlag bool
	runTest      bool
	testName     string
}

func (m *realTMock) Parallel() {
	m.parallelFlag = true
}

func (m *realTMock) Run(testName string, testBody func(t *testing.T)) bool {
	m.testName = testName
	m.runTest = true
	return true
}

type hookTMock struct {
	provider.T

	realT   *realTMock
	errorF  bool
	realTF  bool
	failNow bool
	wgFlag  bool
	wg      *sync.WaitGroup
}

func (m *hookTMock) WG() *sync.WaitGroup {
	m.wgFlag = true
	return m.wg
}

func (m *hookTMock) FailNow() {
	m.failNow = true
}

func (m *hookTMock) Errorf(format string, args ...interface{}) {
	m.errorF = true
}

func (m *hookTMock) SetRealT(t provider.TestingT) {
	m.realTF = true
}

func (m *hookTMock) RealT() provider.TestingT {
	return m.realT
}

type hookProviderMock struct {
	beforeAll  bool
	beforeEach bool
	afterEach  bool
	afterAll   bool

	suiteMeta *suiteMetaMockHooks
	testMeta  *testMetaMockHooks
}

func (m *hookProviderMock) BeforeEachContext() {
	m.beforeEach = true
}

func (m *hookProviderMock) AfterEachContext() {
	m.afterEach = true
}

func (m *hookProviderMock) BeforeAllContext() {
	m.beforeAll = true
}

func (m *hookProviderMock) AfterAllContext() {
	m.afterAll = true
}

func (m *hookProviderMock) GetSuiteMeta() provider.SuiteMeta {
	return m.suiteMeta
}

func (m *hookProviderMock) GetTestMeta() provider.TestMeta {
	return m.testMeta
}

func TestBeforeAllHook(t *testing.T) {
	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
	hookBody := func(t provider.T) {}
	providerMock := &hookProviderMock{
		suiteMeta: &suiteMetaMockHooks{hook: hookBody},
		testMeta:  &testMetaMockHooks{},
	}

	hookFunc := CarriedHook(BeforeAll, providerMock.GetSuiteMeta().GetBeforeAll)
	hookFunc(tMock, providerMock, &sync.WaitGroup{})

	require.True(t, tMock.wgFlag)
	require.True(t, tMock.realTF)

	require.True(t, tMock.realT.runTest)
	require.False(t, tMock.realT.parallelFlag)
	require.Equal(t, string(BeforeAll), tMock.realT.testName)

	require.True(t, providerMock.suiteMeta.baFlag)
	require.False(t, providerMock.suiteMeta.aaFlag)

	require.False(t, providerMock.testMeta.beFlag)
	require.False(t, providerMock.testMeta.aeFlag)
}

func TestBeforeEachHook(t *testing.T) {
	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
	providerMock := &hookProviderMock{
		suiteMeta: &suiteMetaMockHooks{},
		testMeta:  &testMetaMockHooks{be: func(t provider.T) {}},
	}

	hookFunc := CarriedHook(BeforeEach, providerMock.GetTestMeta().GetBeforeEach)
	hookFunc(tMock, providerMock, &sync.WaitGroup{})

	require.True(t, tMock.wgFlag)
	require.True(t, tMock.realTF)

	require.True(t, tMock.realT.runTest)
	require.Equal(t, string(BeforeEach), tMock.realT.testName)

	require.False(t, providerMock.suiteMeta.baFlag)
	require.False(t, providerMock.suiteMeta.aaFlag)

	require.True(t, providerMock.testMeta.beFlag)
	require.False(t, providerMock.testMeta.aeFlag)
}

func TestAfterAllHook(t *testing.T) {
	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
	providerMock := &hookProviderMock{
		suiteMeta: &suiteMetaMockHooks{hook: func(t provider.T) {}},
		testMeta:  &testMetaMockHooks{},
	}

	hookFunc := CarriedHook(AfterAll, providerMock.GetSuiteMeta().GetAfterAll)
	hookFunc(tMock, providerMock, &sync.WaitGroup{})

	require.True(t, tMock.wgFlag)
	require.True(t, tMock.realTF)

	require.True(t, tMock.realT.runTest)
	require.Equal(t, string(AfterAll), tMock.realT.testName)

	require.False(t, providerMock.suiteMeta.baFlag)
	require.True(t, providerMock.suiteMeta.aaFlag)

	require.False(t, providerMock.testMeta.beFlag)
	require.False(t, providerMock.testMeta.aeFlag)
}

func TestAfterEachHook(t *testing.T) {
	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
	providerMock := &hookProviderMock{
		suiteMeta: &suiteMetaMockHooks{},
		testMeta:  &testMetaMockHooks{ae: func(t provider.T) {}},
	}

	hookFunc := CarriedHook(AfterEach, providerMock.GetTestMeta().GetAfterEach)
	hookFunc(tMock, providerMock, &sync.WaitGroup{})

	require.True(t, tMock.wgFlag)
	require.True(t, tMock.realTF)

	require.True(t, tMock.realT.runTest)
	require.Equal(t, string(AfterEach), tMock.realT.testName)

	require.False(t, providerMock.suiteMeta.baFlag)
	require.False(t, providerMock.suiteMeta.aaFlag)

	require.False(t, providerMock.testMeta.beFlag)
	require.True(t, providerMock.testMeta.aeFlag)
}

//
//func TestBeforeAllHook_panic(t *testing.T) {
//	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
//	providerMock := &hookProviderMock{
//		suiteMeta: &suiteMetaMockHooks{hook: func(t provider.T) { panic("whoops") }},
//		testMeta:  &testMetaMockHooks{},
//	}
//
//	hookFunc := CarriedHook(BeforeAll, providerMock.GetSuiteMeta().GetBeforeAll)
//	hookFunc(tMock, providerMock, &sync.WaitGroup{})
//
//	require.True(t, tMock.wgFlag)
//	require.True(t, tMock.failNow)
//	require.True(t, tMock.errorF)
//	require.True(t, tMock.realTF)
//
//	require.True(t, tMock.realT.runTest)
//	require.False(t, tMock.realT.parallelFlag)
//	require.Equal(t, string(BeforeAll), tMock.realT.testName)
//
//	require.True(t, providerMock.suiteMeta.baFlag)
//	require.False(t, providerMock.suiteMeta.aaFlag)
//
//	require.False(t, providerMock.testMeta.beFlag)
//	require.False(t, providerMock.testMeta.aeFlag)
//}
//
//func TestBeforeEachHook_panic(t *testing.T) {
//	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
//	providerMock := &hookProviderMock{
//		suiteMeta: &suiteMetaMockHooks{hook: func(t provider.T) {}},
//		testMeta:  &testMetaMockHooks{},
//	}
//
//	hookFunc := CarriedHook(BeforeEach, providerMock.GetTestMeta().GetBeforeEach)
//	hookFunc(tMock, providerMock, &sync.WaitGroup{})
//
//	require.False(t, tMock.wgFlag)
//	require.True(t, tMock.failNow)
//	require.True(t, tMock.errorF)
//	require.True(t, tMock.realTF)
//
//	require.True(t, tMock.realT.runTest)
//	require.False(t, tMock.realT.parallelFlag)
//	require.Equal(t, string(BeforeEach), tMock.realT.testName)
//
//	require.False(t, providerMock.suiteMeta.baFlag)
//	require.False(t, providerMock.suiteMeta.aaFlag)
//
//	require.True(t, providerMock.testMeta.beFlag)
//	require.False(t, providerMock.testMeta.aeFlag)
//}
//
//func TestAfterAllHook_panic(t *testing.T) {
//	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
//	providerMock := &hookProviderMock{
//		suiteMeta: &suiteMetaMockHooks{hook: func(t provider.T) {}},
//		testMeta:  &testMetaMockHooks{},
//	}
//
//	hookFunc := CarriedHook(AfterAll, providerMock.GetSuiteMeta().GetAfterAll)
//	hookFunc(tMock, providerMock, &sync.WaitGroup{})
//
//	require.True(t, tMock.wgFlag)
//	require.True(t, tMock.failNow)
//	require.True(t, tMock.errorF)
//	require.True(t, tMock.realTF)
//
//	require.True(t, tMock.realT.runTest)
//	require.True(t, tMock.realT.parallelFlag)
//	require.Equal(t, string(AfterAll), tMock.realT.testName)
//
//	require.False(t, providerMock.suiteMeta.baFlag)
//	require.True(t, providerMock.suiteMeta.aaFlag)
//
//	require.False(t, providerMock.testMeta.beFlag)
//	require.False(t, providerMock.testMeta.aeFlag)
//}
//
//func TestAfterEachHook_panic(t *testing.T) {
//	tMock := &hookTMock{wg: &sync.WaitGroup{}, realT: &realTMock{}}
//	providerMock := &hookProviderMock{
//		suiteMeta: &suiteMetaMockHooks{hook: func(t provider.T) {}},
//		testMeta:  &testMetaMockHooks{},
//	}
//
//	hookFunc := CarriedHook(AfterEach, providerMock.GetTestMeta().GetAfterEach)
//	hookFunc(tMock, providerMock, &sync.WaitGroup{})
//
//	require.False(t, tMock.wgFlag)
//	require.True(t, tMock.failNow)
//	require.True(t, tMock.errorF)
//	require.True(t, tMock.realTF)
//
//	require.True(t, tMock.realT.runTest)
//	require.False(t, tMock.realT.parallelFlag)
//	require.Equal(t, string(AfterEach), tMock.realT.testName)
//
//	require.False(t, providerMock.suiteMeta.baFlag)
//	require.False(t, providerMock.suiteMeta.aaFlag)
//
//	require.False(t, providerMock.testMeta.beFlag)
//	require.True(t, providerMock.testMeta.aeFlag)
//}
