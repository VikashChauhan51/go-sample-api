package responses

import (
	"time"
)

type BookResponse struct {
	Id            uint      `json:"id" xml:"id"`
	Title         string    `json:"title" xml:"title"`
	Author        string    `json:"author" xml:"author"`
	Description   string    `json:"description" xml:"description"`
	PublishedDate time.Time `json:"published_date" xml:"published_date"`
}
