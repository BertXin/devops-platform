package common

import (
	"fmt"
)

// 错误类型常量
const (
	// ErrorTypeInternal 内部错误
	ErrorTypeInternal = "InternalError"
	// ErrorTypeRequestParam 请求参数错误
	ErrorTypeRequestParam = "RequestParamError"
	// ErrorTypeUnauthorized 未认证
	ErrorTypeUnauthorized = "UnauthorizedError"
	// ErrorTypeForbidden 无权限
	ErrorTypeForbidden = "ForbiddenError"
	// ErrorTypeNotFound 资源不存在
	ErrorTypeNotFound = "NotFoundError"
)

// Error 自定义错误类型
type Error struct {
	// Type 错误类型
	Type string
	// Message 错误消息
	Message string
	// Cause 原始错误
	Cause error
}

// Error 实现error接口
func (e *Error) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (cause: %s)", e.Type, e.Message, e.Cause.Error())
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Unwrap 返回原始错误
func (e *Error) Unwrap() error {
	return e.Cause
}

// InternalError 创建内部错误
func InternalError(message string, cause error) *Error {
	return &Error{
		Type:    ErrorTypeInternal,
		Message: message,
		Cause:   cause,
	}
}

// RequestParamError 创建请求参数错误
func RequestParamError(message string, cause error) *Error {
	if message == "" {
		if cause != nil {
			message = cause.Error()
		} else {
			message = "请求参数错误"
		}
	}

	return &Error{
		Type:    ErrorTypeRequestParam,
		Message: message,
		Cause:   cause,
	}
}

// UnauthorizedError 创建未认证错误
func UnauthorizedError(message string, cause error) *Error {
	if message == "" {
		message = "用户未登录或登录已过期"
	}

	return &Error{
		Type:    ErrorTypeUnauthorized,
		Message: message,
		Cause:   cause,
	}
}

// ForbiddenError 创建无权限错误
func ForbiddenError(message string, cause error) *Error {
	if message == "" {
		message = "无权限执行该操作"
	}

	return &Error{
		Type:    ErrorTypeForbidden,
		Message: message,
		Cause:   cause,
	}
}

// NotFoundError 创建资源不存在错误
func NotFoundError(message string, cause error) *Error {
	if message == "" {
		message = "资源不存在"
	}

	return &Error{
		Type:    ErrorTypeNotFound,
		Message: message,
		Cause:   cause,
	}
}
