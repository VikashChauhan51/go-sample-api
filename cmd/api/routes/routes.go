package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

func RegisterRoutes(r *gin.RouterGroup, routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			r.GET(route.Path, route.Handler)
		case http.MethodPost:
			r.POST(route.Path, route.Handler)
		case http.MethodPut:
			r.PUT(route.Path, route.Handler)
		case http.MethodPatch:
			r.PATCH(route.Path, route.Handler)
		case http.MethodDelete:
			r.DELETE(route.Path, route.Handler)
		case http.MethodOptions:
			r.OPTIONS(route.Path, route.Handler)
		case http.MethodHead:
			r.HEAD(route.Path, route.Handler)
		default:
			fmt.Printf("Warning: Unsupported HTTP verb '%s' in route %s\n", route.Method, route.Path)
		}
	}
}
