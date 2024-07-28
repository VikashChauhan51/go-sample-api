package objectValues

import (
	"errors"
)

type FirstName struct {
	value string
}

func NewFirstName(value string) (FirstName, error) {
	if value == "" {
		return FirstName{}, errors.New("First name cannot be empty")
	}
	return FirstName{value: value}, nil
}

func (fn FirstName) Value() string {
	return fn.value
}
