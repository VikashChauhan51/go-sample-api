package routes

import (
	"net/http"

	"github.com/VikashChauhan51/go-sample-api/internal/controllers"
	"github.com/VikashChauhan51/go-sample-api/internal/core/interfaces"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/repositories"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/services"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func GetRoutes(db interfaces.Database) *[]Route {
	// Defind all routes
	return &[]Route{
		{"GET", "/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		}},
		{"GET", "/books", controllers.NewBookController(services.NewBookService(repositories.NewBookRepository(db))).GetBooks},
	}
}
