package init

import (
	"devops-platform/internal/deploy-system/authorization/internal/controller"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/internal/deploy-system/authorization/internal/repository"
	"devops-platform/internal/deploy-system/authorization/internal/service"
	"devops-platform/pkg/beans"

	"github.com/sirupsen/logrus"
)

// 使用标准的init函数进行初始化
func init() {
	// 注册仓储
	beans.Register(domain.BeanRepository, repository.NewRepository())

	// 注册权限服务
	beans.Register(domain.BeanAuthorizationService, service.NewAuthorizationService())

	// 注册角色服务
	beans.Register(domain.BeanRoleService, service.NewRoleService())

	// 注册权限服务
	beans.Register(domain.BeanPermissionService, service.NewPermissionService())

	// 注册控制器
	beans.Register(domain.BeanAuthorizationController, controller.NewAuthorizationController())

	logrus.Info("权限模块初始化完成")
}
