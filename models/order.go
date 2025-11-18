package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    UserID          uint      `gorm:"not null;index" json:"user_id"`
    User            *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`

    OrderNumber     string    `gorm:"unique;not null" json:"order_number"`
    TotalAmount     float64   `gorm:"not null;check:total_amount>=0" json:"total_amount"`
    DiscountAmount  float64   `gorm:"default:0" json:"discount_amount"`
    TaxAmount       float64   `gorm:"default:0" json:"tax_amount"`

    Status          string    `gorm:"default:'pending';index" json:"status"`
    PaymentStatus   string    `gorm:"default:'unpaid'" json:"payment_status"`
    PaymentID       string    `json:"payment_id"`

    ShippingAddr    string    `json:"shipping_address"`
    BillingAddr     string    `json:"billing_address"`
    ShippingMethod  string    `json:"shipping_method"`
    TrackingNumber  string    `json:"tracking_number"`

    Notes           string    `gorm:"type:text" json:"notes"`

    // Relations
    OrderItems      []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"order_items"`
    Payment         *Payment `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"payment,omitempty"`

    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Order) TableName() string {
    return "orders"
}