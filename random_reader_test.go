package badio

import (
	"testing"
)

func TestRandomReader(t *testing.T) {
	r := NewRandomReader()
	p := make([]byte, 4096)
	n, err := r.Read(p)
	if err != nil {
		t.Fatalf("%v", err)
	}
	if n != len(p) {
		t.Errorf("Expected to write %d random bytes, got: %d", len(p), n)
	}

	// make sure buffer is non-zero
	var i int
	for ; i < len(p) && p[i] == 0; i++ {
	}

	if i == len(p) {
		t.Errorf("Expected random data in buffer, got zeros.")
	}
}
