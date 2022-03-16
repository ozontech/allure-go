package require

import (
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/provider/pkg/asserts_wrapper/wrapper"
	"github.com/ozontech/allure-go/pkg/provider/pkg/provider"
)

type ProviderT interface {
	WithNewStep(stepName string, f func(ctx provider.StepCtx), params ...allure.Parameter)
	Errorf(format string, args ...interface{})
}

func Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Equal(t, expected, actual, msgAndArgs...)
}

func NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().NotEqual(t, expected, actual, msgAndArgs...)
}

func Error(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Error(t, err, msgAndArgs...)
}

func NoError(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewRequire().NoError(t, err, msgAndArgs...)
}

func NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().NotNil(t, object, msgAndArgs...)
}

func Nil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Nil(t, object, msgAndArgs...)
}

func Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Len(t, object, length, msgAndArgs...)
}

func NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().NotContains(t, s, contains, msgAndArgs...)
}

func Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Contains(t, s, contains, msgAndArgs...)
}

func Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Greater(t, e1, e2, msgAndArgs...)
}

func GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().GreaterOrEqual(t, e1, e2, msgAndArgs...)
}

func Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Less(t, e1, e2, msgAndArgs...)
}

func LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().LessOrEqual(t, e1, e2, msgAndArgs...)
}

func Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Implements(t, interfaceObject, object, msgAndArgs...)
}

func Empty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Empty(t, object, msgAndArgs...)
}

func NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().NotEmpty(t, object, msgAndArgs...)
}

func WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	wrapper.NewRequire().WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

func JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	wrapper.NewRequire().JSONEq(t, expected, actual, msgAndArgs...)
}

func Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().Subset(t, list, subset, msgAndArgs...)
}

func IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire().IsType(t, expectedType, object, msgAndArgs...)
}

func True(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewRequire().True(t, value, msgAndArgs...)
}

func False(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewRequire().False(t, value, msgAndArgs...)
}
