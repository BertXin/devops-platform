package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"devops-platform/internal/deploy-system/department/internal/domain"
	"devops-platform/pkg/types"
	"gorm.io/gorm"
)

// DeptRepository 定义部门数据的数据库操作
type Repository struct {
	repository.Repository
}

// GetByID 获取部门信息通过ID
func (r *Repository) GetByID(ctx context.Context, ID types.Long) (dept *domain.Dept, err error) {
	dept = new(domain.Dept)
	err = r.DB(ctx).First(dept, ID).Error
	if dept.ID != ID {
		dept = nil
		err = gorm.ErrRecordNotFound
	}
	return
}

// Create 创建部门信息
func (r *Repository) Create(ctx context.Context, dept *domain.Dept) (err error) {
	result := r.DB(ctx).Create(dept)
	err = result.Error
	return
}

// Save 保存或更新部门信息
func (r *Repository) Save(ctx context.Context, dept *domain.Dept) (err error) {
	result := r.DB(ctx).Save(dept)
	err = result.Error
	return
}

// FindByName 查询具有特定名称的部门列表
func (r *Repository) FindByName(ctx context.Context, name string, parentid types.Long, pagination types.Pagination) (depts []domain.Dept, total int64, err error) {
	tx := r.DB(ctx).Debug()
	if name != "" {
		tx = tx.Where("name LIKE ?", name)
	}
	if parentid != 0 {
		tx = tx.Where("parent_id = ?", parentid)
	}
	err = tx.Limit(pagination.Limit()).Offset(pagination.Offset()).Find(&depts).Error
	total, err = r.Count(ctx, tx)
	return
}

// List 查询所有部门列表
func (r *Repository) List(ctx context.Context, pagination types.Pagination) (depts []domain.Dept, err error) {
	err = r.DB(ctx).Limit(pagination.Limit()).Offset(pagination.Offset()).Find(&depts).Error
	return
}

// Delete 删除部门信息
func (r *Repository) Delete(ctx context.Context, ID types.Long) (err error) {
	result := r.DB(ctx).Delete(domain.Dept{}, ID)
	err = result.Error
	return
}

// Update 更新部门信息
func (r *Repository) Update(ctx context.Context, dept *domain.Dept) (err error) {
	result := r.DB(ctx).Updates(dept)
	err = result.Error
	return
}

// FindByParentID 根据父部门ID查询子部门列表
func (r *Repository) FindByParentID(ctx context.Context, parentID types.Long, pagination types.Pagination) (depts []domain.Dept, err error) {
	err = r.DB(ctx).Where("parent_id = ?", parentID).Limit(pagination.Limit()).Offset(pagination.Offset()).Find(&depts).Error
	return
}

// Count 统计部门数量
func (r *Repository) Count(ctx context.Context, condition interface{}) (int64, error) {
	var count int64
	result := r.DB(ctx).Model(domain.Dept{}).Where(condition).Count(&count)
	return count, result.Error
}
