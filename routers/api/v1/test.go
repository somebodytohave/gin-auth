package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/gin-auth/pkg/app"
)

// TestAuth 测试认证
// @Summary 测试认证
// @Tags test
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Security ApiKeyAuth
// @Router /api/v1/test [get]
func TestAuth(c *gin.Context) {
	appG := app.GetGin(c)

	appG.ResponseSuc("auth success")
}
