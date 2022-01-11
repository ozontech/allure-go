package provider

import (
	"github.com/koodeex/allure-testify/pkg/allure"
)

/*
Test info
*/

type AllureInfo interface {
	Title(title string)
	Description(description string)
}

// Title changes default test name to title(string)
func (t *T) Title(title string) {
	t.safely(func(result *allure.Result) {
		result.Name = title
	})
}

// Description provides description to test result
func (t *T) Description(description string) {
	t.safely(func(result *allure.Result) {
		result.Description = description
	})
}
