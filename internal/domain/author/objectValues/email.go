package objectValues

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

func NewEmail(value string) (Email, error) {
	if !isValidEmail(value) {
		return Email{}, errors.New("Invalid email format")
	}
	return Email{value: value}, nil
}

func (e Email) Value() string {
	return e.value
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}
