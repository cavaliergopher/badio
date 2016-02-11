package badio

import (
	"fmt"
	"testing"
)

func TestIsBadIOError(t *testing.T) {
	if IsBadIOError(fmt.Errorf("Not a BadIOError")) {
		t.Errorf("IsBadIOError returned true for a non-BadIOError error")
	}

	if !IsBadIOError(newError("A BadIOError")) {
		t.Errorf("IsBadIOError returned false for a BadIOError error")
	}
}

func TestBreakReader(t *testing.T) {
	// reader to generate infinite data stream
	tr := NewNullReader()

	// read buffer (4 bytes at a time)
	p := make([]byte, 4)

	// test breaks at the beginning, end and out of range
	for _, b := range []int64{0, 1, 2, 3, 4, 99, 100, 101} {
		// create new BreakReader
		r := NewBreakReader(tr, b)

		// keep reading until we exceed the desired breakpoint
		var n, o int = 0, 0
		var err error = nil
		for int64(o) <= b && err == nil {
			n, err = r.Read(p)
			o += n
		}

		t.Logf("%d/%d: %v", o, b, err)

		// in range tests
		if int64(o) <= b {
			if !IsBadIOError(err) {
				t.Errorf("Expected break error for break point %d. Got: %v", b, err)
			}

			if int64(o) != b {
				t.Errorf("Expected to read %d bytes. Got: %d", o, n)
			}
		} else {
			// out of range tests
			if IsBadIOError(err) {
				t.Errorf("Expected no break error for break point %d. Got: %v", b, err)
			}

			if n != len(p) {
				t.Errorf("Expected out-of-range read to read %d bytes. Got: %d", len(p), n)
			}
		}
	}
}
