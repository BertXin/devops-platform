package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/pkg/types"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Repository 权限仓储实现
type Repository struct {
	repository.Repository
	enforcer domain.CasbinEnforcer
}

// NewRepository 创建仓储实例
func NewRepository() *Repository {
	return &Repository{}
}

// GetRoleByID 根据ID获取角色
func (r *Repository) GetRoleByID(ctx context.Context, id types.Long) (*domain.Role, error) {
	var role domain.Role
	err := r.DB(ctx).First(&role, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

// GetRoleByCode 根据编码获取角色
func (r *Repository) GetRoleByCode(ctx context.Context, code string) (*domain.Role, error) {
	var role domain.Role
	err := r.DB(ctx).Where("code = ?", code).First(&role).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &role, nil
}

// ListRoles 获取角色列表
func (r *Repository) ListRoles(ctx context.Context, query *domain.RoleQuery) ([]*domain.Role, int64, error) {
	db := r.DB(ctx).Model(&domain.Role{})

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
	db = db.Order("sort_order ASC, id ASC")

	// 查询数据
	var roles []*domain.Role
	err = db.Find(&roles).Error
	if err != nil {
		return nil, 0, err
	}

	return roles, total, nil
}

// SaveRole 保存角色
func (r *Repository) SaveRole(ctx context.Context, role *domain.Role) error {
	return r.DB(ctx).Save(role).Error
}

// DeleteRole 删除角色
func (r *Repository) DeleteRole(ctx context.Context, id types.Long) error {
	// 开启事务
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除角色
		if err := tx.Delete(&domain.Role{}, id).Error; err != nil {
			return err
		}

		// 删除角色权限关联
		if err := tx.Where("role_id = ?", id).Delete(&domain.RolePermission{}).Error; err != nil {
			return err
		}

		// 删除用户角色关联
		if err := tx.Where("role_id = ?", id).Delete(&domain.UserRole{}).Error; err != nil {
			return err
		}

		return nil
	})
}

// ============= 权限相关 =============

// GetPermissionByID 根据ID获取权限
func (r *Repository) GetPermissionByID(ctx context.Context, id types.Long) (*domain.Permission, error) {
	var permission domain.Permission
	err := r.DB(ctx).First(&permission, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &permission, nil
}

// ListPermissions 获取权限列表
func (r *Repository) ListPermissions(ctx context.Context, query *domain.PermissionQuery) ([]*domain.Permission, int64, error) {
	db := r.DB(ctx).Model(&domain.Permission{})

	// 应用查询条件
	if query.Name != "" {
		db = db.Where("name LIKE ?", "%"+query.Name+"%")
	}
	if query.Type != "" {
		db = db.Where("type = ?", query.Type)
	}
	if query.ParentID > 0 {
		db = db.Where("parent_id = ?", query.ParentID)
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
	db = db.Order("sort_order ASC, id ASC")

	// 查询数据
	var permissions []*domain.Permission
	err = db.Find(&permissions).Error
	if err != nil {
		return nil, 0, err
	}

	return permissions, total, nil
}

// SavePermission 保存权限
func (r *Repository) SavePermission(ctx context.Context, permission *domain.Permission) error {
	return r.DB(ctx).Save(permission).Error
}

// DeletePermission 删除权限
func (r *Repository) DeletePermission(ctx context.Context, id types.Long) error {
	// 开启事务
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 检查是否有子权限
		var count int64
		err := tx.Model(&domain.Permission{}).Where("parent_id = ?", id).Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			return errors.New("存在子权限，无法删除")
		}

		// 删除权限
		if err := tx.Delete(&domain.Permission{}, id).Error; err != nil {
			return err
		}

		// 删除角色权限关联
		if err := tx.Where("permission_id = ?", id).Delete(&domain.RolePermission{}).Error; err != nil {
			return err
		}

		return nil
	})
}

// GetAllPermissions 获取所有权限
func (r *Repository) GetAllPermissions(ctx context.Context) ([]*domain.Permission, error) {
	var permissions []*domain.Permission
	err := r.DB(ctx).Order("sort_order ASC, id ASC").Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// GetUserRoles 获取用户的角色列表
func (r *Repository) GetUserRoles(ctx context.Context, userID types.Long) ([]*domain.Role, error) {
	var roles []*domain.Role
	err := r.DB(ctx).Table("t_roles").
		Joins("JOIN t_user_roles ON t_roles.id = t_user_roles.role_id").
		Where("t_user_roles.user_id = ?", userID).
		Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// AssignRolesToUser 为用户分配角色
func (r *Repository) AssignRolesToUser(ctx context.Context, userID types.Long, roleIDs []types.Long) error {
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 先清除用户的所有角色
		if err := tx.Where("user_id = ?", userID).Delete(&domain.UserRole{}).Error; err != nil {
			return err
		}

		// 批量添加新角色
		userRoles := make([]domain.UserRole, 0, len(roleIDs))
		for _, roleID := range roleIDs {
			userRoles = append(userRoles, domain.UserRole{
				UserID: userID,
				RoleID: roleID,
			})
		}
		if len(userRoles) > 0 {
			if err := tx.Create(&userRoles).Error; err != nil {
				return err
			}
		}

		// 更新Casbin关系
		userKey := fmt.Sprintf("%s%d", domain.CasbinUserPrefix, userID)

		// 获取所有角色
		var roles []*domain.Role
		if err := tx.Where("id IN ?", roleIDs).Find(&roles).Error; err != nil {
			return err
		}

		// 先清除用户在Casbin中的所有角色
		if r.enforcer != nil {
			_, err := r.enforcer.DeleteRolesForUser(userKey)
			if err != nil {
				return err
			}

			// 添加新角色到Casbin
			for _, role := range roles {
				roleKey := fmt.Sprintf("%s%s", domain.CasbinRolePrefix, role.Code)
				_, err = r.enforcer.AddRoleForUser(userKey, roleKey)
				if err != nil {
					return err
				}
			}

			// 保存策略
			return r.enforcer.SavePolicy()
		}

		return nil
	})
}

// RemoveRoleFromUser 移除用户的角色
func (r *Repository) RemoveRoleFromUser(ctx context.Context, userID types.Long, roleID types.Long) error {
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除用户角色关联
		err := tx.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&domain.UserRole{}).Error
		if err != nil {
			return err
		}

		// 更新Casbin关系
		if r.enforcer != nil {
			// 获取角色信息
			var role domain.Role
			if err := tx.First(&role, roleID).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
				// 角色不存在，不需要处理Casbin
				return nil
			}

			userKey := fmt.Sprintf("%s%d", domain.CasbinUserPrefix, userID)
			roleKey := fmt.Sprintf("%s%s", domain.CasbinRolePrefix, role.Code)

			_, err = r.enforcer.DeleteRoleForUser(userKey, roleKey)
			if err != nil {
				return err
			}

			// 保存策略
			return r.enforcer.SavePolicy()
		}

		return nil
	})
}

// ============= 角色权限关联 =============

// GetRolePermissions 获取角色的权限列表
func (r *Repository) GetRolePermissions(ctx context.Context, roleID types.Long) ([]*domain.Permission, error) {
	var permissions []*domain.Permission
	err := r.DB(ctx).Table("t_permissions").
		Joins("JOIN t_role_permissions ON t_permissions.id = t_role_permissions.permission_id").
		Where("t_role_permissions.role_id = ?", roleID).
		Find(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

// AssignPermissionsToRole 为角色分配权限
func (r *Repository) AssignPermissionsToRole(ctx context.Context, roleID types.Long, permissionIDs []types.Long) error {
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 获取角色信息
		var role domain.Role
		if err := tx.First(&role, roleID).Error; err != nil {
			return err
		}

		// 删除原有的角色权限关联
		if err := tx.Where("role_id = ?", roleID).Delete(&domain.RolePermission{}).Error; err != nil {
			return err
		}

		// 如果权限ID列表为空，则仅清空权限即可
		if len(permissionIDs) == 0 {
			if r.enforcer != nil {
				// 清空该角色的所有Casbin策略
				roleKey := fmt.Sprintf("%s%s", domain.CasbinRolePrefix, role.Code)
				_, err := r.enforcer.RemoveFilteredPolicy(0, roleKey)
				if err != nil {
					return err
				}
				// 保存策略
				return r.enforcer.SavePolicy()
			}
			return nil
		}

		// 创建新的角色权限关联
		rolePermissions := make([]domain.RolePermission, 0, len(permissionIDs))
		now := time.Now()
		for _, permID := range permissionIDs {
			rolePermissions = append(rolePermissions, domain.RolePermission{
				RoleID:       roleID,
				PermissionID: permID,
				CreatedAt:    now,
			})
		}

		if err := tx.CreateInBatches(rolePermissions, 100).Error; err != nil {
			return err
		}

		// 更新Casbin策略
		if r.enforcer != nil {
			// 清空该角色的所有Casbin策略
			roleKey := fmt.Sprintf("%s%s", domain.CasbinRolePrefix, role.Code)
			_, err := r.enforcer.RemoveFilteredPolicy(0, roleKey)
			if err != nil {
				return err
			}

			// 获取权限信息
			var permissions []*domain.Permission
			if err := tx.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
				return err
			}

			// 添加新权限到Casbin
			for _, perm := range permissions {
				if perm.Type == domain.PermTypeApi && perm.Path != "" {
					// 设置ApiPath和ApiMethod用于Casbin集成
					perm.ApiPath = perm.Path
					perm.ApiMethod = perm.Method

					_, err = r.enforcer.AddPolicy(roleKey, perm.ApiPath, perm.ApiMethod)
					if err != nil {
						return err
					}
				}
			}

			// 保存策略
			return r.enforcer.SavePolicy()
		}

		return nil
	})
}

// RemovePermissionFromRole 移除角色的权限
func (r *Repository) RemovePermissionFromRole(ctx context.Context, roleID types.Long, permissionID types.Long) error {
	return r.DB(ctx).Transaction(func(tx *gorm.DB) error {
		// 删除角色权限关联
		err := tx.Where("role_id = ? AND permission_id = ?", roleID, permissionID).Delete(&domain.RolePermission{}).Error
		if err != nil {
			return err
		}

		// 更新Casbin策略
		if r.enforcer != nil {
			// 获取角色信息
			var role domain.Role
			if err := tx.First(&role, roleID).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
				// 角色不存在，不需要处理Casbin
				return nil
			}

			// 获取权限信息
			var permission domain.Permission
			if err := tx.First(&permission, permissionID).Error; err != nil {
				if !errors.Is(err, gorm.ErrRecordNotFound) {
					return err
				}
				// 权限不存在，不需要处理Casbin
				return nil
			}

			roleKey := fmt.Sprintf("%s%s", domain.CasbinRolePrefix, role.Code)

			if permission.Type == domain.PermTypeApi && permission.Path != "" {
				// 设置ApiPath和ApiMethod用于Casbin集成
				permission.ApiPath = permission.Path
				permission.ApiMethod = permission.Method

				_, err = r.enforcer.RemovePolicy(roleKey, permission.ApiPath, permission.ApiMethod)
				if err != nil {
					return err
				}

				// 保存策略
				return r.enforcer.SavePolicy()
			}
		}

		return nil
	})
}

// ============= Casbin相关 =============

// AddPolicy 添加策略
func (r *Repository) AddPolicy(ctx context.Context, role string, path string, method string) error {
	if r.enforcer != nil {
		_, err := r.enforcer.AddPolicy(role, path, method)
		if err != nil {
			return err
		}
		return r.enforcer.SavePolicy()
	}
	return nil
}

// RemovePolicy 移除策略
func (r *Repository) RemovePolicy(ctx context.Context, role string, path string, method string) error {
	if r.enforcer != nil {
		_, err := r.enforcer.RemovePolicy(role, path, method)
		if err != nil {
			return err
		}
		return r.enforcer.SavePolicy()
	}
	return nil
}

// AddRoleForUser 为用户添加角色
func (r *Repository) AddRoleForUser(ctx context.Context, user string, role string) error {
	if r.enforcer != nil {
		_, err := r.enforcer.AddRoleForUser(user, role)
		if err != nil {
			return err
		}
		return r.enforcer.SavePolicy()
	}
	return nil
}

// RemoveRoleForUser 移除用户的角色
func (r *Repository) RemoveRoleForUser(ctx context.Context, user string, role string) error {
	if r.enforcer != nil {
		_, err := r.enforcer.DeleteRoleForUser(user, role)
		if err != nil {
			return err
		}
		return r.enforcer.SavePolicy()
	}
	return nil
}

// HasPermission 检查权限
func (r *Repository) HasPermission(ctx context.Context, user string, path string, method string) (bool, error) {
	if r.enforcer != nil {
		return r.enforcer.Enforce(user, path, method)
	}
	return false, errors.New("enforcer未初始化")
}
