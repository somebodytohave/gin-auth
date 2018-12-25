package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sun-wenming/gin-auth/pkg/logging"
	"github.com/sun-wenming/gin-auth/pkg/setting"
	"time"
)

var (
	// DB 数据库
	DB *gorm.DB
)

// Model 基类
type Model struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	// * 代表 null
	DeletedAt *time.Time `json:"deleted_at"`
}

// Setup gorm db 初始化
func Setup() {
	var err error
	DB, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		fmt.Println(err.Error())
		logging.Error(err.Error())
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	DB.LogMode(setting.DatabaseSetting.LogMode)
}

// CloseDB 关闭
func CloseDB() {
	defer DB.Close()
}
