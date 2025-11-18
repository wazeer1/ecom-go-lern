package models

import "gorm.io/gorm"

type OrderItem struct {
    ID          uint     `gorm:"primaryKey" json:"id"`
    OrderID     uint     `gorm:"not null;index" json:"order_id"`
    ProductID   uint     `gorm:"not null" json:"product_id"`

    Product     *Product `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT" json:"product,omitempty"`

    Quantity    int      `gorm:"not null;check:quantity>0" json:"quantity"`
    Price   float64  `gorm:"not null" json:"price"`
    Discount    float64  `gorm:"default:0" json:"discount"`
}

func (OrderItem) TableName() string {
    return "order_items"
}

// AfterFind hook to calculate subtotal
func (oi *OrderItem) AfterFind(tx *gorm.DB) (err error) {
    return nil
}

// Subtotal calculates the subtotal for the order item
func (oi *OrderItem) Subtotal() float64 {
    return (oi.Price * float64(oi.Quantity)) * (1 - oi.Discount/100)
}