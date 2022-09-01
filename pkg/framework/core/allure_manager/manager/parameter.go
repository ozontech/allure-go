package manager

import "github.com/ozontech/allure-go/pkg/allure"

// WithParameters adds parameters to report in case of current execution context
func (a *allureManager) WithParameters(params ...*allure.Parameter) {
	a.safely(func(result *allure.Result) {
		result.Parameters = append(result.Parameters, params...)
	})
}

// WithNewParameters adds parameters to report in case of current execution context
func (a *allureManager) WithNewParameters(kv ...interface{}) {
	a.safely(func(result *allure.Result) {
		result.Parameters = append(result.Parameters, allure.NewParameters(kv...)...)
	})
}
