package dto

type Book struct {
	ID     uint   `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
}
