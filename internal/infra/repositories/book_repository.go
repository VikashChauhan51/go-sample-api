package repositories

import (
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	interfaces "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
	repo "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces/repositories"
	"github.com/VikashChauhan51/go-sample-api/internal/dto"
	_ "gorm.io/driver/sqlserver"
)

type BookRepository struct {
	db interfaces.Database
}

func NewBookRepository(db interfaces.Database) repo.BookRepository {
	return &BookRepository{db}
}

func (r *BookRepository) CreateBook(bookDTO *dto.BookCreateDTO) error {

	book := entities.Book{
		Title:  bookDTO.Title,
		Author: bookDTO.Author,
	}
	result := r.db.Create(&book)
	return result.Error
}

func (r *BookRepository) FetchBooksAsync() (*[]entities.Book, error) {
	var books []entities.Book
	result := r.db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return &books, nil
}

func (r *BookRepository) FetchBookByID(id int) (*entities.Book, error) {
	var book entities.Book
	result := r.db.First(&book, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &book, nil
}

func (r *BookRepository) UpdateBook(bookDTO *dto.BookUpdateDTO) error {
	var book entities.Book
	result := r.db.First(&book, bookDTO.ID)
	if result.Error != nil {
		return result.Error
	}
	book.Title = bookDTO.Title
	book.Author = bookDTO.Author
	return r.db.Save(&book).Error
}

func (r *BookRepository) DeleteBook(id int) error {
	result := r.db.Delete(&entities.Book{}, id)
	return result.Error
}
