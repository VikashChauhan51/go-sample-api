package interfaces

import "github.com/VikashChauhan51/go-sample-api/internal/core/entities"

type BookRepository interface {
	FetchBooksAsync() (*[]entities.Book, error)
}
