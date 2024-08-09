package maps

import "testing"

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is test"

	dictionary := Dictionary{word: definition}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertError(t, err, ErrNotFound)
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just test"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is new test"

		dictionary.Update(word, newDefinition)
		assertDefinition(t, dictionary, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just test"

		err := dictionary.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just test"
		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})
	t.Run("word not new", func(t *testing.T) {
		key := "test"
		value := "this is just test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just test"}

	t.Run("word that in dict", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just test"
		assertSearch(t, got, want)
	})
	t.Run("word that not in dict", func(t *testing.T) {
		_, err := dictionary.Search("word")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertError(t, err, ErrNotFound)
	})
}

func assertSearch(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("should add a word:", err)
	}
	assertSearch(t, got, value)
}
