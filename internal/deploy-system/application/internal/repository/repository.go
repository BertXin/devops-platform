package repository

import (
	"context"
	"devops-platform/internal/common/repository"

	"devops-platform/internal/deploy-system/application/internal/domain"
	"devops-platform/pkg/types"
)

// Repository 应用管理仓储接口
type Repository interface {
	// 应用相关
	CreateApplication(ctx context.Context, app *domain.Application) (types.Long, error)
	UpdateApplication(ctx context.Context, app *domain.Application) error
	GetApplicationByID(ctx context.Context, id types.Long) (*domain.Application, error)
	GetApplicationByName(ctx context.Context, name string) (*domain.Application, error)
	ListApplications(ctx context.Context, query *domain.AppQuery) ([]*domain.AppVO, int64, error)
	DeleteApplication(ctx context.Context, id types.Long) error

	// 应用分组相关
	CreateAppGroup(ctx context.Context, group *domain.AppGroup) (types.Long, error)
	UpdateAppGroup(ctx context.Context, group *domain.AppGroup) error
	GetAppGroupByID(ctx context.Context, id types.Long) (*domain.AppGroup, error)
	ListAppGroups(ctx context.Context) ([]*domain.AppGroup, error)
	DeleteAppGroup(ctx context.Context, id types.Long) error

	// 应用-分组关联
	AddAppToGroup(ctx context.Context, appID, groupID types.Long) error
	RemoveAppFromGroup(ctx context.Context, appID, groupID types.Long) error
	GetAppGroups(ctx context.Context, appID types.Long) ([]*domain.AppGroup, error)
	GetGroupApps(ctx context.Context, groupID types.Long) ([]*domain.Application, error)

	// 环境相关
	CreateAppEnv(ctx context.Context, env *domain.AppEnv) (types.Long, error)
	UpdateAppEnv(ctx context.Context, env *domain.AppEnv) error
	GetAppEnvByID(ctx context.Context, id types.Long) (*domain.AppEnv, error)
	ListAppEnvs(ctx context.Context) ([]*domain.AppEnv, error)
	DeleteAppEnv(ctx context.Context, id types.Long) error

	// 发布计划相关
	CreateReleasePlan(ctx context.Context, plan *domain.ReleasePlan) (types.Long, error)
	UpdateReleasePlan(ctx context.Context, plan *domain.ReleasePlan) error
	GetReleasePlanByID(ctx context.Context, id types.Long) (*domain.ReleasePlan, error)
	ListReleasePlans(ctx context.Context, appID types.Long) ([]*domain.ReleasePlan, error)
	DeleteReleasePlan(ctx context.Context, id types.Long) error

	// 部署历史相关
	CreateDeployment(ctx context.Context, deployment *domain.Deployment) (types.Long, error)
	UpdateDeployment(ctx context.Context, deployment *domain.Deployment) error
	GetDeploymentByID(ctx context.Context, id types.Long) (*domain.Deployment, error)
	ListDeployments(ctx context.Context, appID, envID types.Long) ([]*domain.Deployment, error)

	// 部署步骤相关
	CreateDeploymentStep(ctx context.Context, step *domain.DeploymentStep) (types.Long, error)
	UpdateDeploymentStep(ctx context.Context, step *domain.DeploymentStep) error
	GetDeploymentSteps(ctx context.Context, deployID types.Long) ([]*domain.DeploymentStep, error)

	// 镜像仓库相关
	CreateImageRegistry(ctx context.Context, registry *domain.ImageRegistry) (types.Long, error)
	UpdateImageRegistry(ctx context.Context, registry *domain.ImageRegistry) error
	GetImageRegistryByID(ctx context.Context, id types.Long) (*domain.ImageRegistry, error)
	ListImageRegistries(ctx context.Context) ([]*domain.ImageRegistry, error)
	DeleteImageRegistry(ctx context.Context, id types.Long) error

	// 应用HPA相关
	CreateAppHPA(ctx context.Context, hpa *domain.AppHPA) (types.Long, error)
	UpdateAppHPA(ctx context.Context, hpa *domain.AppHPA) error
	GetAppHPAByAppID(ctx context.Context, appID types.Long) (*domain.AppHPA, error)
	DeleteAppHPA(ctx context.Context, id types.Long) error
}

type AppRepository struct {
	repository.Repository
}

func NewAppRepository() *AppRepository {
	return &AppRepository{}
}

// CreateApplication 创建应用
func (r *AppRepository) CreateApplication(ctx context.Context, app *domain.Application) (types.Long, error) {
	if err := r.DB(ctx).Create(app).Error; err != nil {
		return 0, err
	}
	return app.ID, nil
}

// UpdateApplication 更新应用
func (r *AppRepository) UpdateApplication(ctx context.Context, app *domain.Application) error {
	return r.DB(ctx).Updates(app).Error
}

// GetApplicationByID 根据ID获取应用
func (r *AppRepository) GetApplicationByID(ctx context.Context, id types.Long) (*domain.Application, error) {
	var app domain.Application
	if err := r.DB(ctx).First(&app, id).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

// GetApplicationByName 根据名称获取应用
func (r *AppRepository) GetApplicationByName(ctx context.Context, name string) (*domain.Application, error) {
	var app domain.Application
	if err := r.DB(ctx).Where("name = ?", name).First(&app).Error; err != nil {
		return nil, err
	}
	return &app, nil
}

// ListApplications 查询应用列表
func (r *AppRepository) ListApplications(ctx context.Context, query *domain.AppQuery) ([]*domain.AppVO, int64, error) {
	db := r.DB(ctx).Table("app").Select("app.*")

	// 应用条件查询
	if query.Name != "" {
		db = db.Where("app.name LIKE ?", "%"+query.Name+"%")
	}
	if query.Status != "" {
		db = db.Where("app.status = ?", query.Status)
	}

	// 分组查询
	if query.GroupID > 0 {
		db = db.Joins("JOIN relation_app_group_app ON app.id = relation_app_group_app.app_id").
			Where("relation_app_group_app.group_id = ?", query.GroupID)
	}

	// 计算总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (query.Page - 1) * query.Size
	var appVOs []*domain.AppVO

	err := db.Offset(offset).Limit(query.Size).
		Order("app.id DESC").
		Scan(&appVOs).Error

	if err != nil {
		return nil, 0, err
	}

	// 查询每个应用的其他信息
	for _, appVO := range appVOs {
		// 查询应用环境数量
		var envCount int64
		r.DB(ctx).Table("app_env").
			Where("app_id = ?", appVO.ID).
			Count(&envCount)
		appVO.EnvCount = int(envCount)

		// 查询应用所属分组
		var groupIDs []types.Long
		r.DB(ctx).Table("relation_app_group_app").
			Select("group_id").
			Where("app_id = ?", appVO.ID).
			Pluck("group_id", &groupIDs)
		appVO.GroupIDs = groupIDs
	}

	return appVOs, total, nil
}

// DeleteApplication 删除应用
func (r *AppRepository) DeleteApplication(ctx context.Context, id types.Long) error {
	// 软删除，仅修改状态
	return r.DB(ctx).Model(&domain.Application{}).
		Where("id = ?", id).
		Update("status", domain.AppStatusDeleted).Error
}

// CreateAppGroup 创建应用分组
func (r *AppRepository) CreateAppGroup(ctx context.Context, group *domain.AppGroup) (types.Long, error) {
	if err := r.DB(ctx).Create(group).Error; err != nil {
		return 0, err
	}
	return group.ID, nil
}

// UpdateAppGroup 更新应用分组
func (r *AppRepository) UpdateAppGroup(ctx context.Context, group *domain.AppGroup) error {
	return r.DB(ctx).Updates(group).Error
}

// GetAppGroupByID 根据ID获取应用分组
func (r *AppRepository) GetAppGroupByID(ctx context.Context, id types.Long) (*domain.AppGroup, error) {
	var group domain.AppGroup
	if err := r.DB(ctx).First(&group, id).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

// ListAppGroups 查询应用分组列表
func (r *AppRepository) ListAppGroups(ctx context.Context) ([]*domain.AppGroup, error) {
	var groups []*domain.AppGroup
	if err := r.DB(ctx).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

// DeleteAppGroup 删除应用分组
func (r *AppRepository) DeleteAppGroup(ctx context.Context, id types.Long) error {
	return r.DB(ctx).Delete(&domain.AppGroup{}, id).Error
}

// AddAppToGroup 添加应用到分组
func (r *AppRepository) AddAppToGroup(ctx context.Context, appID, groupID types.Long) error {
	relation := &domain.AppGroupRelation{
		AppID:   appID,
		GroupID: groupID,
	}
	return r.DB(ctx).Create(relation).Error
}

// RemoveAppFromGroup 从分组中移除应用
func (r *AppRepository) RemoveAppFromGroup(ctx context.Context, appID, groupID types.Long) error {
	return r.DB(ctx).Where("app_id = ? AND group_id = ?", appID, groupID).
		Delete(&domain.AppGroupRelation{}).Error
}

// GetAppGroups 获取应用所属的分组
func (r *AppRepository) GetAppGroups(ctx context.Context, appID types.Long) ([]*domain.AppGroup, error) {
	var groups []*domain.AppGroup
	err := r.DB(ctx).
		Joins("JOIN relation_app_group_app ON app_group.id = relation_app_group_app.group_id").
		Where("relation_app_group_app.app_id = ?", appID).
		Find(&groups).Error
	return groups, err
}

// GetGroupApps 获取分组中的应用
func (r *AppRepository) GetGroupApps(ctx context.Context, groupID types.Long) ([]*domain.Application, error) {
	var apps []*domain.Application
	err := r.DB(ctx).
		Joins("JOIN relation_app_group_app ON app.id = relation_app_group_app.app_id").
		Where("relation_app_group_app.group_id = ?", groupID).
		Find(&apps).Error
	return apps, err
}

// CreateAppEnv 创建应用环境
func (r *AppRepository) CreateAppEnv(ctx context.Context, env *domain.AppEnv) (types.Long, error) {
	if err := r.DB(ctx).Create(env).Error; err != nil {
		return 0, err
	}
	return env.ID, nil
}

// UpdateAppEnv 更新应用环境
func (r *AppRepository) UpdateAppEnv(ctx context.Context, env *domain.AppEnv) error {
	return r.DB(ctx).Updates(env).Error
}

// GetAppEnvByID 根据ID获取应用环境
func (r *AppRepository) GetAppEnvByID(ctx context.Context, id types.Long) (*domain.AppEnv, error) {
	var env domain.AppEnv
	if err := r.DB(ctx).First(&env, id).Error; err != nil {
		return nil, err
	}
	return &env, nil
}

// ListAppEnvs 查询应用环境列表
func (r *AppRepository) ListAppEnvs(ctx context.Context) ([]*domain.AppEnv, error) {
	var envs []*domain.AppEnv
	if err := r.DB(ctx).Find(&envs).Error; err != nil {
		return nil, err
	}
	return envs, nil
}

// DeleteAppEnv 删除应用环境
func (r *AppRepository) DeleteAppEnv(ctx context.Context, id types.Long) error {
	return r.DB(ctx).Delete(&domain.AppEnv{}, id).Error
}

// CreateReleasePlan 创建发布计划
func (r *AppRepository) CreateReleasePlan(ctx context.Context, plan *domain.ReleasePlan) (types.Long, error) {
	if err := r.DB(ctx).Create(plan).Error; err != nil {
		return 0, err
	}
	return plan.ID, nil
}

// UpdateReleasePlan 更新发布计划
func (r *AppRepository) UpdateReleasePlan(ctx context.Context, plan *domain.ReleasePlan) error {
	return r.DB(ctx).Updates(plan).Error
}

// GetReleasePlanByID 根据ID获取发布计划
func (r *AppRepository) GetReleasePlanByID(ctx context.Context, id types.Long) (*domain.ReleasePlan, error) {
	var plan domain.ReleasePlan
	if err := r.DB(ctx).First(&plan, id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

// ListReleasePlans 查询发布计划列表
func (r *AppRepository) ListReleasePlans(ctx context.Context, appID types.Long) ([]*domain.ReleasePlan, error) {
	var plans []*domain.ReleasePlan
	query := r.DB(ctx)
	if appID > 0 {
		query = query.Where("app_id = ?", appID)
	}
	if err := query.Find(&plans).Error; err != nil {
		return nil, err
	}
	return plans, nil
}

// DeleteReleasePlan 删除发布计划
func (r *AppRepository) DeleteReleasePlan(ctx context.Context, id types.Long) error {
	return r.DB(ctx).Delete(&domain.ReleasePlan{}, id).Error
}

// CreateDeployment 创建部署记录
func (r *AppRepository) CreateDeployment(ctx context.Context, deployment *domain.Deployment) (types.Long, error) {
	if err := r.DB(ctx).Create(deployment).Error; err != nil {
		return 0, err
	}
	return deployment.ID, nil
}

// UpdateDeployment 更新部署记录
func (r *AppRepository) UpdateDeployment(ctx context.Context, deployment *domain.Deployment) error {
	return r.DB(ctx).Updates(deployment).Error
}

// GetDeploymentByID 根据ID获取部署记录
func (r *AppRepository) GetDeploymentByID(ctx context.Context, id types.Long) (*domain.Deployment, error) {
	var deployment domain.Deployment
	if err := r.DB(ctx).First(&deployment, id).Error; err != nil {
		return nil, err
	}

	// 获取部署步骤
	steps, err := r.GetDeploymentSteps(ctx, deployment.ID)
	if err == nil {
		deployment.Steps = steps
	}

	return &deployment, nil
}

// ListDeployments 查询部署历史列表
func (r *AppRepository) ListDeployments(ctx context.Context, appID, envID types.Long) ([]*domain.Deployment, error) {
	var deployments []*domain.Deployment
	query := r.DB(ctx)
	if appID > 0 {
		query = query.Where("app_id = ?", appID)
	}
	if envID > 0 {
		query = query.Where("env_id = ?", envID)
	}
	if err := query.Order("id DESC").Find(&deployments).Error; err != nil {
		return nil, err
	}
	return deployments, nil
}

// CreateDeploymentStep 创建部署步骤
func (r *AppRepository) CreateDeploymentStep(ctx context.Context, step *domain.DeploymentStep) (types.Long, error) {
	if err := r.DB(ctx).Create(step).Error; err != nil {
		return 0, err
	}
	return step.ID, nil
}

// UpdateDeploymentStep 更新部署步骤
func (r *AppRepository) UpdateDeploymentStep(ctx context.Context, step *domain.DeploymentStep) error {
	return r.DB(ctx).Updates(step).Error
}

// GetDeploymentSteps 获取部署步骤列表
func (r *AppRepository) GetDeploymentSteps(ctx context.Context, deployID types.Long) ([]*domain.DeploymentStep, error) {
	var steps []*domain.DeploymentStep
	if err := r.DB(ctx).Where("deploy_id = ?", deployID).
		Order("id ASC").Find(&steps).Error; err != nil {
		return nil, err
	}
	return steps, nil
}

// CreateImageRegistry 创建镜像仓库
func (r *AppRepository) CreateImageRegistry(ctx context.Context, registry *domain.ImageRegistry) (types.Long, error) {
	if err := r.DB(ctx).Create(registry).Error; err != nil {
		return 0, err
	}
	return registry.ID, nil
}

// UpdateImageRegistry 更新镜像仓库
func (r *AppRepository) UpdateImageRegistry(ctx context.Context, registry *domain.ImageRegistry) error {
	return r.DB(ctx).Updates(registry).Error
}

// GetImageRegistryByID 根据ID获取镜像仓库
func (r *AppRepository) GetImageRegistryByID(ctx context.Context, id types.Long) (*domain.ImageRegistry, error) {
	var registry domain.ImageRegistry
	if err := r.DB(ctx).First(&registry, id).Error; err != nil {
		return nil, err
	}
	return &registry, nil
}

// ListImageRegistries 查询镜像仓库列表
func (r *AppRepository) ListImageRegistries(ctx context.Context) ([]*domain.ImageRegistry, error) {
	var registries []*domain.ImageRegistry
	if err := r.DB(ctx).Find(&registries).Error; err != nil {
		return nil, err
	}
	return registries, nil
}

// DeleteImageRegistry 删除镜像仓库
func (r *AppRepository) DeleteImageRegistry(ctx context.Context, id types.Long) error {
	return r.DB(ctx).Delete(&domain.ImageRegistry{}, id).Error
}

// CreateAppHPA 创建应用HPA
func (r *AppRepository) CreateAppHPA(ctx context.Context, hpa *domain.AppHPA) (types.Long, error) {
	if err := r.DB(ctx).Create(hpa).Error; err != nil {
		return 0, err
	}
	return hpa.ID, nil
}

// UpdateAppHPA 更新应用HPA
func (r *AppRepository) UpdateAppHPA(ctx context.Context, hpa *domain.AppHPA) error {
	return r.DB(ctx).Updates(hpa).Error
}

// GetAppHPAByAppID 根据应用ID获取HPA
func (r *AppRepository) GetAppHPAByAppID(ctx context.Context, appID types.Long) (*domain.AppHPA, error) {
	var hpa domain.AppHPA
	if err := r.DB(ctx).Where("app_id = ?", appID).First(&hpa).Error; err != nil {
		return nil, err
	}
	return &hpa, nil
}

// DeleteAppHPA 删除应用HPA
func (r *AppRepository) DeleteAppHPA(ctx context.Context, id types.Long) error {
	return r.DB(ctx).Delete(&domain.AppHPA{}, id).Error
}
