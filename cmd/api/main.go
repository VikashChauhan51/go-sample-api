package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/VikashChauhan51/go-sample-api/cmd/api/routes"
	config "github.com/VikashChauhan51/go-sample-api/configs"
	docs "github.com/VikashChauhan51/go-sample-api/docs"
	"github.com/VikashChauhan51/go-sample-api/internal/core/entities"
	"github.com/VikashChauhan51/go-sample-api/internal/infra/databases"
	"github.com/VikashChauhan51/go-sample-api/pkg/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	env := os.Getenv("GO_ENV")
	config.LoadConfig(env) // Load environment-specific config
	// Connect to the database
	dbConfig := config.Config.Database

	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)
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
	// Apply middlewares
	r.Use(middlewares.LoggerMiddleware())
	docs.SwaggerInfo.BasePath = "/api/v1"
	// Mount the Swagger middleware
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v1 := r.Group("/api/v1")
	{
		allRoutes := routes.ScanRoutes(db)
		routes.RegisterRoutes(v1, allRoutes)
	}
	port := config.Config.Server.Port
	address := fmt.Sprintf(":%d", port)
	fmt.Printf("Starting server on port %d\n", port)
	// Start the server on the configured port
	if err := r.Run(address); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
