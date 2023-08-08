package init

import (
	"devops-platform/internal/deploy-system/login/internal/controller"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/login/internal/service"
	"devops-platform/internal/deploy-system/middleware"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanService, &service.KeyCloakService{})
	c := controller.Controller{}
	beans.Register(domain.BeanController, &c)
	beans.Register(middleware.BeanAuthenticationOAuth2, c.Authentication)
}
