package domain

import (
	"devops-platform/internal/pkg/enum"
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
	"time"
)

// Role 角色实体
type Role struct {
	module.Module
	Name        string      `json:"name" gorm:"size:128;comment:'角色名称'"`
	Code        string      `json:"code" gorm:"size:64;uniqueIndex;comment:'角色唯一标识符'"`
	Description string      `json:"description" gorm:"size:255;comment:'角色描述'"`
	Status      enum.Status `json:"status" gorm:"comment:'状态 1:启用 0:禁用'"`
	SortOrder   int         `json:"sort_order" gorm:"comment:'排序'"`
}

// Validate 验证角色
func (r *Role) Validate() error {
	if r.Name == "" {
		return errors.New("角色名称不能为空")
	}
	if r.Code == "" {
		return errors.New("角色标识不能为空")
	}
	return nil
}

// ToVO 转换为视图对象
func (r *Role) ToVO() *RoleVO {
	return &RoleVO{
		ID:          r.ID,
		Name:        r.Name,
		Code:        r.Code,
		Description: r.Description,
		Status:      r.Status,
		SortOrder:   r.SortOrder,
		CreatedAt:   r.CreatedAt.Time,
		UpdatedAt:   r.LastModifiedAt.Time,
	}
}

// RoleVO 角色视图对象
type RoleVO struct {
	ID          types.Long  `json:"id"`
	Name        string      `json:"name"`
	Code        string      `json:"code"`
	Description string      `json:"description"`
	Status      enum.Status `json:"status"`
	SortOrder   int         `json:"sort_order"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

// Permission 权限实体
type Permission struct {
	module.Module
	ParentID   types.Long  `json:"parent_id" gorm:"comment:'父权限ID'"`
	Name       string      `json:"name" gorm:"size:128;comment:'权限名称'"`
	Type       string      `json:"type" gorm:"size:20;comment:'权限类型: menu, api, button'"`
	Path       string      `json:"path" gorm:"size:200;comment:'路径'"`
	Method     string      `json:"method" gorm:"size:20;comment:'HTTP方法'"`
	Icon       string      `json:"icon" gorm:"size:128;comment:'图标'"`
	Component  string      `json:"component" gorm:"size:128;comment:'组件路径'"`
	Permission string      `json:"permission" gorm:"size:128;comment:'权限标识'"`
	Status     enum.Status `json:"status" gorm:"comment:'状态 1:启用 0:禁用'"`
	Hidden     bool        `json:"hidden" gorm:"comment:'是否隐藏'"`
	SortOrder  int         `json:"sort_order" gorm:"comment:'排序'"`
	ApiPath    string      `json:"-" gorm:"-"` // API路径（用于Casbin集成）
	ApiMethod  string      `json:"-" gorm:"-"` // API方法（用于Casbin集成）
}

// Validate 验证权限
func (p *Permission) Validate() error {
	if p.Name == "" {
		return errors.New("权限名称不能为空")
	}
	if p.Type == "" {
		return errors.New("权限类型不能为空")
	}
	if p.Type == PermTypeApi && p.Path == "" {
		return errors.New("API权限必须指定路径")
	}
	if p.Type == PermTypeButton && p.Permission == "" {
		return errors.New("按钮权限必须指定权限标识")
	}
	return nil
}

// ToVO 转换为视图对象
func (p *Permission) ToVO() *PermissionVO {
	return &PermissionVO{
		ID:         p.ID,
		ParentID:   p.ParentID,
		Name:       p.Name,
		Type:       p.Type,
		Path:       p.Path,
		Method:     p.Method,
		Icon:       p.Icon,
		Component:  p.Component,
		Permission: p.Permission,
		Status:     p.Status,
		Hidden:     p.Hidden,
		SortOrder:  p.SortOrder,
		CreatedAt:  p.CreatedAt.Time,
		UpdatedAt:  p.LastModifiedAt.Time,
	}
}

// PermissionVO 权限视图对象
type PermissionVO struct {
	ID         types.Long      `json:"id"`
	ParentID   types.Long      `json:"parent_id"`
	Name       string          `json:"name"`
	Type       string          `json:"type"`
	Path       string          `json:"path"`
	Method     string          `json:"method"`
	Icon       string          `json:"icon"`
	Component  string          `json:"component"`
	Permission string          `json:"permission"`
	Status     enum.Status     `json:"status"`
	Hidden     bool            `json:"hidden"`
	SortOrder  int             `json:"sort_order"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	Children   []*PermissionVO `json:"children,omitempty"`
}

// MenuVO 菜单视图对象
type MenuVO struct {
	ID        types.Long `json:"id"`
	ParentID  types.Long `json:"parent_id"`
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Component string     `json:"component"`
	Icon      string     `json:"icon"`
	SortOrder int        `json:"sort_order"`
	Hidden    bool       `json:"hidden"`
	Children  []*MenuVO  `json:"children,omitempty"`
}

// RolePermission 角色权限关联
type RolePermission struct {
	ID           types.Long `json:"id" gorm:"primaryKey;autoIncrement"`
	RoleID       types.Long `json:"role_id" gorm:"index:idx_role_perm;comment:'角色ID'"`
	PermissionID types.Long `json:"permission_id" gorm:"index:idx_role_perm;comment:'权限ID'"`
	CreatedAt    time.Time  `json:"created_at"`
}

// UserRole 用户角色关联
type UserRole struct {
	ID        types.Long `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    types.Long `json:"user_id" gorm:"index:idx_user_role;comment:'用户ID'"`
	RoleID    types.Long `json:"role_id" gorm:"index:idx_user_role;comment:'角色ID'"`
	CreatedAt time.Time  `json:"created_at"`
}

// CasbinRule Casbin规则实体
type CasbinRule struct {
	ID    types.Long `json:"id" gorm:"primaryKey;autoIncrement"`
	PType string     `json:"p_type" gorm:"size:100;comment:'策略类型'"`
	V0    string     `json:"v0" gorm:"size:100;comment:'角色或用户'"`
	V1    string     `json:"v1" gorm:"size:100;comment:'资源'"`
	V2    string     `json:"v2" gorm:"size:100;comment:'操作'"`
	V3    string     `json:"v3" gorm:"size:100;comment:'域'"`
	V4    string     `json:"v4" gorm:"size:100;comment:'策略规则'"`
	V5    string     `json:"v5" gorm:"size:100;comment:'扩展字段'"`
}

// CreateRoleCommand 创建角色命令
type CreateRoleCommand struct {
	Name        string `json:"name" binding:"required"`
	Code        string `json:"code" binding:"required"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

// ToRole 转换为角色实体
func (command *CreateRoleCommand) ToRole() (*Role, error) {
	command.Name = strings.TrimSpace(command.Name)
	command.Code = strings.TrimSpace(command.Code)

	role := &Role{
		Name:        command.Name,
		Code:        command.Code,
		Description: command.Description,
		Status:      enum.StatusEnabled,
		SortOrder:   command.SortOrder,
	}

	err := role.Validate()
	if err != nil {
		return nil, common.RequestParamError("", err)
	}

	return role, nil
}

// UpdateRoleCommand 更新角色命令
type UpdateRoleCommand struct {
	ID          types.Long  `json:"-"`
	Name        string      `json:"name" binding:"required"`
	Code        string      `json:"code" binding:"required"`
	Description string      `json:"description"`
	Status      enum.Status `json:"status"`
	SortOrder   int         `json:"sort_order"`
}

// Validate 验证命令参数
func (command *UpdateRoleCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("角色名称不能为空"))
	}

	command.Code = strings.TrimSpace(command.Code)
	if command.Code == "" {
		return common.RequestParamError("", errors.New("角色标识不能为空"))
	}

	return nil
}

// RoleQuery 角色查询条件
type RoleQuery struct {
	Name   string      `json:"name" form:"name"`
	Code   string      `json:"code" form:"code"`
	Status enum.Status `json:"status" form:"status"`
	Page   int         `json:"page" form:"page"`
	Size   int         `json:"size" form:"size"`
}

// CreatePermissionCommand 创建权限命令
type CreatePermissionCommand struct {
	ParentID   types.Long  `json:"parent_id"`
	Name       string      `json:"name" binding:"required"`
	Type       string      `json:"type" binding:"required"`
	Path       string      `json:"path"`
	Method     string      `json:"method"`
	Icon       string      `json:"icon"`
	Component  string      `json:"component"`
	Permission string      `json:"permission"`
	Status     enum.Status `json:"status"`
	Hidden     bool        `json:"hidden"`
	SortOrder  int         `json:"sort_order"`
}

// ToPermission 转换为权限实体
func (command *CreatePermissionCommand) ToPermission() (*Permission, error) {
	command.Name = strings.TrimSpace(command.Name)
	command.Type = strings.TrimSpace(command.Type)

	permission := &Permission{
		ParentID:   command.ParentID,
		Name:       command.Name,
		Type:       command.Type,
		Path:       command.Path,
		Method:     command.Method,
		Icon:       command.Icon,
		Component:  command.Component,
		Permission: command.Permission,
		Status:     command.Status,
		Hidden:     command.Hidden,
		SortOrder:  command.SortOrder,
	}

	if permission.Status == 0 {
		permission.Status = enum.StatusEnabled
	}

	err := permission.Validate()
	if err != nil {
		return nil, common.RequestParamError("", err)
	}

	return permission, nil
}

// UpdatePermissionCommand 更新权限命令
type UpdatePermissionCommand struct {
	ID         types.Long  `json:"-"`
	ParentID   types.Long  `json:"parent_id"`
	Name       string      `json:"name" binding:"required"`
	Type       string      `json:"type" binding:"required"`
	Path       string      `json:"path"`
	Method     string      `json:"method"`
	Icon       string      `json:"icon"`
	Component  string      `json:"component"`
	Permission string      `json:"permission"`
	Status     enum.Status `json:"status"`
	Hidden     bool        `json:"hidden"`
	SortOrder  int         `json:"sort_order"`
}

// Validate 验证命令参数
func (command *UpdatePermissionCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("权限名称不能为空"))
	}

	command.Type = strings.TrimSpace(command.Type)
	if command.Type == "" {
		return common.RequestParamError("", errors.New("权限类型不能为空"))
	}

	if command.Type == PermTypeApi && command.Path == "" {
		return common.RequestParamError("", errors.New("API权限必须指定路径"))
	}

	if command.Type == PermTypeButton && command.Permission == "" {
		return common.RequestParamError("", errors.New("按钮权限必须指定权限标识"))
	}

	return nil
}

// PermissionQuery 权限查询条件
type PermissionQuery struct {
	Name     string      `json:"name" form:"name"`
	Type     string      `json:"type" form:"type"`
	Path     string      `json:"path" form:"path"`
	Status   enum.Status `json:"status" form:"status"`
	ParentID types.Long  `json:"parent_id" form:"parent_id"`
	Page     int         `json:"page" form:"page"`
	Size     int         `json:"size" form:"size"`
}
