package service

import (
	"context"
	_ "devops-platform/internal/common/init"
	"devops-platform/internal/common/web"
	"devops-platform/internal/deploy-system/login/internal/domain"
	_ "devops-platform/internal/deploy-system/user/init"
	"devops-platform/pkg/beans"
	"fmt"
	"os"
	"testing"
)

var s KeyCloakService

func init() {
	fmt.Println("init in test")
	_ = os.Setenv(web.ModeKey, web.ModeUnitTesting)
	beans.Register(domain.BeanService, &s)
	beans.Start()
}
func TestKeyCloakService_Login(t *testing.T) {
	ctx := context.TODO()
	user, err := s.Login(ctx, "xin.zhang", "1997922@Zx")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(user)
}
