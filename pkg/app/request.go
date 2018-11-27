package app

import (
	"github.com/astaxie/beego/validation"
	"github.com/mecm/gin-blog/pkg/logging"
)

// MarkValidErrors 将错误 存入日志
func MarkValidErrors(errors []*validation.Error) {
	for _, err := range errors {
		logging.Info(err.Key, err.Message)
	}
	return
}

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	logging.Error(v...)
	return
}
