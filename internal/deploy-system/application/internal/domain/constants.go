package domain

// 模块Bean名称常量
const (
	// BeanModuleName 模块名称
	BeanModuleName = "ApplicationService"
	// BeanService 应用管理服务Bean名称
	BeanAppService = "applicationService"
	// BeanRepository 仓储Bean名称
	BeanAppRepository = "ApplicationRepository"
	// BeanController 控制器Bean名称
	BeanController = "ApplicationController"
	// BeanDeployService 部署服务Bean名称
	BeanDeployService = "deployService"
	// BeanAppQuery 应用查询Bean名称
	BeanAppQuery = "appQuery"
)

// 应用状态常量
const (
	// AppStatusActive 应用状态-活跃
	AppStatusActive = "active"
	// AppStatusInactive 应用状态-不活跃
	AppStatusInactive = "inactive"
	// AppStatusDeleted 应用状态-已删除
	AppStatusDeleted = "deleted"
)

// 部署状态常量
const (
	// DeployStatusPending 部署状态-待处理
	DeployStatusPending = "pending"
	// DeployStatusRunning 部署状态-运行中
	DeployStatusRunning = "running"
	// DeployStatusSuccess 部署状态-成功
	DeployStatusSuccess = "success"
	// DeployStatusFailed 部署状态-失败
	DeployStatusFailed = "failed"
	// DeployStatusRollback 部署状态-回滚
	DeployStatusRollback = "rollback"
)

// 部署策略常量
const (
	// DeployStrategyRolling 部署策略-滚动更新
	DeployStrategyRolling = "rolling"
	// DeployStrategyBlueGreen 部署策略-蓝绿部署
	DeployStrategyBlueGreen = "blue-green"
	// DeployStrategyCanary 部署策略-金丝雀发布
	DeployStrategyCanary = "canary"
)
