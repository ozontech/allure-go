package internal

import (
	"regexp"
	"strings"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type ProviderT interface {
	FailNow()
	Logf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	BreakResult(string)
	GetResult() *allure.Result
	Provider() provider.Provider
}

func ExtractErrorMessages(output string) string {
	r := regexp.MustCompile(`Messages:(.*)`)
	result := strings.TrimPrefix(r.FindString(output), "Messages:   ")
	left := "\tError:"
	right := "\tTest:"
	if result == "" {
		r2 := regexp.MustCompile(`(?s)` + regexp.QuoteMeta(left) + `(.*?)` + regexp.QuoteMeta(right))
		result = r2.FindString(output)
		result = strings.Trim(strings.TrimSuffix(result, "\tTest:"), " ")
	}
	if result == "" {
		return output
	}
	return result
}

func TestError(errMsg string, testT ProviderT) {
	switch testT.Provider().ExecutionContext().GetName() {
	case TestContextName, BeforeEachContextName:
		testT.BreakResult(errMsg)
		testT.Errorf(errMsg)
		testT.FailNow()
	case AfterEachContextName, AfterAllContextName:
		testT.Logf(errMsg)
		testT.GetResult().StatusDetails.Message = errMsg[:100]
		testT.GetResult().StatusDetails.Trace = errMsg
	case BeforeAllContextName:
		testT.Logf(errMsg)
		testT.GetResult().StatusDetails.Message = errMsg[:100]
		testT.GetResult().StatusDetails.Trace = errMsg
		testT.FailNow()
	}
}
