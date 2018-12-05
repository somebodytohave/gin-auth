package oauth

import (
	"fmt"
	"context"
	"golang.org/x/oauth2/github"
	"os"

	"golang.org/x/oauth2"
)

// Setup 设置oauth2.Config
// 需要在环境变量中设置GITHUB_CLIENT GITHUB_SECRET
func Setup() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{"repo", "user"},
		Endpoint:     github.Endpoint,
	}
}

// GetToken 检索github oauth2令牌
func GetToken(ctx context.Context,conf *oauth2.Config)(*oauth2.Token,error){
	url := conf.AuthCodeURL("state")
	fmt.Printf("Type the following url into your browser and follow the directions on screen: %v\n", url)
	fmt.Println("Paste the code returned in the redirect URL and hit Enter:")

	var code string 
	if _,err := fmt.Scan(&code);err !=nil{
		return nil,err
	}
	// 将授权码转换成Token
	return conf.Exchange(ctx,code)



}