package user_service

import (
	"errors"
	"github.com/mecm/gin-auth/models"
	"github.com/mecm/gin-auth/pkg/e"
	"github.com/mecm/gin-auth/pkg/util"
	"github.com/mecm/gin-auth/pkg/util/reg"
	"github.com/mecm/gin-auth/pkg/util/valid"
)

type User struct {
	ID uint

	UserName string
	Password string
	NickName string
}

// Register 注册用户
func (u *User) Register() error {

	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return err
	}
	maps["password"] = u.Password

	// 用户信息
	userProfile := map[string]interface{}{
		"nickname": u.NickName,
	}
	// 创建 用户信息 与 用户密码
	return models.AddUserLogin(userProfile, maps)
}

// Login 登录用户
func (u *User) Login() error {
	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return err
	}
	// 查询 用户登录信息
	user, err := models.LoginUserLogin(maps)
	if err != nil {
		return err
	}

	inputPwd := u.Password
	hashPwd := user.Password

	// 比较 密码
	if err := util.Compare(inputPwd, hashPwd); err != nil {
		return err
	}

	// 匹配成功  根据 userID 查询 用户信息
	if !(user.UserID > 0) {
		return errors.New(e.GetMsg(e.ERROR_USER_GET_INFO))
	}
	// 匹配正常 根据 userID 查询 用户信息
	exist, err := models.ExistUserByID(user.UserID)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New(e.GetMsg(e.ERROR_USER_GET_INFO))
	}

	return nil
}

// ExistByID 存在 by id
func (u *User) ExistByID() (bool, error) {
	return models.ExistUserByID(u.ID)
}

// 验证 用户名类型
func (u *User) validUserName(maps map[string]interface{}) (map[string]interface{}, error) {
	if valid.Email(u.UserName) {
		maps["login_email"] = u.UserName
		return maps, nil
	}
	if reg.Phone(u.UserName) {
		maps["login_phone"] = u.UserName
		return maps, nil
	}
	if reg.UserName(u.UserName) {
		maps["login_name"] = u.UserName
		return maps, nil
	}
	return nil, errors.New(e.GetMsg(e.ERROR_USER_REG_NAME))
}
