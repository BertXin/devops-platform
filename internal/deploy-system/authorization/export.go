package authorization

import (
	"context"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/pkg/types"
)

//go:generate mockgen -source=export.go -destination=mock/mock_authorization.go -package=mock

// Bean常量
const (
	BeanAuthorizationService = domain.BeanAuthorizationService
	BeanRoleService          = domain.BeanRoleService
	BeanPermissionService    = domain.BeanPermissionService
)

// 领域对象类型别名
type (
	RoleVO                  = domain.RoleVO
	PermissionVO            = domain.PermissionVO
	MenuVO                  = domain.MenuVO
	CreateRoleCommand       = domain.CreateRoleCommand
	UpdateRoleCommand       = domain.UpdateRoleCommand
	RoleQuery               = domain.RoleQuery
	CreatePermissionCommand = domain.CreatePermissionCommand
	UpdatePermissionCommand = domain.UpdatePermissionCommand
	PermissionQuery         = domain.PermissionQuery
)

// 权限服务接口
type AuthorizationService interface {
	// HasPermission 检查用户是否拥有指定权限
	HasPermission(ctx context.Context, userID types.Long, permission string) (bool, error)

	// GetUserRoles 获取用户的角色列表
	GetUserRoles(ctx context.Context, userID types.Long) ([]*domain.RoleVO, error)

	// GetUserPermissions 获取用户的权限列表
	GetUserPermissions(ctx context.Context, userID types.Long) ([]*domain.PermissionVO, error)

	// AssignRolesToUser 为用户分配角色
	AssignRolesToUser(ctx context.Context, userID types.Long, roleIDs []types.Long) error

	// RemoveRoleFromUser 移除用户的角色
	RemoveRoleFromUser(ctx context.Context, userID types.Long, roleID types.Long) error

	// GetUserMenus 获取用户菜单
	GetUserMenus(ctx context.Context, userID types.Long) ([]*domain.MenuVO, error)
}

// 角色服务接口
type RoleService interface {
	// CreateRole 创建角色
	CreateRole(ctx context.Context, command *domain.CreateRoleCommand) (types.Long, error)

	// UpdateRole 更新角色
	UpdateRole(ctx context.Context, command *domain.UpdateRoleCommand) error

	// DeleteRole 删除角色
	DeleteRole(ctx context.Context, roleID types.Long) error

	// GetRoleByID 根据ID获取角色
	GetRoleByID(ctx context.Context, roleID types.Long) (*domain.RoleVO, error)

	// ListRoles 获取角色列表
	ListRoles(ctx context.Context, query *domain.RoleQuery) ([]*domain.RoleVO, int64, error)

	// AssignPermissionsToRole 为角色分配权限
	AssignPermissionsToRole(ctx context.Context, roleID types.Long, permissionIDs []types.Long) error

	// GetRolePermissions 获取角色权限
	GetRolePermissions(ctx context.Context, roleID types.Long) ([]*domain.PermissionVO, error)
}

// 权限查询接口
type PermissionService interface {
	// CreatePermission 创建权限
	CreatePermission(ctx context.Context, command *domain.CreatePermissionCommand) (types.Long, error)

	// UpdatePermission 更新权限
	UpdatePermission(ctx context.Context, command *domain.UpdatePermissionCommand) error

	// DeletePermission 删除权限
	DeletePermission(ctx context.Context, permissionID types.Long) error

	// GetPermissionByID 根据ID获取权限
	GetPermissionByID(ctx context.Context, permissionID types.Long) (*domain.PermissionVO, error)

	// ListPermissions 获取权限列表
	ListPermissions(ctx context.Context, query *domain.PermissionQuery) ([]*domain.PermissionVO, int64, error)

	// GetPermissionTree 获取权限树结构
	GetPermissionTree(ctx context.Context) ([]*domain.PermissionVO, error)
}
