package init

import (
	_ "devops-platform/internal/common/config/init"
	_ "devops-platform/internal/common/database/init"
	_ "devops-platform/internal/common/log/init"
	_ "devops-platform/internal/common/swagger/init"
	_ "devops-platform/internal/common/web/init"
	_ "devops-platform/internal/deploy-system/rbac/init"
)
