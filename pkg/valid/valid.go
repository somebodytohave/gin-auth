package valid

import (
	"gopkg.in/go-playground/validator.v9"
)

// GetValidate GetValidator
func GetValidate() *validator.Validate {
	return validator.New()
}

// Email 验证邮箱
func Email(email string) error {
	errs := GetValidate().Var(email, "required,email")
	// output: Key: "" Error:Field validation for "" failed on the "email" tag
	return errs
}

// Phone 验证手机号
func Phone(phone string) error {
	errs := GetValidate().Var(phone, "required,phone")
	// output: Key: "" Error:Field validation for "" failed on the "email" tag
	return errs
}
