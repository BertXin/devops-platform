# DevOps Platform API 接口文档

## 概述

本文档描述了 DevOps Platform 后端系统的所有 API 接口，包括认证、权限管理、应用管理、组织管理等模块。

## 基础信息

- **Base URL**: `http://localhost:8080`
- **API 版本**: `v1`
- **认证方式**: JWT Token (Bearer Token)
- **数据格式**: JSON

## 通用响应格式

### 成功响应
```json
{
  "code": 200,
  "data": {},
  "message": "success"
}
```

### 分页响应
```json
{
  "code": 200,
  "data": {
    "list": [],
    "total": 100,
    "page": 1,
    "size": 10
  },
  "message": "success"
}
```

### 错误响应
```json
{
  "code": 400,
  "error": "BadRequest",
  "message": "请求参数错误",
  "request_id": "uuid"
}
```

## 1. 认证模块 (Auth)

### 1.1 用户登录
- **URL**: `POST /api/v1/auth/login`
- **描述**: 用户登录获取访问令牌
- **认证**: 无需认证

**请求参数**:
```json
{
  "username": "admin",
  "password": "123456"
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expire_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

### 1.2 用户注册
- **URL**: `POST /api/v1/auth/register`
- **描述**: 注册新用户
- **认证**: 无需认证

**请求参数**:
```json
{
  "username": "newuser",
  "password": "123456",
  "name": "新用户",
  "mobile": "13800138000",
  "email": "user@example.com",
  "avatar": "http://example.com/avatar.jpg",
  "dept_id": 1,
  "role": 2
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "注册成功"
}
```

### 1.3 用户登出
- **URL**: `POST /api/v1/auth/logout`
- **描述**: 用户登出系统
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 1.4 获取当前用户信息
- **URL**: `GET /api/v1/auth/me`
- **描述**: 获取当前登录用户的详细信息
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "username": "admin",
    "nickname": "管理员",
    "email": "admin@example.com",
    "phone": "13800138000",
    "avatar": "http://example.com/avatar.jpg",
    "dept_id": 1,
    "dept_name": "技术部",
    "role_id": 1,
    "status": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

### 1.5 修改密码
- **URL**: `POST /api/v1/auth/change-password`
- **描述**: 修改当前用户密码
- **认证**: 需要认证

**请求参数**:
```json
{
  "old_password": "oldpass123",
  "new_password": "newpass123",
  "confirm_password": "newpass123"
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

## 2. 应用管理模块 (Application)

### 2.1 查询应用列表
- **URL**: `GET /api/v1/apps`
- **描述**: 分页查询应用列表
- **认证**: 需要认证

**查询参数**:
- `name` (string, optional): 应用名称
- `status` (string, optional): 应用状态
- `group_id` (int, optional): 分组ID
- `page` (int, optional): 页码，默认1
- `size` (int, optional): 每页大小，默认10

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "name": "demo-app",
        "description": "演示应用",
        "creator": 1,
        "status": "active",
        "group_ids": [1, 2],
        "env_count": 3,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "size": 10
  },
  "message": "success"
}
```

### 2.2 创建应用
- **URL**: `POST /api/v1/apps`
- **描述**: 创建新应用
- **认证**: 需要认证

**请求参数**:
```json
{
  "name": "new-app",
  "description": "新应用描述",
  "creator": 1
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "创建成功"
}
```

### 2.3 获取应用详情
- **URL**: `GET /api/v1/apps/{id}`
- **描述**: 根据ID获取应用详情
- **认证**: 需要认证

**路径参数**:
- `id` (int): 应用ID

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "name": "demo-app",
    "description": "演示应用",
    "creator": 1,
    "status": "active",
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

### 2.4 更新应用
- **URL**: `PUT /api/v1/apps/{id}`
- **描述**: 更新应用信息
- **认证**: 需要认证

**路径参数**:
- `id` (int): 应用ID

**请求参数**:
```json
{
  "name": "updated-app",
  "description": "更新后的描述"
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 2.5 删除应用
- **URL**: `DELETE /api/v1/apps/{id}`
- **描述**: 删除应用
- **认证**: 需要认证

**路径参数**:
- `id` (int): 应用ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 2.6 配置应用HPA
- **URL**: `POST /api/v1/apps/{id}/hpa`
- **描述**: 配置应用自动伸缩
- **认证**: 需要认证

**路径参数**:
- `id` (int): 应用ID

**请求参数**:
```json
{
  "min_replicas": 1,
  "max_replicas": 10,
  "target_cpu_utilization": 80,
  "target_memory_utilization": 80
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 2.7 获取应用HPA配置
- **URL**: `GET /api/v1/apps/{id}/hpa`
- **描述**: 获取应用HPA配置
- **认证**: 需要认证

**路径参数**:
- `id` (int): 应用ID

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "app_id": 1,
    "min_replicas": 1,
    "max_replicas": 10,
    "target_cpu_utilization": 80,
    "target_memory_utilization": 80,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

### 2.8 查询应用分组列表
- **URL**: `GET /api/v1/app-groups`
- **描述**: 获取所有应用分组
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "前端应用",
      "description": "前端应用分组",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 2.9 创建应用分组
- **URL**: `POST /api/v1/app-groups`
- **描述**: 创建新的应用分组
- **认证**: 需要认证

**请求参数**:
```json
{
  "name": "新分组",
  "description": "分组描述"
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "创建成功"
}
```

### 2.10 添加应用到分组
- **URL**: `POST /api/v1/app-groups/apps`
- **描述**: 将应用添加到分组
- **认证**: 需要认证

**请求参数**:
```json
{
  "group_id": 1,
  "app_id": 1
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 2.11 查询环境列表
- **URL**: `GET /api/v1/envs`
- **描述**: 获取所有环境列表
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "开发环境",
      "cluster_id": 1,
      "namespace": "dev",
      "description": "开发环境",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 2.12 创建环境
- **URL**: `POST /api/v1/envs`
- **描述**: 创建新环境
- **认证**: 需要认证

**请求参数**:
```json
{
  "name": "测试环境",
  "cluster_id": 1,
  "namespace": "test",
  "description": "测试环境"
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "创建成功"
}
```

## 3. 权限管理模块 (Authorization)

### 3.1 获取用户菜单
- **URL**: `GET /api/v1/authorization/menus`
- **描述**: 获取当前用户的菜单权限
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "系统管理",
      "type": "menu",
      "path": "/system",
      "icon": "system",
      "component": "Layout",
      "permission": "system:view",
      "status": 1,
      "hidden": false,
      "sort_order": 1,
      "children": [
        {
          "id": 2,
          "parent_id": 1,
          "name": "用户管理",
          "type": "menu",
          "path": "/system/user",
          "icon": "user",
          "component": "system/user/index",
          "permission": "system:user:view",
          "status": 1,
          "hidden": false,
          "sort_order": 1
        }
      ]
    }
  ],
  "message": "success"
}
```

### 3.2 权限检查
- **URL**: `GET /api/v1/authorization/check/{permission}`
- **描述**: 检查当前用户是否有指定权限
- **认证**: 需要认证

**路径参数**:
- `permission` (string): 权限标识

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "has_permission": true
  },
  "message": "success"
}
```

### 3.3 获取用户角色
- **URL**: `GET /api/v1/authorization/users/{user_id}/roles`
- **描述**: 获取指定用户的角色列表
- **认证**: 需要认证

**路径参数**:
- `user_id` (int): 用户ID

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "name": "管理员",
      "code": "admin",
      "description": "系统管理员",
      "status": 1,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 3.4 为用户分配角色
- **URL**: `POST /api/v1/authorization/users/{user_id}/roles`
- **描述**: 为指定用户分配角色
- **认证**: 需要认证

**路径参数**:
- `user_id` (int): 用户ID

**请求参数**:
```json
{
  "role_ids": [1, 2, 3]
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.5 移除用户角色
- **URL**: `DELETE /api/v1/authorization/users/{user_id}/roles/{role_id}`
- **描述**: 移除用户的指定角色
- **认证**: 需要认证

**路径参数**:
- `user_id` (int): 用户ID
- `role_id` (int): 角色ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.6 获取用户权限
- **URL**: `GET /api/v1/authorization/users/{user_id}/permissions`
- **描述**: 获取指定用户的所有权限
- **认证**: 需要认证

**路径参数**:
- `user_id` (int): 用户ID

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "系统管理",
      "type": "menu",
      "path": "/system",
      "method": "GET",
      "icon": "system",
      "component": "Layout",
      "permission": "system:view",
      "status": 1,
      "hidden": false,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 3.7 创建角色
- **URL**: `POST /api/v1/authorization/roles`
- **描述**: 创建新角色
- **认证**: 需要认证

**请求参数**:
```json
{
  "name": "测试角色",
  "code": "tester",
  "description": "测试人员角色",
  "status": 1,
  "sort_order": 10
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "创建成功"
}
```

### 3.8 查询角色列表
- **URL**: `GET /api/v1/authorization/roles`
- **描述**: 分页查询角色列表
- **认证**: 需要认证

**查询参数**:
- `name` (string, optional): 角色名称
- `code` (string, optional): 角色编码
- `status` (int, optional): 状态
- `page` (int, optional): 页码，默认1
- `size` (int, optional): 每页大小，默认10

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "name": "管理员",
        "code": "admin",
        "description": "系统管理员",
        "status": 1,
        "sort_order": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "size": 10
  },
  "message": "success"
}
```

### 3.9 更新角色
- **URL**: `PUT /api/v1/authorization/roles/{id}`
- **描述**: 更新角色信息
- **认证**: 需要认证

**路径参数**:
- `id` (int): 角色ID

**请求参数**:
```json
{
  "name": "更新后的角色名",
  "description": "更新后的描述",
  "status": 1,
  "sort_order": 10
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.10 删除角色
- **URL**: `DELETE /api/v1/authorization/roles/{id}`
- **描述**: 删除角色
- **认证**: 需要认证

**路径参数**:
- `id` (int): 角色ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.11 获取角色详情
- **URL**: `GET /api/v1/authorization/roles/{id}`
- **描述**: 根据ID获取角色详情
- **认证**: 需要认证

**路径参数**:
- `id` (int): 角色ID

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "name": "管理员",
    "code": "admin",
    "description": "系统管理员",
    "status": 1,
    "sort_order": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

### 3.12 为角色分配权限
- **URL**: `POST /api/v1/authorization/roles/{id}/permissions`
- **描述**: 为角色分配权限
- **认证**: 需要认证

**路径参数**:
- `id` (int): 角色ID

**请求参数**:
```json
{
  "permission_ids": [1, 2, 3, 4, 5]
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.13 获取角色权限
- **URL**: `GET /api/v1/authorization/roles/{id}/permissions`
- **描述**: 获取角色的所有权限
- **认证**: 需要认证

**路径参数**:
- `id` (int): 角色ID

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "系统管理",
      "type": "menu",
      "path": "/system",
      "method": "GET",
      "icon": "system",
      "component": "Layout",
      "permission": "system:view",
      "status": 1,
      "hidden": false,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 3.14 创建权限
- **URL**: `POST /api/v1/authorization/permissions`
- **描述**: 创建新权限
- **认证**: 需要认证

**请求参数**:
```json
{
  "parent_id": 0,
  "name": "用户管理",
  "type": "menu",
  "path": "/system/user",
  "method": "GET",
  "icon": "user",
  "component": "system/user/index",
  "permission": "system:user:view",
  "status": 1,
  "hidden": false,
  "sort_order": 1
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "创建成功"
}
```

### 3.15 查询权限列表
- **URL**: `GET /api/v1/authorization/permissions`
- **描述**: 分页查询权限列表
- **认证**: 需要认证

**查询参数**:
- `name` (string, optional): 权限名称
- `type` (string, optional): 权限类型
- `status` (int, optional): 状态
- `parent_id` (int, optional): 父权限ID
- `page` (int, optional): 页码，默认1
- `size` (int, optional): 每页大小，默认10

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "parent_id": 0,
        "name": "系统管理",
        "type": "menu",
        "path": "/system",
        "method": "GET",
        "icon": "system",
        "component": "Layout",
        "permission": "system:view",
        "status": 1,
        "hidden": false,
        "sort_order": 1,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "size": 10
  },
  "message": "success"
}
```

### 3.16 获取权限树
- **URL**: `GET /api/v1/authorization/permissions/tree`
- **描述**: 获取权限树结构
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "系统管理",
      "type": "menu",
      "path": "/system",
      "method": "GET",
      "icon": "system",
      "component": "Layout",
      "permission": "system:view",
      "status": 1,
      "hidden": false,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "children": [
        {
          "id": 2,
          "parent_id": 1,
          "name": "用户管理",
          "type": "menu",
          "path": "/system/user",
          "method": "GET",
          "icon": "user",
          "component": "system/user/index",
          "permission": "system:user:view",
          "status": 1,
          "hidden": false,
          "sort_order": 1,
          "created_at": "2024-01-01T00:00:00Z",
          "updated_at": "2024-01-01T00:00:00Z"
        }
      ]
    }
  ],
  "message": "success"
}
```

### 3.17 更新权限
- **URL**: `PUT /api/v1/authorization/permissions/detail/{id}`
- **描述**: 更新权限信息
- **认证**: 需要认证

**路径参数**:
- `id` (int): 权限ID

**请求参数**:
```json
{
  "name": "更新后的权限名",
  "path": "/updated/path",
  "icon": "updated-icon",
  "status": 1,
  "sort_order": 10
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.18 删除权限
- **URL**: `DELETE /api/v1/authorization/permissions/detail/{id}`
- **描述**: 删除权限
- **认证**: 需要认证

**路径参数**:
- `id` (int): 权限ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 3.19 获取权限详情
- **URL**: `GET /api/v1/authorization/permissions/detail/{id}`
- **描述**: 根据ID获取权限详情
- **认证**: 需要认证

**路径参数**:
- `id` (int): 权限ID

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "parent_id": 0,
    "name": "系统管理",
    "type": "menu",
    "path": "/system",
    "method": "GET",
    "icon": "system",
    "component": "Layout",
    "permission": "system:view",
    "status": 1,
    "hidden": false,
    "sort_order": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

## 4. 组织管理模块 (Organization)

### 4.1 创建部门
- **URL**: `POST /api/v1/organization/departments`
- **描述**: 创建新部门
- **认证**: 需要认证

**请求参数**:
```json
{
  "parent_id": 0,
  "name": "技术部",
  "code": "tech",
  "description": "技术开发部门",
  "status": 1,
  "sort_order": 1
}
```

**响应数据**:
```json
{
  "code": 201,
  "data": 123,
  "message": "创建成功"
}
```

### 4.2 查询部门列表
- **URL**: `GET /api/v1/organization/departments`
- **描述**: 获取部门列表
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "技术部",
      "code": "tech",
      "description": "技术开发部门",
      "status": 1,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 4.3 获取部门树
- **URL**: `GET /api/v1/organization/departments/tree`
- **描述**: 获取部门树结构
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "技术部",
      "code": "tech",
      "description": "技术开发部门",
      "status": 1,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "children": [
        {
          "id": 2,
          "parent_id": 1,
          "name": "前端组",
          "code": "frontend",
          "description": "前端开发组",
          "status": 1,
          "sort_order": 1,
          "created_at": "2024-01-01T00:00:00Z",
          "updated_at": "2024-01-01T00:00:00Z"
        }
      ]
    }
  ],
  "message": "success"
}
```

### 4.4 获取用户所属部门
- **URL**: `GET /api/v1/organization/departments/user`
- **描述**: 获取当前用户所属的部门
- **认证**: 需要认证

**响应数据**:
```json
{
  "code": 200,
  "data": [
    {
      "id": 1,
      "parent_id": 0,
      "name": "技术部",
      "code": "tech",
      "description": "技术开发部门",
      "status": 1,
      "sort_order": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ],
  "message": "success"
}
```

### 4.5 更新部门
- **URL**: `PUT /api/v1/organization/departments/detail/{id}`
- **描述**: 更新部门信息
- **认证**: 需要认证

**路径参数**:
- `id` (int): 部门ID

**请求参数**:
```json
{
  "name": "更新后的部门名",
  "description": "更新后的描述",
  "status": 1,
  "sort_order": 10
}
```

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 4.6 删除部门
- **URL**: `DELETE /api/v1/organization/departments/detail/{id}`
- **描述**: 删除部门
- **认证**: 需要认证

**路径参数**:
- `id` (int): 部门ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 4.7 获取部门详情
- **URL**: `GET /api/v1/organization/departments/detail/{id}`
- **描述**: 根据ID获取部门详情
- **认证**: 需要认证

**路径参数**:
- `id` (int): 部门ID

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "id": 1,
    "parent_id": 0,
    "name": "技术部",
    "code": "tech",
    "description": "技术开发部门",
    "status": 1,
    "sort_order": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  },
  "message": "success"
}
```

### 4.8 获取部门用户列表
- **URL**: `GET /api/v1/organization/departments/detail/{id}/users`
- **描述**: 获取指定部门的用户列表
- **认证**: 需要认证

**路径参数**:
- `id` (int): 部门ID

**查询参数**:
- `pageNum` (int, optional): 页码，默认1
- `pageSize` (int, optional): 每页数量，默认20
- `username` (string, optional): 用户名
- `realName` (string, optional): 真实姓名

**响应数据**:
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "username": "admin",
        "real_name": "管理员",
        "email": "admin@example.com",
        "mobile": "13800138000",
        "status": 1,
        "is_leader": true,
        "created_at": "2024-01-01T00:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "size": 20
  },
  "message": "success"
}
```

### 4.9 为用户分配部门
- **URL**: `POST /api/v1/organization/departments/users/{userId}/departments/{departmentId}`
- **描述**: 为用户分配部门
- **认证**: 需要认证

**路径参数**:
- `userId` (int): 用户ID
- `departmentId` (int): 部门ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

### 4.10 移除用户部门
- **URL**: `DELETE /api/v1/organization/departments/users/{userId}/departments/{departmentId}`
- **描述**: 移除用户的部门关联
- **认证**: 需要认证

**路径参数**:
- `userId` (int): 用户ID
- `departmentId` (int): 部门ID

**响应数据**:
```json
{
  "code": 200,
  "data": null,
  "message": "success"
}
```

## 5. 健康检查

### 5.1 健康检查
- **URL**: `GET /health`
- **描述**: 系统健康检查
- **认证**: 无需认证

**响应数据**:
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z",
  "service": "devops-platform"
}
```

## 6. 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 201 | 创建成功 |
| 400 | 请求参数错误 |
| 401 | 未认证或认证失败 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 7. 认证说明

### JWT Token 使用

1. 登录成功后，服务器返回 JWT Token
2. 后续请求需要在 Header 中携带 Token：
   ```
   Authorization: Bearer <token>
   ```
3. Token 过期后需要重新登录获取新的 Token

### 权限控制

系统采用 RBAC (基于角色的访问控制) 模型：
- 用户拥有角色
- 角色拥有权限
- 权限控制具体的操作

## 8. 数据模型

### 用户信息 (UserInfo)
```json
{
  "id": 1,
  "username": "admin",
  "nickname": "管理员",
  "email": "admin@example.com",
  "phone": "13800138000",
  "avatar": "http://example.com/avatar.jpg",
  "dept_id": 1,
  "dept_name": "技术部",
  "role_id": 1,
  "status": 1,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### 应用信息 (AppVO)
```json
{
  "id": 1,
  "name": "demo-app",
  "description": "演示应用",
  "creator": 1,
  "status": "active",
  "group_ids": [1, 2],
  "env_count": 3,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### 角色信息 (RoleVO)
```json
{
  "id": 1,
  "name": "管理员",
  "code": "admin",
  "description": "系统管理员",
  "status": 1,
  "sort_order": 1,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}
```

### 权限信息 (PermissionVO)
```json
{
  "id": 1,
  "parent_id": 0,
  "name": "系统管理",
  "type": "menu",
  "path": "/system",
  "method": "GET",
  "icon": "system",
  "component": "Layout",
  "permission": "system:view",
  "status": 1,
  "hidden": false,
  "sort_order": 1,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z",
  "children": []
}
```

### 部门信息 (DepartmentVO)
```json
{
  "id": 1,
  "parent_id": 0,
  "name": "技术部",
  "code": "tech",
  "description": "技术开发部门",
  "status": 1,
  "sort_order": 1,
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z",
  "children": [],
  "users": []
}
```

## 9. 前端对接指南

### 9.1 API 客户端配置

建议在前端项目中创建统一的 API 客户端：

```typescript
// api/index.ts
import axios from 'axios';
import { getToken, removeToken } from '@/utils/auth';

const api = axios.create({
  baseURL: 'http://localhost:8080',
  timeout: 10000,
});

// 请求拦截器
api.interceptors.request.use(
  (config) => {
    const token = getToken();
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
api.interceptors.response.use(
  (response) => {
    return response.data;
  },
  (error) => {
    if (error.response?.status === 401) {
      removeToken();
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);

export default api;
```

### 9.2 API 接口封装示例

```typescript
// api/auth.ts
import api from './index';

export interface LoginRequest {
  username: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  expire_at: string;
}

export interface UserInfo {
  id: number;
  username: string;
  nickname: string;
  email: string;
  phone: string;
  avatar: string;
  dept_id: number;
  dept_name: string;
  role_id: number;
  status: number;
  created_at: string;
  updated_at: string;
}

// 登录
export const login = (data: LoginRequest) => {
  return api.post<LoginResponse>('/api/v1/auth/login', data);
};

// 获取用户信息
export const getUserInfo = () => {
  return api.get<UserInfo>('/api/v1/auth/me');
};

// 登出
export const logout = () => {
  return api.post('/api/v1/auth/logout');
};
```

```typescript
// api/application.ts
import api from './index';

export interface AppQuery {
  name?: string;
  status?: string;
  group_id?: number;
  page?: number;
  size?: number;
}

export interface AppVO {
  id: number;
  name: string;
  description: string;
  creator: number;
  status: string;
  group_ids: number[];
  env_count: number;
  created_at: string;
  updated_at: string;
}

export interface PageResult<T> {
  list: T[];
  total: number;
  page: number;
  size: number;
}

// 查询应用列表
export const getAppList = (params: AppQuery) => {
  return api.get<PageResult<AppVO>>('/api/v1/apps', { params });
};

// 创建应用
export const createApp = (data: any) => {
  return api.post('/api/v1/apps', data);
};

// 获取应用详情
export const getAppDetail = (id: number) => {
  return api.get<AppVO>(`/api/v1/apps/${id}`);
};
```

### 9.3 状态管理集成 (Pinia)

```typescript
// stores/auth.ts
import { defineStore } from 'pinia';
import { login, getUserInfo, logout } from '@/api/auth';
import { setToken, removeToken } from '@/utils/auth';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: '',
    userInfo: null as UserInfo | null,
  }),

  actions: {
    async login(loginData: LoginRequest) {
      try {
        const response = await login(loginData);
        this.token = response.data.token;
        setToken(response.data.token);
        return response;
      } catch (error) {
        throw error;
      }
    },

    async getUserInfo() {
      try {
        const response = await getUserInfo();
        this.userInfo = response.data;
        return response;
      } catch (error) {
        throw error;
      }
    },

    async logout() {
      try {
        await logout();
      } finally {
        this.token = '';
        this.userInfo = null;
        removeToken();
      }
    },
  },
});
```

### 9.4 路由守卫

```typescript
// router/guards.ts
import { useAuthStore } from '@/stores/auth';

export function setupRouterGuards(router: Router) {
  router.beforeEach(async (to, from, next) => {
    const authStore = useAuthStore();
    
    // 白名单路由
    const whiteList = ['/login', '/register'];
    
    if (whiteList.includes(to.path)) {
      next();
      return;
    }
    
    // 检查是否有token
    if (!authStore.token) {
      next('/login');
      return;
    }
    
    // 获取用户信息
    if (!authStore.userInfo) {
      try {
        await authStore.getUserInfo();
        next();
      } catch (error) {
        next('/login');
      }
    } else {
      next();
    }
  });
}
```

### 9.5 错误处理

```typescript
// utils/error.ts
import { Message } from '@arco-design/web-vue';

export function handleApiError(error: any) {
  if (error.response) {
    const { status, data } = error.response;
    
    switch (status) {
      case 400:
        Message.error(data.message || '请求参数错误');
        break;
      case 401:
        Message.error('登录已过期，请重新登录');
        break;
      case 403:
        Message.error('权限不足');
        break;
      case 404:
        Message.error('请求的资源不存在');
        break;
      case 500:
        Message.error('服务器内部错误');
        break;
      default:
        Message.error(data.message || '请求失败');
    }
  } else {
    Message.error('网络错误，请检查网络连接');
  }
}
```

## 10. 注意事项

1. **认证Token**: 所有需要认证的接口都必须在请求头中携带 `Authorization: Bearer <token>`
2. **分页参数**: 分页查询的 `page` 从 1 开始，`size` 默认为 10
3. **时间格式**: 所有时间字段均为 ISO 8601 格式 (YYYY-MM-DDTHH:mm:ssZ)
4. **状态码**: 请根据 HTTP 状态码和响应中的 `code` 字段进行错误处理
5. **权限控制**: 部分接口需要特定权限，请确保用户有相应的权限
6. **数据验证**: 请在前端进行必要的数据验证，避免无效请求

---

**文档版本**: v1.0  
**更新时间**: 2024-01-01  
**维护人员**: DevOps Platform Team