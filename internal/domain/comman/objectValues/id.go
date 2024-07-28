package objectValues

import (
	"database/sql/driver"
	"errors"
)

type PrimaryKey string

func NewPrimaryKey(value string) (PrimaryKey, error) {
	if value == "" {
		return "", errors.New("Primary key cannot be empty")
	}
	return PrimaryKey(value), nil
}

// Implementing the driver.Valuer interface for PrimaryKey
func (pk PrimaryKey) Value() (driver.Value, error) {
	return string(pk), nil
}

// Implementing the sql.Scanner interface for PrimaryKey
func (pk *PrimaryKey) Scan(value interface{}) error {
	intValue, ok := value.(string)
	if !ok {
		return errors.New("type assertion to int64 failed")
	}
	*pk = PrimaryKey(intValue)
	return nil
}
