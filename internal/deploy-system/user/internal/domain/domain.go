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
	Username string       `json:"username" gorm:"uniqueIndex:user_username_uindex"`
	Password string       `json:"-" gorm:"default:null"`
	Name     string       `json:"name"`
	Mobile   string       `json:"mobile"`
	Email    string       `json:"email"`
	Role     enum.SysRole `json:"role"`   // 0:普通用户，1:管理员，2:虚拟用户
	Enable   enum.Enable  `json:"enable"` //1：启用   2：禁用
}

func (u *User) VO() module.User {
	return module.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

type CreateUserCommand struct {
	Username string       `json:"username"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Mobile   string       `json:"mobile"`
	Email    string       `json:"email"`
	Role     enum.SysRole `json:"role"`
}

func (command *CreateUserCommand) ToUser() (*User, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return &User{
		Username: command.Username,
		//Password: command.Password,
		Name:   command.Name,
		Mobile: command.Mobile,
		Email:  command.Email,
		Role:   command.Role,
	}, nil
}

func (command *CreateUserCommand) Validate() error {

	command.Username = strings.TrimSpace(command.Username)
	if command.Username == "" {
		return common.RequestParamError("", errors.New("用户名不能为空"))
	}
	command.Password = strings.TrimSpace(command.Password)
	if command.Password == "" {
		return common.RequestParamError("", errors.New("密码不能为空"))
	}
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("名称不能为空"))
	}

	command.Mobile = strings.TrimSpace(command.Mobile)
	command.Email = strings.TrimSpace(command.Email)

	return nil
}

type ModifyUserCommand struct {
	ID       types.Long   `json:"-"`
	Username string       `json:"username"`
	Password string       `json:"password"`
	Name     string       `json:"name"`
	Mobile   string       `json:"mobile"`
	Email    string       `json:"email"`
	Role     enum.SysRole `json:"role"`
}

func (command *ModifyUserCommand) Validate() error {

	command.Username = strings.TrimSpace(command.Username)
	if command.Username == "" {
		return common.RequestParamError("", errors.New("用户名不能为空"))
	}
	command.Password = strings.TrimSpace(command.Password)
	if command.Password == "" {
		return common.RequestParamError("", errors.New("密码不能为空"))
	}
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("名称不能为空"))
	}
	command.Mobile = strings.TrimSpace(command.Mobile)
	command.Email = strings.TrimSpace(command.Email)

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

type ChangePasswordCommand struct {
	ID       types.Long `json:"-"`
	Password string     `json:"password"`
}
