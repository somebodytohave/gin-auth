package userser

import (
	"github.com/sun-wenming/gin-auth/models/users"
	"github.com/sun-wenming/gin-auth/pkg/util"
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
func (o UserOauth) LoginGithub()  util.Error  {
	maps := make(map[string]interface{})

	maps["oauth_id"] = o.OauthID
	maps["oauth_type"] = o.OauthType
	maps["access_token"] = o.OauthAccessToken
	maps["expires"] = o.OauthExpires

	return users.AddUserOauth(maps)
}

// ExistUserOauth 存在第三方登录
func (o UserOauth) ExistUserOauth() (bool, util.Error) {
	maps := map[string]interface{}{"oauth_id": o.OauthID}
	return users.ExistUserOauth(maps)
}
