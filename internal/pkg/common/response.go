package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

// 内部方法：创建并返回错误响应
func newErrorResponse(status int, errorType string, message string, requestID string) ErrorResponse {
	// 如果没有提供requestID，生成一个新的
	if requestID == "" {
		requestID = uuid.New().String()
	}

	return ErrorResponse{
		Code:      status,
		Error:     errorType,
		Message:   message,
		RequestID: requestID,
	}
}

// ResponseBadRequest 请求参数错误
func ResponseBadRequest(ctx *gin.Context, message string, requestID ...string) {
	rid := ""
	if len(requestID) > 0 {
		rid = requestID[0]
	}

	res := newErrorResponse(http.StatusBadRequest, "BadRequest", message, rid)
	ctx.JSON(http.StatusBadRequest, res)
}

// ResponseUnauthorized 未授权
func ResponseUnauthorized(ctx *gin.Context, message string, requestID ...string) {
	rid := ""
	if len(requestID) > 0 {
		rid = requestID[0]
	}

	res := newErrorResponse(http.StatusUnauthorized, "Unauthorized", message, rid)
	ctx.JSON(http.StatusUnauthorized, res)
}

// ResponseForbidden 无权访问
func ResponseForbidden(ctx *gin.Context, message string, requestID ...string) {
	rid := ""
	if len(requestID) > 0 {
		rid = requestID[0]
	}

	res := newErrorResponse(http.StatusForbidden, "Forbidden", message, rid)
	ctx.JSON(http.StatusForbidden, res)
}

// ResponseNotFound 资源不存在
func ResponseNotFound(ctx *gin.Context, message string, requestID ...string) {
	rid := ""
	if len(requestID) > 0 {
		rid = requestID[0]
	}

	res := newErrorResponse(http.StatusNotFound, "NotFound", message, rid)
	ctx.JSON(http.StatusNotFound, res)
}

// ResponseInternalError 系统内部错误
func ResponseInternalError(ctx *gin.Context, message string, err error, requestID ...string) {
	rid := ""
	if len(requestID) > 0 {
		rid = requestID[0]
	}

	res := newErrorResponse(http.StatusInternalServerError, "InternalServerError", message, rid)

	// 记录错误日志
	logrus.WithError(err).
		WithField("request_id", res.RequestID).
		WithField("path", ctx.Request.URL.Path).
		WithField("method", ctx.Request.Method).
		Error(message)

	ctx.JSON(http.StatusInternalServerError, res)
}

// ResponseError 通用错误处理
func ResponseError(ctx *gin.Context, err error) {
	// 生成请求ID
	requestID := uuid.New().String()

	// 如果是自定义错误，按照错误类型处理
	if commonErr, ok := err.(*Error); ok {
		// 添加日志记录
		logrus.WithError(err).
			WithField("request_id", requestID).
			WithField("error_type", commonErr.Type).
			WithField("path", ctx.Request.URL.Path).
			WithField("method", ctx.Request.Method).
			Error(commonErr.Message)

		switch commonErr.Type {
		case ErrorTypeRequestParam:
			ResponseBadRequest(ctx, commonErr.Message, requestID)
		case ErrorTypeUnauthorized:
			ResponseUnauthorized(ctx, commonErr.Message, requestID)
		case ErrorTypeForbidden:
			ResponseForbidden(ctx, commonErr.Message, requestID)
		case ErrorTypeNotFound:
			ResponseNotFound(ctx, commonErr.Message, requestID)
		default:
			// 内部错误及其他类型使用ResponseInternalError
			ResponseInternalError(ctx, commonErr.Message, commonErr.Cause, requestID)
		}
		return
	}

	// 其他错误当作内部错误处理
	logrus.WithError(err).
		WithField("request_id", requestID).
		WithField("path", ctx.Request.URL.Path).
		WithField("method", ctx.Request.Method).
		Error(err.Error())

	ResponseInternalError(ctx, err.Error(), err, requestID)
}
