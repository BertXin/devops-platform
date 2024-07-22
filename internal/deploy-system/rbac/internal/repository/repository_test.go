package repository

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/rbac/internal/domain"
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
func TestRepository_CreateRole(t *testing.T) {
	role := &domain.Role{
		Name:         "超级管理员",
		Code:         "super",
		PermissionID: 1,
		Desc:         "超级管理员拥有所有权限",
	}
	if err := r.CreateRole(context.TODO(), role); err != nil {
		t.Error(err)
	}
	t.Log("创建成功")
}
func TestRepository_Delete(t *testing.T) {
	if err := r.DeleteRole(context.TODO(), 1); err != nil {
		t.Error(err)
	}
	t.Log("delete success")
}

func TestRepository_UpdateRole(t *testing.T) {
	ctx := context.TODO()
	role, err := r.FindRoleByID(ctx, 2)
	assert.NoError(t, err, "获取部门失败")

	role.PermissionID = 3
	if err := r.UpdateRole(ctx, role); err != nil {
		assert.NoError(t, err, "更新部门失败")
	}
	t.Log("update success")

}
