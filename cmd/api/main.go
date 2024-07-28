package main

import (
	"fmt"
	"os"
	"time"

	"github.com/VikashChauhan51/go-sample-api/cmd/api/routes"
	docs "github.com/VikashChauhan51/go-sample-api/docs"
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/databases"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("No .env file found")
	}
}

func main() {

	loadEnv()
	// Connect to the database
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := databases.NewDBConnection(dsn)
	if err != nil {
		fmt.Printf("failed to connect to database: %v \n", err)
	}

	// Auto migrate database (create tables)
	err = db.AutoMigrate(&entities.Book{})
	if err != nil {
		fmt.Printf("failed to migrate database: %v", err)
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
		allRoutes := routes.ScanRoutes(db)
		routes.RegisterRoutes(v1, allRoutes)
	}

	r.Run()
}
