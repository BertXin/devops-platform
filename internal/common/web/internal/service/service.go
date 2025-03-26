package service

import (
	"context"
	"devops-platform/internal/common/web/internal/domain"
	"devops-platform/internal/pkg/security"

	"github.com/gin-gonic/gin"
)

func GetContext(ctx *gin.Context) (realContext context.Context) {

	if realCtx, ok := ctx.Get(domain.RealContext); ok {
		if realContext, ok = realCtx.(context.Context); ok {
			return
		}
	}
	realContext = context.TODO()

	return
}

// SetContext 设置实际的上下文到gin.Context中
func SetContext(ctx *gin.Context, realContext context.Context) {
	if realContext != nil {
		ctx.Set(domain.RealContext, realContext)
	}
}

func SetCurrentUser(ctx *gin.Context, user *security.UserContext) {
	if user == nil {
		return
	}
	realContext := GetContext(ctx)
	realContext = security.SetUserContext(realContext, user)
	SetContext(ctx, realContext)
}

func Authenticated(ctx *gin.Context) bool {
	return CurrentUser(ctx) != nil
}

func CurrentUser(ctx *gin.Context) (user *security.UserContext) {
	realContext := GetContext(ctx)
	return security.GetUserContext(realContext)
}

func AbortErr(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	ctx.Set(domain.ErrKeyInContext, err)
	if errs, ok := err.(domain.Error); ok {

		ctx.AbortWithStatusJSON(errs.GetStatus(), gin.H{
			"msg":  err.Error(),
			"code": errs.GetCode(),
		})
	} else {
		ctx.AbortWithStatusJSON(400, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
	}
}
