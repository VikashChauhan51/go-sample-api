package dto

type Book struct {
	ID     uint   `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
}

type BookCreateDTO struct {
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
}

type BookUpdateDTO struct {
	ID     int    `json:"id" xml:"id"`
	Title  string `json:"title" xml:"title"`
	Author string `json:"author" xml:"author"`
}
