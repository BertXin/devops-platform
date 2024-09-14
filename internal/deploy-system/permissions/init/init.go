package init

import (
	"devops-platform/internal/deploy-system/permissions/internal/controller"
	"devops-platform/internal/deploy-system/permissions/internal/domain"
	"devops-platform/internal/deploy-system/permissions/internal/repository"
	"devops-platform/internal/deploy-system/permissions/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanRepository, &repository.Repository{})
	beans.Register(domain.BeanService, &service.Service{})
	beans.Register(domain.BeanController, &controller.Controller{})
}
