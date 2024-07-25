package common

import (
	"devops-platform/internal/deploy-system/login"
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

var jwtSecret = []byte("secret")

type Claims struct {
	UserID               types.Long   `json:"user_id"`
	Name                 string       `json:"name"`
	Role                 enum.SysRole `json:"role"`
	jwt.RegisteredClaims              // 内嵌标准的声明
}

// GenToken 生成JWT
func GenerateToken(user *login.LoginUser) (token string, err error) {
	//	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		UserID: user.UserID,
		Name:   user.Username, // 自定义字段
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 定义过期时间
			IssuedAt:  jwt.NewNumericDate(nowTime),    // 签发时间
			NotBefore: jwt.NewNumericDate(nowTime),    // 生效时间
			// 指定token发行人
			Issuer:  "devops-system",
			Subject: "user token",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}
func ParseToken(token string) (*Claims, error) {
	//校验传入的token的格式是否正确
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		logrus.Error("token格式错误")
		return nil, errors.New("token格式错误")
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		logrus.Errorf("解析JWT失败: %+v", err)
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err

}

//func CheckRefreshToken(token string) error {
//	// 1. 尝试解析refresh token
//	claims, err := ParseToken(token)
//	if err != nil {
//		return err
//	}
//	// 2. 检查token是否过期
//	expTime := time.Unix(claims.ExpiresAt, 0)
//	if time.Now().After(expTime) {
//		return errors.New("refresh token expired")
//	}
//	return nil
//}
//
//func GenerateNewAccessToken(refreshToken string) (token string, err error) {
//	// 1. 检验refresh token有效性
//	if err := CheckRefreshToken(refreshToken); err != nil {
//		return "", err
//	}
//	claims, _ := ParseToken(refreshToken)
//
//	user := &login.LoginUser{
//		UserID: claims.UserID,
//	}
//	token, err = GenerateToken(user)
//	return
//}
