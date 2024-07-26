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
