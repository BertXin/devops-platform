package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/department/internal/domain"
	"devops-platform/internal/deploy-system/department/internal/repository"
	"devops-platform/internal/deploy-system/department/internal/service"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	web.Controller
	DeptQuery *repository.Repository
	Service   *service.Service
}

func (c *Controller) injectQuery(getBean func(string) interface{}) {

	deptQuery, ok := getBean(domain.BeanRepository).(*repository.Repository)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanRepository)
		return
	}
	c.DeptQuery = deptQuery
}
func (c *Controller) injectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanService).(*service.Service)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanService)
		return
	}
	c.Service = service
}

func (c *Controller) GetDeptByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}
	dept, err := c.Service.FindDeptByID(ctx, id)
	if err != nil {
		c.ReturnErr(ctx, common.RequestNotFoundError("部门信息不存在"))
		return
	}

	c.ReturnQuerySuccess(ctx, dept.VO())
}
func (c *Controller) CreateDept(ctx *gin.Context) {
	var command domain.CreateDeptCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[创建部门]解析失败", err))
		return
	}
	if err := command.Validate(); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[创建部门]校验失败", err))
		return
	}
	//需要去数据库查询是否有这个部门
	_, total, err := c.Service.FindDeptByName(ctx, command.Name, command.ParentID, types.Pagination{})
	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	if total > 0 {
		c.ReturnErr(ctx, common.RequestParamNilError("部门名称重复"))
		return
	}
	id, err := c.Service.Create(ctx, &command)

	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnCreateSuccess(ctx, id)

}

type FindByNameAndMobileQuery struct {
	Name     string     `json:"name" form:"name"`
	ParentID types.Long `json:"parent_id" form:"parent_id"`
	types.Pagination
}
type FindByNameAndMobileResponse struct {
	Records []domain.Dept `json:"records" ` //记录
	Total   int64         `json:"total" `   //总数
}

func (c *Controller) FindDeptByName(ctx *gin.Context) {

	var query FindByNameAndMobileQuery
	if err := ctx.ShouldBindQuery(&query); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[查询部门]解析失败", err))
		return
	}

	depts, total, err := c.Service.FindDeptByName(ctx, query.Name, query.ParentID, query.Pagination)
	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	if depts == nil {
		c.ReturnErr(ctx, common.RequestNotFoundError("部门信息不存在"))
		return
	}
	c.ReturnFindSuccess(ctx, FindByNameAndMobileResponse{
		Records: depts,
		Total:   total,
	}, total)
}

// 根据部门ID修改部门名称
func (c *Controller) ModifyDeptNameByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}
	var command domain.ModifyDeptCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[修改部门]解析失败", err))
		return
	}
	if err := command.Validate(); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[修改部门]校验失败", err))
		return
	}
	command.ID = id
	if err := c.Service.ModifyDeptNameByID(ctx, &command); err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnModifySuccess(ctx)
}

// 根据部门ID删除部门
func (c *Controller) DeleteDeptByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}
	if err := c.Service.DeleteDeptByID(ctx, id); err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnDeleteSuccess(ctx)
}
