package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/wazeer1/ecommerce-api/handlers"
	"github.com/wazeer1/ecommerce-api/middleware"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ok"})
    })
    
    api := router.Group("/api")
    
    // Public routes
    auth := api.Group("/auth")
    {
        auth.POST("/register", func(c *gin.Context) {
            handlers.Register(c, db)
        })
        auth.POST("/login", func(c *gin.Context) {
            handlers.Login(c, db)
        })
    }
    
    // Public product routes
    products := api.Group("/products")
    {
        products.GET("", func(c *gin.Context) {
            handlers.GetProducts(c, db)
        })
        products.GET("/:id", func(c *gin.Context) {
            handlers.GetProductByID(c, db)  // ✅ FIXED
        })
    }
    
    // Protected routes
    protected := api.Group("")
    protected.Use(middleware.AuthMiddleware())
    {
        // Orders
        protected.POST("/orders", func(c *gin.Context) {
            handlers.CreateOrder(c, db)
        })
        protected.GET("/orders", func(c *gin.Context) {
            handlers.GetUserOrders(c, db)  // ✅ FIXED
        })
        protected.GET("/orders/:id", func(c *gin.Context) {
            handlers.GetOrderByID(c, db)
        })
        
        // Cart
        protected.GET("/cart", func(c *gin.Context) {
            handlers.GetCart(c, db)
        })
        protected.POST("/cart/add", func(c *gin.Context) {
            handlers.AddToCart(c, db)
        })
        protected.DELETE("/cart/:product_id", func(c *gin.Context) {
            handlers.RemoveFromCart(c, db)
        })
    }
    
    // Admin routes
    admin := api.Group("/admin")
    admin.Use(middleware.AuthMiddleware())
    admin.Use(middleware.AdminMiddleware())  // ✅ FIXED
    {
        admin.POST("/products", func(c *gin.Context) {
            handlers.CreateProduct(c, db)
        })
        admin.PUT("/products/:id", func(c *gin.Context) {
            handlers.UpdateProduct(c, db)
        })
        admin.DELETE("/products/:id", func(c *gin.Context) {
            handlers.DeleteProduct(c, db)
        })
    }
}