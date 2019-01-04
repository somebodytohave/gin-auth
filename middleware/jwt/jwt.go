package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/gin-auth/pkg/app"
	"github.com/sun-wenming/gin-auth/pkg/e"
	"github.com/sun-wenming/gin-auth/pkg/util"
	"net/http"
	"time"
)

// JWT middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.GetGin(c)
		var code int
		var data interface{}
		code = e.SUCCESS
		token := c.Request.Header.Get("jwtToken")
		if token == "" {
			code = e.INVALID_PARAMS_WITHOUT_TOKEN
		} else {
			claims, err := util.ParseToken2(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			appG.Response(http.StatusUnauthorized, code, data)
			// 拦截
			c.Abort()
			return
		}
		c.Next()
	}

}
