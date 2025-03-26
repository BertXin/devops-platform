package init

import (
	"devops-platform/internal/deploy-system/organization/internal/controller"
	"devops-platform/internal/deploy-system/organization/internal/domain"
	"devops-platform/internal/deploy-system/organization/internal/repository"
	"devops-platform/internal/deploy-system/organization/internal/service"
	"devops-platform/pkg/beans"

	"github.com/sirupsen/logrus"
)

// 使用标准的init函数进行初始化
func init() {
	// 注册仓库
	beans.Register(domain.BeanDepartmentRepository, repository.NewRepository())

	// 注册服务
	beans.Register(domain.BeanDepartmentService, service.NewDepartmentService())

	// 注册控制器
	beans.Register(domain.BeanDepartmentController, controller.NewDepartmentController())

	logrus.Info("组织结构模块初始化完成")
}
