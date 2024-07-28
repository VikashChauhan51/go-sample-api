package objectValues

import (
	"errors"
)

type LastName struct {
	value string
}

func NewLastName(value string) (LastName, error) {
	if value == "" {
		return LastName{}, errors.New("Last name cannot be empty")
	}
	return LastName{value: value}, nil
}

func (ln LastName) Value() string {
	return ln.value
}
