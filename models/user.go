package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Nickname string `json:"nickname"`
}

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

// ExistUser 检查是否存在此用户
// func ExistUser(username, password string) (bool, error) {
// 	var user User
// 	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
// 	if err != nil && err != gorm.ErrRecordNotFound {
// 		return false, err
// 	}

// 	if user.ID > 0 {
// 		return true, nil
// 	}

// 	return false, nil
// }
