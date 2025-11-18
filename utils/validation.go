package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateEmail(email string) bool {
  pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
  match, _ := regexp.MatchString(pattern, email)
  return match
}

func ValidatePassword(password string) bool {
  if len(password) < 6 {
    return false
  }
  return true
}

func ValidateStruct(data interface{}) error {
  return validate.Struct(data)
}