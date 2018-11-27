package util

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt 加密
// inputPassword 未加密的密码
func Encrypt(inputPassword string) (string, error) {
	// Generate "hash" 加密密码
	hash, err := bcrypt.GenerateFromPassword([]byte(inputPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// 加密后的密码
	password := string(hash)
	return password, nil
}

// Compare 比较密码
func Compare(inputPwd, hashPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(inputPwd))
}
