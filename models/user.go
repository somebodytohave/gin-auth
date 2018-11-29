package models

import (
	"github.com/jinzhu/gorm"
)

// User 用户信息表
type User struct {
	gorm.Model
	Nickname string `json:"nickname"`
}

// GetUser 获取用户信息
func GetUser(id uint) (*User, error) {
	var user User
	err := db.Where("id = ? ", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// AddUser 新增用户信息
func addUser(data map[string]interface{}, tx *gorm.DB) (uint, error) {

	user := User{
		Nickname: data["nickname"].(string),
	}

	if err := tx.Create(&user).Error; err != nil {
		return 0, err
	}
	return user.ID, nil
}

// ExistUserByID 检查是否存在此用户
func ExistUserByID(id uint) (bool, error) {
	var user User
	err := db.Select("id").Where("id = ? ", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
