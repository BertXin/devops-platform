package service

import (
	"context"
	"devops-platform/internal/common/service"
	"errors"
	"github.com/sirupsen/logrus"

	"devops-platform/internal/deploy-system/application/internal/domain"
	"devops-platform/internal/deploy-system/application/internal/repository"
	"devops-platform/pkg/types"
)

// AppService 应用管理服务实现
type AppService struct {
	service.Service
	Repo   *repository.AppRepository
	Logger *logrus.Logger
}

// NewAppService 创建应用管理服务实例
func NewAppService() *AppService {
	return &AppService{}
}

// CreateApplication 创建应用
func (s *AppService) CreateApplication(ctx context.Context, command *domain.CreateAppCommand) (id types.Long, err error) {
	// 检查应用名称是否已存在
	existApp, err := s.Repo.GetApplicationByName(ctx, command.Name)
	if err == nil && existApp != nil {
		return 0, errors.New("应用名称已存在")
	}

	ctx, err = s.BeginTransaction(ctx, "create service app")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create service app")
	}()
	// 创建应用
	app := &domain.Application{
		Name:        command.Name,
		Description: command.Description,
		Creator:     command.Creator,
		Status:      domain.AppStatusActive,
	}

	return s.Repo.CreateApplication(ctx, app)
}

// UpdateApplication 更新应用
func (s *AppService) UpdateApplication(ctx context.Context, command *domain.UpdateAppCommand) (err error) {
	// 获取应用
	app, err := s.Repo.GetApplicationByID(ctx, command.ID)
	if err != nil {
		return err
	}

	// 如果修改了名称，检查名称是否已存在
	if command.Name != "" && command.Name != app.Name {
		existApp, err := s.Repo.GetApplicationByName(ctx, command.Name)
		if err == nil && existApp != nil && existApp.ID != command.ID {
			return errors.New("应用名称已存在")
		}
		app.Name = command.Name
	}

	// 更新其他字段
	if command.Description != "" {
		app.Description = command.Description
	}
	if command.Status != "" {
		app.Status = command.Status
	}
	ctx, err = s.BeginTransaction(ctx, "create service app")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create service app")
	}()
	return s.Repo.UpdateApplication(ctx, app)
}

// GetApplicationByID 根据ID获取应用
func (s *AppService) GetApplicationByID(ctx context.Context, id types.Long) (*domain.Application, error) {
	return s.Repo.GetApplicationByID(ctx, id)
}

// ListApplications 查询应用列表
func (s *AppService) ListApplications(ctx context.Context, query *domain.AppQuery) ([]*domain.AppVO, int64, error) {
	return s.Repo.ListApplications(ctx, query)
}

// DeleteApplication 删除应用
func (s *AppService) DeleteApplication(ctx context.Context, id types.Long) (err error) {
	ctx, err = s.BeginTransaction(ctx, "create service app")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create service app")
	}()
	return s.Repo.DeleteApplication(ctx, id)
}

// CreateAppGroup 创建应用分组
func (s *AppService) CreateAppGroup(ctx context.Context, name, description string) (id types.Long, err error) {
	group := &domain.AppGroup{
		Name:        name,
		Description: description,
	}
	ctx, err = s.BeginTransaction(ctx, "create service app")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create service app")
	}()
	return s.Repo.CreateAppGroup(ctx, group)
}

// UpdateAppGroup 更新应用分组
func (s *AppService) UpdateAppGroup(ctx context.Context, id types.Long, name, description string) (err error) {
	group, err := s.Repo.GetAppGroupByID(ctx, id)
	if err != nil {
		return err
	}

	if name != "" {
		group.Name = name
	}
	if description != "" {
		group.Description = description
	}
	ctx, err = s.BeginTransaction(ctx, "create service app")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create service app")
	}()
	return s.Repo.UpdateAppGroup(ctx, group)
}

// ListAppGroups 查询应用分组列表
func (s *AppService) ListAppGroups(ctx context.Context) ([]*domain.AppGroup, error) {
	return s.Repo.ListAppGroups(ctx)
}

// DeleteAppGroup 删除应用分组
func (s *AppService) DeleteAppGroup(ctx context.Context, id types.Long) error {
	return s.Repo.DeleteAppGroup(ctx, id)
}

// AddAppToGroup 添加应用到分组
func (s *AppService) AddAppToGroup(ctx context.Context, appID, groupID types.Long) (err error) {
	// 检查应用和分组是否存在
	_, err = s.Repo.GetApplicationByID(ctx, appID)
	if err != nil {
		return errors.New("应用不存在")
	}

	_, err = s.Repo.GetAppGroupByID(ctx, groupID)
	if err != nil {
		return errors.New("分组不存在")
	}
	ctx, err = s.BeginTransaction(ctx, "create service app")
	if err != nil {
		return
	}
	defer func() {
		err = s.FinishTransaction(ctx, err, "create service app")
	}()
	return s.Repo.AddAppToGroup(ctx, appID, groupID)
}

// RemoveAppFromGroup 从分组中移除应用
func (s *AppService) RemoveAppFromGroup(ctx context.Context, appID, groupID types.Long) error {
	return s.Repo.RemoveAppFromGroup(ctx, appID, groupID)
}

// GetAppGroups 获取应用所属的分组
func (s *AppService) GetAppGroups(ctx context.Context, appID types.Long) ([]*domain.AppGroup, error) {
	return s.Repo.GetAppGroups(ctx, appID)
}

// CreateAppEnv 创建应用环境
func (s *AppService) CreateAppEnv(ctx context.Context, command *domain.CreateEnvCommand) (types.Long, error) {
	env := &domain.AppEnv{
		Name:        command.Name,
		ClusterID:   command.ClusterID,
		Namespace:   command.Namespace,
		Description: command.Description,
	}
	return s.Repo.CreateAppEnv(ctx, env)
}

// UpdateAppEnv 更新应用环境
func (s *AppService) UpdateAppEnv(ctx context.Context, id types.Long, name, namespace, description string) error {
	env, err := s.Repo.GetAppEnvByID(ctx, id)
	if err != nil {
		return err
	}

	if name != "" {
		env.Name = name
	}
	if namespace != "" {
		env.Namespace = namespace
	}
	if description != "" {
		env.Description = description
	}

	return s.Repo.UpdateAppEnv(ctx, env)
}

// GetAppEnvByID 根据ID获取应用环境
func (s *AppService) GetAppEnvByID(ctx context.Context, id types.Long) (*domain.AppEnv, error) {
	return s.Repo.GetAppEnvByID(ctx, id)
}

// ListAppEnvs 查询应用环境列表
func (s *AppService) ListAppEnvs(ctx context.Context) ([]*domain.AppEnv, error) {
	return s.Repo.ListAppEnvs(ctx)
}

// DeleteAppEnv 删除应用环境
func (s *AppService) DeleteAppEnv(ctx context.Context, id types.Long) error {
	return s.Repo.DeleteAppEnv(ctx, id)
}

// CreateImageRegistry 创建镜像仓库
func (s *AppService) CreateImageRegistry(ctx context.Context, name, url, username, password, email string) (types.Long, error) {
	registry := &domain.ImageRegistry{
		Name:     name,
		URL:      url,
		Username: username,
		Password: password,
		Email:    email,
	}
	return s.Repo.CreateImageRegistry(ctx, registry)
}

// UpdateImageRegistry 更新镜像仓库
func (s *AppService) UpdateImageRegistry(ctx context.Context, id types.Long, name, url, username, password, email string) error {
	registry, err := s.Repo.GetImageRegistryByID(ctx, id)
	if err != nil {
		return err
	}

	if name != "" {
		registry.Name = name
	}
	if url != "" {
		registry.URL = url
	}
	if username != "" {
		registry.Username = username
	}
	if password != "" {
		registry.Password = password
	}
	if email != "" {
		registry.Email = email
	}

	return s.Repo.UpdateImageRegistry(ctx, registry)
}

// ListImageRegistries 查询镜像仓库列表
func (s *AppService) ListImageRegistries(ctx context.Context) ([]*domain.ImageRegistry, error) {
	return s.Repo.ListImageRegistries(ctx)
}

// DeleteImageRegistry 删除镜像仓库
func (s *AppService) DeleteImageRegistry(ctx context.Context, id types.Long) error {
	return s.Repo.DeleteImageRegistry(ctx, id)
}
