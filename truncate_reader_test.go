package badio

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestTruncateReader(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(s); i++ {
		// create a full size buffer
		p := make([]byte, len(s))

		// create truncate read to truncate at i
		r := NewTruncateReader(strings.NewReader(s), int64(i))

		// read one byte at a time
		var n, o int
		var err error
		for x := 0; x < len(s) && err == nil; x++ {
			n, err = r.Read(p[x : x+1])
			o += n
		}

		// ensure we reach EOF
		if err != io.EOF {
			t.Fatalf("Expected io.EOF, got: %v", err)
		}

		// make sure break point was accurate
		if o != i {
			t.Fatalf("Expected to read %d bytes, got: %d", i, n)
		}

		// validate new string
		out := string(bytes.Trim(p, "\x00"))
		if out != s[:i] {
			t.Errorf("Expected '%s', got: '%s'", s[:i], out)
		}

		// make sure next read is io.EOF
		n, err = r.Read(p)
		if n != 0 {
			t.Errorf("Expected to read 0 bytes, got %d", n)
		}

		if err != io.EOF {
			t.Fatalf("Expected io.EOF, got: %v", err)
		}
	}
}

func ExampleNewTruncateReader() {
	s := strings.NewReader("banananananananananana")
	r := NewTruncateReader(s, 6)

	p := make([]byte, 20)
	r.Read(p)

	fmt.Printf("%s\n", bytes.Trim(p, "\x00"))

	// Output: banana
}
