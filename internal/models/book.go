package models

type Book struct {
	ID     uint   `json:"id" xml:"id" gorm:"primary_key"`
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
}
