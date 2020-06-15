package main

//Dictionary is map
type Dictionary map[string]string

const (
	//ErrNotFound is err
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	//ErrWordExists is err
	ErrWordExists = DictionaryErr("already exist")
	//ErrWordDoesNotExist is err
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

//DictionaryErr is string type
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

//Search is func
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

//Add is func
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

//Update is func
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return nil
	}
	return nil
}
