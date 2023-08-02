package kubernetes

import (
	"devops-platform/internal/deploy-system/client/kubernetes/internal/domain"
	"devops-platform/internal/deploy-system/client/kubernetes/internal/service"
)

var NewConfig = domain.NewConfig

type Config = domain.KubernetesConfig

type Service = service.Service

func NewService(cfg Config) (*Service, error) {
	return service.New(cfg)
}
