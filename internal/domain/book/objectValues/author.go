package objectValues

import (
	"errors"
)

type Author struct {
	value string
}

func NewAuthor(value string) (Author, error) {
	if value == "" {
		return Author{}, errors.New("Author cannot be empty")
	}
	return Author{value: value}, nil
}

func (a Author) Value() string {
	return a.value
}
