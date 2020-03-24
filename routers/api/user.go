package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sun-wenming/gin-auth/pkg/app"
	"github.com/sun-wenming/gin-auth/pkg/e"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/util"
	"github.com/sun-wenming/gin-auth/service/userser"
)

type auth struct {
	// UserName 用户名
	UserName string `json:"username" example:"zhangsan" validate:"required,gte=5,lte=30"`
	// PassWord 密码
	PassWord string `json:"password" example:"zhangsan" validate:"required,gte=5,lte=30"`
}

// Register 账号密码注册
// @Summary 账号密码注册
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.auth true "账号密码登录/注册"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/register [post]
func Register(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth auth

	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}
	// 验证
	validate := util.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		appG.ResponseFailValidParam(err)
		return
	}

	userService := userser.User{UserName: mAuth.UserName, Password: mAuth.PassWord}

	// 判断是否存在
	exist, err := userService.ExistByUserName()
	if err != nil {
		appG.ResponseFailError(util.ErrNewCode(e.ErrorUserNameNotExist))
		return
	}

	if exist {
		appG.ResponseFailError(util.ErrNewCode(e.ErrorUserNameExist))
		return
	}

	// 注册
	if err := userService.Register(); err != nil {
		appG.ResponseFailError(err)
		return
	}

	// 注册成功之后 make token
	token, err := util.GenerateToken(mAuth.UserName)
	if err != nil {
		logging.GetLogger().Error(err)
		appG.ResponseFailError(util.ErrNewCode(e.ErrorAuthGenerateToken))
		return
	}
	appG.ResponseSuc(token)
}

// Login 账号密码登录
// @Summary 账号密码登录
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.auth true "用户信息"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Failure 204 {object} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth auth
	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}
	// 验证
	validate := util.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		appG.ResponseFailValidParam(err)
		return
	}

	// 传值
	userService := userser.User{
		UserName: mAuth.UserName,
		Password: mAuth.PassWord,
	}

	token, merr := userService.PwdLogin()
	// 登录查询成功
	if merr != nil {
		appG.ResponseFailError(merr)
		return
	}

	appG.ResponseSuc(token)
}

type phone struct {
	// Phone 手机号
	Phone string `json:"phone" example:"13938738804" validate:"required"`
	// Code 手机号验证码
	Code string `json:"code" example:"123456" validate:"required"`
}

// PhoneLogin 手机号快速登陆/注册
// @Summary 手机号快速登陆/注册
// @Document 如果登录手机号未注册,则自动注册再登录
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param auth body api.phone true "手机号快速登录/注册"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/phonelogin [post]
func PhoneLogin(c *gin.Context) {
	appG := app.GetGin(c)
	var mAuth phone

	// 解析 body json 数据到实体类
	if err := c.ShouldBindJSON(&mAuth); err != nil {
		appG.ResponseFailError(util.ErrNewErr(err))
		return
	}
	// 验证
	validate := util.GetValidate()
	err := validate.Struct(mAuth)
	if err != nil {
		appG.ResponseFailValidParam(err)
		return
	}

	if !util.RegPhone(mAuth.Phone) {
		appG.ResponseFailErrCode(e.ErrorPhoneNotValid)
		return
	}

	// 验证验证码
	code := userser.GetCacheCode(mAuth.Phone)
	if code == "" {
		appG.ResponseFailErrCode(e.ErrorPhoneCodeExpired)
		return
	}
	if mAuth.Code != code {
		appG.ResponseFailErrCode(e.ErrorPhoneCodeNotValid)
		return
	}

	userService := userser.User{UserName: mAuth.Phone, Code: mAuth.Code}

	// 判断是否存在
	exist, merr := userService.ExistByUserName()
	if merr != nil {
		appG.ResponseFailError(merr)
		return
	}

	if !exist { // 注册
		if err := userService.PhoneRegister(); err != nil {
			appG.ResponseFailError(err)
			return
		}
	}

	// 登录 make token
	token, err := util.GenerateToken(mAuth.Phone)
	if err != nil {
		logging.GetLogger().Error(err)
		appG.ResponseFailError(util.ErrNewCode(e.ErrorAuthGenerateToken))
		return
	}
	appG.ResponseSuc(token)
}

// SendCode 发送手机验证码
// @Summary 发送手机验证码
// @accept application/x-www-form-urlencoded
// @Tags auth
// @Produce  json
// @Param phone formData string true "手机号"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth/code [post]
func SendCode(c *gin.Context) {
	appG := app.GetGin(c)
	phone := c.PostForm("phone")
	code, err := userser.SendCode(phone)
	if !util.RegPhone(phone) {
		appG.ResponseFailErrCode(e.ErrorPhoneNotValid)
		return
	}

	if err != nil {
		appG.ResponseFailError(err)
	}
	appG.ResponseSuc(code)
}

// GetUserInfo 获取用户信息
// @Summary 获取用户信息
// @accept application/x-www-form-urlencoded
// @Security ApiKeyAuth
// @Tags user
// @Produce  json
// @Success 200 {object} users.User
// @Router /api/v1/user/getUserInfo [post]
func GetUserInfo(c *gin.Context) {
	appG := app.GetGin(c)

	username, err := util.GetTokenLoginName(c)
	if err != nil {
		appG.ResponseFailError(err)
		return
	}

	userService := userser.User{UserName: username}

	// 判断是否存在
	uid, err := userService.UserLoginGetUserID()
	if err != nil {
		appG.ResponseFailError(err)
		return
	}
	userService.ID = uid

	user, err := userService.GetUserInfo()
	if err != nil {
		appG.ResponseFailError(err)
		return
	}

	appG.ResponseSuc(user)
}
