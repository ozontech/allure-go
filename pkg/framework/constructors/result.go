package constructors

import (
	"github.com/koodeex/allure-testify/pkg/allure"
	"github.com/koodeex/allure-testify/pkg/framework/internal/structs"
)

// ResultHelper provides easy way to construct allure.Result struct
type ResultHelper struct {
}

// NewResultHelper ...
func NewResultHelper() *ResultHelper {
	return &ResultHelper{}
}

func (*ResultHelper) GetNewResultFromOpts(opts *structs.Options) *allure.Result {
	result := allure.NewResult(opts.TestName, opts.FullName).
		WithFrameWork(opts.Framework).
		WithHost(opts.Host).
		WithPackage(opts.PackageName).
		WithThread(opts.ThreadName).
		WithLanguage(opts.Language).
		WithLaunchTags()

	if opts.SuiteName != "" {
		result = result.WithSuite(opts.SuiteName)
	}

	if opts.ParentName != "" {
		result = result.WithParentSuite(opts.ParentName)
	}
	return result
}
