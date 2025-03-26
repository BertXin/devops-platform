package organization

import (
	"context"
	"devops-platform/internal/deploy-system/organization/internal/domain"
	"devops-platform/pkg/types"
)

//go:generate mockgen -source=export.go -destination=mock/mock_organization.go -package=mock

// Bean常量
const (
	BeanDepartmentService = domain.BeanDepartmentService
)

// 领域对象类型别名
type (
	DepartmentVO            = domain.DepartmentVO
	CreateDepartmentCommand = domain.CreateDepartmentCommand
	UpdateDepartmentCommand = domain.UpdateDepartmentCommand
	DepartmentQuery         = domain.DepartmentQuery
	UserVO                  = domain.UserVO
	UserQuery               = domain.UserQuery
)

// DepartmentService 部门服务接口
type DepartmentService interface {
	// CreateDepartment 创建部门
	CreateDepartment(ctx context.Context, command *domain.CreateDepartmentCommand) (types.Long, error)

	// UpdateDepartment 更新部门
	UpdateDepartment(ctx context.Context, command *domain.UpdateDepartmentCommand) error

	// DeleteDepartment 删除部门
	DeleteDepartment(ctx context.Context, id types.Long) error

	// GetDepartmentByID 根据ID获取部门
	GetDepartmentByID(ctx context.Context, id types.Long) (*domain.DepartmentVO, error)

	// ListDepartments 获取部门列表
	ListDepartments(ctx context.Context, query *domain.DepartmentQuery) ([]*domain.DepartmentVO, int64, error)

	// GetDepartmentTree 获取部门树结构
	GetDepartmentTree(ctx context.Context) ([]*domain.DepartmentVO, error)

	// GetUserDepartments 获取用户所属部门
	GetUserDepartments(ctx context.Context, userID types.Long) ([]*domain.DepartmentVO, error)

	// AssignDepartmentToUser 为用户分配部门
	AssignDepartmentToUser(ctx context.Context, userID types.Long, departmentID types.Long, isLeader bool) error

	// RemoveDepartmentFromUser 移除用户部门
	RemoveDepartmentFromUser(ctx context.Context, userID types.Long, departmentID types.Long) error

	// ListDepartmentUsers 获取部门用户列表
	ListDepartmentUsers(ctx context.Context, departmentID types.Long, query *domain.UserQuery) ([]*domain.UserVO, int64, error)
}
