package services

import (
	"github.com/VikashChauhan51/go-sample-api/models"
	"time"
)

// Simulate fetching books from an external API
func FetchBooksAsync() ([]models.Book, error) {

	r := make(chan []models.Book)
	go func() {
		time.Sleep(2 * time.Second) // Simulate external API call
		books := []models.Book{
			{ID: 1, Title: "The Lord of the Rings", Author: "J. R. R. Tolkien"},
			{ID: 2, Title: "Pride and Prejudice", Author: "Jane Austen"},
			{ID: 3, Title: "To Kill a Mockingbird", Author: "Harper Lee"},
		}
		r <- books // Send book list to channel

	}()
	result := <-r
	return result, nil

}
