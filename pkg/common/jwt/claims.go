package jwt

import (
	"devops-platform/pkg/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// RegisteredClaims 是jwt标准声明的别名
type RegisteredClaims = jwt.RegisteredClaims

// NumericDate 是jwt时间类型的别名
type NumericDate = jwt.NumericDate

// Claims JWT令牌的声明结构
type Claims struct {
	UserID           types.Long `json:"user_id"`  // 用户ID
	Username         string     `json:"username"` // 用户名
	Name             string     `json:"name"`     // 真实姓名
	Role             int        `json:"role"`     // 用户角色
	RegisteredClaims            // 内嵌标准的声明
}

// NewNumericDate 创建一个NumericDate
func NewNumericDate(t time.Time) *NumericDate {
	return jwt.NewNumericDate(t)
}
