package interfaces

import "github.com/VikashChauhan51/go-sample-api/internal/models"

type BookService interface {
	FetchBooksAsync() (*[]models.Book, error)
}
