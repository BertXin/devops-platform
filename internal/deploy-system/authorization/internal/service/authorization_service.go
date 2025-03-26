package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/internal/deploy-system/authorization/internal/repository"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

// AuthorizationService 权限服务实现
type AuthorizationService struct {
	service.Service
	Repo   *repository.Repository `inject:"AuthorizationRepository"`
	Logger *logrus.Logger         `inject:"Logger"`
}

func NewAuthorizationService() *AuthorizationService {
	return &AuthorizationService{}
}

// HasPermission 检查用户是否拥有指定权限
func (s *AuthorizationService) HasPermission(ctx context.Context, userID types.Long, permission string) (bool, error) {
	userKey := fmt.Sprintf("%s%d", domain.CasbinUserPrefix, userID)

	// 如果是API权限，检查对应的path和method
	// 此处简化处理，实际应该通过permission标识符找到对应的API信息
	// 这里假设permission的格式为 "api:path:method"

	// 示例实现，实际项目中需要根据权限标识获取对应的资源和操作
	has, err := s.Repo.HasPermission(ctx, userKey, permission, "*")
	if err != nil {
		return false, common.InternalError("检查权限失败", err)
	}

	return has, nil
}

// GetUserRoles 获取用户的角色列表
func (s *AuthorizationService) GetUserRoles(ctx context.Context, userID types.Long) ([]*domain.RoleVO, error) {
	roles, err := s.Repo.GetUserRoles(ctx, userID)
	if err != nil {
		s.Logger.WithError(err).Error("获取用户角色失败")
		return nil, common.InternalError("获取用户角色失败", err)
	}

	roleVOs := make([]*domain.RoleVO, 0, len(roles))
	for _, role := range roles {
		roleVOs = append(roleVOs, role.ToVO())
	}

	return roleVOs, nil
}

// GetUserPermissions 获取用户的权限列表
func (s *AuthorizationService) GetUserPermissions(ctx context.Context, userID types.Long) ([]*domain.PermissionVO, error) {
	// 获取用户的所有角色
	roles, err := s.Repo.GetUserRoles(ctx, userID)
	if err != nil {
		s.Logger.WithError(err).Error("获取用户角色失败")
		return nil, common.InternalError("获取用户角色失败", err)
	}

	// 获取每个角色的权限
	permissionMap := make(map[types.Long]*domain.Permission)
	for _, role := range roles {
		permissions, err := s.Repo.GetRolePermissions(ctx, role.ID)
		if err != nil {
			s.Logger.WithError(err).Error("获取角色权限失败")
			return nil, common.InternalError("获取角色权限失败", err)
		}

		// 去重
		for _, perm := range permissions {
			permissionMap[perm.ID] = perm
		}
	}

	// 转换为VO
	permissionVOs := make([]*domain.PermissionVO, 0, len(permissionMap))
	for _, perm := range permissionMap {
		permissionVOs = append(permissionVOs, perm.ToVO())
	}

	return permissionVOs, nil
}

// AssignRolesToUser 为用户分配角色
func (s *AuthorizationService) AssignRolesToUser(ctx context.Context, userID types.Long, roleIDs []types.Long) (err error) {
	if len(roleIDs) == 0 {
		return common.RequestParamError("", errors.New("角色ID不能为空"))
	}
	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "assign roles to user")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "assign roles to user")
	}()

	err = s.Repo.AssignRolesToUser(ctx, userID, roleIDs)
	if err != nil {
		s.Logger.WithError(err).Error("分配角色失败")
		return common.InternalError("分配角色失败", err)
	}

	return
}

// RemoveRoleFromUser 移除用户的角色
func (s *AuthorizationService) RemoveRoleFromUser(ctx context.Context, userID types.Long, roleID types.Long) (err error) {
	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "remove role from user")
	if err != nil {
		return err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "remove role from user")
	}()

	err = s.Repo.RemoveRoleFromUser(ctx, userID, roleID)
	if err != nil {
		s.Logger.WithError(err).Error("移除角色失败")
		return common.InternalError("移除角色失败", err)
	}

	return
}

// GetUserMenus 获取用户菜单
func (s *AuthorizationService) GetUserMenus(ctx context.Context, userID types.Long) ([]*domain.MenuVO, error) {
	// 获取用户的权限列表
	permissionVOs, err := s.GetUserPermissions(ctx, userID)
	if err != nil {
		return nil, err
	}

	// 过滤出菜单类型的权限
	menuMap := make(map[types.Long]*domain.MenuVO)
	for _, perm := range permissionVOs {
		if perm.Type == domain.PermTypeMenu {
			menuMap[perm.ID] = &domain.MenuVO{
				ID:        perm.ID,
				ParentID:  perm.ParentID,
				Name:      perm.Name,
				Path:      perm.Path,
				Component: perm.Component,
				Icon:      perm.Icon,
				SortOrder: perm.SortOrder,
				Hidden:    perm.Hidden,
				Children:  []*domain.MenuVO{},
			}
		}
	}

	// 构建菜单树
	rootMenus := make([]*domain.MenuVO, 0)
	for _, menu := range menuMap {
		// 如果是顶级菜单，直接添加到结果中
		if menu.ParentID == 0 {
			rootMenus = append(rootMenus, menu)
		} else {
			// 如果有父菜单，添加到父菜单的子菜单中
			if parent, ok := menuMap[menu.ParentID]; ok {
				parent.Children = append(parent.Children, menu)
			}
		}
	}

	return rootMenus, nil
}
