package init

import (
	"devops-platform/internal/deploy-system/department/internal/controller"
	"devops-platform/internal/deploy-system/department/internal/domain"
	"devops-platform/internal/deploy-system/department/internal/repository"
	"devops-platform/internal/deploy-system/department/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanRepository, &repository.Repository{})
	beans.Register(domain.BeanService, &service.Service{})
	beans.Register(domain.BeanController, &controller.Controller{})
}
