package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/organization/internal/domain"
	"devops-platform/internal/deploy-system/organization/internal/service"
	"devops-platform/internal/pkg/common"
	"devops-platform/internal/pkg/enum"
	"devops-platform/pkg/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DepartmentController 部门控制器
type DepartmentController struct {
	web.Controller
	Service *service.DepartmentService `inject:"DepartmentService"`
}

// NewDepartmentController 创建部门控制器实例
func NewDepartmentController() *DepartmentController {
	return &DepartmentController{}
}

// CreateDepartment 创建部门
// @Summary 创建部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param data body domain.CreateDepartmentCommand true "部门创建参数"
// @Success 200 {object} common.Response{data=types.Long} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments [post]
func (c *DepartmentController) CreateDepartment(ctx *gin.Context) {
	var command domain.CreateDepartmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		common.ResponseBadRequest(ctx, err.Error())
		return
	}

	id, err := c.Service.CreateDepartment(ctx, &command)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, id)
}

// UpdateDepartment 更新部门
// @Summary 更新部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门ID"
// @Param data body domain.UpdateDepartmentCommand true "部门更新参数"
// @Success 200 {object} common.Response{data=string} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 404 {object} common.ErrorResponse "部门不存在"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/detail/{id} [put]
func (c *DepartmentController) UpdateDepartment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部门ID")
		return
	}

	var command domain.UpdateDepartmentCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		common.ResponseBadRequest(ctx, err.Error())
		return
	}

	command.ID = types.Long(id)

	err = c.Service.UpdateDepartment(ctx, &command)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, "更新成功")
}

// DeleteDepartment 删除部门
// @Summary 删除部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门ID"
// @Success 200 {object} common.Response{data=string} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 404 {object} common.ErrorResponse "部门不存在"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/detail/{id} [delete]
func (c *DepartmentController) DeleteDepartment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部门ID")
		return
	}

	err = c.Service.DeleteDepartment(ctx, types.Long(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, "删除成功")
}

// GetDepartment 获取部门详情
// @Summary 获取部门详情
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门ID"
// @Success 200 {object} common.Response{data=domain.DepartmentVO} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 404 {object} common.ErrorResponse "部门不存在"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/detail/{id} [get]
func (c *DepartmentController) GetDepartment(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部门ID")
		return
	}

	department, err := c.Service.GetDepartmentByID(ctx, types.Long(id))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	if department == nil {
		common.ResponseNotFound(ctx, "部门不存在")
		return
	}

	common.ResponseSuccess(ctx, department)
}

// ListDepartments 获取部门列表
// @Summary 获取部门列表
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param pageNum query int false "页码，默认1"
// @Param pageSize query int false "每页数量，默认20"
// @Param name query string false "部门名称"
// @Param code query string false "部门编码"
// @Param parentId query int false "父部门ID"
// @Param status query int false "状态：0-禁用，1-启用"
// @Success 200 {object} common.Response{data=common.PageResult{list=[]domain.DepartmentVO}} "成功"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments [get]
func (c *DepartmentController) ListDepartments(ctx *gin.Context) {
	// 构建查询条件
	query := &domain.DepartmentQuery{}
	query.Page = types.GetIntParam(ctx, "pageNum", 1)
	query.Size = types.GetIntParam(ctx, "pageSize", 20)
	query.Name = ctx.Query("name")
	query.Code = ctx.Query("code")
	query.ParentID = types.Long(types.GetIntParam(ctx, "parentId", 0))

	// 修复状态类型转换
	statusVal := types.GetIntParam(ctx, "status", 0)
	query.Status = enum.Status(statusVal)

	departments, total, err := c.Service.ListDepartments(ctx, query)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	// 使用ResponseSuccessWithPageExt函数
	common.ResponseSuccessWithPageExt(ctx, departments, total, query.Page, query.Size)
}

// GetDepartmentTree 获取部门树结构
// @Summary 获取部门树结构
// @Tags 部门管理
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=[]domain.DepartmentVO} "成功"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/tree [get]
func (c *DepartmentController) GetDepartmentTree(ctx *gin.Context) {
	departments, err := c.Service.GetDepartmentTree(ctx)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, departments)
}

// GetUserDepartments 获取用户所属部门
// @Summary 获取用户所属部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param userId query int false "用户ID，默认当前用户"
// @Success 200 {object} common.Response{data=[]domain.DepartmentVO} "成功"
// @Failure 401 {object} common.ErrorResponse "未认证"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/user [get]
func (c *DepartmentController) GetUserDepartments(ctx *gin.Context) {
	// 获取用户ID
	userID := types.Long(types.GetIntParam(ctx, "userId", 0))
	if userID == 0 {
		// 如果未指定用户ID，则获取当前登录用户
		user := c.CurrentUser(ctx)
		if user == nil {
			common.ResponseUnauthorized(ctx, "用户未登录")
			return
		}
		userID = user.UserID
	}

	departments, err := c.Service.GetUserDepartments(ctx, userID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, departments)
}

// AssignDepartmentToUser 为用户分配部门
// @Summary 为用户分配部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param userId path int true "用户ID"
// @Param departmentId path int true "部门ID"
// @Param isLeader query bool false "是否为部门负责人，默认false"
// @Success 200 {object} common.Response{data=string} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/users/{userId}/departments/{departmentId} [post]
func (c *DepartmentController) AssignDepartmentToUser(ctx *gin.Context) {
	// 获取用户ID
	userIDStr := ctx.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的用户ID")
		return
	}

	// 获取部门ID
	departmentIDStr := ctx.Param("departmentId")
	departmentID, err := strconv.ParseInt(departmentIDStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部门ID")
		return
	}

	// 获取是否为部门负责人
	isLeaderStr := ctx.DefaultQuery("isLeader", "false")
	isLeader := isLeaderStr == "true"

	err = c.Service.AssignDepartmentToUser(ctx, types.Long(userID), types.Long(departmentID), isLeader)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, "分配成功")
}

// RemoveDepartmentFromUser 移除用户部门
// @Summary 移除用户部门
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param userId path int true "用户ID"
// @Param departmentId path int true "部门ID"
// @Success 200 {object} common.Response{data=string} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/users/{userId}/departments/{departmentId} [delete]
func (c *DepartmentController) RemoveDepartmentFromUser(ctx *gin.Context) {
	// 获取用户ID
	userIDStr := ctx.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的用户ID")
		return
	}

	// 获取部门ID
	departmentIDStr := ctx.Param("departmentId")
	departmentID, err := strconv.ParseInt(departmentIDStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部门ID")
		return
	}

	err = c.Service.RemoveDepartmentFromUser(ctx, types.Long(userID), types.Long(departmentID))
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, "移除成功")
}

// ListDepartmentUsers 获取部门用户列表
// @Summary 获取部门用户列表
// @Tags 部门管理
// @Accept json
// @Produce json
// @Param id path int true "部门ID"
// @Param pageNum query int false "页码，默认1"
// @Param pageSize query int false "每页数量，默认20"
// @Param username query string false "用户名"
// @Param realName query string false "真实姓名"
// @Success 200 {object} common.Response{data=common.PageResult{list=[]domain.UserVO}} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 500 {object} common.ErrorResponse "服务器内部错误"
// @Router /departments/detail/{id}/users [get]
func (c *DepartmentController) ListDepartmentUsers(ctx *gin.Context) {
	// 获取部门ID
	idStr := ctx.Param("id")
	departmentID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		common.ResponseBadRequest(ctx, "无效的部门ID")
		return
	}

	// 构建查询条件
	query := &domain.UserQuery{}
	query.Page = types.GetIntParam(ctx, "pageNum", 1)
	query.Size = types.GetIntParam(ctx, "pageSize", 20)
	query.Username = ctx.Query("username")
	query.RealName = ctx.Query("realName")

	users, total, err := c.Service.ListDepartmentUsers(ctx, types.Long(departmentID), query)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	// 使用ResponseSuccessWithPageExt函数
	common.ResponseSuccessWithPageExt(ctx, users, total, query.Page, query.Size)
}
