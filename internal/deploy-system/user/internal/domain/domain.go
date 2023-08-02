package domain

import (
	"devops-platform/internal/pkg/enum"
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
)

type User struct {
	module.Module
	Username       string       `json:"username" gorm:"uniqueIndex:user_username_uindex"`
	Name           string       `json:"name"`
	Mobile         string       `json:"mobile"`
	Email          string       `json:"email"`
	Role           enum.SysRole `json:"role"` // 0:普通用户，1:管理员，2:虚拟用户
	OrgDisplayName string       `json:"org_display_name"`
	Avatar         string       `json:"avatar"`
	WxWorkUserID   string       `json:"wx_work_user_id"`
	GitlabUserID   int          `json:"gitlab_user_id"`
	Enable         enum.Enable  `json:"enable"` //1：启用   2：禁用
}

func (u *User) VO() module.User {
	return module.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

type CreateUserCommand struct {
	Username       string       `json:"username"`
	Name           string       `json:"name"`
	Mobile         string       `json:"mobile"`
	Email          string       `json:"email"`
	Role           enum.SysRole `json:"role"`
	OrgDisplayName string       `json:"org_display_name"`
	Avatar         string       `json:"avatar"`
	WxWorkUserID   string       `json:"wx_work_user_id"`
}

func (command *CreateUserCommand) ToUser() (*User, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return &User{
		Username:       command.Username,
		Name:           command.Name,
		Mobile:         command.Mobile,
		Email:          command.Email,
		Role:           command.Role,
		OrgDisplayName: command.OrgDisplayName,
		Avatar:         command.Avatar,
		WxWorkUserID:   command.WxWorkUserID,
	}, nil
}

func (command *CreateUserCommand) Validate() error {

	command.Username = strings.TrimSpace(command.Username)
	if command.Username == "" {
		return common.RequestParamError("", errors.New("用户名不能为空"))
	}

	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("名称不能为空"))
	}
	command.Mobile = strings.TrimSpace(command.Mobile)
	command.Email = strings.TrimSpace(command.Email)
	command.Avatar = strings.TrimSpace(command.Avatar)
	command.WxWorkUserID = strings.TrimSpace(command.WxWorkUserID)

	return nil
}

type ModifyUserCommand struct {
	ID             types.Long   `json:"-"`
	Username       string       `json:"username"`
	Name           string       `json:"name"`
	Mobile         string       `json:"mobile"`
	Email          string       `json:"email"`
	Role           enum.SysRole `json:"role"`
	OrgDisplayName string       `json:"org_display_name"`
	Avatar         string       `json:"avatar"`
	WxWorkUserID   string       `json:"wx_work_user_id"`
	GitlabUserID   int          `json:"gitlab_user_id"`
}

func (command *ModifyUserCommand) Validate() error {

	command.Username = strings.TrimSpace(command.Username)
	if command.Username == "" {
		return common.RequestParamError("", errors.New("用户名不能为空"))
	}

	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("名称不能为空"))
	}

	command.WxWorkUserID = strings.TrimSpace(command.WxWorkUserID)
	if command.WxWorkUserID == "" {
		return common.RequestParamError("", errors.New("企业微信账号不能为空"))
	}

	if command.GitlabUserID < 0 {
		return common.RequestNotFoundError("gitlab账号ID不能小于零")
	}
	command.Mobile = strings.TrimSpace(command.Mobile)
	command.Email = strings.TrimSpace(command.Email)
	command.Avatar = strings.TrimSpace(command.Avatar)

	return nil
}

type ModifyUserRoleCommand struct {
	ID   types.Long   `json:"-"`
	Role enum.SysRole `json:"role"`
}

type ModifyUserStatusCommand struct {
	ID     types.Long  `json:"-"`
	Status enum.Enable `json:"status"`
}

type ModifyUserGitlabUserIDCommand struct {
	ID           types.Long `json:"-"`
	GitlabUserID int        `json:"gitlab_user_id"`
}

type SyncUserMessageCommand struct {
	ID             types.Long `json:"id"`
	Username       string     `json:"username"`
	Name           string     `json:"name"`
	OldMobile      string     `json:"oldMobile"`
	Mobile         string     `json:"mobile"`
	Email          string     `json:"email"`
	Gender         string     `json:"gender"`
	Enable         int        `json:"enable"`
	EmployeeType   int        `json:"employeeType"`
	employeeNumber string     `json:"employeeNumber"`
	CorpFlag       bool       `json:"corpFlag"`
}
