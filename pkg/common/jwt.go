package common

import (
	"devops-platform/internal/deploy-system/login"
	"devops-platform/pkg/types"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserID types.Long `json:"user_id"`
	Name   string     `json:"user_name"`
	jwt.StandardClaims
}

func GenerateToken(user *login.LoginUser) (token string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserID: user.UserID,
		Name:   user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	// 2. 使用HS256算法进行签名
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tokenClaims.SignedString([]byte("secret"))

	return
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
