package main

import (
	"fmt"
	"go-api/config"
	"go-api/internal/routes"
)

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	// Start the Gin server
	r := routes.SetupRoutes()
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	fmt.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}
