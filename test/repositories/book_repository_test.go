package repositories

import (
	"testing"

	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	repo "github.com/VikashChauhan51/go-sample-api/internal/infra/repositories"
	"github.com/VikashChauhan51/go-sample-api/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestFetchBooksAsync(t *testing.T) {
	mockDB := new(mocks.MockDB)
	bookRepository := repo.NewBookRepository(mockDB)

	expectedBooks := []entities.Book{
		{ID: 1, Title: "Title1", Author: "Author1"},
		{ID: 2, Title: "Title2", Author: "Author2"},
	}

	mockDB.On("Find", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]entities.Book)
		*arg = expectedBooks
	}).Return(&gorm.DB{Error: nil}).Once()

	result, err := bookRepository.FetchBooksAsync()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(expectedBooks), len(*result))

	mockDB.AssertExpectations(t)
}
