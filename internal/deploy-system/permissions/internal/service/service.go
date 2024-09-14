package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/permissions/internal/domain"
	"devops-platform/internal/deploy-system/permissions/internal/repository"
	"devops-platform/pkg/types"
)

type Service struct {
	service.Service
	Repo *repository.Repository `inject:"RbacRepository"`
}

// 创建一个角色
func (s *Service) CreateRole(ctx context.Context, command *domain.CreatePermissionCommand) (id types.Long, err error) {
	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "perm service create")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "perm service create")
	}()
	role, err := command.ToPermission()
	if err != nil {
		return
	}
	/*
		增加审计信息
	*/
	role.AuditCreated(ctx)

	err = s.Repo.CreatePerm(ctx, role)
	if err != nil {
		return
	}
	id = role.ID
	return
}

// 根据ID修改角色
//func (s *Service) ModifyRoleByPID(ctx context.Context, command *domain.ModifyPermissionCommand) (err error) {
//	/*
//		开启事务
//	*/
//	ctx, err = s.BeginTransaction(ctx, "perm service ModifyRoleByID")
//	if err != nil {
//		return
//	}
//	defer func() {
//		err = s.FinishTransaction(ctx, err, "perm service ModifyRoleByID")
//	}()
//	role, _ := s.Repo.FindRoleByID(ctx, command.ID)
//	if role == nil {
//		err := common.RequestNotFoundError("角色信息不存在")
//		return err
//
//	}
//	if err = command.Validate(); err != nil {
//		return
//	}
//
//	role.Name = command.Name
//	role.PermissionID = command.PermissionID
//	role.Code = command.Code
//	role.Desc = command.Desc
//
//	role.AuditModified(ctx)
//	if err = s.Repo.UpdateRole(ctx, role); err != nil {
//		return
//	}
//	return
//
//}

// 根据ID删除角色
//func (s *Service) DeleteRoleByID(ctx context.Context, ID types.Long) (err error) {
//	/*
//		开启事务
//	*/
//	ctx, err = s.BeginTransaction(ctx, "perm service DeleteRoleByID")
//	if err != nil {
//		return
//	}
//	defer func() {
//		err = s.FinishTransaction(ctx, err, "perm service DeleteRoleByID")
//	}()
//	role, _ := s.Repo.FindRoleByID(ctx, ID)
//	if role == nil {
//		err := common.RequestNotFoundError("角色信息不存在")
//		return err
//	}
//	/*
//		删除部门
//	*/
//	if err = s.Repo.DeleteRole(ctx, ID); err != nil {
//		return
//	}
//
//	return
//}

// get role by id
//func (s *Service) FindRoleByID(ctx context.Context, ID types.Long) (role *domain.Role, err error) {
//	role, err = s.Repo.FindRoleByID(ctx, ID)
//	if err != nil {
//		return
//	}
//	role.VO()
//	return
//
//}
//func (s *Service) FindRoleByName(ctx context.Context, name string, pagination types.Pagination) (roles []domain.Role, total int64, err error) {
//	roles, total, err = s.Repo.FindRoleByName(ctx, name, pagination)
//	if err != nil {
//		return
//	}
//	for i := range roles {
//		roles[i].VO()
//	}
//	return
//}
