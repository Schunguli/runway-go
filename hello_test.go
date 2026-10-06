package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("in Spanish", func(t *testing.T){
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Sasha", "Spanish")
		want := "Hola, Sasha"
		assertCorrectMessage(t, got, want)
	})
	
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("")
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}

 func assertCorrectMessage(t testing.TB, got, want string) {
func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
 }
}
