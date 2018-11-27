package user_service

import (
	"github.com/mecm/gin-blog/models"
	"github.com/mecm/gin-blog/pkg/util"
)

type User struct {
	ID int

	UserName string
	Password string
}

// Register 注册用户
func (u *User) Register() error {

	password, err := util.Encrypt(u.Password)
	if err != nil {
		return err
	}

	data := map[string]interface{}{
		"username": u.UserName,
		"password": password,
	}

	return models.AddUser(data)
}

func (u *User) Login() error {
	user, err := models.LoginUser(u.UserName)
	if err != nil {
		return err
	}

	inputPwd := u.Password
	hashPwd := user.Password

	// 比较 密码
	if err := util.Compare(inputPwd, hashPwd); err != nil {
		return err
	}

	return nil
}
