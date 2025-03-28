package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// injectRouting 注入路由
func (c *AppController) injectRouting(getBean func(string) interface{}) {
	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}

	// API前缀
	apiPrefix := "/api/v1"

	// 使用JWT认证中间件的路由组
	authRouter := router.Group(apiPrefix)
	authRouter.Use(middleware.JWTAuth())

	// 应用管理路由
	appsGroup := authRouter.Group("/apps")
	{
		appsGroup.GET("", c.ListApplications)         // 查询应用列表
		appsGroup.POST("", c.CreateApplication)       // 创建应用
		appsGroup.GET("/:id", c.GetApplication)       // 获取应用详情
		appsGroup.PUT("/:id", c.UpdateApplication)    // 更新应用
		appsGroup.DELETE("/:id", c.DeleteApplication) // 删除应用

		// 应用HPA配置 - 修改参数名为 :id 以匹配其他路由
		appsGroup.POST("/:id/hpa", c.ConfigureHPA) // 配置HPA
		appsGroup.GET("/:id/hpa", c.GetAppHPA)     // 获取HPA配置
	}

	// 应用分组路由
	groupsGroup := authRouter.Group("/app-groups")
	{
		groupsGroup.GET("", c.ListAppGroups)       // 查询分组列表
		groupsGroup.POST("", c.CreateAppGroup)     // 创建分组
		groupsGroup.POST("/apps", c.AddAppToGroup) // 添加应用到分组
	}

	// 环境管理路由
	envsGroup := authRouter.Group("/envs")
	{
		envsGroup.GET("", c.ListEnvironments)   // 查询环境列表
		envsGroup.POST("", c.CreateEnvironment) // 创建环境
	}

	// 发布管理路由
	releasesGroup := authRouter.Group("/releases")
	{
		releasesGroup.POST("", c.CreateReleasePlan)              // 创建发布计划
		releasesGroup.POST("/:id/execute", c.ExecuteReleasePlan) // 执行发布计划
	}

	// 部署历史路由
	deploymentsGroup := authRouter.Group("/deployments")
	{
		deploymentsGroup.GET("", c.ListDeployments)                  // 查询部署历史
		deploymentsGroup.GET("/:id", c.GetDeployment)                // 获取部署详情
		deploymentsGroup.POST("/:id/rollback", c.RollbackDeployment) // 回滚部署
	}

}
