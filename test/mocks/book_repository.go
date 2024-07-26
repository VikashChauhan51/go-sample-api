package mocks

import (
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/stretchr/testify/mock"
)

type BookRepository struct {
	mock.Mock
}

func (m *BookRepository) FetchBooksAsync() (*[]entities.Book, error) {
	args := m.Called()
	return args.Get(0).(*[]entities.Book), args.Error(1)
}
