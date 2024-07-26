package controllers

import (
	"net/http"

	"github.com/VikashChauhan51/go-sample-api/internal/infra/services"
	"github.com/gin-gonic/gin"
)

// @Summary Get books
// @Description Retrieves a list of books
// @Tags books
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c *gin.Context) {
	books, err := services.BookService.FetchBooksAsync()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": books})
}
