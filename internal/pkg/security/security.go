package security

import (
	"context"
	"devops-platform/internal/pkg/common"
	"devops-platform/pkg/types"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// ----------------- 用户上下文相关定义 -----------------

// UserContext 用户上下文信息
type UserContext struct {
	UserID      types.Long // 用户ID
	Username    string     // 用户名
	RealName    string     // 真实姓名
	Email       string     // 电子邮件
	Phone       string     // 电话号码
	DeptID      types.Long // 部门ID
	Roles       []string   // 角色列表
	Permissions []string   // 权限列表
	TokenString string     // 令牌字符串
	TokenInfo   *TokenInfo // 令牌信息
	IP          string     // 客户端IP
	UserAgent   string     // 用户代理
}

// TokenInfo 令牌信息
type TokenInfo struct {
	Token     string     `json:"token"`      // JWT令牌
	ExpireAt  time.Time  `json:"expire_at"`  // 过期时间
	UserID    types.Long `json:"user_id"`    // 用户ID
	Username  string     `json:"username"`   // 用户名
	RealName  string     `json:"real_name"`  // 真实姓名
	DeptID    types.Long `json:"dept_id"`    // 部门ID
	Role      int        `json:"role"`       // 角色
	LoginTime time.Time  `json:"login_time"` // 登录时间
}

// AuthUser 认证用户信息
type AuthUser struct {
	ID       types.Long `json:"id"`
	Username string     `json:"username"`
	Name     string     `json:"name"`
	Role     int        `json:"role"`
	DeptID   types.Long `json:"dept_id"`
	Token    string     `json:"token"`
}

// 上下文key类型
type contextKey string

// 用户上下文在Context中的Key
const userContextKey = contextKey("user_context")

// 在上下文中设置用户信息
func SetUserContext(ctx context.Context, userContext *UserContext) context.Context {
	if ctx == nil {
		return nil
	}

	if ginCtx, ok := ctx.(*gin.Context); ok {
		// 如果是gin上下文，直接设置值
		ginCtx.Set(string(userContextKey), userContext)
		return ginCtx
	}

	// 标准上下文
	return context.WithValue(ctx, userContextKey, userContext)
}

// 从上下文中获取用户信息
func GetUserContext(ctx context.Context) *UserContext {
	if ctx == nil {
		return nil
	}

	if ginCtx, ok := ctx.(*gin.Context); ok {
		// 如果是gin上下文，从其中获取
		if val, exists := ginCtx.Get(string(userContextKey)); exists {
			if uc, ok := val.(*UserContext); ok {
				return uc
			}
		}
		return nil
	}

	// 标准上下文中获取
	if val := ctx.Value(userContextKey); val != nil {
		if uc, ok := val.(*UserContext); ok {
			return uc
		}
	}
	return nil
}

// CheckPermission 检查用户是否具有指定权限
func CheckPermission(ctx context.Context, permission string) bool {
	userContext := GetUserContext(ctx)
	if userContext == nil {
		return false
	}

	for _, p := range userContext.Permissions {
		if p == permission || p == "*" {
			return true
		}
	}

	return false
}

// RequirePermission Gin中间件，要求用户具有指定权限
func RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userContext := GetUserContext(c)
		if userContext == nil {
			common.ResponseUnauthorized(c, "用户未登录")
			c.Abort()
			return
		}

		if !CheckPermission(c, permission) {
			common.ResponseForbidden(c, fmt.Sprintf("缺少权限: %s", permission))
			c.Abort()
			return
		}

		c.Next()
	}
}

// RequireAuthenticated Gin中间件，要求用户已登录
func RequireAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		userContext := GetUserContext(c)
		if userContext == nil {
			common.ResponseUnauthorized(c, "用户未登录")
			c.Abort()
			return
		}

		c.Next()
	}
}
