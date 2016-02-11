package badio

import (
	"io"
)

type breakReader struct {
	r io.Reader
	n int64
	o int64
}

// NewBreakReader returns a Reader that behaves like r except that it will
// return an error once it has read n bytes.
func NewBreakReader(r io.Reader, n int64) io.Reader {
	return &breakReader{r: r, n: n}
}

func (c *breakReader) Read(p []byte) (int, error) {
	// block further reads if already broken
	if c.o > c.n {
		return 0, newError("Reader is already broken at offset %d (0x%X)", c.o, c.o)
	}

	// enforce break point by reducing buffer size
	if c.n < int64(len(p)) {
		p = p[:c.n]
	}

	// read into buffer
	n, err := c.r.Read(p)
	if err != nil {
		return n, err
	}

	// increment cursor
	c.o += int64(n)

	// break?
	if c.o >= c.n {
		return n, newError("Reader break point at offset %d (0x%X)", c.o, c.o)
	}

	return n, nil
}
