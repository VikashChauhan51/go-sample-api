package entities

type Book struct {
	ID     uint `gorm:"primary_key"`
	Title  string
	Author string
}
