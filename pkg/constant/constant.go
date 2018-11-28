package constant

// ConstantType ...
type ConstantType int

const (
	// 账号名
	LOGIN_NAME ConstantType = iota
	// 手机号
	LOGIN_PHONE
	// 邮箱
	LOGIN_EMAIL
)
