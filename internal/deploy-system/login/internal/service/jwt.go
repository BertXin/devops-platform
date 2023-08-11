package service

import (
	"devops-platform/internal/deploy-system/login/internal/domain"
	"github.com/form3tech-oss/jwt-go"
	"github.com/sirupsen/logrus"
)

func GenerateToken(claims *domain.TokenClaims) (token string, err error) {
	// 1. 定义token中的数据
	//withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"id":       claims.ID,
	//	"username": claims.Username,
	//	"exp":      claims.Exp,
	//})
	withClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 2. 使用秘钥签名token
	token, err = withClaims.SignedString([]byte("secret"))
	if err != nil {
		logrus.Error("请求API出错：", err)
		return
	}
	return
}
