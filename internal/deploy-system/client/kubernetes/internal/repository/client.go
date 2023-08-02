package repository

import (
	"devops-platform/internal/deploy-system/client/kubernetes/internal/domain"
	"k8s.io/client-go/kubernetes"
)

/**
 * kubernetes client
 */
func NewClient(cfg domain.KubernetesConfig) (*kubernetes.Clientset, error) {
	config, err := domain.NewConfig(cfg)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

/**
 * kubernetes client
 */
func New(cfg domain.KubernetesConfig) (*Repository, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Repository{
		client: client,
	}, nil
}
