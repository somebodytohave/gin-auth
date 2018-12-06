package api

import (
	"github.com/mecm/gin-auth/pkg/util/reg"
	"github.com/gin-gonic/gin"
	"github.com/mecm/gin-auth/pkg/app"
	"github.com/mecm/gin-auth/pkg/e"
	"github.com/mecm/gin-auth/pkg/util"
	"github.com/mecm/gin-auth/pkg/util/valid"
	"github.com/mecm/gin-auth/service/user_service"
)

type auth struct {
	// UserName 用户名
	UserName string `json:"username" example:"zhangsan" validate:"required,gte=5,lte=30"`
	// PassWord 密码
	PassWord string `json:"password" example:"zhangsan" validate:"required,gte=5,lte=30"`
	// Type 类型 1：账号密码 2：手机号 3：第三方
	Type int `json:"type" example:"1" validate:"required,oneof=1 2 3"`
}

// Register 注册新用户
// @Summary 注册新用户
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

	userService := user_service.User{UserName: mAuth.UserName, Password: mAuth.UserName}

	exist, err := userService.ExistByName()
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}
	if exist {
		appG.ResponseFailMsg(e.GetMsg(e.ERROR_USER_NAME_EXIST))
		return
	}

	// 注册
	if err := userService.Register(); err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	// 注册成功之后 make token
	token, err := util.GenerateToken(mAuth.UserName, mAuth.PassWord)
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

	// 登录查询成功
	if err := userService.Login(); err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}

	// 生成token
	token, err := util.GenerateToken(mAuth.UserName, mAuth.PassWord)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
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

// PhoneLogin 手机号快速登陆
// @Summary 手机号快速登陆
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

	if !reg.Phone(mAuth.Phone){
		appG.ResponseFailErrCode(e.ERROR_PHONE_NOT_VALID)
		return
	}

	// 验证验证码
	code := user_service.GetCacheCode(mAuth.Phone)
	if code ==""{
		appG.ResponseFailErrCode(e.ERROR_PHONE_CODE_EXPIRED)
		return
	}
	if mAuth.Code != code{
		appG.ResponseFailErrCode(e.ERROR_PHONE_CODE_NOT_VALID)
		return
	} 

	userService := user_service.User{UserName: mAuth.Phone, Code: mAuth.Code}

	exist, err := userService.ExistByName()
	if err != nil {
		appG.ResponseFailMsg(err.Error())
		return
	}
	
	if !exist {// 注册
		if err := userService.PhoneRegister();err !=nil{
			appG.ResponseFailMsg(err.Error())
			return
		}
	}

	// 登录 make token
	token, err := util.GenerateToken(mAuth.Phone, mAuth.Code)
	if err != nil {
		appG.ResponseFailMsg(err.Error())
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
	code,err := user_service.SendCode(phone)
	if !reg.Phone(phone){
		appG.ResponseFailErrCode(e.ERROR_PHONE_NOT_VALID)
		return
	}

	if err !=nil{
		appG.ResponseFailMsg(err.Error())
	}
	appG.ResponseSuc(code)
}
