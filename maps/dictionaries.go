package main

// ErrNotFound is a dictionary error type that indicates key wasn't found
const ErrNotFound = DictionaryErr("could not find the word you were looking for")

// ErrWordExists is a dictionary error type that indicates key is already in dictionary
const ErrWordExists = DictionaryErr("key already exists")

// ErrWordDoesNotExist is a dictionary error type that indicates key does not exist
const ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")

// DictionaryErr is a type that encapsulates all dictionary error types
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary is a generic map with keys and values as string
type Dictionary map[string]string

// Search searches a string with the given key and returns the value
func (d Dictionary) Search(key string) (string, error) {
	definition, isFound := d[key]

	if !isFound {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add adds word into dictionary
func (d Dictionary) Add(word string, definition string) error {
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

// Update updates the existing word in dictionary
func (d Dictionary) Update(word string, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition

	default:
		return err
	}
	return nil
}

// Delete deletes the existing word in dictionary
func (d Dictionary) Delete(word string) error {
	delete(d, word)
	return nil
}

func main() {

}
