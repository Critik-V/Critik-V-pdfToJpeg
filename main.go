package main

import (
	"go-pdf2jpeg/handlers"
	"go-pdf2jpeg/utils"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const envFile = ".env" // Name of the .env file

func init() {
	// Load .env file
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Set Gin mode to release if the application is running in production mode
	if prodMode := utils.IsProduction(); prodMode {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create server instance
	server := gin.Default()

	// CORS middleware configuration to allow only POST requests from the specified origin
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{utils.GetCorsOrigin()},
		AllowMethods:     []string{"POST"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
	}))

	// Routes
	server.POST("/convert", handlers.POSTConvertPdf)

	// Run server
	server.Run(utils.GetPort())
}
