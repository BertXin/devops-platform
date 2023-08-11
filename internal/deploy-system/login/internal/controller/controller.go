package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/login/internal/service"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Controller struct {
	web.Controller
	SsoLoginService *service.KeyCloakService `inject:"SsoLoginService"`
}

func (c *Controller) Login(ctx *gin.Context) {
	var req domain.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		// 返回错误

	}

	resp, err := c.SsoLoginService.LocalLogin(ctx, &req)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, resp)
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

// ParseToken 解析JWT token
func ParseToken(token string) (*domain.TokenClaims, error) {

	// 1. 截取Bearer
	token = strings.TrimPrefix(token, "Bearer ")

	// 2. 解析token
	claim := &domain.TokenClaims{}
	_, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}

	return claim, nil

}
