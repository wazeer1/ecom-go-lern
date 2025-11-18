package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/yourusername/ecommerce-api/models"
)

// GetUserProfile retrieves user profile
func GetUserProfile(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id")

    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    response := models.UserResponse{
        ID:       user.ID,
        Email:    user.Email,
        FullName: user.FullName,
        Phone:    user.Phone,
        Avatar:   user.Avatar,
        IsAdmin:  user.IsAdmin,
    }

    c.JSON(http.StatusOK, response)
}

// UpdateUserProfile updates user profile
func UpdateUserProfile(c *gin.Context, db *gorm.DB) {
    userID := c.Param("id")

    type UpdateRequest struct {
        FullName string `json:"full_name"`
        Phone    string `json:"phone"`
        Avatar   string `json:"avatar"`
    }

    var req UpdateRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := db.Model(&user).Where("id = ?", userID).Updates(req).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Profile updated"})
}