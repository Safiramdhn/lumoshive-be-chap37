package config

import (
	"api-service-shipping/database"
	"log"

	"github.com/spf13/viper"
)

// InitConfig initializes and reads configuration using Viper
func LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	database.InitDB()
}
