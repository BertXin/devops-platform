package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/department/internal/domain"
	"devops-platform/internal/deploy-system/department/internal/repository"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
)

type Service struct {
	service.Service
	Repo *repository.Repository `inject:"DeptRepository"`
}

// Create 创建一个新的部门。
//
// 该方法接收一个上下文对象ctx，用于控制请求的生命周期和传递取消信号。
// 它还接收一个指向domain.CreateDeptCommand的指针，该命令包含了创建部门所需的所有信息。
// 方法返回一个types.Long类型的id，表示新创建部门的唯一标识，以及一个错误对象err。
// 如果创建部门操作成功，id将包含新部门的ID，err将为nil；如果操作失败，err将包含错误信息，id的值将不明确。
func (s *Service) Create(ctx context.Context, command *domain.CreateDeptCommand) (id types.Long, err error) {

	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "dept service create")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "dept service create")
	}()
	dept, err := command.ToDept()
	if err != nil {
		return
	}
	/*
		增加审计信息
	*/
	dept.AuditCreated(ctx)

	if err = s.Repo.Create(ctx, dept); err != nil {
		return
	}
	/*
	 * 已创建的数据库ID
	 */
	id = dept.ID
	return
}

// ModifyDeptNameByID 根据部门ID修改部门名称。
//
// 此函数接收一个上下文Context，用于控制请求的生命周期和传递取消信号。
// 它还接收一个指向domain.ModifyDeptCommand的指针，该命令包含了修改部门名称所需的信息。
// 函数返回一个错误类型，用于指示修改过程中是否发生了错误。
//
// 该函数是Service方法的一部分，设计用于修改企业组织结构中的部门名称。
// 它通过部门ID来唯一标识部门，并应用修改。
func (s *Service) ModifyDeptNameByID(ctx context.Context, command *domain.ModifyDeptCommand) (err error) {

	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "dept service ModifyDeptNameByID")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "dept service ModifyDeptNameByID")
	}()

	dept, _ := s.Repo.GetByID(ctx, command.ID)
	if dept.ID == 0 {
		err := common.RequestNotFoundError("部门信息不存在")
		return err
	}
	dept.Name = command.Name
	/*
		增加审计信息
	*/
	dept.AuditModified(ctx)
	if err = s.Repo.Update(ctx, dept); err != nil {
		return
	}
	return
}

// DeleteDeptByID 根据部门ID删除部门。
//
// 本函数接收一个上下文Context和一个部门ID Long作为参数。
// Context用于控制请求的生命周期，可以传递取消信号或超时设置等。
// ID是待删除的部门的唯一标识。
//
// 函数返回一个错误类型。
// 如果删除操作成功，错误将为nil；如果操作失败，错误将包含失败的原因。
//
// 本函数是Service类型的方法，用于实现对部门的管理操作。
func (s *Service) DeleteDeptByID(ctx context.Context, ID types.Long) (err error) {

	/*
		开启事务
	*/
	ctx, err = s.BeginTransaction(ctx, "dept service delete")
	if err != nil {
		return
	}
	/*
		结束事务
	*/
	defer func() {
		err = s.FinishTransaction(ctx, err, "dept service delete")
	}()
	dept, _ := s.Repo.GetByID(ctx, ID)

	if dept.ID == 0 {
		err := common.RequestNotFoundError("部门信息不存在")
		return err
	}
	/*
		删除部门
	*/
	if err = s.Repo.Delete(ctx, ID); err != nil {
		return
	}
	return
}

// FindDeptByID 根据部门ID查询特定部门的信息。
// 这个方法是在Service类型上的一个方法，专门用于查询部门信息。
// 它接收一个上下文Context，用于控制请求的生命周期，和一个长整型ID作为部门的唯一标识。
// 返回值是一个指向domain.Dept类型的指针，以及可能的错误。
// 通过这个方法，可以高效地根据ID查找并返回对应的部门对象。
func (s *Service) FindDeptByID(ctx context.Context, ID types.Long) (dept *domain.Dept, err error) {

	dept, err = s.Repo.GetByID(ctx, ID)
	if err != nil {
		return
	}
	dept.VO()
	return
}
func (s *Service) FindDeptByName(ctx context.Context, name string, parentid types.Long, pagination types.Pagination) (depts []domain.Dept, total int64, err error) {
	depts, total, err = s.Repo.FindByName(ctx, name, parentid, pagination)
	if err != nil {
		return
	}
	for i := range depts {
		depts[i].VO()
	}
	return
}
func (s *Service) FindDeptByParentID(ctx context.Context, parentID types.Long, pagination types.Pagination) (depts []domain.Dept, err error) {
	depts, err = s.Repo.FindByParentID(ctx, parentID, pagination)
	if err != nil {
		return
	}
	for i := range depts {
		depts[i].VO()
	}
	return
}
func (s *Service) ListDept(ctx context.Context, pagination types.Pagination) (depts []domain.Dept, err error) {
	depts, err = s.Repo.List(ctx, pagination)
	if err != nil {
		return
	}
	for i := range depts {
		depts[i].VO()
	}
	return
}
