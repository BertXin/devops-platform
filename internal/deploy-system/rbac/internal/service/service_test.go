package service

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/rbac/internal/domain"
	"devops-platform/internal/deploy-system/rbac/internal/repository"
	"devops-platform/pkg/beans"
	"fmt"
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

func TestService_CreateRole(t *testing.T) {
	ctx := context.TODO()
	role, err := s.CreateRole(ctx, &domain.CreateRoleCommand{
		Code:         "test",
		Desc:         "test",
		Name:         "test",
		PermissionID: 1,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(role)
}

func TestService_ModifyRoleByID(t *testing.T) {

	ctx := context.TODO()
	err := s.ModifyRoleByID(ctx, &domain.ModifyRoleCommand{
		Code:         "test20",
		Desc:         "test2",
		ID:           4,
		Name:         "test2",
		PermissionID: 2,
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(err)
}

func TestService_DeleteRoleByID(t *testing.T) {
	ctx := context.TODO()
	err := s.DeleteRoleByID(ctx, 4)
	if err != nil {
		t.Error(err)
	}
}

func TestService_FindRoleByID(t *testing.T) {

	ctx := context.TODO()
	role, err := s.FindRoleByID(ctx, 2)
	if err != nil {
		t.Error(err)
	}
	t.Log(role)
}
