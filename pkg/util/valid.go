package util

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"

	"sync"
)

var (
	trans    ut.Translator
	once     sync.Once
	validate *validator.Validate
)

// GetValidate GetValidator
func GetValidate() *validator.Validate {
	once.Do(func() {
		zhh := zh.New()
		uni := ut.New(zhh, zhh)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("zh")
		validate = validator.New()
		zh_translations.RegisterDefaultTranslations(validate, trans)
	})
	return validate
}

// GetTrans 获取翻译
func GetTrans() ut.Translator {
	return trans
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
