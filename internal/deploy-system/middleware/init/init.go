package init

import (
	"devops-platform/internal/deploy-system/middleware/internal/authentication/oauth2"
	"devops-platform/internal/deploy-system/middleware/internal/authentication/verify"
	"devops-platform/internal/deploy-system/middleware/internal/domain"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register(domain.BeanAuthenticationChain, func(getBean func(string) interface{}) {
		oauth2.OAuth2(getBean)
		verify.Verify(getBean)
	})

}
