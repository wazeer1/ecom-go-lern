package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    UserID    uint      `gorm:"not null;index" json:"user_id"`
    ProductID uint      `gorm:"not null;index" json:"product_id"`

    User      *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
    Product   *Product  `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product,omitempty"`

    Quantity  int       `gorm:"not null;check:quantity>0" json:"quantity"`
    AddedAt   time.Time `json:"added_at"`

    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (CartItem) TableName() string {
    return "cart_items"
}