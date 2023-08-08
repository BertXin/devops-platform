package init

import (
	"devops-platform/internal/common/web/internal/domain"
	"devops-platform/internal/common/web/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanGinEngineLifecycle, &service.HttpServerLifecycle{})
	beans.Register(domain.BeanAuthenticationVerify, service.Verify)
}
