package helper

import (
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/wrapper"
)

// NewRequireHelper inits new Require interface
func NewRequireHelper(t ProviderT) AssertsHelper {
	return &a{
		t:       t,
		asserts: wrapper.NewRequire(t),
	}
}
