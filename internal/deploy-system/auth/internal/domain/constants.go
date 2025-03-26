package domain

// 模块Bean名称常量
const (
	// BeanModuleName 模块名称
	BeanModuleName = "AuthService"
	// BeanService 认证服务Bean名称
	BeanService = "authService"
	// BeanRepository 仓储Bean名称
	BeanRepository = "AuthRepository"
	// BeanUserQuery 用户查询Bean名称
	BeanUserQuery = "userQuery"
	// BeanController 控制器Bean名称
	BeanController = "AuthController"
)

// 登录相关常量
const (
	// TokenPrefix 令牌前缀
	TokenPrefix = "Bearer "
	// LoginTypePassword 密码登录类型
	LoginTypePassword = "password"
	// LoginTypeLDAP LDAP登录类型
	LoginTypeLDAP = "ldap"
	// LoginTypeDingTalk 钉钉登录类型
	LoginTypeDingTalk = "dingtalk"
	// LoginTypeMFA MFA二次验证类型
	LoginTypeMFA = "mfa"
)
