package internal

import (
	"regexp"
	"strings"

	"github.com/ozontech/allure-go/pkg/allure"
)

type ProviderT interface {
	FailNow()
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	BreakResult(string)
	GetResult() *allure.Result
}

func ExtractErrorMessages(output string) string {
	r := regexp.MustCompile(`Messages:(.*)`)
	result := strings.Trim(strings.TrimPrefix(r.FindString(output), "Messages:   "), " ")
	if result == "" {
		left := "\tError:"
		right := "\tTest:"
		r2 := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
		result = r2.FindString(output)
		result = strings.Trim(strings.TrimSuffix(result, "\tTest:"), " ")
	}
	if result == "" {
		return output
	}
	return result
}

func TestError(contextName, errMsg string, testT ProviderT) {
	short := errMsg
	if len(errMsg) > 100 {
		short = errMsg[:100]
	}
	switch contextName {
	case TestContextName, BeforeEachContextName:
		testT.BreakResult(errMsg)
		testT.Errorf(errMsg)
		testT.FailNow()
	case AfterEachContextName, AfterAllContextName:
		testT.Logf(errMsg)
		testT.GetResult().SetStatusMessage(short)
		testT.GetResult().SetStatusTrace(errMsg)
	case BeforeAllContextName:
		testT.Logf(errMsg)
		testT.GetResult().SetStatusMessage(short)
		testT.GetResult().SetStatusTrace(errMsg)
		testT.FailNow()
	}
}
