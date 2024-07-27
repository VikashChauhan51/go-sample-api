package routes

import (
	"net/http"

	"github.com/VikashChauhan51/go-sample-api/internal/controllers"
	"github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/repositories"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/services"
	"github.com/gin-gonic/gin"
)

func GetBookRoutes(db interfaces.Database) []Route {
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)
	// Defind all routes
	return []Route{
		{"GET", "/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		}},
		{"GET", "/books", bookController.GetBooks},
	}
}
