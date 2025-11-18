package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
  UserID uint   `json:"user_id"`
  Email  string `json:"email"`
  IsAdmin bool  `json:"is_admin"`
  jwt.RegisteredClaims
}

func GenerateToken(userID uint, email string, isAdmin bool) (string, error) {
  expirationTime := time.Now().Add(24 * time.Hour)
  claims := &Claims{
    UserID: userID,
    Email: email,
    IsAdmin: isAdmin,
    RegisteredClaims: jwt.RegisteredClaims{
      ExpiresAt: jwt.NewNumericDate(expirationTime),
    },
  }
  
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
  claims := &Claims{}
  token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return jwtSecret, nil
  })
  
  if err != nil || !token.Valid {
    return nil, err
  }
  return claims, nil
}