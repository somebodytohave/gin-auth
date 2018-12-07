package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"io"
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

var key = []byte("32-bytes-key-sui-ji-fhaso!nislfg")

// AesEncrypt Aes 加密
func AesEncrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// AesDecrypt Aes 解密
func AesDecrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
