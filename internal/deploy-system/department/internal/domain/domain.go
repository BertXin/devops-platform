package domain

import (
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
)

// Dept 代表部门实体
type Dept struct {
	module.Module
	Name     string     `json:"name" gorm:"comment:'部门名称'"`
	Sort     types.Long `json:"sort" gorm:"comment:'排序'"`
	ParentID types.Long `json:"parent_id" gorm:"comment:'父级部门ID'"` // 假设使用 types.Long 类型
	//Description string     `json:"description" gorm:"comment:'部门描述'"`
}

// DeptVO 定义部门视图对象，用于API响应
type DeptVO struct {
	ID       types.Long `json:"id"`
	Name     string     `json:"name"`
	Sort     types.Long `json:"sort"`
	ParentID types.Long `json:"parent_id"`
	//Description  string     `json:"description"`
}

// VO 方法将 Dept 转换为 DeptVO
func (d *Dept) VO() DeptVO {
	return DeptVO{
		ID:       d.ID,
		Name:     d.Name,
		Sort:     d.Sort,
		ParentID: d.ParentID,
		//Description:  d.Description,
	}
}

// CreateDeptCommand 定义创建部门的命令
type CreateDeptCommand struct {
	Name     string     `json:"name"`
	Sort     types.Long `json:"sort"`
	ParentID types.Long `json:"parent_id"`
	//Description string `json:"description"`
}

// ToDept 方法将 CreateDeptCommand 转换为 Dept
func (command *CreateDeptCommand) ToDept() (*Dept, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return &Dept{
		Name:     command.Name,
		Sort:     command.Sort,
		ParentID: command.ParentID,
		//Description: command.Description,
	}, nil
}

// Validate 方法用于验证 CreateDeptCommand 的数据
func (command *CreateDeptCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("部门名称不能为空"))
	}
	if command.ParentID == 0 {
		return common.RequestParamError("", errors.New("父级部门ID不能为空"))
	}
	return nil
}

// ModifyDeptCommand 定义更新部门的命令
type ModifyDeptCommand struct {
	ID       types.Long `json:"-"`
	Name     string     `json:"name"`
	Sort     int        `json:"sort"`
	ParentID types.Long `json:"parent_id"`
	//Description string     `json:"description"`
}

// Validate 方法用于验证 ModifyDeptCommand 的数据
func (command *ModifyDeptCommand) Validate() error {
	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("部门名称不能为空"))
	}
	if command.ParentID == 0 {
		return common.RequestParamError("", errors.New("父级部门ID不能为0"))
	}
	return nil
}
