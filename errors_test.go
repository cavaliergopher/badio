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
