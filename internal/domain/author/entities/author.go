package entities

import "github.com/VikashChauhan51/go-sample-api/internal/domain/author/objectValues"

type Author struct {
	Id        objectValues.AuthorId `gorm:"primary_key"`
	FirstName objectValues.FirstName
	LastName  objectValues.LastName
	Email     objectValues.Email
	Status    bool
}
