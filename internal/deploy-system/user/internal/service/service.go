package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/internal/deploy-system/user/internal/repository"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
)

type Service struct {
	service.Service
	Repo *repository.Repository `inject:"UserRepository"`
}

// Create
func (s *Service) Create(ctx context.Context, command *domain.CreateUserCommand) (id types.Long, err error) {
	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "user service create")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "user service create")
	}()
	/*
		转换为数据结构体
	*/
	user, err := command.ToUser()

	if err != nil {
		return
	}

	/*
		增加审计信息
	*/
	user.AuditCreated(ctx)
	/*
		创建用户
	*/
	err = s.Repo.Create(ctx, user)
	if err != nil {
		return
	}
	/*
	 * 已创建的数据库ID
	 */
	id = user.ID
	return
}

/*
 * 更新用户信息
 */
func (s *Service) ModifyUserByID(ctx context.Context, command *domain.ModifyUserCommand) (err error) {
	/*
		参数判断
	*/
	err = command.Validate()
	if err != nil {
		return
	}

	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "user service ModifyUserById")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		s.FinishTransaction(ctx, err, "user service ModifyUserById")
	}()
	/*
		查询用户信息对应的ID
	*/
	user, _ := s.Repo.GetByID(ctx, command.ID)
	if user.ID == 0 {
		err := common.RequestNotFoundError("用户信息不存在")
		return err
	}
	/*
	 * 转换为数据库实体
	 */
	user.Name = command.Name
	user.Email = command.Email
	user.Mobile = command.Mobile
	user.Role = command.Role
	user.Username = command.Username
	user.Password = command.Password
	/*
	 * 增加审计信息
	 */
	user.AuditModified(ctx)
	/*
	 * 更新字段
	 */
	err = s.Repo.Save(ctx, user)
	if err != nil {
		return err
	}
	return
}

/*
 * 更新用户角色信息
 */
//ModifyUserRoleByID
func (s *Service) ModifyUserRoleByID(ctx context.Context, command domain.ModifyUserRoleCommand) (err error) {
	/*
	 * 参数判断
	 */
	if !command.Role.ValidRole() {
		err = common.RequestParamError("", errors.New("角色取值不正确"))
		return
	}
	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "user service ModifyUserById")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "user service ModifyUserById")
	}()
	/*
		查询用户信息对应的ID
	*/
	user, _ := s.Repo.GetByID(ctx, command.ID)
	if user.ID == 0 {
		err := common.RequestNotFoundError("用户信息不存在")
		return err
	}

	user.Role = command.Role
	/*
		增加审计信息
	*/
	user.AuditModified(ctx)
	/*
		更新字段
	*/
	err = s.Repo.Save(ctx, user)
	if err != nil {
		return
	}
	return
}

/*
*
禁用或启用用户  1 ：启用 0 ：禁用
*/
func (s *Service) ModifyUserStatusByID(ctx context.Context, command domain.ModifyUserStatusCommand) (err error) {
	/*
		参数判断
	*/
	if !command.Status.ValidStatus() {
		err = common.RequestParamError("", errors.New("状态取值不正确"))
		return
	}

	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "user service ModifyUserStatusById")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "user service ModifyUserStatusById")
	}()
	/*
		查询用户信息对应的ID
	*/
	user, _ := s.Repo.GetByID(ctx, command.ID)
	if user.ID == 0 {
		err := common.RequestNotFoundError("用户信息不存在")
		return err
	}

	user.Enable = command.Status
	/*
	 * 增加审计信息
	 */
	user.AuditModified(ctx)

	/*
	 * 更新字段
	 */
	err = s.Repo.Save(ctx, user)
	if err != nil {
		return err
	}

	return
}

/*
*
按名称修改用户密码
*/
func (s *Service) ModifyUserPasswordByID(ctx context.Context, command domain.ChangePasswordCommand) (err error) {
	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "user service ModifyUserPasswordByID")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "user service ModifyUserPasswordByID")
	}()
	/*
		查询用户信息对应的ID
	*/
	user, _ := s.Repo.GetByID(ctx, command.ID)
	if user.ID == 0 {
		err := common.RequestNotFoundError("用户信息不存在")
		return err
	}

	user.Password = common.SetPassword(command.Password)
	/*
	 * 增加审计信息
	 */
	user.AuditModified(ctx)

	err = s.Repo.Save(ctx, user)
	if err != nil {
		return err
	}

	return
}

func (s *Service) GetByID(ctx context.Context, ID types.Long) (*domain.User, error) {
	return s.Repo.GetByID(ctx, ID)
}

//func (s *Service) GetPasswordByName(ctx context.Context, name string) (password string, err error) {
//	return s.Repo.GetPasswordBy(ctx, name)
//}
