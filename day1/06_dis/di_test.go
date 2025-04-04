package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
    buffer := bytes.Buffer{}  // Create a buffer to capture output
    Greet(&buffer, "Chris")   // Call Greet, passing the buffer

    got := buffer.String()    // Get what was written
    want := "Hello, Chris"

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
