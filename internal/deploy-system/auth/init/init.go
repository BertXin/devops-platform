package init

import (
	"devops-platform/internal/deploy-system/auth/internal/controller"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/internal/deploy-system/auth/internal/repository"
	"devops-platform/internal/deploy-system/auth/internal/service"
	"devops-platform/pkg/beans"

	"github.com/sirupsen/logrus"
)

func init() {
	// 注册仓储层
	beans.Register(domain.BeanRepository, repository.NewRepository())

	// 注册服务层
	beans.Register(domain.BeanService, service.NewAuthService())

	// 注册用户查询服务
	beans.Register(domain.BeanUserQuery, service.NewUserQuery())

	// 注册控制器
	beans.Register(domain.BeanController, controller.NewAuthController())

	logrus.Info("用户模块初始化完成")
}
