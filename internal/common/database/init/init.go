package init

import (
	"devops-platform/internal/common/database/internal/service"
	"devops-platform/pkg/beans"
)

//注册bean
const (
	beanDBCloser = "dbCloser"
)

func init() {
	beans.Register(beanDBCloser, &service.DB{})
}
