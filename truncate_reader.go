package badio

import (
	"io"
)

type truncateReader struct {
	r io.Reader
	n int64
	o int64
}

// EOF is signaled by a zero count with err set to io.EOF.
func NewTruncateReader(r io.Reader, n int64) io.Reader {
	return &truncateReader{r: r, n: n}
}

func (c *truncateReader) Read(p []byte) (n int, err error) {
	// EOF after cutoff
	if c.o > c.n {
		return 0, io.EOF
	}

	// compute read length
	n = len(p)
	if c.o+int64(n) > c.n {
		n = int(c.n - c.o)
	}

	// have we reached EOF?
	if n == 0 {
		return 0, io.EOF
	}

	// real read
	n, err = c.r.Read(p[0:n])
	c.o += int64(n)

	return
}
