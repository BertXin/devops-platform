package domain

import (
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
)

type ImageRepository struct {
	module.Module
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"` //简介
}

type imageRepository struct {
	module.Module
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"` //简介
}

type CreateImageRepositoryCommand struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Description string `json:"description"` //简介
}

type ModifyImageRepositoryCommand struct {
	ID          types.Long `json:"-"`
	Name        string     `json:"name"`
	Address     string     `json:"address"`
	Description string     `json:"description"` //简介
}

func (command *CreateImageRepositoryCommand) ToImageRepository() (*ImageRepository, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}
	return &ImageRepository{
		Name:        command.Name,
		Address:     command.Address,
		Description: command.Description,
	}, nil
}

func (command *CreateImageRepositoryCommand) Validate() error {

	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("仓库名称不能为空"))
	}

	command.Address = strings.TrimSpace(command.Address)
	if command.Address == "" {
		return common.RequestParamError("", errors.New("仓库地址不能为空"))
	}
	command.Description = strings.TrimSpace(command.Description)
	return nil
}
