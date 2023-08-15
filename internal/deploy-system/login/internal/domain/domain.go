package domain

import (
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
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
