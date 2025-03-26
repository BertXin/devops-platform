package controller

import (
	"context"
	"devops-platform/internal/deploy-system/authorization/internal/domain"
	serviceAuth "devops-platform/internal/deploy-system/authorization/internal/service"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 定义本地接口，只包含本控制器需要的方法
type roleService interface {
	CreateRole(ctx context.Context, command *domain.CreateRoleCommand) (types.Long, error)
	UpdateRole(ctx context.Context, command *domain.UpdateRoleCommand) error
	DeleteRole(ctx context.Context, id types.Long) error
	GetRoleByID(ctx context.Context, id types.Long) (*domain.RoleVO, error)
	ListRoles(ctx context.Context, query *domain.RoleQuery) ([]*domain.RoleVO, int64, error)
	AssignPermissionsToRole(ctx context.Context, roleID types.Long, permissionIDs []types.Long) error
	GetRolePermissions(ctx context.Context, roleID types.Long) ([]*domain.PermissionVO, error)
}

// RoleController 角色控制器
type RoleController struct {
	roleService roleService // 使用本地定义的接口
	logger      *logrus.Logger
}

// CreateRole 创建角色
// @Summary 创建角色
// @Description 创建新角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param data body domain.CreateRoleCommand true "角色信息"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles [post]
func (c *RoleController) CreateRole(ctx *gin.Context) {
	var req domain.CreateRoleCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 创建角色
	id, err := c.roleService.CreateRole(ctx, &req)
	if err != nil {
		common.ResponseInternalError(ctx, "创建角色失败", err)
		return
	}

	common.ResponseSuccess(ctx, gin.H{"id": id})
}

// UpdateRole 更新角色
// @Summary 更新角色
// @Description 更新角色信息
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param id path int true "角色ID"
// @Param data body domain.UpdateRoleCommand true "角色信息"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles/:id [put]
func (c *RoleController) UpdateRole(ctx *gin.Context) {
	// 获取角色ID
	roleID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 角色ID必须是数字")
		return
	}

	var req domain.UpdateRoleCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 设置角色ID
	req.ID = types.Long(roleID)

	// 验证参数
	if err := req.Validate(); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 更新角色
	err = c.roleService.UpdateRole(ctx, &req)
	if err != nil {
		common.ResponseInternalError(ctx, "更新角色失败", err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Description 删除指定角色
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param id path int true "角色ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles/:id [delete]
func (c *RoleController) DeleteRole(ctx *gin.Context) {
	// 获取角色ID
	roleID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 角色ID必须是数字")
		return
	}

	// 删除角色
	err = c.roleService.DeleteRole(ctx, types.Long(roleID))
	if err != nil {
		common.ResponseInternalError(ctx, "删除角色失败", err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// GetRoleByID 获取角色详情
// @Summary 获取角色详情
// @Description 获取指定角色的详细信息
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param id path int true "角色ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles/:id [get]
func (c *RoleController) GetRoleByID(ctx *gin.Context) {
	// 获取角色ID
	roleID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 角色ID必须是数字")
		return
	}

	// 获取角色
	role, err := c.roleService.GetRoleByID(ctx, types.Long(roleID))
	if err != nil {
		common.ResponseInternalError(ctx, "获取角色失败", err)
		return
	}

	if role == nil {
		common.ResponseNotFound(ctx, "未找到指定角色")
		return
	}

	common.ResponseSuccess(ctx, role)
}

// ListRoles 获取角色列表
// @Summary 获取角色列表
// @Description 获取角色列表，支持分页和条件查询
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param name query string false "角色名称"
// @Param code query string false "角色编码"
// @Param status query int false "状态"
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles [get]
func (c *RoleController) ListRoles(ctx *gin.Context) {
	var query domain.RoleQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 获取角色列表
	roles, total, err := c.roleService.ListRoles(ctx, &query)
	if err != nil {
		common.ResponseInternalError(ctx, "获取角色列表失败", err)
		return
	}

	// 使用扩展的分页响应
	common.ResponseSuccessWithPageExt(ctx, roles, total, query.Page, query.Size)
}

// AssignPermissionsToRole 为角色分配权限
// @Summary 为角色分配权限
// @Description 为指定角色分配权限
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param id path int true "角色ID"
// @Param data body map[string][]int true "权限ID列表"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles/:id/permissions [post]
func (c *RoleController) AssignPermissionsToRole(ctx *gin.Context) {
	// 获取角色ID
	roleID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 角色ID必须是数字")
		return
	}

	// 获取权限ID列表
	var req struct {
		PermissionIDs []types.Long `json:"permission_ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "参数错误: "+err.Error())
		return
	}

	// 分配权限
	err = c.roleService.AssignPermissionsToRole(ctx, types.Long(roleID), req.PermissionIDs)
	if err != nil {
		common.ResponseInternalError(ctx, "分配权限失败", err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// GetRolePermissions 获取角色权限
// @Summary 获取角色权限
// @Description 获取指定角色的权限列表
// @Tags 角色管理
// @Accept  json
// @Produce  json
// @Param id path int true "角色ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/roles/:id/permissions [get]
func (c *RoleController) GetRolePermissions(ctx *gin.Context) {
	// 获取角色ID
	roleID, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "参数错误: 角色ID必须是数字")
		return
	}

	// 获取角色权限
	permissions, err := c.roleService.GetRolePermissions(ctx, types.Long(roleID))
	if err != nil {
		common.ResponseInternalError(ctx, "获取角色权限失败", err)
		return
	}

	common.ResponseSuccess(ctx, permissions)
}

// InjectService 注入服务
func (c *RoleController) InjectService(getBean func(string) interface{}) {
	// 获取日志对象
	c.logger = logrus.StandardLogger()

	// 先检查Bean是否存在
	beanObj := getBean(domain.BeanRoleService)
	if beanObj == nil {
		c.logger.Errorf("获取[%s]失败: Bean不存在", domain.BeanRoleService)
		return
	}

	// 尝试类型断言
	service, ok := beanObj.(roleService)
	if !ok {
		// 获取Bean的实际类型以便于调试
		c.logger.Errorf("获取[%s]失败: 类型断言错误，期望类型roleService，实际类型%T", domain.BeanRoleService, beanObj)

		// 尝试使用其他类型断言
		if _, ok := beanObj.(*serviceAuth.RoleService); ok {
			c.logger.Infof("Bean是*serviceAuth.RoleService类型，但不满足roleService接口")
		}
		return
	}

	c.roleService = service
	//c.logger.Infof("成功注入[%s]服务", domain.BeanRoleService)
}

// NewRoleController 创建角色控制器实例
func NewRoleController() *RoleController {
	return &RoleController{
		logger: logrus.StandardLogger(),
	}
}
