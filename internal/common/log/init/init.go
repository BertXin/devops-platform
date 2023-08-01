package init

import (
	"devops-platform/internal/common/log/internal/domain"
	"devops-platform/internal/common/log/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanLog, &service.Logger{})
}
