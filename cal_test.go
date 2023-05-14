package main

import "testing"

func TestCalScore(t *testing.T) {
	got := CalScore(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
