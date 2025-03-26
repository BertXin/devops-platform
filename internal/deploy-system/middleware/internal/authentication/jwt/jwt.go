package jwt

import (
	"devops-platform/internal/common/web"
	"devops-platform/internal/pkg/common"
	"devops-platform/internal/pkg/security"
	"devops-platform/pkg/beans"
	"devops-platform/pkg/common/jwt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const BeanAuthenticationJWT = "BeanAuthenticationJWT"

// JWT 注册JWT认证中间件
func JWT(getBean func(string) interface{}) {
	router := getBean(web.BeanGinEngine)
	if router == nil {
		logrus.Errorf("初始化时获取[%s]失败，未找到Gin引擎", web.BeanGinEngine)
		return
	}

	// 检查类型，但不直接使用这个变量
	_, ok := router.(gin.IRoutes)
	if !ok {
		logrus.Errorf("初始化时获取[%s]失败，类型不匹配", web.BeanGinEngine)
		return
	}

	// 注册JWT认证中间件函数
	beans.Register(BeanAuthenticationJWT, JWTAuth)

	logrus.Infof("JWT认证中间件初始化成功")
}

// JWTAuth JWT认证中间件
// 验证请求头中的Authorization字段是否包含有效的JWT令牌
// 如果验证成功，将用户信息存储到上下文中
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.ResponseUnauthorized(c, "未提供认证信息")
			c.Abort()
			return
		}

		// 提取Bearer令牌
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			common.ResponseUnauthorized(c, "认证格式错误")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)

		// 解析令牌
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			common.ResponseUnauthorized(c, "无效的认证令牌")
			c.Abort()
			return
		}

		// 检查令牌是否过期
		expTime := time.Unix(claims.ExpiresAt.Unix(), 0)
		if time.Now().After(expTime) {
			common.ResponseUnauthorized(c, "认证令牌已过期")
			c.Abort()
			return
		}

		// 创建用户上下文
		userContext := &security.UserContext{
			UserID:      claims.UserID,
			Username:    claims.Name,
			RealName:    claims.Name,
			TokenString: tokenString,
			TokenInfo: &security.TokenInfo{
				Token:     tokenString,
				ExpireAt:  expTime,
				UserID:    claims.UserID,
				Username:  claims.Name,
				RealName:  claims.Name,
				Role:      int(claims.Role),
				LoginTime: time.Unix(claims.IssuedAt.Unix(), 0),
			},
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
		}

		// 使用web包存储用户上下文，确保与控制器使用相同的机制
		web.SetCurrentUser(c, userContext)

		c.Next()
	}
}

// OptionalJWTAuth 可选的JWT验证
// 与JWTAuth类似，但如果没有提供令牌或令牌无效，仍然允许请求继续处理
func OptionalJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// 提取Bearer令牌
		const prefix = "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)

		// 解析令牌
		claims, err := jwt.ParseToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		// 检查令牌是否过期
		expTime := time.Unix(claims.ExpiresAt.Unix(), 0)
		if time.Now().After(expTime) {
			c.Next()
			return
		}

		// 创建用户上下文
		userContext := &security.UserContext{
			UserID:      claims.UserID,
			Username:    claims.Name,
			RealName:    claims.Name,
			TokenString: tokenString,
			TokenInfo: &security.TokenInfo{
				Token:     tokenString,
				ExpireAt:  expTime,
				UserID:    claims.UserID,
				Username:  claims.Name,
				RealName:  claims.Name,
				Role:      int(claims.Role),
				LoginTime: time.Unix(claims.IssuedAt.Unix(), 0),
			},
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
		}

		// 使用web包存储用户上下文，确保与控制器使用相同的机制
		web.SetCurrentUser(c, userContext)

		c.Next()
	}
}
