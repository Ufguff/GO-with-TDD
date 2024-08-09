package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound          = DictionaryErr("could not find word in dictionary")
	ErrWordExists        = DictionaryErr("word already exists")
	ErrWordDoesNotExists = DictionaryErr("word does not exists")
)

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExists
	case nil:
		d[word] = definition // рефактор
	default:
		return err
	}

	return nil
}

func (d Dictionary) Search(key string) (string, error) {
	result, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return result, nil
}
func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
