package init

import (
	"devops-platform/internal/deploy-system/client/kubernetes/internal/service"
	"devops-platform/pkg/beans"
)

func init() {
	beans.Register("kubernetes-set-config", service.SetConfig)
}
