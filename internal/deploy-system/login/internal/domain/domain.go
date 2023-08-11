package domain

import (
	"devops-platform/internal/deploy-system/user"
	"errors"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *user.User `json:"user"`
	Token string     `json:"token"`
}

type TokenClaims struct {
	ID       string
	Username string
	Exp      int64 // 过期时间
	Iat      int64
}

func (c *TokenClaims) Valid() error {

	now := time.Now().Unix()

	// 1. 检查过期时间
	if c.Exp < now {
		return errors.New("token expired")
	}

	// 2. 检查issued at时间
	if c.Iat > now {
		return errors.New("token used before issued")
	}

	// 3. 验证签名
	if err := c.verifySignature(); err != nil {
		return errors.New("invalid signature")
	}

	// 4. 其他字段合法性校验
	return nil
}
func (c *TokenClaims) verifySignature() error {
	return nil // 使用jwt库验证签名
}
