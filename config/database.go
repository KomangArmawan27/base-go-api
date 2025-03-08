package config

import (
	"fmt"
	"log"

	// "os"

	"go-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the PostgreSQL connection
func ConnectDatabase() {
	// Load environment variables
	LoadEnv()

	// Get database credentials from .env
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		GetEnv("DB_HOST"),
		GetEnv("DB_USER"),
		GetEnv("DB_PASSWORD"),
		GetEnv("DB_NAME"),
		GetEnv("DB_PORT"),
		GetEnv("DB_SSLMODE"),
	)

	// Open connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	// Auto Migrate Models
	DB.AutoMigrate(&models.User{})

	fmt.Println("✅ Successfully connected to the database!")
}
