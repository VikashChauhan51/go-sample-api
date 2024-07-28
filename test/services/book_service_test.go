package services

import (
	"testing"

	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
	svc "github.com/VikashChauhan51/go-sample-api/internal/infra/services"
	"github.com/VikashChauhan51/go-sample-api/test/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCreateBook(t *testing.T) {
	bookRepository := new(mocks.BookRepository)
	bookService := svc.NewBookService(bookRepository)

	bookDTO := &dto.BookCreateDTO{
		Title:  "Test Title",
		Author: "Test Author",
	}

	bookRepository.On("CreateBook", bookDTO).Return(nil)

	err := bookService.CreateBook(bookDTO)

	assert.NoError(t, err)
	bookRepository.AssertExpectations(t)
}

func TestFetchBooksAsync(t *testing.T) {
	bookRepository := new(mocks.BookRepository)
	bookService := svc.NewBookService(bookRepository)

	expectedBooks := []entities.Book{
		{ID: 1, Title: "Title1", Author: "Author1"},
		{ID: 2, Title: "Title2", Author: "Author2"},
	}

	bookRepository.On("FetchBooksAsync").Return(&expectedBooks, nil)
	result, err := bookService.FetchBooksAsync()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(expectedBooks), len(result))

	bookRepository.AssertExpectations(t)
}

func TestFetchBookByID(t *testing.T) {
	bookRepository := new(mocks.BookRepository)
	bookService := svc.NewBookService(bookRepository)

	expectedBook := &entities.Book{ID: 1, Title: "Title1", Author: "Author1"}

	bookRepository.On("FetchBookByID", 1).Return(expectedBook, nil)

	result, err := bookService.FetchBookByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBook.Title, result.Title)

	bookRepository.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	bookRepository := new(mocks.BookRepository)
	bookService := svc.NewBookService(bookRepository)

	bookDTO := &dto.BookUpdateDTO{
		ID:     1,
		Title:  "Updated Title",
		Author: "Updated Author",
	}

	bookRepository.On("UpdateBook", bookDTO).Return(nil)

	err := bookService.UpdateBook(bookDTO)

	assert.NoError(t, err)
	bookRepository.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	bookRepository := new(mocks.BookRepository)
	bookService := svc.NewBookService(bookRepository)

	bookRepository.On("DeleteBook", 1).Return(nil)

	err := bookService.DeleteBook(1)

	assert.NoError(t, err)
	bookRepository.AssertExpectations(t)
}
