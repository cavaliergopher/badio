package badio

import (
	"fmt"
)

type badReaderError struct {
	s string
}

func newError(format string, a ...interface{}) error {
	return &badReaderError{
		s: fmt.Sprintf(format, a...),
	}
}

func (c *badReaderError) Error() string {
	return c.s
}

func IsBadIOError(err error) bool {
	_, ok := err.(*badReaderError)
	return ok
}
