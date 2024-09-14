package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"devops-platform/internal/deploy-system/permissions/internal/domain"
	"devops-platform/pkg/types"
)

type Repository struct {
	repository.Repository
}

func (r *Repository) CreatePerm(ctx context.Context, perm *domain.Permission) (err error) {
	result := r.DB(ctx).Create(perm)
	err = result.Error
	return
}

func (r *Repository) DeletePerm(ctx context.Context, ID types.Long) (err error) {
	result := r.DB(ctx).Delete(domain.Permission{}, ID)
	err = result.Error
	return
}

func (r *Repository) UpdatePerm(ctx context.Context, perm *domain.Permission) (err error) {
	result := r.DB(ctx).Updates(perm)
	err = result.Error
	return
}

func (r *Repository) FindPermByPID(ctx context.Context, PID types.Long) (perm *domain.Permission, err error) {
	result := r.DB(ctx).Where("p_id = ?", PID).First(&perm)
	err = result.Error
	return
}

func (r *Repository) Count(ctx context.Context, condition interface{}) (int64, error) {
	var count int64
	result := r.DB(ctx).Model(domain.Permission{}).Where(condition).Count(&count)
	return count, result.Error
}
