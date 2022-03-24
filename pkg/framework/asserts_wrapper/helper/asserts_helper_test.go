package helper

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ozontech/allure-go/pkg/allure"
)

type tAssertMock struct {
}

func (p *tAssertMock) Step(step *allure.Step) {
}

func (p *tAssertMock) Errorf(format string, msgAndArgs ...interface{}) {
}

func (p *tAssertMock) FailNow() {
}

func TestNewAssertsHelper(t *testing.T) {
	h := NewAssertsHelper(&tAssertMock{})
	require.NotNil(t, h)
}
