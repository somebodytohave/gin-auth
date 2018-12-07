package user_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mecm/gin-auth/models"
	"github.com/mecm/gin-auth/pkg/e"
	"github.com/mecm/gin-auth/pkg/gredis"
	"github.com/mecm/gin-auth/pkg/logging"
	"github.com/mecm/gin-auth/pkg/util"
	"github.com/mecm/gin-auth/pkg/util/reg"
	"github.com/mecm/gin-auth/pkg/util/valid"
	"github.com/mecm/gin-auth/service/cache_service"
)

// User 用户
type User struct {
	ID uint

	UserName string
	Password string
	Code     string
}

// Register 注册用户
func (u *User) Register() error {
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
	return models.AddUserLogin(maps)
}

// Login 登录用户
func (u *User) Login() error {

	user, err := u.getUserLoginInfo()
	if err != nil {
		return err
	}

	inputPwd := u.Password
	hashPwd := user.Password

	// 比较 密码
	if err := util.Compare(inputPwd, hashPwd); err != nil {
		return err
	}

	if err := existUserInfo(user.UserID); err != nil {
		return err
	}

	return nil
}

// PhoneRegister 手机号注册
func (u *User) PhoneRegister() error {

	maps := map[string]interface{}{
		"login_phone": u.UserName,
	}
	// 创建 用户信息 与 用户密码
	return models.AddUserLogin(maps)

}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo() (*models.User, error) {

	userLogin, err := u.getUserLoginInfo()
	if err != nil {
		return nil, err
	}
	// 是否存在用户
	if err := existUserInfo(userLogin.UserID); err != nil {
		return nil, err
	}

	user, err := models.GetUser(userLogin.UserID)
	if err != nil {
		return nil, err
	}
	return user, nil

}

// SendCode 发送手机验证码
func SendCode(phone string) (string, error) {
	var (
		code string
		err  error
	)
	code = GetCacheCode(phone)
	// 如果没有验证码,随机生成
	if code == "" {
		code = util.GetRandomCode()
	}

	cache := cache_service.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()

	// 发送验证码操作
	// 十分钟验证码缓存
	if err := gredis.Set(key, code, 600); err != nil {
		logging.Warn(e.CACHE_ERROR_SET, err)
	}
	// 便于测试，code返回出去
	return code, err
}

func (u *User) getUserLoginInfo() (*models.UserLogin, error) {
	maps := make(map[string]interface{})
	maps, err := u.validUserName(maps)
	if err != nil {
		return nil, err
	}
	// 查询 用户登录信息
	user, err := models.LoginUserLogin(maps)
	if err != nil {
		return nil, err
	}
	return user, err
}

func existUserInfo(userID uint) error {
	// 匹配成功  根据 userID 查询 用户信息
	if !(userID > 0) {
		return errors.New(e.GetMsg(e.ERROR_USER_GET_INFO))
	}

	exist, err := models.ExistUserByID(userID)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New(e.GetMsg(e.ERROR_USER_GET_INFO))
	}
	return nil
}

// GetCacheCode 获取缓存的验证码
func GetCacheCode(phone string) string {
	cache := cache_service.Phone{Phone: phone}
	key := cache.GetPhoneCodeKey()
	fmt.Println(key)
	if !gredis.Exists(key) {
		return ""
	}
	var code string
	data, err := gredis.Get(key)
	if err != nil {
		logging.Warn(e.CACHE_ERROR_GET, err)
		return ""
	}
	json.Unmarshal(data, &code)
	return code
}

// ExistByName 是否存在
func (u *User) ExistByName() (bool, error) {
	maps := make(map[string]interface{})

	maps, err := u.validUserName(maps)
	if err != nil {
		return true, err
	}
	return models.ExistUserLogin(maps)
}

// 验证 用户名类型
func (u *User) validUserName(maps map[string]interface{}) (map[string]interface{}, error) {
	if valid.Email(u.UserName) {
		maps["login_email"] = u.UserName
		return maps, nil
	}
	// 如果是手机号
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
