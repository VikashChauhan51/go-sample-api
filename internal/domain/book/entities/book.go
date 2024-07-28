package entities

import (
	"github.com/VikashChauhan51/go-sample-api/internal/domain/author/entities"
	authorValues "github.com/VikashChauhan51/go-sample-api/internal/domain/author/objectValues"
	"github.com/VikashChauhan51/go-sample-api/internal/domain/book/objectValues"
)

type Book struct {
	Id            objectValues.BookId   `gorm:"primary_key"`
	Title         objectValues.Title    `gorm:"not null"`
	AuthorId      authorValues.AuthorId `gorm:"not null"`                          // Foreign key reference
	Author        entities.Author       `gorm:"foreignKey:AuthorID;references:ID"` // Association
	Description   string
	PublishedDate objectValues.PublishedDate `gorm:"not null"`
}
