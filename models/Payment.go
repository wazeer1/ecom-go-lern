package models

import (
	"time"
)

type Payment struct {
    ID              uint      `gorm:"primaryKey" json:"id"`
    OrderID         uint      `gorm:"not null;unique" json:"order_id"`

    StripePaymentID string    `gorm:"unique" json:"stripe_payment_id"`
    Amount          float64   `gorm:"not null" json:"amount"`
    Currency        string    `gorm:"default:'USD'" json:"currency"`

    Status          string    `gorm:"default:'pending'" json:"status"`
    PaymentMethod   string    `json:"payment_method"`

    TransactionID   string    `json:"transaction_id"`
    ReceiptURL      string    `json:"receipt_url"`

    ErrorMessage    string    `json:"error_message,omitempty"`

    PaidAt          *time.Time `json:"paid_at"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}

func (Payment) TableName() string {
    return "payments"
}