package oauth

import (
	"golang.org/x/oauth2/github"
	"os"

	"golang.org/x/oauth2"
)

var (
	// GithubOauthConfig 认证配置
	GithubOauthConfig *oauth2.Config
)

// Setup 设置oauth2.Config
// 需要在环境变量中设置GITHUB_CLIENT GITHUB_SECRET
func Setup() {
	GithubOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
}
