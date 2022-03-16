package wrapper

import (
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type ProviderT interface {
	WithNewStep(stepName string, f func(ctx provider.StepCtx), params ...allure.Parameter)
	Errorf(format string, args ...interface{})
}

// AssertsWrapper ...
type AssertsWrapper interface {
	Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
	NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{})
	Error(t ProviderT, err error, msgAndArgs ...interface{})
	NoError(t ProviderT, err error, msgAndArgs ...interface{})
	NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{})
	Nil(t ProviderT, object interface{}, msgAndArgs ...interface{})
	Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{})
	NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{})
	Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{})
	Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
	Empty(t ProviderT, object interface{}, msgAndArgs ...interface{})
	NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{})
	WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{})
	JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{})
	Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{})
	IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{})
	True(t ProviderT, value bool, msgAndArgs ...interface{})
	False(t ProviderT, value bool, msgAndArgs ...interface{})
}
