package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrKeyNotFound = errors.New("Key was not found in dictionary")

func (d Dictionary) Search(query string) (string, error) {
	definition, ok := d[query]
	if !ok {
		return "", ErrKeyNotFound
	}
	return definition, nil
}
