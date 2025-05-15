package allure

import (
	"strings"
)

// Copied from https://cs.opensource.google/go/go/+/refs/tags/go1.24.3:src/errors/join.go;l=40
// Since this go version does not have errors.Join
//
// TODO: delete this in v2
type joinError struct {
	errs []error
}

func (e *joinError) Error() string {
	b := make([]string, 0, len(e.errs))

	for _, err := range e.errs {
		b = append(b, err.Error())
	}

	return strings.Join(b, "\n")
}

func (e *joinError) Unwrap() []error {
	return e.errs
}
