package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yourusername/ecommerce-api/config"
	"github.com/yourusername/ecommerce-api/database"
	"github.com/yourusername/ecommerce-api/middleware"
	"github.com/yourusername/ecommerce-api/routes"
)

func main() {
  godotenv.Load()
  
  db := config.InitDB()
  if err := database.RunMigrations(db); err != nil {
    panic(err)
  }
  
  router := gin.Default()
  
  router.Use(middleware.CORSMiddleware())
  router.Use(middleware.ErrorHandling())
  router.Use(middleware.RateLimitMiddleware())
  
  routes.SetupRoutes(router, db)
  
  fmt.Println("Server running on :8080")
  router.Run(":8080")
}