package init

import (
	"devops-platform/internal/deploy-system/rbac/internal/controller"
	"devops-platform/internal/deploy-system/rbac/internal/domain"
	"devops-platform/internal/deploy-system/rbac/internal/repository"
	"devops-platform/internal/deploy-system/rbac/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanRepository, &repository.Repository{})
	beans.Register(domain.BeanService, &service.Service{})
	beans.Register(domain.BeanController, &controller.Controller{})
}
