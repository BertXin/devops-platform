package domain

import (
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
)

//CREATE TABLE `permission` (
//  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增编号',
//  `created_at` datetime(6) DEFAULT NULL COMMENT '创建时间',
//  `created_by_id` bigint DEFAULT NULL COMMENT '创建人ID',
//  `created_by_name` varchar(16) DEFAULT NULL COMMENT '创建人名称',
//  `last_modified_at` datetime(6) DEFAULT NULL COMMENT '最后修改时间',
//  `last_modified_by_id` bigint DEFAULT NULL COMMENT '最后修改人ID',
//  `last_modified_by_name` varchar(16) DEFAULT NULL COMMENT '最后修改人名称',
//  `p_id` bigint DEFAULT NULL COMMENT '父权限ID',
//  `name` varchar(50) NOT NULL COMMENT '权限名称',
//  `sort` bigint DEFAULT NULL COMMENT '排序',
//  `path` varchar(200) DEFAULT NULL COMMENT '路径',
//  `method` varchar(20) DEFAULT NULL COMMENT '方法',
//  PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='权限';

type Permission struct {
	module.Module
	PID    types.Long `json:"p_id"`
	Name   string     `json:"name"`
	Sort   types.Long `json:"sort"`
	Path   string     `json:"path"`
	Method string     `json:"method"`
}

type PermissionVO struct {
	ID     types.Long `json:"id"`
	PID    types.Long `json:"p_id"`
	Name   string     `json:"name"`
	Sort   types.Long `json:"sort"`
	Path   string     `json:"path"`
	Method string     `json:"method"`
}

func (p *Permission) VO() PermissionVO {
	return PermissionVO{
		ID:     p.ID,
		PID:    p.PID,
		Name:   p.Name,
		Sort:   p.Sort,
		Path:   p.Path,
		Method: p.Method,
	}
}

type CreatePermissionCommand struct {
	Name   string     `json:"name"`
	PID    types.Long `json:"p_id"`
	Sort   types.Long `json:"sort"`
	Path   string     `json:"path"`
	Method string     `json:"method"`
}

func (command *CreatePermissionCommand) ToPermission() (*Permission, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return &Permission{
		Name:   command.Name,
		PID:    command.PID,
		Sort:   command.Sort,
		Path:   command.Path,
		Method: command.Method,
	}, nil
}

func (command *CreatePermissionCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("权限名称不能为空"))
	}
	return nil
}

type ModifyPermissionCommand struct {
	ID     types.Long `json:"-"`
	Name   string     `json:"name"`
	PID    types.Long `json:"p_id"`
	Sort   types.Long `json:"sort"`
	Path   string     `json:"path"`
	Method string     `json:"method"`
}

func (command *ModifyPermissionCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("权限名称不能为空"))
	}
	if command.ID <= 0 {
		return common.RequestParamError("", errors.New("权限编号不能为空"))
	}
	return nil
}
