package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := "Hello, World!"
	want := "Hello, World!"
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
