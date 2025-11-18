package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourusername/ecommerce-api/models"
)

type AddToCartRequest struct {
    ProductID uint `json:"product_id" binding:"required"`
    Quantity  int  `json:"quantity" binding:"required,min=1"`
}

// GetCart retrieves user cart
func GetCart(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")

    var cartItems []models.CartItem
    if err := db.Where("user_id = ?", userID).Preload("Product").Find(&cartItems).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": cartItems})
}

// AddToCart adds product to cart
func AddToCart(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")

    var req AddToCartRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    cartItem := models.CartItem{
        UserID:    userID,
        ProductID: req.ProductID,
        Quantity:  req.Quantity,
    }

    if err := db.Create(&cartItem).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add to cart"})
        return
    }

    c.JSON(http.StatusCreated, cartItem)
}

// RemoveFromCart removes product from cart
func RemoveFromCart(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")
    productID := c.Param("product_id")

    if err := db.Where("user_id = ? AND product_id = ?", userID, productID).
        Delete(&models.CartItem{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to remove"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Item removed from cart"})
}