package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnv retrieves an environment variable
func GetEnv(key string) string {
	return os.Getenv(key)
}
