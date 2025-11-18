package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/wazeer1/ecommerce-api/config"
	"github.com/wazeer1/ecommerce-api/database"
	"github.com/wazeer1/ecommerce-api/middleware"
	"github.com/wazeer1/ecommerce-api/routes"
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