package auth

import (
	"context"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/pkg/types"
)

//go:generate mockgen -source=export.go -destination=mock/mock_auth.go -package=mock

// Bean常量
const (
	BeanAuthService = domain.BeanService   // 认证服务Bean名称
	BeanUserQuery   = domain.BeanUserQuery // 用户查询服务Bean名称
)

// 认证服务接口
type AuthService interface {
	// Login 本地用户登录
	Login(ctx context.Context, username, password string) (*domain.TokenInfo, error)

	// Logout 用户登出
	Logout(ctx context.Context, userID types.Long) error

	// GetUserInfo 获取用户信息
	GetUserInfo(ctx context.Context, userID types.Long) (*domain.UserInfo, error)

	// VerifyToken 验证Token
	VerifyToken(ctx context.Context, token string) (*domain.TokenInfo, error)

	// ChangePassword 修改密码
	ChangePassword(ctx context.Context, userID types.Long, oldPassword, newPassword string) error

	// RegisterUser 注册用户
	RegisterUser(ctx context.Context, command *domain.CreateUserCommand) (types.Long, error)
}

// UserQuery 用户查询接口
type UserQuery interface {
	// GetByUsername 根据用户名查找用户
	GetByUsername(ctx context.Context, username string) (*domain.UserInfo, error)

	// GetByID 根据ID查找用户
	GetByID(ctx context.Context, ID types.Long) (*domain.UserInfo, error)
}

// 领域对象类型别名
type CreateUserCommand = domain.CreateUserCommand
type UserInfo = domain.UserInfo
type TokenInfo = domain.TokenInfo
type ChangePasswordCommand = domain.ChangePasswordCommand
