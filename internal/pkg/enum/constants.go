package enum

import "strconv"

type BuildMode int

const (
	BuildModeFull BuildMode = 0 // 全环境构建
	BuildModeSub  BuildMode = 1 // 分环境构建
)

type DeployMode int

const (
	DeployModeDeployment DeployMode = 0 // 部署deployment
	DeployModeJob        DeployMode = 1 // 部署job
)

type BuildType int

const (
	BuildTypeCustom BuildType = 0 // 自定义分支
	BuildTypeFixed  BuildType = 1 // 固定分支
)

//构建方法
type BuildMethod int

const (
	BuildMethodNone       = 0 //不构建
	BuildMethodJava       = 1 //Java构建
	BuildMethodJavaApi    = 2 //JavaApi构建
	BuildMethodGolang     = 3 //golang构建
	BuildMethodYarn       = 4 //Yarn构建
	BuildMethodNpm        = 5 //Npm构建
	BuildMethodYarnHigh   = 6 //Yarn构建
	BuildMethodNpmHigh    = 7 //Npm构建
	BuildMethodAndroidSdk = 8 //安卓sdk构建
)

var buildMethod = map[BuildMethod]string{
	BuildMethodNone:       "不构建",       //不构建
	BuildMethodJava:       "Java构建",    //Java构建
	BuildMethodJavaApi:    "JavaApi构建", //JavaApi构建
	BuildMethodGolang:     "golang构建",  //golang构建
	BuildMethodYarn:       "Yarn构建",    //Yarn构建
	BuildMethodNpm:        "Npm构建",     //Npm构建
	BuildMethodYarnHigh:   "Yarn高版本构建", //Yarn构建
	BuildMethodNpmHigh:    "Npm高版本构建",  //Npm构建
	BuildMethodAndroidSdk: "安卓sdk构建",   //安卓sdk构建
}

// 人员状态
type Enable int

const (
	EnableStatus  = 1 //启用
	DisableStatus = 2 // 禁用
)

var enable = map[Enable]string{
	EnableStatus:  "启用",
	DisableStatus: "禁用",
}

func (m Enable) String() string {
	return enable[m]
}

func (i Enable) ValidStatus() bool {
	if i.String() == "" {
		return false
	}
	return true
}

func (m BuildMethod) String() string {
	return buildMethod[m]
}

func (m BuildMethod) Int() int {
	return int(m)
}

func (m BuildMethod) Value() string {
	return strconv.Itoa(int(m))
}
func (m BuildMethod) AllowRestart() bool {
	return m != BuildMethodNone && m != BuildMethodJavaApi
}

func (m BuildMethod) HasImageBuild() bool {
	return m != BuildMethodJavaApi && m != BuildMethodAndroidSdk
}
func (m BuildMethod) IsAndroidBuild() bool {
	return m == BuildMethodAndroidSdk
}
func (m BuildMethod) IsJava() bool {
	return m == BuildMethodJavaApi || m == BuildMethodJava
}
func (m BuildMethod) IsYarn() bool {
	return m == BuildMethodYarn
}
func (m BuildMethod) IsNpm() bool {
	return m == BuildMethodNpm
}

func (m BuildMethod) IsYarnHigh() bool {
	return m == BuildMethodYarnHigh
}
func (m BuildMethod) IsNpmHigh() bool {
	return m == BuildMethodNpmHigh
}

func (m BuildMethod) IsFrontEnd() bool {
	return m.IsYarn() && m.IsNpm() && m.IsYarnHigh() && m.IsNpmHigh()
}

// 产品线人员角色
type ProductMemberRole int

func (i ProductMemberRole) String() string {
	return productMemberRole[i]
}

func (i ProductMemberRole) Int() int {
	return int(i)
}

func (i ProductMemberRole) Value() string {
	return strconv.Itoa(int(i))
}

func (i ProductMemberRole) IsLeader() bool {

	return i == ProductMemberRoleLeader ||
		i == ProductMemberRoleTestLeader ||
		i == ProductMemberRoleDeployLeader
}

const (
	ProductMemberRoleLeader ProductMemberRole = iota + 1
	ProductMemberRoleTestLeader
	ProductMemberRoleDeployLeader
	ProductMemberRoleDeveloper
	ProductMemberRoleTester
	ProductMemberRoleDeployer
	ProductMemberRoleSystem
)

var productMemberRole = map[ProductMemberRole]string{
	ProductMemberRoleLeader:       "leader",
	ProductMemberRoleTestLeader:   "test_leader",
	ProductMemberRoleDeployLeader: "deploy_leader", // 发版佬
	ProductMemberRoleDeveloper:    "developer",
	ProductMemberRoleTester:       "tester",
	ProductMemberRoleDeployer:     "deployer", // 发版佬
	ProductMemberRoleSystem:       "system",
}

// 系统角色
type SysRole int

func (i SysRole) String() string {
	return sysRole[i]
}

func (i SysRole) Int() int {
	return int(i)
}

func (i SysRole) Value() string {
	return strconv.Itoa(int(i))
}

func (i SysRole) ValidRole() bool {
	if i.String() == "" {
		return false
	}
	return true
}

const (
	SysRoleGeneralUser SysRole = 0
	SysRoleAdminUser   SysRole = 1
	SysRoleVirtualUser SysRole = 2
)

var sysRole = map[SysRole]string{
	SysRoleGeneralUser: "普通用户",
	SysRoleAdminUser:   "管理员",
	SysRoleVirtualUser: "虚拟用户",
}

//操作
type BuildStatus int

func (s BuildStatus) String() string {
	return buildStatuses[s]
}

func (s BuildStatus) Int() int {
	return int(s)
}

func (s BuildStatus) Value() string {
	return strconv.Itoa(int(s))
}

const (
	BuildStatusCanceled = -2 //已取消
	BuildStatusFail     = -1 //构建失败

	BuildStatusCreated  = 0 //已创建，待构建
	BuildStatusBuilding = 1 //构建中
	BuildStatusBuilt    = 2 //构建结束
	BuildStatusSuccess  = 3 //构建成功

)

var buildStatuses = map[BuildStatus]string{
	BuildStatusCanceled: "已取消",  //已取消
	BuildStatusFail:     "构建失败", //构建失败
	BuildStatusCreated:  "待构建",  //已创建，待构建
	BuildStatusBuilding: "构建中",  //构建中
	BuildStatusBuilt:    "构建结束", //构建中
	BuildStatusSuccess:  "构建成功", //构建成功
}

//操作
type DeployStatus int

func (s DeployStatus) String() string {
	return deployStatuses[s]
}

func (s DeployStatus) Int() int {
	return int(s)
}

func (s DeployStatus) Value() string {
	return strconv.Itoa(int(s))
}

const (
	DeployStatusCanceled  = -2 //已取消
	DeployStatusFail      = -1 //部署失败
	DeployStatusCreated   = 0  //已创建，待部署
	DeployStatusDeploying = 1  //部署中
	DeployStatusSuccess   = 2  //部署成功

)

var deployStatuses = map[DeployStatus]string{
	DeployStatusCanceled:  "已取消",  //已取消
	DeployStatusFail:      "部署失败", //部署失败
	DeployStatusCreated:   "待部署",  //已创建，待部署
	DeployStatusDeploying: "部署中",  //部署中
	DeployStatusSuccess:   "部署成功", //部署成功
}

type ReviewMode int

const (
	ReviewModeNone        ReviewMode = -1 //不评审
	ReviewModeAlways      ReviewMode = 0  // 每次提测时都评审
	ReviewModeIfNotReview ReviewMode = 1  // 首次提测时评审
)
