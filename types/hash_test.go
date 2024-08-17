package types

import "testing"

func TestHashIsZero(t *testing.T) {
	var h Hash
	if !h.IsZero() {
		t.Error("expected hash to be zero")
	}
}
