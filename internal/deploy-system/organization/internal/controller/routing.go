package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/middleware"
	"devops-platform/internal/deploy-system/organization/internal/domain"
	"devops-platform/internal/deploy-system/organization/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Inject 实现依赖注入
func (c *DepartmentController) Inject(getBean func(string) interface{}) {
	// 注入服务
	c.InjectService(getBean)
	// 注入路由
	c.injectRouting(getBean)
}

// InjectService 注入服务
func (c *DepartmentController) InjectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanDepartmentService).(*service.DepartmentService)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanDepartmentService)
		return
	}
	c.Service = service
}

// injectRouting 注入路由
func (c *DepartmentController) injectRouting(getBean func(string) interface{}) {
	router, ok := getBean(web.BeanGinEngine).(gin.IRouter)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", web.BeanGinEngine)
		return
	}

	// API前缀
	apiPrefix := "/api/v1/organization"
	// 组织机构路由组
	organizationGroup := router.Group(apiPrefix)
	// 添加认证中间件
	organizationGroup.Use(middleware.JWTAuth())

	// 部门管理路由
	departments := organizationGroup.Group("/departments")
	{
		// 无参数路由
		// 创建部门
		departments.POST("", c.CreateDepartment)
		// 获取部门列表
		departments.GET("", c.ListDepartments)

		// 固定路径路由
		// 获取部门树
		departments.GET("/tree", c.GetDepartmentTree)
		// 获取用户所属部门
		departments.GET("/user", c.GetUserDepartments)

		// 用户相关路由组
		usersDept := departments.Group("/users")
		{
			// 为用户分配部门
			usersDept.POST("/:userId/departments/:departmentId", c.AssignDepartmentToUser)
			// 移除用户部门
			usersDept.DELETE("/:userId/departments/:departmentId", c.RemoveDepartmentFromUser)
		}

		// ID参数路由 - 使用detail前缀避免路由冲突
		// 更新部门
		departments.PUT("/detail/:id", c.UpdateDepartment)
		// 删除部门
		departments.DELETE("/detail/:id", c.DeleteDepartment)
		// 获取部门详情
		departments.GET("/detail/:id", c.GetDepartment)
		// 获取部门用户列表
		departments.GET("/detail/:id/users", c.ListDepartmentUsers)
	}
}
