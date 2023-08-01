package init

import (
	"devops-platform/internal/common/repository"
	"devops-platform/internal/common/service"
	"devops-platform/internal/deploy-system/user/internal/controller"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanRepository, &repository.Repository{})
	beans.Register(domain.BeanService, &service.Service{})
	beans.Register(domain.BeanController, &controller.Controller{})
}
