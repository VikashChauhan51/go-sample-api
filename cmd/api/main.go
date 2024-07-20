package main

import (
	"fmt"
	"net/http"
	"time"

	docs "github.com/VikashChauhan51/go-sample-api/docs"
	"github.com/VikashChauhan51/go-sample-api/internal/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = "/api/v1"
	// Mount the Swagger middleware
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		for _, route := range routes.Routes {
			switch route.Method {
			case http.MethodGet:
				v1.GET(route.Path, route.Handler)
			case http.MethodPost:
				v1.POST(route.Path, route.Handler)
			case http.MethodPut:
				v1.PUT(route.Path, route.Handler)
			case http.MethodPatch:
				v1.PATCH(route.Path, route.Handler)
			case http.MethodDelete:
				v1.DELETE(route.Path, route.Handler)
			case http.MethodOptions:
				v1.OPTIONS(route.Path, route.Handler)
			case http.MethodHead:
				v1.HEAD(route.Path, route.Handler)
			default:
				fmt.Printf("Warning: Unsupported HTTP verb '%s' in route %s\n", route.Method, route.Path)
			}
		}
	}

	r.Run()
}
