package domain

const (
	// 模块Bean名称常量
	BeanAuthorizationService    = "AuthorizationService"
	BeanRoleService             = "RoleService"
	BeanPermissionService       = "PermissionService"
	BeanRepository              = "AuthorizationRepository"
	BeanAuthorizationController = "AuthorizationController"

	// 权限类型常量
	PermTypeMenu   = "menu"   // 菜单权限
	PermTypeApi    = "api"    // API权限
	PermTypeButton = "button" // 按钮权限

	// Casbin策略类型
	CasbinPTypePolicy = "p" // 策略规则
	CasbinPTypeRole   = "g" // 角色继承规则

	// Casbin模式
	CasbinUserPrefix = "u_" // 用户前缀
	CasbinRolePrefix = "r_" // 角色前缀
)

// Casbin规则模型
// p, role, resource, action, effect
// g, user, role
