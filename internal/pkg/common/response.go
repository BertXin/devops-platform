package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Response 基础响应结构
type Response struct {
	// Code 响应码
	Code int `json:"code"`
	// Data 数据
	Data interface{} `json:"data"`
	// Message 消息
	Message string `json:"message"`
}

// ErrorResponse 错误响应结构
type ErrorResponse struct {
	// Code 响应码
	Code int `json:"code"`
	// Error 错误类型
	Error string `json:"error"`
	// Message 错误消息
	Message string `json:"message"`
	// RequestID 请求ID
	RequestID string `json:"request_id"`
}

// ResponseSuccess 成功响应
func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Data:    data,
		Message: "success",
	})
}

// ResponseSuccessWithPage 带分页的成功响应
func ResponseSuccessWithPage(ctx *gin.Context, list interface{}, total int64) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: PageResult{
			List:  list,
			Total: total,
		},
		Message: "success",
	})
}

// ResponseSuccessWithPageExt 带扩展分页参数的成功响应
func ResponseSuccessWithPageExt(ctx *gin.Context, list interface{}, total int64, page, size int) {
	ctx.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: PageResult{
			List:  list,
			Total: total,
			Page:  page,
			Size:  size,
		},
		Message: "success",
	})
}

// ResponseBadRequest 请求参数错误
func ResponseBadRequest(ctx *gin.Context, message string) {
	requestID := uuid.New().String()
	ctx.JSON(http.StatusBadRequest, ErrorResponse{
		Code:      http.StatusBadRequest,
		Error:     "BadRequest",
		Message:   message,
		RequestID: requestID,
	})
}

// ResponseUnauthorized 未授权
func ResponseUnauthorized(ctx *gin.Context, message string) {
	requestID := uuid.New().String()
	ctx.JSON(http.StatusUnauthorized, ErrorResponse{
		Code:      http.StatusUnauthorized,
		Error:     "Unauthorized",
		Message:   message,
		RequestID: requestID,
	})
}

// ResponseForbidden 无权访问
func ResponseForbidden(ctx *gin.Context, message string) {
	requestID := uuid.New().String()
	ctx.JSON(http.StatusForbidden, ErrorResponse{
		Code:      http.StatusForbidden,
		Error:     "Forbidden",
		Message:   message,
		RequestID: requestID,
	})
}

// ResponseNotFound 资源不存在
func ResponseNotFound(ctx *gin.Context, message string) {
	requestID := uuid.New().String()
	ctx.JSON(http.StatusNotFound, ErrorResponse{
		Code:      http.StatusNotFound,
		Error:     "NotFound",
		Message:   message,
		RequestID: requestID,
	})
}

// ResponseInternalError 系统内部错误
func ResponseInternalError(ctx *gin.Context, message string, err error) {
	requestID := uuid.New().String()

	// 记录错误日志
	// log.GetLogger().WithError(err).WithField("request_id", requestID).Error(message)

	ctx.JSON(http.StatusInternalServerError, ErrorResponse{
		Code:      http.StatusInternalServerError,
		Error:     "InternalServerError",
		Message:   message,
		RequestID: requestID,
	})
}

// ResponseError 通用错误处理
func ResponseError(ctx *gin.Context, err error) {
	// 如果是自定义错误，按照错误类型处理
	if commonErr, ok := err.(*Error); ok {
		switch commonErr.Type {
		case ErrorTypeRequestParam:
			ResponseBadRequest(ctx, commonErr.Message)
		case ErrorTypeUnauthorized:
			ResponseUnauthorized(ctx, commonErr.Message)
		case ErrorTypeForbidden:
			ResponseForbidden(ctx, commonErr.Message)
		case ErrorTypeNotFound:
			ResponseNotFound(ctx, commonErr.Message)
		default:
			ResponseInternalError(ctx, commonErr.Message, commonErr)
		}
		return
	}

	// 其他错误当作内部错误处理
	ResponseInternalError(ctx, err.Error(), err)
}
