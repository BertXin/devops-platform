package beans

import (
"github.com/sirupsen/logrus"
"reflect"
)

func injectBean(bean interface{}) {
	/*
	 * 反射bean的类型
	 */
	beanType := reflect.TypeOf(bean)

	/*
	 * 如果bean类型是指针，则取指针指向的实际类型
	 */
	if beanType.Kind() == reflect.Ptr {
		beanType = beanType.Elem()
	}
	/*
	 * 只对struct进行反射注入
	 */
	if beanType.Kind() != reflect.Struct {
		return
	}
	/*
	 * 反射bean的值
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
		/*
		 * 如果bean的field有inject注解
		 */
		if tag, hasTag := field.Tag.Lookup("inject"); hasTag {
			/*
			 * 从beanFactory中依据inject的tag值反射获取对应的值
			 */
			value := reflect.ValueOf(factory[tag])
			if !value.IsValid() {
				logrus.Panicf("初始化时获取[%s]为空", tag)
				continue
			}
			valueType := reflect.TypeOf(factory[tag])
			fieldValue := beanValue.FieldByName(field.Name)
			/*
			 * 判断依赖注入类型是否一致
			 */
			if valueType.AssignableTo(field.Type) {
				fieldValue.Set(value)
			} else {
				logrus.Panicf("类型不匹配! 不能把bean[%s]注入到[%s/%s.%s]中", tag, beanType.PkgPath(), beanType.Name(), field.Name)
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
