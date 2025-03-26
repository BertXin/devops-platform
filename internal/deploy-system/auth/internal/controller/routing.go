package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/internal/deploy-system/auth/internal/service"
	"devops-platform/internal/deploy-system/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Inject 实现依赖注入
func (c *AuthController) Inject(getBean func(string) interface{}) {
	// 注入服务
	c.InjectService(getBean)
	// 注入路由
	c.injectRouting(getBean)
}

// InjectService 注入服务
func (c *AuthController) InjectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanService).(*service.AuthService)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanService)
		return
	}
	c.Service = service
}

// injectRouting 注入路由
func (c *AuthController) injectRouting(getBean func(string) interface{}) {
	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}

	// 认证相关路由组
	authGroup := router.Group("/auth")
	{
		// 无需认证的路由
		authGroup.POST("/login", c.Login)
		authGroup.POST("/register", c.Register)

		// 健康检查路由
		authGroup.GET("/health", func(c *gin.Context) {
			web.HealthCheck(c, domain.BeanModuleName)
		})

		// 需要认证的路由
		protectedGroup := authGroup.Group("")
		protectedGroup.Use(middleware.JWTAuth())
		protectedGroup.POST("/logout", c.Logout)
		protectedGroup.GET("/me", c.GetUserInfo)
		protectedGroup.POST("/change-password", c.ChangePassword)
	}

	// 添加到忽略URL列表
	web.AddIgnoreUrls("/auth/login", "/auth/register", "/auth/health")
}
