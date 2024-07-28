package interfaces

import "github.com/VikashChauhan51/go-sample-api/internal/dto"

type BookService interface {
	CreateBook(bookDTO *dto.BookCreateDTO) error
	FetchBooksAsync() ([]dto.Book, error)
	FetchBookByID(id int) (dto.Book, error)
	UpdateBook(bookDTO *dto.BookUpdateDTO) error
	DeleteBook(id int) error
}
