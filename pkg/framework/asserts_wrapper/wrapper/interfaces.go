package wrapper

import "time"

// AssertsWrapper ...
type AssertsWrapper interface {
	Equal(provider Provider, expected interface{}, actual interface{}, msgAndArgs ...interface{})
	NotEqual(provider Provider, expected interface{}, actual interface{}, msgAndArgs ...interface{})
	Error(provider Provider, err error, msgAndArgs ...interface{})
	NoError(provider Provider, err error, msgAndArgs ...interface{})
	NotNil(provider Provider, object interface{}, msgAndArgs ...interface{})
	Nil(provider Provider, object interface{}, msgAndArgs ...interface{})
	Len(provider Provider, object interface{}, length int, msgAndArgs ...interface{})
	NotContains(provider Provider, s interface{}, contains interface{}, msgAndArgs ...interface{})
	Contains(provider Provider, s interface{}, contains interface{}, msgAndArgs ...interface{})
	Greater(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	GreaterOrEqual(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	Less(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	LessOrEqual(provider Provider, e1 interface{}, e2 interface{}, msgAndArgs ...interface{})
	Implements(provider Provider, interfaceObject interface{}, object interface{}, msgAndArgs ...interface{})
	Empty(provider Provider, object interface{}, msgAndArgs ...interface{})
	NotEmpty(provider Provider, object interface{}, msgAndArgs ...interface{})
	WithinDuration(provider Provider, expected, actual time.Time, delta time.Duration, msgAndArgs ...interface{})
	JSONEq(provider Provider, expected, actual string, msgAndArgs ...interface{})
	JSONContains(provider Provider, expected, actual string, msgAndArgs ...interface{})
	Subset(provider Provider, list, subset interface{}, msgAndArgs ...interface{})
	IsType(provider Provider, expectedType interface{}, object interface{}, msgAndArgs ...interface{})
	True(provider Provider, value bool, msgAndArgs ...interface{})
	False(provider Provider, value bool, msgAndArgs ...interface{})
	Regexp(provider Provider, rx interface{}, str interface{}, msgAndArgs ...interface{})
}
