package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to guys", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"
		checkError(t, got, want)
	})

	t.Run("say 'Hello World' when string is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		checkError(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Amigo", "Spanish")
		want := "Hola Amigo"
		checkError(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("TransMan", "French")
		want := "Bonjour TransMan"
		checkError(t, got, want)
	})
}

func checkError(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q!", got, want)
	}
}
