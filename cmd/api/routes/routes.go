package routes

import (
	"net/http"

	"github.com/VikashChauhan51/go-sample-api/internal/controllers"
	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// Defind all routes
var Routes = []Route{
	{"GET", "/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	}},
	{"GET", "/books", controllers.GetBooks},
}
