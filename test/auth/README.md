# 认证模块 API 测试

本目录包含对认证模块 (`auth`) API 的测试JSON文件，用于Postman等工具进行接口测试。

## 测试文件说明

- `auth_login.json` - 用户登录接口请求体
- `auth_register.json` - 用户注册接口请求体
- `auth_change_password.json` - 修改密码接口请求体

## API接口列表

| 方法 | 路径 | 描述 | 认证要求 | 对应测试文件 |
|------|------|------|----------|------------|
| POST | `/auth/login` | 用户登录 | 否 | auth_login.json |
| POST | `/auth/register` | 用户注册 | 否 | auth_register.json |
| GET | `/auth/health` | 健康检查 | 否 | - |
| POST | `/auth/logout` | 用户登出 | 是 | - |
| GET | `/auth/me` | 获取当前用户信息 | 是 | - |
| POST | `/auth/change-password` | 修改密码 | 是 | auth_change_password.json |

## 使用方法

### 登录接口

**请求路径**：`POST /auth/login`

**请求体**：
```json
{
  "username": "admin",
  "password": "123456"
}
```

**响应示例**：
```json
{
  "code": 200,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expire": "2023-04-01T12:00:00Z"
  },
  "message": "登录成功"
}
```

### 注册接口

**请求路径**：`POST /auth/register`

**请求体**：
```json
{
  "username": "testuser",
  "password": "Test@123",
  "name": "测试用户",
  "mobile": "13800138000",
  "email": "test@example.com",
  "avatar": "https://www.dnsjia.com/luban/img/head.png",
  "dept_id": 1,
  "role": 2
}
```

**响应示例**：
```json
{
  "code": 201,
  "data": 1,
  "message": "注册成功"
}
```

### 修改密码接口

**请求路径**：`POST /auth/change-password`

**请求头**：
```
Authorization: Bearer {{token}}
```

**请求体**：
```json
{
  "old_password": "123456",
  "new_password": "Abcd@1234"
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

## 测试流程

1. 先调用登录接口获取token
2. 将token设置为环境变量
3. 使用token调用其他需要认证的接口 