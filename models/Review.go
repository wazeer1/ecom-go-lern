package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    ProductID uint      `gorm:"not null;index" json:"product_id"`
    UserID    uint      `gorm:"not null;index" json:"user_id"`

    Product   *Product  `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE" json:"product,omitempty"`
    User      *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"user,omitempty"`

    Rating    int       `gorm:"not null;check:rating>=1 AND rating<=5" json:"rating"`
    Title     string    `gorm:"not null" json:"title"`
    Comment   string    `gorm:"type:text" json:"comment"`
    Helpful   int       `gorm:"default:0" json:"helpful_count"`

    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (Review) TableName() string {
    return "reviews"
}