package services

import (
	"testing"

	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	svc "github.com/VikashChauhan51/go-sample-api/internal/infra/services"
	"github.com/VikashChauhan51/go-sample-api/test/mocks"

	"github.com/stretchr/testify/assert"
)

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
