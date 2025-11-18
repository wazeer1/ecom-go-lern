package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wazeer1/ecommerce-api/models"
	"github.com/wazeer1/ecommerce-api/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterRequest struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required,min=6"`
  FullName string `json:"full_name" binding:"required"`
}

func Register(c *gin.Context, db *gorm.DB) {
  var req RegisterRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    print(err)
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
    return
  }
  
  user := models.User{
    Email:    req.Email,
    Password: string(hashedPassword),
    FullName: req.FullName,
  }
  
  if err := db.Create(&user).Error; err != nil {
    c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
    return
  }
  
  token, _ := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
  c.JSON(http.StatusCreated, gin.H{
    "message": "User registered successfully",
    "token": token,
  })
}

type LoginRequest struct {
  Email    string `json:"email" binding:"required,email"`
  Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context, db *gorm.DB) {
  var req LoginRequest
  if err := c.ShouldBindJSON(&req); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  
  var user models.User
  if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  }
  
  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
    return
  }
  
  token, _ := utils.GenerateToken(user.ID, user.Email, user.IsAdmin)
  c.JSON(http.StatusOK, gin.H{
    "message": "Login successful",
    "token": token,
    "user": user,
  })
}