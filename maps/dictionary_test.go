package main

import "testing"

func TestSerach(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}

	got := Serach(dictionary, "test")
	want := "this is just a test"

	assertStrings(t, got, want)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
