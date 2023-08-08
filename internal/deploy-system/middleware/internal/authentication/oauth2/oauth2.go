package oauth2

import (
	"devops-platform/internal/common/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const BeanAuthenticationOAuth2 = "BeanAuthenticationOAuth2"

func OAuth2(getBean func(string) interface{}) {
	authentication, ok := getBean(BeanAuthenticationOAuth2).(func(*gin.Context))
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", BeanAuthenticationOAuth2)
		return
	}
	router, ok := getBean(web.BeanGinEngine).(gin.IRoutes)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}
	router.Use(authentication)
}
