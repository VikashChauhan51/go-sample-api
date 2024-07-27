package main

import (
	"fmt"
	"time"

	"github.com/VikashChauhan51/go-sample-api/cmd/api/routes"
	docs "github.com/VikashChauhan51/go-sample-api/docs"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/databases"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Connect to the database
	dsn := "sqlserver://username:password@localhost:1433?database=bookstore"
	db, err := databases.NewDBConnection(dsn)
	if err != nil {
		fmt.Printf("failed to connect to database: %v \n", err)
	}
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
		allRoutes, err := routes.ScanRoutes(db)
		if err != nil {
			fmt.Printf("failed to load routes: %v\n", err)
			return
		}
		routes.RegisterRoutes(v1, allRoutes)
	}

	r.Run()
}
