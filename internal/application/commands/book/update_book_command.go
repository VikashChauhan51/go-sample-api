package commands

type UpdateBookCommand struct {
	ID          string `json:"id" xml:"id"`
	Title       string `json:"title" xml:"title"`
	Author      string `json:"author" xml:"author"`
	Description string `json:"description" xml:"description"`
}
