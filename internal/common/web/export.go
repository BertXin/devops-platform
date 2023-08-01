package web

import (
	"devops-platform/internal/common/web/internal/domain"
	"devops-platform/internal/common/web/internal/service"
	"devops-platform/internal/pkg/security"
	"github.com/gin-gonic/gin"
)

//注册bean
const (
	BeanGinEngine            = domain.BeanGinEngine
	ModeUnitTesting          = domain.ModeUnitTesting
	ModeKey                  = domain.ModeKey
	BeanAuthenticationVerify = domain.BeanAuthenticationVerify
)

func SetCurrentUser(ctx *gin.Context, user security.User) {
	service.SetCurrentUser(ctx, user)
}

func Authenticated(ctx *gin.Context) bool {
	return service.Authenticated(ctx)
}

func AddIgnoreUrls(urls ...string) {
	service.AddIgnoreUrls(urls...)
}
