package badio

import (
	"testing"
)

func TestByteReader(t *testing.T) {
	// test bytes reader
	for i := 0; i < 0xFF; i++ {
		b := make([]byte, 4096)
		r := NewByteReader(byte(i))

		n, err := r.Read(b)
		if err != nil {
			t.Fatalf("Error writing 0x%X: %v", i, err)
		}

		if n != len(b) {
			t.Errorf("Expected to read %d bytes, got %d", len(b), n)
		}

		// validate each byte in the written buffer
		for x := 0; x < len(b); x++ {
			if b[x] != byte(i) {
				t.Fatalf("Expected value at buffer offet %d to have value %d", x, i)
			}
		}
	}
}

func TestNullReader(t *testing.T) {
	b := make([]byte, 4096)
	r := NewNullReader()

	n, err := r.Read(b)
	if err != nil {
		t.Fatalf("%v", err)
	}

	if n != len(b) {
		t.Errorf("Expected to read %d bytes, got %d", len(b), n)
	}

	// validate each byte in the written buffer
	for x := 0; x < len(b); x++ {
		if b[x] != 0 {
			t.Fatalf("Expected value at buffer offet %d to be 0", x)
		}
	}
}
