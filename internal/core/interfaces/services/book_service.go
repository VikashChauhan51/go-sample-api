package interfaces

import "github.com/VikashChauhan51/go-sample-api/internal/dto"

type BookService interface {
	FetchBooksAsync() (*[]dto.Book, error)
}
