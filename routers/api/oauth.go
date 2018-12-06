package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/mecm/gin-auth/pkg/app"
	"github.com/mecm/gin-auth/pkg/oauth"
	"github.com/mecm/gin-auth/pkg/util"
	"golang.org/x/oauth2"
	"net/http"
)

var oauthStateString = "random-user"

// LoginGithub github登录
func LoginGithub(c *gin.Context) {
	oauthStateString = util.GetRandomSalt()
	url := oauth.GithubOauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	c.Redirect(http.StatusMovedPermanently, url)
}

// CallBackGithub 登录成功
func CallBackGithub(c *gin.Context) {
	state, _ := c.GetQuery("state")
	code, _ := c.GetQuery("code")

	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	token, err := oauth.GithubOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	oauthClient := oauth.GithubOauthConfig.Client(oauth2.NoContext, token)
	client := github.NewClient(oauthClient)
	user, _, err := client.Users.Get(oauth2.NoContext, "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", *user)
	// TODO 将用户信息存入 数据库
	appG := app.GetGin(c)
	appG.ResponseSuc(*user)

}
