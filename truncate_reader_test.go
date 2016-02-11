package badio

import (
	"io"
	"testing"
)

type truncTest struct {
	bufferSize  int
	truncOffset int64
}

func TestTruncateReader(t *testing.T) {
	tests := []truncTest{
		truncTest{0, 0},
		truncTest{0, 1},
		truncTest{1, 0},
		truncTest{1, 2},
		truncTest{1, 3},
		truncTest{100, 0},
		truncTest{100, 1},
		truncTest{100, 99},
		truncTest{100, 100},
		truncTest{100, 101},
	}

	for i, test := range tests {
		p := make([]byte, test.bufferSize)
		r := NewTruncateReader(NewNullReader(), test.truncOffset)

		var o int = 0
		var err error = nil
		for err == nil {
			var n int
			n, err = r.Read(p)
			o += n
		}

		// make sure got an EOF
		if err != io.EOF {
			t.Errorf("%v", err)
		}

		// make sure the correct number of bytes were read
		if int64(o) != test.truncOffset && test.truncOffset < int64(test.bufferSize) {
			t.Errorf("Read %d bytes instead of %d in test %d (buffer size: %d)", o, test.truncOffset, i+1, len(p))
		}

		// TODO: inspect read bytes and validate the count
	}
}
