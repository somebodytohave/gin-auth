package users

import (
	"github.com/jinzhu/gorm"
	"github.com/sun-wenming/gin-auth/models"
	"github.com/sun-wenming/gin-auth/pkg/e"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/util"
)

// UserLogin 用户密码登陆认证
type UserLogin struct {
	ID         uint `gorm:"primary_key"`
	UserID     uint
	LoginName  string `gorm:"unique"`
	LoginEmail string `gorm:"unique"`
	LoginPhone string `gorm:"unique"`
	Password   string
	status     uint
}

// 验证码登录
func codeLogin() {

}

// AddUserLogin 添加用户账号 与 初始化个人信息
func AddUserLogin(userLogin map[string]interface{}) util.Error {

	tx := models.DB.Begin()

	// 首先创建 user
	userID, err := addUser(tx)
	if err != nil {
		tx.Rollback()
		logging.GetLogger().Error(err)
		return util.ErrNewCode(e.ErrorUserInfoCreate)
	}

	loginInfo := UserLogin{
		UserID: userID,
	}

	if userLogin["password"] != nil {
		loginInfo.Password = userLogin["password"].(string)
	}

	if userLogin["login_name"] != nil {
		loginInfo.LoginName = userLogin["login_name"].(string)
		goto InsertLogin
	}
	if userLogin["login_phone"] != nil {
		loginInfo.LoginPhone = userLogin["login_phone"].(string)
		goto InsertLogin
	}
	if userLogin["login_email"] != nil {
		loginInfo.LoginEmail = userLogin["login_email"].(string)
	}
InsertLogin:

	if err := tx.Create(&loginInfo).Error; err != nil {
		tx.Rollback()
		logging.GetLogger().Error(err)
		return util.ErrNewCode(e.ErrorUserLoginCreate)
	}
	tx.Commit()
	return nil
}

// LoginUserLogin 采用密码方式登录
func LoginUserLogin(maps map[string]interface{}) (*UserLogin, util.Error) {
	var user UserLogin
	if err := models.DB.Where(maps).First(&user).Error; err != nil {
		logging.GetLogger().Error(err)
		return nil, util.ErrNewSql(err)
	}
	return &user, nil
}

// ExistUserLogin 判断是否存在此用户账号
func ExistUserLogin(maps map[string]interface{}) (bool, util.Error) {
	var user UserLogin
	err := models.DB.Select("id").Where(maps).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		logging.GetLogger().Error(err)
		return false, util.ErrNewCode(e.ErrorUserGetLogin)
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// UserLoginGetUserID 通过用户名 获取 用户ID
func UserLoginGetUserID(maps map[string]interface{}) (uint, util.Error) {
	var user UserLogin
	err := models.DB.Select("id,user_id").Where(maps).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, util.ErrNewSql(err)
	}
	if user.ID < 1 { // 判断用户账号是否存在
		return 0, util.ErrNewCode(e.ErrorUserNameNotExist)
	}
	if user.UserID < 1 { // 判断用户信息是否存在
		return 0, util.ErrNewCode(e.ErrorUserInfoEmpty)
	}

	return user.UserID, nil
}
