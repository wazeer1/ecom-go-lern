package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminMiddleware checks if user is admin
func AdminMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        isAdmin, exists := c.Get("is_admin")

        if !exists {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required"})
            c.Abort()
            return
        }

        if !isAdmin.(bool) {
            c.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            c.Abort()
            return
        }

        c.Next()
    }
}