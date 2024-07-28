package objectValues

import (
	"errors"
)

type Title struct {
	value string
}

func NewTitle(value string) (Title, error) {
	if value == "" {
		return Title{}, errors.New("Title cannot be empty")
	}
	return Title{value: value}, nil
}

func (t Title) Value() string {
	return t.value
}
