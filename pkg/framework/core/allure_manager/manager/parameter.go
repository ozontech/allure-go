package manager

import "github.com/ozontech/allure-go/pkg/allure"

// WithParameters adds parameters to report in case of current execution context
func (a *allureManager) WithParameters(params ...*allure.Parameter) {
	a.withResult(func(r *allure.Result) {
		r.Parameters = append(r.Parameters, params...)
	})
}

// WithNewParameters adds parameters to report in case of current execution context
func (a *allureManager) WithNewParameters(kv ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Parameters = append(r.Parameters, allure.NewParameters(kv...)...)
	})
}
