package suite

import (
	"runtime"
	"strings"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

type Suite struct {
	runner runner.TestRunner
}

func (s *Suite) GetRunner() runner.TestRunner {
	return s.runner
}

func (s *Suite) SetRunner(runner runner.TestRunner) {
	s.runner = runner
}

func (s *Suite) RunSuite(t provider.T, suite runner.InternalSuite) {
	t.SkipOnPrint()
	RunSuite(t.RealT(), suite)
}

func RunSuite(t provider.TestingT, suite runner.InternalSuite) map[string]bool {
	return runner.NewSuiteRunner(t, getPackage(2), t.Name(), suite).RunTests()
}

func RunNamedSuite(t provider.TestingT, suiteName string, suite runner.InternalSuite) map[string]bool {
	return runner.NewSuiteRunner(t, getPackage(2), suiteName, suite).RunTests()
}

func getPackage(depth int) string {
	pc, _, _, _ := runtime.Caller(depth)
	funcName := runtime.FuncForPC(pc).Name()
	lastSlash := strings.LastIndexByte(funcName, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(funcName[lastSlash:], '.') + lastSlash
	return funcName[:lastDot]
}
