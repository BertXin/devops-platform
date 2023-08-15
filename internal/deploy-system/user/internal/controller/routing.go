package controller

import (
	"devops-platform/internal/common/web"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (c *Controller) Inject(getBean func(string) interface{}) {
	c.injectQuery(getBean)
	c.injectService(getBean)
	c.injectRouting(getBean)
}

func (c *Controller) injectRouting(getBean func(string) interface{}) {

	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}
	routes := router.Group("user")
	{
		routes.GET("", c.FindByName)
		routes.GET("/:id", c.GetByID)

		routes.POST("", c.CreateUser)
		routes.PATCH("/:id/role", c.ModifyUserRoleByID)
		routes.PATCH("/:id/status", c.ModifyUserStatusByID)
		routes.PATCH("/:id/password", c.ModifyUserPasswordByID)
	}
	//web.AddIgnoreUrls("/user")
}
