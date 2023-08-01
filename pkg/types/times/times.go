package times

import (
	"devops-platform/pkg/types"
	"time"
)

func Now() *types.Time {
	return &types.Time{Time: time.Now()}
}