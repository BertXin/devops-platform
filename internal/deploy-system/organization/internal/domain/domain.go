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

// Department 部门实体
type Department struct {
	module.Module
	ParentID    types.Long  `json:"parent_id" gorm:"comment:'父部门ID'"`
	Name        string      `json:"name" gorm:"size:128;comment:'部门名称'"`
	Code        string      `json:"code" gorm:"size:64;uniqueIndex;comment:'部门编码'"`
	Description string      `json:"description" gorm:"size:255;comment:'描述'"`
	Status      enum.Status `json:"status" gorm:"comment:'状态 1:启用 0:禁用'"`
	Sort        int         `json:"sort_order" gorm:"column:sort;comment:'排序'"`
}

// Validate 验证部门
func (d *Department) Validate() error {
	if d.Name == "" {
		return errors.New("部门名称不能为空")
	}
	if d.Code == "" {
		return errors.New("部门编码不能为空")
	}
	return nil
}

// ToVO 转换为视图对象
func (d *Department) ToVO() *DepartmentVO {
	return &DepartmentVO{
		ID:          d.ID,
		ParentID:    d.ParentID,
		Name:        d.Name,
		Code:        d.Code,
		Description: d.Description,
		Status:      d.Status,
		SortOrder:   d.Sort,
		CreatedAt:   d.CreatedAt.Time,
		UpdatedAt:   d.LastModifiedAt.Time,
	}
}

// DepartmentVO 部门视图对象
type DepartmentVO struct {
	ID          types.Long      `json:"id"`
	ParentID    types.Long      `json:"parent_id"`
	Name        string          `json:"name"`
	Code        string          `json:"code"`
	Description string          `json:"description"`
	Status      enum.Status     `json:"status"`
	SortOrder   int             `json:"sort_order"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	Children    []*DepartmentVO `json:"children,omitempty"`
	Users       []*UserVO       `json:"users,omitempty"`
}

// UserDepartment 用户部门关联
type UserDepartment struct {
	ID           types.Long `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID       types.Long `json:"user_id" gorm:"index:idx_user_dept;comment:'用户ID'"`
	DepartmentID types.Long `json:"department_id" gorm:"index:idx_user_dept;comment:'部门ID'"`
	Type         int        `json:"type" gorm:"comment:'关系类型 0:普通成员 1:负责人'"`
	CreatedAt    time.Time  `json:"created_at"`
}

// UserVO 用户视图对象（简化版，仅包含部门展示需要的字段）
type UserVO struct {
	ID        types.Long  `json:"id"`
	Username  string      `json:"username"`
	RealName  string      `json:"real_name"`
	Email     string      `json:"email"`
	Mobile    string      `json:"mobile"`
	Status    enum.Status `json:"status"`
	IsLeader  bool        `json:"is_leader"` // 是否为部门负责人
	CreatedAt time.Time   `json:"created_at"`
}

// CreateDepartmentCommand 创建部门命令
type CreateDepartmentCommand struct {
	ParentID    types.Long  `json:"parent_id"`
	Name        string      `json:"name" binding:"required"`
	Code        string      `json:"code" binding:"required"`
	Description string      `json:"description"`
	Status      enum.Status `json:"status"`
	SortOrder   int         `json:"sort_order"`
}

// ToDepartment 转换为部门实体
func (command *CreateDepartmentCommand) ToDepartment() (*Department, error) {
	command.Name = strings.TrimSpace(command.Name)
	command.Code = strings.TrimSpace(command.Code)

	dept := &Department{
		ParentID:    command.ParentID,
		Name:        command.Name,
		Code:        command.Code,
		Description: command.Description,
		Status:      command.Status,
		Sort:        command.SortOrder,
	}

	if dept.Status == 0 {
		dept.Status = enum.StatusEnabled
	}

	err := dept.Validate()
	if err != nil {
		return nil, common.RequestParamError("", err)
	}

	return dept, nil
}

// UpdateDepartmentCommand 更新部门命令
type UpdateDepartmentCommand struct {
	ID          types.Long  `json:"-"`
	ParentID    types.Long  `json:"parent_id"`
	Name        string      `json:"name" binding:"required"`
	Code        string      `json:"code" binding:"required"`
	Description string      `json:"description"`
	Status      enum.Status `json:"status"`
	Sort        int         `json:"sort"`
}

// Validate 验证命令参数
func (command *UpdateDepartmentCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("部门名称不能为空"))
	}

	command.Code = strings.TrimSpace(command.Code)
	if command.Code == "" {
		return common.RequestParamError("", errors.New("部门编码不能为空"))
	}

	return nil
}

// DepartmentQuery 部门查询条件
type DepartmentQuery struct {
	Name     string      `json:"name" form:"name"`
	Code     string      `json:"code" form:"code"`
	Status   enum.Status `json:"status" form:"status"`
	ParentID types.Long  `json:"parent_id" form:"parent_id"`
	Page     int         `json:"page" form:"page"`
	Size     int         `json:"size" form:"size"`
}

// UserQuery 用户查询条件
type UserQuery struct {
	Username string      `json:"username" form:"username"`
	RealName string      `json:"real_name" form:"real_name"`
	Email    string      `json:"email" form:"email"`
	Mobile   string      `json:"mobile" form:"mobile"`
	Status   enum.Status `json:"status" form:"status"`
	Page     int         `json:"page" form:"page"`
	Size     int         `json:"size" form:"size"`
}
