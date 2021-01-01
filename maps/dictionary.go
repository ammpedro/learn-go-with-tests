package dictionary

// Dictionary is a map of words and descriptions
type Dictionary map[string]string

// List of error messages
const (
	ErrNotFound         = DictionaryErr("could not find the word that you were looking for")
	ErrWordExists       = DictionaryErr("could not add existing word")
	ErrWordDoesNotExist = DictionaryErr("could not update word, it does not exist")
)

//DictionaryErr is a an error
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search accepts a word and returns its definition
func (d Dictionary) Search(word string) (string, error) {
	definition, wasFound := d[word]

	if !wasFound {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add accepts a word and its definition then adds it to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update accepts a word and its definition and updates the dictionary6Ejgr4YFMc
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}

	return nil
}

// Delete accepts a word and removes it from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
