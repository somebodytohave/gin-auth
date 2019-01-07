package logging

import (
	"github.com/sirupsen/logrus"
	"github.com/sun-wenming/gin-auth/pkg/file"
	"github.com/sun-wenming/gin-auth/pkg/setting"
	"os"
)

var (
	f      *os.File
	logger *logrus.Logger
)

func Setup() {
	// Create a new instance of the logger. You can have any number of instances.
	logger = logrus.New()
	var err error
	//You could set this to any `io.Writer` such as a file
	filePath := getLogFilePath()
	fileName := getLogFileName()
	f, err = file.MustOpen(fileName, filePath)
	if err != nil {
		logger.Fatalln(err)
	}

	// 输出到文件中
	logger.Out = f

	// If you wish to add the calling method as a field
	logger.SetReportCaller(true)

	//logger.Formatter = new(logrus.JSONFormatter)

	//- Fatal：网站挂了，或者极度不正常
	//- Error：跟遇到的用户说对不起，可能有bug
	//- Warn：记录一下，某事又发生了
	//- Info：提示一切正常
	//- debug：没问题，就看看堆栈
	if setting.ServerSetting.RunMode == "release" {
		logger.Level = logrus.WarnLevel
	}

	//logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// GetLogger Logger
func GetLogger() *logrus.Logger {
	return logger
}

//func setPrefix(level Level) {
//	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
//	if ok {
//		logPrefix = fmt.Sprintf("[%s:%d]", filepath.Base(file), line)
//	} else {
//		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
//	}
//	log.SetPrefix(logPrefix)
//}
