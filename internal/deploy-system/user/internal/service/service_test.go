package service

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/user"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/internal/deploy-system/user/internal/repository"
	"devops-platform/internal/pkg/enum"
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

func TestService_Create(t *testing.T) {
	ctx := context.TODO()
	command := &user.CreateUserCommand{
		Username: "130",
		Password: "12334555",
		Name:     "吴不",
		Mobile:   "req",
		Email:    "xin.zhang@yatsenglobal.com",
		Role:     enum.SysRoleAdminUser,
	}
	id, err := s.Create(ctx, command)
	if err != nil {
		t.Fatal("创建用户失败", err.Error())
	}

	fmt.Print("创建用户成功,ID:", id)

}

func TestService_ModifyUserByID(t *testing.T) {
	m := &domain.ModifyUserCommand{
		ID:       1,
		Username: "zhangxin",
		Name:     "zhangxin",
		Mobile:   "zhangxin",
		Email:    "zhangxin",
		Role:     enum.SysRoleVirtualUser,
	}
	err := s.ModifyUserByID(context.TODO(), m)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("成功")
}

func TestModifyUserPasswordByName(t *testing.T) {
	m := &domain.ChangePasswordCommand{
		ID:       16,
		Password: "123456",
	}
	err := s.ModifyUserPasswordByID(context.TODO(), *m)
	if err != nil {
		t.Fatal(err)
	}
}
