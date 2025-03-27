# 组织模块 API 测试

本目录包含对组织模块 (`organization`) API 的测试JSON文件，用于Postman等工具进行接口测试。

## 测试文件说明

- `org_create_department.json` - 创建部门接口请求体
- `org_create_sub_department.json` - 创建子部门接口请求体
- `org_update_department.json` - 更新部门接口请求体
- `org_query_department.json` - 部门查询接口参数
- `org_query_user.json` - 用户查询接口参数
- `org_assign_department.json` - 为用户分配部门接口请求体

## API接口列表

| 方法 | 路径 | 描述 | 对应测试文件 |
|------|------|------|------------|
| POST | `/api/v1/organization/departments` | 创建部门 | org_create_department.json |
| GET | `/api/v1/organization/departments` | 获取部门列表 | org_query_department.json (参数) |
| GET | `/api/v1/organization/departments/tree` | 获取部门树结构 | - |
| GET | `/api/v1/organization/departments/user` | 获取用户所属部门 | - |
| GET | `/api/v1/organization/departments/detail/:id` | 获取部门详情 | - |
| PUT | `/api/v1/organization/departments/detail/:id` | 更新部门 | org_update_department.json |
| DELETE | `/api/v1/organization/departments/detail/:id` | 删除部门 | - |
| GET | `/api/v1/organization/departments/detail/:id/users` | 获取部门用户列表 | org_query_user.json (参数) |
| POST | `/api/v1/organization/departments/users/:userId/departments/:departmentId` | 为用户分配部门 | - |
| DELETE | `/api/v1/organization/departments/users/:userId/departments/:departmentId` | 移除用户部门 | - |

## 使用方法

**注意**：组织模块的所有接口都需要认证，请先通过认证模块登录获取token。

### 创建部门接口

**请求路径**：`POST /api/v1/organization/departments`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "parent_id": 0,
  "name": "研发部",
  "code": "DEV",
  "description": "负责产品研发",
  "status": 1,
  "sort_order": 1
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": 1,
  "message": "success"
}
```

### 更新部门接口

**请求路径**：`PUT /api/v1/organization/departments/detail/1`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "parent_id": 1,
  "name": "后端研发组",
  "code": "DEV-BACKEND",
  "description": "负责后端系统研发",
  "status": 1,
  "sort": 1
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": "更新成功",
  "message": "success"
}
```

### 获取部门列表接口

**请求路径**：`GET /api/v1/organization/departments?name=&code=&status=1&parent_id=0&pageNum=1&pageSize=10`

**请求头**：
```
Authorization: Bearer {{token}}
```

**响应示例**：
```json
{
  "code": 200,
  "data": {
    "list": [
      {
        "id": 1,
        "parent_id": 0,
        "name": "研发部",
        "code": "DEV",
        "description": "负责产品研发",
        "status": 1,
        "sort_order": 1,
        "created_at": "2023-04-01T12:00:00Z",
        "updated_at": "2023-04-01T12:00:00Z"
      }
    ],
    "total": 1,
    "page": 1,
    "size": 10
  },
  "message": "success"
}
```

### 为用户分配部门接口

**请求路径**：`POST /api/v1/organization/departments/users/1/departments/1?isLeader=true`

**请求头**：
```
Authorization: Bearer {{token}}
```

**响应示例**：
```json
{
  "code": 200,
  "data": "分配成功",
  "message": "success"
}
``` 