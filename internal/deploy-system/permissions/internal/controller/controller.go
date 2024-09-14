package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/rbac/internal/domain"
	"devops-platform/internal/deploy-system/rbac/internal/repository"
	"devops-platform/internal/deploy-system/rbac/internal/service"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	web.Controller
	RbacQuery *repository.Repository
	Service   *service.Service
}

func (c *Controller) injectQuery(getBean func(string) interface{}) {

	rbacQuery, ok := getBean(domain.BeanRepository).(*repository.Repository)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanRepository)
		return
	}
	c.RbacQuery = rbacQuery
}

func (c *Controller) injectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanService).(*service.Service)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanService)
		return
	}
	c.Service = service
}

// CreateRole 创建角色
func (c *Controller) CreateRole(ctx *gin.Context) {
	var command domain.CreateRoleCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[创建角色]解析失败", err))
		return
	}
	if err := command.Validate(); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[创建角色]校验失败", err))
		return
	}
	//需要去数据库查询是否有这个角色
	_, total, err := c.Service.FindRoleByName(ctx, command.Name, types.Pagination{})
	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	if total > 0 {
		c.ReturnErr(ctx, common.RequestParamNilError("角色名称重复"))
		return
	}
	id, err := c.Service.CreateRole(ctx, &command)
	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnCreateSuccess(ctx, id)
}

func (c *Controller) DeleteRoleByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[角色ID]解析失败", err))
		return
	}
	if err := c.Service.DeleteRoleByID(ctx, id); err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnDeleteSuccess(ctx)
}

func (c *Controller) FindRoleByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[角色ID]解析失败", err))
		return
	}
	role, err := c.Service.FindRoleByID(ctx, id)
	if err != nil {
		c.ReturnErr(ctx, common.RequestNotFoundError("角色信息不存在"))
		return
	}
	c.ReturnQuerySuccess(ctx, role.VO())
}

type FindByNameAndMobileQuery struct {
	Name string `json:"name" form:"name"`
	types.Pagination
}

type FindByNameAndMobileResponse struct {
	Records []domain.Role `json:"records" ` //记录
	Total   int64         `json:"total" `   //总数
}

func (c *Controller) FindRoleByName(ctx *gin.Context) {
	var query FindByNameAndMobileQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[查询角色]解析失败", err))
		return
	}
	roles, total, err := c.Service.FindRoleByName(ctx, query.Name, query.Pagination)
	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	if roles == nil {
		c.ReturnErr(ctx, common.RequestNotFoundError("角色信息不存在"))
		return
	}
	c.ReturnFindSuccess(ctx, FindByNameAndMobileResponse{
		Records: roles,
		Total:   total,
	}, total)
}

func (c *Controller) ModifyRoleByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[角色ID]解析失败", err))
		return
	}
	var command domain.ModifyRoleCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[修改角色]解析失败", err))
		return
	}
	if err := command.Validate(); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[修改角色]校验失败", err))
		return
	}
	command.ID = id
	if err := c.Service.ModifyRoleByID(ctx, &command); err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnModifySuccess(ctx)
}
