package util

import (
	"gopkg.in/go-playground/validator.v9"
)

// GetValidate GetValidator
func GetValidate() *validator.Validate {
	return validator.New()
}

// ValidEmail 验证邮箱
func ValidEmail(email string) bool {
	errs := GetValidate().Var(email, "email")
	if errs != nil {
		return false
	}
	// output: Key: "" Error:Field validation for "" failed on the "email" tag
	return true
}
