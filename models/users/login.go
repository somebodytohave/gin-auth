package users

import (
	"github.com/sun-wenming/gin-auth/pkg/e"
	"errors"
	"fmt"
	"github.com/sun-wenming/gin-auth/models"

	"github.com/jinzhu/gorm"
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
func AddUserLogin(userLogin map[string]interface{}) error {

	tx := models.DB.Begin()

	// 首先创建 user
	userID, err := addUser(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(userLogin)

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
	fmt.Println(loginInfo)

	if err := tx.Create(&loginInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// LoginUserLogin 采用密码方式登录
func LoginUserLogin(maps map[string]interface{}) (*UserLogin, error) {
	var user UserLogin
	if err := models.DB.Where(maps).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}


// ExistUserLogin 返回用户ID
func ExistUserLogin(maps map[string]interface{}) (uint, error) {
	var user UserLogin
	err := models.DB.Select("id,user_id").Where(maps).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, err
	}
	if user.ID < 1 { // 判断用户账号是否存在
		return 0, errors.New(e.GetMsg(e.ERROR_USER_NAME_NOT_EXIST))
	}
	if user.UserID < 1 { // 判断用户信息是否存在
		return 0, errors.New(e.GetMsg(e.ERROR_USER_INFO_EMPTY))
	}

	return user.UserID, nil
}

