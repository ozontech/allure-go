// Package runner provides a number of ways to run tests and groups of tests,
// both using suites similar to those of `testify`,
// and using the classic test form of the standard `testing` library.
//
// For quick and simple tests it is recommended to use the runner.RunTest function.
// To run `suite.Suite` - static function runner.RunSuite.
// For launching groups of tests - the structure runner.TestRunner.
//
// Note: only allure-testify library suites are supported at the moment
// Learn More about go test: https://golang.org/doc/tutorial/add-a-test
package runner
