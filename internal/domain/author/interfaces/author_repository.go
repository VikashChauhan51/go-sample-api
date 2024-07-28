package interfaces

import (
	"github.com/VikashChauhan51/go-sample-api/internal/domain/author/entities"
)

type AuthorRepository interface {
	FetchAuthorsAsync() (*[]entities.Author, error)
	FetchAuthorByID(id string) (*entities.Author, error)
	DeleteAuthor(id string) error
}
