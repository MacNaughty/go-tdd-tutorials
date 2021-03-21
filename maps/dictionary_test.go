package maps

import "testing"

func TestSearch(t *testing.T) {
	//dictionary := map[string]string{"test": "this is just a test"}
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t1 *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t1 *testing.T) {
		_, err := dictionary.Search("dne")

		assertError(t1, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t1 *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"
		err := dictionary.Add(word, definition)

		assertError(t1, err, nil)
		assertDefinition(t1, dictionary, word, definition)
	})

	t.Run("existing word", func(t1 *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t1, err, ErrWordExists)
		assertDefinition(t1, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t1 *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		dictionary.Update(word, newDefinition)

		assertDefinition(t1, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t1 *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Update(word, definition)

		assertError(t1, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	definition := "this is a test"

	dictionary := Dictionary{word: definition}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func assertDefinition(t1 testing.TB, dictionary Dictionary, word string, definition string) {
	t1.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t1.Fatal("should find added word: ", err)
	}

	if definition != got {
		t1.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t1 testing.TB, got error, want error) {
	t1.Helper()

	if got != want {
		t1.Errorf("got error %q want %q", got, want)
	}
}

func assertStrings(t1 *testing.T, got string, want string) {
	t1.Helper()
	if got != want {
		t1.Errorf("got %s want %s, given %s", got, want, "test")
	}
}
