package user_service

import (
	"errors"
	"github.com/mecm/gin-auth/models"
	"github.com/mecm/gin-auth/pkg/util"
	"github.com/mecm/gin-auth/pkg/valid"
)

type User struct {
	ID int

	UserName string
	Password string
	NickName string
}

// Register 注册用户
func (u *User) Register() error {

	password, err := util.Encrypt(u.Password)

	if err != nil {
		return err
	}

	valid.ValidEmail(u.UserName)

	userLogin := map[string]interface{}{
		"username": u.UserName,
		"password": password,
	}

	userProfile := map[string]interface{}{
		"nickname": u.NickName,
	}

	return models.AddUserLogin(userProfile, userLogin)
}

func (u *User) Login() error {

	user, err := models.LoginUserLogin(u.UserName)
	if err != nil {
		return err
	}

	inputPwd := u.Password
	hashPwd := user.Password

	// 比较 密码
	if err := util.Compare(inputPwd, hashPwd); err != nil {
		return err
	}

	// 匹配正常 根据 userID 查询 用户信息
	if !(user.UserID > 0) {
		return errors.New("未对应的用户")
	}

	models.GetUser(user.UserID)

	return nil
}
