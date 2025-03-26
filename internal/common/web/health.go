package web

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status    string    `json:"status"`
	Service   string    `json:"service"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version,omitempty"`
}

// HealthCheck 处理健康检查请求
// 返回服务状态信息，用于监控和服务发现
func HealthCheck(c *gin.Context, serviceName string) {
	response := HealthResponse{
		Status:    "UP",
		Service:   serviceName,
		Timestamp: time.Now(),
		Version:   "1.0.0", // 这里可以从配置或环境变量获取
	}

	c.JSON(http.StatusOK, response)
}
