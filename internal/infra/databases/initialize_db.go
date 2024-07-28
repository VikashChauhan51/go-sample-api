package databases

import (
	authorEntity "github.com/VikashChauhan51/go-sample-api/internal/domain/author/entities"
	bookEntity "github.com/VikashChauhan51/go-sample-api/internal/domain/book/entities"
	"gorm.io/gorm"
)

func Initialize(db *gorm.DB) error {
	err := db.AutoMigrate(&bookEntity.Book{}, &authorEntity.Author{})
	return err
}
