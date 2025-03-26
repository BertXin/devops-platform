package controller

import (
	"devops-platform/internal/deploy-system/authorization"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// PermissionController 权限控制器
type PermissionController struct {
	permissionService authorization.PermissionService
	logger            *logrus.Logger
}

// CreatePermission 创建权限
// @Summary 创建权限
// @Description 创建新权限
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param data body domain.CreatePermissionCommand true "权限信息"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/permissions [post]
func (c *PermissionController) CreatePermission(ctx *gin.Context) {
	var req domain.CreatePermissionCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 创建权限
	id, err := c.permissionService.CreatePermission(ctx, &req)
	if err != nil {
		common.ResponseInternalError(ctx, "创建权限失败", err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{"id": id})
}

// UpdatePermission 更新权限
// @Summary 更新权限
// @Description 更新权限信息
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param id path int true "权限ID"
// @Param data body domain.UpdatePermissionCommand true "权限信息"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/permissions/detail/{id} [put]
func (c *PermissionController) UpdatePermission(ctx *gin.Context) {
	// 获取权限ID
	permissionID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 权限ID必须是数字")
		return
	}

	var req domain.UpdatePermissionCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 设置权限ID
	req.ID = types.Long(permissionID)

	// 更新权限
	err = c.permissionService.UpdatePermission(ctx, &req)
	if err != nil {
		common.ResponseInternalError(ctx, "更新权限失败", err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// DeletePermission 删除权限
// @Summary 删除权限
// @Description 删除指定权限
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param id path int true "权限ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/permissions/detail/{id} [delete]
func (c *PermissionController) DeletePermission(ctx *gin.Context) {
	// 获取权限ID
	permissionID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 权限ID必须是数字")
		return
	}

	// 删除权限
	err = c.permissionService.DeletePermission(ctx, types.Long(permissionID))
	if err != nil {
		common.ResponseInternalError(ctx, "删除权限失败", err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// GetPermissionByID 获取权限详情
// @Summary 获取权限详情
// @Description 获取指定权限的详细信息
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param id path int true "权限ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/permissions/detail/{id} [get]
func (c *PermissionController) GetPermissionByID(ctx *gin.Context) {
	// 获取权限ID
	permissionID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 权限ID必须是数字")
		return
	}

	// 获取权限
	permission, err := c.permissionService.GetPermissionByID(ctx, types.Long(permissionID))
	if err != nil {
		common.ResponseInternalError(ctx, "获取权限失败", err)
		return
	}

	if permission == nil {
		common.ResponseNotFound(ctx, "未找到指定权限")
		return
	}

	common.ResponseSuccess(ctx, permission)
}

// ListPermissions 获取权限列表
// @Summary 获取权限列表
// @Description 获取权限列表，支持分页和条件查询
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param name query string false "权限名称"
// @Param type query string false "权限类型"
// @Param status query int false "状态"
// @Param parent_id query int false "父权限ID"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/permissions [get]
func (c *PermissionController) ListPermissions(ctx *gin.Context) {
	var query domain.PermissionQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 获取权限列表
	permissions, total, err := c.permissionService.ListPermissions(ctx, &query)
	if err != nil {
		common.ResponseInternalError(ctx, "获取权限列表失败", err)
		return
	}

	// 使用扩展的分页响应
	common.ResponseSuccessWithPageExt(ctx, permissions, total, query.Page, query.Size)
}

// GetPermissionTree 获取权限树结构
// @Summary 获取权限树结构
// @Description 获取权限树结构
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/permissions/tree [get]
func (c *PermissionController) GetPermissionTree(ctx *gin.Context) {
	// 获取权限树
	tree, err := c.permissionService.GetPermissionTree(ctx)
	if err != nil {
		common.ResponseInternalError(ctx, "获取权限树失败", err)
		return
	}

	common.ResponseSuccess(ctx, tree)
}

// InjectService 注入服务
func (c *PermissionController) InjectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanPermissionService).(authorization.PermissionService)
	if !ok {
		logrus.Errorf("初始化时获取[%s]失败", domain.BeanPermissionService)
		return
	}
	c.permissionService = service
}

// NewPermissionController 创建权限控制器实例
func NewPermissionController() *PermissionController {
	return &PermissionController{}
}
