package allure

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ResultHelper provides easy way to construct allure.Result struct
type ResultHelper struct {
}

// NewResultHelper ...
func NewResultHelper() *ResultHelper {
	return &ResultHelper{}
}

// TestingT interface that allow threw testing.T wrapper to the helper
type TestingT interface {
	Name() string
	GetResult() *Result
	GetSuite() string
	GetPackage() string
}

// GetNewResult returns pointer to the new allure.Result exemplar.
func (*ResultHelper) GetNewResult(t TestingT, testName, packageName string, tags ...string) *Result {
	var suiteName string

	host, _ := os.Hostname()
	fullName := fmt.Sprintf("%s/%s", t.Name(), testName)
	callers := strings.Split(t.Name(), "/")

	result := NewResult(testName, fullName).
		WithFrameWork(DefaultVersion).
		WithHost(host).
		WithThread(fullName).
		WithLanguage(runtime.Version()).
		WithLaunchTags()

	var newTags []Label
	for _, tag := range tags {
		newTags = append(newTags, NewLabel(Tag, tag))
	}
	result.SetLabels(newTags...)

	if t.GetResult() != nil {
		prepareResultFromExisted(t.GetResult(), result)
	} else {
		if t.GetPackage() == "" {
			suiteName = callers[len(callers)-1]
		} else {
			suiteName = t.GetSuite()
		}

		if packageName == "" {
			if t.GetPackage() != "" {
				packageName = t.GetPackage()
			} else {
				packageName = "default"
			}
		}

		result = result.
			WithSuite(suiteName).
			WithPackage(packageName)
	}
	return result
}

// prepareResultFromExisted copy data from existed allure.Result
func prepareResultFromExisted(parentResult, childResult *Result) {
	childResult.WithSuite(parentResult.Name)

	if labels := parentResult.GetLabel(Suite); len(labels) != 0 {
		childResult.WithParentSuite(labels[0].Value)
	}
	if labels := parentResult.GetLabel(Package); len(labels) != 0 {
		childResult.SetLabels(labels[0])
	}
	if labels := parentResult.GetLabel(Feature); len(labels) != 0 {
		childResult.SetLabels(labels...)
	}
	if labels := parentResult.GetLabel(Story); len(labels) != 0 {
		childResult.SetLabels(labels...)
	}
	if labels := parentResult.GetLabel(Epic); len(labels) != 0 {
		childResult.SetLabels(labels...)
	}
	if labels := parentResult.GetLabel(Tag); len(labels) != 0 {
		childResult.SetLabels(labels...)
	}
}
