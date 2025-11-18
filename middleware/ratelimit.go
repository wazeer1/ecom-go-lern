package middleware

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
  limiters = make(map[string]*rate.Limiter)
  mu sync.Mutex
)

func RateLimitMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    ip := c.ClientIP()
    mu.Lock()
    
    if _, exists := limiters[ip]; !exists {
      limiters[ip] = rate.NewLimiter(10, 20)
    }
    limiter := limiters[ip]
    mu.Unlock()
    
    if !limiter.Allow() {
      c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
      c.Abort()
      return
    }
    c.Next()
  }
}