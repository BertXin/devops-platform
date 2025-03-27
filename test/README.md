# DevOps平台API测试文件

本目录包含对平台各模块API的测试JSON文件，用于Postman等工具进行接口测试。

## 目录结构

- `auth/` - 认证模块API测试
- `authorization/` - 授权模块API测试
- `organization/` - 组织模块API测试

## 快速开始

### 使用Postman测试

1. 导入各模块的JSON文件到Postman
2. 设置环境变量:
   - `baseUrl`: 例如 `http://localhost:8080`
   - `token`: 有效的JWT令牌（通过登录接口获取）

### 接口认证

- `auth`模块中的登录、注册、健康检查接口不需要认证
- 其他所有接口都需要在请求头中添加认证信息：
  ```
  Authorization: Bearer {{token}}
  ```

## 模块说明

### 认证模块 (auth)

包含用户认证相关的接口测试文件：

- `auth_login.json` - 用户登录
- `auth_register.json` - 用户注册
- `auth_change_password.json` - 修改密码

### 授权模块 (authorization)

包含权限管理相关的接口测试文件：

- `auth_create_role.json` - 创建角色
- `auth_update_role.json` - 更新角色
- `auth_query_role.json` - 查询角色参数
- `auth_assign_role_permissions.json` - 为角色分配权限
- `auth_assign_user_roles.json` - 为用户分配角色
- `auth_create_permission.json` - 创建菜单权限
- `auth_create_api_permission.json` - 创建API权限
- `auth_create_button_permission.json` - 创建按钮权限
- `auth_update_permission.json` - 更新权限
- `auth_query_permission.json` - 查询权限参数

### 组织模块 (organization)

包含组织管理相关的接口测试文件：

- `org_create_department.json` - 创建部门
- `org_create_sub_department.json` - 创建子部门
- `org_update_department.json` - 更新部门
- `org_query_department.json` - 查询部门参数
- `org_query_user.json` - 查询用户参数
- `org_assign_department.json` - 为用户分配部门

## 请求路径

### 认证模块

- 登录: `/auth/login` (POST)
- 注册: `/auth/register` (POST)
- 获取用户信息: `/auth/me` (GET)
- 修改密码: `/auth/change-password` (POST)
- 登出: `/auth/logout` (POST)

### 授权模块

- 角色管理: `/api/v1/authorization/roles` (GET, POST)
- 角色详情: `/api/v1/authorization/roles/:id` (GET, PUT, DELETE)
- 角色权限: `/api/v1/authorization/roles/:id/permissions` (GET, POST)
- 权限管理: `/api/v1/authorization/permissions` (GET, POST)
- 权限树: `/api/v1/authorization/permissions/tree` (GET)
- 权限详情: `/api/v1/authorization/permissions/detail/:id` (GET, PUT, DELETE)

### 组织模块

- 部门管理: `/api/v1/organization/departments` (GET, POST)
- 部门树: `/api/v1/organization/departments/tree` (GET)
- 用户部门: `/api/v1/organization/departments/user` (GET)
- 部门详情: `/api/v1/organization/departments/detail/:id` (GET, PUT, DELETE)
- 部门用户: `/api/v1/organization/departments/detail/:id/users` (GET)
- 分配部门: `/api/v1/organization/departments/users/:userId/departments/:departmentId` (POST, DELETE)

## 响应格式

所有API遵循统一的响应格式：

```json
// 成功响应
{
  "code": 200,
  "data": {},
  "message": "success"
}

// 错误响应
{
  "code": 400,
  "error": "Bad Request",
  "message": "参数错误",
  "request_id": "xyz123"
}
``` 