package valid

import (
	"gopkg.in/go-playground/validator.v9"
)

// GetValidator GetValidator
func GetValidate() *validator.Validate {
	return validator.New()
}
