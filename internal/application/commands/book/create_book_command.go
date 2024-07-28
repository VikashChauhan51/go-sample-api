package commands

type CreateBookCommand struct {
	Title       string `json:"title" xml:"title"`
	Author      string `json:"author" xml:"author"`
	Description string `json:"description" xml:"description"`
}
