package service

import (
	"context"
	"devops-platform/internal/common/config"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/user"
	"github.com/Nerzal/gocloak/v13"
	"github.com/sirupsen/logrus"
)

type KeyCloakService struct {
	client         *gocloak.GoCloak
	config         domain.Config
	UserService    user.Service    `inject:"UserService"`
	UserRepository user.Repository `inject:"UserRepository"`
}

func (s *KeyCloakService) Inject(getBean func(string) interface{}) {
	ok := true

	s.config, ok = getBean(config.BeanSsoLogin).(domain.Config)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", config.BeanSsoLogin)
		return
	}
	s.init()
}

/*
 * 初始化keycloak服务器配置
 */
func (c *KeyCloakService) init() {
	c.client = gocloak.NewClient(c.config.GetAuthUrl())
}

func (s *KeyCloakService) Login(ctx context.Context, username string, password string) (user *gocloak.UserInfo, err error) {
	token, err := s.client.Login(ctx, s.config.GetClientId(), s.config.GetClientSecret(), s.config.GetRealm(), username, password)
	if err != nil {
		return nil, err
	}
	user, err = s.client.GetUserInfo(ctx, token.AccessToken, s.config.GetRealm())
	if err != nil {
		logrus.Error(err)
	}
	return
}
