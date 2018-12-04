package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mecm/gin-auth/docs"
	"github.com/mecm/gin-auth/middleware/jwt"
	"github.com/mecm/gin-auth/pkg/setting"
	"github.com/mecm/gin-auth/routers/api"
	"github.com/mecm/gin-auth/routers/api/v1"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter 返回 框架的实例 包含中间件 配置
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	// 根目录
	root := r.Group("")
	{
		// 用户账号
		account := root.Group("/auth")
		{
			account.POST("/register", api.Register)
			account.POST("/login", api.Login)
		}

		apiv1 := root.Group("/api/v1")

		apiv1.Use(jwt.JWT())
		{
			apiv1.GET("/test", v1.TestAuth)
		}

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
