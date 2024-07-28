package objectValues

import (
	"errors"

	"github.com/VikashChauhan51/go-sample-api/internal/domain/comman/objectValues"
)

type AuthorId struct {
	value objectValues.PrimaryKey
}

func NewAuthorID(value objectValues.PrimaryKey) (AuthorId, error) {
	if value == "" {
		return AuthorId{}, errors.New("ID cannot be empty")
	}
	return AuthorId{value: value}, nil
}

func (id AuthorId) Value() objectValues.PrimaryKey {
	return id.value
}
