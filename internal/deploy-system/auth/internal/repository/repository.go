package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/pkg/types"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Repository 认证仓储实现
type Repository struct {
	repository.Repository
}

func NewRepository() *Repository {
	return &Repository{}
}

// GetByUsername 根据用户名查找用户
func (r *Repository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	err := r.DB(ctx).Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// GetByID 根据ID查找用户
func (r *Repository) GetByID(ctx context.Context, ID types.Long) (*domain.User, error) {
	var user domain.User
	err := r.DB(ctx).First(&user, ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Save 保存用户
func (r *Repository) Save(ctx context.Context, user *domain.User) error {
	return r.DB(ctx).Save(user).Error
}

// UpdatePassword 更新密码
func (r *Repository) UpdatePassword(ctx context.Context, userID types.Long, hashedPassword string) error {
	now := time.Now()
	return r.DB(ctx).Model(&domain.User{}).Where("id = ?", userID).
		Updates(map[string]interface{}{
			"password":         hashedPassword,
			"password_updated": now,
		}).Error
}

// UpdateLastLogin 更新最后登录时间
func (r *Repository) UpdateLastLogin(ctx context.Context, userID types.Long) error {
	return r.DB(ctx).Model(&domain.User{}).Where("id = ?", userID).
		Update("updated_at", gorm.Expr("NOW()")).Error
}

// SaveLoginLog 保存登录日志
func (r *Repository) SaveLoginLog(ctx context.Context, log *domain.LoginLog) error {
	return r.DB(ctx).Create(log).Error
}

// ExistsByUsername 检查用户名是否存在
func (r *Repository) ExistsByUsername(ctx context.Context, username string) (bool, error) {
	var count int64
	err := r.DB(ctx).Model(&domain.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
