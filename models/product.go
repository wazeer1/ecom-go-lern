package models

import "time"

type Product struct {
  ID          uint        `gorm:"primaryKey" json:"id"`
  Name        string      `gorm:"not null" json:"name"`
  Description string      `json:"description"`
  Price       float64     `gorm:"not null" json:"price"`
  Stock       int         `gorm:"default:0" json:"stock"`
  Category    string      `json:"category"`
  Image       string      `json:"image"`
  SKU         string      `gorm:"unique" json:"sku"`
  Discount    float64     `gorm:"default:0" json:"discount"`
  CreatedAt   time.Time   `json:"created_at"`
  UpdatedAt   time.Time   `json:"updated_at"`
  OrderItems  []OrderItem `gorm:"foreignKey:ProductID" json:"order_items,omitempty"`
  Reviews     []Review    `gorm:"foreignKey:ProductID" json:"reviews,omitempty"`
  Status       string    `gorm:"default:'active'" json:"status"`
}