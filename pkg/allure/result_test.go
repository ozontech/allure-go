package allure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResultImplements(t *testing.T) {
	assert.Implements(t, (*WithAttachments)(nil), new(Result))
	assert.Implements(t, (*WithTimer)(nil), new(Result))
	assert.Implements(t, (*Printable)(nil), new(Result))
	assert.Implements(t, (*IResult)(nil), new(Result))
}
