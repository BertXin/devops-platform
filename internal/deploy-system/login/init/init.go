package init

import (
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/login/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanService, &service.KeyCloakService{})
}
