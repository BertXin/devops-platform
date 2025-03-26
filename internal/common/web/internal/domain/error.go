package domain

import (
	"devops-platform/internal/pkg/common"
	"errors"
)

type Error interface {
	GetCode() int
	GetStatus() int
	Error() string
}

func NewUnauthorizedError(message string) error {
	return common.UnauthorizedError(message, errors.New(message))
}
