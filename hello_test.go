package main

import "testing"

func TestHello(t *testing.T) {
	AssesrtCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want % q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		AssesrtCorrectMessage(t, got, want)
	})
	t.Run("saying hello world when empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		AssesrtCorrectMessage(t, got, want)
	})
	t.Run("test in spanish", func(t *testing.T) {
		got := Hello("Eldoe", "Spanish")
		want := "Hola, Eldoe"
		AssesrtCorrectMessage(t, got, want)
	})
}
