package util

import (
	"github.com/mecm/gin-auth/pkg/setting"

	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

// GetPage 保证了各接口的page处理是一致的
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
