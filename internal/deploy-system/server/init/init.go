package init

import (
	"devops-platform/internal/deploy-system/server/internal/domain"
	"devops-platform/internal/deploy-system/server/internal/repository"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanRepository, &repository.Repository{})
}
