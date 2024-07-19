package repository

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/department/internal/domain"
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/beans"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var r Repository

func init() {
	fmt.Println("init in test")
	_ = os.Setenv(web.ModeKey, web.ModeUnitTesting)
	beans.Register(domain.BeanRepository, &r)
	beans.Start()
}

func TestCreate(t *testing.T) {
	user := &domain.Dept{
		Module:   module.Module{},
		Name:     "技术部",
		Sort:     0,
		ParentID: 0,
	}
	err := r.Create(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetDeptByID(t *testing.T) {
	// 根据ID获取部门的测试
	ctx := context.TODO()
	// 假设存在一个ID为1的部门
	dept, err := r.GetByID(ctx, 1)
	assert.NoError(t, err, "根据ID获取部门失败")
	assert.NotNil(t, dept, "部门不存在")
}

func TestUpdateDept(t *testing.T) {
	// 更新部门的测试
	ctx := context.TODO()
	// 假设存在一个ID为1的部门可以更新
	dept, err := r.GetByID(ctx, 1)
	assert.NoError(t, err, "获取部门失败")

	// 修改部门名称
	dept.Name = "IT"

	// 更新部门信息
	err = r.Update(ctx, dept)
	assert.NoError(t, err, "更新部门失败")

	// 检查部门是否更新
	updatedDept, getErr := r.GetByID(ctx, dept.ID)
	assert.NoError(t, getErr, "根据ID获取部门失败")
	assert.Equal(t, dept.Name, updatedDept.Name, "部门名称更新不匹配")
}
