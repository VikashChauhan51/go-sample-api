package objectValues

import (
	"errors"

	"github.com/VikashChauhan51/go-sample-api/internal/domain/comman/objectValues"
)

type BookId struct {
	value objectValues.PrimaryKey
}

func NewBookID(value objectValues.PrimaryKey) (BookId, error) {
	if value == "" {
		return BookId{}, errors.New("ID cannot be empty")
	}
	return BookId{value: value}, nil
}

func (id BookId) Value() objectValues.PrimaryKey {
	return id.value
}
