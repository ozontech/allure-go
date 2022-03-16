package common

import "github.com/ozontech/allure-go/pkg/allure"

// Title changes default test name to title(string)
func (c *common) Title(title string) {
	c.safely(func(result *allure.Result) {
		result.Name = title
	})
}

// Description provides description to test result
func (c *common) Description(description string) {
	c.safely(func(result *allure.Result) {
		result.Description = description
	})
}
