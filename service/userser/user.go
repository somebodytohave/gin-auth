package userser

import (
	"encoding/json"
	"fmt"
	"github.com/sun-wenming/gin-auth/models/users"
	"github.com/sun-wenming/gin-auth/pkg/e"
	"github.com/sun-wenming/gin-auth/pkg/gredis"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/util"
	"github.com/sun-wenming/gin-auth/service/caches"
)

// User 用户
type User struct {
	ID uint

	UserName string
	Password string
	Code     string
}

// Register 注册用户
func (u *User) Register() util.Error {
	maps := make(map[string]interface{})
	maps, err := u.validUserName(maps)
	if err != nil {
		return err
	}

	// 加密
	password, err := util.Encrypt(u.Password)
	if err != nil {
		return err
	}

	maps["password"] = password

	// 创建 用户信息 与 用户密码
	return users.AddUserLogin(maps)
}

// PwdLogin 登录用户
func (u *User) PwdLogin() (string, util.Error) {

	user, err := u.getUserLoginInfo()
	if err != nil {
		return "", err
	}

	inputPwd := u.Password
	hashPwd := user.Password

	// 比较 密码
	if err := util.Compare(inputPwd, hashPwd); err != nil {
		return "", util.ErrNewCode(e.ErrorUserNameExist)
	}

	if err := existUserInfo(user.UserID); err != nil {
		return "", err
	}

	// 生成token
	token, merr := util.GenerateToken(u.UserName)
	if merr != nil {
		logging.GetLogger().Error(merr)
		return "", util.ErrNewCode(e.ErrorAuthGenerateToken)
	}

	return token, nil
}

// PhoneRegister 手机号注册
func (u *User) PhoneRegister() util.Error {

	maps := map[string]interface{}{
		"login_phone": u.UserName,
	}
	// 创建 用户信息 与 用户密码
	return users.AddUserLogin(maps)

}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo() (*users.User,  util.Error) {
	return users.GetUser(u.ID)

}

// SendCode 发送手机验证码
func SendCode(phone string) (string, util.Error) {
	var (
		code string
		err  error
	)
	code = GetCacheCode(phone)
	// 如果没有验证码,随机生成
	if code == "" {
		code = util.GetRandomCode()
	}

	cache := caches.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()

	// 发送验证码操作
	// 十分钟验证码缓存
	if err := gredis.Set(key, code, 600); err != nil {
		logging.GetLogger().Warn(caches.ErrorSet, err)
	}
	if err != nil {
		logging.GetLogger().Error(err)
		return "", util.ErrNewCode(e.ErrorPhoneCodeNotValid)
	}
	// 便于测试，code返回出去
	return code, nil
}

func (u *User) getUserLoginInfo() (*users.UserLogin, util.Error) {
	maps := make(map[string]interface{})
	maps, err := u.validUserName(maps)
	if err != nil {
		return nil, err
	}
	// 查询 用户登录信息
	user, err := users.LoginUserLogin(maps)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func existUserInfo(userID uint)  util.Error {
	// 匹配成功  根据 userID 查询 用户信息
	if !(userID > 0) {
		return util.ErrNewCode(e.ErrorUserGetInfo)
	}

	exist, err := users.ExistUserByID(userID)
	if err != nil {
		return err
	}

	if !exist {
		return util.ErrNewCode(e.ErrorUserInfoEmpty)
	}
	return nil
}

// GetCacheCode 获取缓存的验证码
func GetCacheCode(phone string) string {
	cache := caches.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()
	fmt.Println(key)
	if !gredis.Exists(key) {
		return ""
	}
	var code string
	data, err := gredis.Get(key)
	if err != nil {
		logging.GetLogger().Warn(caches.ErrorGet, err)
		return ""
	}
	json.Unmarshal(data, &code)
	return code
}

// ExistByUserName 是否存在用户账号
func (u *User) ExistByUserName() (bool, util.Error) {
	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return false, err
	}
	return users.ExistUserLogin(maps)
}

// UserLoginGetUserID 返回用户ID
func (u *User) UserLoginGetUserID() (uint, util.Error) {
	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return 0, err
	}
	return users.UserLoginGetUserID(maps)
}

// 验证 用户名类型
func (u *User) validUserName(maps map[string]interface{}) (map[string]interface{}, util.Error) {
	if util.ValidEmail(u.UserName) {
		maps["login_email"] = u.UserName
		return maps, nil
	}
	// 如果是手机号
	if util.RegPhone(u.UserName) {
		maps["login_phone"] = u.UserName
		return maps, nil
	}
	if util.RegUserName(u.UserName) {
		maps["login_name"] = u.UserName
		return maps, nil
	}
	return nil, util.ErrNewCode(e.ErrorUserRegName)
}
