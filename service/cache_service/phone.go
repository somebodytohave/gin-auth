package cache_service

import (
	"github.com/mecm/gin-auth/pkg/e"
	"strconv"
)

type Phone struct {
	Phone string
}

// GetPhoneCodeKey 获取手机号验证码的 key
func (p Phone) GetPhoneCodeKey() string {
	return e.CACHE_PHONE + "_" + p.Phone
}
