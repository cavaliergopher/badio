package badio

import (
	"io"
)

type nullReader struct{}

func NewNullReader() io.Reader {
	return &nullReader{}
}

func (c *nullReader) Read(p []byte) (int, error) {
	var i int = 0
	for ; i < len(p); i++ {
		p[i] = 0
	}

	return i, nil
}
