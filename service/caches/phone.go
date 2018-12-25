package caches

type Phone struct {
	Phone string
}

// GetPhoneCodeKey 获取手机号验证码的 key
func (p Phone) GetPhoneCodeKey() string {
	return CachePhone + "_" + p.Phone
}
