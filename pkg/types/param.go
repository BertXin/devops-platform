package types

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetIntParam 从请求中获取整型参数
func GetIntParam(ctx *gin.Context, key string, defaultValue int) int {
	valueStr := ctx.DefaultQuery(key, strconv.Itoa(defaultValue))
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}
