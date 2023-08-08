package verify

import (
	"devops-platform/internal/common/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Verify(getBean func(string) interface{}) {

	verifier, ok := getBean(web.BeanAuthenticationVerify).(func(*gin.Context))
	if !ok {
		logrus.Infof("初始化时获取[%s]失败", web.BeanAuthenticationVerify)
		return
	}

	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}

	router.Use(verifier)
}
