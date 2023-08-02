package repository

import (
	"context"
	"devops-platform/internal/common/repository"
	"fmt"
	"io"
	apps "k8s.io/api/apps/v1"
	authenticationv1 "k8s.io/api/authentication/v1"
	core "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"time"
)

type Repository struct {
	client *kubernetes.Clientset
	*repository.Repository
}

/**
 * 获取pod日志
 */
func (r *Repository) GetPodLogs(ctx context.Context, namespace string, podName string, containerName string) (log string, err error) {
	request := r.client.CoreV1().Pods(namespace).GetLogs(podName, &core.PodLogOptions{
		Container: containerName,
	})
	rawLog, err := request.DoRaw(ctx)
	if err != nil {
		return
	}
	log = string(rawLog)
	return
}

/**
 * 获取实时pod日志
 */
func (r *Repository) GetRealTimePodLogs(ctx context.Context, namespace string, podName string, containerName string) (readCloser io.ReadCloser, err error) {
	request := r.client.CoreV1().Pods(namespace).GetLogs(podName, &core.PodLogOptions{
		Container: containerName,
		Follow:    true,
	})
	readCloser, err = request.Stream(ctx)
	return
}

/**
 * 重启deployment
 */
func (r *Repository) RestartDeployment(ctx context.Context, namespace string, deploymentName string, systemName string) (err error) {
	deployment, err := r.client.AppsV1().Deployments(namespace).Get(ctx, deploymentName, meta.GetOptions{})
	if err != nil {
		return
	}
	// 更新 Timestamp 触发滚动更新
	if deployment.Spec.Template.Annotations == nil {
		deployment.Spec.Template.Annotations = map[string]string{}
	}
	deployment.Spec.Template.Annotations[fmt.Sprintf("%s/restartedAt", systemName)] = time.Now().String()

	_, err = r.client.AppsV1().Deployments(namespace).Update(ctx, deployment, meta.UpdateOptions{})

	return
}

/**
 * 获取deployment
 */
func (r *Repository) GetDeployment(ctx context.Context, namespace string, deploymentName string) (deployment *apps.Deployment, err error) {
	deployment, err = r.client.AppsV1().Deployments(namespace).Get(ctx, deploymentName, meta.GetOptions{})
	return
}

/**
 * 获取deployments
 */
func (r *Repository) GetDeployments(ctx context.Context, namespace string) (deployments *apps.DeploymentList, err error) {
	deployments, err = r.client.AppsV1().Deployments(namespace).List(ctx, meta.ListOptions{})
	return
}

/**
 * 获取deployment是否就绪
 */
func (r *Repository) DeploymentReady(ctx context.Context, namespace string, deploymentName string) (status bool, err error) {
	deployment, err := r.GetDeployment(ctx, namespace, deploymentName)
	if err != nil {
		return
	}
	/*
	 * 通过deployment.Status.Conditions判断是否部署成功
	 */
	for _, condition := range deployment.Status.Conditions {
		if condition.Type != apps.DeploymentProgressing {
			continue
		}
		if condition.Status != core.ConditionTrue {
			continue
		}
		if condition.Reason != "NewReplicaSetAvailable" {
			continue
		}
		status = true
		break
	}

	return
}

/**
 * 获取deploymentConfig有问题
 */
func (r *Repository) GetDeploymentConfig(ctx context.Context, namespace string, deploymentName string) (confMap *core.ConfigMap, err error) {
	confMap, err = r.client.CoreV1().ConfigMaps(namespace).Get(ctx, deploymentName, meta.GetOptions{})
	return
}

/**
 * 创建CreateServiceAccount
 */
func (r *Repository) CreateServiceAccount(ctx context.Context, namespaces string, saName string) (*v1.ServiceAccount, error) {
	sa := &v1.ServiceAccount{}
	sa.Namespace = namespaces
	sa.Name = saName
	sa.GenerateName = saName
	return r.client.CoreV1().ServiceAccounts(namespaces).Create(ctx, sa, meta.CreateOptions{})
}

/**
 * 删除CreateServiceAccount
 */
func (r *Repository) DeleteServiceAccount(ctx context.Context, namespaces string, saName string) (err error) {
	return r.client.CoreV1().ServiceAccounts(namespaces).Delete(ctx, saName, meta.DeleteOptions{})
}

/**
 * 创建CreateServiceAccountToken
 */
func (r *Repository) CreateServiceAccountToken(ctx context.Context, namespaces string, saName string, expiration *int64) (*authenticationv1.TokenRequest, error) {
	tokenRequest := &authenticationv1.TokenRequest{}
	if namespaces != "" {
		tokenRequest.Namespace = namespaces
	}
	tokenRequest.Name = "token"
	tokenRequest.Spec.ExpirationSeconds = expiration
	return r.client.CoreV1().ServiceAccounts(namespaces).CreateToken(ctx, saName, tokenRequest, meta.CreateOptions{})
}

/**
 * 创建CreateRole
 */
func (r *Repository) CreateRole(ctx context.Context, namespaces string, rules []rbacv1.PolicyRule, name string) (*rbacv1.Role, error) {
	role := &rbacv1.Role{}
	role.Namespace = namespaces
	role.Name = name
	role.Rules = rules
	return r.client.RbacV1().Roles(namespaces).Create(ctx, role, meta.CreateOptions{})
}

/**
 * 删除DeleteRole
 */
func (r *Repository) DeleteRole(ctx context.Context, namespaces string, name string) (err error) {
	return r.client.RbacV1().Roles(namespaces).Delete(ctx, name, meta.DeleteOptions{})
}

/**
 * 创建CreateRoleBinding
 */
func (r *Repository) CreateRoleBinding(ctx context.Context, subject rbacv1.Subject, roleRef rbacv1.RoleRef, namespace string, name string) (*rbacv1.RoleBinding, error) {
	binding := &rbacv1.RoleBinding{}
	binding.Name = name
	binding.Namespace = namespace

	var subjects []rbacv1.Subject
	subjects = append(subjects, subject)
	binding.Subjects = subjects

	binding.RoleRef = roleRef
	return r.client.RbacV1().RoleBindings(namespace).Create(ctx, binding, meta.CreateOptions{})
}

/**
 * 删除DeleteRoleBinding
 */
func (r *Repository) DeleteRoleBinding(ctx context.Context, namespace string, name string) (err error) {
	return r.client.RbacV1().RoleBindings(namespace).Delete(ctx, name, meta.DeleteOptions{})
}

/**
 * 更新UpdateRoleBinding
 */
func (r *Repository) UpdateRoleBinding(ctx context.Context, binding *rbacv1.RoleBinding, namespace string) (*rbacv1.RoleBinding, error) {
	return r.client.RbacV1().RoleBindings(namespace).Update(ctx, binding, meta.UpdateOptions{})
}

/**
 * 查看GetRoleBinding
 */
func (r *Repository) GetRoleBinding(ctx context.Context, namespace string, name string) (*rbacv1.RoleBinding, error) {
	return r.client.RbacV1().RoleBindings(namespace).Get(ctx, name, meta.GetOptions{})
}

/**
 * 创建CreateClusterRole
 */
func (r *Repository) CreateClusterRole(ctx context.Context, rules []rbacv1.PolicyRule, name string) (*rbacv1.ClusterRole, error) {
	role := &rbacv1.ClusterRole{}
	role.Name = name
	role.Rules = rules
	return r.client.RbacV1().ClusterRoles().Update(ctx, role, meta.UpdateOptions{})
}

/**
 * 删除DeleteClusterRole
 */
func (r *Repository) DeleteClusterRole(ctx context.Context, name string) (err error) {
	return r.client.RbacV1().ClusterRoles().Delete(ctx, name, meta.DeleteOptions{})
}
