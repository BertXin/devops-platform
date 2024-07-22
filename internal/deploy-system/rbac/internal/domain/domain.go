package domain

import (
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
)

// ```
//
//	`name` varchar(128) DEFAULT NULL COMMENT '''角色名称''',
//	`desc` varchar(128) DEFAULT NULL COMMENT '''角色描述''',
//	`code` varchar(32) DEFAULT NULL COMMENT '''角色标识''',
//	`permission_id` bigint(20) DEFAULT NULL COMMENT '''权限id外键''',
//
// ```
type Role struct {
	module.Module
	Name         string     `json:"name"`
	Desc         string     `json:"desc"`
	Code         string     `json:"code"`
	PermissionID types.Long `json:"permission_id"`
}

type RoleVO struct {
	ID           types.Long `json:"id"`
	Name         string     `json:"name"`
	Desc         string     `json:"desc"`
	Code         string     `json:"code"`
	PermissionID types.Long `json:"permission_id"`
}

func (r *Role) VO() RoleVO {
	return RoleVO{
		ID:           r.ID,
		Name:         r.Name,
		Desc:         r.Desc,
		Code:         r.Code,
		PermissionID: r.PermissionID,
	}
}

type CreateRoleCommand struct {
	Name         string     `json:"name"`
	Desc         string     `json:"desc"`
	Code         string     `json:"code"`
	PermissionID types.Long `json:"permission_id"`
}

func (command *CreateRoleCommand) ToRole() (*Role, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return &Role{
		Name:         command.Name,
		Desc:         command.Desc,
		Code:         command.Code,
		PermissionID: command.PermissionID,
	}, nil
}

// Validate 方法用于验证 CreateRoleCommand 的数据
func (command *CreateRoleCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("角色名称不能为空"))
	}
	if command.Code == "" {
		return common.RequestParamError("", errors.New("角色标识不能为空"))
	}
	if command.PermissionID == 0 {
		return common.RequestParamError("", errors.New("权限id不能为空"))
	}
	return nil
}

// ModifyRoleCommand 定义更新部门的命令
type ModifyRoleCommand struct {
	ID           types.Long `json:"-"`
	Name         string     `json:"name"`
	Desc         string     `json:"desc"`
	Code         string     `json:"code"`
	PermissionID types.Long `json:"permission_id"`
}

// Validate 方法用于验证 ModifyRoleCommand 的数据
func (command *ModifyRoleCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("角色名称不能为空"))
	}
	if command.Code == "" {
		return common.RequestParamError("", errors.New("角色标识不能为空"))
	}
	if command.PermissionID == 0 {
		return common.RequestParamError("", errors.New("权限id不能为空"))
	}
	return nil
}
