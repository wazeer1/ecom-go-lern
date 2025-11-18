package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourusername/ecommerce-api/models"
)

type CreateProductRequest struct {
    Name        string  `json:"name" binding:"required,min=3"`
    Description string  `json:"description" binding:"required"`
    Price       float64 `json:"price" binding:"required,gt=0"`
    Stock       int     `json:"stock" binding:"required,gte=0"`
    Category    string  `json:"category" binding:"required"`
    SKU         string  `json:"sku" binding:"required"`
    Discount    float64 `json:"discount" binding:"gte=0,lte=100"`
    Image       string  `json:"image"`
}

// GetProducts retrieves all products with pagination
func GetProducts(c *gin.Context, db *gorm.DB) {
    var products []models.Product

    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    category := c.Query("category")
    search := c.Query("search")

    offset := (page - 1) * limit

    query := db

    if category != "" {
        query = query.Where("category = ?", category)
    }

    if search != "" {
        query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
    }

    var total int64
    query.Model(&models.Product{}).Count(&total)

    if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  products,
        "page":  page,
        "limit": limit,
        "total": total,
    })
}

// GetProductByID retrieves single product
func GetProductByID(c *gin.Context, db *gorm.DB) {
    id := c.Param("id")

    var product models.Product
    if err := db.Preload("Reviews").First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// CreateProduct creates new product (admin only)
func CreateProduct(c *gin.Context, db *gorm.DB) {
    var req CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product := models.Product{
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        Stock:       req.Stock,
        Category:    req.Category,
        SKU:         req.SKU,
        Discount:    req.Discount,
        Image:       req.Image,
    }

    if err := db.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusCreated, product)
}

// UpdateProduct updates product (admin only)
func UpdateProduct(c *gin.Context, db *gorm.DB) {
    id := c.Param("id")

    var product models.Product
    if err := db.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    var req CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.Model(&product).Updates(req).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// DeleteProduct deletes product (admin only)
func DeleteProduct(c *gin.Context, db *gorm.DB) {
    id := c.Param("id")

    if err := db.Delete(&models.Product{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}