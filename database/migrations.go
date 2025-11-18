package database

import (
	"fmt"

	"github.com/yourusername/ecommerce-api/models"
	"gorm.io/gorm"
)

// RunMigrations runs all database migrations
func RunMigrations(db *gorm.DB) error {
    fmt.Println("Running database migrations...")

    // Auto migrate all models
    if err := db.AutoMigrate(
        &models.User{},
        &models.Product{},
        &models.Order{},
        &models.OrderItem{},
        &models.Payment{},
        &models.CartItem{},
        &models.Review{},
    ); err != nil {
        return fmt.Errorf("migration failed: %w", err)
    }

    fmt.Println("✅ Migrations completed successfully!")
    return nil
}

// DropAllTables drops all tables (use with caution)
func DropAllTables(db *gorm.DB) error {
    return db.Migrator().DropTable(
        &models.Review{},
        &models.CartItem{},
        &models.Payment{},
        &models.OrderItem{},
        &models.Order{},
        &models.Product{},
        &models.User{},
    )
}

// SeedDatabase seeds initial data
func SeedDatabase(db *gorm.DB) error {
    // Check if data already exists
    var count int64
    db.Model(&models.Product{}).Count(&count)
    if count > 0 {
        return nil
    }

    fmt.Println("Seeding database...")

    // Seed products
    products := []models.Product{
        {
            Name:        "Laptop",
            Description: "High performance laptop",
            Price:       999.99,
            Stock:       50,
            Category:    "Electronics",
            SKU:         "LAPTOP001",
            Discount:    10,
            Image:       "laptop.jpg",
            Status:      "active",
        },
        {
            Name:        "Mouse",
            Description: "Wireless mouse",
            Price:       29.99,
            Stock:       200,
            Category:    "Accessories",
            SKU:         "MOUSE001",
            Discount:    5,
            Image:       "mouse.jpg",
            Status:      "active",
        },
        {
            Name:        "Keyboard",
            Description: "Mechanical keyboard",
            Price:       79.99,
            Stock:       100,
            Category:    "Accessories",
            SKU:         "KEYBOARD001",
            Discount:    0,
            Image:       "keyboard.jpg",
            Status:      "active",
        },
    }

    for _, product := range products {
        if err := db.Create(&product).Error; err != nil {
            return err
        }
    }

    fmt.Println("✅ Database seeded successfully!")
    return nil
}