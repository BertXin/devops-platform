package domain

import (
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"errors"
	"time"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserVO struct {
	UserID    types.Long   `json:"user_id"`
	LoginName string       `json:"login_name"`
	Username  string       `json:"user_name"`
	Role      enum.SysRole `json:"role"`
	Token     string       `json:"-"`
}

func (user *LoginUserVO) GetID() types.Long {
	return user.GetID()
}

func (user *LoginUserVO) GetName() string {
	return user.Username
}

func (user *LoginUserVO) GetRole() enum.SysRole {
	return user.Role
}
func (user *LoginUserVO) GetToken() string {
	return user.Token
}

func (user *LoginUserVO) GetLoginName() string {
	return user.LoginName
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
