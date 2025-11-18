package models

import (
	"time"
)

type User struct {
  ID        uint      `gorm:"primaryKey" json:"id"`
  Email     string    `gorm:"unique;not null" json:"email"`
  Password  string    `gorm:"-" json:"-"`
  FullName  string    `json:"full_name"`
  Phone     string    `json:"phone"`
  IsAdmin   bool      `gorm:"default:false" json:"is_admin"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  Orders    []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`
  Avatar    string    `json:"avatar"`
}

type UserResponse struct {
    ID       uint   `json:"id"`
    Email    string `json:"email"`
    FullName string `json:"full_name"`
    Phone    string `json:"phone"`
    Avatar   string `json:"avatar"`
    IsAdmin  bool   `json:"is_admin"`
}