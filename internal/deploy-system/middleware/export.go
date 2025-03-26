package middleware

import (
	"devops-platform/internal/deploy-system/middleware/internal/authentication/jwt"
)

const BeanAuthenticationJWT = jwt.BeanAuthenticationJWT

// JWTAuth 导出JWT认证中间件
var JWTAuth = jwt.JWTAuth

// OptionalJWTAuth 导出可选的JWT认证中间件
var OptionalJWTAuth = jwt.OptionalJWTAuth
