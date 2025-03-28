package domain

import (
	"devops-platform/internal/pkg/enum"
	"devops-platform/internal/pkg/module"
	"devops-platform/internal/pkg/security"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"fmt"
	"strings"
	"time"
)

// User 用户实体
type User struct {
	module.Module              // 包含ID、创建时间、更新时间、删除时间等基础字段
	UID             string     `json:"uid" gorm:"column:uid;comment:'用戶uid'"`
	Username        string     `json:"username" gorm:"uniqueIndex:uk_username;comment:'用户名'"`
	Password        string     `json:"-" gorm:"comment:'用户密码'"`
	Phone           string     `json:"phone" gorm:"comment:'手机号码'"`
	Email           string     `json:"email" gorm:"comment:'邮箱'"`
	Nickname        string     `json:"nickname" gorm:"comment:'用户昵称'"`
	Avatar          string     `json:"avatar" gorm:"default:https://www.dnsjia.com/luban/img/head.png;comment:'用户头像'"`
	Status          int8       `json:"status" gorm:"type:tinyint(1);default:1;comment:'用户状态(1:正常 0:禁用)'"`
	MFASecret       string     `json:"-" gorm:"column:mfa_secret;type:text;comment:'mfa密钥'"`
	RoleID          types.Long `json:"role_id" gorm:"comment:'角色id外键'"`
	DeptID          types.Long `json:"dept_id" gorm:"comment:'部门id外键'"`
	Title           string     `json:"title" gorm:"comment:'职位'"`
	CreateBy        string     `json:"create_by" gorm:"comment:'创建来源,ldap/local/dingtalk'"`
	PasswordUpdated *time.Time `json:"password_updated" gorm:"comment:'密码更新时间'"`
	UpdatedAt       *time.Time `json:"updatedAt" gorm:"comment:'最后登录时间'"`
}

// TableName 指定数据库表名
func (User) TableName() string {
	return "user"
}

// ToAuthUser 转换为认证用户
func (u *User) ToAuthUser(token string) *security.AuthUser {
	return &security.AuthUser{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Nickname,
		Role:     int(u.RoleID),
		DeptID:   u.DeptID,
		Token:    token,
	}
}

// UserInfo 用户信息DTO
func (u *User) ToUserInfo(deptName string) *UserInfo {
	return &UserInfo{
		ID:       u.ID,
		Username: u.Username,
		Name:     u.Nickname,
		Mobile:   u.Phone,
		Email:    u.Email,
		Avatar:   u.Avatar,
		DeptID:   u.DeptID,
		DeptName: deptName,
		RoleID:   u.RoleID,
		Status:   u.Status,
	}
}

// SetPassword 设置密码
func (u *User) SetPassword(plainPassword string) error {
	if plainPassword == "" {
		return errors.New("密码不能为空")
	}
	u.Password = common.HashPassword(plainPassword)
	now := time.Now()
	u.PasswordUpdated = &now
	return nil
}

// ValidatePassword 验证密码
func (u *User) ValidatePassword(plainPassword string) error {
	return common.ValidatePassword(u.Password, plainPassword)
}

// VerifyPassword 验证密码并返回布尔值
func (u *User) VerifyPassword(plainPassword string) bool {
	err := u.ValidatePassword(plainPassword)
	return err == nil
}

// UpdateLastLogin 更新最后登录时间
func (u *User) UpdateLastLogin() {
	// 不再需要单独的LastLogin字段，直接使用UpdatedAt
	now := time.Now()
	u.UpdatedAt = &now
}

// UserInfo 用户信息视图对象
type UserInfo struct {
	ID       types.Long `json:"id"`
	Username string     `json:"username"`
	Name     string     `json:"name"`
	Mobile   string     `json:"mobile"`
	Email    string     `json:"email"`
	Avatar   string     `json:"avatar"`
	DeptID   types.Long `json:"dept_id"`
	DeptName string     `json:"dept_name"`
	RoleID   types.Long `json:"role_id"`
	Status   int8       `json:"status"`
	Roles    []RoleInfo `json:"roles,omitempty"` // 用户角色列表
}

// RoleInfo 角色信息
type RoleInfo struct {
	ID   types.Long `json:"id"`
	Name string     `json:"name"`
	Code string     `json:"code"`
}

// TokenInfo 令牌信息
type TokenInfo struct {
	Token    string     `json:"token"`     // JWT令牌
	ExpireAt time.Time  `json:"expire_at"` // 过期时间
	UserID   types.Long `json:"user_id"`   // 用户ID
	Username string     `json:"username"`  // 用户名
	Name     string     `json:"name"`      // 用户姓名
	Role     int        `json:"role"`      // 用户角色
}

// LoginLog 登录日志
type LoginLog struct {
	ID        types.Long `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    types.Long `json:"user_id" gorm:"comment:'用户ID'"`
	Username  string     `json:"username" gorm:"comment:'用户名'"`
	IP        string     `json:"ip" gorm:"comment:'登录IP'"`
	UserAgent string     `json:"user_agent" gorm:"comment:'用户代理'"`
	LoginType string     `json:"login_type" gorm:"comment:'登录类型'"`
	Status    int        `json:"status" gorm:"comment:'状态 1成功 0失败'"`
	Message   string     `json:"message" gorm:"comment:'消息'"`
	CreatedAt time.Time  `json:"created_at" gorm:"comment:'创建时间'"`
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// CreateUserCommand 创建用户命令
type CreateUserCommand struct {
	Username string       `json:"username"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Mobile   string       `json:"mobile"`
	Email    string       `json:"email"`
	Avatar   string       `json:"avatar"`
	DeptID   types.Long   `json:"dept_id"`
	Role     enum.SysRole `json:"role"`
}

// ToUser 转换为用户实体
func (command *CreateUserCommand) ToUser() (*User, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}

	// 生成UID
	uid := generateUID()

	user := &User{
		UID:      uid,
		Username: command.Username,
		Nickname: command.Name,
		Phone:    command.Mobile,
		Email:    command.Email,
		Avatar:   command.Avatar,
		DeptID:   command.DeptID,
		RoleID:   types.Long(command.Role),
		Status:   1,       // 默认启用
		CreateBy: "local", // 本地创建
	}

	err = user.SetPassword(command.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// generateUID 生成用户唯一标识
func generateUID() string {
	return fmt.Sprintf("U%d", time.Now().UnixNano()/1000000)
}

// Validate 验证命令参数
func (command *CreateUserCommand) Validate() error {
	command.Username = strings.TrimSpace(command.Username)
	if command.Username == "" {
		return errors.New("用户名不能为空")
	}

	command.Password = strings.TrimSpace(command.Password)
	if command.Password == "" {
		return errors.New("密码不能为空")
	}

	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return errors.New("姓名不能为空")
	}

	command.Mobile = strings.TrimSpace(command.Mobile)
	command.Email = strings.TrimSpace(command.Email)

	if command.DeptID <= 0 {
		return errors.New("必须指定部门")
	}

	return nil
}

// ChangePasswordCommand 修改密码命令
type ChangePasswordCommand struct {
	UserID      types.Long `json:"-"`
	OldPassword string     `json:"old_password" binding:"required"`
	NewPassword string     `json:"new_password" binding:"required,min=6,max=20"`
}

// Validate 验证命令参数
func (command *ChangePasswordCommand) Validate() error {
	if command.OldPassword == "" {
		return errors.New("原密码不能为空")
	}

	if command.NewPassword == "" {
		return errors.New("新密码不能为空")
	}

	if command.NewPassword == command.OldPassword {
		return errors.New("新密码不能与旧密码相同")
	}

	return nil
}

// SessionUser 会话用户信息
type SessionUser struct {
	ID       types.Long
	Username string
	Token    string
}

// GetID 获取用户ID
func (u *SessionUser) GetID() types.Long {
	return u.ID
}

// GetUsername 获取用户名
func (u *SessionUser) GetUsername() string {
	return u.Username
}

// GetToken 获取用户令牌
func (u *SessionUser) GetToken() string {
	return u.Token
}

// GetName 实现security.User接口
func (u *SessionUser) GetName() string {
	return u.Username
}

// ToUserContext 将SessionUser转换为security.UserContext
func (u *SessionUser) ToUserContext() *security.UserContext {
	return &security.UserContext{
		UserID:      u.ID,
		Username:    u.Username,
		RealName:    u.Username,
		TokenString: u.Token,
		TokenInfo: &security.TokenInfo{
			Token:    u.Token,
			UserID:   u.ID,
			Username: u.Username,
		},
	}
}
