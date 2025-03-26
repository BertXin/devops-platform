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

// PermissionService 权限服务实现
type PermissionService struct {
	service.Service
	repo   *repository.Repository `inject:"AuthorizationRepository"`
	logger *logrus.Logger         `inject:"Logger"`
}

func NewPermissionService() *PermissionService {
	return &PermissionService{}
}

// CreatePermission 创建权限
func (s *PermissionService) CreatePermission(ctx context.Context, command *domain.CreatePermissionCommand) (permissionID types.Long, err error) {
	// 如果指定了父权限，检查父权限是否存在
	if command.ParentID > 0 {
		parent, err := s.repo.GetPermissionByID(ctx, command.ParentID)
		if err != nil {
			return 0, common.InternalError("查询父权限失败", err)
		}

		if parent == nil {
			return 0, common.RequestParamError("", errors.New("父权限不存在"))
		}
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "create permission")
	if err != nil {
		return 0, err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create permission")
	}()

	// 转换为权限实体
	permission, err := command.ToPermission()
	if err != nil {
		return 0, err
	}

	// 保存权限
	err = s.repo.SavePermission(ctx, permission)
	if err != nil {
		return 0, common.InternalError("保存权限失败", err)
	}

	return permission.ID, nil
}

// UpdatePermission 更新权限
func (s *PermissionService) UpdatePermission(ctx context.Context, command *domain.UpdatePermissionCommand) (err error) {
	// 检查权限是否存在
	permission, err := s.repo.GetPermissionByID(ctx, command.ID)
	if err != nil {
		return common.InternalError("查询权限失败", err)
	}

	if permission == nil {
		return common.RequestParamError("", errors.New("权限不存在"))
	}

	// 如果指定了父权限，检查父权限是否存在
	if command.ParentID > 0 && command.ParentID != permission.ParentID {
		parent, err := s.repo.GetPermissionByID(ctx, command.ParentID)
		if err != nil {
			return common.InternalError("查询父权限失败", err)
		}

		if parent == nil {
			return common.RequestParamError("", errors.New("父权限不存在"))
		}

		// 检查是否形成循环依赖
		if command.ID == command.ParentID {
			return common.RequestParamError("", errors.New("不能将自己设为父权限"))
		}
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "update permission")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "update permission")
	}()

	// 验证更新参数
	err = command.Validate()
	if err != nil {
		return common.RequestParamError("", err)
	}

	// 更新权限信息
	permission.ParentID = command.ParentID
	permission.Name = command.Name
	permission.Type = command.Type
	permission.Path = command.Path
	permission.Method = command.Method
	permission.Icon = command.Icon
	permission.Component = command.Component
	permission.Permission = command.Permission
	permission.Status = command.Status
	permission.Hidden = command.Hidden
	permission.SortOrder = command.SortOrder

	// 保存权限
	err = s.repo.SavePermission(ctx, permission)
	if err != nil {
		return common.InternalError("更新权限失败", err)
	}

	return nil
}

// DeletePermission 删除权限
func (s *PermissionService) DeletePermission(ctx context.Context, permissionID types.Long) (err error) {
	// 检查权限是否存在
	permission, err := s.repo.GetPermissionByID(ctx, permissionID)
	if err != nil {
		return common.InternalError("查询权限失败", err)
	}

	if permission == nil {
		return common.RequestParamError("", errors.New("权限不存在"))
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "delete permission")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "delete permission")
	}()

	// 删除权限
	err = s.repo.DeletePermission(ctx, permissionID)
	if err != nil {
		return common.InternalError("删除权限失败", err)
	}

	return nil
}

// GetPermissionByID 根据ID获取权限
func (s *PermissionService) GetPermissionByID(ctx context.Context, permissionID types.Long) (*domain.PermissionVO, error) {
	// 查询权限
	permission, err := s.repo.GetPermissionByID(ctx, permissionID)
	if err != nil {
		return nil, common.InternalError("查询权限失败", err)
	}

	if permission == nil {
		return nil, nil
	}

	return permission.ToVO(), nil
}

// ListPermissions 获取权限列表
func (s *PermissionService) ListPermissions(ctx context.Context, query *domain.PermissionQuery) ([]*domain.PermissionVO, int64, error) {
	// 查询权限列表
	permissions, total, err := s.repo.ListPermissions(ctx, query)
	if err != nil {
		return nil, 0, common.InternalError("查询权限列表失败", err)
	}

	// 转换为VO
	permissionVOs := make([]*domain.PermissionVO, 0, len(permissions))
	for _, perm := range permissions {
		permissionVOs = append(permissionVOs, perm.ToVO())
	}

	return permissionVOs, total, nil
}

// GetPermissionTree 获取权限树结构
func (s *PermissionService) GetPermissionTree(ctx context.Context) ([]*domain.PermissionVO, error) {
	// 获取所有权限
	permissions, err := s.repo.GetAllPermissions(ctx)
	if err != nil {
		return nil, common.InternalError("获取权限列表失败", err)
	}

	// 转换为VO
	permMap := make(map[types.Long]*domain.PermissionVO)
	for _, perm := range permissions {
		vo := perm.ToVO()
		vo.Children = make([]*domain.PermissionVO, 0)
		permMap[vo.ID] = vo
	}

	// 构建树结构
	var roots []*domain.PermissionVO
	for _, perm := range permMap {
		if perm.ParentID == 0 {
			// 根节点
			roots = append(roots, perm)
		} else {
			// 子节点，添加到父节点的子列表中
			if parent, ok := permMap[perm.ParentID]; ok {
				parent.Children = append(parent.Children, perm)
			} else {
				// 父节点不存在，作为根节点处理
				roots = append(roots, perm)
			}
		}
	}

	return roots, nil
}
