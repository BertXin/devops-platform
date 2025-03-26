package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/organization/internal/domain"
	"devops-platform/internal/deploy-system/organization/internal/repository"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"errors"

	"github.com/sirupsen/logrus"
)

// DepartmentService 部门服务实现
type DepartmentService struct {
	service.Service
	Repo   *repository.Repository `inject:"DepartmentRepository"`
	Logger *logrus.Logger         `inject:"Logger"`
}

// NewDepartmentService 创建部门服务实例
func NewDepartmentService() *DepartmentService {
	return &DepartmentService{}
}

// CreateDepartment 创建部门
func (s *DepartmentService) CreateDepartment(ctx context.Context, command *domain.CreateDepartmentCommand) (id types.Long, err error) {
	// 检查部门编码是否已存在
	existDept, err := s.Repo.GetDepartmentByCode(ctx, command.Code)
	if err != nil {
		s.Logger.WithError(err).Error("检查部门编码失败")
		return 0, common.InternalError("检查部门编码失败", err)
	}

	if existDept != nil {
		return 0, common.RequestParamError("", errors.New("部门编码已存在"))
	}

	// 如果指定了父部门，检查父部门是否存在
	if command.ParentID > 0 {
		parentDept, err := s.Repo.GetDepartmentByID(ctx, command.ParentID)
		if err != nil {
			s.Logger.WithError(err).Error("查询父部门失败")
			return 0, common.InternalError("查询父部门失败", err)
		}

		if parentDept == nil {
			return 0, common.RequestParamError("", errors.New("父部门不存在"))
		}
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "department service create")
	if err != nil {
		return 0, err
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "department service create")
	}()

	// 转换为部门实体
	department, err := command.ToDepartment()
	if err != nil {
		return 0, err
	}

	// 保存部门
	err = s.Repo.SaveDepartment(ctx, department)
	if err != nil {
		s.Logger.WithError(err).Error("保存部门失败")
		return 0, common.InternalError("保存部门失败", err)
	}

	return department.ID, nil
}

// UpdateDepartment 更新部门
func (s *DepartmentService) UpdateDepartment(ctx context.Context, command *domain.UpdateDepartmentCommand) (err error) {
	// 检查部门是否存在
	department, err := s.Repo.GetDepartmentByID(ctx, command.ID)
	if err != nil {
		s.Logger.WithError(err).Error("查询部门失败")
		return common.InternalError("查询部门失败", err)
	}

	if department == nil {
		return common.RequestParamError("", errors.New("部门不存在"))
	}

	// 检查部门编码是否已被其他部门使用
	if department.Code != command.Code {
		existDept, err := s.Repo.GetDepartmentByCode(ctx, command.Code)
		if err != nil {
			s.Logger.WithError(err).Error("检查部门编码失败")
			return common.InternalError("检查部门编码失败", err)
		}

		if existDept != nil && existDept.ID != command.ID {
			return common.RequestParamError("", errors.New("部门编码已被使用"))
		}
	}

	// 如果指定了父部门，检查父部门是否存在
	if command.ParentID > 0 && command.ParentID != department.ParentID {
		parentDept, err := s.Repo.GetDepartmentByID(ctx, command.ParentID)
		if err != nil {
			s.Logger.WithError(err).Error("查询父部门失败")
			return common.InternalError("查询父部门失败", err)
		}

		if parentDept == nil {
			return common.RequestParamError("", errors.New("父部门不存在"))
		}

		// 检查是否形成循环依赖
		if command.ID == command.ParentID {
			return common.RequestParamError("", errors.New("不能将自己设为父部门"))
		}

		// TODO: 检查更深层次的循环依赖（如果A的父部门是B，B的父部门是C，要设置C的父部门为A）
	}

	// 验证更新参数
	err = command.Validate()
	if err != nil {
		return err
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "department service update")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "department service update")
	}()

	// 更新部门信息
	department.ParentID = command.ParentID
	department.Name = command.Name
	department.Code = command.Code
	department.Description = command.Description
	department.Status = command.Status
	department.Sort = command.Sort

	// 保存部门
	err = s.Repo.SaveDepartment(ctx, department)
	if err != nil {
		s.Logger.WithError(err).Error("更新部门失败")
		return common.InternalError("更新部门失败", err)
	}

	return nil
}

// DeleteDepartment 删除部门
func (s *DepartmentService) DeleteDepartment(ctx context.Context, id types.Long) (err error) {
	// 检查部门是否存在
	department, err := s.Repo.GetDepartmentByID(ctx, id)
	if err != nil {
		s.Logger.WithError(err).Error("查询部门失败")
		return common.InternalError("查询部门失败", err)
	}

	if department == nil {
		return common.RequestParamError("", errors.New("部门不存在"))
	}

	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "department service delete")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "department service delete")
	}()
	// 删除部门
	err = s.Repo.DeleteDepartment(ctx, id)
	if err != nil {
		s.Logger.WithError(err).Error("删除部门失败")
		return common.InternalError("删除部门失败", err)
	}

	return nil
}

// GetDepartmentByID 根据ID获取部门
func (s *DepartmentService) GetDepartmentByID(ctx context.Context, id types.Long) (*domain.DepartmentVO, error) {
	// 查询部门
	department, err := s.Repo.GetDepartmentByID(ctx, id)
	if err != nil {
		s.Logger.WithError(err).Error("查询部门失败")
		return nil, common.InternalError("查询部门失败", err)
	}

	if department == nil {
		return nil, nil
	}

	return department.ToVO(), nil
}

// ListDepartments 获取部门列表
func (s *DepartmentService) ListDepartments(ctx context.Context, query *domain.DepartmentQuery) ([]*domain.DepartmentVO, int64, error) {
	// 查询部门列表
	departments, total, err := s.Repo.ListDepartments(ctx, query)
	if err != nil {
		s.Logger.WithError(err).Error("查询部门列表失败")
		return nil, 0, common.InternalError("查询部门列表失败", err)
	}

	// 转换为VO
	departmentVOs := make([]*domain.DepartmentVO, 0, len(departments))
	for _, dept := range departments {
		departmentVOs = append(departmentVOs, dept.ToVO())
	}

	return departmentVOs, total, nil
}

// GetDepartmentTree 获取部门树结构
func (s *DepartmentService) GetDepartmentTree(ctx context.Context) ([]*domain.DepartmentVO, error) {
	// 获取所有部门
	departments, err := s.Repo.GetAllDepartments(ctx)
	if err != nil {
		s.Logger.WithError(err).Error("获取部门列表失败")
		return nil, common.InternalError("获取部门列表失败", err)
	}

	// 转换为VO
	deptsMap := make(map[types.Long]*domain.DepartmentVO)
	for _, dept := range departments {
		deptVO := dept.ToVO()
		deptVO.Children = make([]*domain.DepartmentVO, 0)
		deptsMap[dept.ID] = deptVO
	}

	// 构建树形结构
	roots := make([]*domain.DepartmentVO, 0)
	for _, deptVO := range deptsMap {
		if deptVO.ParentID == 0 {
			// 顶级部门
			roots = append(roots, deptVO)
		} else {
			// 子部门
			if parent, ok := deptsMap[deptVO.ParentID]; ok {
				parent.Children = append(parent.Children, deptVO)
			}
		}
	}

	return roots, nil
}

// GetUserDepartments 获取用户所属部门
func (s *DepartmentService) GetUserDepartments(ctx context.Context, userID types.Long) ([]*domain.DepartmentVO, error) {
	departments, err := s.Repo.GetUserDepartments(ctx, userID)
	if err != nil {
		s.Logger.WithError(err).Error("获取用户部门失败")
		return nil, common.InternalError("获取用户部门失败", err)
	}

	// 转换为VO
	departmentVOs := make([]*domain.DepartmentVO, 0, len(departments))
	for _, dept := range departments {
		departmentVOs = append(departmentVOs, dept.ToVO())
	}

	return departmentVOs, nil
}

// AssignDepartmentToUser 为用户分配部门
func (s *DepartmentService) AssignDepartmentToUser(ctx context.Context, userID types.Long, departmentID types.Long, isLeader bool) (err error) {
	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "assign department to user")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "assign department to user")
	}()
	err = s.Repo.AssignDepartmentToUser(ctx, userID, departmentID, isLeader)
	if err != nil {
		s.Logger.WithError(err).Error("分配部门失败")
		return common.InternalError("分配部门失败", err)
	}

	return nil
}

// RemoveDepartmentFromUser 移除用户部门
func (s *DepartmentService) RemoveDepartmentFromUser(ctx context.Context, userID types.Long, departmentID types.Long) (err error) {
	// 创建事务上下文
	ctx, err = s.BeginTransaction(ctx, "remove department from user")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "remove department from user")
	}()
	err = s.Repo.RemoveDepartmentFromUser(ctx, userID, departmentID)
	if err != nil {
		s.Logger.WithError(err).Error("移除部门失败")
		return common.InternalError("移除部门失败", err)
	}

	return nil
}

// ListDepartmentUsers 获取部门用户列表
func (s *DepartmentService) ListDepartmentUsers(ctx context.Context, departmentID types.Long, query *domain.UserQuery) ([]*domain.UserVO, int64, error) {
	users, total, err := s.Repo.ListDepartmentUsers(ctx, departmentID, query)
	if err != nil {
		s.Logger.WithError(err).Error("获取部门用户列表失败")
		return nil, 0, common.InternalError("获取部门用户列表失败", err)
	}

	return users, total, nil
}
