package entities

import "github.com/VikashChauhan51/go-sample-api/internal/domain/author/objectValues"

type Author struct {
	Id        objectValues.AuthorId  `gorm:"type:varchar(255);primary_key"`
	FirstName objectValues.FirstName `gorm:"type:varchar(255);not null"`
	LastName  objectValues.LastName  `gorm:"type:varchar(255);not null"`
	Email     objectValues.Email     `gorm:"type:varchar(255);not null;unique"`
	Status    bool
}
