package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

// 默认的JWT签名密钥
var defaultSecret = []byte("devops-platform-jwt-secret")

// GenerateToken 使用默认密钥生成JWT令牌
func GenerateToken(claims Claims) (string, error) {
	return GenerateTokenWithSecret(claims, nil)
}

// GenerateTokenWithClaims 使用自定义声明和密钥生成JWT令牌
func GenerateTokenWithClaims(claims Claims, secretKey string) (string, error) {
	var secret []byte
	if secretKey == "" {
		secret = defaultSecret
	} else {
		secret = []byte(secretKey)
	}

	return GenerateTokenWithSecret(claims, secret)
}

// GenerateTokenWithSecret 使用指定的密钥生成JWT令牌
func GenerateTokenWithSecret(claims Claims, secret []byte) (string, error) {
	if secret == nil || len(secret) == 0 {
		secret = defaultSecret
	}

	// 创建JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名生成完整的JWT
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken 使用默认密钥解析JWT令牌
func ParseToken(tokenString string) (*Claims, error) {
	return ParseTokenWithSecret(tokenString, nil)
}

// ParseTokenWithKey 使用指定密钥解析JWT令牌
func ParseTokenWithKey(tokenString string, secretKey string) (*Claims, error) {
	var secret []byte
	if secretKey == "" {
		secret = defaultSecret
	} else {
		secret = []byte(secretKey)
	}

	return ParseTokenWithSecret(tokenString, secret)
}

// ParseTokenWithSecret 使用指定的密钥解析JWT令牌
func ParseTokenWithSecret(tokenString string, secret []byte) (*Claims, error) {
	if secret == nil || len(secret) == 0 {
		secret = defaultSecret
	}

	// 验证令牌格式
	parts := strings.Split(tokenString, ".")
	if len(parts) != 3 {
		return nil, errors.New("令牌格式错误")
	}

	// 解析令牌
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("不支持的签名算法")
		}
		return secret, nil
	})

	if err != nil {
		logrus.WithError(err).Error("解析JWT令牌失败")
		return nil, err
	}

	// 验证令牌有效性
	if !token.Valid {
		return nil, errors.New("无效的令牌")
	}

	return claims, nil
}

// SetSecret 设置默认的JWT签名密钥
func SetSecret(secret string) {
	if secret != "" {
		defaultSecret = []byte(secret)
	}
}

// IsTokenExpired 检查令牌是否已过期
func IsTokenExpired(claims *Claims) bool {
	if claims == nil || claims.ExpiresAt == nil {
		return true
	}

	return time.Now().After(time.Unix(claims.ExpiresAt.Unix(), 0))
}
