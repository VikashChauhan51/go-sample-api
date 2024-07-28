package repositories

import (
	"testing"

	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
	repo "github.com/VikashChauhan51/go-sample-api/internal/infra/repositories"
	"github.com/VikashChauhan51/go-sample-api/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateBook(t *testing.T) {
	mockDB := new(mocks.MockDB)
	bookRepository := repo.NewBookRepository(mockDB)

	bookDTO := &dto.BookCreateDTO{
		Title:  "Test Title",
		Author: "Test Author",
	}

	mockDB.On("Create", mock.Anything).Return(&gorm.DB{Error: nil}).Once()

	err := bookRepository.CreateBook(bookDTO)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

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

func TestFetchBookByID(t *testing.T) {
	mockDB := new(mocks.MockDB)
	bookRepository := repo.NewBookRepository(mockDB)

	expectedBook := &entities.Book{ID: 1, Title: "Title1", Author: "Author1"}

	mockDB.On("First", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.Book)
		*arg = *expectedBook
	}).Return(&gorm.DB{Error: nil}).Once()

	result, err := bookRepository.FetchBookByID(1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, expectedBook, result)

	mockDB.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	mockDB := new(mocks.MockDB)
	bookRepository := repo.NewBookRepository(mockDB)

	bookDTO := &dto.BookUpdateDTO{
		ID:     1,
		Title:  "Updated Title",
		Author: "Updated Author",
	}

	mockDB.On("First", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*entities.Book)
		*arg = entities.Book{ID: 1, Title: "Title1", Author: "Author1"}
	}).Return(&gorm.DB{Error: nil}).Once()

	mockDB.On("Save", mock.Anything).Return(&gorm.DB{Error: nil}).Once()

	err := bookRepository.UpdateBook(bookDTO)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	mockDB := new(mocks.MockDB)
	bookRepository := repo.NewBookRepository(mockDB)

	mockDB.On("Delete", mock.Anything, mock.Anything).Return(&gorm.DB{Error: nil}).Once()

	err := bookRepository.DeleteBook(1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
