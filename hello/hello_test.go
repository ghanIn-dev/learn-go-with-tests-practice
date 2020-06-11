package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("seokhyun")
	want := "Hello seokhyun"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
