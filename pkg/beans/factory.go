package beans

import (
	"errors"
	"github.com/sirupsen/logrus"
)

const (
	BeanStopWaiter = "generalStopWaiter"
)

type injectable interface {
	Inject(func(string) interface{})
}

type preInjectable interface {
	PreInject(func(string) interface{})
}

var factory = make(map[string]interface{})

var starts postStarts
var stops preStops

func Register(name string, bean interface{}) {
	var err error
	if name == "" {
		err = errors.New("bean名称不能为空")
	} else if bean == nil {
		err = errors.New("bean不能为空")
	} else if _, ok := factory[name]; ok {
		err = errors.New("bean名称重复")
	} else {
		factory[name] = bean
	}

	if err != nil {
		logrus.WithError(err).Panicf("注册bean[%s]失败", name)
	}
	if p, ok := bean.(postStart); ok {
		starts = append(starts, p)
	}
	if p, ok := bean.(preStop); ok {
		stops = append(stops, p)
	}
}

func get(name string) interface{} {
	return factory[name]
}

func Start() {
	/*
		先期注入
	*/
	preInject()
	/*
		进行bean注入
	*/
	inject()
	/*
		延迟启动项
	*/
	starts.start()

	if hasStopWait() {
		/**
		 * 等待停止信号
		 */
		stopWait()
		/**
		 * 进行清理
		 */
		stops.stop()

		logrus.Info("应用关闭成功")
	}
}

func inject() {

	for _, value := range factory {
		if f, ok := value.(func(func(string) interface{})); ok {
			f(get)
		}
	}
	for _, value := range factory {
		if bean, ok := value.(injectable); ok {
			bean.Inject(get)
		}
		injectBean(value)
	}
}

func preInject() {
	for _, value := range factory {
		if bean, ok := value.(preInjectable); ok {
			bean.PreInject(get)
		}
	}
}

func stopWait() {
	if generalStopWaiter, ok := factory[BeanStopWaiter].(func()); ok {
		generalStopWaiter()
	}
}

func hasStopWait() bool {
	_, ok := factory[BeanStopWaiter].(func())
	return ok
}

func RegisterStopWaiter(generalStopWaiter func()) {
	Register(BeanStopWaiter, generalStopWaiter)
}

/*
1. Bean注册
- 提供Register方法注册Bean,可以注册名为name的Bean对象
- 注册时会做一些校验,比如name不能为空,bean不能为空等
- 注册时会保存Bean到factory map中

2. Bean获取
- 提供get方法从factory中获取已注册的Bean

3. 依赖注入
- 提供inject方法,遍历所有Bean,如果实现了injectable接口,会调用其Inject方法完成依赖注入
- Inject方法会传入get方法,可以根据名称获取依赖的Bean

4. 生命周期管理
- 提供Start方法,用于启动所有Bean
- 会首先执行preInject做先期依赖注入
- 然后执行inject做一般依赖注入
- 最后启动实现了postStart接口的延迟启动Bean
- 提供Stop方法,用于关闭所有Bean
- 会执行实现了preStop接口的Bean的预处理动作
- 然后等待全局的停止信号
- 最后执行实现了preStop接口的Bean的关闭动作

5. 停止等待
- 提供RegisterStopWaiter方法注册停止等待处理函数
- Start时会等待这个函数的执行
整体而言,这个Bean容器实现了依赖注入、生命周期管理、停止处理等服务,可以用于管理各种Bean的创建、依赖注入和销毁。
*/
