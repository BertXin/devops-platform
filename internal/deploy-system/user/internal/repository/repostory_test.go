package repository

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/user/internal/domain"
	"devops-platform/internal/pkg/module"
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

func TestCreate(t *testing.T) {
	user := &domain.User{
		Module:         module.Module{},
		Username:       "zhangxiu233",
		Name:           "test",
		Mobile:         "",
		Email:          "",
		Role:           0,
		OrgDisplayName: "",
		Avatar:         "",
		WxWorkUserID:   "",
		GitlabUserID:   0,
		Enable:         1,
	}
	err := r.Create(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}
}
