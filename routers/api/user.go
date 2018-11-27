package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mecm/gin-blog/pkg/app"
	"github.com/mecm/gin-blog/pkg/logging"
	"github.com/mecm/gin-blog/pkg/util"
	"github.com/mecm/gin-blog/pkg/valid"
	"github.com/mecm/gin-blog/service/user_service"
)

type auth struct {
	// UserName 用户名
	UserName string `json:"username" example:"zhangsan" validate:"required,gte=5,lte=50"`
	// PassWord 密码
	PassWord string `json:"password" example:"zhangsan" validate:"required,gte=5,lte=50"`
}

// Register 注册新用户
// @Summary 注册新用户
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.auth true "用户信息"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth auth
	logging.Info(mAuth)

	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}
	// 验证
	validate := valid.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	userService := user_service.User{UserName: mAuth.UserName, Password: mAuth.PassWord}
	if err := userService.Register(); err != nil {
		logging.Info(err)
		appG.ResponseFailMsg(err.Error())
		return
	}

	password, err := util.Encrypt(mAuth.PassWord)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	// 注册成功之后
	token, err := util.GenerateToken(mAuth.UserName, password)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	appG.ResponseSuc(token)

}

// Login 登录
// @Summary 登录
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.auth true "用户信息"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth auth
	logging.Info(mAuth)

	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}
	// 验证
	validate := valid.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	userService := user_service.User{
		UserName: mAuth.UserName,
		Password: mAuth.PassWord,
	}

	if err := userService.Login(); err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	// auth
	password, err := util.Encrypt(mAuth.PassWord)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	token, err := util.GenerateToken(mAuth.UserName, password)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	appG.ResponseSuc(token)
}
