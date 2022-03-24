package helper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
)

type tRequireMock struct {
}

func (p *tRequireMock) Step(step *allure.Step) {
}

func (p *tRequireMock) Errorf(format string, msgAndArgs ...interface{}) {
}

func (p *tRequireMock) FailNow() {
}

func TestNewRequireHelper(t *testing.T) {
	h := NewAssertsHelper(&tRequireMock{})
	require.NotNil(t, h)
}
