package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"devops-platform/internal/common/web/internal/domain"
	"devops-platform/internal/common/web/internal/service"
	"devops-platform/internal/pkg/security"
	"devops-platform/pkg/types"
)

type Controller struct {
}

func (c *Controller) GetLongParam(ctx *gin.Context, param string) (result types.Long, err error) {
	result, err = types.StringToLong(ctx.Param(param))
	return
}

func (c *Controller) GetContext(ctx *gin.Context) (realContext context.Context) {
	return service.GetContext(ctx)
}

func (c *Controller) SetContext(ctx *gin.Context, realContext context.Context) {
	service.SetContext(ctx, realContext)
}

func (c *Controller) SetCurrentUser(ctx *gin.Context, user security.User) {
	realContext := c.GetContext(ctx)
	realContext = security.SetCurrentUser(realContext, user)
	c.SetContext(ctx, realContext)
}

func (c *Controller) CurrentUser(ctx *gin.Context) security.User {
	realCtx := c.GetContext(ctx)
	return security.CurrentUser(realCtx)
}

func (c *Controller) AbortErr(ctx *gin.Context, err error) {
	service.AbortErr(ctx, err)
}

func (c *Controller) ReturnErr(ctx *gin.Context, err error) {
	if err == nil {
		return
	}
	ctx.Set(domain.ErrKeyInContext, err)
	if errs, ok := err.(domain.Error); ok {
		ctx.JSON(errs.GetStatus(), gin.H{
			"msg":  err.Error(),
			"code": errs.GetCode(),
		})
	} else {
		ctx.JSON(400, gin.H{
			"msg":  err.Error(),
			"code": 400,
		})
	}
}

func (c *Controller) ReturnSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func (c *Controller) ReturnCreateSuccess(ctx *gin.Context, id types.Long) {
	ctx.JSON(http.StatusCreated, gin.H{"id": id.String(), "msg": "create success"})
}

func (c *Controller) ReturnModifySuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "modify success"})
}

func (c *Controller) ReturnDeleteSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "delete success"})
}

func (c *Controller) ReturnQuerySuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func (c *Controller) ReturnFindSuccess(ctx *gin.Context, records interface{}, total int64) {
	ctx.JSON(http.StatusOK, FindResponse{
		Records: records,
		Total:   total,
	})
}

type FindResponse struct {
	Records interface{}
	Total   int64
}
