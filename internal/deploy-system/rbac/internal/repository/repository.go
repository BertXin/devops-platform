package repository

import (
	"context"
	"devops-platform/internal/common/repository"

	"devops-platform/internal/deploy-system/rbac/internal/domain"
	"devops-platform/pkg/types"
)

type Repository struct {
	repository.Repository
}

func (r *Repository) CreateRole(ctx context.Context, role *domain.Role) (err error) {
	result := r.DB(ctx).Create(role)
	err = result.Error
	return
}

func (r *Repository) UpdateRole(ctx context.Context, role *domain.Role) (err error) {
	result := r.DB(ctx).Updates(role)
	err = result.Error
	return
}

func (r *Repository) DeleteRole(ctx context.Context, ID types.Long) (err error) {
	result := r.DB(ctx).Delete(domain.Role{}, ID)
	err = result.Error
	return
}
func (r *Repository) FindRoleByID(ctx context.Context, ID types.Long) (role *domain.Role, err error) {
	result := r.DB(ctx).First(&role, ID)
	err = result.Error
	return
}
func (r *Repository) FindRoleByName(ctx context.Context, name string, pagination types.Pagination) (roles []domain.Role, total int64, err error) {
	tx := r.DB(ctx).Debug()
	if name != "" {
		tx = tx.Where("name LIKE ?", "%"+name+"%")
	}
	err = tx.Model(&domain.Role{}).Count(&total).Limit(pagination.PageSize).Offset(pagination.Offset()).Find(&roles).Error
	total, err = r.Count(ctx, tx)
	return
}

func (r *Repository) Count(ctx context.Context, condition interface{}) (int64, error) {
	var count int64
	result := r.DB(ctx).Model(domain.Role{}).Where(condition).Count(&count)
	return count, result.Error
}
