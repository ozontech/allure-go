package suite

import (
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestSuite_SetRunner(t *testing.T) {
	suite := Suite{}
	r := runner.NewRunner(new(testing.T), "suiteName")
	suite.SetRunner(r)
	require.Equal(t, r, suite.runner)
}

func TestSuite_GetRunner(t *testing.T) {
	r := runner.NewRunner(new(testing.T), "suiteName")
	suite := Suite{runner: r}
	require.Equal(t, r, suite.GetRunner())
}

type TestSuiteRunSuite struct {
	Suite
	s2 *TestSuiteRunSuite2
}

func (s *TestSuiteRunSuite) TestSome(t provider.T) {
	s.s2 = new(TestSuiteRunSuite2)
	s.RunSuite(t, s.s2)
}

type TestSuiteRunSuite2 struct {
	Suite

	wg        sync.WaitGroup
	testSome1 bool

	beforeAll  bool
	beforeEach bool
	afterEach  bool
	afterAll   bool
}

func (s *TestSuiteRunSuite2) BeforeAll(t provider.T) {
	s.beforeAll = true
}

func (s *TestSuiteRunSuite2) BeforeEach(t provider.T) {
	s.beforeEach = true
}

func (s *TestSuiteRunSuite2) AfterEach(t provider.T) {
	s.afterEach = true
}

func (s *TestSuiteRunSuite2) AfterAll(t provider.T) {
	s.afterAll = true
}

func (s *TestSuiteRunSuite2) TestSome(t provider.T) {
	s.testSome1 = true
}

func TestSuite_RunSuite(t *testing.T) {
	allureDir := "./allure-results"
	defer os.RemoveAll(allureDir)

	suite := new(TestSuiteRunSuite)
	mockT := &suiteRunnerTMock{t: t}
	r := NewSuiteRunner(mockT, "packageName", "suiteName", suite)
	r.RunTests()

	// subtests have been run
	require.True(t, suite.s2.testSome1)
	require.True(t, suite.s2.beforeEach)
	require.True(t, suite.s2.beforeAll)
	require.True(t, suite.s2.afterEach)
	require.True(t, suite.s2.afterAll)
}
