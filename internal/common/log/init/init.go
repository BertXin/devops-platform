package init

import (
	"devops-platform/internal/common/log/internal/domain"
	"devops-platform/internal/common/log/internal/service"
	"devops-platform/pkg/beans"

	"github.com/sirupsen/logrus"
)

func init() {
	// 注册日志组件
	beans.Register(domain.BeanLog, &service.Logger{})

	// 注册Logger实例，为其他服务提供依赖注入
	logger := logrus.New()
	beans.Register("Logger", logger)
}
