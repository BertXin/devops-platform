package controller

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/internal/deploy-system/user/internal/repository"
	"devops-platform/internal/deploy-system/user/internal/service"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Controller struct {
	web.Controller
	UserQuery *repository.Repository
	Service   *service.Service
}

func (c *Controller) injectQuery(getBean func(string) interface{}) {

	userQuery, ok := getBean(domain.BeanRepository).(*repository.Repository)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanRepository)
		return
	}
	c.UserQuery = userQuery
}

func (c *Controller) injectService(getBean func(string) interface{}) {
	service, ok := getBean(domain.BeanService).(*service.Service)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", domain.BeanService)
		return
	}
	c.Service = service
}

// @Summary 获取用户信息
// @Tags user
// @Accept json
// @Produce  json
// @Param id path int64 true "用户ID"
// @Success 200 {object} domain.User
// @Router /user/{id} [get]
func (c *Controller) GetByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}
	user, err := c.Service.GetByID(c.GetContext(ctx), id)

	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	if user == nil {
		c.ReturnErr(ctx, common.RequestNotFoundError("用户信息不存在"))
		return
	}
	//c.ReturnQuerySuccess(ctx,user)
	ctx.JSON(http.StatusOK, user)
}

// @Summary 创建用户
// @Tags user
// @Accept json
// @Produce json
// @Param object body domain.CreateUserCommand true "创建用户"
// @Success 201 string string "{"id":"1", "msg": "create success"}"
// @Router /user [post]
func (c *Controller) CreateUser(ctx *gin.Context) {

	var command domain.CreateUserCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}

	command.Validate()
	//密码加密
	command.Password = common.SetPassword(command.Password)

	id, err := c.Service.Create(ctx, &command)

	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	c.ReturnCreateSuccess(ctx, id)
}

type FindByNameAndMobileQuery struct {
	Name   string     `json:"name" form:"name"`
	ID     types.Long `json:"id" form:"id"`
	Mobile string     `json:"mobile" form:"mobile"`
	Enable int64      `json:"enable" form:"enable"`
	types.Pagination
}
type FindByNameAndMobileResponse struct {
	Records []domain.User `json:"records" ` //记录
	Total   int64         `json:"total" `   //总数
}

// @Summary 根据名称查询用户信息
// @Tags user
// @Accept json
// @Produce  json
// @Param name query string false "用户名称"
// @Param page_size query string false "每页数据条数"
// @Param page query string false "第几页"
// @Success 200 {object} FindByNameAndMobileResponse
// @Router /user [get]
func (c *Controller) FindByName(ctx *gin.Context) {
	var query FindByNameAndMobileQuery

	if err := ctx.ShouldBindQuery(&query); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}

	users, total, err := c.UserQuery.FindByNameAndMobile(c.GetContext(ctx), query.ID, query.Name, query.Mobile, query.Enable, query.Pagination)
	if err != nil {
		c.ReturnErr(ctx, common.ServiceError(500, err))
		return
	}
	ctx.JSON(http.StatusOK, &FindByNameAndMobileResponse{
		Records: users,
		Total:   total,
	})
}

// @Summary 更新用户角色信息
// @Tags user
// @Accept json
// @Produce  json
// @Param id path int64 true "用户ID"
// @Param object body domain.ModifyUserRoleCommand true "更新用户角色信息"
// @Success 200 object object "{"msg": "modify success"}"
// @Router /user/{id}/role [patch]
func (c *Controller) ModifyUserRoleByID(ctx *gin.Context) {

	id, err := c.GetLongParam(ctx, "id")

	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}
	var command domain.ModifyUserRoleCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}
	command.ID = id

	if err = c.Service.ModifyUserRoleByID(ctx, command); err != nil {
		c.ReturnErr(ctx, err)
		return
	}

	c.ReturnModifySuccess(ctx)
}

// @Summary 更新用户状态
// @Tags user
// @Accept json
// @Produce  json
// @Param id path int64 true "用户ID"
// @Param object body domain.ModifyUserStatusCommand true "更新用户状态  1:启用  2：禁用"
// @Success 200 object object "{"msg": "modify success"}"
// @Router /user/{id}/status [patch]
// @Security ApiKeyAuth
func (c *Controller) ModifyUserStatusByID(ctx *gin.Context) {
	id, err := c.GetLongParam(ctx, "id")

	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}

	var command domain.ModifyUserStatusCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}
	command.ID = id
	if err = c.Service.ModifyUserStatusByID(ctx, command); err != nil {
		c.ReturnErr(ctx, err)
		return
	}
	c.ReturnModifySuccess(ctx)
}

// @Summary 更新用户密码
// @Tags user
// @Accept json
// @Produce  json
// @Param id path int64 true "用户ID"
// @Param object body domain.ChangePasswordCommand true "更新用户密码信息"
// @Success 200 object object "{"msg": "modify success"}"
// @Router /user/{id}/password [patch]
func (c *Controller) ModifyUserPasswordByID(ctx *gin.Context) {
	//oldPassword := ctx.Query("oldPassword")
	//newPassword := ctx.Query("newPassword")
	id, err := c.GetLongParam(ctx, "id")
	if err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参[用户ID]解析失败", err))
		return
	}

	var command domain.ChangePasswordCommand
	if err := ctx.ShouldBindJSON(&command); err != nil {
		c.ReturnErr(ctx, common.RequestParamError("入参解析失败", err))
		return
	}

	command.ID = id
	//command.Password = c.Service.SetPassword(command.Password)

	if err = c.Service.ModifyUserPasswordByID(ctx, command); err != nil {
		c.ReturnErr(ctx, err)
	}
	c.ReturnModifySuccess(ctx)
}
