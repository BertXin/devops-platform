package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/internal/deploy-system/authorization/internal/service"
	"devops-platform/internal/deploy-system/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Inject 实现依赖注入
func (c *AuthorizationController) Inject(getBean func(string) interface{}) {
	// 注入服务
	c.InjectService(getBean)
	// 注入路由
	c.injectRouting(getBean)
}

// InjectService 注入服务
func (c *AuthorizationController) InjectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanAuthorizationService).(*service.AuthorizationService)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanAuthorizationService)
		return
	}
	c.Service = service
}

// injectRouting 注入路由
func (c *AuthorizationController) injectRouting(getBean func(string) interface{}) {
	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}

	//获取其他控制器
	roleController := NewRoleController()
	permController := NewPermissionController()

	// 注入其他控制器的服务
	roleController.InjectService(getBean)
	permController.InjectService(getBean)

	// 权限认证相关路由，需要登录认证
	authzRouter := router.Group("/api/v1/authorization")
	authzRouter.Use(middleware.JWTAuth())
	{

		// 当前用户菜单
		authzRouter.GET("/menus", c.GetUserMenus)
		// 权限检查
		authzRouter.GET("/check/:permission", c.HasPermission)

		// === 用户相关路由 ===
		// 1. 先注册用户根路由组
		usersRouter := authzRouter.Group("/users")
		// 2. 按照路径层级依次创建子路由组
		{
			// 用户ID参数路由组
			userIDRouter := usersRouter.Group("/:user_id")
			{
				// 基本操作
				userIDRouter.GET("/roles", c.GetUserRoles)
				userIDRouter.POST("/roles", c.AssignRolesToUser)
				userIDRouter.GET("/permissions", c.GetUserPermissions)

				// 嵌套的带角色ID的路由
				userRolesRouter := userIDRouter.Group("/roles")
				{
					userRolesRouter.DELETE("/:role_id", c.RemoveRoleFromUser)
				}
			}
		}

		// === 角色相关路由 ===
		// 1. 先注册角色根路由组
		rolesRouter := authzRouter.Group("/roles")
		{
			// 2. 先注册不带参数的路由
			rolesRouter.POST("", roleController.CreateRole)
			rolesRouter.GET("", roleController.ListRoles)

			// 3. 以独立路由方式注册ID参数路由，避免使用Group
			rolesRouter.PUT("/:id", roleController.UpdateRole)
			rolesRouter.DELETE("/:id", roleController.DeleteRole)
			rolesRouter.GET("/:id", roleController.GetRoleByID)
			rolesRouter.POST("/:id/permissions", roleController.AssignPermissionsToRole)
			rolesRouter.GET("/:id/permissions", roleController.GetRolePermissions)
		}

		// === 权限相关路由 ===
		// 1. 首先注册不带参数或带固定路径的路由
		permRoute := authzRouter.Group("/permissions")
		permRoute.POST("", permController.CreatePermission)
		permRoute.GET("", permController.ListPermissions)
		permRoute.GET("/tree", permController.GetPermissionTree)

		// 2. 单独注册带ID参数的路由，使用唯一的路径变体防止冲突
		permRoute.PUT("/detail/:id", permController.UpdatePermission)
		permRoute.DELETE("/detail/:id", permController.DeletePermission)
		permRoute.GET("/detail/:id", permController.GetPermissionByID)
	}
}
