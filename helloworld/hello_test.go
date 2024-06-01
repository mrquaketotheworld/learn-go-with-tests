package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to name", func(t *testing.T) {
		got := Hello("Bob", "")
		want := "Hello, Bob"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world, if name empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", french)
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
