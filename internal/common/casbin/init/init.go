package init

import (
	"devops-platform/internal/common/casbin/internal/domain"
	"devops-platform/internal/common/casbin/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	// 注册Casbin Enforcer (使用正式的Bean名称)
	beans.Register(domain.BeanEnforcer, &service.CasbinEnforcer{})
}
