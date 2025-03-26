package domain

// casbin配置
type casbin struct {
	// 模型文件路径
	Model string `toml:"model"`
	// 策略适配器类型：file, mysql
	Adapter string `toml:"adapter"`
	// 策略文件路径（当adapter=file时使用）
	Policy string `toml:"policy"`
	// 是否自动加载策略
	AutoLoad bool `toml:"auto_load"`
	// 自动加载间隔时间（秒）
	AutoLoadInterval int `toml:"auto_load_interval"`
}

// GetModel 获取模型文件路径
func (c *casbin) GetModel() string {
	return c.Model
}

// GetAdapter 获取适配器类型
func (c *casbin) GetAdapter() string {
	return c.Adapter
}

// GetPolicy 获取策略文件路径
func (c *casbin) GetPolicy() string {
	return c.Policy
}

// IsAutoLoad 是否自动加载
func (c *casbin) IsAutoLoad() bool {
	return c.AutoLoad
}

// GetAutoLoadInterval 获取自动加载间隔时间
func (c *casbin) GetAutoLoadInterval() int {
	return c.AutoLoadInterval
}
