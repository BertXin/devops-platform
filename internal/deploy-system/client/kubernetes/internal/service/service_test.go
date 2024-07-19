package service

import (
	"context"
	"devops-platform/internal/deploy-system/client/kubernetes/internal/domain"
	"fmt"
	"testing"
)

func TestService_GetPodLogs(t *testing.T) {
	//
	s, err := New(getConfig())
	if err != nil {
		t.Fatal(err)
	}
	logs, err := s.GetPodLogs(context.TODO(), "niffty-apps", "auth-deploy-bbb7d4c94-8h8zh", "auth")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(logs)
}

func getConfig() domain.KubernetesConfig {
	return getConfig()
}
