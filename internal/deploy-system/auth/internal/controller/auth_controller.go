package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/auth/internal/domain"
	"devops-platform/internal/deploy-system/auth/internal/service"
	"devops-platform/internal/pkg/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController 认证控制器
type AuthController struct {
	web.Controller
	Service *service.AuthService `inject:"authService"`
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

// Login 用户登录
// @Summary 用户登录
// @Description 使用用户名和密码登录
// @Tags 认证
// @Accept json
// @Produce json
// @Param data body domain.LoginRequest true "登录信息"
// @Success 200 {object} common.Response{data=web.TokenResponse} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 401 {object} common.ErrorResponse "认证失败"
// @Router /auth/login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var req domain.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "请求参数错误: "+err.Error())
		return
	}

	tokenInfo, err := c.Service.Login(ctx, req.Username, req.Password)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	// 存储用户信息到上下文
	sessionUser := &domain.SessionUser{
		ID:       tokenInfo.UserID,
		Username: tokenInfo.Username,
		Token:    tokenInfo.Token,
	}
	c.SetCurrentUser(ctx, sessionUser.ToUserContext())

	// 返回认证令牌
	c.ReturnTokenSuccess(ctx, tokenInfo.Token)
}

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出系统
// @Tags 认证
// @Accept json
// @Produce json
// @Success 200 {object} common.Response "成功"
// @Failure 401 {object} common.ErrorResponse "未认证"
// @Router /auth/logout [post]
func (c *AuthController) Logout(ctx *gin.Context) {
	user := c.CurrentUser(ctx)
	if user == nil {
		common.ResponseUnauthorized(ctx, "用户未登录")
		return
	}

	err := c.Service.Logout(ctx, user.UserID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	c.ReturnSuccess(ctx)
}

// GetUserInfo 获取当前用户信息
// @Summary 获取当前用户信息
// @Description 获取当前用户的详细信息
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=domain.UserInfo} "成功"
// @Failure 401 {object} common.ErrorResponse "未认证"
// @Router /auth/me [get]
func (c *AuthController) GetUserInfo(ctx *gin.Context) {
	user := c.CurrentUser(ctx)
	if user == nil {
		common.ResponseUnauthorized(ctx, "用户未登录")
		return
	}

	userInfo, err := c.Service.GetUserInfo(ctx, user.UserID)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, userInfo)
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 修改当前用户的密码
// @Tags 用户
// @Accept json
// @Produce json
// @Param data body domain.ChangePasswordCommand true "密码信息"
// @Success 200 {object} common.Response "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Failure 401 {object} common.ErrorResponse "未认证"
// @Router /auth/change-password [post]
func (c *AuthController) ChangePassword(ctx *gin.Context) {
	user := c.CurrentUser(ctx)
	if user == nil {
		common.ResponseUnauthorized(ctx, "用户未登录")
		return
	}

	var req domain.ChangePasswordCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "请求参数错误: "+err.Error())
		return
	}

	err := req.Validate()
	if err != nil {
		common.ResponseBadRequest(ctx, err.Error())
		return
	}

	err = c.Service.ChangePassword(ctx, user.UserID, req.OldPassword, req.NewPassword)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	common.ResponseSuccess(ctx, nil)
}

// Register 用户注册
// @Summary 用户注册
// @Description 注册新用户
// @Tags 认证
// @Accept json
// @Produce json
// @Param data body domain.CreateUserCommand true "用户信息"
// @Success 201 {object} common.Response{data=string} "成功"
// @Failure 400 {object} common.ErrorResponse "请求参数错误"
// @Router /auth/register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	var req domain.CreateUserCommand
	if err := ctx.ShouldBindJSON(&req); err != nil {
		common.ResponseBadRequest(ctx, "请求参数错误: "+err.Error())
		return
	}

	err := req.Validate()
	if err != nil {
		common.ResponseBadRequest(ctx, err.Error())
		return
	}

	userID, err := c.Service.RegisterUser(ctx, &req)
	if err != nil {
		common.ResponseError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, common.Response{
		Code:    http.StatusCreated,
		Data:    userID,
		Message: "注册成功",
	})
}
