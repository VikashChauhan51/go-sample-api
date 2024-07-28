package entities

import (
	"github.com/VikashChauhan51/go-sample-api/internal/domain/author/entities"
	authorValues "github.com/VikashChauhan51/go-sample-api/internal/domain/author/objectValues"
	"github.com/VikashChauhan51/go-sample-api/internal/domain/book/objectValues"
)

type Book struct {
	Id            objectValues.BookId        `gorm:"primary_key"`
	Title         objectValues.Title         `gorm:"type:varchar(255);not null"`
	AuthorId      authorValues.AuthorId      `gorm:"type:varchar(255);not null"`        // Foreign key reference
	Author        entities.Author            `gorm:"foreignKey:AuthorId;references:Id"` // Association
	Description   string                     `gorm:"type:varchar(255); null"`
	PublishedDate objectValues.PublishedDate `gorm:"type:date;not null"`
}
