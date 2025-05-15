package manager

import (
	"fmt"

	"github.com/ozontech/allure-go/pkg/allure"
)

// Title changes default test name to title(using fmt.Sprint)
func (a *allureManager) Title(args ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Name = fmt.Sprint(args...)
	})
}

// Titlef changes default test name to title(using fmt.Sprintf)
func (a *allureManager) Titlef(format string, args ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Name = fmt.Sprintf(format, args...)
	})
}

// Description provides description to test result(using fmt.Sprint)
func (a *allureManager) Description(args ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Description = fmt.Sprint(args...)
	})
}

// Descriptionf provides description to test result(using fmt.Sprintf)
func (a *allureManager) Descriptionf(format string, args ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Description = fmt.Sprintf(format, args...)
	})
}

// Stage provides staqe to test result(using fmt.Sprint)
func (a *allureManager) Stage(args ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Stage = fmt.Sprint(args...)
	})
}

// Stagef provides staqe to test result(using fmt.Sprintf)
func (a *allureManager) Stagef(format string, args ...interface{}) {
	a.withResult(func(r *allure.Result) {
		r.Stage = fmt.Sprintf(format, args...)
	})
}
