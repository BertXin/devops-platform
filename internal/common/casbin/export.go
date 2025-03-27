package casbin

import (
	"devops-platform/internal/common/casbin/internal/domain"
	"devops-platform/internal/common/casbin/internal/service"
)

// 注册bean
const (
	BeanEnforcer = domain.BeanEnforcer
)

// Enforce 执行权限验证
func Enforce(rvals ...interface{}) (bool, error) {
	return service.Enforce(rvals...)
}

// LoadPolicy 加载策略
func LoadPolicy() error {
	return service.LoadPolicy()
}

// SavePolicy 保存策略
func SavePolicy() error {
	return service.SavePolicy()
}

// AddPolicy 添加策略
func AddPolicy(params ...interface{}) (bool, error) {
	return service.AddPolicy(params...)
}

// RemovePolicy 删除策略
func RemovePolicy(params ...interface{}) (bool, error) {
	return service.RemovePolicy(params...)
}

// RemoveFilteredPolicy 按条件删除策略
func RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	return service.RemoveFilteredPolicy(fieldIndex, fieldValues...)
}

// AddRoleForUser 为用户添加角色
func AddRoleForUser(user string, role string) (bool, error) {
	return service.AddRoleForUser(user, role)
}

// DeleteRoleForUser 删除用户角色
func DeleteRoleForUser(user string, role string) (bool, error) {
	return service.DeleteRoleForUser(user, role)
}

// DeleteRolesForUser 删除用户所有角色
func DeleteRolesForUser(user string) (bool, error) {
	return service.DeleteRolesForUser(user)
}

// GetRolesForUser 获取用户所有角色
func GetRolesForUser(name string) ([]string, error) {
	return service.GetRolesForUser(name)
}

// HasRoleForUser 判断用户是否拥有指定角色
func HasRoleForUser(name string, role string) (bool, error) {
	return service.HasRoleForUser(name, role)
}

// GetUsersForRole 获取拥有指定角色的所有用户
func GetUsersForRole(name string) ([]string, error) {
	return service.GetUsersForRole(name)
}

// GetAllRoles 获取所有角色
func GetAllRoles() ([]string, error) {
	return service.GetAllRoles()
}

// GetAllObjects 获取所有资源
func GetAllObjects() ([]string, error) {
	return service.GetAllObjects()
}

// GetAllSubjects 获取所有主体
func GetAllSubjects() ([]string, error) {
	return service.GetAllSubjects()
}
