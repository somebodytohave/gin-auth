package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	Debug           bool
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	PrefixUrl      string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string
	LogSavePath    string
	LogSaveName    string
	LogFileExt     string
	TimeFormat     string
}

// AppSetting APP 配置
var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting 服务器配置
var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	LogMode     bool
}

// DatabaseSetting 数据库配置
var DatabaseSetting = &Database{}

type Redis struct {
	Host     string
	Password string
	// 最大空闲连接数
	MaxIdle int
	// 在给定时间内，允许分配的最大连接数（当为零时，没有限制）
	MaxActive int
	// 在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）
	IdleTimeout time.Duration
}

// RedisSetting Redis 缓存
var RedisSetting = &Redis{}

var cfg *ini.File

// Setup 初始化各项配置
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	mapTo("redis", RedisSetting)

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
