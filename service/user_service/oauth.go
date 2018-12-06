package user_service

import (
	"github.com/mecm/gin-auth/models"
)

// UserOatuh UserOatuh
type UserOatuh struct {
	ID               uint
	OauthType        uint
	OauthAccessToken string
	OauthExpires     string
	NickName         string
	status           uint
}

// LoginGithub 注册认证登录
func (o UserOatuh) LoginGithub() error {
	maps := make(map[string]interface{})

	// 用户信息
	userProfile := map[string]interface{}{
		"nickname": o.NickName,
	}

	maps["access_token"] = o.OauthAccessToken
	maps["oauth_type"] = o.OauthType
	maps["expires"] = o.OauthExpires

	return models.AddUserOatuh(userProfile, maps)
}
