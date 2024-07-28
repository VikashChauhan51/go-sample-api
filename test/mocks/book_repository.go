package mocks

import (
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
	"github.com/stretchr/testify/mock"
)

type BookRepository struct {
	mock.Mock
}

func (m *BookRepository) CreateBook(bookDTO *dto.BookCreateDTO) error {
	args := m.Called(bookDTO)
	return args.Error(0)
}

func (m *BookRepository) FetchBooksAsync() (*[]entities.Book, error) {
	args := m.Called()
	return args.Get(0).(*[]entities.Book), args.Error(1)
}

func (m *BookRepository) FetchBookByID(id int) (*entities.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *BookRepository) UpdateBook(bookDTO *dto.BookUpdateDTO) error {
	args := m.Called(bookDTO)
	return args.Error(0)
}

func (m *BookRepository) DeleteBook(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
