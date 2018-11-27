package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password" gorm:"unique"`
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.Where("id = ? ", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func AddUser(data map[string]interface{}) error {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(username string) (*User, error) {
	user := User{
		Username: username,
	}
	if err := db.Where("username = ? ", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

// ExistUser 检查是否存在此用户
func ExistUser(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}
