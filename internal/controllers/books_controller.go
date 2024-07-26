package controllers

import (
	"net/http"

	svc "github.com/VikashChauhan51/go-sample-api/internal/core/interfaces/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService svc.BookService
}

func NewBookController(service svc.BookService) *BookController {
	return &BookController{service}
}

// @Summary Get books
// @Description Retrieves a list of books
// @Tags books
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func (b *BookController) GetBooks(c *gin.Context) {
	books, err := b.bookService.FetchBooksAsync()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
