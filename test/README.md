# 运维平台API测试说明文档

本文档提供了对运维平台API接口进行测试的详细指南，包括测试数据和测试流程。

## 目录结构

```
tests/
├── README.md                           # 本文档
├── auth/                               # 认证模块测试数据
│   ├── auth_register_admin.json        # 管理员注册
│   ├── auth_register_developer.json    # 开发人员注册
│   ├── auth_register_tester.json       # 测试人员注册
│   ├── auth_login.json                 # 用户登录
│   └── auth_change_password.json       # 修改密码
├── authorization/                      # 授权模块测试数据
│   ├── role_platform_admin.json        # 平台管理员角色
│   ├── role_developer.json             # 开发人员角色
│   ├── role_tester.json                # 测试人员角色
│   ├── permission_menu.json            # 菜单权限
│   ├── permission_api.json             # API权限
│   ├── permission_button.json          # 按钮权限
│   ├── role_assign_permissions.json    # 为角色分配权限
│   └── user_assign_roles.json          # 为用户分配角色
└── organization/                       # 组织模块测试数据
    ├── department_tech.json            # 技术部
    ├── department_dev.json             # 开发组
    ├── department_qa.json              # 测试组
    ├── department_ops.json             # 运维组
    ├── user_assign_department.json     # 为用户分配部门
    └── department_update.json          # 更新部门信息
```

## API基础URL

```
http://localhost:8080
```

## 认证令牌

大多数API需要在请求头中包含JWT令牌进行认证：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

## 测试流程

### 1. 前置准备

1. 确保运维平台服务已正常启动
2. 创建Postman环境变量:
   - `base_url`: 您的API基础URL
   - `token`: 用于存储JWT令牌
   - `user_id`: 存储当前用户ID
   - `role_id`: 存储角色ID
   - `dept_id`: 存储部门ID

### 2. 组织结构创建

按顺序执行以下请求:

1. 创建技术部门
   - POST `{{base_url}}/api/v1/organization/departments`
   - Body: `tests/organization/department_tech.json`
   - 保存返回的部门ID: `dept_id = {{response.data}}`

2. 创建开发组
   - POST `{{base_url}}/api/v1/organization/departments`
   - Body: `tests/organization/department_dev.json`

3. 创建测试组
   - POST `{{base_url}}/api/v1/organization/departments`
   - Body: `tests/organization/department_qa.json`

4. 创建运维组
   - POST `{{base_url}}/api/v1/organization/departments`
   - Body: `tests/organization/department_ops.json`

5. 获取部门树结构(验证)
   - GET `{{base_url}}/api/v1/organization/departments/tree`

### 3. 角色权限创建

1. 创建平台管理员角色
   - POST `{{base_url}}/api/v1/authorization/roles`
   - Body: `tests/authorization/role_platform_admin.json`
   - 保存返回的角色ID: `role_id = {{response.data}}`

2. 创建开发人员角色
   - POST `{{base_url}}/api/v1/authorization/roles`
   - Body: `tests/authorization/role_developer.json`

3. 创建测试人员角色
   - POST `{{base_url}}/api/v1/authorization/roles`
   - Body: `tests/authorization/role_tester.json`

4. 创建菜单权限
   - POST `{{base_url}}/api/v1/authorization/permissions`
   - Body: `tests/authorization/permission_menu.json`

5. 创建API权限
   - POST `{{base_url}}/api/v1/authorization/permissions`
   - Body: `tests/authorization/permission_api.json`

6. 创建按钮权限
   - POST `{{base_url}}/api/v1/authorization/permissions`
   - Body: `tests/authorization/permission_button.json`

7. 为角色分配权限
   - POST `{{base_url}}/api/v1/authorization/roles/{{role_id}}/permissions`
   - Body: `tests/authorization/role_assign_permissions.json`

8. 获取权限树(验证)
   - GET `{{base_url}}/api/v1/authorization/permissions/tree`

### 4. 用户管理

1. 注册管理员用户
   - POST `{{base_url}}/auth/register`
   - Body: `tests/auth/auth_register_admin.json`
   - 保存返回的用户ID: `user_id = {{response.data}}`

2. 注册开发人员用户
   - POST `{{base_url}}/auth/register`
   - Body: `tests/auth/auth_register_developer.json`

3. 注册测试人员用户
   - POST `{{base_url}}/auth/register`
   - Body: `tests/auth/auth_register_tester.json`

4. 登录管理员用户
   - POST `{{base_url}}/auth/login`
   - Body: `tests/auth/auth_login.json`
   - 保存返回的令牌: `token = {{response.data.token}}`

5. 为用户分配角色
   - POST `{{base_url}}/api/v1/authorization/users/{{user_id}}/roles`
   - Body: `tests/authorization/user_assign_roles.json`
   - Headers: `Authorization: Bearer {{token}}`

6. 为用户分配部门
   - POST `{{base_url}}/api/v1/organization/departments/users/{{user_id}}/departments/{{dept_id}}`
   - Body: `tests/organization/user_assign_department.json`
   - Headers: `Authorization: Bearer {{token}}`

### 5. 功能验证

1. 获取用户信息
   - GET `{{base_url}}/auth/me`
   - Headers: `Authorization: Bearer {{token}}`

2. 获取用户角色
   - GET `{{base_url}}/api/v1/authorization/users/{{user_id}}/roles`
   - Headers: `Authorization: Bearer {{token}}`

3. 获取用户权限
   - GET `{{base_url}}/api/v1/authorization/users/{{user_id}}/permissions`
   - Headers: `Authorization: Bearer {{token}}`

4. 获取用户菜单
   - GET `{{base_url}}/api/v1/authorization/menus`
   - Headers: `Authorization: Bearer {{token}}`

5. 修改密码
   - POST `{{base_url}}/auth/change-password`
   - Body: `tests/auth/auth_change_password.json`
   - Headers: `Authorization: Bearer {{token}}`

6. 更新部门信息
   - PUT `{{base_url}}/api/v1/organization/departments/detail/{{dept_id}}`
   - Body: `tests/organization/department_update.json`
   - Headers: `Authorization: Bearer {{token}}`

7. 获取部门用户列表
   - GET `{{base_url}}/api/v1/organization/departments/detail/{{dept_id}}/users?page=1&size=10`
   - Headers: `Authorization: Bearer {{token}}`

### 6. 清理测试

1. 用户登出
   - POST `{{base_url}}/auth/logout`
   - Headers: `Authorization: Bearer {{token}}`

## 注意事项

1. 所有请求应设置Content-Type为application/json
2. 大多数请求需要在请求头中加入Authorization: Bearer {token}
3. 错误响应格式：{"code": 错误码, "error": "错误类型", "message": "详细信息", "request_id": "请求ID"}
4. 成功响应格式：{"code": 200, "data": 数据内容, "message": "success"}

## 故障排除

如果遇到测试问题，请检查:

1. 服务是否正常运行
2. 请求URL和方法是否正确
3. 请求头中的认证令牌是否有效
4. 请求体JSON格式是否正确

## Postman集合

您可以将这些测试数据和流程导入到Postman中，创建一个完整的测试集合，以便更方便地进行测试。 