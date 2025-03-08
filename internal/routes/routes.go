package routes

import (
	"go-api/internal/handlers"
	"go-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes initializes API routes
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Apply logger middleware
	r.Use(middleware.LoggerMiddleware())

	// Public route (Login)
	r.POST("/api/v1/login", handlers.LoginHandler)

	// Protected routes
	protected := r.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware())

	// Admin routes (only admin can modify data)
	admin := protected.Group("/")
	admin.Use(middleware.RoleMiddleware("admin"))

	// Manager routes (only manager or above can modify data)
	manager := protected.Group("/")
	manager.Use(middleware.RoleMiddleware("manager"))

	// User routes
	admin.GET("/users", handlers.GetUsers)
	admin.GET("/user", handlers.GetUserByID)
	admin.POST("/user/create", handlers.CreateUser)
	admin.PUT("/user/update", handlers.UpdateUser)
	admin.DELETE("/user/delete", handlers.DeleteUser)

	return r
}
