package init

import (
	_ "devops-platform/internal/deploy-system/client/kubernetes/init"
	_ "devops-platform/internal/deploy-system/login/init"
	_ "devops-platform/internal/deploy-system/server/init"
	_ "devops-platform/internal/deploy-system/user/init"
)
