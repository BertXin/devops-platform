package service

import (
	"context"
	"devops-platform/internal/common/config"
	"devops-platform/internal/deploy-system/client/kubernetes/internal/domain"
	"devops-platform/internal/deploy-system/client/kubernetes/internal/repository"
	"github.com/sirupsen/logrus"
	"io"
	apps "k8s.io/api/apps/v1"
	authenticationv1 "k8s.io/api/authentication/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
)

var cfg domain.Config

func SetConfig(getBean func(string) interface{}) {

	ok := false

	cfg, ok = getBean(config.BeanTekton).(domain.Config)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", config.BeanTekton)
		return
	}
}

type Service struct {
	repository *repository.Repository
}

func New(cfg domain.KubernetesConfig) (*Service, error) {
	repo, err := repository.New(cfg)
	if err != nil {
		return nil, err
	}
	return &Service{
		repository: repo,
	}, nil
}

/**
 * 获取pod日志
 */
func (s *Service) GetPodLogs(ctx context.Context, namespace string, podName string, containerName string) (string, error) {
	return s.repository.GetPodLogs(ctx, namespace, podName, containerName)
}

/**
 * 获取pod实时日志
 */
func (s *Service) GetRealTimePodLogs(ctx context.Context, namespace string, podName string, containerName string) (readCloser io.ReadCloser, err error) {
	return s.repository.GetRealTimePodLogs(ctx, namespace, podName, containerName)
}

/**
 * 获取deployment是否就绪
 */
func (s *Service) DeploymentReady(ctx context.Context, namespace string, deploymentName string) (status bool, err error) {
	return s.repository.DeploymentReady(ctx, namespace, deploymentName)
}

/**
 * 获取deployments
 */
func (s *Service) GetDeployments(ctx context.Context, namespace string) (deployments []apps.Deployment, err error) {
	var (
		deploymentList *apps.DeploymentList
	)
	if deploymentList, err = s.repository.GetDeployments(ctx, namespace); err != nil {
		return
	}
	deployments = deploymentList.Items

	return
}

/**
 * 获取deployment
 */
func (s *Service) GetDeployment(ctx context.Context, namespace string, deploymentName string) (deployment *apps.Deployment, err error) {
	return s.repository.GetDeployment(ctx, namespace, deploymentName)
}

/**
 * 重启deployment
 */
func (s *Service) RestartDeployment(ctx context.Context, namespaces string, deploymentName string) (err error) {
	return s.repository.RestartDeployment(ctx, namespaces, deploymentName, cfg.GetSystemName())
}

func (s *Service) CreateServiceAccount(ctx context.Context, namespaces string, saName string) (*v1.ServiceAccount, error) {
	return s.repository.CreateServiceAccount(ctx, namespaces, saName)
}

func (s *Service) DeleteServiceAccount(ctx context.Context, namespaces string, saName string) (err error) {
	return s.repository.DeleteServiceAccount(ctx, namespaces, saName)
}

func (s *Service) CreateRole(ctx context.Context, namespace string, rules []rbacv1.PolicyRule, name string) (err error) {
	if namespace == "" {
		_, err = s.repository.CreateClusterRole(ctx, rules, name)
	} else {
		_, err = s.repository.CreateRole(ctx, namespace, rules, name)
	}
	return
}

func (s *Service) DeleteRole(ctx context.Context, namespace string, name string) (err error) {
	if namespace == "" {
		err = s.repository.DeleteClusterRole(ctx, name)
	} else {
		err = s.repository.DeleteRole(ctx, namespace, name)
	}
	return
}

func (s *Service) GetRoleBinding(ctx context.Context, namespace string, name string) (*rbacv1.RoleBinding, error) {
	return s.repository.GetRoleBinding(ctx, namespace, name)
}

func (s *Service) UpdateRoleBinding(ctx context.Context, binding *rbacv1.RoleBinding, namespace string) (*rbacv1.RoleBinding, error) {
	return s.repository.UpdateRoleBinding(ctx, binding, namespace)
}

func (s *Service) CreateRoleBinding(ctx context.Context, subject rbacv1.Subject, roleRef rbacv1.RoleRef, namespace string, name string) (*rbacv1.RoleBinding, error) {
	return s.repository.CreateRoleBinding(ctx, subject, roleRef, namespace, name)
}

func (s *Service) DeleteRoleBinding(ctx context.Context, namespace string, name string) (err error) {
	return s.repository.DeleteRoleBinding(ctx, namespace, name)
}

func (s *Service) CreateServiceAccountToken(ctx context.Context, namespaces string, saName string, expiration *int64) (*authenticationv1.TokenRequest, error) {
	return s.repository.CreateServiceAccountToken(ctx, namespaces, saName, expiration)
}
