package domain

import (
	"devops-platform/internal/pkg/module"
	"time"

	"devops-platform/pkg/types"
)

// Application 应用实体
type Application struct {
	module.Module
	ID          types.Long   `json:"id" gorm:"primaryKey"`
	Name        string       `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Description string       `json:"description" gorm:"size:500"`
	Creator     types.Long   `json:"creator" gorm:"not null"`
	Status      string       `json:"status" gorm:"size:20;not null;default:'active'"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	Groups      []*AppGroup  `json:"groups,omitempty" gorm:"-"`
	Envs        []AppEnv     `json:"envs,omitempty" gorm:"-"`
	Deployments []Deployment `json:"deployments,omitempty" gorm:"-"`
}

// AppGroup 应用分组实体
type AppGroup struct {
	module.Module
	Name        string             `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Description string             `json:"description" gorm:"size:500"`
	Apps        []AppGroupRelation `json:"apps,omitempty" gorm:"-"`
}

// AppGroupRelation 应用-分组关联表
type AppGroupRelation struct {
	ID      types.Long `json:"id" gorm:"primaryKey"`
	GroupID types.Long `json:"group_id" gorm:"not null;index"`
	AppID   types.Long `json:"app_id" gorm:"not null;index"`
}

// AppEnv 应用环境实体
type AppEnv struct {
	module.Module
	Name        string     `json:"name" gorm:"size:100;not null"`
	ClusterID   types.Long `json:"cluster_id" gorm:"not null;index"`
	Namespace   string     `json:"namespace" gorm:"size:100;not null"`
	Description string     `json:"description" gorm:"size:500"`
}

// AppHPA 应用HPA配置
type AppHPA struct {
	module.Module
	AppID        types.Long `json:"app_id" gorm:"not null;index"`
	MinReplicas  int        `json:"min_replicas" gorm:"not null;default:1"`
	MaxReplicas  int        `json:"max_replicas" gorm:"not null;default:10"`
	TargetCPU    int        `json:"target_cpu" gorm:"not null;default:80"`
	TargetMemory int        `json:"target_memory" gorm:"default:0"`
}

// ImageRegistry 镜像仓库配置
type ImageRegistry struct {
	module.Module
	Name     string `json:"name" gorm:"size:100;not null;uniqueIndex"`
	URL      string `json:"url" gorm:"size:200;not null"`
	Username string `json:"username" gorm:"size:100"`
	Password string `json:"password" gorm:"size:100"`
	Email    string `json:"email" gorm:"size:200"`
}

// AppImageRegistry 应用-镜像仓库关联
type AppImageRegistry struct {
	module.Module
	AppID      types.Long `json:"app_id" gorm:"not null;index"`
	RegistryID types.Long `json:"registry_id" gorm:"not null;index"`
}

// ReleasePlan 发布计划
type ReleasePlan struct {
	module.Module
	AppID    types.Long `json:"app_id" gorm:"not null;index"`
	EnvID    types.Long `json:"env_id" gorm:"not null;index"`
	Version  string     `json:"version" gorm:"size:50;not null"`
	Strategy string     `json:"strategy" gorm:"size:50;not null;default:'rolling'"`
	Status   string     `json:"status" gorm:"size:20;not null;default:'pending'"`
}

// Deployment 部署记录
type Deployment struct {
	module.Module
	AppID     types.Long        `json:"app_id" gorm:"not null;index"`
	EnvID     types.Long        `json:"env_id" gorm:"not null;index"`
	Version   string            `json:"version" gorm:"size:50;not null"`
	Status    string            `json:"status" gorm:"size:20;not null;default:'pending'"`
	StartTime time.Time         `json:"start_time" gorm:"not null"`
	EndTime   *time.Time        `json:"end_time"`
	Steps     []*DeploymentStep `json:"steps,omitempty" gorm:"-"`
}

// DeploymentStep 部署步骤记录
type DeploymentStep struct {
	module.Module
	DeployID  types.Long `json:"deploy_id" gorm:"not null;index"`
	Name      string     `json:"name" gorm:"size:100;not null"`
	Status    string     `json:"status" gorm:"size:20;not null;default:'pending'"`
	Message   string     `json:"message" gorm:"size:1000"`
	StartTime time.Time  `json:"start_time" gorm:"not null"`
	EndTime   *time.Time `json:"end_time"`
}

// CreateAppCommand 创建应用命令
type CreateAppCommand struct {
	Name        string     `json:"name" binding:"required,max=100"`
	Description string     `json:"description" binding:"max=500"`
	Creator     types.Long `json:"creator"`
}

// UpdateAppCommand 更新应用命令
type UpdateAppCommand struct {
	ID          types.Long `json:"id" binding:"required"`
	Name        string     `json:"name" binding:"max=100"`
	Description string     `json:"description" binding:"max=500"`
	Status      string     `json:"status" binding:"max=20"`
}

// CreateEnvCommand 创建环境命令
type CreateEnvCommand struct {
	Name        string     `json:"name" binding:"required,max=100"`
	ClusterID   types.Long `json:"cluster_id" binding:"required"`
	Namespace   string     `json:"namespace" binding:"required,max=100"`
	Description string     `json:"description" binding:"max=500"`
}

// CreateReleaseCommand 创建发布命令
type CreateReleaseCommand struct {
	AppID    types.Long `json:"app_id" binding:"required"`
	EnvID    types.Long `json:"env_id" binding:"required"`
	Version  string     `json:"version" binding:"required,max=50"`
	Strategy string     `json:"strategy" binding:"required,max=50"`
}

// AppQuery 应用查询参数
type AppQuery struct {
	Name    string     `json:"name" form:"name"`
	Status  string     `json:"status" form:"status"`
	GroupID types.Long `json:"group_id" form:"group_id"`
	Page    int        `json:"page" form:"page,default=1"`
	Size    int        `json:"size" form:"size,default=10"`
}

// AppVO 应用视图对象
type AppVO struct {
	ID          types.Long   `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Creator     types.Long   `json:"creator"`
	Status      string       `json:"status"`
	GroupIDs    []types.Long `json:"group_ids,omitempty"`
	EnvCount    int          `json:"env_count"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}

// TableName 返回应用表名
func (Application) TableName() string {
	return "app"
}

// TableName 返回应用分组表名
func (AppGroup) TableName() string {
	return "app_group"
}

// TableName 返回应用分组关联表名
func (AppGroupRelation) TableName() string {
	return "relation_app_group_app"
}

// TableName 返回应用环境表名
func (AppEnv) TableName() string {
	return "app_env"
}

// TableName 返回应用HPA表名
func (AppHPA) TableName() string {
	return "app_hpa"
}

// TableName 返回镜像仓库表名
func (ImageRegistry) TableName() string {
	return "app_image_registry"
}

// TableName 返回发布计划表名
func (ReleasePlan) TableName() string {
	return "app_release_plan"
}

// TableName 返回部署历史表名
func (Deployment) TableName() string {
	return "deploy_history"
}

// TableName 返回部署步骤表名
func (DeploymentStep) TableName() string {
	return "deploy_steps"
}
