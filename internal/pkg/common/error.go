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

func Unauthorized(code int, err error) *Error {
	return &Error{
		error:  err,
		code:   code,
		status: 401,
	}
}

func ServiceError(code int, err error) *Error {
	return &Error{
		error:  err,
		code:   code,
		status: 500,
	}
}

func RequestError(code int, err error) *Error {
	return &Error{
		error:  err,
		code:   code,
		status: 400,
	}
}

func RequestParamError(message string, err error) *Error {
	if message != "" {
		err = fmt.Errorf(message+" %v", err)
	}

	return &Error{
		error:  err,
		code:   400,
		status: 400,
	}
}

func NewRequestParamError(message string) *Error {

	return &Error{
		error:  errors.New(message),
		code:   400,
		status: 400,
	}
}

func NewRequestParamErrorf(message string, a ...interface{}) *Error {

	return &Error{
		error:  fmt.Errorf(message, a...),
		code:   400,
		status: 400,
	}
}

func RequestNotFoundError(err string) *Error {
	return &Error{
		error:  errors.New(err),
		code:   404,
		status: 404,
	}
}

func WarpServiceError(code int, message string, err error) *Error {
	return &Error{
		error:  fmt.Errorf(message+" %v", err),
		code:   code,
		status: 500,
	}
}
func WarpRequestError(code int, message string, err error) *Error {
	return &Error{
		error:  fmt.Errorf(message+" %v", err),
		code:   code,
		status: 400,
	}
}

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
