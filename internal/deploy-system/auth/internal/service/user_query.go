package service

import (
	"context"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/internal/deploy-system/auth/internal/repository"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"

	"github.com/sirupsen/logrus"
)

// UserQueryService 用户查询服务
type UserQueryService struct {
	service.Service
	repo   repository.Repository `inject:"UserRepository"`
	logger *logrus.Logger        `inject:"Logger"`
}

func NewUserQueryService() *UserQueryService {
	return &UserQueryService{}
}

// GetByUsername 根据用户名查询用户信息
func (q *UserQueryService) GetByUsername(ctx context.Context, username string) (*domain.UserInfo, error) {
	user, err := q.repo.GetByUsername(ctx, username)
	if err != nil {
		q.logger.WithError(err).Error("查询用户失败")
		return nil, common.InternalError("系统内部错误", err)
	}

	if user == nil {
		return nil, common.NotFoundError("用户不存在", nil)
	}

	// 这里应该通过部门服务查询部门名称，简化处理返回空
	return user.ToUserInfo(""), nil
}

// GetByID 根据ID查询用户信息
func (q *UserQueryService) GetByID(ctx context.Context, ID types.Long) (*domain.UserInfo, error) {
	user, err := q.repo.GetByID(ctx, ID)
	if err != nil {
		q.logger.WithError(err).Error("查询用户失败")
		return nil, common.InternalError("系统内部错误", err)
	}

	if user == nil {
		return nil, common.NotFoundError("用户不存在", nil)
	}

	// 这里应该通过部门服务查询部门名称，简化处理返回空
	return user.ToUserInfo(""), nil
}
