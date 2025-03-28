package controller

import (
	"devops-platform/internal/pkg/common"
	"strconv"

	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/application/internal/domain"
	"devops-platform/internal/deploy-system/application/internal/service"
	"devops-platform/pkg/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AppController 应用管理控制器
type AppController struct {
	web.Controller
	AppService    *service.AppService
	DeployService *service.DeployService
	AppQuery      *service.AppQuery
}

// NewAppController 创建应用管理控制器
func NewAppController() *AppController {
	return &AppController{}
}

// Inject 实现依赖注入
func (c *AppController) Inject(getBean func(string) interface{}) {
	c.InjectService(getBean)
	c.injectRouting(getBean)
}

// InjectService 注入服务
func (c *AppController) InjectService(getBean func(string) interface{}) {
	appService, ok := getBean(domain.BeanAppService).(*service.AppService)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanAppService)
		return
	}
	c.AppService = appService

	deployService, ok := getBean(domain.BeanDeployService).(*service.DeployService)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanDeployService)
		return
	}
	c.DeployService = deployService

	appQuery, ok := getBean(domain.BeanAppQuery).(*service.AppQuery)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanAppQuery)
		return
	}
	c.AppQuery = appQuery
}

// CreateApplication 创建应用
// @Summary 创建应用
// @Description 创建新应用
// @Tags 应用管理
// @Accept json
// @Produce json
// @Param data body domain.CreateAppCommand true "应用信息"
// @Success 200 {object} common.Response{data=string}
// @Router /api/v1/apps [post]
func (c *AppController) CreateApplication(ctx *gin.Context) {
	var command domain.CreateAppCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 从上下文获取当前用户ID
	userID, exists := ctx.Get("userID")
	if exists {
		command.Creator = userID.(types.Long)
	}

	id, err := c.AppService.CreateApplication(ctx, &command)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, id)
}

// UpdateApplication 更新应用
// @Summary 更新应用
// @Description 更新应用信息
// @Tags 应用管理
// @Accept json
// @Produce json
// @Param id path int true "应用ID"
// @Param data body domain.UpdateAppCommand true "应用信息"
// @Success 200 {object} common.Response
// @Router /api/v1/apps/{id} [put]
func (c *AppController) UpdateApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的应用ID")
		return
	}

	var command domain.UpdateAppCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}
	command.ID = types.Long(id)

	if err := c.AppService.UpdateApplication(ctx, &command); err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// GetApplication 获取应用
// @Summary 获取应用详情
// @Description 根据ID获取应用详情
// @Tags 应用管理
// @Produce json
// @Param id path int true "应用ID"
// @Success 200 {object} common.Response
// @Router /api/v1/apps/{id} [get]
func (c *AppController) GetApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的应用ID")
		return
	}

	app, err := c.AppQuery.GetApplicationByID(ctx, types.Long(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	// 获取应用分组
	groups, _ := c.AppQuery.GetAppGroups(ctx, app.ID)
	app.Groups = groups

	common.ResponseSuccess(ctx, app)
}

// ListApplications 查询应用列表
// @Summary 查询应用列表
// @Description 查询应用列表
// @Tags 应用管理
// @Produce json
// @Param name query string false "应用名称"
// @Param status query string false "应用状态"
// @Param group_id query int false "分组ID"
// @Param page query int false "页码" default(1)
// @Param size query int false "每页大小" default(10)
// @Success 200 {object} common.Response
// @Router /api/v1/apps [get]
func (c *AppController) ListApplications(ctx *gin.Context) {
	var query domain.AppQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 设置默认值
	if query.Page <= 0 {
		query.Page = 1
	}
	if query.Size <= 0 {
		query.Size = 10
	}

	apps, total, err := c.AppQuery.ListApplications(ctx, &query)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccessWithPageExt(ctx, apps, total, query.Page, query.Size)
}

// DeleteApplication 删除应用
// @Summary 删除应用
// @Description 删除应用
// @Tags 应用管理
// @Produce json
// @Param id path int true "应用ID"
// @Success 200 {object} common.Response
// @Router /api/v1/apps/{id} [delete]
func (c *AppController) DeleteApplication(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的应用ID")
		return
	}

	if err := c.AppService.DeleteApplication(ctx, types.Long(id)); err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// CreateAppGroup 创建应用分组
// @Summary 创建应用分组
// @Description 创建应用分组
// @Tags 应用分组
// @Accept json
// @Produce json
// @Param data body object true "分组信息"
// @Success 200 {object} common.Response
// @Router /api/v1/app-groups [post]
func (c *AppController) CreateAppGroup(ctx *gin.Context) {
	var params struct {
		Name        string `json:"name" binding:"required"`
		Description string `json:"description"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	id, err := c.AppService.CreateAppGroup(ctx, params.Name, params.Description)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{
		"id": id,
	})
}

// ListAppGroups 查询应用分组列表
// @Summary 查询应用分组列表
// @Description 查询应用分组列表
// @Tags 应用分组
// @Produce json
// @Success 200 {object} common.Response
// @Router /api/v1/app-groups [get]
func (c *AppController) ListAppGroups(ctx *gin.Context) {
	groups, err := c.AppService.ListAppGroups(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, groups)
}

// AddAppToGroup 添加应用到分组
// @Summary 添加应用到分组
// @Description 添加应用到分组
// @Tags 应用分组
// @Accept json
// @Produce json
// @Param data body object true "关联信息"
// @Success 200 {object} common.Response
// @Router /api/v1/app-groups/apps [post]
func (c *AppController) AddAppToGroup(ctx *gin.Context) {
	var params struct {
		AppID   types.Long `json:"app_id" binding:"required"`
		GroupID types.Long `json:"group_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	if err := c.AppService.AddAppToGroup(ctx, params.AppID, params.GroupID); err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// CreateEnvironment 创建应用环境
// @Summary 创建应用环境
// @Description 创建应用环境
// @Tags 应用环境
// @Accept json
// @Produce json
// @Param data body domain.CreateEnvCommand true "环境信息"
// @Success 200 {object} common.Response
// @Router /api/v1/envs [post]
func (c *AppController) CreateEnvironment(ctx *gin.Context) {
	var command domain.CreateEnvCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	id, err := c.AppService.CreateAppEnv(ctx, &command)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{
		"id": id,
	})
}

// ListEnvironments 查询环境列表
// @Summary 查询环境列表
// @Description 查询环境列表
// @Tags 应用环境
// @Produce json
// @Success 200 {object} common.Response
// @Router /api/v1/envs [get]
func (c *AppController) ListEnvironments(ctx *gin.Context) {
	envs, err := c.AppQuery.ListAppEnvs(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, envs)
}

// CreateReleasePlan 创建发布计划
// @Summary 创建发布计划
// @Description 创建发布计划
// @Tags 发布管理
// @Accept json
// @Produce json
// @Param data body domain.CreateReleaseCommand true "发布计划信息"
// @Success 200 {object} common.Response
// @Router /api/v1/releases [post]
func (c *AppController) CreateReleasePlan(ctx *gin.Context) {
	var command domain.CreateReleaseCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	id, err := c.DeployService.CreateReleasePlan(ctx, &command)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{
		"id": id,
	})
}

// ExecuteReleasePlan 执行发布计划
// @Summary 执行发布计划
// @Description 执行发布计划
// @Tags 发布管理
// @Produce json
// @Param id path int true "发布计划ID"
// @Success 200 {object} common.Response
// @Router /api/v1/releases/{id}/execute [post]
func (c *AppController) ExecuteReleasePlan(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的发布计划ID")
		return
	}

	deployID, err := c.DeployService.ExecuteReleasePlan(ctx, types.Long(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{
		"deploy_id": deployID,
	})
}

// GetDeployment 获取部署记录
// @Summary 获取部署记录
// @Description 获取部署记录详情
// @Tags 发布管理
// @Produce json
// @Param id path int true "部署ID"
// @Success 200 {object} common.Response
// @Router /api/v1/deployments/{id} [get]
func (c *AppController) GetDeployment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部署ID")
		return
	}

	deployment, err := c.AppQuery.GetDeploymentByID(ctx, types.Long(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, deployment)
}

// ListDeployments 查询部署历史
// @Summary 查询部署历史
// @Description 查询部署历史列表
// @Tags 发布管理
// @Produce json
// @Param app_id query int false "应用ID"
// @Param env_id query int false "环境ID"
// @Success 200 {object} common.Response
// @Router /api/v1/deployments [get]
func (c *AppController) ListDeployments(ctx *gin.Context) {
	appIDStr := ctx.Query("app_id")
	envIDStr := ctx.Query("env_id")

	var appID, envID types.Long

	if appIDStr != "" {
		id, err := strconv.ParseInt(appIDStr, 10, 64)
		if err != nil {
			common.ResponseBadRequest(ctx, "无效的应用ID")
			return
		}
		appID = types.Long(id)
	}

	if envIDStr != "" {
		id, err := strconv.ParseInt(envIDStr, 10, 64)
		if err != nil {
			common.ResponseBadRequest(ctx, "无效的环境ID")
			return
		}
		envID = types.Long(id)
	}

	deployments, err := c.AppQuery.ListDeployments(ctx, appID, envID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, deployments)
}

// RollbackDeployment 回滚部署
// @Summary 回滚部署
// @Description 回滚部署
// @Tags 发布管理
// @Produce json
// @Param id path int true "部署ID"
// @Success 200 {object} common.Response
// @Router /api/v1/deployments/{id}/rollback [post]
func (c *AppController) RollbackDeployment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部署ID")
		return
	}

	if err := c.DeployService.RollbackDeployment(ctx, types.Long(id)); err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// ConfigureHPA 配置应用HPA
// @Summary 配置应用HPA
// @Description 配置应用HPA自动伸缩
// @Tags 应用管理
// @Accept json
// @Produce json
// @Param id path int true "应用ID"
// @Param data body object true "HPA配置"
// @Success 200 {object} common.Response
// @Router /api/v1/apps/{id}/hpa [post]
func (c *AppController) ConfigureHPA(ctx *gin.Context) {
	idStr := ctx.Param("id")
	appID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的应用ID")
		return
	}

	var params struct {
		MinReplicas  int `json:"min_replicas" binding:"required,min=1"`
		MaxReplicas  int `json:"max_replicas" binding:"required,min=1"`
		TargetCPU    int `json:"target_cpu" binding:"required,min=1,max=100"`
		TargetMemory int `json:"target_memory"`
	}
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 最大副本数必须大于等于最小副本数
	if params.MaxReplicas < params.MinReplicas {
		common.ResponseBadRequest(ctx, "最大副本数必须大于等于最小副本数")
		return
	}

	id, err := c.DeployService.CreateHPA(ctx, types.Long(appID), params.MinReplicas, params.MaxReplicas, params.TargetCPU, params.TargetMemory)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{
		"id": id,
	})
}

// GetAppHPA 获取应用HPA配置
// @Summary 获取应用HPA配置
// @Description 获取应用HPA自动伸缩配置
// @Tags 应用管理
// @Produce json
// @Param id path int true "应用ID"
// @Success 200 {object} common.Response
// @Router /api/v1/apps/{id}/hpa [get]
func (c *AppController) GetAppHPA(ctx *gin.Context) {
	idStr := ctx.Param("id")
	appID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的应用ID")
		return
	}

	hpa, err := c.AppQuery.GetAppHPA(ctx, types.Long(appID))
	if err != nil {
		// HPA配置不存在时，返回空配置而不是错误
		if err.Error() == "record not found" {
			common.ResponseSuccess(ctx, nil)
			return
		}
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, hpa)
}
