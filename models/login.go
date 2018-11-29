package models

// UserLogin 用户密码登陆认证
type UserLogin struct {
	ID         uint `gorm:"primary_key"`
	UserID     uint
	LoginName  string `gorm:"unique"`
	LoginEmail string `gorm:"unique"`
	LoginPhone string `gorm:"unique"`
	Password   string `gorm:"unique"`
	status     uint
}

// AddUserLogin 添加用户账号 与 初始化个人信息
func AddUserLogin(userProfile, userLogin map[string]interface{}) error {

	tx := db.Begin()

	// 首先创建 user
	userID, err := addUser(userProfile, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	loginInfo := UserLogin{
		UserID:    userID,
		LoginName: userLogin["username"].(string),
		Password:  userLogin["password"].(string),
	}
	if err := tx.Create(&loginInfo).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

// LoginUserLogin 采用密码方式登录
func LoginUserLogin(maps map[string]interface{}) (*UserLogin, error) {
	var user UserLogin
	if err := db.Where(maps).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

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
