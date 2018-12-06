package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// UserOatuh 第三方登录认证
type UserOatuh struct {
	ID               uint `gorm:"primary_key"`
	UserID           uint
	OauthType        uint
	OauthAccessToken string `gorm:"unique"`
	OauthExpires     string
	status           uint
}

// AddUserOatuh 添加用户账号 与 初始化个人信息
func AddUserOatuh(userOatuh map[string]interface{}) error {

	tx := db.Begin()

	// 首先创建 user
	userID, err := addUser(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(userOatuh)

	oauthInfo := UserOatuh{
		UserID:           userID,
		OauthType:        userOatuh["oauth_type"].(uint),
		OauthAccessToken: userOatuh["access_token"].(string),
		OauthExpires:     userOatuh["expires"].(string),
	}

	if err := tx.Create(&oauthInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// LoginUserOatuh 采用密码方式登录
// func LoginUserOatuh(maps map[string]interface{}) (*UserOatuh, error) {
// 	var user UserOatuh
// 	if err := db.Where(maps).First(&user).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// ExistUserOatuh 判断用户账号是否存在
func ExistUserOatuh(maps map[string]interface{}) (bool, error) {
	var user UserOatuh
	err := db.Select("id").Where(maps).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
