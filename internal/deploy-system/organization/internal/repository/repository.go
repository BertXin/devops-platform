package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"devops-platform/internal/deploy-system/organization/internal/domain"
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Repository 部门仓储实现
type Repository struct {
	repository.Repository
}

// NewRepository 创建仓储实例
func NewRepository() *Repository {
	return &Repository{}
}

// GetDepartmentByID 根据ID获取部门
func (r *Repository) GetDepartmentByID(ctx context.Context, id types.Long) (*domain.Department, error) {
	var department domain.Department
	err := r.DB(ctx).First(&department, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &department, nil
}

// GetDepartmentByCode 根据编码获取部门
func (r *Repository) GetDepartmentByCode(ctx context.Context, code string) (*domain.Department, error) {
	var department domain.Department
	err := r.DB(ctx).Where("code = ?", code).First(&department).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &department, nil
}

// ListDepartments 获取部门列表
func (r *Repository) ListDepartments(ctx context.Context, query *domain.DepartmentQuery) ([]*domain.Department, int64, error) {
	db := r.DB(ctx).Model(&domain.Department{})

	// 应用查询条件
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Code != "" {
		db = db.Where("code LIKE ?", "%"+query.Code+"%")
	}
	if query.Status > 0 {
		db = db.Where("status = ?", query.Status)
	}
	if query.ParentID > 0 {
		db = db.Where("parenid = ?", query.ParentID)
	}

	// 获取总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页
	page := query.Page
	if page <= 0 {
		page = 1
	}
	size := query.Size
	if size <= 0 {
		size = 10
	}

	db = db.Offset((page - 1) * size).Limit(size)
	db = db.Order("sort ASC, id ASC")

	// 查询数据
	var departments []*domain.Department
	err = db.Find(&departments).Error
	if err != nil {
		return nil, 0, err
	}

	return departments, total, nil
}

// SaveDepartment 保存部门
func (r *Repository) SaveDepartment(ctx context.Context, department *domain.Department) error {
	return r.DB(ctx).Save(department).Error
}

// DeleteDepartment 删除部门
func (r *Repository) DeleteDepartment(ctx context.Context, id types.Long) error {
	// 开启事务
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查是否有子部门
		var count int64
		err := tx.Model(&domain.Department{}).Where("parenid = ?", id).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("存在子部门，无法删除")
		}

		// 删除部门
		if err := tx.Delete(&domain.Department{}, id).Error; err != nil {
			return err
		}

		// 删除用户部门关联
		if err := tx.Where("department_id = ?", id).Delete(&domain.UserDepartment{}).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetAllDepartments 获取所有部门
func (r *Repository) GetAllDepartments(ctx context.Context) ([]*domain.Department, error) {
	var departments []*domain.Department
	err := r.DB(ctx).
		Where("status = ?", enum.StatusEnabled).
		Order("sort ASC, id ASC").
		Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// ============= 用户部门关联 =============

// GetUserDepartments 获取用户所属部门
func (r *Repository) GetUserDepartments(ctx context.Context, userID types.Long) ([]*domain.Department, error) {
	var departments []*domain.Department
	err := r.DB(ctx).Table("department").
		Joins("JOIN user_department ON user_department.department_id = department.id").
		Where("user_department.user_id = ? AND department.status = ?", userID, enum.StatusEnabled).
		Find(&departments).Error
	if err != nil {
		return nil, err
	}
	return departments, nil
}

// AssignDepartmentToUser 为用户分配部门
func (r *Repository) AssignDepartmentToUser(ctx context.Context, userID types.Long, departmentID types.Long, isLeader bool) error {
	// 检查部门是否存在
	department, err := r.GetDepartmentByID(ctx, departmentID)
	if err != nil {
		return err
	}
	if department == nil {
		return errors.New("部门不存在")
	}

	// 开启事务
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查用户是否已经在该部门
		var userDept domain.UserDepartment
		err := tx.Where("user_id = ? AND department_id = ?", userID, departmentID).First(&userDept).Error
		if err == nil {
			// 已存在关联，更新关系类型
			userDept.Type = domain.UserDeptTypeNormal
			if isLeader {
				userDept.Type = domain.UserDeptTypeLeader
			}
			return tx.Save(&userDept).Error
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			// 其他错误
			return err
		}

		// 不存在关联，创建新关联
		userDept = domain.UserDepartment{
			UserID:       userID,
			DepartmentID: departmentID,
			Type:         domain.UserDeptTypeNormal,
			CreatedAt:    time.Now(),
		}
		if isLeader {
			userDept.Type = domain.UserDeptTypeLeader
		}

		return tx.Create(&userDept).Error
	})
}

// RemoveDepartmentFromUser 移除用户部门
func (r *Repository) RemoveDepartmentFromUser(ctx context.Context, userID types.Long, departmentID types.Long) error {
	return r.DB(ctx).Where("user_id = ? AND department_id = ?", userID, departmentID).Delete(&domain.UserDepartment{}).Error
}

// ListDepartmentUsers 获取部门用户列表
func (r *Repository) ListDepartmentUsers(ctx context.Context, departmentID types.Long, query *domain.UserQuery) ([]*domain.UserVO, int64, error) {
	// 用于查询的自定义结构体
	type UserWithType struct {
		ID        types.Long
		Username  string
		RealName  string
		Email     string
		Mobile    string
		Status    enum.Status
		CreatedAt time.Time
		Type      int
	}

	db := r.DB(ctx).Table("users").
		Select("users.id, users.username, users.nickname, users.email, users.phone, users.status, users.created_at, user_department.type").
		Joins("JOIN user_department ON user_department.user_id = users.id").
		Where("user_department.department_id = ?", departmentID)

	// 应用查询条件
	if query.Username != "" {
		db = db.Where("users.username LIKE ?", "%"+query.Username+"%")
	}
	if query.RealName != "" {
		db = db.Where("users.nickname LIKE ?", "%"+query.RealName+"%")
	}
	if query.Status > 0 {
		db = db.Where("users.status = ?", query.Status)
	}

	// 获取总数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页
	page := query.Page
	if page <= 0 {
		page = 1
	}
	size := query.Size
	if size <= 0 {
		size = 10
	}

	db = db.Offset((page - 1) * size).Limit(size)
	db = db.Order("users.id ASC")

	// 查询数据
	var users []UserWithType
	err = db.Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	// 转换为VO
	userVOs := make([]*domain.UserVO, 0, len(users))
	for _, user := range users {
		userVOs = append(userVOs, &domain.UserVO{
			ID:        user.ID,
			Username:  user.Username,
			RealName:  user.RealName,
			Email:     user.Email,
			Mobile:    user.Mobile,
			Status:    user.Status,
			IsLeader:  user.Type == domain.UserDeptTypeLeader,
			CreatedAt: user.CreatedAt,
		})
	}

	return userVOs, total, nil
}

// GetDepartmentUserType 获取用户在部门中的类型
func (r *Repository) GetDepartmentUserType(ctx context.Context, userID types.Long, departmentID types.Long) (int, error) {
	var userDept domain.UserDepartment
	err := r.DB(ctx).Where("user_id = ? AND department_id = ?", userID, departmentID).First(&userDept).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.UserDeptTypeNormal, nil
		}
		return 0, err
	}
	return userDept.Type, nil
}
