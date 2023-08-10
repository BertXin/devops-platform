package domain

import (
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"github.com/Nerzal/gocloak/v13"
)

type BaseR struct {
	Code int32 `json:"code"`

	Msg string `json:"msg"`
}

/*
 登录返回的TOKEN对象
*/
type SsoTokenVO struct {
	// 认证成功的凭证
	AccessToken string `json:"access_token"`
	IDToken     string `json:"id_token"`
	// 凭证有效期，单位：秒
	ExpiresIn int `json:"expires_in"`
	// 用于刷新的凭证有效期，单位：秒
	RefreshExpiresIn int `json:"refresh_expires_in"`
	// 用于刷新的凭证
	RefreshToken string `json:"refresh_token"`
	// 凭证类型
	TokenType       string `json:"token_type"`
	NotBeforePolicy int    `json:"not-before-policy"`
	SessionState    string `json:"session_state"`
	// 凭证权限范围
	Scope string `json:"scope"`
}

type SsoCheckTokenVO struct {
	Name string `json:"name"`
}

func ToUser(userinfo *gocloak.UserInfo) *SsoCheckTokenVO {
	return &SsoCheckTokenVO{Name: *userinfo.Name}
}

type LoginUserVO struct {
	UserID    types.Long   `json:"user_id"`
	LoginName string       `json:"login_name"`
	Username  string       `json:"user_name"`
	Role      enum.SysRole `json:"role"`
	Token     string       `json:"-"`
}

func (user *LoginUserVO) GetID() types.Long {
	return user.UserID
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

type UserDetail struct {
	UserID types.Long `json:"userId"`

	LoginName string `json:"username"`

	Password string `json:"password"`

	UserName string `json:"realname"`

	Phone string `json:"phone"`

	Email string `json:"email"`

	Avatar string `json:"avatar"`

	WxUserId string `json:"wxUserId"`

	Gender int8 `json:"gender"`

	UserType int8 `json:"userType"`

	OrgName string `json:"orgName"`
}

type UserDetailVO struct {
	BaseR
	Data struct {
		SysUser UserDetail `json:"sysUser"`
	} `json:"data"`
}
