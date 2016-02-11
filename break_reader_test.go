package badio

import (
	"fmt"
	"strings"
	"testing"
)

func TestBreakReader(t *testing.T) {
	// reader to generate infinite stream of 0x01
	tr := NewSequenceReader([]byte{0x01})

	tests := 1024
	for i := 0; i < tests; i++ {
		// create a big buffer
		p := make([]byte, tests)
		r := NewBreakReader(tr, int64(i))

		// read one byte at a time
		var n, o int
		var err error
		for x := 0; x < tests && err == nil; x++ {
			n, err = r.Read(p[x : x+1])
			o += n
		}

		// ensure an error happened
		if !IsBadIOError(err) {
			t.Fatalf("Expected BadIOError, got: %v", err)
		}

		// make sure break point was accurate
		if o != i {
			t.Fatalf("Expected to read %d bytes, got: %d", i, n)
		}

		// count actual read bytes
		n = 0
		for x := 0; x < len(p); x++ {
			if p[x] != 0 {
				n++
			}
		}

		if n != i {
			t.Fatalf("Expected %d bytes to be changed, got %d", i, n)
		}

		// make sure next read is an error
		n, err = r.Read(p)
		if n != 0 {
			t.Errorf("Expected to read 0 bytes, got %d", n)
		}

		if !IsBadIOError(err) {
			t.Fatalf("Expected BadIOError, got: %v", err)
		}
	}
}

func ExampleNewBreakReader() {
	s := strings.NewReader("banananananananananana")
	r := NewBreakReader(s, 6)

	p := make([]byte, 20)
	_, err := r.Read(p)

	fmt.Printf("Error: %v\n", err)

	// Output: Error: Reader break point at offset 6 (0x6)
}
