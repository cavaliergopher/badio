package badio

import (
	"io"
)

type nullReader struct{}

// NewNullReader returns a Reader that implements Read by returning an infinite
// stream of zeros, analogous to `cat /dev/zero`.
func NewNullReader() io.Reader {
	return &nullReader{}
}

func (c *nullReader) Read(p []byte) (int, error) {
	var i int
	for ; i < len(p); i++ {
		p[i] = 0
	}

	return i, nil
}
