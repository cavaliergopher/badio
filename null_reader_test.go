package badio

import (
	"bytes"
	"testing"
)

func TestNullReader(t *testing.T) {
	b := bytes.Repeat([]byte{0xFF}, 1024)
	r := NewNullReader()

	n, err := r.Read(b)

	if err != nil {
		t.Fatalf("%v", err)
	}

	if n != len(b) {
		t.Errorf("Expected to read %d bytes, got %d", len(b), n)
	}

	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			t.Fatalf("Non-zero at offset %d", i)
		}
	}
}
