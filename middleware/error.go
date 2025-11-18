package middleware

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
  return func(c *gin.Context) {
    defer func() {
      if err := recover(); err != nil {
        c.JSON(500, gin.H{"error": "Internal server error"})
      }
    }()
    c.Next()
  }
}