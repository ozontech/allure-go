package suite

import (
	"reflect"
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

func (s *Suite) RunSuite(t provider.T, suite runner.TestSuite) runner.SuiteResult {
	t.SkipOnPrint()
	parts := strings.Split(t.RealT().Name(), "/")
	parentName := parts[len(parts)-3]

	return runner.NewSuiteRunnerWithParent(t.RealT(), getPackage(2), cleanName(getSuiteName(suite)), parentName, suite).RunTests()
}

func (s *Suite) RunNamedSuite(t provider.T, suiteName string, suite runner.TestSuite) runner.SuiteResult {
	t.SkipOnPrint()
	parts := strings.Split(t.RealT().Name(), "/")
	parentName := parts[len(parts)-3]

	return runner.NewSuiteRunnerWithParent(t.RealT(), getPackage(2), suiteName, parentName, suite).RunTests()
}

func RunSuite(t provider.TestingT, suite runner.TestSuite) runner.SuiteResult {
	return runner.NewSuiteRunner(t, getPackage(2), getSuiteName(suite), suite).RunTests()
}

func RunNamedSuite(t provider.TestingT, suiteName string, suite runner.TestSuite) runner.SuiteResult {
	return runner.NewSuiteRunner(t, getPackage(2), suiteName, suite).RunTests()
}

func getSuiteName(suite interface{}) string {
	s := reflect.TypeOf(suite)
	if s.Kind() == reflect.Ptr {
		return s.Elem().Name()
	}

	return s.Name()
}

func cleanName(fullName string) string {
	nameParts := strings.Split(fullName, "/")

	removeIDs := make([]int, 0, len(nameParts))

	for idx, namePart := range nameParts {
		if strings.HasSuffix(namePart, "_Tests") {
			removeIDs = append(removeIDs, idx)
		}
	}

	for _, idx := range removeIDs {
		nameParts = remove(nameParts, idx)
	}

	return strings.Join(nameParts, "/")
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
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
