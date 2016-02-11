package badio

import (
	"fmt"
	"testing"
)

func TestZeroLengthSequenceReader(t *testing.T) {
	r := NewSequenceReader([]byte{})
	p := make([]byte, 1024)
	if n, err := r.Read(p); err == nil || n != 0 {
		t.Errorf("Expected error on zero-length sequence. Read %d bytes with err: %v", n, err)
	}
}

func TestSequenceReader(t *testing.T) {
	s := "_.-^-."
	r := NewSequenceReader([]byte(s))
	p := make([]byte, 1024)
	o := 0
	for i := 0; o < len(p); i++ {
		// increase buffer window size each iteration
		b := i
		if o+b > len(p) {
			b = len(p) - o
		}

		// read sequence
		n, err := r.Read(p[o : o+b])
		if err != nil {
			t.Fatalf("%v", err)
		}

		if n != b {
			t.Errorf("Expected to read %d bytes, got %d", b, n)
		}

		// move cursor
		o += i
	}

	// validate output
	for i := 0; i < len(p); i += len(s) {
		l := len(s)
		if i+l > len(p) {
			l = len(p) - i
		}

		if string(p[i:i+l]) != s[:l] {
			t.Fatalf("Bad sequence output as offset %d", i)
		}
	}
}

func ExampleNewSequenceReader() {
	r := NewSequenceReader([]byte("na"))

	p := make([]byte, 20)
	r.Read(p)

	fmt.Printf("ba%s\n", p)

	// Output: banananananananananana
}
