package user_service

import (
	"github.com/sun-wenming/gin-auth/models"
)

// UserOauth UserOauth
type UserOauth struct {
	ID               uint
	OauthType        uint
	OauthID          string
	OauthAccessToken string
	OauthExpires     string
	NickName         string
	status           uint
}

// LoginGithub 注册认证登录
func (o UserOauth) LoginGithub() error {
	maps := make(map[string]interface{})

	maps["oauth_id"] = o.OauthID
	maps["oauth_type"] = o.OauthType
	maps["access_token"] = o.OauthAccessToken
	maps["expires"] = o.OauthExpires

	return models.AddUserOauth(maps)
}

func (o UserOauth) ExistUserOauth() (bool, error) {
	maps := map[string]interface{}{"oauth_id": o.OauthID}
	return models.ExistUserOauth(maps)
}
