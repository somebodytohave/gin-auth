package util

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"reflect"
	"strings"
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

		customFieldName()

		validate = validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return "{{" + name + "}}"
		})

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


// 自定义字段名称
// 参考自 : https://github.com/syssam/go-playground-sample/blob/master/main.go
func customFieldName() {
	var errAdd error
	errAdd = trans.Add("{{username}}", "用户名", false)
	errAdd = trans.Add("{{password}}", "密码", false)
	logging.GetLogger().Warn("添加自定义字段翻译失败", errAdd)
}
