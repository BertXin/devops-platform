package init

import (
	"devops-platform/internal/deploy-system/application/internal/controller"
	"devops-platform/internal/deploy-system/application/internal/domain"
	"devops-platform/internal/deploy-system/application/internal/repository"
	"devops-platform/internal/deploy-system/application/internal/service"
	"devops-platform/pkg/beans"

	"github.com/sirupsen/logrus"
)

func init() {
	// 注册仓储层
	beans.Register(domain.BeanAppRepository, repository.NewAppRepository())

	// 注册服务层
	beans.Register(domain.BeanAppService, service.NewAppService())
	beans.Register(domain.BeanDeployService, service.NewDeployService())
	beans.Register(domain.BeanAppQuery, service.NewAppQuery())

	// 注册控制器
	beans.Register(domain.BeanController, controller.NewAppController())

	logrus.Info("应用管理模块初始化完成")
}
