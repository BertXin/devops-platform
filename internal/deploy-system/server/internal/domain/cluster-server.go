package domain

import (
	"devops-platform/internal/pkg/module"
	"devops-platform/pkg/common"
	"devops-platform/pkg/types"
	"errors"
	"strings"
)

/**
 * Kubernetes 集群
 */
type ClusterServer struct {
	module.Module
	Name        string `json:"name"`
	Config      Config `gorm:"embedded" json:"config"`
	Description string `json:"description"` //简介
}

type ProductLineEnvironment struct {
	module.Module
	ClusterServerID types.Long    `json:"-"`
	ClusterServer   ClusterServer `json:"cluster_server"`
}

/**
 * Kubernetes 集群
 */
type clusterServer struct {
	module.Module
	Name              string     `json:"name"`
	Config            Config     `gorm:"embedded" json:"config"`
	ImageRepositoryID types.Long `json:"-"`
	Description       string     `json:"description"` //简介
}

/**
 * Kubernetes 证书配置
 */
type Config struct {
	Address               string `json:"address"`
	CertificateAuthority  string `json:"-"`
	UserClientCertificate string `json:"-"`
	UserClientKey         string `json:"-"`
}

func (c *Config) GetAddress() string {
	return c.Address
}
func (c *Config) GetCertificateAuthority() string {
	return c.CertificateAuthority
}
func (c *Config) GetUserClientCertificate() string {
	return c.UserClientCertificate
}
func (c *Config) GetUserClientKey() string {
	return c.UserClientKey
}

type CreateClusterServerCommand struct {
	Name                  string `json:"name"`
	Description           string `json:"description"` //简介
	Address               string `json:"address"`
	CertificateAuthority  string `json:"certificate_authority"`
	UserClientCertificate string `json:"user_client_certificate"`
	UserClientKey         string `json:"user_client_key"`
}

func (command *CreateClusterServerCommand) ToClusterServer() (*ClusterServer, error) {
	err := command.Validate()
	if err != nil {
		return nil, err
	}

	return &ClusterServer{
		Name: command.Name,
		Config: Config{
			Address:               command.Address,
			CertificateAuthority:  command.CertificateAuthority,
			UserClientCertificate: command.UserClientCertificate,
			UserClientKey:         command.UserClientKey,
		},
		Description: command.Description,
	}, nil
}

func (command *CreateClusterServerCommand) Validate() error {

	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("名称不能为空"))
	}

	command.Address = strings.TrimSpace(command.Address)
	if command.Address == "" {
		return common.RequestParamError("", errors.New("集群服务器地址不能为空"))
	}

	command.CertificateAuthority = strings.TrimSpace(command.CertificateAuthority)
	if command.CertificateAuthority == "" {
		return common.RequestParamError("", errors.New("服务端证书不能为空"))
	}

	command.UserClientCertificate = strings.TrimSpace(command.UserClientCertificate)
	if command.UserClientCertificate == "" {
		return common.RequestParamError("", errors.New("客户端证书不能为空"))
	}

	command.UserClientKey = strings.TrimSpace(command.UserClientKey)
	if command.UserClientKey == "" {
		return common.RequestParamError("", errors.New("客户端密钥不能为空"))
	}

	command.Description = strings.TrimSpace(command.Description)

	return nil
}

type ModifyCertificateCommand struct {
	ID                    types.Long `json:"-"`
	CertificateAuthority  string     `json:"certificate_authority"`
	UserClientCertificate string     `json:"user_client_certificate"`
	UserClientKey         string     `json:"user_client_key"`
}

func (command *ModifyCertificateCommand) Validate() error {

	command.CertificateAuthority = strings.TrimSpace(command.CertificateAuthority)
	if command.CertificateAuthority == "" {
		return common.RequestParamError("", errors.New("服务端证书不能为空"))
	}

	command.UserClientCertificate = strings.TrimSpace(command.UserClientCertificate)
	if command.UserClientCertificate == "" {
		return common.RequestParamError("", errors.New("客户端证书不能为空"))
	}

	command.UserClientKey = strings.TrimSpace(command.UserClientKey)
	if command.UserClientKey == "" {
		return common.RequestParamError("", errors.New("客户端密钥不能为空"))
	}

	return nil
}

type ModifyInfoCommand struct {
	ID                types.Long `json:"-"`
	Name              string     `json:"name"`
	ImageRepositoryID types.Long `json:"image_repository_id"`
	Description       string     `json:"description"` //简介
	Address           string     `json:"address"`
}

func (command *ModifyInfoCommand) Validate() error {

	command.Address = strings.TrimSpace(command.Address)
	if command.Address == "" {
		return common.RequestParamError("", errors.New("url地址不能为空"))
	}

	command.Name = strings.TrimSpace(command.Name)
	if command.Name == "" {
		return common.RequestParamError("", errors.New("名称不能为空"))
	}

	if command.ImageRepositoryID < 0 {
		return common.RequestNotFoundError("镜像ID必须大于零")
	}
	command.Description = strings.TrimSpace(command.Description)

	return nil
}
