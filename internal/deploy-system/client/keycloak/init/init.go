package init

import (
	"devops-platform/internal/deploy-system/client/keycloak/internal/domain"
	"devops-platform/internal/deploy-system/client/keycloak/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanService, &service.Service{})
}
