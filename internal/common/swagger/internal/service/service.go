package service

import (
	_ "devops-platform/docs"
	"devops-platform/internal/common/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Swagger struct {
}

func (s *Swagger) Inject(getBean func(string) interface{}) {

	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
