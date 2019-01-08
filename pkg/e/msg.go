package e

const (
	SUCCESS                        = 200
	ERROR                          = 500
	ErrorInvalidParamsWithoutToken = 401
	ErrorInvalidParams             = 400

	// 认证
	ErrorAuthParseTokenFail = iota + 20001
	ErrorAuthCheckTokenTimeout
	ErrorAuthGenerateToken
	ErrorAuthToken

	// --- 客户端错误
	ErrorUserGetInfo = iota + 40001
	ErrorUserGetLogin
	ErrorUserRegName
	ErrorUserNameExist
	ErrorPhoneNotValid
	ErrorPhoneCodeSend
	ErrorPhoneCodeExpired
	ErrorPhoneCodeNotValid
	ErrorUserName
	ErrorUserNameNotExist
	ErrorUserInfoEmpty
	ErrorUserLoginEmpty
	ErrorUserPwd
	// --- end

	// --- 服务器错误
	ErrorExecSql = iota + 50001
	ErrorPasswordEncrypt
	ErrorUserInfoCreate
	ErrorUserLoginCreate
	ErrorOauthState
	ErrorOauthCode
	ErrorOauthInfo

	// --- end

)

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	ErrorInvalidParams:             "请求参数错误",
	ErrorInvalidParamsWithoutToken: "不存在Token参数",

	// 认证
	ErrorAuthParseTokenFail:    "Token解析失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthGenerateToken:     "Token生成失败",
	ErrorAuthToken:             "Token错误",
	// --- end

	ErrorUserGetInfo:       "获取到用户失败.",
	ErrorUserGetLogin:       "获取到帐户失败.",
	ErrorUserRegName:       "用户名输入格式错误.",
	ErrorUserNameExist:     "用户名已存在.",
	ErrorPhoneNotValid:     "手机号验证失败.",
	ErrorPhoneCodeSend:     "验证码发送失败.",
	ErrorPhoneCodeExpired:  "验证码已过期.",
	ErrorPhoneCodeNotValid: "验证码验证失败.",
	ErrorUserNameNotExist:  "用户名不存在.",
	ErrorUserInfoEmpty:     "用户不存在.",
	ErrorUserLoginEmpty:    "帐户不存在.",
	ErrorUserPwd:           "密码错误.",

	// --- 服务器错误
	ErrorPasswordEncrypt: "密码加密失败.",
	ErrorUserInfoCreate:  "用户信息创建失败.",
	ErrorUserLoginCreate: "用户帐户创建失败.",
	ErrorOauthState:      "三方登录状态码错误.",
	ErrorOauthCode:       "三方登录获取token失败.",
	ErrorOauthInfo:       "三方登录获取信息失败.",
	ErrorUserName:       "获取用户名失败.",
	// --- end
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
