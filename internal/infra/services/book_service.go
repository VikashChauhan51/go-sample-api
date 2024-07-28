package services

import (
	repo "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces/repositories"
	svc "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces/services"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
)

type BookService struct {
	bookRepository repo.BookRepository
}

func NewBookService(bookRepository repo.BookRepository) svc.BookService {
	return &BookService{bookRepository}
}

func (b *BookService) CreateBook(bookDTO *dto.BookCreateDTO) error {
	return b.bookRepository.CreateBook(bookDTO)
}

func (b *BookService) FetchBooksAsync() ([]dto.Book, error) {

	books, err := b.bookRepository.FetchBooksAsync()
	if err != nil {
		return nil, err
	}

	result := make([]dto.Book, len(*books))
	// map result
	for index, val := range *books {
		result[index] = dto.Book{
			ID:     val.ID,
			Title:  val.Title,
			Author: val.Author,
		}
	}
	return result, nil

}

func (b *BookService) FetchBookByID(id int) (dto.Book, error) {
	book, err := b.bookRepository.FetchBookByID(id)
	if err != nil {
		return dto.Book{}, err
	}
	return dto.Book{
		ID:     book.ID,
		Title:  book.Title,
		Author: book.Author,
	}, nil
}

func (b *BookService) UpdateBook(bookDTO *dto.BookUpdateDTO) error {

	return b.bookRepository.UpdateBook(bookDTO)
}

func (b *BookService) DeleteBook(id int) error {

	return b.bookRepository.DeleteBook(id)
}
