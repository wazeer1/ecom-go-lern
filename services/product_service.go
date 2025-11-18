package services

import (
	"gorm.io/gorm"

	"github.com/wazeer1/ecommerce-api/models"
)

// Product listing with pagination and optional filters
func ListProducts(db *gorm.DB, page int, limit int, category string, search string) ([]models.Product, int64, error) {
    var products []models.Product
    var total int64

    offset := (page - 1) * limit
    query := db.Model(&models.Product{})

    if category != "" {
        query = query.Where("category = ?", category)
    }
    if search != "" {
        // Using LIKE for cross-db compatibility
        like := "%" + search + "%"
        query = query.Where("name LIKE ? OR description LIKE ?", like, like)
    }

    if err := query.Count(&total).Error; err != nil {
        return nil, 0, err
    }
    if err := query.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
        return nil, 0, err
    }
    return products, total, nil
}

func GetProductByID(db *gorm.DB, id string) (*models.Product, error) {
    var product models.Product
    if err := db.Preload("Reviews").First(&product, id).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

// CreateProductInput represents fields allowed for creation
type CreateProductInput struct {
    Name        string
    Description string
    Price       float64
    Stock       int
    Category    string
    SKU         string
    Discount    float64
    Image       string
}

func CreateProduct(db *gorm.DB, in CreateProductInput) (*models.Product, error) {
    product := models.Product{
        Name:        in.Name,
        Description: in.Description,
        Price:       in.Price,
        Stock:       in.Stock,
        Category:    in.Category,
        SKU:         in.SKU,
        Discount:    in.Discount,
        Image:       in.Image,
    }
    if err := db.Create(&product).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

// UpdateProductInput mirrors CreateProductInput for simplicity
type UpdateProductInput struct {
    Name        string
    Description string
    Price       float64
    Stock       int
    Category    string
    SKU         string
    Discount    float64
    Image       string
}

func UpdateProduct(db *gorm.DB, id string, in UpdateProductInput) (*models.Product, error) {
    var product models.Product
    if err := db.First(&product, id).Error; err != nil {
        return nil, err
    }
    // Use map to avoid zero-values overwriting when fields are omitted in handler mapping
    updates := map[string]interface{}{
        "name":        in.Name,
        "description": in.Description,
        "price":       in.Price,
        "stock":       in.Stock,
        "category":    in.Category,
        "sku":         in.SKU,
        "discount":    in.Discount,
        "image":       in.Image,
    }
    if err := db.Model(&product).Updates(updates).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func DeleteProduct(db *gorm.DB, id string) error {
    return db.Delete(&models.Product{}, id).Error
}