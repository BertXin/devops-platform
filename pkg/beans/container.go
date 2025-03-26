package beans

// Container 定义依赖容器接口
type Container interface {
	// GetBean 根据名称获取bean
	GetBean(name string) interface{}
	// GetDB 获取数据库连接
	GetDB() interface{}
}

// 实现Container接口的工厂方法
type containerImpl struct{}

// 创建容器实例
func NewContainer() Container {
	return &containerImpl{}
}

// GetBean 根据名称获取bean
func (c *containerImpl) GetBean(name string) interface{} {
	return factory[name]
}

// GetDB 获取数据库连接
func (c *containerImpl) GetDB() interface{} {
	return factory["DB"]
}
