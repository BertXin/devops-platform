package common

import (
	"errors"
	"fmt"
)

type Error struct {
	error
	code   int
	status int
}

func (e *Error) GetCode() int {
	return e.code
}
func (e *Error) GetStatus() int {
	return e.status
}
func (e *Error) Unwrap() error {
	return e.error
}

func (e *Error) Error() string {
	return e.error.Error()
}

//Unauthorized 封装未授权错误
func Unauthorized(code int, err error) *Error {
	return &Error{
		error:  err,
		code:   code,
		status: 401,
	}
}

//ServiceError 封装服务端错误
func ServiceError(code int, err error) *Error {
	return &Error{
		error:  err,
		code:   code,
		status: 500,
	}
}

//RequestError 封装请求错误
func RequestError(code int, err error) *Error {
	return &Error{
		error:  err,
		code:   code,
		status: 400,
	}
}

//RequestParamError 封装请求参数错误
func RequestParamError(message string, err error) *Error {
	return &Error{
		error:  fmt.Errorf(message+" %v", err),
		code:   400,
		status: 400,
	}
}

//RequestNotFoundError 封装请求未找到错误
func RequestNotFoundError(err string) *Error {
	return &Error{
		error:  errors.New(err),
		code:   404,
		status: 404,
	}
}

// 封装服务端错误并添加消息
func WarpServiceError(code int, message string, err error) *Error {
	return &Error{
		error:  fmt.Errorf(message+" %v", err),
		code:   code,
		status: 500,
	}
}

//WarpRequestError 封装请求错误并添加消息
func WarpRequestError(code int, message string, err error) *Error {
	return &Error{
		error:  fmt.Errorf(message+" %v", err),
		code:   code,
		status: 400,
	}
}

// 封装任意错误
func WarpError(err error) *Error {
	if err == nil {
		return nil
	}
	var e *Error
	if errors.As(err, &e) {
		return e
	} else {
		return ServiceError(500, err)
	}
}
