package routes

import (
	"github.com/VikashChauhan51/go-sample-api/internal/controllers"
	"github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/repositories"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/services"
)

func GetBookRoutes(db interfaces.Database) []Route {
	bookRepository := repositories.NewBookRepository(db)
	bookService := services.NewBookService(bookRepository)
	bookController := controllers.NewBookController(bookService)
	// Defind all routes
	return []Route{
		{"GET", "/books", bookController.GetBooks},
		{"GET", "/book/:id", bookController.GetBookByID},
		{"POST", "/books", bookController.CreateBook},
		{"PUT", "/books/:id", bookController.UpdateBook},
		{"DELETE", "/book/:id", bookController.DeleteBook},
	}
}
