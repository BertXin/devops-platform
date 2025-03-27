package init

import (
	"devops-platform/internal/common/casbin/internal/domain"
	"devops-platform/internal/common/casbin/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	// 注册Casbin Enforcer
	beans.Register(domain.BeanEnforcer, &service.CasbinEnforcer{})
	// 如果需要添加生命周期管理，可以像web模块一样增加生命周期Bean
}
