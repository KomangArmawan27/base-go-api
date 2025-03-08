package middleware

import (
	"go-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware is a middleware that checks if the user has the required role
type roles struct {
	RequiredRole string `json:"requiredRole"`
	UserRole     string `json:"userRole"`
}

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		// Debug log to see if the role is present
		if !exists {
			// c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Missing role", "role": role})
			utils.Response(c, http.StatusForbidden, false, "Forbidden: Missing role", nil)
			c.Abort()
			return
		}

		// Type assertion to make sure the role is a string
		userRole, ok := role.(string)
		if !ok {
			utils.Response(c, http.StatusForbidden, false, "Forbidden: Invalid role type", nil)
			c.Abort()
			return
		}

		// build response
		roles := roles{
			RequiredRole: requiredRole,
			UserRole:     userRole,
		}
		if userRole != requiredRole {
			utils.Response(c, http.StatusForbidden, false, "Forbidden: Insufficient permissions", roles)
			c.Abort()
			return
		}

		c.Next()
	}
}
