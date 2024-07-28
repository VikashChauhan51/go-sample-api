package interfaces

import (
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
)

type BookRepository interface {
	CreateBook(bookDTO *dto.BookCreateDTO) error
	FetchBooksAsync() (*[]entities.Book, error)
	FetchBookByID(id int) (*entities.Book, error)
	UpdateBook(bookDTO *dto.BookUpdateDTO) error
	DeleteBook(id int) error
}
