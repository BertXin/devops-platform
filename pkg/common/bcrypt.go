package common

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

//SetPassword加密password
func SetPassword(password string) (hashString string) {
	// 生成密码哈希
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 转换为字符串
	hashString = string(hash)
	return
}

//ValidatePassword解密password
func ValidatePassword(encryptedPassword string, password string) bool {
	// 2. 使用bcrypt进行密码匹配
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))

	// 3. 返回错误或nil
	if err != nil {
		logrus.Error("密码不正确", err)
		return false
	}
	return true
}
