package require

import (
	"time"

	"github.com/ozontech/allure-go/pkg/allure"
	"github.com/ozontech/allure-go/pkg/framework/asserts_wrapper/wrapper"
)

type ProviderT interface {
	Step(step *allure.Step)
	Errorf(format string, args ...interface{})
	FailNow()
}

func Equal(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Equal(t, expected, actual, msgAndArgs...)
}

func NotEqual(t ProviderT, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotEqual(t, expected, actual, msgAndArgs...)
}

func Error(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Error(t, err, msgAndArgs...)
}

func NoError(t ProviderT, err error, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NoError(t, err, msgAndArgs...)
}

func NotNil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotNil(t, object, msgAndArgs...)
}

func Nil(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Nil(t, object, msgAndArgs...)
}

func Len(t ProviderT, object interface{}, length int, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Len(t, object, length, msgAndArgs...)
}

func NotContains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotContains(t, s, contains, msgAndArgs...)
}

func Contains(t ProviderT, s interface{}, contains interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Contains(t, s, contains, msgAndArgs...)
}

func Greater(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Greater(t, e1, e2, msgAndArgs...)
}

func GreaterOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).GreaterOrEqual(t, e1, e2, msgAndArgs...)
}

func Less(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Less(t, e1, e2, msgAndArgs...)
}

func LessOrEqual(t ProviderT, e1 interface{}, e2 interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).LessOrEqual(t, e1, e2, msgAndArgs...)
}

func Implements(t ProviderT, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Implements(t, interfaceObject, object, msgAndArgs...)
}

func Empty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Empty(t, object, msgAndArgs...)
}

func NotEmpty(t ProviderT, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).NotEmpty(t, object, msgAndArgs...)
}

func WithinDuration(t ProviderT, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

func JSONEq(t ProviderT, expected, actual string, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).JSONEq(t, expected, actual, msgAndArgs...)
}

func Subset(t ProviderT, list, subset interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).Subset(t, list, subset, msgAndArgs...)
}

func IsType(t ProviderT, expectedType interface{}, object interface{}, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).IsType(t, expectedType, object, msgAndArgs...)
}

func True(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).True(t, value, msgAndArgs...)
}

func False(t ProviderT, value bool, msgAndArgs ...interface{}) {
	wrapper.NewRequire(t).False(t, value, msgAndArgs...)
}
