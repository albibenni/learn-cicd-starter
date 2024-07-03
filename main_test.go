package main

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	want := 2 + 3
	if want != 5 {
		t.Fatalf(`got %q, want 5`, want)
	}
}
