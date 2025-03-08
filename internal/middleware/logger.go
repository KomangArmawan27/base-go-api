package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggerMiddleware logs request details
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Log request details
		log.Printf(
			"%s | %s | %s | %d | %v",
			c.Request.Method,
			c.Request.RequestURI,
			c.ClientIP(),
			c.Writer.Status(),
			duration,
		)
	}
}
