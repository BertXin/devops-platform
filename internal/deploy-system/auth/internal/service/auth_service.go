package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/internal/deploy-system/auth/internal/repository"
	"devops-platform/internal/pkg/common"
	"devops-platform/internal/pkg/security"
	"devops-platform/pkg/common/jwt"
	"devops-platform/pkg/types"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// AuthService 认证服务实现
type AuthService struct {
	service.Service
	Repo   *repository.Repository `inject:"AuthRepository"`
	Logger *logrus.Logger         `inject:"Logger"`
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

// Login 用户登录
func (s *AuthService) Login(ctx context.Context, username, password, ip, userAgent string) (*domain.TokenInfo, error) {
	// 查找用户
	user, err := s.Repo.GetByUsername(ctx, username)
	if err != nil {
		s.Logger.WithError(err).Error("查询用户失败")
		return nil, common.InternalError("系统内部错误", err)
	}

	if user == nil {
		// 记录登录失败日志
		s.saveLoginLog(ctx, 0, username, domain.LoginTypePassword, 0, "用户不存在", ip, userAgent)
		return nil, common.UnauthorizedError("用户名或密码错误", nil)
	}

	// 验证密码
	if !user.VerifyPassword(password) {
		// 记录登录失败日志
		s.saveLoginLog(ctx, user.ID, user.Username, domain.LoginTypePassword, 0, "密码错误", ip, userAgent)
		return nil, common.UnauthorizedError("用户名或密码错误", nil)
	}

	// 检查用户状态
	if user.Status != 1 {
		s.saveLoginLog(ctx, user.ID, user.Username, domain.LoginTypePassword, 0, "账户已禁用", ip, userAgent)
		return nil, common.ForbiddenError("账户已被禁用", nil)
	}

	// 生成JWT令牌
	tokenString, expireTime, err := s.generateToken(user)
	if err != nil {
		s.Logger.WithError(err).WithField("userId", user.ID).Error("生成令牌失败")
		return nil, common.InternalError("生成令牌失败", err)
	}

	// 创建事务上下文
	txCtx, err := s.Tx(ctx)
	if err != nil {
		return nil, common.InternalError("创建事务失败", err)
	}
	defer func() {
		// 根据err状态提交或回滚事务
		if err != nil {
			s.RollbackTx(txCtx, err, "auth service login")
		} else {
			s.CommitTx(txCtx, "auth service login")
		}
	}()

	// 更新最后登录时间
	user.UpdateLastLogin()
	err = s.Repo.Save(txCtx, user)
	if err != nil {
		s.Logger.WithError(err).WithField("userId", user.ID).Error("更新登录时间失败")
		return nil, common.InternalError("更新登录时间失败", err)
	}

	// 记录登录成功日志
	s.saveLoginLog(txCtx, user.ID, user.Username, domain.LoginTypePassword, 1, "登录成功", ip, userAgent)

	// 构建令牌信息
	tokenInfo := &domain.TokenInfo{
		Token:    tokenString,
		ExpireAt: expireTime,
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Nickname,
		Role:     int(user.RoleID),
	}

	return tokenInfo, nil
}

// generateToken 生成JWT令牌
func (s *AuthService) generateToken(user *domain.User) (string, time.Time, error) {
	// 设置过期时间
	expireTime := time.Now().Add(3 * time.Hour)

	claims := jwt.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Name:     user.Nickname,
		Role:     int(user.RoleID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "devops-platform",
			Subject:   user.Username,
		},
	}

	tokenString, err := jwt.GenerateToken(claims)
	return tokenString, expireTime, err
}

// Logout 用户登出
func (s *AuthService) Logout(ctx context.Context, userID types.Long) error {
	// 可以实现令牌黑名单机制，使当前令牌失效
	// 这里简化处理，仅记录日志
	s.Logger.WithField("userId", userID).Info("用户登出")
	return nil
}

// GetUserInfo 获取用户信息
func (s *AuthService) GetUserInfo(ctx context.Context, userID types.Long) (*domain.UserInfo, error) {
	user, err := s.Repo.GetByID(ctx, userID)
	if err != nil {
		return nil, common.InternalError("查询用户失败", err)
	}

	if user == nil {
		return nil, common.NotFoundError("用户不存在", nil)
	}

	// 这里可以通过deptService获取部门名称，简化处理返回空
	deptName := ""

	return user.ToUserInfo(deptName), nil
}

// VerifyToken 验证Token
func (s *AuthService) VerifyToken(ctx context.Context, token string) (*security.TokenInfo, error) {
	// 验证令牌
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return nil, common.UnauthorizedError("无效的令牌", err)
	}

	// 检查是否过期
	if jwt.IsTokenExpired(claims) {
		return nil, common.UnauthorizedError("令牌已过期", nil)
	}

	// 构建令牌信息
	tokenInfo := &security.TokenInfo{
		Token:     token,
		ExpireAt:  time.Unix(claims.ExpiresAt.Unix(), 0),
		UserID:    claims.UserID,
		Username:  claims.Username,
		RealName:  claims.Name,
		Role:      claims.Role,
		LoginTime: time.Unix(claims.IssuedAt.Unix(), 0),
	}

	return tokenInfo, nil
}

// ChangePassword 修改密码
func (s *AuthService) ChangePassword(ctx context.Context, userID types.Long, oldPassword, newPassword string) (err error) {

	// 查找用户
	user, err := s.Repo.GetByID(ctx, userID)
	if err != nil {
		return common.InternalError("查询用户失败", err)
	}

	if user == nil {
		return common.NotFoundError("用户不存在", nil)
	}

	// 验证旧密码
	if !user.VerifyPassword(oldPassword) {
		return common.UnauthorizedError("原密码错误", nil)
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "auth service change password")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "auth service change password")
	}()

	// 设置新密码
	err = user.SetPassword(newPassword)
	if err != nil {
		return common.InternalError("设置密码失败", err)
	}

	// 更新用户
	user.AuditModified(ctx)
	return s.Repo.Save(ctx, user)
}

// RegisterUser 注册用户
func (s *AuthService) RegisterUser(ctx context.Context, command *domain.CreateUserCommand) (id types.Long, err error) {
	// 检查用户名是否已存在
	exists, err := s.Repo.ExistsByUsername(ctx, command.Username)
	if err != nil {
		return 0, common.InternalError("检查用户名失败", err)
	}

	if exists {
		return 0, common.RequestParamError("用户名已存在", nil)
	}
	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "auth service change password")
	if err != nil {
		return 0, err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "auth service change password")
	}()

	// 转换为用户实体
	user, err := command.ToUser()
	if err != nil {
		return 0, common.RequestParamError(err.Error(), err)
	}

	// 添加审计信息
	user.AuditCreated(ctx)

	// 保存用户
	err = s.Repo.Save(ctx, user)
	if err != nil {
		return 0, common.InternalError("保存用户失败", err)
	}

	return user.ID, nil
}

// saveLoginLog 保存登录日志
func (s *AuthService) saveLoginLog(ctx context.Context, userID types.Long, username, loginType string, status int, message, ip, userAgent string) {
	// 创建登录日志
	log := &domain.LoginLog{
		UserID:    userID,
		Username:  username,
		IP:        ip,
		UserAgent: userAgent,
		LoginType: loginType,
		Status:    status,
		Message:   message,
		CreatedAt: time.Now(),
	}

	// 保存日志
	err := s.Repo.SaveLoginLog(ctx, log)
	if err != nil {
		s.Logger.WithError(err).WithFields(map[string]interface{}{
			"userID":   userID,
			"username": username,
			"message":  message,
		}).Error("保存登录日志失败")
	}
}

// userQuery 用户查询服务实现
type userQuery struct {
	repo   repository.Repository
	logger logrus.Logger
}

// GetByUsername 根据用户名查询用户
func (q *userQuery) GetByUsername(ctx context.Context, username string) (*domain.UserInfo, error) {
	user, err := q.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	// 这里可以通过deptService获取部门名称，简化处理返回空
	deptName := ""

	return user.ToUserInfo(deptName), nil
}

// GetByID 根据ID查询用户
func (q *userQuery) GetByID(ctx context.Context, ID types.Long) (*domain.UserInfo, error) {
	user, err := q.repo.GetByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	// 这里可以通过deptService获取部门名称，简化处理返回空
	deptName := ""

	return user.ToUserInfo(deptName), nil
}
