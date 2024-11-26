package main

import (
	"api-service-shipping/config"
	"api-service-shipping/controller"
	"api-service-shipping/database"
	"api-service-shipping/router"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	// Load the configuration
	config.LoadConfig()
}

func main() {
	// Initialize Gin engine
	r := gin.Default()

	// Initialize the database connection
	db := database.GetDatabase()

	// Initialize controllers
	mainController := controller.NewMainController(db)

	// Setup the router with the controllers
	router.APIRouter(r, mainController)

	// Start the server on the configured port
	port := viper.GetString("PORT")
	if port == "" {
		log.Fatal("PORT is not defined in the configuration")
	}
	log.Printf("Server running on port %s", port)
	r.Run(":" + port)
}
