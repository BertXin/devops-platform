package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"errors"
	"gorm.io/gorm"
)

/**
 * 数据库操作
 * 必须包含repository.Repository
 */
type Repository struct {
	repository.Repository
}

/**
 * 根据ID获取用户信息
 */
//GetByID
func (r *Repository) GetByID(ctx context.Context, ID types.Long) (user *domain.User, err error) {
	user = new(domain.User)
	err = r.DB(ctx).First(user, ID).Error
	if user.ID != ID {
		user = nil
	}
	return
}

/**
 * 创建用户信息
 */
//Create
func (r *Repository) Create(ctx context.Context, user *domain.User) (err error) {
	user.Enable = enum.EnableStatus
	result := r.DB(ctx).Create(user)
	err = result.Error
	return
}

/**
 * 根据名称获取密码
 */
func (r *Repository) GetPasswordByUsername(ctx context.Context, username string) (password string, err error) {
	var user domain.User
	err = r.DB(ctx).Model(&user).Where("username = ?", username).Select("password").First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return user.Password, err
}

/**
 * 保存、更新用户信息
 */
func (r *Repository) Save(ctx context.Context, user *domain.User) (err error) {
	result := r.DB(ctx).Save(user)
	err = result.Error
	return
}

/**
 * 根据名称查找用户信息
 */
func (r *Repository) FindByNameAndMobile(ctx context.Context, id types.Long, name string, mobile string, enable int64, pagination types.Pagination) (users []domain.User, total int64, err error) {

	tx := r.DB(ctx).Debug()
	if id != 0 {
		tx = tx.Where("id = ?", id)
	}
	if name != "" {
		tx = tx.Where("name like ?", "%"+name+"%")
	}
	if mobile != "" {
		tx = tx.Where("mobile like ?", "%"+mobile+"%")
	}
	if enable != 0 {
		tx = tx.Where("enable = ?", enable)
	}

	err = tx.Model(&domain.User{}).Count(&total).Error
	if total == 0 || err != nil {
		return
	}
	err = tx.Limit(pagination.Limit()).Offset(pagination.Offset()).Find(&users).Error
	return
}

/**
 * 根据用户名获取用户信息
 */
func (r *Repository) GetByUsername(ctx context.Context, username string) (user *domain.User, err error) {
	user = &domain.User{
		Username: username,
	}
	if err := r.DB(ctx).Where(user).First(user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
