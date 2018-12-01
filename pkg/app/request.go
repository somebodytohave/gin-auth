package app

import (
	"github.com/mecm/gin-auth/pkg/logging"
)

// MarkError 将错误 存入日志
func MarkError(v ...interface{}) {
	logging.Info(v...)
	return
}
