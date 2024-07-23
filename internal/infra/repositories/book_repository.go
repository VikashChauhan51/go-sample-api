package repositories

import (
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	interfaces "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces/repositories"
	_ "gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) FetchBooksAsync() (*[]entities.Book, error) {
	var books []entities.Book
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return &books, nil
}
