package middleware

import (
	"go-api/internal/auth"
	"go-api/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks if a request has a valid JWT
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			utils.Response(c, http.StatusUnauthorized, false, "Missing or invalid token", nil)
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims, err := auth.ValidateToken(tokenString)
		if err != nil {
			utils.Response(c, http.StatusUnauthorized, false, "Invalid token", nil)
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("role", claims.Role)
		c.Next()
	}
}
