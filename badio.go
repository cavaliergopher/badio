package badio

import (
	"io"
)

type badReader struct {
	r             io.Reader
	doBreak       bool
	breakAtOffset int64
	offset        int64
	isBroken      bool
}

func (c *badReader) Read(p []byte) (int, error) {
	// block further reads if already broken
	if c.isBroken {
		return 0, newError("Reader is already broken at offset %d (0x%X)", c.offset, c.offset)
	}

	// enforce break point by reducing buffer size
	if c.doBreak && c.breakAtOffset < int64(len(p)) {
		p = p[:c.breakAtOffset]
	}

	// read into buffer
	n, err := c.r.Read(p)
	if err != nil {
		return n, err
	}

	// increment cursor
	c.offset += int64(n)

	// break?
	if c.doBreak && c.offset == c.breakAtOffset {
		c.isBroken = true
		return n, newError("Reader break point at offset %d (0x%X)", c.offset, c.offset)
	}

	return n, nil
}

func NewBreakReader(r io.Reader, breakAtOffset int64) io.Reader {
	return &badReader{
		r: r,
		//doBreak:       true,
		breakAtOffset: breakAtOffset,
	}
}
