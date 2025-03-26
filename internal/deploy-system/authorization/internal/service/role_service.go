package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/internal/deploy-system/authorization/internal/repository"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"errors"

	"github.com/sirupsen/logrus"
)

// RoleService 角色服务实现
type RoleService struct {
	service.Service
	repo   *repository.Repository `inject:"AuthorizationRepository"`
	logger *logrus.Logger         `inject:"Logger"`
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

// CreateRole 创建角色
func (s *RoleService) CreateRole(ctx context.Context, command *domain.CreateRoleCommand) (roleID types.Long, err error) {
	// 检查角色编码是否已存在
	existRole, err := s.repo.GetRoleByCode(ctx, command.Code)
	if err != nil {
		return 0, common.InternalError("检查角色编码失败", err)
	}

	if existRole != nil {
		return 0, common.RequestParamError("", errors.New("角色编码已存在"))
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "create role")
	if err != nil {
		return 0, err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create role")
	}()

	// 转换为角色实体
	role, err := command.ToRole()
	if err != nil {
		return 0, err
	}

	// 保存角色
	err = s.repo.SaveRole(ctx, role)
	if err != nil {
		return 0, common.InternalError("保存角色失败", err)
	}

	return role.ID, nil
}

// UpdateRole 更新角色
func (s *RoleService) UpdateRole(ctx context.Context, command *domain.UpdateRoleCommand) (err error) {
	// 检查角色是否存在
	role, err := s.repo.GetRoleByID(ctx, command.ID)
	if err != nil {
		return common.InternalError("查询角色失败", err)
	}

	if role == nil {
		return common.RequestParamError("", errors.New("角色不存在"))
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "update role")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "update role")
	}()

	// 验证更新参数
	err = command.Validate()
	if err != nil {
		return err
	}

	// 检查角色编码是否已存在
	if command.Code != role.Code {
		existRole, err := s.repo.GetRoleByCode(ctx, command.Code)
		if err != nil {
			return common.InternalError("检查角色编码失败", err)
		}

		if existRole != nil {
			return common.RequestParamError("", errors.New("角色编码已存在"))
		}
	}

	// 更新角色信息
	role.Name = command.Name
	role.Code = command.Code
	role.Description = command.Description
	role.Status = command.Status

	// 保存角色
	err = s.repo.SaveRole(ctx, role)
	if err != nil {
		return common.InternalError("更新角色失败", err)
	}

	return nil
}

// DeleteRole 删除角色
func (s *RoleService) DeleteRole(ctx context.Context, roleID types.Long) (err error) {
	// 检查角色是否存在
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return common.InternalError("查询角色失败", err)
	}

	if role == nil {
		return common.RequestParamError("", errors.New("角色不存在"))
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "delete role")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "delete role")
	}()

	// 删除角色
	err = s.repo.DeleteRole(ctx, roleID)
	if err != nil {
		return common.InternalError("删除角色失败", err)
	}

	return nil
}

// GetRoleByID 根据ID获取角色
func (s *RoleService) GetRoleByID(ctx context.Context, roleID types.Long) (*domain.RoleVO, error) {
	// 查询角色
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return nil, common.InternalError("查询角色失败", err)
	}

	if role == nil {
		return nil, nil
	}

	return role.ToVO(), nil
}

// ListRoles 获取角色列表
func (s *RoleService) ListRoles(ctx context.Context, query *domain.RoleQuery) ([]*domain.RoleVO, int64, error) {
	// 查询角色列表
	roles, total, err := s.repo.ListRoles(ctx, query)
	if err != nil {
		return nil, 0, common.InternalError("查询角色列表失败", err)
	}

	// 转换为VO
	roleVOs := make([]*domain.RoleVO, 0, len(roles))
	for _, role := range roles {
		roleVOs = append(roleVOs, role.ToVO())
	}

	return roleVOs, total, nil
}

// AssignPermissionsToRole 为角色分配权限
func (s *RoleService) AssignPermissionsToRole(ctx context.Context, roleID types.Long, permissionIDs []types.Long) (err error) {
	if len(permissionIDs) == 0 {
		return common.RequestParamError("", errors.New("权限ID不能为空"))
	}

	// 检查角色是否存在
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return common.InternalError("查询角色失败", err)
	}

	if role == nil {
		return common.RequestParamError("", errors.New("角色不存在"))
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "assign permissions to role")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "assign permissions to role")
	}()

	// 分配权限
	err = s.repo.AssignPermissionsToRole(ctx, roleID, permissionIDs)
	if err != nil {
		return common.InternalError("分配权限失败", err)
	}

	return nil
}

// GetRolePermissions 获取角色权限
func (s *RoleService) GetRolePermissions(ctx context.Context, roleID types.Long) ([]*domain.PermissionVO, error) {
	// 检查角色是否存在
	role, err := s.repo.GetRoleByID(ctx, roleID)
	if err != nil {
		return nil, common.InternalError("查询角色失败", err)
	}

	if role == nil {
		return nil, common.RequestParamError("", errors.New("角色不存在"))
	}

	// 获取角色权限
	permissions, err := s.repo.GetRolePermissions(ctx, roleID)
	if err != nil {
		return nil, common.InternalError("获取角色权限失败", err)
	}

	// 转换为VO
	permissionVOs := make([]*domain.PermissionVO, 0, len(permissions))
	for _, perm := range permissions {
		permissionVOs = append(permissionVOs, perm.ToVO())
	}

	return permissionVOs, nil
}
