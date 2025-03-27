# 授权模块 API 测试

本目录包含对授权模块 (`authorization`) API 的测试JSON文件，用于Postman等工具进行接口测试。

## 测试文件说明

### 角色管理

- `auth_create_role.json` - 创建角色接口请求体
- `auth_update_role.json` - 更新角色接口请求体
- `auth_query_role.json` - 角色查询接口参数
- `auth_assign_role_permissions.json` - 为角色分配权限接口请求体

### 权限管理

- `auth_create_permission.json` - 创建菜单权限接口请求体
- `auth_create_api_permission.json` - 创建API权限接口请求体
- `auth_create_button_permission.json` - 创建按钮权限接口请求体
- `auth_update_permission.json` - 更新权限接口请求体
- `auth_query_permission.json` - 权限查询接口参数

### 用户授权

- `auth_assign_user_roles.json` - 为用户分配角色接口请求体

## API接口列表

| 方法 | 路径 | 描述 | 对应测试文件 |
|------|------|------|------------|
| POST | `/api/v1/authorization/roles` | 创建角色 | auth_create_role.json |
| GET | `/api/v1/authorization/roles` | 获取角色列表 | auth_query_role.json (参数) |
| PUT | `/api/v1/authorization/roles/:id` | 更新角色 | auth_update_role.json |
| DELETE | `/api/v1/authorization/roles/:id` | 删除角色 | - |
| GET | `/api/v1/authorization/roles/:id` | 获取角色详情 | - |
| POST | `/api/v1/authorization/roles/:id/permissions` | 为角色分配权限 | auth_assign_role_permissions.json |
| GET | `/api/v1/authorization/roles/:id/permissions` | 获取角色权限 | - |
| POST | `/api/v1/authorization/permissions` | 创建权限 | auth_create_permission.json |
| GET | `/api/v1/authorization/permissions` | 获取权限列表 | auth_query_permission.json (参数) |
| GET | `/api/v1/authorization/permissions/tree` | 获取权限树结构 | - |
| PUT | `/api/v1/authorization/permissions/detail/:id` | 更新权限 | auth_update_permission.json |
| DELETE | `/api/v1/authorization/permissions/detail/:id` | 删除权限 | - |
| GET | `/api/v1/authorization/permissions/detail/:id` | 获取权限详情 | - |
| GET | `/api/v1/authorization/users/:user_id/roles` | 获取用户角色 | - |
| POST | `/api/v1/authorization/users/:user_id/roles` | 为用户分配角色 | auth_assign_user_roles.json |
| GET | `/api/v1/authorization/users/:user_id/permissions` | 获取用户权限 | - |
| DELETE | `/api/v1/authorization/users/:user_id/roles/:role_id` | 移除用户角色 | - |
| GET | `/api/v1/authorization/menus` | 获取用户菜单 | - |
| GET | `/api/v1/authorization/check/:permission` | 检查权限 | - |

## 使用方法

**注意**：授权模块的所有接口都需要认证，请先通过认证模块登录获取token。

### 创建角色接口

**请求路径**：`POST /api/v1/authorization/roles`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "name": "管理员角色",
  "code": "admin",
  "description": "系统管理员角色，拥有所有权限",
  "sort_order": 1
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "message": "success"
}
```

### 创建权限接口

**请求路径**：`POST /api/v1/authorization/permissions`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "parent_id": 0,
  "name": "系统管理",
  "type": "menu",
  "path": "/system",
  "icon": "setting",
  "component": "Layout",
  "permission": "system",
  "status": 1,
  "hidden": false,
  "sort_order": 1
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": {
    "id": 1
  },
  "message": "success"
}
```

### 为角色分配权限接口

**请求路径**：`POST /api/v1/authorization/roles/1/permissions`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "permission_ids": [1, 2, 3, 4, 5]
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 为用户分配角色接口

**请求路径**：`POST /api/v1/authorization/users/1/roles`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "role_ids": [1, 2, 3]
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
``` 