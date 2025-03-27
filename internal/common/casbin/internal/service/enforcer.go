package service

import (
	"devops-platform/internal/common/casbin/internal/domain"
	"devops-platform/internal/common/config"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 定义casbin配置接口
type casbinConfig interface {
	GetModel() string
	GetAdapter() string
	GetPolicy() string
	IsAutoLoad() bool
	GetAutoLoadInterval() int
}

// CasbinEnforcer 实现
type CasbinEnforcer struct {
	enforcer *casbin.Enforcer
	config   casbinConfig
}

// 全局enforcer实例
var globalEnforcer domain.CasbinEnforcer

// Enforce 全局执行权限验证
func Enforce(rvals ...interface{}) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.Enforce(rvals...)
}

// LoadPolicy 全局加载策略
func LoadPolicy() error {
	if globalEnforcer == nil {
		return nil
	}
	return globalEnforcer.LoadPolicy()
}

// SavePolicy 全局保存策略
func SavePolicy() error {
	if globalEnforcer == nil {
		return nil
	}
	return globalEnforcer.SavePolicy()
}

// AddPolicy 全局添加策略
func AddPolicy(params ...interface{}) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.AddPolicy(params...)
}

// RemovePolicy 全局删除策略
func RemovePolicy(params ...interface{}) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.RemovePolicy(params...)
}

// RemoveFilteredPolicy 全局按条件删除策略
func RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.RemoveFilteredPolicy(fieldIndex, fieldValues...)
}

// AddRoleForUser 全局为用户添加角色
func AddRoleForUser(user string, role string) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.AddRoleForUser(user, role)
}

// DeleteRoleForUser 全局删除用户角色
func DeleteRoleForUser(user string, role string) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.DeleteRoleForUser(user, role)
}

// DeleteRolesForUser 全局删除用户所有角色
func DeleteRolesForUser(user string) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.DeleteRolesForUser(user)
}

// GetRolesForUser 全局获取用户所有角色
func GetRolesForUser(name string) ([]string, error) {
	if globalEnforcer == nil {
		return nil, nil
	}
	return globalEnforcer.GetRolesForUser(name)
}

// GetUsersForRole 全局获取拥有指定角色的所有用户
func GetUsersForRole(name string) ([]string, error) {
	if globalEnforcer == nil {
		return nil, nil
	}
	return globalEnforcer.GetUsersForRole(name)
}

// HasRoleForUser 全局判断用户是否拥有指定角色
func HasRoleForUser(name string, role string) (bool, error) {
	if globalEnforcer == nil {
		return false, nil
	}
	return globalEnforcer.HasRoleForUser(name, role)
}

// GetAllRoles 全局获取所有角色
func GetAllRoles() ([]string, error) {
	if globalEnforcer == nil {
		return nil, nil
	}
	return globalEnforcer.GetAllRoles()
}

// GetAllObjects 全局获取所有资源
func GetAllObjects() ([]string, error) {
	if globalEnforcer == nil {
		return nil, nil
	}
	return globalEnforcer.GetAllObjects()
}

// GetAllSubjects 全局获取所有主体
func GetAllSubjects() ([]string, error) {
	if globalEnforcer == nil {
		return nil, nil
	}
	return globalEnforcer.GetAllSubjects()
}

// 预注入依赖
func (e *CasbinEnforcer) PreInject(getBean func(string) interface{}) {
	// 获取配置
	beanName := config.BeanCasbin
	cfg, ok := getBean(beanName).(casbinConfig)
	if !ok {
		logrus.Panicf("初始化时获取[%s]失败", beanName)
		return
	}
	e.config = cfg

	// 加载模型
	var m model.Model
	var err error
	if m, err = model.NewModelFromFile(cfg.GetModel()); err != nil {
		logrus.Panicf("加载Casbin模型文件错误: %s", err.Error())
		return
	}

	// 创建适配器
	var adapter interface{}
	switch cfg.GetAdapter() {
	case "file":
		adapter = fileadapter.NewAdapter(cfg.GetPolicy())
	case "mysql":
		// 获取数据库连接
		db := getBean("gorm-db")
		if db == nil {
			logrus.Panic("获取数据库连接失败")
			return
		}
		if adapter, err = gormadapter.NewAdapterByDB(db.(*gorm.DB)); err != nil {
			logrus.Panicf("创建Casbin GORM适配器错误: %s", err.Error())
			return
		}
	default:
		logrus.Panicf("不支持的Casbin适配器类型: %s", cfg.GetAdapter())
		return
	}

	// 创建enforcer
	if e.enforcer, err = casbin.NewEnforcer(m, adapter); err != nil {
		logrus.Panicf("创建Casbin Enforcer错误: %s", err.Error())
		return
	}

	// 启动自动加载
	if cfg.IsAutoLoad() {
		e.enableAutoLoad(time.Duration(cfg.GetAutoLoadInterval()) * time.Second)
	}

	// 设置全局实例
	globalEnforcer = e
}

// enableAutoLoad 启用自动加载策略
func (e *CasbinEnforcer) enableAutoLoad(d time.Duration) {
	e.enforcer.EnableAutoSave(true)

	// 简单定时加载策略
	// 注意：这里我们使用非阻塞的goroutine来定期重新加载策略
	go func() {
		ticker := time.NewTicker(d)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				logrus.Debug("正在自动重新加载Casbin策略...")
				err := e.enforcer.LoadPolicy()
				if err != nil {
					logrus.WithError(err).Error("自动加载策略失败")
				}
			}
		}
	}()
}

// Enforce 执行权限验证
func (e *CasbinEnforcer) Enforce(rvals ...interface{}) (bool, error) {
	return e.enforcer.Enforce(rvals...)
}

// LoadPolicy 加载策略
func (e *CasbinEnforcer) LoadPolicy() error {
	return e.enforcer.LoadPolicy()
}

// SavePolicy 保存策略
func (e *CasbinEnforcer) SavePolicy() error {
	return e.enforcer.SavePolicy()
}

// AddPolicy 添加策略
func (e *CasbinEnforcer) AddPolicy(params ...interface{}) (bool, error) {
	return e.enforcer.AddPolicy(params...)
}

// RemovePolicy 删除策略
func (e *CasbinEnforcer) RemovePolicy(params ...interface{}) (bool, error) {
	return e.enforcer.RemovePolicy(params...)
}

// RemoveFilteredPolicy 按条件删除策略
func (e *CasbinEnforcer) RemoveFilteredPolicy(fieldIndex int, fieldValues ...string) (bool, error) {
	return e.enforcer.RemoveFilteredPolicy(fieldIndex, fieldValues...)
}

// AddRoleForUser 为用户添加角色
func (e *CasbinEnforcer) AddRoleForUser(user string, role string) (bool, error) {
	return e.enforcer.AddRoleForUser(user, role)
}

// DeleteRoleForUser 删除用户角色
func (e *CasbinEnforcer) DeleteRoleForUser(user string, role string) (bool, error) {
	return e.enforcer.DeleteRoleForUser(user, role)
}

// DeleteRolesForUser 删除用户所有角色
func (e *CasbinEnforcer) DeleteRolesForUser(user string) (bool, error) {
	return e.enforcer.DeleteRolesForUser(user)
}

// GetRolesForUser 获取用户所有角色
func (e *CasbinEnforcer) GetRolesForUser(name string) ([]string, error) {
	return e.enforcer.GetRolesForUser(name)
}

// GetUsersForRole 获取拥有指定角色的所有用户
func (e *CasbinEnforcer) GetUsersForRole(name string) ([]string, error) {
	return e.enforcer.GetUsersForRole(name)
}

// HasRoleForUser 判断用户是否拥有指定角色
func (e *CasbinEnforcer) HasRoleForUser(name string, role string) (bool, error) {
	return e.enforcer.HasRoleForUser(name, role)
}

// GetAllRoles 获取所有角色
func (e *CasbinEnforcer) GetAllRoles() ([]string, error) {
	return e.enforcer.GetAllRoles()
}

// GetAllObjects 获取所有资源
func (e *CasbinEnforcer) GetAllObjects() ([]string, error) {
	return e.enforcer.GetAllObjects()
}

// GetAllSubjects 获取所有主体
func (e *CasbinEnforcer) GetAllSubjects() ([]string, error) {
	return e.enforcer.GetAllSubjects()
}

// StartAutoLoadPolicy 启动自动加载策略
func (e *CasbinEnforcer) StartAutoLoadPolicy(d interface{}) {
	if duration, ok := d.(time.Duration); ok {
		e.enableAutoLoad(duration)
	}
}
