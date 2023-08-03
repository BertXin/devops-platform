package service

import (
	"devops-platform/internal/common/config"
	"devops-platform/internal/deploy-system/client/keycloak/internal/domain"
	"github.com/Nerzal/gocloak/v13"
	"github.com/sirupsen/logrus"
)

type Service struct {
	client *gocloak.GoCloak
	config domain.Config
}

func (s *Service) Inject(getBean func(string) interface{}) {
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
func (c *Service) init() {
	c.client = gocloak.NewClient(c.config.GetAuthUrl())
}
