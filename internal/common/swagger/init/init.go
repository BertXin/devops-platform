package swagger

import (
	"devops-platform/internal/common/swagger/internal/service"
	"devops-platform/pkg/beans"
)

//注册bean
const (
	beanSwagger = "swagger"
)

func init() {
	beans.Register(beanSwagger, &service.Swagger{})
}
