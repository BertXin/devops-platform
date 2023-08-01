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
		Username:     "130xxafewrwqx",
		Name:         "吴不",
		Mobile:       "req",
		Email:        "xin.zhang@yatsenglobal.com",
		Avatar:       "无头像",
		WxWorkUserID: "130xxx",
		Role:         enum.SysRoleAdminUser,
	}
	id, err := s.Create(ctx, command)
	if err != nil {
		t.Fatal("创建用户失败", err.Error())
	}

	fmt.Print("创建用户成功,ID:", id)

}
