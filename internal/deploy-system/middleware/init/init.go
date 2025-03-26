package init

import (
	"devops-platform/internal/deploy-system/middleware/internal/authentication/jwt"
	"devops-platform/internal/deploy-system/middleware/internal/domain"
	"devops-platform/pkg/beans"

	"github.com/sirupsen/logrus"
)

func init() {
	beans.Register(domain.BeanAuthenticationChain, func(getBean func(string) interface{}) {
		// 注册JWT认证中间件
		jwt.JWT(getBean)

		// 注册OAuth2认证中间件
		//oauth2.OAuth2(getBean)

		// 注册验证中间件
		//verify.Verify(getBean)

		logrus.Info("认证中间件初始化完成")
	})
}
