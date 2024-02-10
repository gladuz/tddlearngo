package hello_world

import "testing"

func TestHello(t *testing.T) {

	t.Run("Saying hello chris", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying with empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying with lang", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying with french", func(t *testing.T) {
		got := Hello("Jorj", "French")
		want := "Bonjour, Jorj"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	// To tell the test suite that this function is helper so that it reports the calling line number
	// Basically treats the function as inline code
	t.Helper()
	if got != want {
		t.Errorf("got %q want %v", got, want)
	}
}
