package domain

import (
	"devops-platform/internal/pkg/common"
	"errors"
	"net/http"
)

type Error interface {
	GetCode() int
	GetStatus() int
	Error() string
}

func NewUnauthorizedError(message string) error {
	return common.Unauthorized(http.StatusUnauthorized, errors.New(message))
}
