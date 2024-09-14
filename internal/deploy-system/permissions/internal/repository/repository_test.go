package repository

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/permissions/internal/domain"
	"devops-platform/pkg/beans"
	"fmt"
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

func TestPRepository_CreatePerm(t *testing.T) {
	perm := &domain.Permission{
		PID:    0,
		Name:   "系统管理",
		Method: "GET",
		Path:   "/system",
		Sort:   1,
	}
	if err := r.CreatePerm(context.TODO(), perm); err != nil {
		t.Error(err)
	}
	t.Log("创建成功")
}

func TestRepository_FindPermByPID(t *testing.T) {
	perm, err := r.FindPermByPID(context.TODO(), 0)
	if err != nil {
		t.Error(err)
	}
	t.Log(perm)

}
