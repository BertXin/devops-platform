package service

import (
	"context"

	"devops-platform/internal/deploy-system/application/internal/domain"
	"devops-platform/internal/deploy-system/application/internal/repository"
	"devops-platform/pkg/types"
)

// AppQuery 应用查询服务实现
type AppQuery struct {
	repo repository.Repository
}

// NewAppQuery 创建应用查询服务实例
func NewAppQuery() *AppQuery {
	return &AppQuery{}
}

// GetApplicationByID 根据ID获取应用
func (q *AppQuery) GetApplicationByID(ctx context.Context, id types.Long) (*domain.Application, error) {
	return q.repo.GetApplicationByID(ctx, id)
}

// GetApplicationByName 根据名称获取应用
func (q *AppQuery) GetApplicationByName(ctx context.Context, name string) (*domain.Application, error) {
	return q.repo.GetApplicationByName(ctx, name)
}

// ListApplications 查询应用列表
func (q *AppQuery) ListApplications(ctx context.Context, query *domain.AppQuery) ([]*domain.AppVO, int64, error) {
	return q.repo.ListApplications(ctx, query)
}

// GetAppGroups 获取应用所属的分组
func (q *AppQuery) GetAppGroups(ctx context.Context, appID types.Long) ([]*domain.AppGroup, error) {
	return q.repo.GetAppGroups(ctx, appID)
}

// GetGroupApps 获取分组中的应用
func (q *AppQuery) GetGroupApps(ctx context.Context, groupID types.Long) ([]*domain.Application, error) {
	return q.repo.GetGroupApps(ctx, groupID)
}

// GetAppEnvByID 根据ID获取应用环境
func (q *AppQuery) GetAppEnvByID(ctx context.Context, id types.Long) (*domain.AppEnv, error) {
	return q.repo.GetAppEnvByID(ctx, id)
}

// ListAppEnvs 查询应用环境列表
func (q *AppQuery) ListAppEnvs(ctx context.Context) ([]*domain.AppEnv, error) {
	return q.repo.ListAppEnvs(ctx)
}

// GetReleasePlanByID 根据ID获取发布计划
func (q *AppQuery) GetReleasePlanByID(ctx context.Context, id types.Long) (*domain.ReleasePlan, error) {
	return q.repo.GetReleasePlanByID(ctx, id)
}

// ListReleasePlans 查询发布计划列表
func (q *AppQuery) ListReleasePlans(ctx context.Context, appID types.Long) ([]*domain.ReleasePlan, error) {
	return q.repo.ListReleasePlans(ctx, appID)
}

// GetDeploymentByID 根据ID获取部署记录
func (q *AppQuery) GetDeploymentByID(ctx context.Context, id types.Long) (*domain.Deployment, error) {
	return q.repo.GetDeploymentByID(ctx, id)
}

// ListDeployments 查询部署历史列表
func (q *AppQuery) ListDeployments(ctx context.Context, appID, envID types.Long) ([]*domain.Deployment, error) {
	return q.repo.ListDeployments(ctx, appID, envID)
}

// GetDeploymentSteps 获取部署步骤列表
func (q *AppQuery) GetDeploymentSteps(ctx context.Context, deployID types.Long) ([]*domain.DeploymentStep, error) {
	return q.repo.GetDeploymentSteps(ctx, deployID)
}

// GetImageRegistryByID 根据ID获取镜像仓库
func (q *AppQuery) GetImageRegistryByID(ctx context.Context, id types.Long) (*domain.ImageRegistry, error) {
	return q.repo.GetImageRegistryByID(ctx, id)
}

// ListImageRegistries 查询镜像仓库列表
func (q *AppQuery) ListImageRegistries(ctx context.Context) ([]*domain.ImageRegistry, error) {
	return q.repo.ListImageRegistries(ctx)
}

// GetAppHPA 获取应用HPA配置
func (q *AppQuery) GetAppHPA(ctx context.Context, appID types.Long) (*domain.AppHPA, error) {
	return q.repo.GetAppHPAByAppID(ctx, appID)
}
