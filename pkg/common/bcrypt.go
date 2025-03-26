package common

import (
	"golang.org/x/crypto/bcrypt"
)

// SetPassword加密password
func HashPassword(password string) (hashString string) {
	// 生成密码哈希
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 转换为字符串
	hashString = string(hash)
	return
}

type PasswordMismatchError struct{}

func (e *PasswordMismatchError) Error() string {
	return "password mismatch"
}

// ValidatePassword解密password
func ValidatePassword(encryptedPassword string, password string) error {
	// 2. 使用bcrypt进行密码匹配
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	//定义密码验证错误的返回
	if err != nil {
		return &PasswordMismatchError{}
	}
	return nil
}
