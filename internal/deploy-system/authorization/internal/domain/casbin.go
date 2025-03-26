package domain

// CasbinEnforcer 定义Casbin Enforcer接口
type CasbinEnforcer interface {
	// 主要权限验证方法
	Enforce(rvals ...interface{}) (bool, error)

	// 策略管理方法
	LoadPolicy() error
	SavePolicy() error
	AddPolicy(params ...interface{}) (bool, error)
	RemovePolicy(params ...interface{}) (bool, error)
	RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error)

	// 角色管理方法
	AddRoleForUser(user string, role string) (bool, error)
	DeleteRoleForUser(user string, role string) (bool, error)
	DeleteRolesForUser(user string) (bool, error)
	GetRolesForUser(name string) ([]string, error)
	GetUsersForRole(name string) ([]string, error)
	HasRoleForUser(name string, role string) (bool, error)

	// 信息查询方法
	GetAllRoles() ([]string, error)
	GetAllObjects() ([]string, error)
	GetAllSubjects() ([]string, error)
}
