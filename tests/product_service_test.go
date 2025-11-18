package tests

import (
	"fmt"
	"strconv"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/wazeer1/ecommerce-api/models"
	"github.com/wazeer1/ecommerce-api/services"
)

func setupTestDB(t *testing.T) *gorm.DB {
    t.Helper()
    db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    if err != nil {
        t.Fatalf("failed to open sqlite: %v", err)
    }
    if err := db.AutoMigrate(&models.Product{}, &models.Review{}); err != nil {
        t.Fatalf("failed to migrate: %v", err)
    }
    return db
}

func TestCreateProduct(t *testing.T) {
    db := setupTestDB(t)

    in := services.CreateProductInput{
        Name:        "Test Product",
        Description: "A product for testing",
        Price:       19.99,
        Stock:       10,
        Category:    "test",
        SKU:         "TEST-001",
        Discount:    5,
        Image:       "",
    }

    p, err := services.CreateProduct(db, in)
    if err != nil {
        t.Fatalf("CreateProduct error: %v", err)
    }
    if p.ID == 0 {
        t.Fatalf("expected persisted product with ID, got 0")
    }
}

func TestListProducts(t *testing.T) {
    db := setupTestDB(t)

    // Seed
    for i := 1; i <= 15; i++ {
        _, err := services.CreateProduct(db, services.CreateProductInput{
            Name:        fmt.Sprintf("Prod-%02d", i),
            Description: "desc",
            Price:       float64(i),
            Stock:       i,
            Category:    "cat",
            SKU:         fmt.Sprintf("SKU-%02d", i),
        })
        if err != nil {
            t.Fatalf("seed create error: %v", err)
        }
    }

    products, total, err := services.ListProducts(db, 1, 10, "", "")
    if err != nil {
        t.Fatalf("ListProducts error: %v", err)
    }
    if total != 15 {
        t.Fatalf("expected total 15, got %d", total)
    }
    if len(products) != 10 {
        t.Fatalf("expected 10 products on page 1, got %d", len(products))
    }
}

func TestGetUpdateDeleteProduct(t *testing.T) {
    db := setupTestDB(t)

    created, err := services.CreateProduct(db, services.CreateProductInput{
        Name:     "Prod",
        Price:    9.99,
        Stock:    3,
        Category: "cat",
        SKU:      "SKU-X",
    })
    if err != nil {
        t.Fatalf("CreateProduct error: %v", err)
    }

    // Get
    got, err := services.GetProductByID(db, strconv.FormatUint(uint64(created.ID), 10))
    if err != nil || got.ID != created.ID {
        t.Fatalf("GetProductByID failed: err=%v id=%d", err, got.ID)
    }

    // Update
    up, err := services.UpdateProduct(db, strconv.FormatUint(uint64(created.ID), 10), services.UpdateProductInput{Price: 12.34, Stock: 5})
    if err != nil {
        t.Fatalf("UpdateProduct failed: %v", err)
    }
    if up.Price != 12.34 || up.Stock != 5 {
        t.Fatalf("UpdateProduct did not persist fields: price=%v stock=%v", up.Price, up.Stock)
    }

    // Delete
    if err := services.DeleteProduct(db, strconv.FormatUint(uint64(created.ID), 10)); err != nil {
        t.Fatalf("DeleteProduct failed: %v", err)
    }
    if _, err := services.GetProductByID(db, strconv.FormatUint(uint64(created.ID), 10)); err == nil {
        t.Fatalf("expected error when fetching deleted product")
    }
}