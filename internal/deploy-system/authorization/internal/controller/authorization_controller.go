package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/authorization/internal/service"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AuthorizationController 权限控制器
type AuthorizationController struct {
	web.Controller
	Service *service.AuthorizationService `inject:"AuthorizationService"`
}

// NewAuthorizationController 创建权限控制器实例
func NewAuthorizationController() *AuthorizationController {
	return &AuthorizationController{}
}

// GetUserRoles 获取用户角色
// @Summary 获取用户角色
// @Description 获取指定用户的角色列表
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/users/:user_id/roles [get]
func (c *AuthorizationController) GetUserRoles(ctx *gin.Context) {
	// 获取用户ID
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "用户ID必须是数字")
		return
	}

	// 获取用户角色
	roles, err := c.Service.GetUserRoles(ctx, types.Long(userID))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, roles)
}

// GetUserPermissions 获取用户权限
// @Summary 获取用户权限
// @Description 获取指定用户的权限列表
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/users/:user_id/permissions [get]
func (c *AuthorizationController) GetUserPermissions(ctx *gin.Context) {
	// 获取用户ID
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "用户ID必须是数字")
		return
	}

	// 获取用户权限
	permissions, err := c.Service.GetUserPermissions(ctx, types.Long(userID))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, permissions)
}

// AssignRolesToUser 为用户分配角色
// @Summary 为用户分配角色
// @Description 为指定用户分配角色
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户ID"
// @Param data body map[string][]int true "角色ID列表"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/users/:user_id/roles [post]
func (c *AuthorizationController) AssignRolesToUser(ctx *gin.Context) {
	// 获取用户ID
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "用户ID必须是数字")
		return
	}

	// 获取角色ID列表
	var req struct {
		RoleIDs []types.Long `json:"role_ids" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, err.Error())
		return
	}

	// 分配角色
	err = c.Service.AssignRolesToUser(ctx, types.Long(userID), req.RoleIDs)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// RemoveRoleFromUser 移除用户角色
// @Summary 移除用户角色
// @Description 移除用户的指定角色
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param user_id path int true "用户ID"
// @Param role_id path int true "角色ID"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/users/:user_id/roles/:role_id [delete]
func (c *AuthorizationController) RemoveRoleFromUser(ctx *gin.Context) {
	// 获取用户ID
	userID, err := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "用户ID必须是数字")
		return
	}

	// 获取角色ID
	roleID, err := strconv.ParseInt(ctx.Param("role_id"), 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "角色ID必须是数字")
		return
	}

	// 移除角色
	err = c.Service.RemoveRoleFromUser(ctx, types.Long(userID), types.Long(roleID))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// GetUserMenus 获取用户菜单
// @Summary 获取用户菜单
// @Description 获取当前用户的菜单列表
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/menus [get]
func (c *AuthorizationController) GetUserMenus(ctx *gin.Context) {
	// 从上下文获取用户信息
	user := c.CurrentUser(ctx)
	if user == nil {
		common.ResponseUnauthorized(ctx, "用户未登录")
		return
	}

	// 获取用户菜单
	menus, err := c.Service.GetUserMenus(ctx, user.UserID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, menus)
}

// HasPermission 检查权限
// @Summary 检查权限
// @Description 检查当前用户是否拥有指定权限
// @Tags 权限管理
// @Accept  json
// @Produce  json
// @Param permission path string true "权限标识"
// @Success 200 {object} common.Response
// @Router /api/v1/authorization/check/:permission [get]
func (c *AuthorizationController) HasPermission(ctx *gin.Context) {
	user := c.CurrentUser(ctx)
	if user == nil {
		common.ResponseUnauthorized(ctx, "用户未登录")
		return
	}

	permission := ctx.Param("permission")
	if permission == "" {
		common.ResponseBadRequest(ctx, "权限标识不能为空")
		return
	}

	has, err := c.Service.HasPermission(ctx, user.UserID, permission)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, has)
}
