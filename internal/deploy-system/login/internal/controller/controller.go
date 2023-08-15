package controller

import (
	"context"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/login/internal/service"
	"devops-platform/pkg/common"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const LoginError int = 50000

type Controller struct {
	web.Controller
	SsoLoginService *service.KeyCloakService `inject:"SsoLoginService"`
}

func (c *Controller) LocalLogin(ctx *gin.Context) {
	var req domain.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}
	//验证登录
	user, err := c.SsoLoginService.LocalLogin(ctx, &req)
	if err != nil {
		c.ReturnErr(ctx, common.WarpError(err))
		return
	}

	c.SetCurrentUser(ctx, user)

	token, err := common.GenerateToken(user)
	if err != nil {
		c.ReturnErr(ctx, common.Unauthorized(LoginError, errors.New("获取token失败")))
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (c *Controller) Authentication(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	//if len(token) == 0 {
	//	token = ctx.Query("token")
	//}
	//
	///*
	//* 没有token或token的开头不是"Bearer " 则直接进行下一步
	// */
	//token = strings.ToLower(token)
	//if len(token) == 0 || !strings.HasPrefix(token, domain.TokenPrefix) {
	//	ctx.Next()
	//	return
	//}
	//token = strings.TrimPrefix(token, domain.TokenPrefix)

	claims, err := common.ParseToken(token)
	if err != nil {
		fmt.Print(err)
		return
	}
	user := c.parserUserFromClaims(ctx, claims)

	c.SetCurrentUser(ctx, user)
	ctx.Next()
}

func (c *Controller) parserUserFromClaims(ctx context.Context, claims *common.Claims) (loginUser *domain.LoginUserVO) {
	// 1. 从claims中获取用户ID
	userId := claims.UserID
	// 2. 使用用户ID从数据库查询用户信息
	user, err := c.SsoLoginService.UserService.GetByID(ctx, userId)
	if err != nil {
		return nil
	}
	loginUser = &domain.LoginUserVO{
		UserID:    user.ID,
		LoginName: user.Name,
		Username:  user.Username,
		Role:      user.Role,
	}
	return loginUser
}
