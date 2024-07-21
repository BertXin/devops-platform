package service

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/department/internal/domain"
	"devops-platform/internal/deploy-system/department/internal/repository"
	"devops-platform/pkg/beans"
	"devops-platform/pkg/types"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var s Service

func init() {
	fmt.Println("init in test")
	_ = os.Setenv(web.ModeKey, web.ModeUnitTesting)
	beans.Register(domain.BeanRepository, &repository.Repository{})
	beans.Register(domain.BeanService, &s)
	beans.Start()
}
func TestService_Create(t *testing.T) {
	ctx := context.TODO()

	_, err := s.Create(ctx, &domain.CreateDeptCommand{
		Name:     "测试",
		ParentID: 0,
		Sort:     0,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}
func TestService_DeleteDeptByID(t *testing.T) {
	ctx := context.TODO()

	err := s.DeleteDeptByID(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}
func TestService_ModifyDeptParentIDByID(t *testing.T) {
	ctx := context.TODO()
	err := s.ModifyDeptParentIDByID(ctx, &domain.ModifyDeptCommand{
		ID:       2,
		ParentID: 1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}
func TestService_ModifyDeptNameByID(t *testing.T) {
	ctx := context.TODO()
	err := s.ModifyDeptNameByID(ctx, &domain.ModifyDeptCommand{
		ID:   2,
		Name: "测试2",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}
func TestService_FindDeptByID(t *testing.T) {
	ctx := context.TODO()
	dept, err := s.FindDeptByID(ctx, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(dept.VO())
	t.Log(err)
}

func TestService_FindDeptByName(t *testing.T) {
	ctx := context.TODO()
	depts, total, err := s.FindDeptByName(ctx, "测试", 0, types.Pagination{})
	if err != nil {
		assert.NoError(t, err, "根据名称查询部门失败")
	}
	fmt.Println(depts)
	t.Log(total)
}
