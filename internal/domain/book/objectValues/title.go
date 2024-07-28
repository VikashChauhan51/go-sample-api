package objectValues

import (
	"database/sql/driver"
	"errors"
)

type Title string

func NewTitle(value string) (Title, error) {
	if value == "" {
		return "", errors.New("Title cannot be empty")
	}
	return Title(value), nil
}

func (t Title) Value() (driver.Value, error) {
	return string(t), nil
}

func (t *Title) Scan(value interface{}) error {
	strValue, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string failed")
	}
	*t = Title(strValue)
	return nil
}
