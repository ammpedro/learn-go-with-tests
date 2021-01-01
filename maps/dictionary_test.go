package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {

	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}

	t.Run("should return definition for known word", func(t *testing.T) {

		got, _ := dictionary.Search(word)
		want := definition

		assertStrings(t, got, want)
	})
	t.Run("should return definition for unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {

	t.Run("should handle new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("should handle existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "test existing word")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should handle existing word", func(t *testing.T) {
		word := "test"
		definition := "this is the initial definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "this is an updated definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("should handle new word", func(t *testing.T) {
		word := "test"
		definition := "this is the initial definition"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should handle existing word", func(t *testing.T) {
		word := "test"
		definition := "this is the initial definition"
		dictionary := Dictionary{word: definition}

		dictionary.Delete(word)
		_, err := dictionary.Search(word)
		if err != ErrNotFound {
			t.Errorf("expected word: %q to be deleted", word)
		}
	})

	t.Run("should handle empty dictionary", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		// delete(d, key) does not throw an error when key to be deleted does not exist
		dictionary.Delete(word)
		_, err := dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error: %q, but want %q", got, want)
	}
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatalf("expected to find added word: %q but got error: %q", word, err)
	}

	assertStrings(t, got, definition)
}
