package application

import (
	"context"

	"devops-platform/internal/deploy-system/application/internal/domain"
	"devops-platform/pkg/types"
)

//go:generate mockgen -source=export.go -destination=mock/mock_application.go -package=mock

// Bean常量
const (
	BeanAppService    = domain.BeanService       // 应用管理服务Bean名称
	BeanDeployService = domain.BeanDeployService // 部署服务Bean名称
	BeanAppQuery      = domain.BeanAppQuery      // 应用查询服务Bean名称
)

// AppService 应用管理服务接口
type AppService interface {
	// CreateApplication 创建应用
	CreateApplication(ctx context.Context, command *domain.CreateAppCommand) (types.Long, error)

	// UpdateApplication 更新应用
	UpdateApplication(ctx context.Context, command *domain.UpdateAppCommand) error

	// DeleteApplication 删除应用
	DeleteApplication(ctx context.Context, id types.Long) error

	// CreateAppGroup 创建应用分组
	CreateAppGroup(ctx context.Context, name, description string) (types.Long, error)

	// UpdateAppGroup 更新应用分组
	UpdateAppGroup(ctx context.Context, id types.Long, name, description string) error

	// DeleteAppGroup 删除应用分组
	DeleteAppGroup(ctx context.Context, id types.Long) error

	// AddAppToGroup 添加应用到分组
	AddAppToGroup(ctx context.Context, appID, groupID types.Long) error

	// RemoveAppFromGroup 从分组中移除应用
	RemoveAppFromGroup(ctx context.Context, appID, groupID types.Long) error

	// CreateAppEnv 创建应用环境
	CreateAppEnv(ctx context.Context, command *domain.CreateEnvCommand) (types.Long, error)

	// UpdateAppEnv 更新应用环境
	UpdateAppEnv(ctx context.Context, id types.Long, name, namespace, description string) error

	// DeleteAppEnv 删除应用环境
	DeleteAppEnv(ctx context.Context, id types.Long) error

	// CreateImageRegistry 创建镜像仓库
	CreateImageRegistry(ctx context.Context, name, url, username, password, email string) (types.Long, error)

	// UpdateImageRegistry 更新镜像仓库
	UpdateImageRegistry(ctx context.Context, id types.Long, name, url, username, password, email string) error

	// DeleteImageRegistry 删除镜像仓库
	DeleteImageRegistry(ctx context.Context, id types.Long) error
}

// AppQuery 应用查询接口
type AppQueryServer interface {
	// GetApplicationByID 根据ID获取应用
	GetApplicationByID(ctx context.Context, id types.Long) (*domain.Application, error)

	// GetApplicationByName 根据名称获取应用
	GetApplicationByName(ctx context.Context, name string) (*domain.Application, error)

	// ListApplications 查询应用列表
	ListApplications(ctx context.Context, query *domain.AppQuery) ([]*domain.AppVO, int64, error)

	// GetAppGroups 获取应用所属的分组
	GetAppGroups(ctx context.Context, appID types.Long) ([]*domain.AppGroup, error)

	// GetGroupApps 获取分组中的应用
	GetGroupApps(ctx context.Context, groupID types.Long) ([]*domain.Application, error)

	// ListAppEnvs 查询应用环境列表
	ListAppEnvs(ctx context.Context) ([]*domain.AppEnv, error)

	// GetAppEnvByID 根据ID获取应用环境
	GetAppEnvByID(ctx context.Context, id types.Long) (*domain.AppEnv, error)

	// ListImageRegistries 查询镜像仓库列表
	ListImageRegistries(ctx context.Context) ([]*domain.ImageRegistry, error)

	// GetAppHPA 获取应用HPA配置
	GetAppHPA(ctx context.Context, appID types.Long) (*domain.AppHPA, error)
}

// DeployService 部署服务接口
type DeployService interface {
	// CreateReleasePlan 创建发布计划
	CreateReleasePlan(ctx context.Context, command *domain.CreateReleaseCommand) (types.Long, error)

	// ExecuteReleasePlan 执行发布计划
	ExecuteReleasePlan(ctx context.Context, planID types.Long) (types.Long, error)

	// GetDeployment 获取部署记录
	GetDeployment(ctx context.Context, id types.Long) (*domain.Deployment, error)

	// ListDeployments 查询部署历史列表
	ListDeployments(ctx context.Context, appID, envID types.Long) ([]*domain.Deployment, error)

	// RollbackDeployment 回滚部署
	RollbackDeployment(ctx context.Context, id types.Long) error

	// CreateHPA 创建/更新应用HPA配置
	CreateHPA(ctx context.Context, appID types.Long, minReplicas, maxReplicas, targetCPU, targetMemory int) (types.Long, error)

	// DeleteAppHPA 删除应用HPA配置
	DeleteAppHPA(ctx context.Context, appID types.Long) error
}

// 领域对象类型别名
type Application = domain.Application
type AppGroup = domain.AppGroup
type AppEnv = domain.AppEnv
type AppHPA = domain.AppHPA
type Deployment = domain.Deployment
type DeploymentStep = domain.DeploymentStep
type ReleasePlan = domain.ReleasePlan
type ImageRegistry = domain.ImageRegistry
type CreateAppCommand = domain.CreateAppCommand
type UpdateAppCommand = domain.UpdateAppCommand
type CreateEnvCommand = domain.CreateEnvCommand
type CreateReleaseCommand = domain.CreateReleaseCommand
type AppQuery = domain.AppQuery
type AppVO = domain.AppVO
