package qrcode

import (
	"image/jpeg"
	"github.com/boombuler/barcode"
	"github.com/mecm/gin-blog/pkg/file"
	"github.com/mecm/gin-blog/pkg/util"
	"github.com/boombuler/barcode/qr"
	"github.com/mecm/gin-blog/pkg/setting"
)

const (
	EXT_JPG = ".jpg"
)

// QrCode 二维码 实体类
type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

// NewQrCode 创建一个 qrode
func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

// CheckEncode 检查二维码是否存在
func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if file.CheckNotExist(src) == true {
		return false
	}

	return true
}

// Encode 返回带有名称和路径的二维码
func (q *QrCode) Encode(path string) (string, string, error) {
	// 二维码名字
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	// 保存的地址
	src := path + name
	// 检查是否存在
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}
		// 返回带有高度和宽度的条形码
		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}
		
		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}

func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl +"/"+ GetQrCodePath() + name
}

func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetQrCodePath()
}

