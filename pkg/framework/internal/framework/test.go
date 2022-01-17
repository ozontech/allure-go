package framework

import (
	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider"
)

type InternalTest struct {
	result       *allure.Result
	testFunction func(*provider.T)
	testName     string
}

func NewTest(test func(*provider.T), result *allure.Result) *InternalTest {
	return &InternalTest{
		result:       result,
		testFunction: test,
		testName:     result.Name,
	}
}
