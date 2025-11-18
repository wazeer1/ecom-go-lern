package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourusername/ecommerce-api/models"
)

type CreateOrderRequest struct {
    Items           []OrderItemRequest `json:"items" binding:"required,min=1"`
    ShippingAddress string             `json:"shipping_address" binding:"required"`
}

type OrderItemRequest struct {
    ProductID uint `json:"product_id" binding:"required"`
    Quantity  int  `json:"quantity" binding:"required,min=1"`
}

// CreateOrder creates new order
func CreateOrder(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")

    var req CreateOrderRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var totalAmount float64
    var orderItems []models.OrderItem

    for _, item := range req.Items {
        var product models.Product
        if err := db.First(&product, item.ProductID).Error; err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
            return
        }

        if product.Stock < item.Quantity {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient stock"})
            return
        }

        unitPrice := product.Price * (1 - product.Discount/100)
        totalAmount += unitPrice * float64(item.Quantity)

        orderItem := models.OrderItem{
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            Price: unitPrice,
            Discount:  product.Discount,
        }
        orderItems = append(orderItems, orderItem)
    }

    orderNumber := fmt.Sprintf("ORD-%d-%d", userID, 12345)

    order := models.Order{
        UserID:          userID,
        OrderNumber:     orderNumber,
        TotalAmount:     totalAmount,
        Status:          "pending",
        ShippingAddr:    req.ShippingAddress,
        OrderItems:      orderItems,
    }

    if err := db.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusCreated, order)
}

// GetUserOrders retrieves user orders
func GetUserOrders(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    offset := (page - 1) * limit

    var orders []models.Order
    if err := db.Where("user_id = ?", userID).Offset(offset).Limit(limit).
        Preload("OrderItems").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": orders, "page": page, "limit": limit})
}

// GetOrderByID retrieves single order
func GetOrderByID(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")
    orderID := c.Param("id")

    var order models.Order
    if err := db.Where("id = ? AND user_id = ?", orderID, userID).
        Preload("OrderItems").First(&order).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

// CancelOrder cancels order
func CancelOrder(c *gin.Context, db *gorm.DB) {
    userID := c.GetUint("user_id")
    orderID := c.Param("id")

    var order models.Order
    if err := db.First(&order, orderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if order.UserID != userID {
        c.JSON(http.StatusForbidden, gin.H{"error": "Cannot cancel this order"})
        return
    }

    if err := db.Model(&order).Update("status", "cancelled").Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order cancelled"})
}

// GetAllOrders retrieves all orders (admin)
func GetAllOrders(c *gin.Context, db *gorm.DB) {
    var orders []models.Order
    if err := db.Preload("User").Preload("OrderItems").Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": orders})
}

// UpdateOrderStatus updates order status (admin)
func UpdateOrderStatus(c *gin.Context, db *gorm.DB) {
    id := c.Param("id")

    type StatusRequest struct {
        Status string `json:"status" binding:"required"`
    }

    var req StatusRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Model(&models.Order{}).Where("id = ?", id).Update("status", req.Status).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
}