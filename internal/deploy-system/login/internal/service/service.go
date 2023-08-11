package service

import (
	"context"
	"devops-platform/internal/common/config"
	"devops-platform/internal/deploy-system/login/internal/domain"
	"devops-platform/internal/deploy-system/user"
	"github.com/Nerzal/gocloak/v13"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
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
 根据用户密码换取token信息
*/
//func (s *KeyCloakService) GetSsoToken(ctx context.Context, username string, password string) (token string, err error) {
//	login, err := s.client.Login(ctx, s.config.GetClientId(), s.config.GetClientSecret(), s.config.GetRealm(), username, password)
//	if err != nil {
//		logrus.Error("请求API出错：", err)
//		return
//	}
//	accessToken := login.AccessToken
//	return accessToken, err
//}

/*
 * 根据token换取登录用户信息
 */
//func (s *KeyCloakService) CheckToken(ctx context.Context, token string) (checkToken *domain.SsoCheckTokenVO, err error) {
//	info, err := s.client.GetUserInfo(ctx, token, s.config.GetRealm())
//	if err != nil {
//		logrus.Error("请求API出错：", err)
//		return
//	}
//	return domain.ToUser(info), err
//
//}

/*
登录
*/
//func (s *KeyCloakService) Login(ctx context.Context, username string, pwd string) (token string, err error) {
//	//查询用户
//	userinfo, err := s.UserRepository.GetByUsername(ctx, username)
//	if err != nil {
//		logrus.Error("查询不到用户", err)
//		return
//	}
//	//根据用户获取密码
//	password, err := s.UserRepository.GetPasswordByUsername(ctx, userinfo.Username)
//	if err != nil {
//		logrus.Error("获取密码识别", err)
//		return
//	}
//	//验证密码
//	if err := s.ValidatePassword(password, pwd); err != nil {
//		// 返回错误
//		logrus.Error("密码验证错误", err)
//		return
//	}
//	// 4. 生成JWT token
//	//s.JwtService.GenerateToken(username)
//	return "", err
//}

/*
本地登录
*/
func (s *KeyCloakService) LocalLogin(ctx context.Context, login *domain.LoginRequest) (domain.LoginResponse, error) {
	userinfo, err := s.UserRepository.GetByUsername(ctx, login.Username)
	if err != nil {
		logrus.Error("查询不到用户", err)
		return domain.LoginResponse{nil, ""}, err
	}
	//根据用户获取密码
	password, err := s.UserRepository.GetPasswordByUsername(ctx, userinfo.Username)
	if err != nil {
		logrus.Error("获取密码失败", err)
		return domain.LoginResponse{nil, ""}, err
	}
	//验证密码
	if err := s.ValidatePassword(password, login.Password); err != nil {
		logrus.Error("密码验证错误", err)
		return domain.LoginResponse{nil, ""}, err
	}
	//创建jwtToken
	claims := domain.TokenClaims{
		ID:       login.Username,
		Username: login.Username,
		Exp:      time.Now().Add(24 * time.Hour).Unix(),
	}
	token, _ := GenerateToken(&claims)
	return domain.LoginResponse{
		User:  userinfo,
		Token: token,
	}, nil
}

/*
 获取用户明细，不存在则新增
*/
//func (s *KeyCloakService) GetAndCreate(ctx context.Context, token string) (loginUser domain.LoginUserVO, err error) {
//
//}

func (s *KeyCloakService) GetaCodeURL() {
}

/*
密码验证
*/
func (s *KeyCloakService) ValidatePassword(encryptedPassword string, password string) error {
	// 1. 把encryptedPassword从字符串转换成[]byte
	encryptedPasswordBytes := []byte(encryptedPassword)

	// 2. 使用bcrypt进行密码匹配
	err := bcrypt.CompareHashAndPassword(encryptedPasswordBytes, []byte(password))

	// 3. 返回错误或nil
	if err != nil {
		logrus.Error("密码不正确", err)
		return err
	}
	return nil
}
