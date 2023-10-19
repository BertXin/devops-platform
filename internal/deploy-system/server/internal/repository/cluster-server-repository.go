package repository

import (
	"context"
	"devops-platform/internal/deploy-system/server/internal/domain"
	"devops-platform/pkg/types"
)

func (r *Repository) GetClusterServerByID(ctx context.Context, ID types.Long) (clusterServer *domain.ClusterServer, err error) {
	clusterServer = new(domain.ClusterServer)
	err = r.DB(ctx).Preload("ImageRepository").Take(clusterServer, ID).Error
	return
}

func (r *Repository) CreateClusterServer(ctx context.Context, ClusterServer *domain.ClusterServer) (err error) {
	err = r.DB(ctx).Omit("ImageRepository").Create(ClusterServer).Error
	return
}
