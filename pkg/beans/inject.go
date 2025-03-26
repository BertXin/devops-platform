package beans

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

// Inject 依赖注入接口
type Inject interface {
	// Inject 依赖注入方法
	Inject(container Container)
}

func injectBean(bean interface{}) {
	// 优先处理实现了Inject接口的bean，使用自定义的注入逻辑
	if inject, ok := bean.(Inject); ok {
		// 创建容器
		container := NewContainer()
		// 调用Inject方法进行依赖注入
		inject.Inject(container)
		return
	}

	/*
	 * 只给struct类型注入，减少不必要的反射，增加性能
	 */
	beanType := reflect.TypeOf(bean)
	if beanType == nil || (beanType.Kind() != reflect.Struct && beanType.Kind() != reflect.Ptr) {
		return
	}
	/*
	 * 如果是指针，则获取指针指向的值的Type
	 */
	if beanType.Kind() == reflect.Ptr {
		beanType = beanType.Elem()
	}
	/*
	 * 如果bean的Type不是struct类型
	 */
	if beanType.Kind() != reflect.Struct {
		return
	}
	/*
	 * 通过Type获取对应的Value
	 */
	beanValue := reflect.ValueOf(bean)
	/*
	 * 如果bean值是指针，则取实际值（指针无法获取field）
	 */
	if beanValue.Kind() == reflect.Ptr {
		beanValue = beanValue.Elem()
	}
	/*
	 * 1. 遍历bean的field
	 * 2. 如果bean的field有inject注解，则依据名称进行注入
	 */
	for i, size := 0, beanType.NumField(); i < size; i++ {
		/*
		 * bean的field的类型
		 */
		field := beanType.Field(i)

		// 检查字段是否导出，如果字段首字母小写（非导出），跳过注入
		if field.PkgPath != "" {
			continue
		}

		/*
		 * 如果bean的field有inject注解
		 */
		if tag, hasTag := field.Tag.Lookup("inject"); hasTag {
			/*
			 * 从beanFactory中依据inject的tag值反射获取对应的值
			 */
			injectedValue, exists := factory[tag]
			if !exists || injectedValue == nil {
				// 使用Error而不是Panic，允许程序继续运行
				logrus.Errorf("初始化时获取[%s]为空 (bean: %s/%s.%s)", tag, beanType.PkgPath(), beanType.Name(), field.Name)
				continue
			}

			value := reflect.ValueOf(injectedValue)
			valueType := reflect.TypeOf(injectedValue)
			fieldValue := beanValue.FieldByName(field.Name)

			/*
			 * 判断依赖注入类型是否一致
			 */
			if valueType.AssignableTo(field.Type) {
				// 确保字段可设置
				if fieldValue.CanSet() {
					fieldValue.Set(value)
				} else {
					logrus.Errorf("字段[%s]不可设置，可能是非导出字段", field.Name)
				}
			} else {
				// 使用Error而不是Panic
				logrus.Errorf("类型不匹配! 不能把bean[%s]注入到[%s/%s.%s]中", tag, beanType.PkgPath(), beanType.Name(), field.Name)
			}
		}
	}
}

/*
1. 接收一个bean对象
2. 使用reflect获取bean的类型(Type)和值(Value)
3. 遍历bean的所有字段(Field)
4. 检查每个字段是否有"inject"标记
5. 如果有,则根据"inject"的名称,使用反射从factory中获取对应的对象
6. 检查这个对象的类型是否匹配字段类型
7. 如果匹配,则使用反射将对象设置到字段上
关键点:
1. 使用反射提供类型安全的依赖注入
2. 支持针对指针和非指针的bean对象注入
3. 只对struct类型注入,减少不必要的反射
4. 校验类型匹配,防止错误注入
5. 使用factory管理依赖对象的创建,实现解耦
6. panic失败,方便测试调试
整体上是一个典型的依赖注入实现,可以自动解析bean中的依赖关系,并从factory填充这些依赖。
*/
