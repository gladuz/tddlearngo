package maps

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("search for word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStringEqual(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")
		if err != ErrKeyNotFound {
			t.Fatal("expected an error")
		}
	})
}

func assertStringEqual(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("wanted %q, got %q, given %q", want, got, "test")
	}
}
