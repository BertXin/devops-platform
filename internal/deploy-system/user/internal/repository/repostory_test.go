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
		Module:   module.Module{},
		Username: "zhangxiu233222",
		Password: "123456",
		Name:     "test",
		Mobile:   "",
		Email:    "",
		Role:     0,
		Enable:   1,
	}
	err := r.Create(context.TODO(), user)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRepository_GetPasswordByUsername(t *testing.T) {
	password, err := r.GetPasswordByUsername(context.TODO(), "xin.zhang")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(password)
	//$2a$10$q1tKmdT8.cfZH1Lw3mhYLe3pGMUhy7zZIAFYgrSZuXGGQgfCtnToO
	//$2a$10$VazwsWi9v/cBe9Dm4OXCyOFXBXl1D1mPkryNyR4Xik/2TyZIqX9Fa
	//$2a$10$5so.oA32gCWS3/GdZw6iPu/if60bXxq9CMRnHR7a2fbdhfPgjplZy
	//$2a$10$Yvfz08oaCbnEkugFOIuCveKfbMmKAqqQzKh/F/s0O4IilSJT6SQqy
}

func TestRepository_GetByID(t *testing.T) {
	user, err := r.GetByID(context.TODO(), 103)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}

func TestRepository_GetByUsername(t *testing.T) {
	user, err := r.GetByUsername(context.TODO(), "xin.zhang1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
}
