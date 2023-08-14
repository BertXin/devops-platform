package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/login/internal/service"
	"devops-platform/pkg/common"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const LoginError int = 50000

type Controller struct {
	web.Controller
	SsoLoginService *service.KeyCloakService `inject:"SsoLoginService"`
}

func (c *Controller) Login(ctx *gin.Context) {
	var req domain.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}
	//验证登录
	user, err := c.SsoLoginService.LocalLogin(ctx, &req)
	if err != nil {
		return
	}

	c.SetCurrentUser(ctx, user)
	token, err := common.GenerateToken(user)
	if err != nil {
		c.ReturnErr(ctx, common.Unauthorized(LoginError, errors.New("换取token失败")))
		return
	}
	ctx.JSON(http.StatusOK, token)
}

func (c *Controller) Authentication(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if len(token) == 0 {
		token = ctx.Query("token")
	}
	/*
	 * 没有token或token的开头不是"Bearer " 则直接进行下一步
	 */
	token = strings.ToLower(token)
	if len(token) == 0 || !strings.HasPrefix(token, domain.TokenPrefix) {
		ctx.Next()
		return
	}
	token = strings.TrimPrefix(token, domain.TokenPrefix)

	ctx.Next()
}

// parseTokenFromHeader 解析JWT token
func parseTokenFromHeader(ctx *gin.Context) string {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	// 按空格分割
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		return ""
	}

	return parts[1]

}
