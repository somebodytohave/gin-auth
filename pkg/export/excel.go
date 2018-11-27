package export

import (
	"github.com/mecm/gin-blog/pkg/setting"
)

// GetExcelFullUrl 获取 访问 url
func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

// GetExcelPath 获取 Excel 路径
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath excel 保存路径
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
