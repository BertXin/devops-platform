package common

import (
	"devops-platform/internal/deploy-system/login"
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("secret")

// Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	UserID types.Long
	Name   string
	Role   enum.SysRole
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(user *login.LoginUser) (token string, err error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		UserID: user.UserID,
		Name:   user.Username,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			// 指定token发行人
			Issuer:  "devops-system",
			Subject: "user token",
		},
	}
	// 2. 使用HS256算法进行签名
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err = tokenClaims.SignedString(jwtSecret)
	return
}

// ParseToken 验证token
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func CheckRefreshToken(token string) error {
	// 1. 尝试解析refresh token
	claims, err := ParseToken(token)
	if err != nil {
		return err
	}
	// 2. 检查token是否过期
	expTime := time.Unix(claims.ExpiresAt, 0)
	if time.Now().After(expTime) {
		return errors.New("refresh token expired")
	}
	return nil
}

func GenerateNewAccessToken(refreshToken string) (token string, err error) {
	// 1. 检验refresh token有效性
	if err := CheckRefreshToken(refreshToken); err != nil {
		return "", err
	}
	claims, _ := ParseToken(refreshToken)

	user := &login.LoginUser{
		UserID: claims.UserID,
	}
	token, err = GenerateToken(user)
	return
}
