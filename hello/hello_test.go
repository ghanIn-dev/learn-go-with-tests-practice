package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessge := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to peole", func(t *testing.T) {
		got := Hello("seokhyun", "")
		want := "Hello seokhyun"

		assertCorrectMessge(t, got, want)
	})

	t.Run("saying 'Hello world' when an empty stirng is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello world"

		assertCorrectMessge(t, got, want)
	})

	t.Run("in Korean", func(t *testing.T) {
		got := Hello("석현", "Korea")
		want := "Hello 석현"
		assertCorrectMessge(t, got, want)
	})
}
