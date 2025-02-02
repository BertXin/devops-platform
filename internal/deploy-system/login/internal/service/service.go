package service

import (
	"context"
	"devops-platform/internal/common/config"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/user"
	"devops-platform/pkg/common"
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

/*
本地登录
*/
func (s *KeyCloakService) LocalLogin(ctx context.Context, login *domain.LoginRequest) (loginuser *domain.LoginUserVO, err error) {
	userinfo, err := s.UserRepository.GetByUsername(ctx, login.Username)
	if err != nil {
		logrus.Error("查询不到用户", err)
		return nil, err
	}
	//根据用户获取密码
	password, err := s.UserRepository.GetPasswordByUsername(ctx, userinfo.Username)
	if err != nil {
		logrus.Error("获取密码失败", err)
		return nil, err
	}
	//验证密码
	if err := common.ValidatePassword(password, login.Password); err != nil {
		// 返回错误
		logrus.Error("密码错误", err)
		return nil, err
	}

	return &domain.LoginUserVO{
		UserID:    userinfo.ID,
		LoginName: userinfo.Username,
		Username:  userinfo.Name,
		Role:      userinfo.Role,
	}, nil

}
