package service

import (
	"context"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/internal/deploy-system/auth/internal/repository"
	"devops-platform/internal/deploy-system/organization"
	"devops-platform/pkg/types"
	"fmt"

	"github.com/sirupsen/logrus"
)

// UserQuery 用户查询服务实现
type UserQuery struct {
	Repo              repository.Repository
	Logger            *logrus.Logger
	DepartmentService organization.DepartmentService
}

func NewUserQuery() *UserQuery {
	return &UserQuery{}
}

// GetByUsername 根据用户名查询用户
func (q *UserQuery) GetByUsername(ctx context.Context, username string) (*domain.UserInfo, error) {
	user, err := q.Repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	deptName := ""
	if user.DeptID > 0 && q.DepartmentService != nil {
		dept, err := q.DepartmentService.GetDepartmentByID(ctx, user.DeptID)
		if err == nil && dept != nil {
			deptName = dept.Name
		} else if err != nil && q.Logger != nil {
			q.Logger.WithError(err).WithField("deptID", user.DeptID).Warn("获取部门信息失败")
		}
	}

	return user.ToUserInfo(deptName), nil
}

// GetByID 根据ID查询用户
func (q *UserQuery) GetByID(ctx context.Context, ID types.Long) (*domain.UserInfo, error) {
	user, err := q.Repo.GetByID(ctx, ID)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	if user == nil {
		return nil, nil
	}

	deptName := ""
	if user.DeptID > 0 && q.DepartmentService != nil {
		dept, err := q.DepartmentService.GetDepartmentByID(ctx, user.DeptID)
		if err == nil && dept != nil {
			deptName = dept.Name
		} else if err != nil && q.Logger != nil {
			q.Logger.WithError(err).WithField("deptID", user.DeptID).Warn("获取部门信息失败")
		}
	}

	return user.ToUserInfo(deptName), nil
}
