package users

import (
	"github.com/jinzhu/gorm"
	"github.com/sun-wenming/gin-auth/models"
	"github.com/sun-wenming/gin-auth/pkg/e"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/util"
)

// User 用户信息表
type User struct {
	models.Model
	Nickname string `json:"nickname"`
}

// GetUser 获取用户信息
func GetUser(id uint) (*User, util.Error) {
	var user User
	err := models.DB.Where("id = ? ", id).First(&user).Error
	if err != nil {
		logging.GetLogger().Error(err)
		return nil, util.ErrNewCode(e.ErrorUserGetInfo)
	}
	return &user, nil
}

// AddUser 新增用户信息
func addUser(tx *gorm.DB) (uint, error) {
	var user User
	if err := tx.Create(&user).Error; err != nil {
		logging.GetLogger().Error(err)
		return 0, err
	}
	return user.ID, nil
}

// ExistUserByID 检查是否存在此用户
func ExistUserByID(id uint) (bool, util.Error) {
	var user User
	err := models.DB.Select("id").Where("id = ? ", id).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.GetLogger().Error(err)
		return false, util.ErrNewCode(e.ErrorUserGetInfo)
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
