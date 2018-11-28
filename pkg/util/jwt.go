package util

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mecm/gin-auth/pkg/setting"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

// Claims 声明
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成 token
func GenerateToken(username, password string) (string, error) {

	password, err := Encrypt(password)
	if err != nil {
		return "", err
	}

	// 现在的时间
	nowTime := time.Now()
	// 过期的时间
	expireTime := nowTime.Add(3 * time.Hour)
	// 初始化 声明
	claims := Claims{
		username, password, jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-auth",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 获取完整签名之后的 token
	return tokenClaims.SignedString(jwtSecret)
}

// ParseToken 解析 token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
