package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/login/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	web.Controller
	SsoLoginService *service.KeyCloakService `inject:"SsoLoginService"`
}

func (c *Controller) Authentication(ctx *gin.Context) {
	ctx.Next()
}
